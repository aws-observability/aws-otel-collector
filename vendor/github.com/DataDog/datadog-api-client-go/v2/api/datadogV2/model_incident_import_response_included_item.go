// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentImportResponseIncludedItem - An object related to an incident that is included in the response.
type IncidentImportResponseIncludedItem struct {
	IncidentUserData   *IncidentUserData
	IncidentTypeObject *IncidentTypeObject

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// IncidentUserDataAsIncidentImportResponseIncludedItem is a convenience function that returns IncidentUserData wrapped in IncidentImportResponseIncludedItem.
func IncidentUserDataAsIncidentImportResponseIncludedItem(v *IncidentUserData) IncidentImportResponseIncludedItem {
	return IncidentImportResponseIncludedItem{IncidentUserData: v}
}

// IncidentTypeObjectAsIncidentImportResponseIncludedItem is a convenience function that returns IncidentTypeObject wrapped in IncidentImportResponseIncludedItem.
func IncidentTypeObjectAsIncidentImportResponseIncludedItem(v *IncidentTypeObject) IncidentImportResponseIncludedItem {
	return IncidentImportResponseIncludedItem{IncidentTypeObject: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *IncidentImportResponseIncludedItem) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into IncidentUserData
	err = datadog.Unmarshal(data, &obj.IncidentUserData)
	if err == nil {
		if obj.IncidentUserData != nil && obj.IncidentUserData.UnparsedObject == nil {
			jsonIncidentUserData, _ := datadog.Marshal(obj.IncidentUserData)
			if string(jsonIncidentUserData) == "{}" && string(data) != "{}" { // empty struct
				obj.IncidentUserData = nil
			} else {
				match++
			}
		} else {
			obj.IncidentUserData = nil
		}
	} else {
		obj.IncidentUserData = nil
	}

	// try to unmarshal data into IncidentTypeObject
	err = datadog.Unmarshal(data, &obj.IncidentTypeObject)
	if err == nil {
		if obj.IncidentTypeObject != nil && obj.IncidentTypeObject.UnparsedObject == nil {
			jsonIncidentTypeObject, _ := datadog.Marshal(obj.IncidentTypeObject)
			if string(jsonIncidentTypeObject) == "{}" { // empty struct
				obj.IncidentTypeObject = nil
			} else {
				match++
			}
		} else {
			obj.IncidentTypeObject = nil
		}
	} else {
		obj.IncidentTypeObject = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.IncidentUserData = nil
		obj.IncidentTypeObject = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj IncidentImportResponseIncludedItem) MarshalJSON() ([]byte, error) {
	if obj.IncidentUserData != nil {
		return datadog.Marshal(&obj.IncidentUserData)
	}

	if obj.IncidentTypeObject != nil {
		return datadog.Marshal(&obj.IncidentTypeObject)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *IncidentImportResponseIncludedItem) GetActualInstance() interface{} {
	if obj.IncidentUserData != nil {
		return obj.IncidentUserData
	}

	if obj.IncidentTypeObject != nil {
		return obj.IncidentTypeObject
	}

	// all schemas are nil
	return nil
}
