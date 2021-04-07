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
	"os"
	"reflect"
	"testing"

	"github.com/crossdock/crossdock-go/require"
	"github.com/spf13/cobra"
	"go.opentelemetry.io/collector/service"

	"github.com/aws-observability/aws-otel-collector/pkg/defaultcomponents"
	"github.com/stretchr/testify/assert"
)

func TestGetCfgFactoryReturn(t *testing.T) {
	cfgFunc := GetCfgFactory()
	assert.True(t, reflect.TypeOf(cfgFunc).Kind() == reflect.Func)
}

func TestGetCfgFactoryConfig(t *testing.T) {
	factories, _ := defaultcomponents.Components()
	params := service.Parameters{
		Factories: factories,
	}

	t.Run("test_invalid_config", func(t *testing.T) {
		app, err := service.New(params)
		require.NoError(t, err)
		err = app.Command().ParseFlags([]string{
			"--config=invalid-path/otelcol-config.yaml",
		})
		require.NoError(t, err)
		cfgFunc := GetCfgFactory()
		_, err = cfgFunc(app.Command(), factories)
		require.Error(t, err)
	})

	t.Run("test_valid_config", func(t *testing.T) {
		app, err := service.New(params)
		require.NoError(t, err)
		err = app.Command().ParseFlags([]string{
			"--config=testdata/config.yaml",
		})
		require.NoError(t, err)
		cfgFunc := GetCfgFactory()
		_, err = cfgFunc(app.Command(), factories)
		require.NoError(t, err)
	})
}

func TestGetCfgFactoryContainer(t *testing.T) {
	os.Setenv("AOT_CONFIG_CONTENT", "extensions:\n  health_check:\n  pprof:\n    endpoint: 0.0.0.0:1777\nreceivers:\n  otlp:\n    protocols:\n      grpc:\n        endpoint: 0.0.0.0:4317\nprocessors:\n  batch:\nexporters:\n  logging:\n    loglevel: debug\n  awsxray:\n    local_mode: true\n    region: 'us-west-2'\n  awsemf:\n    region: 'us-west-2'\nservice:\n  pipelines:\n    traces:\n      receivers: [prometheusreceiver]\n      exporters: [logging,awsxray]\n    metrics:\n      receivers: [prometheusreceiver]\n      exporters: [awsemf]\n  extensions: [pprof]")
	defer os.Unsetenv("AOT_CONFIG_CONTENT")
	cmd := &cobra.Command{}
	factories, _ := defaultcomponents.Components()
	cfgFunc := GetCfgFactory()
	cfgModel, err := cfgFunc(cmd, factories)
	if err != nil {
		t.Log(err)
	}
	assert.True(t, err == nil)
	assert.True(t, cfgModel != nil)
	if cfgModel == nil {
		return
	}
	assert.True(t, cfgModel.Receivers != nil && cfgModel.Receivers["otlp"] != nil)
	assert.True(t, cfgModel.Receivers != nil && cfgModel.Receivers["prometheus"] == nil)
	assert.True(t, cfgModel.Exporters != nil && cfgModel.Exporters["awsemf"] != nil)
	assert.True(t, cfgModel.Processors != nil && cfgModel.Extensions["pprof"] != nil)
}
