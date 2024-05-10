// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UserTeamTeamType User team team type
type UserTeamTeamType string

// List of UserTeamTeamType.
const (
	USERTEAMTEAMTYPE_TEAM UserTeamTeamType = "team"
)

var allowedUserTeamTeamTypeEnumValues = []UserTeamTeamType{
	USERTEAMTEAMTYPE_TEAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UserTeamTeamType) GetAllowedValues() []UserTeamTeamType {
	return allowedUserTeamTeamTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UserTeamTeamType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UserTeamTeamType(value)
	return nil
}

// NewUserTeamTeamTypeFromValue returns a pointer to a valid UserTeamTeamType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUserTeamTeamTypeFromValue(v string) (*UserTeamTeamType, error) {
	ev := UserTeamTeamType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UserTeamTeamType: valid values are %v", v, allowedUserTeamTeamTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UserTeamTeamType) IsValid() bool {
	for _, existing := range allowedUserTeamTeamTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserTeamTeamType value.
func (v UserTeamTeamType) Ptr() *UserTeamTeamType {
	return &v
}
