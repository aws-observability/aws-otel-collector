// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetAgentInfoResourceType The type of Agent info resource.
type FleetAgentInfoResourceType string

// List of FleetAgentInfoResourceType.
const (
	FLEETAGENTINFORESOURCETYPE_DATADOG_AGENT_KEY FleetAgentInfoResourceType = "datadog_agent_key"
)

var allowedFleetAgentInfoResourceTypeEnumValues = []FleetAgentInfoResourceType{
	FLEETAGENTINFORESOURCETYPE_DATADOG_AGENT_KEY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FleetAgentInfoResourceType) GetAllowedValues() []FleetAgentInfoResourceType {
	return allowedFleetAgentInfoResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FleetAgentInfoResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FleetAgentInfoResourceType(value)
	return nil
}

// NewFleetAgentInfoResourceTypeFromValue returns a pointer to a valid FleetAgentInfoResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFleetAgentInfoResourceTypeFromValue(v string) (*FleetAgentInfoResourceType, error) {
	ev := FleetAgentInfoResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FleetAgentInfoResourceType: valid values are %v", v, allowedFleetAgentInfoResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FleetAgentInfoResourceType) IsValid() bool {
	for _, existing := range allowedFleetAgentInfoResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FleetAgentInfoResourceType value.
func (v FleetAgentInfoResourceType) Ptr() *FleetAgentInfoResourceType {
	return &v
}
