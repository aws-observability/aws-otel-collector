// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentResponseIncludedItem - An object related to an incident that is included in the response.
type IncidentResponseIncludedItem struct {
	IncidentUserData *IncidentUserData
	AttachmentData   *AttachmentData

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// IncidentUserDataAsIncidentResponseIncludedItem is a convenience function that returns IncidentUserData wrapped in IncidentResponseIncludedItem.
func IncidentUserDataAsIncidentResponseIncludedItem(v *IncidentUserData) IncidentResponseIncludedItem {
	return IncidentResponseIncludedItem{IncidentUserData: v}
}

// AttachmentDataAsIncidentResponseIncludedItem is a convenience function that returns AttachmentData wrapped in IncidentResponseIncludedItem.
func AttachmentDataAsIncidentResponseIncludedItem(v *AttachmentData) IncidentResponseIncludedItem {
	return IncidentResponseIncludedItem{AttachmentData: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *IncidentResponseIncludedItem) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into AttachmentData
	err = datadog.Unmarshal(data, &obj.AttachmentData)
	if err == nil {
		if obj.AttachmentData != nil && obj.AttachmentData.UnparsedObject == nil {
			jsonAttachmentData, _ := datadog.Marshal(obj.AttachmentData)
			if string(jsonAttachmentData) == "{}" { // empty struct
				obj.AttachmentData = nil
			} else {
				match++
			}
		} else {
			obj.AttachmentData = nil
		}
	} else {
		obj.AttachmentData = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.IncidentUserData = nil
		obj.AttachmentData = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj IncidentResponseIncludedItem) MarshalJSON() ([]byte, error) {
	if obj.IncidentUserData != nil {
		return datadog.Marshal(&obj.IncidentUserData)
	}

	if obj.AttachmentData != nil {
		return datadog.Marshal(&obj.AttachmentData)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *IncidentResponseIncludedItem) GetActualInstance() interface{} {
	if obj.IncidentUserData != nil {
		return obj.IncidentUserData
	}

	if obj.AttachmentData != nil {
		return obj.AttachmentData
	}

	// all schemas are nil
	return nil
}
