// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SharedDashboardIncluded - Resource included with a shared dashboard.
type SharedDashboardIncluded struct {
	SharedDashboardIncludedDashboard *SharedDashboardIncludedDashboard
	SharedDashboardIncludedUser      *SharedDashboardIncludedUser

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SharedDashboardIncludedDashboardAsSharedDashboardIncluded is a convenience function that returns SharedDashboardIncludedDashboard wrapped in SharedDashboardIncluded.
func SharedDashboardIncludedDashboardAsSharedDashboardIncluded(v *SharedDashboardIncludedDashboard) SharedDashboardIncluded {
	return SharedDashboardIncluded{SharedDashboardIncludedDashboard: v}
}

// SharedDashboardIncludedUserAsSharedDashboardIncluded is a convenience function that returns SharedDashboardIncludedUser wrapped in SharedDashboardIncluded.
func SharedDashboardIncludedUserAsSharedDashboardIncluded(v *SharedDashboardIncludedUser) SharedDashboardIncluded {
	return SharedDashboardIncluded{SharedDashboardIncludedUser: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *SharedDashboardIncluded) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SharedDashboardIncludedDashboard
	err = datadog.Unmarshal(data, &obj.SharedDashboardIncludedDashboard)
	if err == nil {
		if obj.SharedDashboardIncludedDashboard != nil && obj.SharedDashboardIncludedDashboard.UnparsedObject == nil {
			jsonSharedDashboardIncludedDashboard, _ := datadog.Marshal(obj.SharedDashboardIncludedDashboard)
			if string(jsonSharedDashboardIncludedDashboard) == "{}" { // empty struct
				obj.SharedDashboardIncludedDashboard = nil
			} else {
				match++
			}
		} else {
			obj.SharedDashboardIncludedDashboard = nil
		}
	} else {
		obj.SharedDashboardIncludedDashboard = nil
	}

	// try to unmarshal data into SharedDashboardIncludedUser
	err = datadog.Unmarshal(data, &obj.SharedDashboardIncludedUser)
	if err == nil {
		if obj.SharedDashboardIncludedUser != nil && obj.SharedDashboardIncludedUser.UnparsedObject == nil {
			jsonSharedDashboardIncludedUser, _ := datadog.Marshal(obj.SharedDashboardIncludedUser)
			if string(jsonSharedDashboardIncludedUser) == "{}" { // empty struct
				obj.SharedDashboardIncludedUser = nil
			} else {
				match++
			}
		} else {
			obj.SharedDashboardIncludedUser = nil
		}
	} else {
		obj.SharedDashboardIncludedUser = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.SharedDashboardIncludedDashboard = nil
		obj.SharedDashboardIncludedUser = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj SharedDashboardIncluded) MarshalJSON() ([]byte, error) {
	if obj.SharedDashboardIncludedDashboard != nil {
		return datadog.Marshal(&obj.SharedDashboardIncludedDashboard)
	}

	if obj.SharedDashboardIncludedUser != nil {
		return datadog.Marshal(&obj.SharedDashboardIncludedUser)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *SharedDashboardIncluded) GetActualInstance() interface{} {
	if obj.SharedDashboardIncludedDashboard != nil {
		return obj.SharedDashboardIncludedDashboard
	}

	if obj.SharedDashboardIncludedUser != nil {
		return obj.SharedDashboardIncludedUser
	}

	// all schemas are nil
	return nil
}
