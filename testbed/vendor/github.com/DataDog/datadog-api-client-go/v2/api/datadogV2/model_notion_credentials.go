// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotionCredentials - The definition of the `NotionCredentials` object.
type NotionCredentials struct {
	NotionAPIKey *NotionAPIKey

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// NotionAPIKeyAsNotionCredentials is a convenience function that returns NotionAPIKey wrapped in NotionCredentials.
func NotionAPIKeyAsNotionCredentials(v *NotionAPIKey) NotionCredentials {
	return NotionCredentials{NotionAPIKey: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *NotionCredentials) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into NotionAPIKey
	err = datadog.Unmarshal(data, &obj.NotionAPIKey)
	if err == nil {
		if obj.NotionAPIKey != nil && obj.NotionAPIKey.UnparsedObject == nil {
			jsonNotionAPIKey, _ := datadog.Marshal(obj.NotionAPIKey)
			if string(jsonNotionAPIKey) == "{}" { // empty struct
				obj.NotionAPIKey = nil
			} else {
				match++
			}
		} else {
			obj.NotionAPIKey = nil
		}
	} else {
		obj.NotionAPIKey = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.NotionAPIKey = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj NotionCredentials) MarshalJSON() ([]byte, error) {
	if obj.NotionAPIKey != nil {
		return datadog.Marshal(&obj.NotionAPIKey)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *NotionCredentials) GetActualInstance() interface{} {
	if obj.NotionAPIKey != nil {
		return obj.NotionAPIKey
	}

	// all schemas are nil
	return nil
}
