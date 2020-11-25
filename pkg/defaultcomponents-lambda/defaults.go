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
	 "go.opentelemetry.io/collector/exporter/loggingexporter"
	 "go.opentelemetry.io/collector/exporter/otlphttpexporter"
	 "go.opentelemetry.io/collector/exporter/prometheusexporter"
	 "go.opentelemetry.io/collector/receiver/otlpreceiver"
 )
 
 // LambdaComponents returns a set of stripped components used by the
 // OpenTelemetry collector built for Lambda env.
 func LambdaComponents() (
	 component.Factories,
	 error,
 ) {
	 var errs []error
 
	 receivers, err := component.MakeReceiverFactoryMap(
		 otlpreceiver.NewFactory(),
	 )
	 if err != nil {
		 errs = append(errs, err)
	 }
 
	 exporters, err := component.MakeExporterFactoryMap(
		 awsxrayexporter.NewFactory(),
		 awsemfexporter.NewFactory(),
		 prometheusexporter.NewFactory(),
		 loggingexporter.NewFactory(),
		 otlphttpexporter.NewFactory(),
	 )
	 if err != nil {
		 errs = append(errs, err)
	 }
 
	 factories := component.Factories{
		 Receivers:  receivers,
		 Exporters:  exporters,
	 }
 
	 return factories, componenterror.CombineErrors(errs)
 }