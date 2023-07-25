// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// FindingType The JSON:API type for findings.
type FindingType string

// List of FindingType.
const (
	FINDINGTYPE_FINDING FindingType = "finding"
)

var allowedFindingTypeEnumValues = []FindingType{
	FINDINGTYPE_FINDING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FindingType) GetAllowedValues() []FindingType {
	return allowedFindingTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FindingType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FindingType(value)
	return nil
}

// NewFindingTypeFromValue returns a pointer to a valid FindingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFindingTypeFromValue(v string) (*FindingType, error) {
	ev := FindingType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FindingType: valid values are %v", v, allowedFindingTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FindingType) IsValid() bool {
	for _, existing := range allowedFindingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FindingType value.
func (v FindingType) Ptr() *FindingType {
	return &v
}

// NullableFindingType handles when a null is used for FindingType.
type NullableFindingType struct {
	value *FindingType
	isSet bool
}

// Get returns the associated value.
func (v NullableFindingType) Get() *FindingType {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableFindingType) Set(val *FindingType) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableFindingType) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableFindingType) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableFindingType initializes the struct as if Set has been called.
func NewNullableFindingType(val *FindingType) *NullableFindingType {
	return &NullableFindingType{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableFindingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableFindingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
