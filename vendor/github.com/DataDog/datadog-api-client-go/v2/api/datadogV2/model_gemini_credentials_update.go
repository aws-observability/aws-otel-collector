// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeminiCredentialsUpdate - The definition of the `GeminiCredentialsUpdate` object.
type GeminiCredentialsUpdate struct {
	GeminiAPIKeyUpdate *GeminiAPIKeyUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// GeminiAPIKeyUpdateAsGeminiCredentialsUpdate is a convenience function that returns GeminiAPIKeyUpdate wrapped in GeminiCredentialsUpdate.
func GeminiAPIKeyUpdateAsGeminiCredentialsUpdate(v *GeminiAPIKeyUpdate) GeminiCredentialsUpdate {
	return GeminiCredentialsUpdate{GeminiAPIKeyUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *GeminiCredentialsUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GeminiAPIKeyUpdate
	err = datadog.Unmarshal(data, &obj.GeminiAPIKeyUpdate)
	if err == nil {
		if obj.GeminiAPIKeyUpdate != nil && obj.GeminiAPIKeyUpdate.UnparsedObject == nil {
			jsonGeminiAPIKeyUpdate, _ := datadog.Marshal(obj.GeminiAPIKeyUpdate)
			if string(jsonGeminiAPIKeyUpdate) == "{}" { // empty struct
				obj.GeminiAPIKeyUpdate = nil
			} else {
				match++
			}
		} else {
			obj.GeminiAPIKeyUpdate = nil
		}
	} else {
		obj.GeminiAPIKeyUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.GeminiAPIKeyUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj GeminiCredentialsUpdate) MarshalJSON() ([]byte, error) {
	if obj.GeminiAPIKeyUpdate != nil {
		return datadog.Marshal(&obj.GeminiAPIKeyUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *GeminiCredentialsUpdate) GetActualInstance() interface{} {
	if obj.GeminiAPIKeyUpdate != nil {
		return obj.GeminiAPIKeyUpdate
	}

	// all schemas are nil
	return nil
}
