// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnthropicCredentialsUpdate - The definition of the `AnthropicCredentialsUpdate` object.
type AnthropicCredentialsUpdate struct {
	AnthropicAPIKeyUpdate *AnthropicAPIKeyUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// AnthropicAPIKeyUpdateAsAnthropicCredentialsUpdate is a convenience function that returns AnthropicAPIKeyUpdate wrapped in AnthropicCredentialsUpdate.
func AnthropicAPIKeyUpdateAsAnthropicCredentialsUpdate(v *AnthropicAPIKeyUpdate) AnthropicCredentialsUpdate {
	return AnthropicCredentialsUpdate{AnthropicAPIKeyUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *AnthropicCredentialsUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AnthropicAPIKeyUpdate
	err = datadog.Unmarshal(data, &obj.AnthropicAPIKeyUpdate)
	if err == nil {
		if obj.AnthropicAPIKeyUpdate != nil && obj.AnthropicAPIKeyUpdate.UnparsedObject == nil {
			jsonAnthropicAPIKeyUpdate, _ := datadog.Marshal(obj.AnthropicAPIKeyUpdate)
			if string(jsonAnthropicAPIKeyUpdate) == "{}" { // empty struct
				obj.AnthropicAPIKeyUpdate = nil
			} else {
				match++
			}
		} else {
			obj.AnthropicAPIKeyUpdate = nil
		}
	} else {
		obj.AnthropicAPIKeyUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.AnthropicAPIKeyUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj AnthropicCredentialsUpdate) MarshalJSON() ([]byte, error) {
	if obj.AnthropicAPIKeyUpdate != nil {
		return datadog.Marshal(&obj.AnthropicAPIKeyUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *AnthropicCredentialsUpdate) GetActualInstance() interface{} {
	if obj.AnthropicAPIKeyUpdate != nil {
		return obj.AnthropicAPIKeyUpdate
	}

	// all schemas are nil
	return nil
}
