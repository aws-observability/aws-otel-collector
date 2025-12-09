// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AsanaCredentialsUpdate - The definition of the `AsanaCredentialsUpdate` object.
type AsanaCredentialsUpdate struct {
	AsanaAccessTokenUpdate *AsanaAccessTokenUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// AsanaAccessTokenUpdateAsAsanaCredentialsUpdate is a convenience function that returns AsanaAccessTokenUpdate wrapped in AsanaCredentialsUpdate.
func AsanaAccessTokenUpdateAsAsanaCredentialsUpdate(v *AsanaAccessTokenUpdate) AsanaCredentialsUpdate {
	return AsanaCredentialsUpdate{AsanaAccessTokenUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *AsanaCredentialsUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AsanaAccessTokenUpdate
	err = datadog.Unmarshal(data, &obj.AsanaAccessTokenUpdate)
	if err == nil {
		if obj.AsanaAccessTokenUpdate != nil && obj.AsanaAccessTokenUpdate.UnparsedObject == nil {
			jsonAsanaAccessTokenUpdate, _ := datadog.Marshal(obj.AsanaAccessTokenUpdate)
			if string(jsonAsanaAccessTokenUpdate) == "{}" { // empty struct
				obj.AsanaAccessTokenUpdate = nil
			} else {
				match++
			}
		} else {
			obj.AsanaAccessTokenUpdate = nil
		}
	} else {
		obj.AsanaAccessTokenUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.AsanaAccessTokenUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj AsanaCredentialsUpdate) MarshalJSON() ([]byte, error) {
	if obj.AsanaAccessTokenUpdate != nil {
		return datadog.Marshal(&obj.AsanaAccessTokenUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *AsanaCredentialsUpdate) GetActualInstance() interface{} {
	if obj.AsanaAccessTokenUpdate != nil {
		return obj.AsanaAccessTokenUpdate
	}

	// all schemas are nil
	return nil
}
