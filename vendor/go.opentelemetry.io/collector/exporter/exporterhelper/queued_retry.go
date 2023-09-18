// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package exporterhelper // import "go.opentelemetry.io/collector/exporter/exporterhelper"

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/cenkalti/backoff/v4"
	"go.opencensus.io/metric/metricdata"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer/consumererror"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exporterhelper/internal"
	"go.opentelemetry.io/collector/internal/obsreportconfig/obsmetrics"
)

const defaultQueueSize = 1000

var errSendingQueueIsFull = errors.New("sending_queue is full")

// QueueSettings defines configuration for queueing batches before sending to the consumerSender.
type QueueSettings struct {
	// Enabled indicates whether to not enqueue batches before sending to the consumerSender.
	Enabled bool `mapstructure:"enabled"`
	// NumConsumers is the number of consumers from the queue.
	NumConsumers int `mapstructure:"num_consumers"`
	// QueueSize is the maximum number of batches allowed in queue at a given time.
	QueueSize int `mapstructure:"queue_size"`
	// StorageID if not empty, enables the persistent storage and uses the component specified
	// as a storage extension for the persistent queue
	StorageID *component.ID `mapstructure:"storage"`
}

// NewDefaultQueueSettings returns the default settings for QueueSettings.
func NewDefaultQueueSettings() QueueSettings {
	return QueueSettings{
		Enabled:      true,
		NumConsumers: 10,
		// By default, batches are 8192 spans, for a total of up to 8 million spans in the queue
		// This can be estimated at 1-4 GB worth of maximum memory usage
		// This default is probably still too high, and may be adjusted further down in a future release
		QueueSize: defaultQueueSize,
	}
}

// Validate checks if the QueueSettings configuration is valid
func (qCfg *QueueSettings) Validate() error {
	if !qCfg.Enabled {
		return nil
	}

	if qCfg.QueueSize <= 0 {
		return errors.New("queue size must be positive")
	}

	return nil
}

type queueSender struct {
	baseRequestSender
	fullName         string
	id               component.ID
	signal           component.DataType
	queue            internal.ProducerConsumerQueue
	traceAttribute   attribute.KeyValue
	logger           *zap.Logger
	requeuingEnabled bool
}

func newQueueSender(id component.ID, signal component.DataType, queue internal.ProducerConsumerQueue, logger *zap.Logger) *queueSender {
	return &queueSender{
		fullName:       id.String(),
		id:             id,
		signal:         signal,
		queue:          queue,
		traceAttribute: attribute.String(obsmetrics.ExporterKey, id.String()),
		logger:         logger,
		// TODO: this can be further exposed as a config param rather than relying on a type of queue
		requeuingEnabled: queue != nil && queue.IsPersistent(),
	}
}

func (qs *queueSender) onTemporaryFailure(logger *zap.Logger, req internal.Request, err error) error {
	if !qs.requeuingEnabled || qs.queue == nil {
		logger.Error(
			"Exporting failed. No more retries left. Dropping data.",
			zap.Error(err),
			zap.Int("dropped_items", req.Count()),
		)
		return err
	}

	if qs.queue.Produce(req) {
		logger.Error(
			"Exporting failed. Putting back to the end of the queue.",
			zap.Error(err),
		)
	} else {
		logger.Error(
			"Exporting failed. Queue did not accept requeuing request. Dropping data.",
			zap.Error(err),
			zap.Int("dropped_items", req.Count()),
		)
	}
	return err
}

// start is invoked during service startup.
func (qs *queueSender) start(ctx context.Context, host component.Host, set exporter.CreateSettings) error {
	if qs.queue == nil {
		return nil
	}

	err := qs.queue.Start(ctx, host, internal.QueueSettings{
		CreateSettings: set,
		DataType:       qs.signal,
		Callback: func(item internal.Request) {
			_ = qs.nextSender.send(item)
			item.OnProcessingFinished()
		},
	})
	if err != nil {
		return err
	}

	// Start reporting queue length metric
	err = globalInstruments.queueSize.UpsertEntry(func() int64 {
		return int64(qs.queue.Size())
	}, metricdata.NewLabelValue(qs.fullName))
	if err != nil {
		return fmt.Errorf("failed to create retry queue size metric: %w", err)
	}
	err = globalInstruments.queueCapacity.UpsertEntry(func() int64 {
		return int64(qs.queue.Capacity())
	}, metricdata.NewLabelValue(qs.fullName))
	if err != nil {
		return fmt.Errorf("failed to create retry queue capacity metric: %w", err)
	}

	return nil
}

