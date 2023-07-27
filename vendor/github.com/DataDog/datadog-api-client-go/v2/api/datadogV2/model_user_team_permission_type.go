// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// UserTeamPermissionType User team permission type
type UserTeamPermissionType string

// List of UserTeamPermissionType.
const (
	USERTEAMPERMISSIONTYPE_USER_TEAM_PERMISSIONS UserTeamPermissionType = "user_team_permissions"
)

var allowedUserTeamPermissionTypeEnumValues = []UserTeamPermissionType{
	USERTEAMPERMISSIONTYPE_USER_TEAM_PERMISSIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UserTeamPermissionType) GetAllowedValues() []UserTeamPermissionType {
	return allowedUserTeamPermissionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UserTeamPermissionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UserTeamPermissionType(value)
	return nil
}

// NewUserTeamPermissionTypeFromValue returns a pointer to a valid UserTeamPermissionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUserTeamPermissionTypeFromValue(v string) (*UserTeamPermissionType, error) {
	ev := UserTeamPermissionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UserTeamPermissionType: valid values are %v", v, allowedUserTeamPermissionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UserTeamPermissionType) IsValid() bool {
	for _, existing := range allowedUserTeamPermissionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserTeamPermissionType value.
func (v UserTeamPermissionType) Ptr() *UserTeamPermissionType {
	return &v
}

// NullableUserTeamPermissionType handles when a null is used for UserTeamPermissionType.
type NullableUserTeamPermissionType struct {
	value *UserTeamPermissionType
	isSet bool
}

// Get returns the associated value.
func (v NullableUserTeamPermissionType) Get() *UserTeamPermissionType {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableUserTeamPermissionType) Set(val *UserTeamPermissionType) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableUserTeamPermissionType) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableUserTeamPermissionType) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableUserTeamPermissionType initializes the struct as if Set has been called.
func NewNullableUserTeamPermissionType(val *UserTeamPermissionType) *NullableUserTeamPermissionType {
	return &NullableUserTeamPermissionType{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableUserTeamPermissionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableUserTeamPermissionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
