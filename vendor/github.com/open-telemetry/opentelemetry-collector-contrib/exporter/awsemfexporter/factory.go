// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awsemfexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsemfexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
<<<<<<< HEAD
=======
	"go.opentelemetry.io/collector/featuregate"
>>>>>>> main
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsemfexporter/internal/metadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/awsutil"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/resourcetotelemetry"
)

<<<<<<< HEAD
=======
var defaultNoRollupfg = featuregate.GlobalRegistry().MustRegister("awsemf.nodimrollupdefault", featuregate.StageAlpha,
	featuregate.WithRegisterFromVersion("v0.83.0"),
	featuregate.WithRegisterDescription("Changes the default AWS EMF Exporter Dimension rollup option to "+
		"NoDimensionRollup"))

>>>>>>> main
// NewFactory creates a factory for AWS EMF exporter.
func NewFactory() exporter.Factory {
	return exporter.NewFactory(
		metadata.Type,
		createDefaultConfig,
		exporter.WithMetrics(createMetricsExporter, metadata.MetricsStability))
}

// CreateDefaultConfig creates the default configuration for exporter.
func createDefaultConfig() component.Config {
<<<<<<< HEAD
=======
	var defaultDimensionRollupOption string
	if defaultNoRollupfg.IsEnabled() {
		defaultDimensionRollupOption = "NoDimensionRollup"
	} else {
		defaultDimensionRollupOption = "ZeroAndSingleDimensionRollup"
	}
>>>>>>> main
	return &Config{
		AWSSessionSettings:              awsutil.CreateDefaultSessionConfig(),
		LogGroupName:                    "",
		LogStreamName:                   "",
		Namespace:                       "",
<<<<<<< HEAD
		DimensionRollupOption:           "ZeroAndSingleDimensionRollup",
=======
		DimensionRollupOption:           defaultDimensionRollupOption,
>>>>>>> main
		Version:                         "1",
		RetainInitialValueOfDeltaMetric: false,
		OutputDestination:               "cloudwatch",
		logger:                          zap.NewNop(),
	}
}

// createMetricsExporter creates a metrics exporter based on this config.
func createMetricsExporter(ctx context.Context, params exporter.CreateSettings, config component.Config) (exporter.Metrics, error) {
	expCfg := config.(*Config)

	emfExp, err := newEmfExporter(expCfg, params)
	if err != nil {
		return nil, err
	}

	exporter, err := exporterhelper.NewMetricsExporter(
		ctx,
		params,
		config,
		emfExp.pushMetricsData,
		exporterhelper.WithShutdown(emfExp.shutdown),
	)
	if err != nil {
		return nil, err
	}

	return resourcetotelemetry.WrapMetricsExporter(expCfg.ResourceToTelemetrySettings, exporter), nil
}
