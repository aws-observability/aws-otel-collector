// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DowntimeScheduleUpdateRequest - Schedule for the downtime.
type DowntimeScheduleUpdateRequest struct {
	DowntimeScheduleRecurrencesUpdateRequest   *DowntimeScheduleRecurrencesUpdateRequest
	DowntimeScheduleOneTimeCreateUpdateRequest *DowntimeScheduleOneTimeCreateUpdateRequest

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// DowntimeScheduleRecurrencesUpdateRequestAsDowntimeScheduleUpdateRequest is a convenience function that returns DowntimeScheduleRecurrencesUpdateRequest wrapped in DowntimeScheduleUpdateRequest.
func DowntimeScheduleRecurrencesUpdateRequestAsDowntimeScheduleUpdateRequest(v *DowntimeScheduleRecurrencesUpdateRequest) DowntimeScheduleUpdateRequest {
	return DowntimeScheduleUpdateRequest{DowntimeScheduleRecurrencesUpdateRequest: v}
}

// DowntimeScheduleOneTimeCreateUpdateRequestAsDowntimeScheduleUpdateRequest is a convenience function that returns DowntimeScheduleOneTimeCreateUpdateRequest wrapped in DowntimeScheduleUpdateRequest.
func DowntimeScheduleOneTimeCreateUpdateRequestAsDowntimeScheduleUpdateRequest(v *DowntimeScheduleOneTimeCreateUpdateRequest) DowntimeScheduleUpdateRequest {
	return DowntimeScheduleUpdateRequest{DowntimeScheduleOneTimeCreateUpdateRequest: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *DowntimeScheduleUpdateRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DowntimeScheduleRecurrencesUpdateRequest
	err = datadog.Unmarshal(data, &obj.DowntimeScheduleRecurrencesUpdateRequest)
	if err == nil {
		if obj.DowntimeScheduleRecurrencesUpdateRequest != nil && obj.DowntimeScheduleRecurrencesUpdateRequest.UnparsedObject == nil {
			jsonDowntimeScheduleRecurrencesUpdateRequest, _ := datadog.Marshal(obj.DowntimeScheduleRecurrencesUpdateRequest)
			if string(jsonDowntimeScheduleRecurrencesUpdateRequest) == "{}" && string(data) != "{}" { // empty struct
				obj.DowntimeScheduleRecurrencesUpdateRequest = nil
			} else {
				match++
			}
		} else {
			obj.DowntimeScheduleRecurrencesUpdateRequest = nil
		}
	} else {
		obj.DowntimeScheduleRecurrencesUpdateRequest = nil
	}

	// try to unmarshal data into DowntimeScheduleOneTimeCreateUpdateRequest
	err = datadog.Unmarshal(data, &obj.DowntimeScheduleOneTimeCreateUpdateRequest)
	if err == nil {
		if obj.DowntimeScheduleOneTimeCreateUpdateRequest != nil && obj.DowntimeScheduleOneTimeCreateUpdateRequest.UnparsedObject == nil {
			jsonDowntimeScheduleOneTimeCreateUpdateRequest, _ := datadog.Marshal(obj.DowntimeScheduleOneTimeCreateUpdateRequest)
			if string(jsonDowntimeScheduleOneTimeCreateUpdateRequest) == "{}" && string(data) != "{}" { // empty struct
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
		obj.DowntimeScheduleRecurrencesUpdateRequest = nil
		obj.DowntimeScheduleOneTimeCreateUpdateRequest = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj DowntimeScheduleUpdateRequest) MarshalJSON() ([]byte, error) {
	if obj.DowntimeScheduleRecurrencesUpdateRequest != nil {
		return datadog.Marshal(&obj.DowntimeScheduleRecurrencesUpdateRequest)
	}

	if obj.DowntimeScheduleOneTimeCreateUpdateRequest != nil {
		return datadog.Marshal(&obj.DowntimeScheduleOneTimeCreateUpdateRequest)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *DowntimeScheduleUpdateRequest) GetActualInstance() interface{} {
	if obj.DowntimeScheduleRecurrencesUpdateRequest != nil {
		return obj.DowntimeScheduleRecurrencesUpdateRequest
	}

	if obj.DowntimeScheduleOneTimeCreateUpdateRequest != nil {
		return obj.DowntimeScheduleOneTimeCreateUpdateRequest
	}

	// all schemas are nil
	return nil
}
