// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsAnnotationLabelValueValue - The value for this label. Must comply with the label schema type constraints.
type LLMObsAnnotationLabelValueValue struct {
	AnyValueNumber                        *float64
	AnyValueString                        *string
	LLMObsAnnotationLabelValueStringArray *[]string
	AnyValueBoolean                       *bool

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// AnyValueNumberAsLLMObsAnnotationLabelValueValue is a convenience function that returns float64 wrapped in LLMObsAnnotationLabelValueValue.
func AnyValueNumberAsLLMObsAnnotationLabelValueValue(v *float64) LLMObsAnnotationLabelValueValue {
	return LLMObsAnnotationLabelValueValue{AnyValueNumber: v}
}

// AnyValueStringAsLLMObsAnnotationLabelValueValue is a convenience function that returns string wrapped in LLMObsAnnotationLabelValueValue.
func AnyValueStringAsLLMObsAnnotationLabelValueValue(v *string) LLMObsAnnotationLabelValueValue {
	return LLMObsAnnotationLabelValueValue{AnyValueString: v}
}

// LLMObsAnnotationLabelValueStringArrayAsLLMObsAnnotationLabelValueValue is a convenience function that returns []string wrapped in LLMObsAnnotationLabelValueValue.
func LLMObsAnnotationLabelValueStringArrayAsLLMObsAnnotationLabelValueValue(v *[]string) LLMObsAnnotationLabelValueValue {
	return LLMObsAnnotationLabelValueValue{LLMObsAnnotationLabelValueStringArray: v}
}

// AnyValueBooleanAsLLMObsAnnotationLabelValueValue is a convenience function that returns bool wrapped in LLMObsAnnotationLabelValueValue.
func AnyValueBooleanAsLLMObsAnnotationLabelValueValue(v *bool) LLMObsAnnotationLabelValueValue {
	return LLMObsAnnotationLabelValueValue{AnyValueBoolean: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *LLMObsAnnotationLabelValueValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
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

	// try to unmarshal data into LLMObsAnnotationLabelValueStringArray
	err = datadog.Unmarshal(data, &obj.LLMObsAnnotationLabelValueStringArray)
	if err == nil {
		if obj.LLMObsAnnotationLabelValueStringArray != nil {
			jsonLLMObsAnnotationLabelValueStringArray, _ := datadog.Marshal(obj.LLMObsAnnotationLabelValueStringArray)
			if string(jsonLLMObsAnnotationLabelValueStringArray) == "{}" && string(data) != "{}" { // empty struct
				obj.LLMObsAnnotationLabelValueStringArray = nil
			} else {
				match++
			}
		} else {
			obj.LLMObsAnnotationLabelValueStringArray = nil
		}
	} else {
		obj.LLMObsAnnotationLabelValueStringArray = nil
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
		obj.AnyValueNumber = nil
		obj.AnyValueString = nil
		obj.LLMObsAnnotationLabelValueStringArray = nil
		obj.AnyValueBoolean = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj LLMObsAnnotationLabelValueValue) MarshalJSON() ([]byte, error) {
	if obj.AnyValueNumber != nil {
		return datadog.Marshal(&obj.AnyValueNumber)
	}

	if obj.AnyValueString != nil {
		return datadog.Marshal(&obj.AnyValueString)
	}

	if obj.LLMObsAnnotationLabelValueStringArray != nil {
		return datadog.Marshal(&obj.LLMObsAnnotationLabelValueStringArray)
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
func (obj *LLMObsAnnotationLabelValueValue) GetActualInstance() interface{} {
	if obj.AnyValueNumber != nil {
		return obj.AnyValueNumber
	}

	if obj.AnyValueString != nil {
		return obj.AnyValueString
	}

	if obj.LLMObsAnnotationLabelValueStringArray != nil {
		return obj.LLMObsAnnotationLabelValueStringArray
	}

	if obj.AnyValueBoolean != nil {
		return obj.AnyValueBoolean
	}

	// all schemas are nil
	return nil
}
