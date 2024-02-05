// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PermissionsType Permissions resource type.
type PermissionsType string

// List of PermissionsType.
const (
	PERMISSIONSTYPE_PERMISSIONS PermissionsType = "permissions"
)

var allowedPermissionsTypeEnumValues = []PermissionsType{
	PERMISSIONSTYPE_PERMISSIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PermissionsType) GetAllowedValues() []PermissionsType {
	return allowedPermissionsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PermissionsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PermissionsType(value)
	return nil
}

// NewPermissionsTypeFromValue returns a pointer to a valid PermissionsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPermissionsTypeFromValue(v string) (*PermissionsType, error) {
	ev := PermissionsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PermissionsType: valid values are %v", v, allowedPermissionsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PermissionsType) IsValid() bool {
	for _, existing := range allowedPermissionsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PermissionsType value.
func (v PermissionsType) Ptr() *PermissionsType {
	return &v
}
