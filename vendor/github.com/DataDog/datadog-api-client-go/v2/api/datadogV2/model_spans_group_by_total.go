// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SpansGroupByTotal - A resulting object to put the given computes in over all the matching records.
type SpansGroupByTotal struct {
	SpansGroupByTotalBoolean *bool
	SpansGroupByTotalString  *string
	SpansGroupByTotalNumber  *float64

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SpansGroupByTotalBooleanAsSpansGroupByTotal is a convenience function that returns bool wrapped in SpansGroupByTotal.
func SpansGroupByTotalBooleanAsSpansGroupByTotal(v *bool) SpansGroupByTotal {
	return SpansGroupByTotal{SpansGroupByTotalBoolean: v}
}

// SpansGroupByTotalStringAsSpansGroupByTotal is a convenience function that returns string wrapped in SpansGroupByTotal.
func SpansGroupByTotalStringAsSpansGroupByTotal(v *string) SpansGroupByTotal {
	return SpansGroupByTotal{SpansGroupByTotalString: v}
}

// SpansGroupByTotalNumberAsSpansGroupByTotal is a convenience function that returns float64 wrapped in SpansGroupByTotal.
func SpansGroupByTotalNumberAsSpansGroupByTotal(v *float64) SpansGroupByTotal {
	return SpansGroupByTotal{SpansGroupByTotalNumber: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *SpansGroupByTotal) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SpansGroupByTotalBoolean
	err = datadog.Unmarshal(data, &obj.SpansGroupByTotalBoolean)
	if err == nil {
		if obj.SpansGroupByTotalBoolean != nil {
			jsonSpansGroupByTotalBoolean, _ := datadog.Marshal(obj.SpansGroupByTotalBoolean)
			if string(jsonSpansGroupByTotalBoolean) == "{}" { // empty struct
				obj.SpansGroupByTotalBoolean = nil
			} else {
				match++
			}
		} else {
			obj.SpansGroupByTotalBoolean = nil
		}
	} else {
		obj.SpansGroupByTotalBoolean = nil
	}

	// try to unmarshal data into SpansGroupByTotalString
	err = datadog.Unmarshal(data, &obj.SpansGroupByTotalString)
	if err == nil {
		if obj.SpansGroupByTotalString != nil {
			jsonSpansGroupByTotalString, _ := datadog.Marshal(obj.SpansGroupByTotalString)
			if string(jsonSpansGroupByTotalString) == "{}" { // empty struct
				obj.SpansGroupByTotalString = nil
			} else {
				match++
			}
		} else {
			obj.SpansGroupByTotalString = nil
		}
	} else {
		obj.SpansGroupByTotalString = nil
	}

	// try to unmarshal data into SpansGroupByTotalNumber
	err = datadog.Unmarshal(data, &obj.SpansGroupByTotalNumber)
	if err == nil {
		if obj.SpansGroupByTotalNumber != nil {
			jsonSpansGroupByTotalNumber, _ := datadog.Marshal(obj.SpansGroupByTotalNumber)
			if string(jsonSpansGroupByTotalNumber) == "{}" { // empty struct
				obj.SpansGroupByTotalNumber = nil
			} else {
				match++
			}
		} else {
			obj.SpansGroupByTotalNumber = nil
		}
	} else {
		obj.SpansGroupByTotalNumber = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.SpansGroupByTotalBoolean = nil
		obj.SpansGroupByTotalString = nil
		obj.SpansGroupByTotalNumber = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj SpansGroupByTotal) MarshalJSON() ([]byte, error) {
	if obj.SpansGroupByTotalBoolean != nil {
		return datadog.Marshal(&obj.SpansGroupByTotalBoolean)
	}

	if obj.SpansGroupByTotalString != nil {
		return datadog.Marshal(&obj.SpansGroupByTotalString)
	}

	if obj.SpansGroupByTotalNumber != nil {
		return datadog.Marshal(&obj.SpansGroupByTotalNumber)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *SpansGroupByTotal) GetActualInstance() interface{} {
	if obj.SpansGroupByTotalBoolean != nil {
		return obj.SpansGroupByTotalBoolean
	}

	if obj.SpansGroupByTotalString != nil {
		return obj.SpansGroupByTotalString
	}

	if obj.SpansGroupByTotalNumber != nil {
		return obj.SpansGroupByTotalNumber
	}

	// all schemas are nil
	return nil
}
