package k8sclient

import (
	v1 "k8s.io/api/core/v1"
)

type podInfo struct {
	namespace string
	phase     v1.PodPhase
}
