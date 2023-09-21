// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
<<<<<<< HEAD
	"encoding/json"
=======
	"github.com/goccy/go-json"
>>>>>>> main
)

// TeamIncluded - Included resources related to the team
type TeamIncluded struct {
	User               *User
	TeamLink           *TeamLink
	UserTeamPermission *UserTeamPermission

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// UserAsTeamIncluded is a convenience function that returns User wrapped in TeamIncluded.
func UserAsTeamIncluded(v *User) TeamIncluded {
	return TeamIncluded{User: v}
}

// TeamLinkAsTeamIncluded is a convenience function that returns TeamLink wrapped in TeamIncluded.
func TeamLinkAsTeamIncluded(v *TeamLink) TeamIncluded {
	return TeamIncluded{TeamLink: v}
}

// UserTeamPermissionAsTeamIncluded is a convenience function that returns UserTeamPermission wrapped in TeamIncluded.
func UserTeamPermissionAsTeamIncluded(v *UserTeamPermission) TeamIncluded {
	return TeamIncluded{UserTeamPermission: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *TeamIncluded) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into User
	err = json.Unmarshal(data, &obj.User)
	if err == nil {
		if obj.User != nil && obj.User.UnparsedObject == nil {
			jsonUser, _ := json.Marshal(obj.User)
			if string(jsonUser) == "{}" { // empty struct
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

	// try to unmarshal data into TeamLink
	err = json.Unmarshal(data, &obj.TeamLink)
	if err == nil {
		if obj.TeamLink != nil && obj.TeamLink.UnparsedObject == nil {
			jsonTeamLink, _ := json.Marshal(obj.TeamLink)
			if string(jsonTeamLink) == "{}" { // empty struct
				obj.TeamLink = nil
			} else {
				match++
			}
		} else {
			obj.TeamLink = nil
		}
	} else {
		obj.TeamLink = nil
	}

	// try to unmarshal data into UserTeamPermission
	err = json.Unmarshal(data, &obj.UserTeamPermission)
	if err == nil {
		if obj.UserTeamPermission != nil && obj.UserTeamPermission.UnparsedObject == nil {
			jsonUserTeamPermission, _ := json.Marshal(obj.UserTeamPermission)
			if string(jsonUserTeamPermission) == "{}" { // empty struct
				obj.UserTeamPermission = nil
			} else {
				match++
			}
		} else {
			obj.UserTeamPermission = nil
		}
	} else {
		obj.UserTeamPermission = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.User = nil
		obj.TeamLink = nil
		obj.UserTeamPermission = nil
		return json.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj TeamIncluded) MarshalJSON() ([]byte, error) {
	if obj.User != nil {
		return json.Marshal(&obj.User)
	}

	if obj.TeamLink != nil {
		return json.Marshal(&obj.TeamLink)
	}

	if obj.UserTeamPermission != nil {
		return json.Marshal(&obj.UserTeamPermission)
	}

	if obj.UnparsedObject != nil {
		return json.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *TeamIncluded) GetActualInstance() interface{} {
	if obj.User != nil {
		return obj.User
	}

	if obj.TeamLink != nil {
		return obj.TeamLink
	}

	if obj.UserTeamPermission != nil {
		return obj.UserTeamPermission
	}

	// all schemas are nil
	return nil
}
<<<<<<< HEAD

// NullableTeamIncluded handles when a null is used for TeamIncluded.
type NullableTeamIncluded struct {
	value *TeamIncluded
	isSet bool
}

// Get returns the associated value.
func (v NullableTeamIncluded) Get() *TeamIncluded {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableTeamIncluded) Set(val *TeamIncluded) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableTeamIncluded) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableTeamIncluded) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTeamIncluded initializes the struct as if Set has been called.
func NewNullableTeamIncluded(val *TeamIncluded) *NullableTeamIncluded {
	return &NullableTeamIncluded{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableTeamIncluded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableTeamIncluded) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return json.Unmarshal(src, &v.value)
}
=======
>>>>>>> main
