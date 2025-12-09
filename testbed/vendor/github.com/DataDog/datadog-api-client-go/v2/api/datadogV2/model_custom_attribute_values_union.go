// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomAttributeValuesUnion - Union of supported value for a custom attribute
type CustomAttributeValuesUnion struct {
	CustomAttributeStringValue      *string
	CustomAttributeMultiStringValue *[]string
	CustomAttributeNumberValue      *float64
	CustomAttributeMultiNumberValue *[]float64

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// CustomAttributeStringValueAsCustomAttributeValuesUnion is a convenience function that returns string wrapped in CustomAttributeValuesUnion.
func CustomAttributeStringValueAsCustomAttributeValuesUnion(v *string) CustomAttributeValuesUnion {
	return CustomAttributeValuesUnion{CustomAttributeStringValue: v}
}

// CustomAttributeMultiStringValueAsCustomAttributeValuesUnion is a convenience function that returns []string wrapped in CustomAttributeValuesUnion.
func CustomAttributeMultiStringValueAsCustomAttributeValuesUnion(v *[]string) CustomAttributeValuesUnion {
	return CustomAttributeValuesUnion{CustomAttributeMultiStringValue: v}
}

// CustomAttributeNumberValueAsCustomAttributeValuesUnion is a convenience function that returns float64 wrapped in CustomAttributeValuesUnion.
func CustomAttributeNumberValueAsCustomAttributeValuesUnion(v *float64) CustomAttributeValuesUnion {
	return CustomAttributeValuesUnion{CustomAttributeNumberValue: v}
}

// CustomAttributeMultiNumberValueAsCustomAttributeValuesUnion is a convenience function that returns []float64 wrapped in CustomAttributeValuesUnion.
func CustomAttributeMultiNumberValueAsCustomAttributeValuesUnion(v *[]float64) CustomAttributeValuesUnion {
	return CustomAttributeValuesUnion{CustomAttributeMultiNumberValue: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *CustomAttributeValuesUnion) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CustomAttributeStringValue
	err = datadog.Unmarshal(data, &obj.CustomAttributeStringValue)
	if err == nil {
		if obj.CustomAttributeStringValue != nil {
			jsonCustomAttributeStringValue, _ := datadog.Marshal(obj.CustomAttributeStringValue)
			if string(jsonCustomAttributeStringValue) == "{}" { // empty struct
				obj.CustomAttributeStringValue = nil
			} else {
				match++
			}
		} else {
			obj.CustomAttributeStringValue = nil
		}
	} else {
		obj.CustomAttributeStringValue = nil
	}

	// try to unmarshal data into CustomAttributeMultiStringValue
	err = datadog.Unmarshal(data, &obj.CustomAttributeMultiStringValue)
	if err == nil {
		if obj.CustomAttributeMultiStringValue != nil {
			jsonCustomAttributeMultiStringValue, _ := datadog.Marshal(obj.CustomAttributeMultiStringValue)
			if string(jsonCustomAttributeMultiStringValue) == "{}" && string(data) != "{}" { // empty struct
				obj.CustomAttributeMultiStringValue = nil
			} else {
				match++
			}
		} else {
			obj.CustomAttributeMultiStringValue = nil
		}
	} else {
		obj.CustomAttributeMultiStringValue = nil
	}

	// try to unmarshal data into CustomAttributeNumberValue
	err = datadog.Unmarshal(data, &obj.CustomAttributeNumberValue)
	if err == nil {
		if obj.CustomAttributeNumberValue != nil {
			jsonCustomAttributeNumberValue, _ := datadog.Marshal(obj.CustomAttributeNumberValue)
			if string(jsonCustomAttributeNumberValue) == "{}" { // empty struct
				obj.CustomAttributeNumberValue = nil
			} else {
				match++
			}
		} else {
			obj.CustomAttributeNumberValue = nil
		}
	} else {
		obj.CustomAttributeNumberValue = nil
	}

	// try to unmarshal data into CustomAttributeMultiNumberValue
	err = datadog.Unmarshal(data, &obj.CustomAttributeMultiNumberValue)
	if err == nil {
		if obj.CustomAttributeMultiNumberValue != nil {
			jsonCustomAttributeMultiNumberValue, _ := datadog.Marshal(obj.CustomAttributeMultiNumberValue)
			if string(jsonCustomAttributeMultiNumberValue) == "{}" && string(data) != "{}" { // empty struct
				obj.CustomAttributeMultiNumberValue = nil
			} else {
				match++
			}
		} else {
			obj.CustomAttributeMultiNumberValue = nil
		}
	} else {
		obj.CustomAttributeMultiNumberValue = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.CustomAttributeStringValue = nil
		obj.CustomAttributeMultiStringValue = nil
		obj.CustomAttributeNumberValue = nil
		obj.CustomAttributeMultiNumberValue = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj CustomAttributeValuesUnion) MarshalJSON() ([]byte, error) {
	if obj.CustomAttributeStringValue != nil {
		return datadog.Marshal(&obj.CustomAttributeStringValue)
	}

	if obj.CustomAttributeMultiStringValue != nil {
		return datadog.Marshal(&obj.CustomAttributeMultiStringValue)
	}

	if obj.CustomAttributeNumberValue != nil {
		return datadog.Marshal(&obj.CustomAttributeNumberValue)
	}

	if obj.CustomAttributeMultiNumberValue != nil {
		return datadog.Marshal(&obj.CustomAttributeMultiNumberValue)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *CustomAttributeValuesUnion) GetActualInstance() interface{} {
	if obj.CustomAttributeStringValue != nil {
		return obj.CustomAttributeStringValue
	}

	if obj.CustomAttributeMultiStringValue != nil {
		return obj.CustomAttributeMultiStringValue
	}

	if obj.CustomAttributeNumberValue != nil {
		return obj.CustomAttributeNumberValue
	}

	if obj.CustomAttributeMultiNumberValue != nil {
		return obj.CustomAttributeMultiNumberValue
	}

	// all schemas are nil
	return nil
}
