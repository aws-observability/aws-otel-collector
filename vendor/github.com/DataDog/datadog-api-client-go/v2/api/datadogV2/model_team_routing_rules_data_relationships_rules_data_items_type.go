// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamRoutingRulesDataRelationshipsRulesDataItemsType Indicates that the resource is of type 'team_routing_rules'.
type TeamRoutingRulesDataRelationshipsRulesDataItemsType string

// List of TeamRoutingRulesDataRelationshipsRulesDataItemsType.
const (
	TEAMROUTINGRULESDATARELATIONSHIPSRULESDATAITEMSTYPE_TEAM_ROUTING_RULES TeamRoutingRulesDataRelationshipsRulesDataItemsType = "team_routing_rules"
)

var allowedTeamRoutingRulesDataRelationshipsRulesDataItemsTypeEnumValues = []TeamRoutingRulesDataRelationshipsRulesDataItemsType{
	TEAMROUTINGRULESDATARELATIONSHIPSRULESDATAITEMSTYPE_TEAM_ROUTING_RULES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TeamRoutingRulesDataRelationshipsRulesDataItemsType) GetAllowedValues() []TeamRoutingRulesDataRelationshipsRulesDataItemsType {
	return allowedTeamRoutingRulesDataRelationshipsRulesDataItemsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TeamRoutingRulesDataRelationshipsRulesDataItemsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TeamRoutingRulesDataRelationshipsRulesDataItemsType(value)
	return nil
}

// NewTeamRoutingRulesDataRelationshipsRulesDataItemsTypeFromValue returns a pointer to a valid TeamRoutingRulesDataRelationshipsRulesDataItemsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTeamRoutingRulesDataRelationshipsRulesDataItemsTypeFromValue(v string) (*TeamRoutingRulesDataRelationshipsRulesDataItemsType, error) {
	ev := TeamRoutingRulesDataRelationshipsRulesDataItemsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TeamRoutingRulesDataRelationshipsRulesDataItemsType: valid values are %v", v, allowedTeamRoutingRulesDataRelationshipsRulesDataItemsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TeamRoutingRulesDataRelationshipsRulesDataItemsType) IsValid() bool {
	for _, existing := range allowedTeamRoutingRulesDataRelationshipsRulesDataItemsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TeamRoutingRulesDataRelationshipsRulesDataItemsType value.
func (v TeamRoutingRulesDataRelationshipsRulesDataItemsType) Ptr() *TeamRoutingRulesDataRelationshipsRulesDataItemsType {
	return &v
}
