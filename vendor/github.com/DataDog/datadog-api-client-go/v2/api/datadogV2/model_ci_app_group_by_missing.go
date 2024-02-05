// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CIAppGroupByMissing - The value to use for logs that don't have the facet used to group-by.
type CIAppGroupByMissing struct {
	CIAppGroupByMissingString *string
	CIAppGroupByMissingNumber *float64

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// CIAppGroupByMissingStringAsCIAppGroupByMissing is a convenience function that returns string wrapped in CIAppGroupByMissing.
func CIAppGroupByMissingStringAsCIAppGroupByMissing(v *string) CIAppGroupByMissing {
	return CIAppGroupByMissing{CIAppGroupByMissingString: v}
}

// CIAppGroupByMissingNumberAsCIAppGroupByMissing is a convenience function that returns float64 wrapped in CIAppGroupByMissing.
func CIAppGroupByMissingNumberAsCIAppGroupByMissing(v *float64) CIAppGroupByMissing {
	return CIAppGroupByMissing{CIAppGroupByMissingNumber: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *CIAppGroupByMissing) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CIAppGroupByMissingString
	err = datadog.Unmarshal(data, &obj.CIAppGroupByMissingString)
	if err == nil {
		if obj.CIAppGroupByMissingString != nil {
			jsonCIAppGroupByMissingString, _ := datadog.Marshal(obj.CIAppGroupByMissingString)
			if string(jsonCIAppGroupByMissingString) == "{}" { // empty struct
				obj.CIAppGroupByMissingString = nil
			} else {
				match++
			}
		} else {
			obj.CIAppGroupByMissingString = nil
		}
	} else {
		obj.CIAppGroupByMissingString = nil
	}

	// try to unmarshal data into CIAppGroupByMissingNumber
	err = datadog.Unmarshal(data, &obj.CIAppGroupByMissingNumber)
	if err == nil {
		if obj.CIAppGroupByMissingNumber != nil {
			jsonCIAppGroupByMissingNumber, _ := datadog.Marshal(obj.CIAppGroupByMissingNumber)
			if string(jsonCIAppGroupByMissingNumber) == "{}" { // empty struct
				obj.CIAppGroupByMissingNumber = nil
			} else {
				match++
			}
		} else {
			obj.CIAppGroupByMissingNumber = nil
		}
	} else {
		obj.CIAppGroupByMissingNumber = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.CIAppGroupByMissingString = nil
		obj.CIAppGroupByMissingNumber = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj CIAppGroupByMissing) MarshalJSON() ([]byte, error) {
	if obj.CIAppGroupByMissingString != nil {
		return datadog.Marshal(&obj.CIAppGroupByMissingString)
	}

	if obj.CIAppGroupByMissingNumber != nil {
		return datadog.Marshal(&obj.CIAppGroupByMissingNumber)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *CIAppGroupByMissing) GetActualInstance() interface{} {
	if obj.CIAppGroupByMissingString != nil {
		return obj.CIAppGroupByMissingString
	}

	if obj.CIAppGroupByMissingNumber != nil {
		return obj.CIAppGroupByMissingNumber
	}

	// all schemas are nil
	return nil
}
