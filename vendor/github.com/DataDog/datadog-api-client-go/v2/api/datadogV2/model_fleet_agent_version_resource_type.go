// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetAgentVersionResourceType The type of Agent version resource.
type FleetAgentVersionResourceType string

// List of FleetAgentVersionResourceType.
const (
	FLEETAGENTVERSIONRESOURCETYPE_AGENT_VERSION FleetAgentVersionResourceType = "agent_version"
)

var allowedFleetAgentVersionResourceTypeEnumValues = []FleetAgentVersionResourceType{
	FLEETAGENTVERSIONRESOURCETYPE_AGENT_VERSION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FleetAgentVersionResourceType) GetAllowedValues() []FleetAgentVersionResourceType {
	return allowedFleetAgentVersionResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FleetAgentVersionResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FleetAgentVersionResourceType(value)
	return nil
}

// NewFleetAgentVersionResourceTypeFromValue returns a pointer to a valid FleetAgentVersionResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFleetAgentVersionResourceTypeFromValue(v string) (*FleetAgentVersionResourceType, error) {
	ev := FleetAgentVersionResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FleetAgentVersionResourceType: valid values are %v", v, allowedFleetAgentVersionResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FleetAgentVersionResourceType) IsValid() bool {
	for _, existing := range allowedFleetAgentVersionResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FleetAgentVersionResourceType value.
func (v FleetAgentVersionResourceType) Ptr() *FleetAgentVersionResourceType {
	return &v
}
