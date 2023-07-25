// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package dynatraceexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/dynatraceexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	exp "go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exporterhelper"

	dtconfig "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/dynatraceexporter/config"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/dynatraceexporter/internal/metadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/resourcetotelemetry"
)

// NewFactory creates a Dynatrace exporter factory
func NewFactory() exp.Factory {
	return exp.NewFactory(
		metadata.Type,
		createDefaultConfig,
		exp.WithMetrics(createMetricsExporter, metadata.MetricsStability),
	)
}

// createDefaultConfig creates the default exporter configuration
func createDefaultConfig() component.Config {
	return &dtconfig.Config{
		RetrySettings: exporterhelper.NewDefaultRetrySettings(),
		QueueSettings: exporterhelper.NewDefaultQueueSettings(),
		ResourceToTelemetrySettings: resourcetotelemetry.Settings{
			Enabled: false,
		},

		APIToken:           "",
		HTTPClientSettings: confighttp.HTTPClientSettings{Endpoint: ""},

		Tags:              []string{},
		DefaultDimensions: make(map[string]string),
	}
}

// createMetricsExporter creates a metrics exporter based on this
func createMetricsExporter(
	ctx context.Context,
	set exp.CreateSettings,
	c component.Config,
) (exp.Metrics, error) {

	cfg := c.(*dtconfig.Config)

	exp := newMetricsExporter(set, cfg)

	exporter, err := exporterhelper.NewMetricsExporter(
		ctx,
		set,
		cfg,
		exp.PushMetricsData,
		exporterhelper.WithQueue(cfg.QueueSettings),
		exporterhelper.WithRetry(cfg.RetrySettings),
		exporterhelper.WithStart(exp.start),
	)
	if err != nil {
		return nil, err
	}
	return resourcetotelemetry.WrapMetricsExporter(cfg.ResourceToTelemetrySettings, exporter), nil
}
