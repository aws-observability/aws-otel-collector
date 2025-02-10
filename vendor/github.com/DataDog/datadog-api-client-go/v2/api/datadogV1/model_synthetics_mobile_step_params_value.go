// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsMobileStepParamsValue - Values used in the step for in multiple step types.
type SyntheticsMobileStepParamsValue struct {
	SyntheticsMobileStepParamsValueString *string
	SyntheticsMobileStepParamsValueNumber *int64

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SyntheticsMobileStepParamsValueStringAsSyntheticsMobileStepParamsValue is a convenience function that returns string wrapped in SyntheticsMobileStepParamsValue.
func SyntheticsMobileStepParamsValueStringAsSyntheticsMobileStepParamsValue(v *string) SyntheticsMobileStepParamsValue {
	return SyntheticsMobileStepParamsValue{SyntheticsMobileStepParamsValueString: v}
}

// SyntheticsMobileStepParamsValueNumberAsSyntheticsMobileStepParamsValue is a convenience function that returns int64 wrapped in SyntheticsMobileStepParamsValue.
func SyntheticsMobileStepParamsValueNumberAsSyntheticsMobileStepParamsValue(v *int64) SyntheticsMobileStepParamsValue {
	return SyntheticsMobileStepParamsValue{SyntheticsMobileStepParamsValueNumber: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *SyntheticsMobileStepParamsValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SyntheticsMobileStepParamsValueString
	err = datadog.Unmarshal(data, &obj.SyntheticsMobileStepParamsValueString)
	if err == nil {
		if obj.SyntheticsMobileStepParamsValueString != nil {
			jsonSyntheticsMobileStepParamsValueString, _ := datadog.Marshal(obj.SyntheticsMobileStepParamsValueString)
			if string(jsonSyntheticsMobileStepParamsValueString) == "{}" { // empty struct
				obj.SyntheticsMobileStepParamsValueString = nil
			} else {
				match++
			}
		} else {
			obj.SyntheticsMobileStepParamsValueString = nil
		}
	} else {
		obj.SyntheticsMobileStepParamsValueString = nil
	}

	// try to unmarshal data into SyntheticsMobileStepParamsValueNumber
	err = datadog.Unmarshal(data, &obj.SyntheticsMobileStepParamsValueNumber)
	if err == nil {
		if obj.SyntheticsMobileStepParamsValueNumber != nil {
			jsonSyntheticsMobileStepParamsValueNumber, _ := datadog.Marshal(obj.SyntheticsMobileStepParamsValueNumber)
			if string(jsonSyntheticsMobileStepParamsValueNumber) == "{}" { // empty struct
				obj.SyntheticsMobileStepParamsValueNumber = nil
			} else {
				match++
			}
		} else {
			obj.SyntheticsMobileStepParamsValueNumber = nil
		}
	} else {
		obj.SyntheticsMobileStepParamsValueNumber = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.SyntheticsMobileStepParamsValueString = nil
		obj.SyntheticsMobileStepParamsValueNumber = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj SyntheticsMobileStepParamsValue) MarshalJSON() ([]byte, error) {
	if obj.SyntheticsMobileStepParamsValueString != nil {
		return datadog.Marshal(&obj.SyntheticsMobileStepParamsValueString)
	}

	if obj.SyntheticsMobileStepParamsValueNumber != nil {
		return datadog.Marshal(&obj.SyntheticsMobileStepParamsValueNumber)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *SyntheticsMobileStepParamsValue) GetActualInstance() interface{} {
	if obj.SyntheticsMobileStepParamsValueString != nil {
		return obj.SyntheticsMobileStepParamsValueString
	}

	if obj.SyntheticsMobileStepParamsValueNumber != nil {
		return obj.SyntheticsMobileStepParamsValueNumber
	}

	// all schemas are nil
	return nil
}
