// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
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

	// try to unmarshal data into TeamLink
	err = datadog.Unmarshal(data, &obj.TeamLink)
	if err == nil {
		if obj.TeamLink != nil && obj.TeamLink.UnparsedObject == nil {
			jsonTeamLink, _ := datadog.Marshal(obj.TeamLink)
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
	err = datadog.Unmarshal(data, &obj.UserTeamPermission)
	if err == nil {
		if obj.UserTeamPermission != nil && obj.UserTeamPermission.UnparsedObject == nil {
			jsonUserTeamPermission, _ := datadog.Marshal(obj.UserTeamPermission)
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
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj TeamIncluded) MarshalJSON() ([]byte, error) {
	if obj.User != nil {
		return datadog.Marshal(&obj.User)
	}

	if obj.TeamLink != nil {
		return datadog.Marshal(&obj.TeamLink)
	}

	if obj.UserTeamPermission != nil {
		return datadog.Marshal(&obj.UserTeamPermission)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
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
