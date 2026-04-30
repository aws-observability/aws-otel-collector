// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RetentionReturnCriteriaTimeIntervalUnit Unit of time for retention return criteria interval.
type RetentionReturnCriteriaTimeIntervalUnit string

// List of RetentionReturnCriteriaTimeIntervalUnit.
const (
	RETENTIONRETURNCRITERIATIMEINTERVALUNIT_DAY   RetentionReturnCriteriaTimeIntervalUnit = "day"
	RETENTIONRETURNCRITERIATIMEINTERVALUNIT_WEEK  RetentionReturnCriteriaTimeIntervalUnit = "week"
	RETENTIONRETURNCRITERIATIMEINTERVALUNIT_MONTH RetentionReturnCriteriaTimeIntervalUnit = "month"
)

var allowedRetentionReturnCriteriaTimeIntervalUnitEnumValues = []RetentionReturnCriteriaTimeIntervalUnit{
	RETENTIONRETURNCRITERIATIMEINTERVALUNIT_DAY,
	RETENTIONRETURNCRITERIATIMEINTERVALUNIT_WEEK,
	RETENTIONRETURNCRITERIATIMEINTERVALUNIT_MONTH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RetentionReturnCriteriaTimeIntervalUnit) GetAllowedValues() []RetentionReturnCriteriaTimeIntervalUnit {
	return allowedRetentionReturnCriteriaTimeIntervalUnitEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RetentionReturnCriteriaTimeIntervalUnit) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RetentionReturnCriteriaTimeIntervalUnit(value)
	return nil
}

// NewRetentionReturnCriteriaTimeIntervalUnitFromValue returns a pointer to a valid RetentionReturnCriteriaTimeIntervalUnit
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRetentionReturnCriteriaTimeIntervalUnitFromValue(v string) (*RetentionReturnCriteriaTimeIntervalUnit, error) {
	ev := RetentionReturnCriteriaTimeIntervalUnit(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RetentionReturnCriteriaTimeIntervalUnit: valid values are %v", v, allowedRetentionReturnCriteriaTimeIntervalUnitEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RetentionReturnCriteriaTimeIntervalUnit) IsValid() bool {
	for _, existing := range allowedRetentionReturnCriteriaTimeIntervalUnitEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RetentionReturnCriteriaTimeIntervalUnit value.
func (v RetentionReturnCriteriaTimeIntervalUnit) Ptr() *RetentionReturnCriteriaTimeIntervalUnit {
	return &v
}
