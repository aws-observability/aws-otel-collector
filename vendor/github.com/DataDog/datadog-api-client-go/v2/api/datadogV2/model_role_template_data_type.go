// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RoleTemplateDataType Roles resource type.
type RoleTemplateDataType string

// List of RoleTemplateDataType.
const (
	ROLETEMPLATEDATATYPE_ROLES RoleTemplateDataType = "roles"
)

var allowedRoleTemplateDataTypeEnumValues = []RoleTemplateDataType{
	ROLETEMPLATEDATATYPE_ROLES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RoleTemplateDataType) GetAllowedValues() []RoleTemplateDataType {
	return allowedRoleTemplateDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RoleTemplateDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RoleTemplateDataType(value)
	return nil
}

// NewRoleTemplateDataTypeFromValue returns a pointer to a valid RoleTemplateDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRoleTemplateDataTypeFromValue(v string) (*RoleTemplateDataType, error) {
	ev := RoleTemplateDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RoleTemplateDataType: valid values are %v", v, allowedRoleTemplateDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RoleTemplateDataType) IsValid() bool {
	for _, existing := range allowedRoleTemplateDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RoleTemplateDataType value.
func (v RoleTemplateDataType) Ptr() *RoleTemplateDataType {
	return &v
}
