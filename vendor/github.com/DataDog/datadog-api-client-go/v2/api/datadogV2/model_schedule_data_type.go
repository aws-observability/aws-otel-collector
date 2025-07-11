// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleDataType Schedules resource type.
type ScheduleDataType string

// List of ScheduleDataType.
const (
	SCHEDULEDATATYPE_SCHEDULES ScheduleDataType = "schedules"
)

var allowedScheduleDataTypeEnumValues = []ScheduleDataType{
	SCHEDULEDATATYPE_SCHEDULES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScheduleDataType) GetAllowedValues() []ScheduleDataType {
	return allowedScheduleDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScheduleDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScheduleDataType(value)
	return nil
}

// NewScheduleDataTypeFromValue returns a pointer to a valid ScheduleDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScheduleDataTypeFromValue(v string) (*ScheduleDataType, error) {
	ev := ScheduleDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScheduleDataType: valid values are %v", v, allowedScheduleDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScheduleDataType) IsValid() bool {
	for _, existing := range allowedScheduleDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScheduleDataType value.
func (v ScheduleDataType) Ptr() *ScheduleDataType {
	return &v
}
