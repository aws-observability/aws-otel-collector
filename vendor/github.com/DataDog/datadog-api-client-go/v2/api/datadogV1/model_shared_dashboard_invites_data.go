// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"encoding/json"
)

// SharedDashboardInvitesData - An object or list of objects containing the information for an invitation to a shared dashboard.
type SharedDashboardInvitesData struct {
	SharedDashboardInvitesDataObject *SharedDashboardInvitesDataObject
	SharedDashboardInvitesDataList   *[]SharedDashboardInvitesDataObject

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SharedDashboardInvitesDataObjectAsSharedDashboardInvitesData is a convenience function that returns SharedDashboardInvitesDataObject wrapped in SharedDashboardInvitesData.
func SharedDashboardInvitesDataObjectAsSharedDashboardInvitesData(v *SharedDashboardInvitesDataObject) SharedDashboardInvitesData {
	return SharedDashboardInvitesData{SharedDashboardInvitesDataObject: v}
}

// SharedDashboardInvitesDataListAsSharedDashboardInvitesData is a convenience function that returns []SharedDashboardInvitesDataObject wrapped in SharedDashboardInvitesData.
func SharedDashboardInvitesDataListAsSharedDashboardInvitesData(v *[]SharedDashboardInvitesDataObject) SharedDashboardInvitesData {
	return SharedDashboardInvitesData{SharedDashboardInvitesDataList: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *SharedDashboardInvitesData) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SharedDashboardInvitesDataObject
	err = json.Unmarshal(data, &obj.SharedDashboardInvitesDataObject)
	if err == nil {
		if obj.SharedDashboardInvitesDataObject != nil && obj.SharedDashboardInvitesDataObject.UnparsedObject == nil {
			jsonSharedDashboardInvitesDataObject, _ := json.Marshal(obj.SharedDashboardInvitesDataObject)
			if string(jsonSharedDashboardInvitesDataObject) == "{}" { // empty struct
				obj.SharedDashboardInvitesDataObject = nil
			} else {
				match++
			}
		} else {
			obj.SharedDashboardInvitesDataObject = nil
		}
	} else {
		obj.SharedDashboardInvitesDataObject = nil
	}

	// try to unmarshal data into SharedDashboardInvitesDataList
	err = json.Unmarshal(data, &obj.SharedDashboardInvitesDataList)
	if err == nil {
		if obj.SharedDashboardInvitesDataList != nil {
			jsonSharedDashboardInvitesDataList, _ := json.Marshal(obj.SharedDashboardInvitesDataList)
			if string(jsonSharedDashboardInvitesDataList) == "{}" { // empty struct
				obj.SharedDashboardInvitesDataList = nil
			} else {
				match++
			}
		} else {
			obj.SharedDashboardInvitesDataList = nil
		}
	} else {
		obj.SharedDashboardInvitesDataList = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.SharedDashboardInvitesDataObject = nil
		obj.SharedDashboardInvitesDataList = nil
		return json.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj SharedDashboardInvitesData) MarshalJSON() ([]byte, error) {
	if obj.SharedDashboardInvitesDataObject != nil {
		return json.Marshal(&obj.SharedDashboardInvitesDataObject)
	}

	if obj.SharedDashboardInvitesDataList != nil {
		return json.Marshal(&obj.SharedDashboardInvitesDataList)
	}

	if obj.UnparsedObject != nil {
		return json.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *SharedDashboardInvitesData) GetActualInstance() interface{} {
	if obj.SharedDashboardInvitesDataObject != nil {
		return obj.SharedDashboardInvitesDataObject
	}

	if obj.SharedDashboardInvitesDataList != nil {
		return obj.SharedDashboardInvitesDataList
	}

	// all schemas are nil
	return nil
}

// NullableSharedDashboardInvitesData handles when a null is used for SharedDashboardInvitesData.
type NullableSharedDashboardInvitesData struct {
	value *SharedDashboardInvitesData
	isSet bool
}

// Get returns the associated value.
func (v NullableSharedDashboardInvitesData) Get() *SharedDashboardInvitesData {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableSharedDashboardInvitesData) Set(val *SharedDashboardInvitesData) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableSharedDashboardInvitesData) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableSharedDashboardInvitesData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableSharedDashboardInvitesData initializes the struct as if Set has been called.
func NewNullableSharedDashboardInvitesData(val *SharedDashboardInvitesData) *NullableSharedDashboardInvitesData {
	return &NullableSharedDashboardInvitesData{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableSharedDashboardInvitesData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableSharedDashboardInvitesData) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return json.Unmarshal(src, &v.value)
}
