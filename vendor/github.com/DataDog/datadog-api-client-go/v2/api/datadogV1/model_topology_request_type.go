// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TopologyRequestType Widget request type.
type TopologyRequestType string

// List of TopologyRequestType.
const (
	TOPOLOGYREQUESTTYPE_TOPOLOGY TopologyRequestType = "topology"
)

var allowedTopologyRequestTypeEnumValues = []TopologyRequestType{
	TOPOLOGYREQUESTTYPE_TOPOLOGY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TopologyRequestType) GetAllowedValues() []TopologyRequestType {
	return allowedTopologyRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TopologyRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TopologyRequestType(value)
	return nil
}

// NewTopologyRequestTypeFromValue returns a pointer to a valid TopologyRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTopologyRequestTypeFromValue(v string) (*TopologyRequestType, error) {
	ev := TopologyRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TopologyRequestType: valid values are %v", v, allowedTopologyRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TopologyRequestType) IsValid() bool {
	for _, existing := range allowedTopologyRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TopologyRequestType value.
func (v TopologyRequestType) Ptr() *TopologyRequestType {
	return &v
}
