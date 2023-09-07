// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/goccy/go-json"
)

// SpansGroupByMissing - The value to use for spans that don't have the facet used to group by.
type SpansGroupByMissing struct {
	SpansGroupByMissingString *string
	SpansGroupByMissingNumber *float64

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SpansGroupByMissingStringAsSpansGroupByMissing is a convenience function that returns string wrapped in SpansGroupByMissing.
func SpansGroupByMissingStringAsSpansGroupByMissing(v *string) SpansGroupByMissing {
	return SpansGroupByMissing{SpansGroupByMissingString: v}
}

// SpansGroupByMissingNumberAsSpansGroupByMissing is a convenience function that returns float64 wrapped in SpansGroupByMissing.
func SpansGroupByMissingNumberAsSpansGroupByMissing(v *float64) SpansGroupByMissing {
	return SpansGroupByMissing{SpansGroupByMissingNumber: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *SpansGroupByMissing) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SpansGroupByMissingString
	err = json.Unmarshal(data, &obj.SpansGroupByMissingString)
	if err == nil {
		if obj.SpansGroupByMissingString != nil {
			jsonSpansGroupByMissingString, _ := json.Marshal(obj.SpansGroupByMissingString)
			if string(jsonSpansGroupByMissingString) == "{}" { // empty struct
				obj.SpansGroupByMissingString = nil
			} else {
				match++
			}
		} else {
			obj.SpansGroupByMissingString = nil
		}
	} else {
		obj.SpansGroupByMissingString = nil
	}

	// try to unmarshal data into SpansGroupByMissingNumber
	err = json.Unmarshal(data, &obj.SpansGroupByMissingNumber)
	if err == nil {
		if obj.SpansGroupByMissingNumber != nil {
			jsonSpansGroupByMissingNumber, _ := json.Marshal(obj.SpansGroupByMissingNumber)
			if string(jsonSpansGroupByMissingNumber) == "{}" { // empty struct
				obj.SpansGroupByMissingNumber = nil
			} else {
				match++
			}
		} else {
			obj.SpansGroupByMissingNumber = nil
		}
	} else {
		obj.SpansGroupByMissingNumber = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.SpansGroupByMissingString = nil
		obj.SpansGroupByMissingNumber = nil
		return json.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj SpansGroupByMissing) MarshalJSON() ([]byte, error) {
	if obj.SpansGroupByMissingString != nil {
		return json.Marshal(&obj.SpansGroupByMissingString)
	}

	if obj.SpansGroupByMissingNumber != nil {
		return json.Marshal(&obj.SpansGroupByMissingNumber)
	}

	if obj.UnparsedObject != nil {
		return json.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *SpansGroupByMissing) GetActualInstance() interface{} {
	if obj.SpansGroupByMissingString != nil {
		return obj.SpansGroupByMissingString
	}

	if obj.SpansGroupByMissingNumber != nil {
		return obj.SpansGroupByMissingNumber
	}

	// all schemas are nil
	return nil
}
