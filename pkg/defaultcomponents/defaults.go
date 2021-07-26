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

package defaultcomponents // import "aws-observability.io/collector/defaultcomponents

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsemfexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsprometheusremotewriteexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/dynatraceexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/fileexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/logzioexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sapmexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/healthcheckextension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/ecsobserver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/pprofextension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/attributesprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/filterprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricstransformprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/probabilisticsamplerprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourceprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/spanprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsecscontainermetricsreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsxrayreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jaegerreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/statsdreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/zipkinreceiver"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer/consumererror"
	"go.opentelemetry.io/collector/exporter/loggingexporter"
	"go.opentelemetry.io/collector/exporter/otlpexporter"
	"go.opentelemetry.io/collector/exporter/otlphttpexporter"
	"go.opentelemetry.io/collector/receiver/otlpreceiver"
	"go.opentelemetry.io/collector/service/defaultcomponents"
)

// Components register OTel components for aws-otel-collector distribution
func Components() (component.Factories, error) {
	errs := []error{}
	factories, err := defaultcomponents.Components()
	if err != nil {
		return component.Factories{}, err
	}

	// enable extensions
	extensions := []component.ExtensionFactory{
		ecsobserver.NewFactory(),
		healthcheckextension.NewFactory(),
		pprofextension.NewFactory(),
	}

	for _, ext := range factories.Extensions {
		extensions = append(extensions, ext)
	}

	factories.Extensions, err = component.MakeExtensionFactoryMap(extensions...)
	if err != nil {
		errs = append(errs, err)
	}

	// enable the selected receivers
	factories.Receivers, err = component.MakeReceiverFactoryMap(
		prometheusreceiver.NewFactory(),
		otlpreceiver.NewFactory(),
		awsecscontainermetricsreceiver.NewFactory(),
		awscontainerinsightreceiver.NewFactory(),
		awsxrayreceiver.NewFactory(),
		statsdreceiver.NewFactory(),
		jaegerreceiver.NewFactory(),
		zipkinreceiver.NewFactory(),
	)
	if err != nil {
		errs = append(errs, err)
	}

	// enable the selected processors
	processors := []component.ProcessorFactory{
		attributesprocessor.NewFactory(),
		resourceprocessor.NewFactory(),
		probabilisticsamplerprocessor.NewFactory(),
		spanprocessor.NewFactory(),
		filterprocessor.NewFactory(),
		metricstransformprocessor.NewFactory(),
		resourcedetectionprocessor.NewFactory(),
	}
	for _, pr := range factories.Processors {
		processors = append(processors, pr)
	}
	factories.Processors, err = component.MakeProcessorFactoryMap(processors...)
	if err != nil {
		errs = append(errs, err)
	}

	// enable the selected exporters
	factories.Exporters, err = component.MakeExporterFactoryMap(
		awsxrayexporter.NewFactory(),
		awsemfexporter.NewFactory(),
		awsprometheusremotewriteexporter.NewFactory(),
		prometheusexporter.NewFactory(),
		loggingexporter.NewFactory(),
		fileexporter.NewFactory(),
		otlpexporter.NewFactory(),
		otlphttpexporter.NewFactory(),
		dynatraceexporter.NewFactory(),
		sapmexporter.NewFactory(),
		signalfxexporter.NewFactory(),
		// disable it until security team reviews it
		//splunkhecexporter.NewFactory(),
		// New Relic offers native OTLP ingest. Sign up at https://forms.gle/fa2pWcQxgVQYMggEA
		//newrelicexporter.NewFactory(),
		datadogexporter.NewFactory(),
		logzioexporter.NewFactory(),
	)
	if err != nil {
		errs = append(errs, err)
	}

	return factories, consumererror.Combine(errs)
}
