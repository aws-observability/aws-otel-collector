// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RetentionCohortCriteriaTimeIntervalType Type of time interval for cohort criteria.
type RetentionCohortCriteriaTimeIntervalType string

// List of RetentionCohortCriteriaTimeIntervalType.
const (
	RETENTIONCOHORTCRITERIATIMEINTERVALTYPE_CALENDAR RetentionCohortCriteriaTimeIntervalType = "calendar"
)

var allowedRetentionCohortCriteriaTimeIntervalTypeEnumValues = []RetentionCohortCriteriaTimeIntervalType{
	RETENTIONCOHORTCRITERIATIMEINTERVALTYPE_CALENDAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RetentionCohortCriteriaTimeIntervalType) GetAllowedValues() []RetentionCohortCriteriaTimeIntervalType {
	return allowedRetentionCohortCriteriaTimeIntervalTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RetentionCohortCriteriaTimeIntervalType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RetentionCohortCriteriaTimeIntervalType(value)
	return nil
}

// NewRetentionCohortCriteriaTimeIntervalTypeFromValue returns a pointer to a valid RetentionCohortCriteriaTimeIntervalType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRetentionCohortCriteriaTimeIntervalTypeFromValue(v string) (*RetentionCohortCriteriaTimeIntervalType, error) {
	ev := RetentionCohortCriteriaTimeIntervalType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RetentionCohortCriteriaTimeIntervalType: valid values are %v", v, allowedRetentionCohortCriteriaTimeIntervalTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RetentionCohortCriteriaTimeIntervalType) IsValid() bool {
	for _, existing := range allowedRetentionCohortCriteriaTimeIntervalTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RetentionCohortCriteriaTimeIntervalType value.
func (v RetentionCohortCriteriaTimeIntervalType) Ptr() *RetentionCohortCriteriaTimeIntervalType {
	return &v
}
