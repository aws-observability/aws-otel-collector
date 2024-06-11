// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AuthNMappingUpdateRelationships - Relationship of AuthN Mapping update object to a Role or Team.
type AuthNMappingUpdateRelationships struct {
	AuthNMappingRelationshipToRole *AuthNMappingRelationshipToRole
	AuthNMappingRelationshipToTeam *AuthNMappingRelationshipToTeam

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// AuthNMappingRelationshipToRoleAsAuthNMappingUpdateRelationships is a convenience function that returns AuthNMappingRelationshipToRole wrapped in AuthNMappingUpdateRelationships.
func AuthNMappingRelationshipToRoleAsAuthNMappingUpdateRelationships(v *AuthNMappingRelationshipToRole) AuthNMappingUpdateRelationships {
	return AuthNMappingUpdateRelationships{AuthNMappingRelationshipToRole: v}
}

// AuthNMappingRelationshipToTeamAsAuthNMappingUpdateRelationships is a convenience function that returns AuthNMappingRelationshipToTeam wrapped in AuthNMappingUpdateRelationships.
func AuthNMappingRelationshipToTeamAsAuthNMappingUpdateRelationships(v *AuthNMappingRelationshipToTeam) AuthNMappingUpdateRelationships {
	return AuthNMappingUpdateRelationships{AuthNMappingRelationshipToTeam: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *AuthNMappingUpdateRelationships) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AuthNMappingRelationshipToRole
	err = datadog.Unmarshal(data, &obj.AuthNMappingRelationshipToRole)
	if err == nil {
		if obj.AuthNMappingRelationshipToRole != nil && obj.AuthNMappingRelationshipToRole.UnparsedObject == nil {
			jsonAuthNMappingRelationshipToRole, _ := datadog.Marshal(obj.AuthNMappingRelationshipToRole)
			if string(jsonAuthNMappingRelationshipToRole) == "{}" { // empty struct
				obj.AuthNMappingRelationshipToRole = nil
			} else {
				match++
			}
		} else {
			obj.AuthNMappingRelationshipToRole = nil
		}
	} else {
		obj.AuthNMappingRelationshipToRole = nil
	}

	// try to unmarshal data into AuthNMappingRelationshipToTeam
	err = datadog.Unmarshal(data, &obj.AuthNMappingRelationshipToTeam)
	if err == nil {
		if obj.AuthNMappingRelationshipToTeam != nil && obj.AuthNMappingRelationshipToTeam.UnparsedObject == nil {
			jsonAuthNMappingRelationshipToTeam, _ := datadog.Marshal(obj.AuthNMappingRelationshipToTeam)
			if string(jsonAuthNMappingRelationshipToTeam) == "{}" { // empty struct
				obj.AuthNMappingRelationshipToTeam = nil
			} else {
				match++
			}
		} else {
			obj.AuthNMappingRelationshipToTeam = nil
		}
	} else {
		obj.AuthNMappingRelationshipToTeam = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.AuthNMappingRelationshipToRole = nil
		obj.AuthNMappingRelationshipToTeam = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj AuthNMappingUpdateRelationships) MarshalJSON() ([]byte, error) {
	if obj.AuthNMappingRelationshipToRole != nil {
		return datadog.Marshal(&obj.AuthNMappingRelationshipToRole)
	}

	if obj.AuthNMappingRelationshipToTeam != nil {
		return datadog.Marshal(&obj.AuthNMappingRelationshipToTeam)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *AuthNMappingUpdateRelationships) GetActualInstance() interface{} {
	if obj.AuthNMappingRelationshipToRole != nil {
		return obj.AuthNMappingRelationshipToRole
	}

	if obj.AuthNMappingRelationshipToTeam != nil {
		return obj.AuthNMappingRelationshipToTeam
	}

	// all schemas are nil
	return nil
}