// shutdown is invoked during service shutdown.
func (qs *queueSender) shutdown() {
	if qs.queue != nil {
		// Cleanup queue metrics reporting
		_ = globalInstruments.queueSize.UpsertEntry(func() int64 {
			return int64(0)
		}, metricdata.NewLabelValue(qs.fullName))

		// Stop the queued sender, this will drain the queue and will call the retry (which is stopped) that will only
		// try once every request.
		qs.queue.Stop()
	}
}

// RetrySettings defines configuration for retrying batches in case of export failure.
// The current supported strategy is exponential backoff.
type RetrySettings struct {
	// Enabled indicates whether to not retry sending batches in case of export failure.
	Enabled bool `mapstructure:"enabled"`
	// InitialInterval the time to wait after the first failure before retrying.
	InitialInterval time.Duration `mapstructure:"initial_interval"`
	// RandomizationFactor is a random factor used to calculate next backoffs
	// Randomized interval = RetryInterval * (1 ± RandomizationFactor)
	RandomizationFactor float64 `mapstructure:"randomization_factor"`
	// Multiplier is the value multiplied by the backoff interval bounds
	Multiplier float64 `mapstructure:"multiplier"`
	// MaxInterval is the upper bound on backoff interval. Once this value is reached the delay between
	// consecutive retries will always be `MaxInterval`.
	MaxInterval time.Duration `mapstructure:"max_interval"`
	// MaxElapsedTime is the maximum amount of time (including retries) spent trying to send a request/batch.
	// Once this value is reached, the data is discarded.
	MaxElapsedTime time.Duration `mapstructure:"max_elapsed_time"`
}

// NewDefaultRetrySettings returns the default settings for RetrySettings.
func NewDefaultRetrySettings() RetrySettings {
	return RetrySettings{
		Enabled:             true,
		InitialInterval:     5 * time.Second,
		RandomizationFactor: backoff.DefaultRandomizationFactor,
		Multiplier:          backoff.DefaultMultiplier,
		MaxInterval:         30 * time.Second,
		MaxElapsedTime:      5 * time.Minute,
	}
}

// send implements the requestSender interface
func (qs *queueSender) send(req internal.Request) error {
	if qs.queue == nil {
		err := qs.nextSender.send(req)
		if err != nil {
			qs.logger.Error(
				"Exporting failed. Dropping data. Try enabling sending_queue to survive temporary failures.",
				zap.Int("dropped_items", req.Count()),
			)
		}
		return err
	}

	// Prevent cancellation and deadline to propagate to the context stored in the queue.
	// The grpc/http based receivers will cancel the request context after this function returns.
	req.SetContext(noCancellationContext{Context: req.Context()})

	span := trace.SpanFromContext(req.Context())
	if !qs.queue.Produce(req) {
		qs.logger.Error(
			"Dropping data because sending_queue is full. Try increasing queue_size.",
			zap.Int("dropped_items", req.Count()),
		)
		span.AddEvent("Dropped item, sending_queue is full.", trace.WithAttributes(qs.traceAttribute))
		return errSendingQueueIsFull
	}

	span.AddEvent("Enqueued item.", trace.WithAttributes(qs.traceAttribute))
	return nil
}

// TODO: Clean this by forcing all exporters to return an internal error type that always include the information about retries.
type throttleRetry struct {
	err   error
	delay time.Duration
}

func (t throttleRetry) Error() string {
	return "Throttle (" + t.delay.String() + "), error: " + t.err.Error()
}

func (t throttleRetry) Unwrap() error {
	return t.err
}

// NewThrottleRetry creates a new throttle retry error.
func NewThrottleRetry(err error, delay time.Duration) error {
	return throttleRetry{
		err:   err,
		delay: delay,
	}
}

type onRequestHandlingFinishedFunc func(*zap.Logger, internal.Request, error) error

