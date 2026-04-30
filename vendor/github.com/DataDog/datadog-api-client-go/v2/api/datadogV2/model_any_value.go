// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnyValue - Represents any valid JSON value.
type AnyValue struct {
	AnyValueString  *string
	AnyValueNumber  *float64
	AnyValueObject  map[string]interface{}
	AnyValueArray   *[]AnyValueItem
	AnyValueBoolean *bool

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// AnyValueStringAsAnyValue is a convenience function that returns string wrapped in AnyValue.
func AnyValueStringAsAnyValue(v *string) AnyValue {
	return AnyValue{AnyValueString: v}
}

// AnyValueNumberAsAnyValue is a convenience function that returns float64 wrapped in AnyValue.
func AnyValueNumberAsAnyValue(v *float64) AnyValue {
	return AnyValue{AnyValueNumber: v}
}

// AnyValueObjectAsAnyValue is a convenience function that returns map[string]interface{} wrapped in AnyValue.
func AnyValueObjectAsAnyValue(v map[string]interface{}) AnyValue {
	return AnyValue{AnyValueObject: v}
}

// AnyValueArrayAsAnyValue is a convenience function that returns []AnyValueItem wrapped in AnyValue.
func AnyValueArrayAsAnyValue(v *[]AnyValueItem) AnyValue {
	return AnyValue{AnyValueArray: v}
}

// AnyValueBooleanAsAnyValue is a convenience function that returns bool wrapped in AnyValue.
func AnyValueBooleanAsAnyValue(v *bool) AnyValue {
	return AnyValue{AnyValueBoolean: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *AnyValue) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into AnyValueArray
	err = datadog.Unmarshal(data, &obj.AnyValueArray)
	if err == nil {
		if obj.AnyValueArray != nil {
			jsonAnyValueArray, _ := datadog.Marshal(obj.AnyValueArray)
			if string(jsonAnyValueArray) == "{}" && string(data) != "{}" { // empty struct
				obj.AnyValueArray = nil
			} else {
				match++
			}
		} else {
			obj.AnyValueArray = nil
		}
	} else {
		obj.AnyValueArray = nil
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
		obj.AnyValueArray = nil
		obj.AnyValueBoolean = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj AnyValue) MarshalJSON() ([]byte, error) {
	if obj.AnyValueString != nil {
		return datadog.Marshal(&obj.AnyValueString)
	}

	if obj.AnyValueNumber != nil {
		return datadog.Marshal(&obj.AnyValueNumber)
	}

	if obj.AnyValueObject != nil {
		return datadog.Marshal(&obj.AnyValueObject)
	}

	if obj.AnyValueArray != nil {
		return datadog.Marshal(&obj.AnyValueArray)
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
func (obj *AnyValue) GetActualInstance() interface{} {
	if obj.AnyValueString != nil {
		return obj.AnyValueString
	}

	if obj.AnyValueNumber != nil {
		return obj.AnyValueNumber
	}

	if obj.AnyValueObject != nil {
		return obj.AnyValueObject
	}

	if obj.AnyValueArray != nil {
		return obj.AnyValueArray
	}

	if obj.AnyValueBoolean != nil {
		return obj.AnyValueBoolean
	}

	// all schemas are nil
	return nil
}

// NullableAnyValue handles when a null is used for AnyValue.
type NullableAnyValue struct {
	value *AnyValue
	isSet bool
}

// Get returns the associated value.
func (v NullableAnyValue) Get() *AnyValue {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableAnyValue) Set(val *AnyValue) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableAnyValue) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableAnyValue) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableAnyValue initializes the struct as if Set has been called.
func NewNullableAnyValue(val *AnyValue) *NullableAnyValue {
	return &NullableAnyValue{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableAnyValue) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableAnyValue) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
