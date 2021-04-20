package common

const (
	EKS            = "eks"
	KubeSecurePort = "10250"
	BearerToken    = "/var/run/secrets/kubernetes.io/serviceaccount/token"
	CAFile         = "/var/run/secrets/kubernetes.io/serviceaccount/ca.crt"

	Kubernetes       = "kubernetes"
	K8sNamespace     = "Namespace"
	PodIdKey         = "PodId"
	PodNameKey       = "PodName"
	K8sPodNameKey    = "K8sPodName"
	ContainerNamekey = "ContainerName"
	ContainerIdkey   = "ContainerId"
	PodOwnersKey     = "PodOwners"

	PodStatus       = "pod_status"
	ContainerStatus = "container_status"

	ContainerStatusReason          = "container_status_reason"
	ContainerLastTerminationReason = "container_last_termination_reason"

	Timestamp = "Timestamp"

	//Pod Owners
	ReplicaSet            = "ReplicaSet"
	ReplicationController = "ReplicationController"
	StatefulSet           = "StatefulSet"
	DaemonSet             = "DaemonSet"
	Deployment            = "Deployment"
	Job                   = "Job"
	CronJob               = "CronJob"
)
