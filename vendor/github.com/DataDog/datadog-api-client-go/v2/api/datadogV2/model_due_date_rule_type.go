// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DueDateRuleType The JSON:API type for due date rules.
type DueDateRuleType string

// List of DueDateRuleType.
const (
	DUEDATERULETYPE_DUE_DATE_RULES DueDateRuleType = "due_date_rules"
)

var allowedDueDateRuleTypeEnumValues = []DueDateRuleType{
	DUEDATERULETYPE_DUE_DATE_RULES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DueDateRuleType) GetAllowedValues() []DueDateRuleType {
	return allowedDueDateRuleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DueDateRuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DueDateRuleType(value)
	return nil
}

// NewDueDateRuleTypeFromValue returns a pointer to a valid DueDateRuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDueDateRuleTypeFromValue(v string) (*DueDateRuleType, error) {
	ev := DueDateRuleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DueDateRuleType: valid values are %v", v, allowedDueDateRuleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DueDateRuleType) IsValid() bool {
	for _, existing := range allowedDueDateRuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DueDateRuleType value.
func (v DueDateRuleType) Ptr() *DueDateRuleType {
	return &v
}
