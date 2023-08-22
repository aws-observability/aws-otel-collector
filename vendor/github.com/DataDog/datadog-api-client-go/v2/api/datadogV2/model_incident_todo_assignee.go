// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
)

// IncidentTodoAssignee - A todo assignee.
type IncidentTodoAssignee struct {
	IncidentTodoAssigneeHandle    *string
	IncidentTodoAnonymousAssignee *IncidentTodoAnonymousAssignee

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// IncidentTodoAssigneeHandleAsIncidentTodoAssignee is a convenience function that returns string wrapped in IncidentTodoAssignee.
func IncidentTodoAssigneeHandleAsIncidentTodoAssignee(v *string) IncidentTodoAssignee {
	return IncidentTodoAssignee{IncidentTodoAssigneeHandle: v}
}

// IncidentTodoAnonymousAssigneeAsIncidentTodoAssignee is a convenience function that returns IncidentTodoAnonymousAssignee wrapped in IncidentTodoAssignee.
func IncidentTodoAnonymousAssigneeAsIncidentTodoAssignee(v *IncidentTodoAnonymousAssignee) IncidentTodoAssignee {
	return IncidentTodoAssignee{IncidentTodoAnonymousAssignee: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *IncidentTodoAssignee) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into IncidentTodoAssigneeHandle
	err = json.Unmarshal(data, &obj.IncidentTodoAssigneeHandle)
	if err == nil {
		if obj.IncidentTodoAssigneeHandle != nil {
			jsonIncidentTodoAssigneeHandle, _ := json.Marshal(obj.IncidentTodoAssigneeHandle)
			if string(jsonIncidentTodoAssigneeHandle) == "{}" { // empty struct
				obj.IncidentTodoAssigneeHandle = nil
			} else {
				match++
			}
		} else {
			obj.IncidentTodoAssigneeHandle = nil
		}
	} else {
		obj.IncidentTodoAssigneeHandle = nil
	}

	// try to unmarshal data into IncidentTodoAnonymousAssignee
	err = json.Unmarshal(data, &obj.IncidentTodoAnonymousAssignee)
	if err == nil {
		if obj.IncidentTodoAnonymousAssignee != nil && obj.IncidentTodoAnonymousAssignee.UnparsedObject == nil {
			jsonIncidentTodoAnonymousAssignee, _ := json.Marshal(obj.IncidentTodoAnonymousAssignee)
			if string(jsonIncidentTodoAnonymousAssignee) == "{}" { // empty struct
				obj.IncidentTodoAnonymousAssignee = nil
			} else {
				match++
			}
		} else {
			obj.IncidentTodoAnonymousAssignee = nil
		}
	} else {
		obj.IncidentTodoAnonymousAssignee = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.IncidentTodoAssigneeHandle = nil
		obj.IncidentTodoAnonymousAssignee = nil
		return json.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj IncidentTodoAssignee) MarshalJSON() ([]byte, error) {
	if obj.IncidentTodoAssigneeHandle != nil {
		return json.Marshal(&obj.IncidentTodoAssigneeHandle)
	}

	if obj.IncidentTodoAnonymousAssignee != nil {
		return json.Marshal(&obj.IncidentTodoAnonymousAssignee)
	}

	if obj.UnparsedObject != nil {
		return json.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *IncidentTodoAssignee) GetActualInstance() interface{} {
	if obj.IncidentTodoAssigneeHandle != nil {
		return obj.IncidentTodoAssigneeHandle
	}

	if obj.IncidentTodoAnonymousAssignee != nil {
		return obj.IncidentTodoAnonymousAssignee
	}

	// all schemas are nil
	return nil
}

// NullableIncidentTodoAssignee handles when a null is used for IncidentTodoAssignee.
type NullableIncidentTodoAssignee struct {
	value *IncidentTodoAssignee
	isSet bool
}

// Get returns the associated value.
func (v NullableIncidentTodoAssignee) Get() *IncidentTodoAssignee {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableIncidentTodoAssignee) Set(val *IncidentTodoAssignee) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableIncidentTodoAssignee) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableIncidentTodoAssignee) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableIncidentTodoAssignee initializes the struct as if Set has been called.
func NewNullableIncidentTodoAssignee(val *IncidentTodoAssignee) *NullableIncidentTodoAssignee {
	return &NullableIncidentTodoAssignee{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableIncidentTodoAssignee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableIncidentTodoAssignee) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return json.Unmarshal(src, &v.value)
}
