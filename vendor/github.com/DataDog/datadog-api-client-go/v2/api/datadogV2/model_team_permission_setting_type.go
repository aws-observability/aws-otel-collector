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

// TeamPermissionSettingType Team permission setting type
type TeamPermissionSettingType string

// List of TeamPermissionSettingType.
const (
	TEAMPERMISSIONSETTINGTYPE_TEAM_PERMISSION_SETTINGS TeamPermissionSettingType = "team_permission_settings"
)

var allowedTeamPermissionSettingTypeEnumValues = []TeamPermissionSettingType{
	TEAMPERMISSIONSETTINGTYPE_TEAM_PERMISSION_SETTINGS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TeamPermissionSettingType) GetAllowedValues() []TeamPermissionSettingType {
	return allowedTeamPermissionSettingTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TeamPermissionSettingType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TeamPermissionSettingType(value)
	return nil
}

// NewTeamPermissionSettingTypeFromValue returns a pointer to a valid TeamPermissionSettingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTeamPermissionSettingTypeFromValue(v string) (*TeamPermissionSettingType, error) {
	ev := TeamPermissionSettingType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TeamPermissionSettingType: valid values are %v", v, allowedTeamPermissionSettingTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TeamPermissionSettingType) IsValid() bool {
	for _, existing := range allowedTeamPermissionSettingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TeamPermissionSettingType value.
func (v TeamPermissionSettingType) Ptr() *TeamPermissionSettingType {
	return &v
}
<<<<<<< HEAD

// NullableTeamPermissionSettingType handles when a null is used for TeamPermissionSettingType.
type NullableTeamPermissionSettingType struct {
	value *TeamPermissionSettingType
	isSet bool
}

// Get returns the associated value.
func (v NullableTeamPermissionSettingType) Get() *TeamPermissionSettingType {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableTeamPermissionSettingType) Set(val *TeamPermissionSettingType) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableTeamPermissionSettingType) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableTeamPermissionSettingType) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTeamPermissionSettingType initializes the struct as if Set has been called.
func NewNullableTeamPermissionSettingType(val *TeamPermissionSettingType) *NullableTeamPermissionSettingType {
	return &NullableTeamPermissionSettingType{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableTeamPermissionSettingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableTeamPermissionSettingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
=======
>>>>>>> main
