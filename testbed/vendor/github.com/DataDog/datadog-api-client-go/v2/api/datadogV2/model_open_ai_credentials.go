// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAICredentials - The definition of the `OpenAICredentials` object.
type OpenAICredentials struct {
	OpenAIAPIKey *OpenAIAPIKey

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// OpenAIAPIKeyAsOpenAICredentials is a convenience function that returns OpenAIAPIKey wrapped in OpenAICredentials.
func OpenAIAPIKeyAsOpenAICredentials(v *OpenAIAPIKey) OpenAICredentials {
	return OpenAICredentials{OpenAIAPIKey: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *OpenAICredentials) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into OpenAIAPIKey
	err = datadog.Unmarshal(data, &obj.OpenAIAPIKey)
	if err == nil {
		if obj.OpenAIAPIKey != nil && obj.OpenAIAPIKey.UnparsedObject == nil {
			jsonOpenAIAPIKey, _ := datadog.Marshal(obj.OpenAIAPIKey)
			if string(jsonOpenAIAPIKey) == "{}" { // empty struct
				obj.OpenAIAPIKey = nil
			} else {
				match++
			}
		} else {
			obj.OpenAIAPIKey = nil
		}
	} else {
		obj.OpenAIAPIKey = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.OpenAIAPIKey = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj OpenAICredentials) MarshalJSON() ([]byte, error) {
	if obj.OpenAIAPIKey != nil {
		return datadog.Marshal(&obj.OpenAIAPIKey)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *OpenAICredentials) GetActualInstance() interface{} {
	if obj.OpenAIAPIKey != nil {
		return obj.OpenAIAPIKey
	}

	// all schemas are nil
	return nil
}
