// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ActionQueryShowToastOnError - Whether to display a toast to the user when the query returns an error.
type ActionQueryShowToastOnError struct {
	Bool   *bool
	String *string

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// BoolAsActionQueryShowToastOnError is a convenience function that returns bool wrapped in ActionQueryShowToastOnError.
func BoolAsActionQueryShowToastOnError(v *bool) ActionQueryShowToastOnError {
	return ActionQueryShowToastOnError{Bool: v}
}

// StringAsActionQueryShowToastOnError is a convenience function that returns string wrapped in ActionQueryShowToastOnError.
func StringAsActionQueryShowToastOnError(v *string) ActionQueryShowToastOnError {
	return ActionQueryShowToastOnError{String: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ActionQueryShowToastOnError) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Bool
	err = datadog.Unmarshal(data, &obj.Bool)
	if err == nil {
		if obj.Bool != nil {
			jsonBool, _ := datadog.Marshal(obj.Bool)
			if string(jsonBool) == "{}" { // empty struct
				obj.Bool = nil
			} else {
				match++
			}
		} else {
			obj.Bool = nil
		}
	} else {
		obj.Bool = nil
	}

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

	if match != 1 { // more than 1 match
		// reset to nil
		obj.Bool = nil
		obj.String = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ActionQueryShowToastOnError) MarshalJSON() ([]byte, error) {
	if obj.Bool != nil {
		return datadog.Marshal(&obj.Bool)
	}

	if obj.String != nil {
		return datadog.Marshal(&obj.String)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ActionQueryShowToastOnError) GetActualInstance() interface{} {
	if obj.Bool != nil {
		return obj.Bool
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}
