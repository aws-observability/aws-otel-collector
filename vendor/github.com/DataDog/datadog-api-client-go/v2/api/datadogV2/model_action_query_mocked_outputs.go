// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ActionQueryMockedOutputs - The mocked outputs of the action query. This is useful for testing the app without actually running the action.
type ActionQueryMockedOutputs struct {
	String                         *string
	ActionQueryMockedOutputsObject *ActionQueryMockedOutputsObject

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// StringAsActionQueryMockedOutputs is a convenience function that returns string wrapped in ActionQueryMockedOutputs.
func StringAsActionQueryMockedOutputs(v *string) ActionQueryMockedOutputs {
	return ActionQueryMockedOutputs{String: v}
}

// ActionQueryMockedOutputsObjectAsActionQueryMockedOutputs is a convenience function that returns ActionQueryMockedOutputsObject wrapped in ActionQueryMockedOutputs.
func ActionQueryMockedOutputsObjectAsActionQueryMockedOutputs(v *ActionQueryMockedOutputsObject) ActionQueryMockedOutputs {
	return ActionQueryMockedOutputs{ActionQueryMockedOutputsObject: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ActionQueryMockedOutputs) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into ActionQueryMockedOutputsObject
	err = datadog.Unmarshal(data, &obj.ActionQueryMockedOutputsObject)
	if err == nil {
		if obj.ActionQueryMockedOutputsObject != nil && obj.ActionQueryMockedOutputsObject.UnparsedObject == nil {
			jsonActionQueryMockedOutputsObject, _ := datadog.Marshal(obj.ActionQueryMockedOutputsObject)
			if string(jsonActionQueryMockedOutputsObject) == "{}" { // empty struct
				obj.ActionQueryMockedOutputsObject = nil
			} else {
				match++
			}
		} else {
			obj.ActionQueryMockedOutputsObject = nil
		}
	} else {
		obj.ActionQueryMockedOutputsObject = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.String = nil
		obj.ActionQueryMockedOutputsObject = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ActionQueryMockedOutputs) MarshalJSON() ([]byte, error) {
	if obj.String != nil {
		return datadog.Marshal(&obj.String)
	}

	if obj.ActionQueryMockedOutputsObject != nil {
		return datadog.Marshal(&obj.ActionQueryMockedOutputsObject)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ActionQueryMockedOutputs) GetActualInstance() interface{} {
	if obj.String != nil {
		return obj.String
	}

	if obj.ActionQueryMockedOutputsObject != nil {
		return obj.ActionQueryMockedOutputsObject
	}

	// all schemas are nil
	return nil
}
