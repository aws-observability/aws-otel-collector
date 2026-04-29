// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StatusPageArrayIncluded - An included resource related to a status page.
type StatusPageArrayIncluded struct {
	StatusPagesUser *StatusPagesUser

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// StatusPagesUserAsStatusPageArrayIncluded is a convenience function that returns StatusPagesUser wrapped in StatusPageArrayIncluded.
func StatusPagesUserAsStatusPageArrayIncluded(v *StatusPagesUser) StatusPageArrayIncluded {
	return StatusPageArrayIncluded{StatusPagesUser: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *StatusPageArrayIncluded) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into StatusPagesUser
	err = datadog.Unmarshal(data, &obj.StatusPagesUser)
	if err == nil {
		if obj.StatusPagesUser != nil && obj.StatusPagesUser.UnparsedObject == nil {
			jsonStatusPagesUser, _ := datadog.Marshal(obj.StatusPagesUser)
			if string(jsonStatusPagesUser) == "{}" { // empty struct
				obj.StatusPagesUser = nil
			} else {
				match++
			}
		} else {
			obj.StatusPagesUser = nil
		}
	} else {
		obj.StatusPagesUser = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.StatusPagesUser = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj StatusPageArrayIncluded) MarshalJSON() ([]byte, error) {
	if obj.StatusPagesUser != nil {
		return datadog.Marshal(&obj.StatusPagesUser)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *StatusPageArrayIncluded) GetActualInstance() interface{} {
	if obj.StatusPagesUser != nil {
		return obj.StatusPagesUser
	}

	// all schemas are nil
	return nil
}
