// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ActionQuerySpecInputs - The inputs to the action query. These are the values that are passed to the action when it is triggered.
type ActionQuerySpecInputs struct {
	String               *string
	ActionQuerySpecInput map[string]interface{}

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// StringAsActionQuerySpecInputs is a convenience function that returns string wrapped in ActionQuerySpecInputs.
func StringAsActionQuerySpecInputs(v *string) ActionQuerySpecInputs {
	return ActionQuerySpecInputs{String: v}
}

// ActionQuerySpecInputAsActionQuerySpecInputs is a convenience function that returns map[string]interface{} wrapped in ActionQuerySpecInputs.
func ActionQuerySpecInputAsActionQuerySpecInputs(v map[string]interface{}) ActionQuerySpecInputs {
	return ActionQuerySpecInputs{ActionQuerySpecInput: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ActionQuerySpecInputs) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into String
	err = datadog.Unmarshal(data, &obj.String)
	if err == nil {
		if obj.String != nil {
			jsonString, _ := datadog.Marshal(obj.String)
			if string(jsonString) == "{}" { // empty struct
				obj.String = nil
			} else {
				match++
			}
		} else {
			obj.String = nil
		}
	} else {
		obj.String = nil
	}

	// try to unmarshal data into ActionQuerySpecInput
	err = datadog.Unmarshal(data, &obj.ActionQuerySpecInput)
	if err == nil {
		if obj.ActionQuerySpecInput != nil {
			jsonActionQuerySpecInput, _ := datadog.Marshal(obj.ActionQuerySpecInput)
			if string(jsonActionQuerySpecInput) == "{}" && string(data) != "{}" { // empty struct
				obj.ActionQuerySpecInput = nil
			} else {
				match++
			}
		} else {
			obj.ActionQuerySpecInput = nil
		}
	} else {
		obj.ActionQuerySpecInput = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.String = nil
		obj.ActionQuerySpecInput = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ActionQuerySpecInputs) MarshalJSON() ([]byte, error) {
	if obj.String != nil {
		return datadog.Marshal(&obj.String)
	}

	if obj.ActionQuerySpecInput != nil {
		return datadog.Marshal(&obj.ActionQuerySpecInput)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ActionQuerySpecInputs) GetActualInstance() interface{} {
	if obj.String != nil {
		return obj.String
	}

	if obj.ActionQuerySpecInput != nil {
		return obj.ActionQuerySpecInput
	}

	// all schemas are nil
	return nil
}
