package awscontainerinsightprocessor

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor/processorhelper"
)

const (
	// The value of "type" key in configuration.
	typeStr = "awscontainerinsightprocessor"
)

var processorCapabilities = component.ProcessorCapabilities{MutatesConsumedData: true}

// NewFactory returns a new factory for the awscontainerinsight processor.
func NewFactory() component.ProcessorFactory {
	return processorhelper.NewFactory(
		typeStr,
		createDefaultConfig,
		processorhelper.WithMetrics(createMetricsProcessor))
}

func createDefaultConfig() config.Processor {
	return &Config{
		ProcessorSettings: config.ProcessorSettings{
			TypeVal: typeStr,
			NameVal: typeStr,
		},
		TagService:      true,
		PrefFullPodName: false,
	}
}

func createMetricsProcessor(
	ctx context.Context,
	params component.ProcessorCreateParams,
	cfg config.Processor,
	nextConsumer consumer.Metrics,
) (component.MetricsProcessor, error) {
	oCfg := cfg.(*Config)
	metricsProcessor := newAwsContainerInsightProcessor(params.Logger, oCfg)

	return processorhelper.NewMetricsProcessor(
		cfg,
		nextConsumer,
		metricsProcessor,
		processorhelper.WithCapabilities(processorCapabilities))
}
