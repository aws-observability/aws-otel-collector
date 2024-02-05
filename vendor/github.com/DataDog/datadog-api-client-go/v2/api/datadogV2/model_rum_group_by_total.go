// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMGroupByTotal - A resulting object to put the given computes in over all the matching records.
type RUMGroupByTotal struct {
	RUMGroupByTotalBoolean *bool
	RUMGroupByTotalString  *string
	RUMGroupByTotalNumber  *float64

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// RUMGroupByTotalBooleanAsRUMGroupByTotal is a convenience function that returns bool wrapped in RUMGroupByTotal.
func RUMGroupByTotalBooleanAsRUMGroupByTotal(v *bool) RUMGroupByTotal {
	return RUMGroupByTotal{RUMGroupByTotalBoolean: v}
}

// RUMGroupByTotalStringAsRUMGroupByTotal is a convenience function that returns string wrapped in RUMGroupByTotal.
func RUMGroupByTotalStringAsRUMGroupByTotal(v *string) RUMGroupByTotal {
	return RUMGroupByTotal{RUMGroupByTotalString: v}
}

// RUMGroupByTotalNumberAsRUMGroupByTotal is a convenience function that returns float64 wrapped in RUMGroupByTotal.
func RUMGroupByTotalNumberAsRUMGroupByTotal(v *float64) RUMGroupByTotal {
	return RUMGroupByTotal{RUMGroupByTotalNumber: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *RUMGroupByTotal) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into RUMGroupByTotalBoolean
	err = datadog.Unmarshal(data, &obj.RUMGroupByTotalBoolean)
	if err == nil {
		if obj.RUMGroupByTotalBoolean != nil {
			jsonRUMGroupByTotalBoolean, _ := datadog.Marshal(obj.RUMGroupByTotalBoolean)
			if string(jsonRUMGroupByTotalBoolean) == "{}" { // empty struct
				obj.RUMGroupByTotalBoolean = nil
			} else {
				match++
			}
		} else {
			obj.RUMGroupByTotalBoolean = nil
		}
	} else {
		obj.RUMGroupByTotalBoolean = nil
	}

	// try to unmarshal data into RUMGroupByTotalString
	err = datadog.Unmarshal(data, &obj.RUMGroupByTotalString)
	if err == nil {
		if obj.RUMGroupByTotalString != nil {
			jsonRUMGroupByTotalString, _ := datadog.Marshal(obj.RUMGroupByTotalString)
			if string(jsonRUMGroupByTotalString) == "{}" { // empty struct
				obj.RUMGroupByTotalString = nil
			} else {
				match++
			}
		} else {
			obj.RUMGroupByTotalString = nil
		}
	} else {
		obj.RUMGroupByTotalString = nil
	}

	// try to unmarshal data into RUMGroupByTotalNumber
	err = datadog.Unmarshal(data, &obj.RUMGroupByTotalNumber)
	if err == nil {
		if obj.RUMGroupByTotalNumber != nil {
			jsonRUMGroupByTotalNumber, _ := datadog.Marshal(obj.RUMGroupByTotalNumber)
			if string(jsonRUMGroupByTotalNumber) == "{}" { // empty struct
				obj.RUMGroupByTotalNumber = nil
			} else {
				match++
			}
		} else {
			obj.RUMGroupByTotalNumber = nil
		}
	} else {
		obj.RUMGroupByTotalNumber = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.RUMGroupByTotalBoolean = nil
		obj.RUMGroupByTotalString = nil
		obj.RUMGroupByTotalNumber = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj RUMGroupByTotal) MarshalJSON() ([]byte, error) {
	if obj.RUMGroupByTotalBoolean != nil {
		return datadog.Marshal(&obj.RUMGroupByTotalBoolean)
	}

	if obj.RUMGroupByTotalString != nil {
		return datadog.Marshal(&obj.RUMGroupByTotalString)
	}

	if obj.RUMGroupByTotalNumber != nil {
		return datadog.Marshal(&obj.RUMGroupByTotalNumber)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *RUMGroupByTotal) GetActualInstance() interface{} {
	if obj.RUMGroupByTotalBoolean != nil {
		return obj.RUMGroupByTotalBoolean
	}

	if obj.RUMGroupByTotalString != nil {
		return obj.RUMGroupByTotalString
	}

	if obj.RUMGroupByTotalNumber != nil {
		return obj.RUMGroupByTotalNumber
	}

	// all schemas are nil
	return nil
}
