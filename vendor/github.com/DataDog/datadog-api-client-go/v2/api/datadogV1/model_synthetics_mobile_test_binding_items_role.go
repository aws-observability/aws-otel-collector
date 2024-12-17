// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsMobileTestBindingItemsRole The definition of `SyntheticsMobileTestBindingItemsRole` object.
type SyntheticsMobileTestBindingItemsRole string

// List of SyntheticsMobileTestBindingItemsRole.
const (
	SYNTHETICSMOBILETESTBINDINGITEMSROLE_EDITOR SyntheticsMobileTestBindingItemsRole = "editor"
	SYNTHETICSMOBILETESTBINDINGITEMSROLE_VIEWER SyntheticsMobileTestBindingItemsRole = "viewer"
)

var allowedSyntheticsMobileTestBindingItemsRoleEnumValues = []SyntheticsMobileTestBindingItemsRole{
	SYNTHETICSMOBILETESTBINDINGITEMSROLE_EDITOR,
	SYNTHETICSMOBILETESTBINDINGITEMSROLE_VIEWER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsMobileTestBindingItemsRole) GetAllowedValues() []SyntheticsMobileTestBindingItemsRole {
	return allowedSyntheticsMobileTestBindingItemsRoleEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsMobileTestBindingItemsRole) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsMobileTestBindingItemsRole(value)
	return nil
}

// NewSyntheticsMobileTestBindingItemsRoleFromValue returns a pointer to a valid SyntheticsMobileTestBindingItemsRole
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsMobileTestBindingItemsRoleFromValue(v string) (*SyntheticsMobileTestBindingItemsRole, error) {
	ev := SyntheticsMobileTestBindingItemsRole(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsMobileTestBindingItemsRole: valid values are %v", v, allowedSyntheticsMobileTestBindingItemsRoleEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsMobileTestBindingItemsRole) IsValid() bool {
	for _, existing := range allowedSyntheticsMobileTestBindingItemsRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsMobileTestBindingItemsRole value.
func (v SyntheticsMobileTestBindingItemsRole) Ptr() *SyntheticsMobileTestBindingItemsRole {
	return &v
}
