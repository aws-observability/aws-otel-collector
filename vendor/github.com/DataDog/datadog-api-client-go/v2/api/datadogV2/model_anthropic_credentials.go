// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnthropicCredentials - The definition of the `AnthropicCredentials` object.
type AnthropicCredentials struct {
	AnthropicAPIKey *AnthropicAPIKey

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// AnthropicAPIKeyAsAnthropicCredentials is a convenience function that returns AnthropicAPIKey wrapped in AnthropicCredentials.
func AnthropicAPIKeyAsAnthropicCredentials(v *AnthropicAPIKey) AnthropicCredentials {
	return AnthropicCredentials{AnthropicAPIKey: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *AnthropicCredentials) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AnthropicAPIKey
	err = datadog.Unmarshal(data, &obj.AnthropicAPIKey)
	if err == nil {
		if obj.AnthropicAPIKey != nil && obj.AnthropicAPIKey.UnparsedObject == nil {
			jsonAnthropicAPIKey, _ := datadog.Marshal(obj.AnthropicAPIKey)
			if string(jsonAnthropicAPIKey) == "{}" { // empty struct
				obj.AnthropicAPIKey = nil
			} else {
				match++
			}
		} else {
			obj.AnthropicAPIKey = nil
		}
	} else {
		obj.AnthropicAPIKey = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.AnthropicAPIKey = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj AnthropicCredentials) MarshalJSON() ([]byte, error) {
	if obj.AnthropicAPIKey != nil {
		return datadog.Marshal(&obj.AnthropicAPIKey)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *AnthropicCredentials) GetActualInstance() interface{} {
	if obj.AnthropicAPIKey != nil {
		return obj.AnthropicAPIKey
	}

	// all schemas are nil
	return nil
}
