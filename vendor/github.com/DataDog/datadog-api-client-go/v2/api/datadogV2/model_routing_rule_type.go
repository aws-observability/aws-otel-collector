// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RoutingRuleType Team routing rules resource type.
type RoutingRuleType string

// List of RoutingRuleType.
const (
	ROUTINGRULETYPE_TEAM_ROUTING_RULES RoutingRuleType = "team_routing_rules"
)

var allowedRoutingRuleTypeEnumValues = []RoutingRuleType{
	ROUTINGRULETYPE_TEAM_ROUTING_RULES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RoutingRuleType) GetAllowedValues() []RoutingRuleType {
	return allowedRoutingRuleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RoutingRuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RoutingRuleType(value)
	return nil
}

// NewRoutingRuleTypeFromValue returns a pointer to a valid RoutingRuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRoutingRuleTypeFromValue(v string) (*RoutingRuleType, error) {
	ev := RoutingRuleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RoutingRuleType: valid values are %v", v, allowedRoutingRuleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RoutingRuleType) IsValid() bool {
	for _, existing := range allowedRoutingRuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RoutingRuleType value.
func (v RoutingRuleType) Ptr() *RoutingRuleType {
	return &v
}
