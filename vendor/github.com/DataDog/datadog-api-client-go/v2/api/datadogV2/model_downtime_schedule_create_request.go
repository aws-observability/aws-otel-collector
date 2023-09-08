// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/goccy/go-json"
)

// DowntimeScheduleCreateRequest - Schedule for the downtime.
type DowntimeScheduleCreateRequest struct {
	DowntimeScheduleRecurrencesCreateRequest   *DowntimeScheduleRecurrencesCreateRequest
	DowntimeScheduleOneTimeCreateUpdateRequest *DowntimeScheduleOneTimeCreateUpdateRequest

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// DowntimeScheduleRecurrencesCreateRequestAsDowntimeScheduleCreateRequest is a convenience function that returns DowntimeScheduleRecurrencesCreateRequest wrapped in DowntimeScheduleCreateRequest.
func DowntimeScheduleRecurrencesCreateRequestAsDowntimeScheduleCreateRequest(v *DowntimeScheduleRecurrencesCreateRequest) DowntimeScheduleCreateRequest {
	return DowntimeScheduleCreateRequest{DowntimeScheduleRecurrencesCreateRequest: v}
}

// DowntimeScheduleOneTimeCreateUpdateRequestAsDowntimeScheduleCreateRequest is a convenience function that returns DowntimeScheduleOneTimeCreateUpdateRequest wrapped in DowntimeScheduleCreateRequest.
func DowntimeScheduleOneTimeCreateUpdateRequestAsDowntimeScheduleCreateRequest(v *DowntimeScheduleOneTimeCreateUpdateRequest) DowntimeScheduleCreateRequest {
	return DowntimeScheduleCreateRequest{DowntimeScheduleOneTimeCreateUpdateRequest: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *DowntimeScheduleCreateRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DowntimeScheduleRecurrencesCreateRequest
	err = json.Unmarshal(data, &obj.DowntimeScheduleRecurrencesCreateRequest)
	if err == nil {
		if obj.DowntimeScheduleRecurrencesCreateRequest != nil && obj.DowntimeScheduleRecurrencesCreateRequest.UnparsedObject == nil {
			jsonDowntimeScheduleRecurrencesCreateRequest, _ := json.Marshal(obj.DowntimeScheduleRecurrencesCreateRequest)
			if string(jsonDowntimeScheduleRecurrencesCreateRequest) == "{}" { // empty struct
				obj.DowntimeScheduleRecurrencesCreateRequest = nil
			} else {
				match++
			}
		} else {
			obj.DowntimeScheduleRecurrencesCreateRequest = nil
		}
	} else {
		obj.DowntimeScheduleRecurrencesCreateRequest = nil
	}

	// try to unmarshal data into DowntimeScheduleOneTimeCreateUpdateRequest
	err = json.Unmarshal(data, &obj.DowntimeScheduleOneTimeCreateUpdateRequest)
	if err == nil {
		if obj.DowntimeScheduleOneTimeCreateUpdateRequest != nil && obj.DowntimeScheduleOneTimeCreateUpdateRequest.UnparsedObject == nil {
			jsonDowntimeScheduleOneTimeCreateUpdateRequest, _ := json.Marshal(obj.DowntimeScheduleOneTimeCreateUpdateRequest)
			if string(jsonDowntimeScheduleOneTimeCreateUpdateRequest) == "{}" { // empty struct
				obj.DowntimeScheduleOneTimeCreateUpdateRequest = nil
			} else {
				match++
			}
		} else {
			obj.DowntimeScheduleOneTimeCreateUpdateRequest = nil
		}
	} else {
		obj.DowntimeScheduleOneTimeCreateUpdateRequest = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.DowntimeScheduleRecurrencesCreateRequest = nil
		obj.DowntimeScheduleOneTimeCreateUpdateRequest = nil
		return json.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj DowntimeScheduleCreateRequest) MarshalJSON() ([]byte, error) {
	if obj.DowntimeScheduleRecurrencesCreateRequest != nil {
		return json.Marshal(&obj.DowntimeScheduleRecurrencesCreateRequest)
	}

	if obj.DowntimeScheduleOneTimeCreateUpdateRequest != nil {
		return json.Marshal(&obj.DowntimeScheduleOneTimeCreateUpdateRequest)
	}

	if obj.UnparsedObject != nil {
		return json.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *DowntimeScheduleCreateRequest) GetActualInstance() interface{} {
	if obj.DowntimeScheduleRecurrencesCreateRequest != nil {
		return obj.DowntimeScheduleRecurrencesCreateRequest
	}

	if obj.DowntimeScheduleOneTimeCreateUpdateRequest != nil {
		return obj.DowntimeScheduleOneTimeCreateUpdateRequest
	}

	// all schemas are nil
	return nil
}
