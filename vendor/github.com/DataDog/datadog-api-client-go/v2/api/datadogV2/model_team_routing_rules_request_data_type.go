// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamRoutingRulesRequestDataType Team routing rules resource type.
type TeamRoutingRulesRequestDataType string

// List of TeamRoutingRulesRequestDataType.
const (
	TEAMROUTINGRULESREQUESTDATATYPE_TEAM_ROUTING_RULES TeamRoutingRulesRequestDataType = "team_routing_rules"
)

var allowedTeamRoutingRulesRequestDataTypeEnumValues = []TeamRoutingRulesRequestDataType{
	TEAMROUTINGRULESREQUESTDATATYPE_TEAM_ROUTING_RULES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TeamRoutingRulesRequestDataType) GetAllowedValues() []TeamRoutingRulesRequestDataType {
	return allowedTeamRoutingRulesRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TeamRoutingRulesRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TeamRoutingRulesRequestDataType(value)
	return nil
}

// NewTeamRoutingRulesRequestDataTypeFromValue returns a pointer to a valid TeamRoutingRulesRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTeamRoutingRulesRequestDataTypeFromValue(v string) (*TeamRoutingRulesRequestDataType, error) {
	ev := TeamRoutingRulesRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TeamRoutingRulesRequestDataType: valid values are %v", v, allowedTeamRoutingRulesRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TeamRoutingRulesRequestDataType) IsValid() bool {
	for _, existing := range allowedTeamRoutingRulesRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TeamRoutingRulesRequestDataType value.
func (v TeamRoutingRulesRequestDataType) Ptr() *TeamRoutingRulesRequestDataType {
	return &v
}
