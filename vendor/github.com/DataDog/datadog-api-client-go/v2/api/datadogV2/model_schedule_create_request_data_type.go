// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleCreateRequestDataType Schedules resource type.
type ScheduleCreateRequestDataType string

// List of ScheduleCreateRequestDataType.
const (
	SCHEDULECREATEREQUESTDATATYPE_SCHEDULES ScheduleCreateRequestDataType = "schedules"
)

var allowedScheduleCreateRequestDataTypeEnumValues = []ScheduleCreateRequestDataType{
	SCHEDULECREATEREQUESTDATATYPE_SCHEDULES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScheduleCreateRequestDataType) GetAllowedValues() []ScheduleCreateRequestDataType {
	return allowedScheduleCreateRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScheduleCreateRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScheduleCreateRequestDataType(value)
	return nil
}

// NewScheduleCreateRequestDataTypeFromValue returns a pointer to a valid ScheduleCreateRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScheduleCreateRequestDataTypeFromValue(v string) (*ScheduleCreateRequestDataType, error) {
	ev := ScheduleCreateRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScheduleCreateRequestDataType: valid values are %v", v, allowedScheduleCreateRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScheduleCreateRequestDataType) IsValid() bool {
	for _, existing := range allowedScheduleCreateRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScheduleCreateRequestDataType value.
func (v ScheduleCreateRequestDataType) Ptr() *ScheduleCreateRequestDataType {
	return &v
}
