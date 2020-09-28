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
	"os"
	"reflect"
	"testing"

	"aws-observability.io/collector/pkg/defaultcomponents"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/config"
)

func TestGetCfgFactory(t *testing.T) {
	cfgFunc := GetCfgFactory()
	assert.True(t, reflect.TypeOf(cfgFunc).Kind() == reflect.Func)
}

func TestGetCfgFactoryContainer(t *testing.T) {
	os.Setenv("AOT_CONFIG_CONTENT", "extensions:\n  health_check:\n  pprof:\n    endpoint: 0.0.0.0:1777\nreceivers:\n  otlp:\n    protocols:\n      grpc:\n        endpoint: 0.0.0.0:55680\nprocessors:\n  batch:\n  queued_retry:\nexporters:\n  logging:\n    loglevel: debug\n  awsxray:\n    local_mode: true\n    region: 'us-west-2'\n  awsemf:\n    region: 'us-west-2'\nservice:\n  pipelines:\n    traces:\n      receivers: [prometheusreceiver]\n      exporters: [logging,awsxray]\n    metrics:\n      receivers: [prometheusreceiver]\n      exporters: [awsemf]\n  extensions: [pprof]")
	v := config.NewViper()
	factories, _ := defaultcomponents.Components()
	cfgFunc := GetCfgFactory()
	cfgModel, err := cfgFunc(v, factories)
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
	assert.True(t, cfgModel.Processors != nil && cfgModel.Processors["queued_retry"] != nil)
	assert.True(t, cfgModel.Processors != nil && cfgModel.Extensions["pprof"] != nil)
}
