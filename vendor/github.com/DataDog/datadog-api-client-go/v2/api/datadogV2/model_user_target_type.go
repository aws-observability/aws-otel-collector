// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UserTargetType Indicates that the resource is of type `users`.
type UserTargetType string

// List of UserTargetType.
const (
	USERTARGETTYPE_USERS UserTargetType = "users"
)

var allowedUserTargetTypeEnumValues = []UserTargetType{
	USERTARGETTYPE_USERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UserTargetType) GetAllowedValues() []UserTargetType {
	return allowedUserTargetTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UserTargetType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UserTargetType(value)
	return nil
}

// NewUserTargetTypeFromValue returns a pointer to a valid UserTargetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUserTargetTypeFromValue(v string) (*UserTargetType, error) {
	ev := UserTargetType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UserTargetType: valid values are %v", v, allowedUserTargetTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UserTargetType) IsValid() bool {
	for _, existing := range allowedUserTargetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserTargetType value.
func (v UserTargetType) Ptr() *UserTargetType {
	return &v
}
