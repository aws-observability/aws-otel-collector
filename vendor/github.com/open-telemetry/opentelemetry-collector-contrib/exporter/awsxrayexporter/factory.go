// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awsxrayexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
<<<<<<< HEAD
=======
	"go.opentelemetry.io/collector/featuregate"
>>>>>>> main

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter/internal/metadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/awsutil"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray/telemetry"
)

<<<<<<< HEAD
=======
var skipTimestampValidationFeatureGate = featuregate.GlobalRegistry().MustRegister(
	"exporter.awsxray.skiptimestampvalidation",
	featuregate.StageAlpha,
	featuregate.WithRegisterDescription("Remove XRay's timestamp validation on first 32 bits of trace ID"),
	featuregate.WithRegisterFromVersion("v0.84.0"))

>>>>>>> main
// NewFactory creates a factory for AWS-Xray exporter.
func NewFactory() exporter.Factory {
	return exporter.NewFactory(
		metadata.Type,
		createDefaultConfig,
		exporter.WithTraces(createTracesExporter, metadata.TracesStability))
}

func createDefaultConfig() component.Config {
	return &Config{
<<<<<<< HEAD
		AWSSessionSettings: awsutil.CreateDefaultSessionConfig(),
=======
		AWSSessionSettings:      awsutil.CreateDefaultSessionConfig(),
		skipTimestampValidation: skipTimestampValidationFeatureGate.IsEnabled(),
>>>>>>> main
	}
}

func createTracesExporter(
	_ context.Context,
	params exporter.CreateSettings,
	cfg component.Config,
) (exporter.Traces, error) {
	eCfg := cfg.(*Config)
	return newTracesExporter(eCfg, params, &awsutil.Conn{}, telemetry.GlobalRegistry())
}
