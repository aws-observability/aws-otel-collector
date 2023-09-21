// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusreceiver/internal"

import (
	"context"
	"regexp"
	"time"

	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/storage"
	"go.opentelemetry.io/collector/consumer"
<<<<<<< HEAD
	"go.opentelemetry.io/collector/featuregate"
=======
>>>>>>> main
	"go.opentelemetry.io/collector/obsreport"
	"go.opentelemetry.io/collector/receiver"
)

// appendable translates Prometheus scraping diffs into OpenTelemetry format.
type appendable struct {
	sink                 consumer.Metrics
	metricAdjuster       MetricsAdjuster
	useStartTimeMetric   bool
<<<<<<< HEAD
=======
	trimSuffixes         bool
>>>>>>> main
	startTimeMetricRegex *regexp.Regexp
	externalLabels       labels.Labels

	settings receiver.CreateSettings
	obsrecv  *obsreport.Receiver
<<<<<<< HEAD
	registry *featuregate.Registry
=======
>>>>>>> main
}

// NewAppendable returns a storage.Appendable instance that emits metrics to the sink.
func NewAppendable(
	sink consumer.Metrics,
	set receiver.CreateSettings,
	gcInterval time.Duration,
	useStartTimeMetric bool,
	startTimeMetricRegex *regexp.Regexp,
	useCreatedMetric bool,
	externalLabels labels.Labels,
<<<<<<< HEAD
	registry *featuregate.Registry) (storage.Appendable, error) {
=======
	trimSuffixes bool) (storage.Appendable, error) {
>>>>>>> main
	var metricAdjuster MetricsAdjuster
	if !useStartTimeMetric {
		metricAdjuster = NewInitialPointAdjuster(set.Logger, gcInterval, useCreatedMetric)
	} else {
		metricAdjuster = NewStartTimeMetricAdjuster(set.Logger, startTimeMetricRegex)
	}

	obsrecv, err := obsreport.NewReceiver(obsreport.ReceiverSettings{ReceiverID: set.ID, Transport: transport, ReceiverCreateSettings: set})
	if err != nil {
		return nil, err
	}

	return &appendable{
		sink:                 sink,
		settings:             set,
		metricAdjuster:       metricAdjuster,
		useStartTimeMetric:   useStartTimeMetric,
		startTimeMetricRegex: startTimeMetricRegex,
		externalLabels:       externalLabels,
		obsrecv:              obsrecv,
<<<<<<< HEAD
		registry:             registry,
=======
		trimSuffixes:         trimSuffixes,
>>>>>>> main
	}, nil
}

func (o *appendable) Appender(ctx context.Context) storage.Appender {
<<<<<<< HEAD
	return newTransaction(ctx, o.metricAdjuster, o.sink, o.externalLabels, o.settings, o.obsrecv, o.registry)
=======
	return newTransaction(ctx, o.metricAdjuster, o.sink, o.externalLabels, o.settings, o.obsrecv, o.trimSuffixes)
>>>>>>> main
}
