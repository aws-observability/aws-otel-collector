// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAICredentialsUpdate - The definition of the `OpenAICredentialsUpdate` object.
type OpenAICredentialsUpdate struct {
	OpenAIAPIKeyUpdate *OpenAIAPIKeyUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// OpenAIAPIKeyUpdateAsOpenAICredentialsUpdate is a convenience function that returns OpenAIAPIKeyUpdate wrapped in OpenAICredentialsUpdate.
func OpenAIAPIKeyUpdateAsOpenAICredentialsUpdate(v *OpenAIAPIKeyUpdate) OpenAICredentialsUpdate {
	return OpenAICredentialsUpdate{OpenAIAPIKeyUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *OpenAICredentialsUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into OpenAIAPIKeyUpdate
	err = datadog.Unmarshal(data, &obj.OpenAIAPIKeyUpdate)
	if err == nil {
		if obj.OpenAIAPIKeyUpdate != nil && obj.OpenAIAPIKeyUpdate.UnparsedObject == nil {
			jsonOpenAIAPIKeyUpdate, _ := datadog.Marshal(obj.OpenAIAPIKeyUpdate)
			if string(jsonOpenAIAPIKeyUpdate) == "{}" { // empty struct
				obj.OpenAIAPIKeyUpdate = nil
			} else {
				match++
			}
		} else {
			obj.OpenAIAPIKeyUpdate = nil
		}
	} else {
		obj.OpenAIAPIKeyUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.OpenAIAPIKeyUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj OpenAICredentialsUpdate) MarshalJSON() ([]byte, error) {
	if obj.OpenAIAPIKeyUpdate != nil {
		return datadog.Marshal(&obj.OpenAIAPIKeyUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *OpenAICredentialsUpdate) GetActualInstance() interface{} {
	if obj.OpenAIAPIKeyUpdate != nil {
		return obj.OpenAIAPIKeyUpdate
	}

	// all schemas are nil
	return nil
}
