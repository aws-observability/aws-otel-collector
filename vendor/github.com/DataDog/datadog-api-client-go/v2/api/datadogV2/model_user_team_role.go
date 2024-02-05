// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UserTeamRole The user's role within the team
type UserTeamRole string

// List of UserTeamRole.
const (
	USERTEAMROLE_ADMIN UserTeamRole = "admin"
)

var allowedUserTeamRoleEnumValues = []UserTeamRole{
	USERTEAMROLE_ADMIN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UserTeamRole) GetAllowedValues() []UserTeamRole {
	return allowedUserTeamRoleEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UserTeamRole) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UserTeamRole(value)
	return nil
}

// NewUserTeamRoleFromValue returns a pointer to a valid UserTeamRole
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUserTeamRoleFromValue(v string) (*UserTeamRole, error) {
	ev := UserTeamRole(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UserTeamRole: valid values are %v", v, allowedUserTeamRoleEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UserTeamRole) IsValid() bool {
	for _, existing := range allowedUserTeamRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserTeamRole value.
func (v UserTeamRole) Ptr() *UserTeamRole {
	return &v
}

// NullableUserTeamRole handles when a null is used for UserTeamRole.
type NullableUserTeamRole struct {
	value *UserTeamRole
	isSet bool
}

// Get returns the associated value.
func (v NullableUserTeamRole) Get() *UserTeamRole {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableUserTeamRole) Set(val *UserTeamRole) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableUserTeamRole) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableUserTeamRole) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableUserTeamRole initializes the struct as if Set has been called.
func NewNullableUserTeamRole(val *UserTeamRole) *NullableUserTeamRole {
	return &NullableUserTeamRole{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableUserTeamRole) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableUserTeamRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return datadog.Unmarshal(src, &v.value)
}
