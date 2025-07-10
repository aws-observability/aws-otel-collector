// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RoutingRuleRelationshipsPolicyDataType Indicates that the resource is of type 'policies'.
type RoutingRuleRelationshipsPolicyDataType string

// List of RoutingRuleRelationshipsPolicyDataType.
const (
	ROUTINGRULERELATIONSHIPSPOLICYDATATYPE_POLICIES RoutingRuleRelationshipsPolicyDataType = "policies"
)

var allowedRoutingRuleRelationshipsPolicyDataTypeEnumValues = []RoutingRuleRelationshipsPolicyDataType{
	ROUTINGRULERELATIONSHIPSPOLICYDATATYPE_POLICIES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RoutingRuleRelationshipsPolicyDataType) GetAllowedValues() []RoutingRuleRelationshipsPolicyDataType {
	return allowedRoutingRuleRelationshipsPolicyDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RoutingRuleRelationshipsPolicyDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RoutingRuleRelationshipsPolicyDataType(value)
	return nil
}

// NewRoutingRuleRelationshipsPolicyDataTypeFromValue returns a pointer to a valid RoutingRuleRelationshipsPolicyDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRoutingRuleRelationshipsPolicyDataTypeFromValue(v string) (*RoutingRuleRelationshipsPolicyDataType, error) {
	ev := RoutingRuleRelationshipsPolicyDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RoutingRuleRelationshipsPolicyDataType: valid values are %v", v, allowedRoutingRuleRelationshipsPolicyDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RoutingRuleRelationshipsPolicyDataType) IsValid() bool {
	for _, existing := range allowedRoutingRuleRelationshipsPolicyDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RoutingRuleRelationshipsPolicyDataType value.
func (v RoutingRuleRelationshipsPolicyDataType) Ptr() *RoutingRuleRelationshipsPolicyDataType {
	return &v
}
