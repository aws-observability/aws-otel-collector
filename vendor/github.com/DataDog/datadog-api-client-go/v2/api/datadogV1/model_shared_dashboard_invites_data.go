// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
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
	err = datadog.Unmarshal(data, &obj.SharedDashboardInvitesDataObject)
	if err == nil {
		if obj.SharedDashboardInvitesDataObject != nil && obj.SharedDashboardInvitesDataObject.UnparsedObject == nil {
			jsonSharedDashboardInvitesDataObject, _ := datadog.Marshal(obj.SharedDashboardInvitesDataObject)
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
	err = datadog.Unmarshal(data, &obj.SharedDashboardInvitesDataList)
	if err == nil {
		if obj.SharedDashboardInvitesDataList != nil {
			jsonSharedDashboardInvitesDataList, _ := datadog.Marshal(obj.SharedDashboardInvitesDataList)
			if string(jsonSharedDashboardInvitesDataList) == "{}" && string(data) != "{}" { // empty struct
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
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj SharedDashboardInvitesData) MarshalJSON() ([]byte, error) {
	if obj.SharedDashboardInvitesDataObject != nil {
		return datadog.Marshal(&obj.SharedDashboardInvitesDataObject)
	}

	if obj.SharedDashboardInvitesDataList != nil {
		return datadog.Marshal(&obj.SharedDashboardInvitesDataList)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
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
