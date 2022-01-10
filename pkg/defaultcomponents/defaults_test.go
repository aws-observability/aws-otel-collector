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

	"github.com/stretchr/testify/require"
)

const (
	exportersCount  = 13
	receiversCount  = 8
	extensionsCount = 6
	processorCount  = 12
)

func TestComponents(t *testing.T) {
	factories, err := Components()
	require.NoError(t, err)
	exporters := factories.Exporters
	require.Len(t, exporters, exportersCount)
	// aws exporters
	require.NotNil(t, exporters["awsxray"])
	require.NotNil(t, exporters["awsemf"])
	require.NotNil(t, exporters["awsprometheusremotewrite"])
	// core exporters
	require.NotNil(t, exporters["logging"])
	require.NotNil(t, exporters["otlp"])
	require.NotNil(t, exporters["otlphttp"])
	// other exporters
	require.NotNil(t, exporters["file"])
	require.NotNil(t, exporters["datadog"])
	require.NotNil(t, exporters["dynatrace"])
	require.NotNil(t, exporters["prometheus"])
	require.NotNil(t, exporters["sapm"])
	require.NotNil(t, exporters["signalfx"])
	require.NotNil(t, exporters["logzio"])

	receivers := factories.Receivers
	require.Len(t, receivers, receiversCount)
	// aws receivers
	require.NotNil(t, receivers["awsecscontainermetrics"])
	require.NotNil(t, receivers["awscontainerinsightreceiver"])
	require.NotNil(t, receivers["awsxray"])
	require.NotNil(t, receivers["statsd"])
	// core receivers
	require.NotNil(t, receivers["otlp"])
	// other receivers
	require.NotNil(t, receivers["prometheus"])
	require.NotNil(t, receivers["zipkin"])
	require.NotNil(t, receivers["jaeger"])

	extensions := factories.Extensions
	require.Len(t, extensions, extensionsCount)
	// aws extensions
	require.NotNil(t, extensions["awsproxy"])
	require.NotNil(t, extensions["ecs_observer"])
	// core extensions
	require.NotNil(t, extensions["zpages"])
	require.NotNil(t, extensions["memory_ballast"])
	// other extensions
	require.NotNil(t, extensions["pprof"])
	require.NotNil(t, extensions["health_check"])

	processors := factories.Processors
	require.Len(t, processors, processorCount)
	// aws processors
	require.NotNil(t, processors["experimental_metricsgeneration"])
	// core processors
	require.NotNil(t, processors["batch"])
	require.NotNil(t, processors["memory_limiter"])
	// other processors
	require.NotNil(t, processors["attributes"])
	require.NotNil(t, processors["resource"])
	require.NotNil(t, processors["probabilistic_sampler"])
	require.NotNil(t, processors["span"])
	require.NotNil(t, processors["filter"])
	require.NotNil(t, processors["metricstransform"])
	require.NotNil(t, processors["resourcedetection"])
	require.NotNil(t, processors["cumulativetodelta"])
	require.NotNil(t, processors["deltatorate"])
}
