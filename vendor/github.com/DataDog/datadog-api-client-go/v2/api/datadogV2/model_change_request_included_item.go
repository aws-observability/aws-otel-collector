// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeRequestIncludedItem - An included resource item in the change request response.
type ChangeRequestIncludedItem struct {
	ChangeRequestIncludedUser     *ChangeRequestIncludedUser
	ChangeRequestIncludedDecision *ChangeRequestIncludedDecision

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ChangeRequestIncludedUserAsChangeRequestIncludedItem is a convenience function that returns ChangeRequestIncludedUser wrapped in ChangeRequestIncludedItem.
func ChangeRequestIncludedUserAsChangeRequestIncludedItem(v *ChangeRequestIncludedUser) ChangeRequestIncludedItem {
	return ChangeRequestIncludedItem{ChangeRequestIncludedUser: v}
}

// ChangeRequestIncludedDecisionAsChangeRequestIncludedItem is a convenience function that returns ChangeRequestIncludedDecision wrapped in ChangeRequestIncludedItem.
func ChangeRequestIncludedDecisionAsChangeRequestIncludedItem(v *ChangeRequestIncludedDecision) ChangeRequestIncludedItem {
	return ChangeRequestIncludedItem{ChangeRequestIncludedDecision: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ChangeRequestIncludedItem) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ChangeRequestIncludedUser
	err = datadog.Unmarshal(data, &obj.ChangeRequestIncludedUser)
	if err == nil {
		if obj.ChangeRequestIncludedUser != nil && obj.ChangeRequestIncludedUser.UnparsedObject == nil {
			jsonChangeRequestIncludedUser, _ := datadog.Marshal(obj.ChangeRequestIncludedUser)
			if string(jsonChangeRequestIncludedUser) == "{}" { // empty struct
				obj.ChangeRequestIncludedUser = nil
			} else {
				match++
			}
		} else {
			obj.ChangeRequestIncludedUser = nil
		}
	} else {
		obj.ChangeRequestIncludedUser = nil
	}

	// try to unmarshal data into ChangeRequestIncludedDecision
	err = datadog.Unmarshal(data, &obj.ChangeRequestIncludedDecision)
	if err == nil {
		if obj.ChangeRequestIncludedDecision != nil && obj.ChangeRequestIncludedDecision.UnparsedObject == nil {
			jsonChangeRequestIncludedDecision, _ := datadog.Marshal(obj.ChangeRequestIncludedDecision)
			if string(jsonChangeRequestIncludedDecision) == "{}" { // empty struct
				obj.ChangeRequestIncludedDecision = nil
			} else {
				match++
			}
		} else {
			obj.ChangeRequestIncludedDecision = nil
		}
	} else {
		obj.ChangeRequestIncludedDecision = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ChangeRequestIncludedUser = nil
		obj.ChangeRequestIncludedDecision = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ChangeRequestIncludedItem) MarshalJSON() ([]byte, error) {
	if obj.ChangeRequestIncludedUser != nil {
		return datadog.Marshal(&obj.ChangeRequestIncludedUser)
	}

	if obj.ChangeRequestIncludedDecision != nil {
		return datadog.Marshal(&obj.ChangeRequestIncludedDecision)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ChangeRequestIncludedItem) GetActualInstance() interface{} {
	if obj.ChangeRequestIncludedUser != nil {
		return obj.ChangeRequestIncludedUser
	}

	if obj.ChangeRequestIncludedDecision != nil {
		return obj.ChangeRequestIncludedDecision
	}

	// all schemas are nil
	return nil
}
