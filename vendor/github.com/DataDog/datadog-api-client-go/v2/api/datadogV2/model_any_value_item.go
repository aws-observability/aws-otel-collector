// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnyValueItem - A single item in an array of arbitrary values, which can be a string, number, object, or boolean.
type AnyValueItem struct {
	AnyValueString  *string
	AnyValueNumber  *float64
	AnyValueObject  map[string]interface{}
	AnyValueBoolean *bool

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// AnyValueStringAsAnyValueItem is a convenience function that returns string wrapped in AnyValueItem.
func AnyValueStringAsAnyValueItem(v *string) AnyValueItem {
	return AnyValueItem{AnyValueString: v}
}

// AnyValueNumberAsAnyValueItem is a convenience function that returns float64 wrapped in AnyValueItem.
func AnyValueNumberAsAnyValueItem(v *float64) AnyValueItem {
	return AnyValueItem{AnyValueNumber: v}
}

// AnyValueObjectAsAnyValueItem is a convenience function that returns map[string]interface{} wrapped in AnyValueItem.
func AnyValueObjectAsAnyValueItem(v map[string]interface{}) AnyValueItem {
	return AnyValueItem{AnyValueObject: v}
}

// AnyValueBooleanAsAnyValueItem is a convenience function that returns bool wrapped in AnyValueItem.
func AnyValueBooleanAsAnyValueItem(v *bool) AnyValueItem {
	return AnyValueItem{AnyValueBoolean: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *AnyValueItem) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AnyValueString
	err = datadog.Unmarshal(data, &obj.AnyValueString)
	if err == nil {
		if obj.AnyValueString != nil {
			jsonAnyValueString, _ := datadog.Marshal(obj.AnyValueString)
			if string(jsonAnyValueString) == "{}" { // empty struct
				obj.AnyValueString = nil
			} else {
				match++
			}
		} else {
			obj.AnyValueString = nil
		}
	} else {
		obj.AnyValueString = nil
	}

	// try to unmarshal data into AnyValueNumber
	err = datadog.Unmarshal(data, &obj.AnyValueNumber)
	if err == nil {
		if obj.AnyValueNumber != nil {
			jsonAnyValueNumber, _ := datadog.Marshal(obj.AnyValueNumber)
			if string(jsonAnyValueNumber) == "{}" { // empty struct
				obj.AnyValueNumber = nil
			} else {
				match++
			}
		} else {
			obj.AnyValueNumber = nil
		}
	} else {
		obj.AnyValueNumber = nil
	}

	// try to unmarshal data into AnyValueObject
	err = datadog.Unmarshal(data, &obj.AnyValueObject)
	if err == nil {
		if obj.AnyValueObject != nil {
			jsonAnyValueObject, _ := datadog.Marshal(obj.AnyValueObject)
			if string(jsonAnyValueObject) == "{}" && string(data) != "{}" { // empty struct
				obj.AnyValueObject = nil
			} else {
				match++
			}
		} else {
			obj.AnyValueObject = nil
		}
	} else {
		obj.AnyValueObject = nil
	}

	// try to unmarshal data into AnyValueBoolean
	err = datadog.Unmarshal(data, &obj.AnyValueBoolean)
	if err == nil {
		if obj.AnyValueBoolean != nil {
			jsonAnyValueBoolean, _ := datadog.Marshal(obj.AnyValueBoolean)
			if string(jsonAnyValueBoolean) == "{}" { // empty struct
				obj.AnyValueBoolean = nil
			} else {
				match++
			}
		} else {
			obj.AnyValueBoolean = nil
		}
	} else {
		obj.AnyValueBoolean = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.AnyValueString = nil
		obj.AnyValueNumber = nil
		obj.AnyValueObject = nil
		obj.AnyValueBoolean = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj AnyValueItem) MarshalJSON() ([]byte, error) {
	if obj.AnyValueString != nil {
		return datadog.Marshal(&obj.AnyValueString)
	}

	if obj.AnyValueNumber != nil {
		return datadog.Marshal(&obj.AnyValueNumber)
	}

	if obj.AnyValueObject != nil {
		return datadog.Marshal(&obj.AnyValueObject)
	}

	if obj.AnyValueBoolean != nil {
		return datadog.Marshal(&obj.AnyValueBoolean)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *AnyValueItem) GetActualInstance() interface{} {
	if obj.AnyValueString != nil {
		return obj.AnyValueString
	}

	if obj.AnyValueNumber != nil {
		return obj.AnyValueNumber
	}

	if obj.AnyValueObject != nil {
		return obj.AnyValueObject
	}

	if obj.AnyValueBoolean != nil {
		return obj.AnyValueBoolean
	}

	// all schemas are nil
	return nil
}
