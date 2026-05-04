// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentUserDefinedFieldCollected The lifecycle stage at which the app prompts users to fill out this field. Cannot be set on required fields.
type IncidentUserDefinedFieldCollected string

// List of IncidentUserDefinedFieldCollected.
const (
	INCIDENTUSERDEFINEDFIELDCOLLECTED_ACTIVE    IncidentUserDefinedFieldCollected = "active"
	INCIDENTUSERDEFINEDFIELDCOLLECTED_STABLE    IncidentUserDefinedFieldCollected = "stable"
	INCIDENTUSERDEFINEDFIELDCOLLECTED_RESOLVED  IncidentUserDefinedFieldCollected = "resolved"
	INCIDENTUSERDEFINEDFIELDCOLLECTED_COMPLETED IncidentUserDefinedFieldCollected = "completed"
)

var allowedIncidentUserDefinedFieldCollectedEnumValues = []IncidentUserDefinedFieldCollected{
	INCIDENTUSERDEFINEDFIELDCOLLECTED_ACTIVE,
	INCIDENTUSERDEFINEDFIELDCOLLECTED_STABLE,
	INCIDENTUSERDEFINEDFIELDCOLLECTED_RESOLVED,
	INCIDENTUSERDEFINEDFIELDCOLLECTED_COMPLETED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentUserDefinedFieldCollected) GetAllowedValues() []IncidentUserDefinedFieldCollected {
	return allowedIncidentUserDefinedFieldCollectedEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentUserDefinedFieldCollected) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentUserDefinedFieldCollected(value)
	return nil
}

// NewIncidentUserDefinedFieldCollectedFromValue returns a pointer to a valid IncidentUserDefinedFieldCollected
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentUserDefinedFieldCollectedFromValue(v string) (*IncidentUserDefinedFieldCollected, error) {
	ev := IncidentUserDefinedFieldCollected(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentUserDefinedFieldCollected: valid values are %v", v, allowedIncidentUserDefinedFieldCollectedEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentUserDefinedFieldCollected) IsValid() bool {
	for _, existing := range allowedIncidentUserDefinedFieldCollectedEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentUserDefinedFieldCollected value.
func (v IncidentUserDefinedFieldCollected) Ptr() *IncidentUserDefinedFieldCollected {
	return &v
}

// NullableIncidentUserDefinedFieldCollected handles when a null is used for IncidentUserDefinedFieldCollected.
type NullableIncidentUserDefinedFieldCollected struct {
	value *IncidentUserDefinedFieldCollected
	isSet bool
}

// Get returns the associated value.
func (v NullableIncidentUserDefinedFieldCollected) Get() *IncidentUserDefinedFieldCollected {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableIncidentUserDefinedFieldCollected) Set(val *IncidentUserDefinedFieldCollected) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableIncidentUserDefinedFieldCollected) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableIncidentUserDefinedFieldCollected) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableIncidentUserDefinedFieldCollected initializes the struct as if Set has been called.
func NewNullableIncidentUserDefinedFieldCollected(val *IncidentUserDefinedFieldCollected) *NullableIncidentUserDefinedFieldCollected {
	return &NullableIncidentUserDefinedFieldCollected{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableIncidentUserDefinedFieldCollected) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableIncidentUserDefinedFieldCollected) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return datadog.Unmarshal(src, &v.value)
}
