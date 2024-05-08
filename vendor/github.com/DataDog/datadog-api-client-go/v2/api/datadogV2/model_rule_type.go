// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RuleType The JSON:API type for scorecard rules.
type RuleType string

// List of RuleType.
const (
	RULETYPE_RULE RuleType = "rule"
)

var allowedRuleTypeEnumValues = []RuleType{
	RULETYPE_RULE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RuleType) GetAllowedValues() []RuleType {
	return allowedRuleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RuleType(value)
	return nil
}

// NewRuleTypeFromValue returns a pointer to a valid RuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRuleTypeFromValue(v string) (*RuleType, error) {
	ev := RuleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RuleType: valid values are %v", v, allowedRuleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RuleType) IsValid() bool {
	for _, existing := range allowedRuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RuleType value.
func (v RuleType) Ptr() *RuleType {
	return &v
}
