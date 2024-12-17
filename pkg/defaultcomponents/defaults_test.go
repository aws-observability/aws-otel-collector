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

package defaultcomponents

import (
	"testing"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/featuregate"

	"github.com/stretchr/testify/assert"
)

const (
	exportersCount  = 15
	receiversCount  = 10
	extensionsCount = 7
	processorCount  = 15
)

// Assert that the components behind feature gate are not in the default
func TestComponents(t *testing.T) {
	factories, err := Components()
	assert.NoError(t, err)
	exporters := factories.Exporters
	assert.Len(t, exporters, exportersCount)
	// aws exporters
	assert.NotNil(t, exporters[component.MustNewType("awsxray")])
	assert.NotNil(t, exporters[component.MustNewType("awsemf")])
	// core exporters
	assert.NotNil(t, exporters[component.MustNewType("debug")])
	assert.NotNil(t, exporters[component.MustNewType("otlp")])
	assert.NotNil(t, exporters[component.MustNewType("otlphttp")])
	// other exporters
	assert.NotNil(t, exporters[component.MustNewType("file")])
	assert.NotNil(t, exporters[component.MustNewType("prometheus")])
	assert.NotNil(t, exporters[component.MustNewType("prometheusremotewrite")])
	assert.NotNil(t, exporters[component.MustNewType("kafka")])
	assert.NotNil(t, exporters[component.MustNewType("loadbalancing")])

	receivers := factories.Receivers
	assert.Len(t, receivers, receiversCount)
	// aws receivers
	assert.NotNil(t, receivers[component.MustNewType("awsecscontainermetrics")])
	assert.NotNil(t, receivers[component.MustNewType("awscontainerinsightreceiver")])
	assert.NotNil(t, receivers[component.MustNewType("awsxray")])
	assert.NotNil(t, receivers[component.MustNewType("statsd")])
	assert.NotNil(t, exporters[component.MustNewType("awscloudwatchlogs")])

	// core receivers
	assert.NotNil(t, receivers[component.MustNewType("otlp")])
	// other receivers
	assert.NotNil(t, receivers[component.MustNewType("prometheus")])
	assert.NotNil(t, receivers[component.MustNewType("zipkin")])
	assert.NotNil(t, receivers[component.MustNewType("jaeger")])
	assert.NotNil(t, receivers[component.MustNewType("kafka")])
	assert.NotNil(t, receivers[component.MustNewType("filelog")])

	extensions := factories.Extensions
	assert.Len(t, extensions, extensionsCount)
	// aws extensions
	assert.NotNil(t, extensions[component.MustNewType("awsproxy")])
	assert.NotNil(t, extensions[component.MustNewType("ecs_observer")])
	assert.NotNil(t, extensions[component.MustNewType("sigv4auth")])
	// core extensions
	assert.NotNil(t, extensions[component.MustNewType("zpages")])
	// other extensions
	assert.NotNil(t, extensions[component.MustNewType("pprof")])
	assert.NotNil(t, extensions[component.MustNewType("health_check")])
	assert.NotNil(t, extensions[component.MustNewType("file_storage")])

	processors := factories.Processors
	assert.Len(t, processors, processorCount)
	// aws processors
	assert.NotNil(t, processors[component.MustNewType("metricsgeneration")])
	// core processors
	assert.NotNil(t, processors[component.MustNewType("batch")])
	assert.NotNil(t, processors[component.MustNewType("memory_limiter")])
	// other processors
	assert.NotNil(t, processors[component.MustNewType("attributes")])
	assert.NotNil(t, processors[component.MustNewType("resource")])
	assert.NotNil(t, processors[component.MustNewType("probabilistic_sampler")])
	assert.NotNil(t, processors[component.MustNewType("span")])
	assert.NotNil(t, processors[component.MustNewType("filter")])
	assert.NotNil(t, processors[component.MustNewType("metricstransform")])
	assert.NotNil(t, processors[component.MustNewType("resourcedetection")])
	assert.NotNil(t, processors[component.MustNewType("cumulativetodelta")])
	assert.NotNil(t, processors[component.MustNewType("deltatorate")])
	assert.NotNil(t, processors[component.MustNewType("groupbytrace")])
	assert.NotNil(t, processors[component.MustNewType("tail_sampling")])
	assert.NotNil(t, processors[component.MustNewType("k8sattributes")])

	// Ensure that the components behind feature gates are included
	assert.NotNil(t, exporters[component.MustNewType("datadog")])
	assert.NotNil(t, exporters[component.MustNewType("sapm")])
	assert.NotNil(t, exporters[component.MustNewType("signalfx")])
	assert.NotNil(t, exporters[component.MustNewType("logzio")])
}

func TestEnableFeatureGate(t *testing.T) {

	testCases := []struct {
		desc        string
		featureName string
		component   component.Type
	}{
		{
			desc:        "disable datadog exporter",
			featureName: "adot.exporter.datadogexporter.deprecation",
			component:   component.MustNewType("datadog"),
		},
		{
			desc:        "disable logzio exporter",
			featureName: "adot.exporter.logzioexporter.deprecation",
			component:   component.MustNewType("logzio"),
		},
		{
			desc:        "disable sapm exporter",
			featureName: "adot.exporter.sapmexporter.deprecation",
			component:   component.MustNewType("sapm"),
		},
		{
			desc:        "disable signalfx exporter",
			featureName: "adot.exporter.signalfxexporter.deprecation",
			component:   component.MustNewType("signalfx"),
		},
	}
	expectedLen := exportersCount

	for _, tc := range testCases {
		expectedLen--
		t.Run(tc.desc, func(t *testing.T) {
			err := featuregate.GlobalRegistry().Set(tc.featureName, true)
			assert.NoError(t, err)

			factories, err := Components()
			assert.NoError(t, err)

			exporters := factories.Exporters
			assert.Len(t, exporters, expectedLen)
			assert.Nil(t, exporters[tc.component])
		})
	}
}
