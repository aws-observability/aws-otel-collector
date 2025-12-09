// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RulesetRespDataType Ruleset resource type.
type RulesetRespDataType string

// List of RulesetRespDataType.
const (
	RULESETRESPDATATYPE_RULESET RulesetRespDataType = "ruleset"
)

var allowedRulesetRespDataTypeEnumValues = []RulesetRespDataType{
	RULESETRESPDATATYPE_RULESET,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RulesetRespDataType) GetAllowedValues() []RulesetRespDataType {
	return allowedRulesetRespDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RulesetRespDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RulesetRespDataType(value)
	return nil
}

// NewRulesetRespDataTypeFromValue returns a pointer to a valid RulesetRespDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRulesetRespDataTypeFromValue(v string) (*RulesetRespDataType, error) {
	ev := RulesetRespDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RulesetRespDataType: valid values are %v", v, allowedRulesetRespDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RulesetRespDataType) IsValid() bool {
	for _, existing := range allowedRulesetRespDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RulesetRespDataType value.
func (v RulesetRespDataType) Ptr() *RulesetRespDataType {
	return &v
}
