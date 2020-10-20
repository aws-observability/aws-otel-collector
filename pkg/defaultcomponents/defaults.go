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

package defaultcomponents // import "aws-observability.io/collector/defaultcomponents

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsemfexporter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenterror"
	"go.opentelemetry.io/collector/exporter/fileexporter"
	"go.opentelemetry.io/collector/exporter/loggingexporter"
	"go.opentelemetry.io/collector/exporter/otlpexporter"
	"go.opentelemetry.io/collector/exporter/prometheusexporter"
	"go.opentelemetry.io/collector/receiver/otlpreceiver"
	"go.opentelemetry.io/collector/receiver/prometheusreceiver"
	"go.opentelemetry.io/collector/service/defaultcomponents"
)

// Components register OTel components for aws-otel-collector distribution
func Components() (component.Factories, error) {
	errs := []error{}
	factories, err := defaultcomponents.Components()
	if err != nil {
		return component.Factories{}, err
	}

	// enable the selected receivers
	factories.Receivers, err = component.MakeReceiverFactoryMap(
		prometheusreceiver.NewFactory(),
		otlpreceiver.NewFactory(),
	)
	if err != nil {
		errs = append(errs, err)
	}

	// enable the selected exporters
	factories.Exporters, err = component.MakeExporterFactoryMap(
		awsxrayexporter.NewFactory(),
		awsemfexporter.NewFactory(),
		prometheusexporter.NewFactory(),
		loggingexporter.NewFactory(),
		fileexporter.NewFactory(),
		otlpexporter.NewFactory(),
	)
	if err != nil {
		errs = append(errs, err)
	}

	return factories, componenterror.CombineErrors(errs)
}
