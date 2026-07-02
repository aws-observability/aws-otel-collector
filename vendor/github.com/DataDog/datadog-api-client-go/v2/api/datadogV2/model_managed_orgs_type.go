// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ManagedOrgsType The resource type for managed organizations.
type ManagedOrgsType string

// List of ManagedOrgsType.
const (
	MANAGEDORGSTYPE_MANAGED_ORGS ManagedOrgsType = "managed_orgs"
)

var allowedManagedOrgsTypeEnumValues = []ManagedOrgsType{
	MANAGEDORGSTYPE_MANAGED_ORGS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ManagedOrgsType) GetAllowedValues() []ManagedOrgsType {
	return allowedManagedOrgsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ManagedOrgsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ManagedOrgsType(value)
	return nil
}

// NewManagedOrgsTypeFromValue returns a pointer to a valid ManagedOrgsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewManagedOrgsTypeFromValue(v string) (*ManagedOrgsType, error) {
	ev := ManagedOrgsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ManagedOrgsType: valid values are %v", v, allowedManagedOrgsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ManagedOrgsType) IsValid() bool {
	for _, existing := range allowedManagedOrgsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ManagedOrgsType value.
func (v ManagedOrgsType) Ptr() *ManagedOrgsType {
	return &v
}
