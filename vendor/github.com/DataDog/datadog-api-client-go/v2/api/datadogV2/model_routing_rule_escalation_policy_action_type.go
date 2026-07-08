// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RoutingRuleEscalationPolicyActionType Indicates that the action pages an escalation policy. This action can be set once per routing rule item, and is mutually exclusive with the top-level `policy_id` field on the routing rule.
type RoutingRuleEscalationPolicyActionType string

// List of RoutingRuleEscalationPolicyActionType.
const (
	ROUTINGRULEESCALATIONPOLICYACTIONTYPE_ESCALATION_POLICY RoutingRuleEscalationPolicyActionType = "escalation_policy"
)

var allowedRoutingRuleEscalationPolicyActionTypeEnumValues = []RoutingRuleEscalationPolicyActionType{
	ROUTINGRULEESCALATIONPOLICYACTIONTYPE_ESCALATION_POLICY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RoutingRuleEscalationPolicyActionType) GetAllowedValues() []RoutingRuleEscalationPolicyActionType {
	return allowedRoutingRuleEscalationPolicyActionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RoutingRuleEscalationPolicyActionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RoutingRuleEscalationPolicyActionType(value)
	return nil
}

// NewRoutingRuleEscalationPolicyActionTypeFromValue returns a pointer to a valid RoutingRuleEscalationPolicyActionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRoutingRuleEscalationPolicyActionTypeFromValue(v string) (*RoutingRuleEscalationPolicyActionType, error) {
	ev := RoutingRuleEscalationPolicyActionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RoutingRuleEscalationPolicyActionType: valid values are %v", v, allowedRoutingRuleEscalationPolicyActionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RoutingRuleEscalationPolicyActionType) IsValid() bool {
	for _, existing := range allowedRoutingRuleEscalationPolicyActionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RoutingRuleEscalationPolicyActionType value.
func (v RoutingRuleEscalationPolicyActionType) Ptr() *RoutingRuleEscalationPolicyActionType {
	return &v
}
