// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ComponentGridPropertiesIsVisible - Whether the grid component and its children are visible. If a string, it must be a valid JavaScript expression that evaluates to a boolean.
type ComponentGridPropertiesIsVisible struct {
	String *string
	Bool   *bool

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// StringAsComponentGridPropertiesIsVisible is a convenience function that returns string wrapped in ComponentGridPropertiesIsVisible.
func StringAsComponentGridPropertiesIsVisible(v *string) ComponentGridPropertiesIsVisible {
	return ComponentGridPropertiesIsVisible{String: v}
}

// BoolAsComponentGridPropertiesIsVisible is a convenience function that returns bool wrapped in ComponentGridPropertiesIsVisible.
func BoolAsComponentGridPropertiesIsVisible(v *bool) ComponentGridPropertiesIsVisible {
	return ComponentGridPropertiesIsVisible{Bool: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ComponentGridPropertiesIsVisible) UnmarshalJSON(data []byte) error {
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

	if match != 1 { // more than 1 match
		// reset to nil
		obj.String = nil
		obj.Bool = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ComponentGridPropertiesIsVisible) MarshalJSON() ([]byte, error) {
	if obj.String != nil {
		return datadog.Marshal(&obj.String)
	}

	if obj.Bool != nil {
		return datadog.Marshal(&obj.Bool)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ComponentGridPropertiesIsVisible) GetActualInstance() interface{} {
	if obj.String != nil {
		return obj.String
	}

	if obj.Bool != nil {
		return obj.Bool
	}

	// all schemas are nil
	return nil
}
