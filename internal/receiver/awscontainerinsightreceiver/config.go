package awscontainerinsightreceiver

import (
	"time"

	"go.opentelemetry.io/collector/config"
)

// Config defines configuration for aws ecs container metrics receiver.
type Config struct {
	config.ReceiverSettings `mapstructure:",squash"`

	// CollectionInterval is the interval at which metrics should be collected. The default is 60 second.
	CollectionInterval time.Duration `mapstructure:"collection_interval"`

	//ContainerOrchestrator is the type of container orchestration service, e.g. eks or ecs. The default is eks.
	ContainerOrchestrator string `mapstructure:"container_orchestrator"`
}
