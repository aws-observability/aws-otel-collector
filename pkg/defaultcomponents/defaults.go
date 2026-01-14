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
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awscloudwatchlogsexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsemfexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/fileexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/kafkaexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/loadbalancingexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/logzioexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusremotewriteexporter"
	//lint:ignore SA1019 sapmexporter is deprecated upstream but still supported in ADOT
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sapmexporter" //nolint:staticcheck
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/awsproxy"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/healthcheckextension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/ecsobserver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/pprofextension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/sigv4authextension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage/filestorage"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/attributesprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/cumulativetodeltaprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatorateprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/filterprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/groupbytraceprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/k8sattributesprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricsgenerationprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricstransformprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/probabilisticsamplerprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourceprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/spanprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsecscontainermetricsreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsxrayreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/filelogreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jaegerreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kafkareceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/statsdreceiver"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/zipkinreceiver"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/debugexporter"
	"go.opentelemetry.io/collector/exporter/otlpexporter"
	"go.opentelemetry.io/collector/exporter/otlphttpexporter"
	"go.opentelemetry.io/collector/extension"
	"go.opentelemetry.io/collector/extension/zpagesextension"
	"go.opentelemetry.io/collector/featuregate"
	"go.opentelemetry.io/collector/otelcol"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/collector/processor/batchprocessor"
	"go.opentelemetry.io/collector/processor/memorylimiterprocessor"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/otlpreceiver"
	"go.opentelemetry.io/collector/service/telemetry/otelconftelemetry"
	"go.uber.org/multierr"
)

var datadogExporterFeatureGateDeprecation = featuregate.GlobalRegistry().MustRegister("adot.exporter.datadogexporter.deprecation",
	featuregate.StageAlpha,
	featuregate.WithRegisterDescription("Removes the Datadog exporter from the set of configurable exporters "))

var logzioExporterFeatureGateDeprecation = featuregate.GlobalRegistry().MustRegister("adot.exporter.logzioexporter.deprecation",
	featuregate.StageAlpha,
	featuregate.WithRegisterDescription("Removes the Logzio Exporter from the set of configurable exporters "))

var sapmExporterFeatureGateDeprecation = featuregate.GlobalRegistry().MustRegister("adot.exporter.sapmexporter.deprecation",
	featuregate.StageAlpha,
	featuregate.WithRegisterDescription("Removes the SAPM Exporter from the set of configurable exporters"))

var signalfxExporterFeatureGateDeprecation = featuregate.GlobalRegistry().MustRegister("adot.exporter.signalfxexporter.deprecation",
	featuregate.StageAlpha,
	featuregate.WithRegisterDescription("Removes the SignalFx Metrics Exporter from the set of configurable exporters"))

// Components register OTel components for ADOT-collector distribution
func Components() (otelcol.Factories, error) {
	var errs error
	extensionsList := []extension.Factory{
		awsproxy.NewFactory(),
		ecsobserver.NewFactory(),
		healthcheckextension.NewFactory(),
		pprofextension.NewFactory(),
		sigv4authextension.NewFactory(),
		zpagesextension.NewFactory(),
		filestorage.NewFactory(),
	}

	extensions, err := otelcol.MakeFactoryMap[extension.Factory](extensionsList...)

	if err != nil {
		errs = multierr.Append(errs, err)
	}

	receiverList := []receiver.Factory{
		awsecscontainermetricsreceiver.NewFactory(),
		awscontainerinsightreceiver.NewFactory(),
		awsxrayreceiver.NewFactory(),
		statsdreceiver.NewFactory(),
		prometheusreceiver.NewFactory(),
		kafkareceiver.NewFactory(),
		jaegerreceiver.NewFactory(),
		zipkinreceiver.NewFactory(),
		otlpreceiver.NewFactory(),
		filelogreceiver.NewFactory(),
	}

	receivers, err := otelcol.MakeFactoryMap[receiver.Factory](receiverList...)

	if err != nil {
		errs = multierr.Append(errs, err)
	}

	processorList := []processor.Factory{
		groupbytraceprocessor.NewFactory(),
		tailsamplingprocessor.NewFactory(),
		attributesprocessor.NewFactory(),
		resourceprocessor.NewFactory(),
		probabilisticsamplerprocessor.NewFactory(),
		spanprocessor.NewFactory(),
		filterprocessor.NewFactory(),
		metricstransformprocessor.NewFactory(),
		resourcedetectionprocessor.NewFactory(),
		metricsgenerationprocessor.NewFactory(),
		cumulativetodeltaprocessor.NewFactory(),
		deltatorateprocessor.NewFactory(),
		batchprocessor.NewFactory(),
		memorylimiterprocessor.NewFactory(),
		k8sattributesprocessor.NewFactory(),
	}
	processors, err := otelcol.MakeFactoryMap[processor.Factory](processorList...)

	if err != nil {
		errs = multierr.Append(errs, err)
	}

	// enable the selected exporters
	exporterList := []exporter.Factory{
		awsemfexporter.NewFactory(),
		prometheusremotewriteexporter.NewFactory(),
		prometheusexporter.NewFactory(),
		fileexporter.NewFactory(),
		kafkaexporter.NewFactory(),
		debugexporter.NewFactory(),
		otlpexporter.NewFactory(),
		otlphttpexporter.NewFactory(),
		awsxrayexporter.NewFactory(),
		loadbalancingexporter.NewFactory(),
		awscloudwatchlogsexporter.NewFactory(),
	}
	if !datadogExporterFeatureGateDeprecation.IsEnabled() {
		exporterList = append(exporterList, datadogexporter.NewFactory())
	}
	if !logzioExporterFeatureGateDeprecation.IsEnabled() {
		exporterList = append(exporterList, logzioexporter.NewFactory())
	}
	if !sapmExporterFeatureGateDeprecation.IsEnabled() {
		exporterList = append(exporterList, sapmexporter.NewFactory())
	}
	if !signalfxExporterFeatureGateDeprecation.IsEnabled() {
		exporterList = append(exporterList, signalfxexporter.NewFactory())
	}

	exporters, err := otelcol.MakeFactoryMap[exporter.Factory](exporterList...)

	if err != nil {
		errs = multierr.Append(errs, err)
	}

	factories := otelcol.Factories{
		Telemetry:  otelconftelemetry.NewFactory(),
		Extensions: extensions,
		Receivers:  receivers,
		Processors: processors,
		Exporters:  exporters,
	}

	return factories, errs
}