type retrySender struct {
	baseRequestSender
	traceAttribute     attribute.KeyValue
	cfg                RetrySettings
	stopCh             chan struct{}
	logger             *zap.Logger
	onTemporaryFailure onRequestHandlingFinishedFunc
}

func newRetrySender(id component.ID, rCfg RetrySettings, logger *zap.Logger, onTemporaryFailure onRequestHandlingFinishedFunc) *retrySender {
	if onTemporaryFailure == nil {
		onTemporaryFailure = func(logger *zap.Logger, req internal.Request, err error) error {
			return err
		}
	}
	return &retrySender{
		traceAttribute:     attribute.String(obsmetrics.ExporterKey, id.String()),
		cfg:                rCfg,
		stopCh:             make(chan struct{}),
		logger:             logger,
		onTemporaryFailure: onTemporaryFailure,
	}
}

func (rs *retrySender) shutdown() {
	close(rs.stopCh)
}

// send implements the requestSender interface
func (rs *retrySender) send(req internal.Request) error {
	if !rs.cfg.Enabled {
		err := rs.nextSender.send(req)
		if err != nil {
			rs.logger.Error(
				"Exporting failed. Try enabling retry_on_failure config option to retry on retryable errors",
				zap.Error(err),
			)
		}
		return err
	}

	// Do not use NewExponentialBackOff since it calls Reset and the code here must
	// call Reset after changing the InitialInterval (this saves an unnecessary call to Now).
	expBackoff := backoff.ExponentialBackOff{
		InitialInterval:     rs.cfg.InitialInterval,
		RandomizationFactor: rs.cfg.RandomizationFactor,
		Multiplier:          rs.cfg.Multiplier,
		MaxInterval:         rs.cfg.MaxInterval,
		MaxElapsedTime:      rs.cfg.MaxElapsedTime,
		Stop:                backoff.Stop,
		Clock:               backoff.SystemClock,
	}
	expBackoff.Reset()
	span := trace.SpanFromContext(req.Context())
	retryNum := int64(0)
	for {
		span.AddEvent(
			"Sending request.",
			trace.WithAttributes(rs.traceAttribute, attribute.Int64("retry_num", retryNum)))

		err := rs.nextSender.send(req)
		if err == nil {
			return nil
		}

		// Immediately drop data on permanent errors.
		if consumererror.IsPermanent(err) {
			rs.logger.Error(
				"Exporting failed. The error is not retryable. Dropping data.",
				zap.Error(err),
				zap.Int("dropped_items", req.Count()),
			)
			return err
		}

		// Give the request a chance to extract signal data to retry if only some data
		// failed to process.
		req = req.OnError(err)

		backoffDelay := expBackoff.NextBackOff()
		if backoffDelay == backoff.Stop {
			// throw away the batch
			err = fmt.Errorf("max elapsed time expired %w", err)
			return rs.onTemporaryFailure(rs.logger, req, err)
		}

		throttleErr := throttleRetry{}
		isThrottle := errors.As(err, &throttleErr)
		if isThrottle {
			backoffDelay = max(backoffDelay, throttleErr.delay)
		}

		backoffDelayStr := backoffDelay.String()
		span.AddEvent(
			"Exporting failed. Will retry the request after interval.",
			trace.WithAttributes(
				rs.traceAttribute,
				attribute.String("interval", backoffDelayStr),
				attribute.String("error", err.Error())))
		rs.logger.Info(
			"Exporting failed. Will retry the request after interval.",
			zap.Error(err),
			zap.String("interval", backoffDelayStr),
		)
		retryNum++

		// back-off, but get interrupted when shutting down or request is cancelled or timed out.
		select {
		case <-req.Context().Done():
			return fmt.Errorf("Request is cancelled or timed out %w", err)
		case <-rs.stopCh:
			return rs.onTemporaryFailure(rs.logger, req, fmt.Errorf("interrupted due to shutdown %w", err))
		case <-time.After(backoffDelay):
		}
	}
}

// max returns the larger of x or y.
func max(x, y time.Duration) time.Duration {
	if x < y {
		return y
	}
	return x
}

type noCancellationContext struct {
	context.Context
}

func (noCancellationContext) Deadline() (deadline time.Time, ok bool) {
	return
}

func (noCancellationContext) Done() <-chan struct{} {
	return nil
}

func (noCancellationContext) Err() error {
	return nil
}
