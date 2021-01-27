/*
 * Copyright 2020 Amazon.com, Inc. or its affiliates. All Rights Reserved.
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
	"flag"
	"github.com/spf13/cobra"
	"os"
	"reflect"
	"testing"

	"go.opentelemetry.io/collector/service/builder"

	"github.com/aws-observability/aws-otel-collector/pkg/defaultcomponents"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/config"
)

func TestGetCfgFactoryReturn(t *testing.T) {
	cfgFunc := GetCfgFactory()
	assert.True(t, reflect.TypeOf(cfgFunc).Kind() == reflect.Func)
}

func TestGetCfgFactoryWithoutConfig(t *testing.T) {
	newFlag := flag.NewFlagSet("newFlag", flag.ContinueOnError)
	builder.Flags(newFlag)
	v := config.NewViper()
	cmd := &cobra.Command{}
	factories, _ := defaultcomponents.Components()
	cfgFunc := GetCfgFactory()
	_, err := cfgFunc(v, cmd, factories)
	assert.True(t, err != nil)
}

func TestGetCfgFactoryContainer(t *testing.T) {
	os.Setenv("AOT_CONFIG_CONTENT", "extensions:\n  health_check:\n  pprof:\n    endpoint: 0.0.0.0:1777\nreceivers:\n  otlp:\n    protocols:\n      grpc:\n        endpoint: 0.0.0.0:55680\nprocessors:\n  batch:\nexporters:\n  logging:\n    loglevel: debug\n  awsxray:\n    local_mode: true\n    region: 'us-west-2'\n  awsemf:\n    region: 'us-west-2'\nservice:\n  pipelines:\n    traces:\n      receivers: [prometheusreceiver]\n      exporters: [logging,awsxray]\n    metrics:\n      receivers: [prometheusreceiver]\n      exporters: [awsemf]\n  extensions: [pprof]")
	defer os.Unsetenv("AOT_CONFIG_CONTENT")
	v := config.NewViper()
	cmd := &cobra.Command{}
	factories, _ := defaultcomponents.Components()
	cfgFunc := GetCfgFactory()
	cfgModel, err := cfgFunc(v, cmd, factories)
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
