// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// Weekday A day of the week.
type Weekday string

// List of Weekday.
const (
	WEEKDAY_MONDAY    Weekday = "monday"
	WEEKDAY_TUESDAY   Weekday = "tuesday"
	WEEKDAY_WEDNESDAY Weekday = "wednesday"
	WEEKDAY_THURSDAY  Weekday = "thursday"
	WEEKDAY_FRIDAY    Weekday = "friday"
	WEEKDAY_SATURDAY  Weekday = "saturday"
	WEEKDAY_SUNDAY    Weekday = "sunday"
)

var allowedWeekdayEnumValues = []Weekday{
	WEEKDAY_MONDAY,
	WEEKDAY_TUESDAY,
	WEEKDAY_WEDNESDAY,
	WEEKDAY_THURSDAY,
	WEEKDAY_FRIDAY,
	WEEKDAY_SATURDAY,
	WEEKDAY_SUNDAY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *Weekday) GetAllowedValues() []Weekday {
	return allowedWeekdayEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *Weekday) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = Weekday(value)
	return nil
}

// NewWeekdayFromValue returns a pointer to a valid Weekday
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWeekdayFromValue(v string) (*Weekday, error) {
	ev := Weekday(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for Weekday: valid values are %v", v, allowedWeekdayEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v Weekday) IsValid() bool {
	for _, existing := range allowedWeekdayEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Weekday value.
func (v Weekday) Ptr() *Weekday {
	return &v
}
