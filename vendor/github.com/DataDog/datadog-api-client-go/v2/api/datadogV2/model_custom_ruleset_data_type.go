// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomRulesetDataType Resource type
type CustomRulesetDataType string

// List of CustomRulesetDataType.
const (
	CUSTOMRULESETDATATYPE_CUSTOM_RULESET CustomRulesetDataType = "custom_ruleset"
)

var allowedCustomRulesetDataTypeEnumValues = []CustomRulesetDataType{
	CUSTOMRULESETDATATYPE_CUSTOM_RULESET,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CustomRulesetDataType) GetAllowedValues() []CustomRulesetDataType {
	return allowedCustomRulesetDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CustomRulesetDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CustomRulesetDataType(value)
	return nil
}

// NewCustomRulesetDataTypeFromValue returns a pointer to a valid CustomRulesetDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCustomRulesetDataTypeFromValue(v string) (*CustomRulesetDataType, error) {
	ev := CustomRulesetDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CustomRulesetDataType: valid values are %v", v, allowedCustomRulesetDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CustomRulesetDataType) IsValid() bool {
	for _, existing := range allowedCustomRulesetDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomRulesetDataType value.
func (v CustomRulesetDataType) Ptr() *CustomRulesetDataType {
	return &v
}
