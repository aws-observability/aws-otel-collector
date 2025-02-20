// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// APIKeyResponseIncludedItem - An object related to an API key.
type APIKeyResponseIncludedItem struct {
	User      *User
	LeakedKey *LeakedKey

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// UserAsAPIKeyResponseIncludedItem is a convenience function that returns User wrapped in APIKeyResponseIncludedItem.
func UserAsAPIKeyResponseIncludedItem(v *User) APIKeyResponseIncludedItem {
	return APIKeyResponseIncludedItem{User: v}
}

// LeakedKeyAsAPIKeyResponseIncludedItem is a convenience function that returns LeakedKey wrapped in APIKeyResponseIncludedItem.
func LeakedKeyAsAPIKeyResponseIncludedItem(v *LeakedKey) APIKeyResponseIncludedItem {
	return APIKeyResponseIncludedItem{LeakedKey: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *APIKeyResponseIncludedItem) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into LeakedKey
	err = datadog.Unmarshal(data, &obj.LeakedKey)
	if err == nil {
		if obj.LeakedKey != nil && obj.LeakedKey.UnparsedObject == nil {
			jsonLeakedKey, _ := datadog.Marshal(obj.LeakedKey)
			if string(jsonLeakedKey) == "{}" { // empty struct
				obj.LeakedKey = nil
			} else {
				match++
			}
		} else {
			obj.LeakedKey = nil
		}
	} else {
		obj.LeakedKey = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.User = nil
		obj.LeakedKey = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj APIKeyResponseIncludedItem) MarshalJSON() ([]byte, error) {
	if obj.User != nil {
		return datadog.Marshal(&obj.User)
	}

	if obj.LeakedKey != nil {
		return datadog.Marshal(&obj.LeakedKey)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *APIKeyResponseIncludedItem) GetActualInstance() interface{} {
	if obj.User != nil {
		return obj.User
	}

	if obj.LeakedKey != nil {
		return obj.LeakedKey
	}

	// all schemas are nil
	return nil
}
