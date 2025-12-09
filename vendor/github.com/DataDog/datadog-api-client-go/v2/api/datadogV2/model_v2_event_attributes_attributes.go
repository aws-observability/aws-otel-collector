// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// V2EventAttributesAttributes - JSON object for category-specific attributes.
type V2EventAttributesAttributes struct {
	ChangeEventAttributes *ChangeEventAttributes
	AlertEventAttributes  *AlertEventAttributes

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ChangeEventAttributesAsV2EventAttributesAttributes is a convenience function that returns ChangeEventAttributes wrapped in V2EventAttributesAttributes.
func ChangeEventAttributesAsV2EventAttributesAttributes(v *ChangeEventAttributes) V2EventAttributesAttributes {
	return V2EventAttributesAttributes{ChangeEventAttributes: v}
}

// AlertEventAttributesAsV2EventAttributesAttributes is a convenience function that returns AlertEventAttributes wrapped in V2EventAttributesAttributes.
func AlertEventAttributesAsV2EventAttributesAttributes(v *AlertEventAttributes) V2EventAttributesAttributes {
	return V2EventAttributesAttributes{AlertEventAttributes: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *V2EventAttributesAttributes) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ChangeEventAttributes
	err = datadog.Unmarshal(data, &obj.ChangeEventAttributes)
	if err == nil {
		if obj.ChangeEventAttributes != nil && obj.ChangeEventAttributes.UnparsedObject == nil {
			jsonChangeEventAttributes, _ := datadog.Marshal(obj.ChangeEventAttributes)
			if string(jsonChangeEventAttributes) == "{}" && string(data) != "{}" { // empty struct
				obj.ChangeEventAttributes = nil
			} else {
				match++
			}
		} else {
			obj.ChangeEventAttributes = nil
		}
	} else {
		obj.ChangeEventAttributes = nil
	}

	// try to unmarshal data into AlertEventAttributes
	err = datadog.Unmarshal(data, &obj.AlertEventAttributes)
	if err == nil {
		if obj.AlertEventAttributes != nil && obj.AlertEventAttributes.UnparsedObject == nil {
			jsonAlertEventAttributes, _ := datadog.Marshal(obj.AlertEventAttributes)
			if string(jsonAlertEventAttributes) == "{}" && string(data) != "{}" { // empty struct
				obj.AlertEventAttributes = nil
			} else {
				match++
			}
		} else {
			obj.AlertEventAttributes = nil
		}
	} else {
		obj.AlertEventAttributes = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ChangeEventAttributes = nil
		obj.AlertEventAttributes = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj V2EventAttributesAttributes) MarshalJSON() ([]byte, error) {
	if obj.ChangeEventAttributes != nil {
		return datadog.Marshal(&obj.ChangeEventAttributes)
	}

	if obj.AlertEventAttributes != nil {
		return datadog.Marshal(&obj.AlertEventAttributes)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *V2EventAttributesAttributes) GetActualInstance() interface{} {
	if obj.ChangeEventAttributes != nil {
		return obj.ChangeEventAttributes
	}

	if obj.AlertEventAttributes != nil {
		return obj.AlertEventAttributes
	}

	// all schemas are nil
	return nil
}
