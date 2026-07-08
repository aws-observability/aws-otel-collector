// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RulesetStatusRespDataType Ruleset status resource type.
type RulesetStatusRespDataType string

// List of RulesetStatusRespDataType.
const (
	RULESETSTATUSRESPDATATYPE_RULESET_STATUS RulesetStatusRespDataType = "ruleset_status"
)

var allowedRulesetStatusRespDataTypeEnumValues = []RulesetStatusRespDataType{
	RULESETSTATUSRESPDATATYPE_RULESET_STATUS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RulesetStatusRespDataType) GetAllowedValues() []RulesetStatusRespDataType {
	return allowedRulesetStatusRespDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RulesetStatusRespDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RulesetStatusRespDataType(value)
	return nil
}

// NewRulesetStatusRespDataTypeFromValue returns a pointer to a valid RulesetStatusRespDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRulesetStatusRespDataTypeFromValue(v string) (*RulesetStatusRespDataType, error) {
	ev := RulesetStatusRespDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RulesetStatusRespDataType: valid values are %v", v, allowedRulesetStatusRespDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RulesetStatusRespDataType) IsValid() bool {
	for _, existing := range allowedRulesetStatusRespDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RulesetStatusRespDataType value.
func (v RulesetStatusRespDataType) Ptr() *RulesetStatusRespDataType {
	return &v
}
