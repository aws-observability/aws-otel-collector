// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsDowntimeWeekdayPosition The position of a weekday within a month for a monthly Synthetics downtime recurrence. `1` through `4` select the first through fourth occurrence of the weekday in the month, and `-1` selects the last occurrence.
type SyntheticsDowntimeWeekdayPosition int64

// List of SyntheticsDowntimeWeekdayPosition.
const (
	SYNTHETICSDOWNTIMEWEEKDAYPOSITION_FIRST  SyntheticsDowntimeWeekdayPosition = 1
	SYNTHETICSDOWNTIMEWEEKDAYPOSITION_SECOND SyntheticsDowntimeWeekdayPosition = 2
	SYNTHETICSDOWNTIMEWEEKDAYPOSITION_THIRD  SyntheticsDowntimeWeekdayPosition = 3
	SYNTHETICSDOWNTIMEWEEKDAYPOSITION_FOURTH SyntheticsDowntimeWeekdayPosition = 4
	SYNTHETICSDOWNTIMEWEEKDAYPOSITION_LAST   SyntheticsDowntimeWeekdayPosition = -1
)

var allowedSyntheticsDowntimeWeekdayPositionEnumValues = []SyntheticsDowntimeWeekdayPosition{
	SYNTHETICSDOWNTIMEWEEKDAYPOSITION_FIRST,
	SYNTHETICSDOWNTIMEWEEKDAYPOSITION_SECOND,
	SYNTHETICSDOWNTIMEWEEKDAYPOSITION_THIRD,
	SYNTHETICSDOWNTIMEWEEKDAYPOSITION_FOURTH,
	SYNTHETICSDOWNTIMEWEEKDAYPOSITION_LAST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsDowntimeWeekdayPosition) GetAllowedValues() []SyntheticsDowntimeWeekdayPosition {
	return allowedSyntheticsDowntimeWeekdayPositionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsDowntimeWeekdayPosition) UnmarshalJSON(src []byte) error {
	var value int64
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsDowntimeWeekdayPosition(value)
	return nil
}

// NewSyntheticsDowntimeWeekdayPositionFromValue returns a pointer to a valid SyntheticsDowntimeWeekdayPosition
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsDowntimeWeekdayPositionFromValue(v int64) (*SyntheticsDowntimeWeekdayPosition, error) {
	ev := SyntheticsDowntimeWeekdayPosition(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsDowntimeWeekdayPosition: valid values are %v", v, allowedSyntheticsDowntimeWeekdayPositionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsDowntimeWeekdayPosition) IsValid() bool {
	for _, existing := range allowedSyntheticsDowntimeWeekdayPositionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsDowntimeWeekdayPosition value.
func (v SyntheticsDowntimeWeekdayPosition) Ptr() *SyntheticsDowntimeWeekdayPosition {
	return &v
}
