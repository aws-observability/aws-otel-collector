package awscontainerinsightreceiver

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
)

// Factory for awscontainerinsightreceiver
const (
	// Key to invoke this receiver
	typeStr = "awscontainerinsightreceiver"

	// Default collection interval. Every 60s the receiver will collect metrics
	defaultCollectionInterval = 60 * time.Second

	// Default container orchestrator service is aws eks
	defaultContainerOrchestrator = "eks"
)

// NewFactory creates a factory for AWS container insight receiver
func NewFactory() component.ReceiverFactory {
	return receiverhelper.NewFactory(
		typeStr,
		createDefaultConfig,
		receiverhelper.WithMetrics(createMetricsReceiver))
}

// createDefaultConfig returns a default config for the receiver.
func createDefaultConfig() config.Receiver {
	return &Config{
		ReceiverSettings: config.ReceiverSettings{
			TypeVal: typeStr,
			NameVal: typeStr,
		},
		CollectionInterval:    defaultCollectionInterval,
		ContainerOrchestrator: defaultContainerOrchestrator,
	}
}

// CreateMetricsReceiver creates an AWS Container Insight receiver.
func createMetricsReceiver(
	_ context.Context,
	params component.ReceiverCreateParams,
	baseCfg config.Receiver,
	consumer consumer.Metrics,
) (component.MetricsReceiver, error) {

	rCfg := baseCfg.(*Config)
	return New(params.Logger, rCfg, consumer)
}
