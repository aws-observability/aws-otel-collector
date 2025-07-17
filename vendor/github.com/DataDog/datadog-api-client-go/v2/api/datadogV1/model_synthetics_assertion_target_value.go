// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsAssertionTargetValue - Value used by the operator in assertions. Can be either a number or string.
type SyntheticsAssertionTargetValue struct {
	SyntheticsAssertionTargetValueNumber *float64
	SyntheticsAssertionTargetValueString *string

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SyntheticsAssertionTargetValueNumberAsSyntheticsAssertionTargetValue is a convenience function that returns float64 wrapped in SyntheticsAssertionTargetValue.
func SyntheticsAssertionTargetValueNumberAsSyntheticsAssertionTargetValue(v *float64) SyntheticsAssertionTargetValue {
	return SyntheticsAssertionTargetValue{SyntheticsAssertionTargetValueNumber: v}
}

// SyntheticsAssertionTargetValueStringAsSyntheticsAssertionTargetValue is a convenience function that returns string wrapped in SyntheticsAssertionTargetValue.
func SyntheticsAssertionTargetValueStringAsSyntheticsAssertionTargetValue(v *string) SyntheticsAssertionTargetValue {
	return SyntheticsAssertionTargetValue{SyntheticsAssertionTargetValueString: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *SyntheticsAssertionTargetValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SyntheticsAssertionTargetValueNumber
	err = datadog.Unmarshal(data, &obj.SyntheticsAssertionTargetValueNumber)
	if err == nil {
		if obj.SyntheticsAssertionTargetValueNumber != nil {
			jsonSyntheticsAssertionTargetValueNumber, _ := datadog.Marshal(obj.SyntheticsAssertionTargetValueNumber)
			if string(jsonSyntheticsAssertionTargetValueNumber) == "{}" { // empty struct
				obj.SyntheticsAssertionTargetValueNumber = nil
			} else {
				match++
			}
		} else {
			obj.SyntheticsAssertionTargetValueNumber = nil
		}
	} else {
		obj.SyntheticsAssertionTargetValueNumber = nil
	}

	// try to unmarshal data into SyntheticsAssertionTargetValueString
	err = datadog.Unmarshal(data, &obj.SyntheticsAssertionTargetValueString)
	if err == nil {
		if obj.SyntheticsAssertionTargetValueString != nil {
			jsonSyntheticsAssertionTargetValueString, _ := datadog.Marshal(obj.SyntheticsAssertionTargetValueString)
			if string(jsonSyntheticsAssertionTargetValueString) == "{}" { // empty struct
				obj.SyntheticsAssertionTargetValueString = nil
			} else {
				match++
			}
		} else {
			obj.SyntheticsAssertionTargetValueString = nil
		}
	} else {
		obj.SyntheticsAssertionTargetValueString = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.SyntheticsAssertionTargetValueNumber = nil
		obj.SyntheticsAssertionTargetValueString = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj SyntheticsAssertionTargetValue) MarshalJSON() ([]byte, error) {
	if obj.SyntheticsAssertionTargetValueNumber != nil {
		return datadog.Marshal(&obj.SyntheticsAssertionTargetValueNumber)
	}

	if obj.SyntheticsAssertionTargetValueString != nil {
		return datadog.Marshal(&obj.SyntheticsAssertionTargetValueString)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *SyntheticsAssertionTargetValue) GetActualInstance() interface{} {
	if obj.SyntheticsAssertionTargetValueNumber != nil {
		return obj.SyntheticsAssertionTargetValueNumber
	}

	if obj.SyntheticsAssertionTargetValueString != nil {
		return obj.SyntheticsAssertionTargetValueString
	}

	// all schemas are nil
	return nil
}
