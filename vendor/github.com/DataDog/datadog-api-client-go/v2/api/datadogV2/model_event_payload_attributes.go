// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventPayloadAttributes - JSON object for category-specific attributes. Schema is different per event category.
type EventPayloadAttributes struct {
	ChangeEventCustomAttributes *ChangeEventCustomAttributes
	AlertEventCustomAttributes  *AlertEventCustomAttributes

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ChangeEventCustomAttributesAsEventPayloadAttributes is a convenience function that returns ChangeEventCustomAttributes wrapped in EventPayloadAttributes.
func ChangeEventCustomAttributesAsEventPayloadAttributes(v *ChangeEventCustomAttributes) EventPayloadAttributes {
	return EventPayloadAttributes{ChangeEventCustomAttributes: v}
}

// AlertEventCustomAttributesAsEventPayloadAttributes is a convenience function that returns AlertEventCustomAttributes wrapped in EventPayloadAttributes.
func AlertEventCustomAttributesAsEventPayloadAttributes(v *AlertEventCustomAttributes) EventPayloadAttributes {
	return EventPayloadAttributes{AlertEventCustomAttributes: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *EventPayloadAttributes) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ChangeEventCustomAttributes
	err = datadog.Unmarshal(data, &obj.ChangeEventCustomAttributes)
	if err == nil {
		if obj.ChangeEventCustomAttributes != nil && obj.ChangeEventCustomAttributes.UnparsedObject == nil {
			jsonChangeEventCustomAttributes, _ := datadog.Marshal(obj.ChangeEventCustomAttributes)
			if string(jsonChangeEventCustomAttributes) == "{}" { // empty struct
				obj.ChangeEventCustomAttributes = nil
			} else {
				match++
			}
		} else {
			obj.ChangeEventCustomAttributes = nil
		}
	} else {
		obj.ChangeEventCustomAttributes = nil
	}

	// try to unmarshal data into AlertEventCustomAttributes
	err = datadog.Unmarshal(data, &obj.AlertEventCustomAttributes)
	if err == nil {
		if obj.AlertEventCustomAttributes != nil && obj.AlertEventCustomAttributes.UnparsedObject == nil {
			jsonAlertEventCustomAttributes, _ := datadog.Marshal(obj.AlertEventCustomAttributes)
			if string(jsonAlertEventCustomAttributes) == "{}" { // empty struct
				obj.AlertEventCustomAttributes = nil
			} else {
				match++
			}
		} else {
			obj.AlertEventCustomAttributes = nil
		}
	} else {
		obj.AlertEventCustomAttributes = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ChangeEventCustomAttributes = nil
		obj.AlertEventCustomAttributes = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj EventPayloadAttributes) MarshalJSON() ([]byte, error) {
	if obj.ChangeEventCustomAttributes != nil {
		return datadog.Marshal(&obj.ChangeEventCustomAttributes)
	}

	if obj.AlertEventCustomAttributes != nil {
		return datadog.Marshal(&obj.AlertEventCustomAttributes)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *EventPayloadAttributes) GetActualInstance() interface{} {
	if obj.ChangeEventCustomAttributes != nil {
		return obj.ChangeEventCustomAttributes
	}

	if obj.AlertEventCustomAttributes != nil {
		return obj.AlertEventCustomAttributes
	}

	// all schemas are nil
	return nil
}
