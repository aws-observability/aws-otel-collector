// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UserInvitationsType User invitations type.
type UserInvitationsType string

// List of UserInvitationsType.
const (
	USERINVITATIONSTYPE_USER_INVITATIONS UserInvitationsType = "user_invitations"
)

var allowedUserInvitationsTypeEnumValues = []UserInvitationsType{
	USERINVITATIONSTYPE_USER_INVITATIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UserInvitationsType) GetAllowedValues() []UserInvitationsType {
	return allowedUserInvitationsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UserInvitationsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UserInvitationsType(value)
	return nil
}

// NewUserInvitationsTypeFromValue returns a pointer to a valid UserInvitationsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUserInvitationsTypeFromValue(v string) (*UserInvitationsType, error) {
	ev := UserInvitationsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UserInvitationsType: valid values are %v", v, allowedUserInvitationsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UserInvitationsType) IsValid() bool {
	for _, existing := range allowedUserInvitationsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserInvitationsType value.
func (v UserInvitationsType) Ptr() *UserInvitationsType {
	return &v
}
