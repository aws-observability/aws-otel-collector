// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotionCredentialsUpdate - The definition of the `NotionCredentialsUpdate` object.
type NotionCredentialsUpdate struct {
	NotionAPIKeyUpdate *NotionAPIKeyUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// NotionAPIKeyUpdateAsNotionCredentialsUpdate is a convenience function that returns NotionAPIKeyUpdate wrapped in NotionCredentialsUpdate.
func NotionAPIKeyUpdateAsNotionCredentialsUpdate(v *NotionAPIKeyUpdate) NotionCredentialsUpdate {
	return NotionCredentialsUpdate{NotionAPIKeyUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *NotionCredentialsUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into NotionAPIKeyUpdate
	err = datadog.Unmarshal(data, &obj.NotionAPIKeyUpdate)
	if err == nil {
		if obj.NotionAPIKeyUpdate != nil && obj.NotionAPIKeyUpdate.UnparsedObject == nil {
			jsonNotionAPIKeyUpdate, _ := datadog.Marshal(obj.NotionAPIKeyUpdate)
			if string(jsonNotionAPIKeyUpdate) == "{}" { // empty struct
				obj.NotionAPIKeyUpdate = nil
			} else {
				match++
			}
		} else {
			obj.NotionAPIKeyUpdate = nil
		}
	} else {
		obj.NotionAPIKeyUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.NotionAPIKeyUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj NotionCredentialsUpdate) MarshalJSON() ([]byte, error) {
	if obj.NotionAPIKeyUpdate != nil {
		return datadog.Marshal(&obj.NotionAPIKeyUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *NotionCredentialsUpdate) GetActualInstance() interface{} {
	if obj.NotionAPIKeyUpdate != nil {
		return obj.NotionAPIKeyUpdate
	}

	// all schemas are nil
	return nil
}
