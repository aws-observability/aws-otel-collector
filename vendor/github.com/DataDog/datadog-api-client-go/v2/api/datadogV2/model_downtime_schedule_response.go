// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DowntimeScheduleResponse - The schedule that defines when the monitor starts, stops, and recurs. There are two types of schedules:
// one-time and recurring. Recurring schedules may have up to five RRULE-based recurrences. If no schedules are
// provided, the downtime will begin immediately and never end.
type DowntimeScheduleResponse struct {
	DowntimeScheduleRecurrencesResponse *DowntimeScheduleRecurrencesResponse
	DowntimeScheduleOneTimeResponse     *DowntimeScheduleOneTimeResponse

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// DowntimeScheduleRecurrencesResponseAsDowntimeScheduleResponse is a convenience function that returns DowntimeScheduleRecurrencesResponse wrapped in DowntimeScheduleResponse.
func DowntimeScheduleRecurrencesResponseAsDowntimeScheduleResponse(v *DowntimeScheduleRecurrencesResponse) DowntimeScheduleResponse {
	return DowntimeScheduleResponse{DowntimeScheduleRecurrencesResponse: v}
}

// DowntimeScheduleOneTimeResponseAsDowntimeScheduleResponse is a convenience function that returns DowntimeScheduleOneTimeResponse wrapped in DowntimeScheduleResponse.
func DowntimeScheduleOneTimeResponseAsDowntimeScheduleResponse(v *DowntimeScheduleOneTimeResponse) DowntimeScheduleResponse {
	return DowntimeScheduleResponse{DowntimeScheduleOneTimeResponse: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *DowntimeScheduleResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DowntimeScheduleRecurrencesResponse
	err = datadog.Unmarshal(data, &obj.DowntimeScheduleRecurrencesResponse)
	if err == nil {
		if obj.DowntimeScheduleRecurrencesResponse != nil && obj.DowntimeScheduleRecurrencesResponse.UnparsedObject == nil {
			jsonDowntimeScheduleRecurrencesResponse, _ := datadog.Marshal(obj.DowntimeScheduleRecurrencesResponse)
			if string(jsonDowntimeScheduleRecurrencesResponse) == "{}" { // empty struct
				obj.DowntimeScheduleRecurrencesResponse = nil
			} else {
				match++
			}
		} else {
			obj.DowntimeScheduleRecurrencesResponse = nil
		}
	} else {
		obj.DowntimeScheduleRecurrencesResponse = nil
	}

	// try to unmarshal data into DowntimeScheduleOneTimeResponse
	err = datadog.Unmarshal(data, &obj.DowntimeScheduleOneTimeResponse)
	if err == nil {
		if obj.DowntimeScheduleOneTimeResponse != nil && obj.DowntimeScheduleOneTimeResponse.UnparsedObject == nil {
			jsonDowntimeScheduleOneTimeResponse, _ := datadog.Marshal(obj.DowntimeScheduleOneTimeResponse)
			if string(jsonDowntimeScheduleOneTimeResponse) == "{}" { // empty struct
				obj.DowntimeScheduleOneTimeResponse = nil
			} else {
				match++
			}
		} else {
			obj.DowntimeScheduleOneTimeResponse = nil
		}
	} else {
		obj.DowntimeScheduleOneTimeResponse = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.DowntimeScheduleRecurrencesResponse = nil
		obj.DowntimeScheduleOneTimeResponse = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj DowntimeScheduleResponse) MarshalJSON() ([]byte, error) {
	if obj.DowntimeScheduleRecurrencesResponse != nil {
		return datadog.Marshal(&obj.DowntimeScheduleRecurrencesResponse)
	}

	if obj.DowntimeScheduleOneTimeResponse != nil {
		return datadog.Marshal(&obj.DowntimeScheduleOneTimeResponse)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *DowntimeScheduleResponse) GetActualInstance() interface{} {
	if obj.DowntimeScheduleRecurrencesResponse != nil {
		return obj.DowntimeScheduleRecurrencesResponse
	}

	if obj.DowntimeScheduleOneTimeResponse != nil {
		return obj.DowntimeScheduleOneTimeResponse
	}

	// all schemas are nil
	return nil
}
