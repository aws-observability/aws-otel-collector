// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// TeamLinkType Team link type
type TeamLinkType string

// List of TeamLinkType.
const (
	TEAMLINKTYPE_TEAM_LINKS TeamLinkType = "team_links"
)

var allowedTeamLinkTypeEnumValues = []TeamLinkType{
	TEAMLINKTYPE_TEAM_LINKS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TeamLinkType) GetAllowedValues() []TeamLinkType {
	return allowedTeamLinkTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TeamLinkType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TeamLinkType(value)
	return nil
}

// NewTeamLinkTypeFromValue returns a pointer to a valid TeamLinkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTeamLinkTypeFromValue(v string) (*TeamLinkType, error) {
	ev := TeamLinkType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TeamLinkType: valid values are %v", v, allowedTeamLinkTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TeamLinkType) IsValid() bool {
	for _, existing := range allowedTeamLinkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TeamLinkType value.
func (v TeamLinkType) Ptr() *TeamLinkType {
	return &v
}

// NullableTeamLinkType handles when a null is used for TeamLinkType.
type NullableTeamLinkType struct {
	value *TeamLinkType
	isSet bool
}

// Get returns the associated value.
func (v NullableTeamLinkType) Get() *TeamLinkType {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableTeamLinkType) Set(val *TeamLinkType) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableTeamLinkType) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableTeamLinkType) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTeamLinkType initializes the struct as if Set has been called.
func NewNullableTeamLinkType(val *TeamLinkType) *NullableTeamLinkType {
	return &NullableTeamLinkType{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableTeamLinkType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableTeamLinkType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
