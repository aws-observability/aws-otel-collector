// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package componenttest // import "go.opentelemetry.io/collector/component/componenttest"

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel/attribute"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
	"go.uber.org/zap"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configtelemetry"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

const (
	// Names used by the metrics and labels are hard coded here in order to avoid
	// inadvertent changes: at this point changing metric names and labels should
	// be treated as a breaking changing and requires a good justification.
	// Changes to metric names or labels can break alerting, dashboards, etc
	// that are used to monitor the Collector in production deployments.
	// DO NOT SWITCH THE VARIABLES BELOW TO SIMILAR ONES DEFINED ON THE PACKAGE.
	receiverTag  = "receiver"
	scraperTag   = "scraper"
	transportTag = "transport"
	exporterTag  = "exporter"
)

type TestTelemetry struct {
	id           component.ID
	ts           component.TelemetrySettings
	SpanRecorder *tracetest.SpanRecorder
	reader       *sdkmetric.ManualReader
}

// CheckExporterTraces checks that for the current exported values for trace exporter metrics match given values.
func (tts *TestTelemetry) CheckExporterTraces(sentSpans, sendFailedSpans int64) error {
	return checkExporterTraces(tts.reader, tts.id, sentSpans, sendFailedSpans)
}

// CheckExporterMetrics checks that for the current exported values for metrics exporter metrics match given values.
func (tts *TestTelemetry) CheckExporterMetrics(sentMetricsPoints, sendFailedMetricsPoints int64) error {
	return checkExporterMetrics(tts.reader, tts.id, sentMetricsPoints, sendFailedMetricsPoints)
}

func (tts *TestTelemetry) CheckExporterEnqueueFailedMetrics(enqueueFailed int64) error {
	return checkExporterEnqueueFailed(tts.reader, tts.id, "metric_points", enqueueFailed)
}

func (tts *TestTelemetry) CheckExporterEnqueueFailedTraces(enqueueFailed int64) error {
	return checkExporterEnqueueFailed(tts.reader, tts.id, "spans", enqueueFailed)
}

func (tts *TestTelemetry) CheckExporterEnqueueFailedLogs(enqueueFailed int64) error {
	return checkExporterEnqueueFailed(tts.reader, tts.id, "log_records", enqueueFailed)
}

// CheckExporterLogs checks that for the current exported values for logs exporter metrics match given values.
func (tts *TestTelemetry) CheckExporterLogs(sentLogRecords, sendFailedLogRecords int64) error {
	return checkExporterLogs(tts.reader, tts.id, sentLogRecords, sendFailedLogRecords)
}

func (tts *TestTelemetry) CheckExporterMetricGauge(metric string, val int64, extraAttrs ...attribute.KeyValue) error {
	attrs := attributesForExporterMetrics(tts.id, extraAttrs...)
	return checkIntGauge(tts.reader, metric, val, attrs)
}

// CheckReceiverTraces checks that for the current exported values for trace receiver metrics match given values.
func (tts *TestTelemetry) CheckReceiverTraces(protocol string, acceptedSpans, droppedSpans int64) error {
	return checkReceiverTraces(tts.reader, tts.id, protocol, acceptedSpans, droppedSpans)
}

// CheckReceiverLogs checks that for the current exported values for logs receiver metrics match given values.
func (tts *TestTelemetry) CheckReceiverLogs(protocol string, acceptedLogRecords, droppedLogRecords int64) error {
	return checkReceiverLogs(tts.reader, tts.id, protocol, acceptedLogRecords, droppedLogRecords)
}

// CheckReceiverMetrics checks that for the current exported values for metrics receiver metrics match given values.
func (tts *TestTelemetry) CheckReceiverMetrics(protocol string, acceptedMetricPoints, droppedMetricPoints int64) error {
	return checkReceiverMetrics(tts.reader, tts.id, protocol, acceptedMetricPoints, droppedMetricPoints)
}

// CheckScraperMetrics checks that for the current exported values for metrics scraper metrics match given values.
func (tts *TestTelemetry) CheckScraperMetrics(receiver component.ID, scraper component.ID, scrapedMetricPoints, erroredMetricPoints int64) error {
	return checkScraperMetrics(tts.reader, receiver, scraper, scrapedMetricPoints, erroredMetricPoints)
}

// Shutdown unregisters any views and shuts down the SpanRecorder
func (tts *TestTelemetry) Shutdown(ctx context.Context) error {
	return errors.Join(
		tts.ts.TracerProvider.(*sdktrace.TracerProvider).Shutdown(ctx),
		tts.ts.MeterProvider.(*sdkmetric.MeterProvider).Shutdown(ctx))
}

// TelemetrySettings returns the TestTelemetry's TelemetrySettings
func (tts *TestTelemetry) TelemetrySettings() component.TelemetrySettings {
	return tts.ts
}

// SetupTelemetry sets up the testing environment to check the metrics recorded by receivers, producers, or exporters.
// The caller must pass the ID of the component being tested. The ID will be used by the CreateSettings and Check methods.
// The caller must defer a call to `Shutdown` on the returned TestTelemetry.
func SetupTelemetry(id component.ID) (TestTelemetry, error) {
	settings := TestTelemetry{
		id:           id,
		reader:       sdkmetric.NewManualReader(),
		SpanRecorder: new(tracetest.SpanRecorder),
	}

	mp := sdkmetric.NewMeterProvider(
		sdkmetric.WithResource(resource.Empty()),
		sdkmetric.WithReader(settings.reader),
	)

	settings.ts = component.TelemetrySettings{
		Logger:         zap.NewNop(),
		TracerProvider: sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(settings.SpanRecorder)),
		MeterProvider:  mp,
		MetricsLevel:   configtelemetry.LevelDetailed,
		Resource:       pcommon.NewResource(),
	}

	return settings, nil
}
