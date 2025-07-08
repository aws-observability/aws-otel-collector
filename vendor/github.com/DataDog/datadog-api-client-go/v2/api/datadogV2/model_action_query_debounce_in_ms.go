// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ActionQueryDebounceInMs - The minimum time in milliseconds that must pass before the query can be triggered again. This is useful for preventing accidental double-clicks from triggering the query multiple times.
type ActionQueryDebounceInMs struct {
	Float64 *float64
	String  *string

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// Float64AsActionQueryDebounceInMs is a convenience function that returns float64 wrapped in ActionQueryDebounceInMs.
func Float64AsActionQueryDebounceInMs(v *float64) ActionQueryDebounceInMs {
	return ActionQueryDebounceInMs{Float64: v}
}

// StringAsActionQueryDebounceInMs is a convenience function that returns string wrapped in ActionQueryDebounceInMs.
func StringAsActionQueryDebounceInMs(v *string) ActionQueryDebounceInMs {
	return ActionQueryDebounceInMs{String: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ActionQueryDebounceInMs) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Float64
	err = datadog.Unmarshal(data, &obj.Float64)
	if err == nil {
		if obj.Float64 != nil {
			jsonFloat64, _ := datadog.Marshal(obj.Float64)
			if string(jsonFloat64) == "{}" { // empty struct
				obj.Float64 = nil
			} else {
				match++
			}
		} else {
			obj.Float64 = nil
		}
	} else {
		obj.Float64 = nil
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
		obj.Float64 = nil
		obj.String = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ActionQueryDebounceInMs) MarshalJSON() ([]byte, error) {
	if obj.Float64 != nil {
		return datadog.Marshal(&obj.Float64)
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
func (obj *ActionQueryDebounceInMs) GetActualInstance() interface{} {
	if obj.Float64 != nil {
		return obj.Float64
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}
