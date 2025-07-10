// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleUpdateRequestDataType Schedules resource type.
type ScheduleUpdateRequestDataType string

// List of ScheduleUpdateRequestDataType.
const (
	SCHEDULEUPDATEREQUESTDATATYPE_SCHEDULES ScheduleUpdateRequestDataType = "schedules"
)

var allowedScheduleUpdateRequestDataTypeEnumValues = []ScheduleUpdateRequestDataType{
	SCHEDULEUPDATEREQUESTDATATYPE_SCHEDULES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScheduleUpdateRequestDataType) GetAllowedValues() []ScheduleUpdateRequestDataType {
	return allowedScheduleUpdateRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScheduleUpdateRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScheduleUpdateRequestDataType(value)
	return nil
}

// NewScheduleUpdateRequestDataTypeFromValue returns a pointer to a valid ScheduleUpdateRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScheduleUpdateRequestDataTypeFromValue(v string) (*ScheduleUpdateRequestDataType, error) {
	ev := ScheduleUpdateRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScheduleUpdateRequestDataType: valid values are %v", v, allowedScheduleUpdateRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScheduleUpdateRequestDataType) IsValid() bool {
	for _, existing := range allowedScheduleUpdateRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScheduleUpdateRequestDataType value.
func (v ScheduleUpdateRequestDataType) Ptr() *ScheduleUpdateRequestDataType {
	return &v
}
