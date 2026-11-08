/*
 * Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License").
 * You may not use this file except in compliance with the License.
 * A copy of the License is located at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * or in the "license" file accompanying this file. This file is distributed
 * on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
 * express or implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

package config

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/confmap/provider/envprovider"
	"go.opentelemetry.io/collector/confmap/provider/fileprovider"
	"go.opentelemetry.io/collector/exporter/otlphttpexporter"
	"go.opentelemetry.io/collector/otelcol"
	"go.opentelemetry.io/collector/pipeline"

	"github.com/aws-observability/aws-otel-collector/pkg/defaultcomponents"
)

// loadBundledConfig loads a configuration from the config directory with the collector's own
// components, the way the collector does at startup.
func loadBundledConfig(t *testing.T, name string) *otelcol.Config {
	t.Helper()
	factories, err := defaultcomponents.Components()
	require.NoError(t, err)

	provider, err := otelcol.NewConfigProvider(otelcol.ConfigProviderSettings{
		ResolverSettings: confmap.ResolverSettings{
			URIs:              []string{"file:" + filepath.Join("..", "..", "config", name)},
			ProviderFactories: []confmap.ProviderFactory{fileprovider.NewFactory(), envprovider.NewFactory()},
			DefaultScheme:     "env",
		},
	})
	require.NoError(t, err)

	cfg, err := provider.Get(context.Background(), factories)
	require.NoError(t, err)
	require.NoError(t, cfg.Validate())
	return cfg
}

func TestCloudWatchOTLPConfigs(t *testing.T) {
	t.Setenv("AWS_REGION", "us-west-2")

	for _, name := range []string{"cloudwatch-otlp.yaml", filepath.Join("ecs", "ecs-cloudwatch-otlp.yaml")} {
		t.Run(name, func(t *testing.T) {
			cfg := loadBundledConfig(t, name)

			// The configurations receive OTLP only.
			assert.Equal(t, []component.ID{component.MustNewID("otlp")}, keys(cfg.Receivers))

			traces := cfg.Service.Pipelines[pipeline.NewID(pipeline.SignalTraces)]
			metrics := cfg.Service.Pipelines[pipeline.NewID(pipeline.SignalMetrics)]
			require.NotNil(t, traces)
			require.NotNil(t, metrics)
			assert.Equal(t, []component.ID{component.MustNewIDWithName("otlp_http", "traces")}, traces.Exporters)
			assert.Equal(t, []component.ID{component.MustNewIDWithName("otlp_http", "metrics")}, metrics.Exporters)

			tracesExporter := cfg.Exporters[component.MustNewIDWithName("otlp_http", "traces")].(*otlphttpexporter.Config)
			metricsExporter := cfg.Exporters[component.MustNewIDWithName("otlp_http", "metrics")].(*otlphttpexporter.Config)
			assert.Equal(t, "https://xray.us-west-2.amazonaws.com/v1/traces", tracesExporter.TracesEndpoint)
			assert.Equal(t, "https://monitoring.us-west-2.amazonaws.com/v1/metrics", metricsExporter.MetricsEndpoint)
		})
	}
}

func keys(m map[component.ID]component.Config) []component.ID {
	ids := make([]component.ID, 0, len(m))
	for id := range m {
		ids = append(ids, id)
	}
	return ids
}
