// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RolesType Roles type.
type RolesType string

// List of RolesType.
const (
	ROLESTYPE_ROLES RolesType = "roles"
)

var allowedRolesTypeEnumValues = []RolesType{
	ROLESTYPE_ROLES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RolesType) GetAllowedValues() []RolesType {
	return allowedRolesTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RolesType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RolesType(value)
	return nil
}

// NewRolesTypeFromValue returns a pointer to a valid RolesType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRolesTypeFromValue(v string) (*RolesType, error) {
	ev := RolesType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RolesType: valid values are %v", v, allowedRolesTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RolesType) IsValid() bool {
	for _, existing := range allowedRolesTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RolesType value.
func (v RolesType) Ptr() *RolesType {
	return &v
}
