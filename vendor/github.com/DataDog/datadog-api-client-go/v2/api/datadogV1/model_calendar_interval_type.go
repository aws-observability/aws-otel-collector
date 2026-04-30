// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CalendarIntervalType Type of calendar interval.
type CalendarIntervalType string

// List of CalendarIntervalType.
const (
	CALENDARINTERVALTYPE_DAY     CalendarIntervalType = "day"
	CALENDARINTERVALTYPE_WEEK    CalendarIntervalType = "week"
	CALENDARINTERVALTYPE_MONTH   CalendarIntervalType = "month"
	CALENDARINTERVALTYPE_YEAR    CalendarIntervalType = "year"
	CALENDARINTERVALTYPE_QUARTER CalendarIntervalType = "quarter"
	CALENDARINTERVALTYPE_MINUTE  CalendarIntervalType = "minute"
	CALENDARINTERVALTYPE_HOUR    CalendarIntervalType = "hour"
)

var allowedCalendarIntervalTypeEnumValues = []CalendarIntervalType{
	CALENDARINTERVALTYPE_DAY,
	CALENDARINTERVALTYPE_WEEK,
	CALENDARINTERVALTYPE_MONTH,
	CALENDARINTERVALTYPE_YEAR,
	CALENDARINTERVALTYPE_QUARTER,
	CALENDARINTERVALTYPE_MINUTE,
	CALENDARINTERVALTYPE_HOUR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CalendarIntervalType) GetAllowedValues() []CalendarIntervalType {
	return allowedCalendarIntervalTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CalendarIntervalType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CalendarIntervalType(value)
	return nil
}

// NewCalendarIntervalTypeFromValue returns a pointer to a valid CalendarIntervalType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCalendarIntervalTypeFromValue(v string) (*CalendarIntervalType, error) {
	ev := CalendarIntervalType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CalendarIntervalType: valid values are %v", v, allowedCalendarIntervalTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CalendarIntervalType) IsValid() bool {
	for _, existing := range allowedCalendarIntervalTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CalendarIntervalType value.
func (v CalendarIntervalType) Ptr() *CalendarIntervalType {
	return &v
}
