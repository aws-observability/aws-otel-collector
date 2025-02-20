// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DowntimeResponseIncludedItem - An object related to a downtime.
type DowntimeResponseIncludedItem struct {
	User                        *User
	DowntimeMonitorIncludedItem *DowntimeMonitorIncludedItem

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// UserAsDowntimeResponseIncludedItem is a convenience function that returns User wrapped in DowntimeResponseIncludedItem.
func UserAsDowntimeResponseIncludedItem(v *User) DowntimeResponseIncludedItem {
	return DowntimeResponseIncludedItem{User: v}
}

// DowntimeMonitorIncludedItemAsDowntimeResponseIncludedItem is a convenience function that returns DowntimeMonitorIncludedItem wrapped in DowntimeResponseIncludedItem.
func DowntimeMonitorIncludedItemAsDowntimeResponseIncludedItem(v *DowntimeMonitorIncludedItem) DowntimeResponseIncludedItem {
	return DowntimeResponseIncludedItem{DowntimeMonitorIncludedItem: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *DowntimeResponseIncludedItem) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into User
	err = datadog.Unmarshal(data, &obj.User)
	if err == nil {
		if obj.User != nil && obj.User.UnparsedObject == nil {
			jsonUser, _ := datadog.Marshal(obj.User)
			if string(jsonUser) == "{}" && string(data) != "{}" { // empty struct
				obj.User = nil
			} else {
				match++
			}
		} else {
			obj.User = nil
		}
	} else {
		obj.User = nil
	}

	// try to unmarshal data into DowntimeMonitorIncludedItem
	err = datadog.Unmarshal(data, &obj.DowntimeMonitorIncludedItem)
	if err == nil {
		if obj.DowntimeMonitorIncludedItem != nil && obj.DowntimeMonitorIncludedItem.UnparsedObject == nil {
			jsonDowntimeMonitorIncludedItem, _ := datadog.Marshal(obj.DowntimeMonitorIncludedItem)
			if string(jsonDowntimeMonitorIncludedItem) == "{}" && string(data) != "{}" { // empty struct
				obj.DowntimeMonitorIncludedItem = nil
			} else {
				match++
			}
		} else {
			obj.DowntimeMonitorIncludedItem = nil
		}
	} else {
		obj.DowntimeMonitorIncludedItem = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.User = nil
		obj.DowntimeMonitorIncludedItem = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj DowntimeResponseIncludedItem) MarshalJSON() ([]byte, error) {
	if obj.User != nil {
		return datadog.Marshal(&obj.User)
	}

	if obj.DowntimeMonitorIncludedItem != nil {
		return datadog.Marshal(&obj.DowntimeMonitorIncludedItem)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *DowntimeResponseIncludedItem) GetActualInstance() interface{} {
	if obj.User != nil {
		return obj.User
	}

	if obj.DowntimeMonitorIncludedItem != nil {
		return obj.DowntimeMonitorIncludedItem
	}

	// all schemas are nil
	return nil
}
