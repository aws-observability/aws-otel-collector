// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ShiftIncluded - Included data for shift operations.
type ShiftIncluded struct {
	ScheduleUser *ScheduleUser

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ScheduleUserAsShiftIncluded is a convenience function that returns ScheduleUser wrapped in ShiftIncluded.
func ScheduleUserAsShiftIncluded(v *ScheduleUser) ShiftIncluded {
	return ShiftIncluded{ScheduleUser: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ShiftIncluded) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ScheduleUser
	err = datadog.Unmarshal(data, &obj.ScheduleUser)
	if err == nil {
		if obj.ScheduleUser != nil && obj.ScheduleUser.UnparsedObject == nil {
			jsonScheduleUser, _ := datadog.Marshal(obj.ScheduleUser)
			if string(jsonScheduleUser) == "{}" { // empty struct
				obj.ScheduleUser = nil
			} else {
				match++
			}
		} else {
			obj.ScheduleUser = nil
		}
	} else {
		obj.ScheduleUser = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ScheduleUser = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ShiftIncluded) MarshalJSON() ([]byte, error) {
	if obj.ScheduleUser != nil {
		return datadog.Marshal(&obj.ScheduleUser)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ShiftIncluded) GetActualInstance() interface{} {
	if obj.ScheduleUser != nil {
		return obj.ScheduleUser
	}

	// all schemas are nil
	return nil
}
