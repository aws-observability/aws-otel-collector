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
	"go.opentelemetry.io/collector/featuregate"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	exportersCount  = 15
	receiversCount  = 9
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
	assert.NotNil(t, exporters["awsxray"])
	assert.NotNil(t, exporters["awsemf"])
	// core exporters
	assert.NotNil(t, exporters["logging"])
	assert.NotNil(t, exporters["otlp"])
	assert.NotNil(t, exporters["otlphttp"])
	// other exporters
	assert.NotNil(t, exporters["file"])
	assert.NotNil(t, exporters["datadog"])
	assert.NotNil(t, exporters["dynatrace"])
	assert.NotNil(t, exporters["prometheus"])
	assert.NotNil(t, exporters["sapm"])
	assert.NotNil(t, exporters["signalfx"])
	assert.NotNil(t, exporters["logzio"])
	assert.NotNil(t, exporters["prometheusremotewrite"])
	assert.NotNil(t, exporters["kafka"])
	assert.NotNil(t, exporters["loadbalancing"])

	receivers := factories.Receivers
	assert.Len(t, receivers, receiversCount)
	// aws receivers
	assert.NotNil(t, receivers["awsecscontainermetrics"])
	assert.NotNil(t, receivers["awscontainerinsightreceiver"])
	assert.NotNil(t, receivers["awsxray"])
	assert.NotNil(t, receivers["statsd"])

	// core receivers
	assert.NotNil(t, receivers["otlp"])
	// other receivers
	assert.NotNil(t, receivers["prometheus"])
	assert.NotNil(t, receivers["zipkin"])
	assert.NotNil(t, receivers["jaeger"])
	assert.NotNil(t, receivers["kafka"])

	extensions := factories.Extensions
	assert.Len(t, extensions, extensionsCount)
	// aws extensions
	assert.NotNil(t, extensions["awsproxy"])
	assert.NotNil(t, extensions["ecs_observer"])
	assert.NotNil(t, extensions["sigv4auth"])
	// core extensions
	assert.NotNil(t, extensions["zpages"])
	assert.NotNil(t, extensions["memory_ballast"])
	// other extensions
	assert.NotNil(t, extensions["pprof"])
	assert.NotNil(t, extensions["health_check"])

	processors := factories.Processors
	assert.Len(t, processors, processorCount)
	// aws processors
	assert.NotNil(t, processors["experimental_metricsgeneration"])
	// core processors
	assert.NotNil(t, processors["batch"])
	assert.NotNil(t, processors["memory_limiter"])
	// other processors
	assert.NotNil(t, processors["attributes"])
	assert.NotNil(t, processors["resource"])
	assert.NotNil(t, processors["probabilistic_sampler"])
	assert.NotNil(t, processors["span"])
	assert.NotNil(t, processors["filter"])
	assert.NotNil(t, processors["metricstransform"])
	assert.NotNil(t, processors["resourcedetection"])
	assert.NotNil(t, processors["cumulativetodelta"])
	assert.NotNil(t, processors["deltatorate"])
	assert.NotNil(t, processors["groupbytrace"])
	assert.NotNil(t, processors["tail_sampling"])
	assert.NotNil(t, processors["k8sattributes"])

	// Ensure that the components behind feature gates aren't included
	assert.Nil(t, receivers["filelog"])
	assert.Nil(t, exporters["awscloudwatchlogs"])
	assert.Nil(t, extensions["file_storage"])
}

func TestEnabledFeatureGate(t *testing.T) {
	testCases := []struct {
		featureName string
		expectedLen int
	}{
		{"adot.receiver.filelog", receiversCount + 1},
		{"adot.exporter.awscloudwatchlogs", exportersCount + 1},
		{"adot.extension.file_storage", extensionsCount + 1},
	}

	for _, tc := range testCases {
		t.Run(tc.featureName, func(t *testing.T) {
			err := featuregate.GlobalRegistry().Set(tc.featureName, true)
			assert.NoError(t, err)

			factories, err := Components()
			assert.NoError(t, err)

			switch tc.featureName {
			case "adot.receiver.filelog":
				receivers := factories.Receivers
				assert.Len(t, receivers, tc.expectedLen)
				assert.NotNil(t, receivers["filelog"])

			case "adot.exporter.awscloudwatchlogs":
				exporters := factories.Exporters
				assert.Len(t, exporters, tc.expectedLen)
				assert.NotNil(t, exporters["awscloudwatchlogs"])

			case "adot.extension.file_storage":
				extensions := factories.Extensions
				assert.Len(t, extensions, tc.expectedLen)
				assert.NotNil(t, extensions["file_storage"])
			}
		})
	}
}
