// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsDowntimeWeekday A day of the week for a Synthetics downtime recurrence.
type SyntheticsDowntimeWeekday string

// List of SyntheticsDowntimeWeekday.
const (
	SYNTHETICSDOWNTIMEWEEKDAY_MONDAY    SyntheticsDowntimeWeekday = "MO"
	SYNTHETICSDOWNTIMEWEEKDAY_TUESDAY   SyntheticsDowntimeWeekday = "TU"
	SYNTHETICSDOWNTIMEWEEKDAY_WEDNESDAY SyntheticsDowntimeWeekday = "WE"
	SYNTHETICSDOWNTIMEWEEKDAY_THURSDAY  SyntheticsDowntimeWeekday = "TH"
	SYNTHETICSDOWNTIMEWEEKDAY_FRIDAY    SyntheticsDowntimeWeekday = "FR"
	SYNTHETICSDOWNTIMEWEEKDAY_SATURDAY  SyntheticsDowntimeWeekday = "SA"
	SYNTHETICSDOWNTIMEWEEKDAY_SUNDAY    SyntheticsDowntimeWeekday = "SU"
)

var allowedSyntheticsDowntimeWeekdayEnumValues = []SyntheticsDowntimeWeekday{
	SYNTHETICSDOWNTIMEWEEKDAY_MONDAY,
	SYNTHETICSDOWNTIMEWEEKDAY_TUESDAY,
	SYNTHETICSDOWNTIMEWEEKDAY_WEDNESDAY,
	SYNTHETICSDOWNTIMEWEEKDAY_THURSDAY,
	SYNTHETICSDOWNTIMEWEEKDAY_FRIDAY,
	SYNTHETICSDOWNTIMEWEEKDAY_SATURDAY,
	SYNTHETICSDOWNTIMEWEEKDAY_SUNDAY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsDowntimeWeekday) GetAllowedValues() []SyntheticsDowntimeWeekday {
	return allowedSyntheticsDowntimeWeekdayEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsDowntimeWeekday) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsDowntimeWeekday(value)
	return nil
}

// NewSyntheticsDowntimeWeekdayFromValue returns a pointer to a valid SyntheticsDowntimeWeekday
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsDowntimeWeekdayFromValue(v string) (*SyntheticsDowntimeWeekday, error) {
	ev := SyntheticsDowntimeWeekday(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsDowntimeWeekday: valid values are %v", v, allowedSyntheticsDowntimeWeekdayEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsDowntimeWeekday) IsValid() bool {
	for _, existing := range allowedSyntheticsDowntimeWeekdayEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsDowntimeWeekday value.
func (v SyntheticsDowntimeWeekday) Ptr() *SyntheticsDowntimeWeekday {
	return &v
}
