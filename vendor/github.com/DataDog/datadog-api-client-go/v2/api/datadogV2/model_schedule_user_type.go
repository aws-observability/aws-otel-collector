// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleUserType Users resource type.
type ScheduleUserType string

// List of ScheduleUserType.
const (
	SCHEDULEUSERTYPE_USERS ScheduleUserType = "users"
)

var allowedScheduleUserTypeEnumValues = []ScheduleUserType{
	SCHEDULEUSERTYPE_USERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScheduleUserType) GetAllowedValues() []ScheduleUserType {
	return allowedScheduleUserTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScheduleUserType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScheduleUserType(value)
	return nil
}

// NewScheduleUserTypeFromValue returns a pointer to a valid ScheduleUserType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScheduleUserTypeFromValue(v string) (*ScheduleUserType, error) {
	ev := ScheduleUserType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScheduleUserType: valid values are %v", v, allowedScheduleUserTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScheduleUserType) IsValid() bool {
	for _, existing := range allowedScheduleUserTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScheduleUserType value.
func (v ScheduleUserType) Ptr() *ScheduleUserType {
	return &v
}
