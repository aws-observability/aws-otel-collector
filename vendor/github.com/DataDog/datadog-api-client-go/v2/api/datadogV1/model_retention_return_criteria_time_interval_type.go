// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RetentionReturnCriteriaTimeIntervalType Type of time interval for return criteria.
type RetentionReturnCriteriaTimeIntervalType string

// List of RetentionReturnCriteriaTimeIntervalType.
const (
	RETENTIONRETURNCRITERIATIMEINTERVALTYPE_FIXED RetentionReturnCriteriaTimeIntervalType = "fixed"
)

var allowedRetentionReturnCriteriaTimeIntervalTypeEnumValues = []RetentionReturnCriteriaTimeIntervalType{
	RETENTIONRETURNCRITERIATIMEINTERVALTYPE_FIXED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RetentionReturnCriteriaTimeIntervalType) GetAllowedValues() []RetentionReturnCriteriaTimeIntervalType {
	return allowedRetentionReturnCriteriaTimeIntervalTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RetentionReturnCriteriaTimeIntervalType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RetentionReturnCriteriaTimeIntervalType(value)
	return nil
}

// NewRetentionReturnCriteriaTimeIntervalTypeFromValue returns a pointer to a valid RetentionReturnCriteriaTimeIntervalType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRetentionReturnCriteriaTimeIntervalTypeFromValue(v string) (*RetentionReturnCriteriaTimeIntervalType, error) {
	ev := RetentionReturnCriteriaTimeIntervalType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RetentionReturnCriteriaTimeIntervalType: valid values are %v", v, allowedRetentionReturnCriteriaTimeIntervalTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RetentionReturnCriteriaTimeIntervalType) IsValid() bool {
	for _, existing := range allowedRetentionReturnCriteriaTimeIntervalTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RetentionReturnCriteriaTimeIntervalType value.
func (v RetentionReturnCriteriaTimeIntervalType) Ptr() *RetentionReturnCriteriaTimeIntervalType {
	return &v
}
