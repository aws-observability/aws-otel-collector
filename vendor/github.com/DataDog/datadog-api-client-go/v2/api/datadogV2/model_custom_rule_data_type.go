// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomRuleDataType Resource type
type CustomRuleDataType string

// List of CustomRuleDataType.
const (
	CUSTOMRULEDATATYPE_CUSTOM_RULE CustomRuleDataType = "custom_rule"
)

var allowedCustomRuleDataTypeEnumValues = []CustomRuleDataType{
	CUSTOMRULEDATATYPE_CUSTOM_RULE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CustomRuleDataType) GetAllowedValues() []CustomRuleDataType {
	return allowedCustomRuleDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CustomRuleDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CustomRuleDataType(value)
	return nil
}

// NewCustomRuleDataTypeFromValue returns a pointer to a valid CustomRuleDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCustomRuleDataTypeFromValue(v string) (*CustomRuleDataType, error) {
	ev := CustomRuleDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CustomRuleDataType: valid values are %v", v, allowedCustomRuleDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CustomRuleDataType) IsValid() bool {
	for _, existing := range allowedCustomRuleDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomRuleDataType value.
func (v CustomRuleDataType) Ptr() *CustomRuleDataType {
	return &v
}
