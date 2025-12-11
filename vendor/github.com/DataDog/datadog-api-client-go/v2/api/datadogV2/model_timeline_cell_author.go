// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimelineCellAuthor - author of the timeline cell
type TimelineCellAuthor struct {
	TimelineCellAuthorUser *TimelineCellAuthorUser

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// TimelineCellAuthorUserAsTimelineCellAuthor is a convenience function that returns TimelineCellAuthorUser wrapped in TimelineCellAuthor.
func TimelineCellAuthorUserAsTimelineCellAuthor(v *TimelineCellAuthorUser) TimelineCellAuthor {
	return TimelineCellAuthor{TimelineCellAuthorUser: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *TimelineCellAuthor) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TimelineCellAuthorUser
	err = datadog.Unmarshal(data, &obj.TimelineCellAuthorUser)
	if err == nil {
		if obj.TimelineCellAuthorUser != nil && obj.TimelineCellAuthorUser.UnparsedObject == nil {
			jsonTimelineCellAuthorUser, _ := datadog.Marshal(obj.TimelineCellAuthorUser)
			if string(jsonTimelineCellAuthorUser) == "{}" && string(data) != "{}" { // empty struct
				obj.TimelineCellAuthorUser = nil
			} else {
				match++
			}
		} else {
			obj.TimelineCellAuthorUser = nil
		}
	} else {
		obj.TimelineCellAuthorUser = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.TimelineCellAuthorUser = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj TimelineCellAuthor) MarshalJSON() ([]byte, error) {
	if obj.TimelineCellAuthorUser != nil {
		return datadog.Marshal(&obj.TimelineCellAuthorUser)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *TimelineCellAuthor) GetActualInstance() interface{} {
	if obj.TimelineCellAuthorUser != nil {
		return obj.TimelineCellAuthorUser
	}

	// all schemas are nil
	return nil
}
