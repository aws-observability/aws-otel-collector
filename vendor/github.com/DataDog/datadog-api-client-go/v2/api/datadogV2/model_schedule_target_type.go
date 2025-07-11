// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleTargetType Indicates that the resource is of type `schedules`.
type ScheduleTargetType string

// List of ScheduleTargetType.
const (
	SCHEDULETARGETTYPE_SCHEDULES ScheduleTargetType = "schedules"
)

var allowedScheduleTargetTypeEnumValues = []ScheduleTargetType{
	SCHEDULETARGETTYPE_SCHEDULES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScheduleTargetType) GetAllowedValues() []ScheduleTargetType {
	return allowedScheduleTargetTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScheduleTargetType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScheduleTargetType(value)
	return nil
}

// NewScheduleTargetTypeFromValue returns a pointer to a valid ScheduleTargetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScheduleTargetTypeFromValue(v string) (*ScheduleTargetType, error) {
	ev := ScheduleTargetType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScheduleTargetType: valid values are %v", v, allowedScheduleTargetTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScheduleTargetType) IsValid() bool {
	for _, existing := range allowedScheduleTargetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScheduleTargetType value.
func (v ScheduleTargetType) Ptr() *ScheduleTargetType {
	return &v
}
