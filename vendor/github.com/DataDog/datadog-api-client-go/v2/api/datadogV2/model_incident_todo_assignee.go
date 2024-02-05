// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
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
	err = datadog.Unmarshal(data, &obj.IncidentTodoAssigneeHandle)
	if err == nil {
		if obj.IncidentTodoAssigneeHandle != nil {
			jsonIncidentTodoAssigneeHandle, _ := datadog.Marshal(obj.IncidentTodoAssigneeHandle)
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
	err = datadog.Unmarshal(data, &obj.IncidentTodoAnonymousAssignee)
	if err == nil {
		if obj.IncidentTodoAnonymousAssignee != nil && obj.IncidentTodoAnonymousAssignee.UnparsedObject == nil {
			jsonIncidentTodoAnonymousAssignee, _ := datadog.Marshal(obj.IncidentTodoAnonymousAssignee)
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
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj IncidentTodoAssignee) MarshalJSON() ([]byte, error) {
	if obj.IncidentTodoAssigneeHandle != nil {
		return datadog.Marshal(&obj.IncidentTodoAssigneeHandle)
	}

	if obj.IncidentTodoAnonymousAssignee != nil {
		return datadog.Marshal(&obj.IncidentTodoAnonymousAssignee)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
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
