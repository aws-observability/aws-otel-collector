// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentImportFieldAttributes - Dynamic fields for which selections can be made, with field names as keys.
type IncidentImportFieldAttributes struct {
	IncidentImportFieldAttributesSingleValue   *IncidentImportFieldAttributesSingleValue
	IncidentImportFieldAttributesMultipleValue *IncidentImportFieldAttributesMultipleValue

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// IncidentImportFieldAttributesSingleValueAsIncidentImportFieldAttributes is a convenience function that returns IncidentImportFieldAttributesSingleValue wrapped in IncidentImportFieldAttributes.
func IncidentImportFieldAttributesSingleValueAsIncidentImportFieldAttributes(v *IncidentImportFieldAttributesSingleValue) IncidentImportFieldAttributes {
	return IncidentImportFieldAttributes{IncidentImportFieldAttributesSingleValue: v}
}

// IncidentImportFieldAttributesMultipleValueAsIncidentImportFieldAttributes is a convenience function that returns IncidentImportFieldAttributesMultipleValue wrapped in IncidentImportFieldAttributes.
func IncidentImportFieldAttributesMultipleValueAsIncidentImportFieldAttributes(v *IncidentImportFieldAttributesMultipleValue) IncidentImportFieldAttributes {
	return IncidentImportFieldAttributes{IncidentImportFieldAttributesMultipleValue: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *IncidentImportFieldAttributes) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into IncidentImportFieldAttributesSingleValue
	err = datadog.Unmarshal(data, &obj.IncidentImportFieldAttributesSingleValue)
	if err == nil {
		if obj.IncidentImportFieldAttributesSingleValue != nil && obj.IncidentImportFieldAttributesSingleValue.UnparsedObject == nil {
			jsonIncidentImportFieldAttributesSingleValue, _ := datadog.Marshal(obj.IncidentImportFieldAttributesSingleValue)
			if string(jsonIncidentImportFieldAttributesSingleValue) == "{}" && string(data) != "{}" { // empty struct
				obj.IncidentImportFieldAttributesSingleValue = nil
			} else {
				match++
			}
		} else {
			obj.IncidentImportFieldAttributesSingleValue = nil
		}
	} else {
		obj.IncidentImportFieldAttributesSingleValue = nil
	}

	// try to unmarshal data into IncidentImportFieldAttributesMultipleValue
	err = datadog.Unmarshal(data, &obj.IncidentImportFieldAttributesMultipleValue)
	if err == nil {
		if obj.IncidentImportFieldAttributesMultipleValue != nil && obj.IncidentImportFieldAttributesMultipleValue.UnparsedObject == nil {
			jsonIncidentImportFieldAttributesMultipleValue, _ := datadog.Marshal(obj.IncidentImportFieldAttributesMultipleValue)
			if string(jsonIncidentImportFieldAttributesMultipleValue) == "{}" && string(data) != "{}" { // empty struct
				obj.IncidentImportFieldAttributesMultipleValue = nil
			} else {
				match++
			}
		} else {
			obj.IncidentImportFieldAttributesMultipleValue = nil
		}
	} else {
		obj.IncidentImportFieldAttributesMultipleValue = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.IncidentImportFieldAttributesSingleValue = nil
		obj.IncidentImportFieldAttributesMultipleValue = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj IncidentImportFieldAttributes) MarshalJSON() ([]byte, error) {
	if obj.IncidentImportFieldAttributesSingleValue != nil {
		return datadog.Marshal(&obj.IncidentImportFieldAttributesSingleValue)
	}

	if obj.IncidentImportFieldAttributesMultipleValue != nil {
		return datadog.Marshal(&obj.IncidentImportFieldAttributesMultipleValue)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *IncidentImportFieldAttributes) GetActualInstance() interface{} {
	if obj.IncidentImportFieldAttributesSingleValue != nil {
		return obj.IncidentImportFieldAttributesSingleValue
	}

	if obj.IncidentImportFieldAttributesMultipleValue != nil {
		return obj.IncidentImportFieldAttributesMultipleValue
	}

	// all schemas are nil
	return nil
}
