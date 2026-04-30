// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentUserDefinedFieldCategory The section in which the field appears: "what_happened" or "why_it_happened". When null, the field appears in the Attributes section.
type IncidentUserDefinedFieldCategory string

// List of IncidentUserDefinedFieldCategory.
const (
	INCIDENTUSERDEFINEDFIELDCATEGORY_WHAT_HAPPENED   IncidentUserDefinedFieldCategory = "what_happened"
	INCIDENTUSERDEFINEDFIELDCATEGORY_WHY_IT_HAPPENED IncidentUserDefinedFieldCategory = "why_it_happened"
)

var allowedIncidentUserDefinedFieldCategoryEnumValues = []IncidentUserDefinedFieldCategory{
	INCIDENTUSERDEFINEDFIELDCATEGORY_WHAT_HAPPENED,
	INCIDENTUSERDEFINEDFIELDCATEGORY_WHY_IT_HAPPENED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentUserDefinedFieldCategory) GetAllowedValues() []IncidentUserDefinedFieldCategory {
	return allowedIncidentUserDefinedFieldCategoryEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentUserDefinedFieldCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentUserDefinedFieldCategory(value)
	return nil
}

// NewIncidentUserDefinedFieldCategoryFromValue returns a pointer to a valid IncidentUserDefinedFieldCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentUserDefinedFieldCategoryFromValue(v string) (*IncidentUserDefinedFieldCategory, error) {
	ev := IncidentUserDefinedFieldCategory(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentUserDefinedFieldCategory: valid values are %v", v, allowedIncidentUserDefinedFieldCategoryEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentUserDefinedFieldCategory) IsValid() bool {
	for _, existing := range allowedIncidentUserDefinedFieldCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentUserDefinedFieldCategory value.
func (v IncidentUserDefinedFieldCategory) Ptr() *IncidentUserDefinedFieldCategory {
	return &v
}

// NullableIncidentUserDefinedFieldCategory handles when a null is used for IncidentUserDefinedFieldCategory.
type NullableIncidentUserDefinedFieldCategory struct {
	value *IncidentUserDefinedFieldCategory
	isSet bool
}

// Get returns the associated value.
func (v NullableIncidentUserDefinedFieldCategory) Get() *IncidentUserDefinedFieldCategory {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableIncidentUserDefinedFieldCategory) Set(val *IncidentUserDefinedFieldCategory) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableIncidentUserDefinedFieldCategory) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableIncidentUserDefinedFieldCategory) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableIncidentUserDefinedFieldCategory initializes the struct as if Set has been called.
func NewNullableIncidentUserDefinedFieldCategory(val *IncidentUserDefinedFieldCategory) *NullableIncidentUserDefinedFieldCategory {
	return &NullableIncidentUserDefinedFieldCategory{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableIncidentUserDefinedFieldCategory) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableIncidentUserDefinedFieldCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return datadog.Unmarshal(src, &v.value)
}
