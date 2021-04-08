package k8sclient

import (
	v1 "k8s.io/api/core/v1"
)

type nodeInfo struct {
	conditions []*nodeCondition
}

type nodeCondition struct {
	Type   v1.NodeConditionType
	Status v1.ConditionStatus
}
