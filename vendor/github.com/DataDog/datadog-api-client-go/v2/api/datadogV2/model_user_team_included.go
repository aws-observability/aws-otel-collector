// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UserTeamIncluded - Included resources related to the team membership
type UserTeamIncluded struct {
	User *User
	Team *Team

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// UserAsUserTeamIncluded is a convenience function that returns User wrapped in UserTeamIncluded.
func UserAsUserTeamIncluded(v *User) UserTeamIncluded {
	return UserTeamIncluded{User: v}
}

// TeamAsUserTeamIncluded is a convenience function that returns Team wrapped in UserTeamIncluded.
func TeamAsUserTeamIncluded(v *Team) UserTeamIncluded {
	return UserTeamIncluded{Team: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *UserTeamIncluded) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into Team
	err = datadog.Unmarshal(data, &obj.Team)
	if err == nil {
		if obj.Team != nil && obj.Team.UnparsedObject == nil {
			jsonTeam, _ := datadog.Marshal(obj.Team)
			if string(jsonTeam) == "{}" { // empty struct
				obj.Team = nil
			} else {
				match++
			}
		} else {
			obj.Team = nil
		}
	} else {
		obj.Team = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.User = nil
		obj.Team = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj UserTeamIncluded) MarshalJSON() ([]byte, error) {
	if obj.User != nil {
		return datadog.Marshal(&obj.User)
	}

	if obj.Team != nil {
		return datadog.Marshal(&obj.Team)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *UserTeamIncluded) GetActualInstance() interface{} {
	if obj.User != nil {
		return obj.User
	}

	if obj.Team != nil {
		return obj.Team
	}

	// all schemas are nil
	return nil
}
