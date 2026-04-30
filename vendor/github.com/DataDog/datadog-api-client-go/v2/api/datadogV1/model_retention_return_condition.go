// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RetentionReturnCondition Condition for counting user return.
type RetentionReturnCondition string

// List of RetentionReturnCondition.
const (
	RETENTIONRETURNCONDITION_CONVERSION_ON          RetentionReturnCondition = "conversion_on"
	RETENTIONRETURNCONDITION_CONVERSION_ON_OR_AFTER RetentionReturnCondition = "conversion_on_or_after"
)

var allowedRetentionReturnConditionEnumValues = []RetentionReturnCondition{
	RETENTIONRETURNCONDITION_CONVERSION_ON,
	RETENTIONRETURNCONDITION_CONVERSION_ON_OR_AFTER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RetentionReturnCondition) GetAllowedValues() []RetentionReturnCondition {
	return allowedRetentionReturnConditionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RetentionReturnCondition) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RetentionReturnCondition(value)
	return nil
}

// NewRetentionReturnConditionFromValue returns a pointer to a valid RetentionReturnCondition
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRetentionReturnConditionFromValue(v string) (*RetentionReturnCondition, error) {
	ev := RetentionReturnCondition(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RetentionReturnCondition: valid values are %v", v, allowedRetentionReturnConditionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RetentionReturnCondition) IsValid() bool {
	for _, existing := range allowedRetentionReturnConditionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RetentionReturnCondition value.
func (v RetentionReturnCondition) Ptr() *RetentionReturnCondition {
	return &v
}
