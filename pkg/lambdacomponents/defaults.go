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

package lambdacomponents

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsemfexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusremotewriteexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/basicauthextension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/sigv4authextension"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/loggingexporter"
	"go.opentelemetry.io/collector/exporter/otlpexporter"
	"go.opentelemetry.io/collector/exporter/otlphttpexporter"
	"go.opentelemetry.io/collector/extension"
	"go.opentelemetry.io/collector/otelcol"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/otlpreceiver"
	"go.uber.org/multierr"
)

// Components returns a set of stripped components used by the
// OpenTelemetry collector built for Lambda env.
func Components() (
	otelcol.Factories,
	error,
) {
	var errs error

	extensions, err := extension.MakeFactoryMap(
		basicauthextension.NewFactory(),
		sigv4authextension.NewFactory(),
	)
	if err != nil {
		errs = multierr.Append(errs, err)
	}

	receivers, err := receiver.MakeFactoryMap(
		otlpreceiver.NewFactory(),
	)
	if err != nil {
		errs = multierr.Append(errs, err)
	}

	exporters, err := exporter.MakeFactoryMap(
		awsxrayexporter.NewFactory(),
		awsemfexporter.NewFactory(),
		prometheusexporter.NewFactory(),
		prometheusremotewriteexporter.NewFactory(),
		loggingexporter.NewFactory(),
		otlpexporter.NewFactory(),
		otlphttpexporter.NewFactory(),
	)
	if err != nil {
		errs = multierr.Append(errs, err)
	}

	factories := otelcol.Factories{
		Extensions: extensions,
		Receivers:  receivers,
		Exporters:  exporters,
	}

	return factories, errs
}
