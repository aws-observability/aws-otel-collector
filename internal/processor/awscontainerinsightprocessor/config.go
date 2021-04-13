package awscontainerinsightprocessor

import (
	"go.opentelemetry.io/collector/config"
)

// Config defines configuration for k8s attributes processor.
type Config struct {
	config.ProcessorSettings `mapstructure:",squash"`

	//Whether to add the associated service name as attribute. The default is true
	TagService bool `mapstructure:"add_service_as_attribute"`

	//The "PodName" attribute is set based on the name of those well-known controllers like Daemonset, Job, ReplicaSet, ReplicationController, ...
	//If it can not be set that way and PrefFullPodName is true, the "PodName" attribute is set to the pod's own name.
	PrefFullPodName bool `mapstructure:"prefer_full_pod_name"`
}
