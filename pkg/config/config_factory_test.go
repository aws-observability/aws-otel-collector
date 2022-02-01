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
	"os"
	"testing"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/pprofextension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsxrayreceiver"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/service"

	"github.com/aws-observability/aws-otel-collector/pkg/defaultcomponents"
)

func TestGetCfgFactoryConfig(t *testing.T) {
	factories, _ := defaultcomponents.Components()
	params := service.CollectorSettings{
		Factories: factories,
	}

	resetFlag := func() {
		configFlag = new(stringArrayValue)
	}

	t.Run("test_invalid_path", func(t *testing.T) {
		cmd := &cobra.Command{
			Use:          params.BuildInfo.Command,
			Version:      params.BuildInfo.Version,
			SilenceUsage: true,
		}
		cmd.Flags().AddGoFlagSet(Flags())
		err := cmd.ParseFlags([]string{
			"--config=invalid-path/otelcol-config.yaml",
		})
		require.NoError(t, err)
		provider := GetConfigProvider()
		_, err = provider.Get(context.Background(), factories)
		require.Error(t, err)
		resetFlag()
	})

	t.Run("test_config_with_env_var_set", func(t *testing.T) {
		const expectedEndpoint = "0.0.0.0:2000"
		os.Setenv("XRAY_ENDPOINT", expectedEndpoint)
		defer os.Unsetenv("XRAY_ENDPOINT")
		cmd := &cobra.Command{
			Use:          params.BuildInfo.Command,
			Version:      params.BuildInfo.Version,
			SilenceUsage: true,
		}

		cmd.Flags().AddGoFlagSet(Flags())
		err := cmd.ParseFlags([]string{
			"--config=testdata/config.yaml",
		})
		require.NoError(t, err)
		provider := GetConfigProvider()
		cfg, err := provider.Get(context.Background(), factories)
		require.NoError(t, err)
		require.NotNil(t, cfg)
		receiver := cfg.Receivers[config.NewComponentID("awsxray")].(*awsxrayreceiver.Config)
		require.NotNil(t, receiver)
		require.Equal(t, expectedEndpoint, receiver.Endpoint)
		resetFlag()
	})

	t.Run("test_config_without_env_var_set", func(t *testing.T) {
		cmd := &cobra.Command{
			Use:          params.BuildInfo.Command,
			Version:      params.BuildInfo.Version,
			SilenceUsage: true,
		}

		cmd.Flags().AddGoFlagSet(Flags())
		err := cmd.ParseFlags([]string{
			"--config=testdata/config.yaml",
		})
		require.NoError(t, err)
		provider := GetConfigProvider()
		cfg, err := provider.Get(context.Background(), factories)
		require.NoError(t, err)
		require.NotNil(t, cfg)
		receiver := cfg.Receivers[config.NewComponentID("awsxray")].(*awsxrayreceiver.Config)
		require.NotNil(t, receiver)
		require.Empty(t, receiver.Endpoint)
		resetFlag()
	})
}

func TestGetMapProviderContainer(t *testing.T) {
	const expectedEndpoint = "0.0.0.0:1777"
	os.Setenv("PPROF_ENDPOINT", expectedEndpoint)
	defer os.Unsetenv("PPROF_ENDPOINT")

	os.Setenv(envKey, "extensions:\n  health_check:\n  pprof:\n    endpoint: '${PPROF_ENDPOINT}'\nreceivers:\n  otlp:\n    protocols:\n      grpc:\n        endpoint: 0.0.0.0:4317\nprocessors:\n  batch:\nexporters:\n  logging:\n    loglevel: debug\n  awsxray:\n    local_mode: true\n    region: 'us-west-2'\n  awsemf:\n    region: 'us-west-2'\nservice:\n  pipelines:\n    traces:\n      receivers: [otlp]\n      exporters: [logging,awsxray]\n    metrics:\n      receivers: [otlp]\n      exporters: [awsemf]\n  extensions: [pprof]")
	defer os.Unsetenv(envKey)

	factories, _ := defaultcomponents.Components()
	provider := GetConfigProvider()
	cfg, err := provider.Get(context.Background(), factories)
	require.NoError(t, err)
	require.NotNil(t, cfg)
	extension := cfg.Extensions[config.NewComponentID("pprof")].(*pprofextension.Config)
	require.NotNil(t, extension)
	require.Equal(t, expectedEndpoint, extension.TCPAddr.Endpoint)

	require.NotNil(t, cfg.Receivers[config.NewComponentID("otlp")])
	require.NotNil(t, cfg.Exporters[config.NewComponentID("awsemf")])
}
