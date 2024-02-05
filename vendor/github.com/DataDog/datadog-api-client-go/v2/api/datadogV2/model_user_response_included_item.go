// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UserResponseIncludedItem - An object related to a user.
type UserResponseIncludedItem struct {
	Organization *Organization
	Permission   *Permission
	Role         *Role

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// OrganizationAsUserResponseIncludedItem is a convenience function that returns Organization wrapped in UserResponseIncludedItem.
func OrganizationAsUserResponseIncludedItem(v *Organization) UserResponseIncludedItem {
	return UserResponseIncludedItem{Organization: v}
}

// PermissionAsUserResponseIncludedItem is a convenience function that returns Permission wrapped in UserResponseIncludedItem.
func PermissionAsUserResponseIncludedItem(v *Permission) UserResponseIncludedItem {
	return UserResponseIncludedItem{Permission: v}
}

// RoleAsUserResponseIncludedItem is a convenience function that returns Role wrapped in UserResponseIncludedItem.
func RoleAsUserResponseIncludedItem(v *Role) UserResponseIncludedItem {
	return UserResponseIncludedItem{Role: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *UserResponseIncludedItem) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Organization
	err = datadog.Unmarshal(data, &obj.Organization)
	if err == nil {
		if obj.Organization != nil && obj.Organization.UnparsedObject == nil {
			jsonOrganization, _ := datadog.Marshal(obj.Organization)
			if string(jsonOrganization) == "{}" { // empty struct
				obj.Organization = nil
			} else {
				match++
			}
		} else {
			obj.Organization = nil
		}
	} else {
		obj.Organization = nil
	}

	// try to unmarshal data into Permission
	err = datadog.Unmarshal(data, &obj.Permission)
	if err == nil {
		if obj.Permission != nil && obj.Permission.UnparsedObject == nil {
			jsonPermission, _ := datadog.Marshal(obj.Permission)
			if string(jsonPermission) == "{}" { // empty struct
				obj.Permission = nil
			} else {
				match++
			}
		} else {
			obj.Permission = nil
		}
	} else {
		obj.Permission = nil
	}

	// try to unmarshal data into Role
	err = datadog.Unmarshal(data, &obj.Role)
	if err == nil {
		if obj.Role != nil && obj.Role.UnparsedObject == nil {
			jsonRole, _ := datadog.Marshal(obj.Role)
			if string(jsonRole) == "{}" { // empty struct
				obj.Role = nil
			} else {
				match++
			}
		} else {
			obj.Role = nil
		}
	} else {
		obj.Role = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.Organization = nil
		obj.Permission = nil
		obj.Role = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj UserResponseIncludedItem) MarshalJSON() ([]byte, error) {
	if obj.Organization != nil {
		return datadog.Marshal(&obj.Organization)
	}

	if obj.Permission != nil {
		return datadog.Marshal(&obj.Permission)
	}

	if obj.Role != nil {
		return datadog.Marshal(&obj.Role)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *UserResponseIncludedItem) GetActualInstance() interface{} {
	if obj.Organization != nil {
		return obj.Organization
	}

	if obj.Permission != nil {
		return obj.Permission
	}

	if obj.Role != nil {
		return obj.Role
	}

	// all schemas are nil
	return nil
}
