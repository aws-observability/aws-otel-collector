// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeminiCredentials - The definition of the `GeminiCredentials` object.
type GeminiCredentials struct {
	GeminiAPIKey *GeminiAPIKey

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// GeminiAPIKeyAsGeminiCredentials is a convenience function that returns GeminiAPIKey wrapped in GeminiCredentials.
func GeminiAPIKeyAsGeminiCredentials(v *GeminiAPIKey) GeminiCredentials {
	return GeminiCredentials{GeminiAPIKey: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *GeminiCredentials) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GeminiAPIKey
	err = datadog.Unmarshal(data, &obj.GeminiAPIKey)
	if err == nil {
		if obj.GeminiAPIKey != nil && obj.GeminiAPIKey.UnparsedObject == nil {
			jsonGeminiAPIKey, _ := datadog.Marshal(obj.GeminiAPIKey)
			if string(jsonGeminiAPIKey) == "{}" { // empty struct
				obj.GeminiAPIKey = nil
			} else {
				match++
			}
		} else {
			obj.GeminiAPIKey = nil
		}
	} else {
		obj.GeminiAPIKey = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.GeminiAPIKey = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj GeminiCredentials) MarshalJSON() ([]byte, error) {
	if obj.GeminiAPIKey != nil {
		return datadog.Marshal(&obj.GeminiAPIKey)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *GeminiCredentials) GetActualInstance() interface{} {
	if obj.GeminiAPIKey != nil {
		return obj.GeminiAPIKey
	}

	// all schemas are nil
	return nil
}
