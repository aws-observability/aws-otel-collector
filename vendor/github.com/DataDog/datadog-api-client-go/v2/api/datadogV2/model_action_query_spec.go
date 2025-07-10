// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ActionQuerySpec - The definition of the action query.
type ActionQuerySpec struct {
	String                *string
	ActionQuerySpecObject *ActionQuerySpecObject

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// StringAsActionQuerySpec is a convenience function that returns string wrapped in ActionQuerySpec.
func StringAsActionQuerySpec(v *string) ActionQuerySpec {
	return ActionQuerySpec{String: v}
}

// ActionQuerySpecObjectAsActionQuerySpec is a convenience function that returns ActionQuerySpecObject wrapped in ActionQuerySpec.
func ActionQuerySpecObjectAsActionQuerySpec(v *ActionQuerySpecObject) ActionQuerySpec {
	return ActionQuerySpec{ActionQuerySpecObject: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ActionQuerySpec) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into ActionQuerySpecObject
	err = datadog.Unmarshal(data, &obj.ActionQuerySpecObject)
	if err == nil {
		if obj.ActionQuerySpecObject != nil && obj.ActionQuerySpecObject.UnparsedObject == nil {
			jsonActionQuerySpecObject, _ := datadog.Marshal(obj.ActionQuerySpecObject)
			if string(jsonActionQuerySpecObject) == "{}" { // empty struct
				obj.ActionQuerySpecObject = nil
			} else {
				match++
			}
		} else {
			obj.ActionQuerySpecObject = nil
		}
	} else {
		obj.ActionQuerySpecObject = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.String = nil
		obj.ActionQuerySpecObject = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ActionQuerySpec) MarshalJSON() ([]byte, error) {
	if obj.String != nil {
		return datadog.Marshal(&obj.String)
	}

	if obj.ActionQuerySpecObject != nil {
		return datadog.Marshal(&obj.ActionQuerySpecObject)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ActionQuerySpec) GetActualInstance() interface{} {
	if obj.String != nil {
		return obj.String
	}

	if obj.ActionQuerySpecObject != nil {
		return obj.ActionQuerySpecObject
	}

	// all schemas are nil
	return nil
}
