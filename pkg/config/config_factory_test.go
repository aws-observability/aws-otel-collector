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
	"fmt"
	"log"
	"path/filepath"
	"testing"

	"go.opentelemetry.io/collector/featuregate"
	"go.opentelemetry.io/collector/otelcol"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/pprofextension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsxrayreceiver"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"

	"github.com/aws-observability/aws-otel-collector/pkg/defaultcomponents"
)

func getValidTestConfigPath() string {
	return filepath.Join("testdata", "config.yaml")
}
func TestGetCfgFactoryConfig(t *testing.T) {
	params := otelcol.CollectorSettings{
		Factories: defaultcomponents.Components,
	}
	factories, err := params.Factories()
	require.NoError(t, err)

	t.Run("test_invalid_path", func(t *testing.T) {
		cmd := &cobra.Command{
			Use:          params.BuildInfo.Command,
			Version:      params.BuildInfo.Version,
			SilenceUsage: true,
		}
		flagSet := Flags(featuregate.NewRegistry())
		cmd.Flags().AddGoFlagSet(flagSet)
		err := cmd.ParseFlags([]string{
			"--config=invalid-path/otelcol-config.yaml",
		})
		require.NoError(t, err)

		provider, err := otelcol.NewConfigProvider(GetConfigProviderSettings(flagSet))
		if err != nil {
			log.Panicf("Err on creating Config Provider: %v\n", err)
		}
		_, err = provider.Get(context.Background(), factories)
		require.Error(t, err)
	})

	t.Run("test_config_with_env_var_set", func(t *testing.T) {
		const expectedEndpoint = "0.0.0.0:2000"
		t.Setenv("XRAY_ENDPOINT", expectedEndpoint)
		cmd := &cobra.Command{
			Use:          params.BuildInfo.Command,
			Version:      params.BuildInfo.Version,
			SilenceUsage: true,
		}
		flagSet := Flags(featuregate.NewRegistry())

		cmd.Flags().AddGoFlagSet(flagSet)
		err := cmd.ParseFlags([]string{
			fmt.Sprintf("--config=%s", getValidTestConfigPath()),
		})
		require.NoError(t, err)
		provider, err := otelcol.NewConfigProvider(GetConfigProviderSettings(flagSet))
		if err != nil {
			log.Panicf("Err on creating Config Provider: %v\n", err)
		}
		cfg, err := provider.Get(context.Background(), factories)
		require.NoError(t, err)
		require.NotNil(t, cfg)
		receiver := cfg.Receivers[component.MustNewID("awsxray")].(*awsxrayreceiver.Config)
		require.NotNil(t, receiver)
		require.Equal(t, expectedEndpoint, receiver.Endpoint)
	})

	t.Run("test_config_without_env_var_set", func(t *testing.T) {
		cmd := &cobra.Command{
			Use:          params.BuildInfo.Command,
			Version:      params.BuildInfo.Version,
			SilenceUsage: true,
		}
		flagSet := Flags(featuregate.NewRegistry())
		cmd.Flags().AddGoFlagSet(flagSet)
		err := cmd.ParseFlags([]string{
			fmt.Sprintf("--config=%s", getValidTestConfigPath()),
		})
		require.NoError(t, err)
		provider, err := otelcol.NewConfigProvider(GetConfigProviderSettings(flagSet))
		if err != nil {
			log.Panicf("Err on creating Config Provider: %v\n", err)
		}
		cfg, err := provider.Get(context.Background(), factories)
		require.NoError(t, err)
		require.NotNil(t, cfg)
		receiver := cfg.Receivers[component.MustNewID("awsxray")].(*awsxrayreceiver.Config)
		require.NotNil(t, receiver)
		require.Empty(t, receiver.Endpoint)
	})
}

func TestGetMapProviderContainer(t *testing.T) {
	const expectedEndpoint = "0.0.0.0:1777"
	t.Setenv("PPROF_ENDPOINT", expectedEndpoint)

	t.Setenv(envKey, "extensions:\n  health_check:\n  pprof:\n    endpoint: '${PPROF_ENDPOINT}'\nreceivers:\n  otlp:\n    protocols:\n      grpc:\n        endpoint: 0.0.0.0:4317\nprocessors:\n  batch:\nexporters:\n  debug:\n    verbosity: detailed\n  awsxray:\n    local_mode: true\n    region: 'us-west-2'\n  awsemf:\n    region: 'us-west-2'\nservice:\n  pipelines:\n    traces:\n      receivers: [otlp]\n      exporters: [debug,awsxray]\n    metrics:\n      receivers: [otlp]\n      exporters: [awsemf]\n  extensions: [pprof]")

	factories, _ := defaultcomponents.Components()
	provider, err := otelcol.NewConfigProvider(GetConfigProviderSettings(Flags(featuregate.NewRegistry())))
	if err != nil {
		log.Panicf("Err on creating Config Provider: %v\n", err)
	}
	cfg, err := provider.Get(context.Background(), factories)
	require.NoError(t, err)
	require.NotNil(t, cfg)
	extension := cfg.Extensions[component.MustNewID("pprof")].(*pprofextension.Config)
	require.NotNil(t, extension)
	require.Equal(t, expectedEndpoint, extension.TCPAddr.Endpoint)

	require.NotNil(t, cfg.Receivers[component.MustNewID("otlp")])
	require.NotNil(t, cfg.Exporters[component.MustNewID("awsemf")])
}
