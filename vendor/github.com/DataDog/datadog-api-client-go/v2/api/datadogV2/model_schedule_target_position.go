// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleTargetPosition Specifies the position of a schedule target (example `previous`, `current`, or `next`).
type ScheduleTargetPosition string

// List of ScheduleTargetPosition.
const (
	SCHEDULETARGETPOSITION_PREVIOUS ScheduleTargetPosition = "previous"
	SCHEDULETARGETPOSITION_CURRENT  ScheduleTargetPosition = "current"
	SCHEDULETARGETPOSITION_NEXT     ScheduleTargetPosition = "next"
)

var allowedScheduleTargetPositionEnumValues = []ScheduleTargetPosition{
	SCHEDULETARGETPOSITION_PREVIOUS,
	SCHEDULETARGETPOSITION_CURRENT,
	SCHEDULETARGETPOSITION_NEXT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScheduleTargetPosition) GetAllowedValues() []ScheduleTargetPosition {
	return allowedScheduleTargetPositionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScheduleTargetPosition) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScheduleTargetPosition(value)
	return nil
}

// NewScheduleTargetPositionFromValue returns a pointer to a valid ScheduleTargetPosition
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScheduleTargetPositionFromValue(v string) (*ScheduleTargetPosition, error) {
	ev := ScheduleTargetPosition(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScheduleTargetPosition: valid values are %v", v, allowedScheduleTargetPositionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScheduleTargetPosition) IsValid() bool {
	for _, existing := range allowedScheduleTargetPositionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScheduleTargetPosition value.
func (v ScheduleTargetPosition) Ptr() *ScheduleTargetPosition {
	return &v
}
