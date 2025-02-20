// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AuthNMappingIncluded - Included data in the AuthN Mapping response.
type AuthNMappingIncluded struct {
	SAMLAssertionAttribute *SAMLAssertionAttribute
	Role                   *Role
	AuthNMappingTeam       *AuthNMappingTeam

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SAMLAssertionAttributeAsAuthNMappingIncluded is a convenience function that returns SAMLAssertionAttribute wrapped in AuthNMappingIncluded.
func SAMLAssertionAttributeAsAuthNMappingIncluded(v *SAMLAssertionAttribute) AuthNMappingIncluded {
	return AuthNMappingIncluded{SAMLAssertionAttribute: v}
}

// RoleAsAuthNMappingIncluded is a convenience function that returns Role wrapped in AuthNMappingIncluded.
func RoleAsAuthNMappingIncluded(v *Role) AuthNMappingIncluded {
	return AuthNMappingIncluded{Role: v}
}

// AuthNMappingTeamAsAuthNMappingIncluded is a convenience function that returns AuthNMappingTeam wrapped in AuthNMappingIncluded.
func AuthNMappingTeamAsAuthNMappingIncluded(v *AuthNMappingTeam) AuthNMappingIncluded {
	return AuthNMappingIncluded{AuthNMappingTeam: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *AuthNMappingIncluded) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SAMLAssertionAttribute
	err = datadog.Unmarshal(data, &obj.SAMLAssertionAttribute)
	if err == nil {
		if obj.SAMLAssertionAttribute != nil && obj.SAMLAssertionAttribute.UnparsedObject == nil {
			jsonSAMLAssertionAttribute, _ := datadog.Marshal(obj.SAMLAssertionAttribute)
			if string(jsonSAMLAssertionAttribute) == "{}" { // empty struct
				obj.SAMLAssertionAttribute = nil
			} else {
				match++
			}
		} else {
			obj.SAMLAssertionAttribute = nil
		}
	} else {
		obj.SAMLAssertionAttribute = nil
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

	// try to unmarshal data into AuthNMappingTeam
	err = datadog.Unmarshal(data, &obj.AuthNMappingTeam)
	if err == nil {
		if obj.AuthNMappingTeam != nil && obj.AuthNMappingTeam.UnparsedObject == nil {
			jsonAuthNMappingTeam, _ := datadog.Marshal(obj.AuthNMappingTeam)
			if string(jsonAuthNMappingTeam) == "{}" && string(data) != "{}" { // empty struct
				obj.AuthNMappingTeam = nil
			} else {
				match++
			}
		} else {
			obj.AuthNMappingTeam = nil
		}
	} else {
		obj.AuthNMappingTeam = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.SAMLAssertionAttribute = nil
		obj.Role = nil
		obj.AuthNMappingTeam = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj AuthNMappingIncluded) MarshalJSON() ([]byte, error) {
	if obj.SAMLAssertionAttribute != nil {
		return datadog.Marshal(&obj.SAMLAssertionAttribute)
	}

	if obj.Role != nil {
		return datadog.Marshal(&obj.Role)
	}

	if obj.AuthNMappingTeam != nil {
		return datadog.Marshal(&obj.AuthNMappingTeam)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *AuthNMappingIncluded) GetActualInstance() interface{} {
	if obj.SAMLAssertionAttribute != nil {
		return obj.SAMLAssertionAttribute
	}

	if obj.Role != nil {
		return obj.Role
	}

	if obj.AuthNMappingTeam != nil {
		return obj.AuthNMappingTeam
	}

	// all schemas are nil
	return nil
}
