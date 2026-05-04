// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomRuleRevisionDataType Resource type
type CustomRuleRevisionDataType string

// List of CustomRuleRevisionDataType.
const (
	CUSTOMRULEREVISIONDATATYPE_CUSTOM_RULE_REVISION CustomRuleRevisionDataType = "custom_rule_revision"
)

var allowedCustomRuleRevisionDataTypeEnumValues = []CustomRuleRevisionDataType{
	CUSTOMRULEREVISIONDATATYPE_CUSTOM_RULE_REVISION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CustomRuleRevisionDataType) GetAllowedValues() []CustomRuleRevisionDataType {
	return allowedCustomRuleRevisionDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CustomRuleRevisionDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CustomRuleRevisionDataType(value)
	return nil
}

// NewCustomRuleRevisionDataTypeFromValue returns a pointer to a valid CustomRuleRevisionDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCustomRuleRevisionDataTypeFromValue(v string) (*CustomRuleRevisionDataType, error) {
	ev := CustomRuleRevisionDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CustomRuleRevisionDataType: valid values are %v", v, allowedCustomRuleRevisionDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CustomRuleRevisionDataType) IsValid() bool {
	for _, existing := range allowedCustomRuleRevisionDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomRuleRevisionDataType value.
func (v CustomRuleRevisionDataType) Ptr() *CustomRuleRevisionDataType {
	return &v
}
