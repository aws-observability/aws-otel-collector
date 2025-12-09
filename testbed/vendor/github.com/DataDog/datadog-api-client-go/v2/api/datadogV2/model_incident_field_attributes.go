// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentFieldAttributes - Dynamic fields for which selections can be made, with field names as keys.
type IncidentFieldAttributes struct {
	IncidentFieldAttributesSingleValue   *IncidentFieldAttributesSingleValue
	IncidentFieldAttributesMultipleValue *IncidentFieldAttributesMultipleValue

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// IncidentFieldAttributesSingleValueAsIncidentFieldAttributes is a convenience function that returns IncidentFieldAttributesSingleValue wrapped in IncidentFieldAttributes.
func IncidentFieldAttributesSingleValueAsIncidentFieldAttributes(v *IncidentFieldAttributesSingleValue) IncidentFieldAttributes {
	return IncidentFieldAttributes{IncidentFieldAttributesSingleValue: v}
}

// IncidentFieldAttributesMultipleValueAsIncidentFieldAttributes is a convenience function that returns IncidentFieldAttributesMultipleValue wrapped in IncidentFieldAttributes.
func IncidentFieldAttributesMultipleValueAsIncidentFieldAttributes(v *IncidentFieldAttributesMultipleValue) IncidentFieldAttributes {
	return IncidentFieldAttributes{IncidentFieldAttributesMultipleValue: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *IncidentFieldAttributes) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into IncidentFieldAttributesSingleValue
	err = datadog.Unmarshal(data, &obj.IncidentFieldAttributesSingleValue)
	if err == nil {
		if obj.IncidentFieldAttributesSingleValue != nil && obj.IncidentFieldAttributesSingleValue.UnparsedObject == nil {
			jsonIncidentFieldAttributesSingleValue, _ := datadog.Marshal(obj.IncidentFieldAttributesSingleValue)
			if string(jsonIncidentFieldAttributesSingleValue) == "{}" && string(data) != "{}" { // empty struct
				obj.IncidentFieldAttributesSingleValue = nil
			} else {
				match++
			}
		} else {
			obj.IncidentFieldAttributesSingleValue = nil
		}
	} else {
		obj.IncidentFieldAttributesSingleValue = nil
	}

	// try to unmarshal data into IncidentFieldAttributesMultipleValue
	err = datadog.Unmarshal(data, &obj.IncidentFieldAttributesMultipleValue)
	if err == nil {
		if obj.IncidentFieldAttributesMultipleValue != nil && obj.IncidentFieldAttributesMultipleValue.UnparsedObject == nil {
			jsonIncidentFieldAttributesMultipleValue, _ := datadog.Marshal(obj.IncidentFieldAttributesMultipleValue)
			if string(jsonIncidentFieldAttributesMultipleValue) == "{}" && string(data) != "{}" { // empty struct
				obj.IncidentFieldAttributesMultipleValue = nil
			} else {
				match++
			}
		} else {
			obj.IncidentFieldAttributesMultipleValue = nil
		}
	} else {
		obj.IncidentFieldAttributesMultipleValue = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.IncidentFieldAttributesSingleValue = nil
		obj.IncidentFieldAttributesMultipleValue = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj IncidentFieldAttributes) MarshalJSON() ([]byte, error) {
	if obj.IncidentFieldAttributesSingleValue != nil {
		return datadog.Marshal(&obj.IncidentFieldAttributesSingleValue)
	}

	if obj.IncidentFieldAttributesMultipleValue != nil {
		return datadog.Marshal(&obj.IncidentFieldAttributesMultipleValue)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *IncidentFieldAttributes) GetActualInstance() interface{} {
	if obj.IncidentFieldAttributesSingleValue != nil {
		return obj.IncidentFieldAttributesSingleValue
	}

	if obj.IncidentFieldAttributesMultipleValue != nil {
		return obj.IncidentFieldAttributesMultipleValue
	}

	// all schemas are nil
	return nil
}
