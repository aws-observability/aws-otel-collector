// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMGroupByMissing - The value to use for logs that don't have the facet used to group by.
type RUMGroupByMissing struct {
	RUMGroupByMissingString *string
	RUMGroupByMissingNumber *float64

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// RUMGroupByMissingStringAsRUMGroupByMissing is a convenience function that returns string wrapped in RUMGroupByMissing.
func RUMGroupByMissingStringAsRUMGroupByMissing(v *string) RUMGroupByMissing {
	return RUMGroupByMissing{RUMGroupByMissingString: v}
}

// RUMGroupByMissingNumberAsRUMGroupByMissing is a convenience function that returns float64 wrapped in RUMGroupByMissing.
func RUMGroupByMissingNumberAsRUMGroupByMissing(v *float64) RUMGroupByMissing {
	return RUMGroupByMissing{RUMGroupByMissingNumber: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *RUMGroupByMissing) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into RUMGroupByMissingString
	err = datadog.Unmarshal(data, &obj.RUMGroupByMissingString)
	if err == nil {
		if obj.RUMGroupByMissingString != nil {
			jsonRUMGroupByMissingString, _ := datadog.Marshal(obj.RUMGroupByMissingString)
			if string(jsonRUMGroupByMissingString) == "{}" { // empty struct
				obj.RUMGroupByMissingString = nil
			} else {
				match++
			}
		} else {
			obj.RUMGroupByMissingString = nil
		}
	} else {
		obj.RUMGroupByMissingString = nil
	}

	// try to unmarshal data into RUMGroupByMissingNumber
	err = datadog.Unmarshal(data, &obj.RUMGroupByMissingNumber)
	if err == nil {
		if obj.RUMGroupByMissingNumber != nil {
			jsonRUMGroupByMissingNumber, _ := datadog.Marshal(obj.RUMGroupByMissingNumber)
			if string(jsonRUMGroupByMissingNumber) == "{}" { // empty struct
				obj.RUMGroupByMissingNumber = nil
			} else {
				match++
			}
		} else {
			obj.RUMGroupByMissingNumber = nil
		}
	} else {
		obj.RUMGroupByMissingNumber = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.RUMGroupByMissingString = nil
		obj.RUMGroupByMissingNumber = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj RUMGroupByMissing) MarshalJSON() ([]byte, error) {
	if obj.RUMGroupByMissingString != nil {
		return datadog.Marshal(&obj.RUMGroupByMissingString)
	}

	if obj.RUMGroupByMissingNumber != nil {
		return datadog.Marshal(&obj.RUMGroupByMissingNumber)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *RUMGroupByMissing) GetActualInstance() interface{} {
	if obj.RUMGroupByMissingString != nil {
		return obj.RUMGroupByMissingString
	}

	if obj.RUMGroupByMissingNumber != nil {
		return obj.RUMGroupByMissingNumber
	}

	// all schemas are nil
	return nil
}
