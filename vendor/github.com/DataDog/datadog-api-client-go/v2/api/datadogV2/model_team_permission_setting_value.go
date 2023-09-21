// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
<<<<<<< HEAD
	"encoding/json"
	"fmt"
=======
	"fmt"

	"github.com/goccy/go-json"
>>>>>>> main
)

// TeamPermissionSettingValue What type of user is allowed to perform the specified action
type TeamPermissionSettingValue string

// List of TeamPermissionSettingValue.
const (
	TEAMPERMISSIONSETTINGVALUE_ADMINS             TeamPermissionSettingValue = "admins"
	TEAMPERMISSIONSETTINGVALUE_MEMBERS            TeamPermissionSettingValue = "members"
	TEAMPERMISSIONSETTINGVALUE_ORGANIZATION       TeamPermissionSettingValue = "organization"
	TEAMPERMISSIONSETTINGVALUE_USER_ACCESS_MANAGE TeamPermissionSettingValue = "user_access_manage"
	TEAMPERMISSIONSETTINGVALUE_TEAMS_MANAGE       TeamPermissionSettingValue = "teams_manage"
)

var allowedTeamPermissionSettingValueEnumValues = []TeamPermissionSettingValue{
	TEAMPERMISSIONSETTINGVALUE_ADMINS,
	TEAMPERMISSIONSETTINGVALUE_MEMBERS,
	TEAMPERMISSIONSETTINGVALUE_ORGANIZATION,
	TEAMPERMISSIONSETTINGVALUE_USER_ACCESS_MANAGE,
	TEAMPERMISSIONSETTINGVALUE_TEAMS_MANAGE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TeamPermissionSettingValue) GetAllowedValues() []TeamPermissionSettingValue {
	return allowedTeamPermissionSettingValueEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TeamPermissionSettingValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TeamPermissionSettingValue(value)
	return nil
}

// NewTeamPermissionSettingValueFromValue returns a pointer to a valid TeamPermissionSettingValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTeamPermissionSettingValueFromValue(v string) (*TeamPermissionSettingValue, error) {
	ev := TeamPermissionSettingValue(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TeamPermissionSettingValue: valid values are %v", v, allowedTeamPermissionSettingValueEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TeamPermissionSettingValue) IsValid() bool {
	for _, existing := range allowedTeamPermissionSettingValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TeamPermissionSettingValue value.
func (v TeamPermissionSettingValue) Ptr() *TeamPermissionSettingValue {
	return &v
}
<<<<<<< HEAD

// NullableTeamPermissionSettingValue handles when a null is used for TeamPermissionSettingValue.
type NullableTeamPermissionSettingValue struct {
	value *TeamPermissionSettingValue
	isSet bool
}

// Get returns the associated value.
func (v NullableTeamPermissionSettingValue) Get() *TeamPermissionSettingValue {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableTeamPermissionSettingValue) Set(val *TeamPermissionSettingValue) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableTeamPermissionSettingValue) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableTeamPermissionSettingValue) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTeamPermissionSettingValue initializes the struct as if Set has been called.
func NewNullableTeamPermissionSettingValue(val *TeamPermissionSettingValue) *NullableTeamPermissionSettingValue {
	return &NullableTeamPermissionSettingValue{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableTeamPermissionSettingValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableTeamPermissionSettingValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
=======
>>>>>>> main
