// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamRoutingRulesDataType Team routing rules resource type.
type TeamRoutingRulesDataType string

// List of TeamRoutingRulesDataType.
const (
	TEAMROUTINGRULESDATATYPE_TEAM_ROUTING_RULES TeamRoutingRulesDataType = "team_routing_rules"
)

var allowedTeamRoutingRulesDataTypeEnumValues = []TeamRoutingRulesDataType{
	TEAMROUTINGRULESDATATYPE_TEAM_ROUTING_RULES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TeamRoutingRulesDataType) GetAllowedValues() []TeamRoutingRulesDataType {
	return allowedTeamRoutingRulesDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TeamRoutingRulesDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TeamRoutingRulesDataType(value)
	return nil
}

// NewTeamRoutingRulesDataTypeFromValue returns a pointer to a valid TeamRoutingRulesDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTeamRoutingRulesDataTypeFromValue(v string) (*TeamRoutingRulesDataType, error) {
	ev := TeamRoutingRulesDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TeamRoutingRulesDataType: valid values are %v", v, allowedTeamRoutingRulesDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TeamRoutingRulesDataType) IsValid() bool {
	for _, existing := range allowedTeamRoutingRulesDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TeamRoutingRulesDataType value.
func (v TeamRoutingRulesDataType) Ptr() *TeamRoutingRulesDataType {
	return &v
}
