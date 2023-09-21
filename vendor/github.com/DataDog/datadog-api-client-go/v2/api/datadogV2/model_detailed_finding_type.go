// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
<<<<<<< HEAD
	"encoding/json"
	"fmt"
=======
	"fmt"

	"github.com/goccy/go-json"
>>>>>>> main
)

// DetailedFindingType The JSON:API type for findings that have the message and resource configuration.
type DetailedFindingType string

// List of DetailedFindingType.
const (
	DETAILEDFINDINGTYPE_DETAILED_FINDING DetailedFindingType = "detailed_finding"
)

var allowedDetailedFindingTypeEnumValues = []DetailedFindingType{
	DETAILEDFINDINGTYPE_DETAILED_FINDING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DetailedFindingType) GetAllowedValues() []DetailedFindingType {
	return allowedDetailedFindingTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DetailedFindingType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DetailedFindingType(value)
	return nil
}

// NewDetailedFindingTypeFromValue returns a pointer to a valid DetailedFindingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDetailedFindingTypeFromValue(v string) (*DetailedFindingType, error) {
	ev := DetailedFindingType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DetailedFindingType: valid values are %v", v, allowedDetailedFindingTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DetailedFindingType) IsValid() bool {
	for _, existing := range allowedDetailedFindingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DetailedFindingType value.
func (v DetailedFindingType) Ptr() *DetailedFindingType {
	return &v
}
<<<<<<< HEAD

// NullableDetailedFindingType handles when a null is used for DetailedFindingType.
type NullableDetailedFindingType struct {
	value *DetailedFindingType
	isSet bool
}

// Get returns the associated value.
func (v NullableDetailedFindingType) Get() *DetailedFindingType {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableDetailedFindingType) Set(val *DetailedFindingType) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableDetailedFindingType) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableDetailedFindingType) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDetailedFindingType initializes the struct as if Set has been called.
func NewNullableDetailedFindingType(val *DetailedFindingType) *NullableDetailedFindingType {
	return &NullableDetailedFindingType{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableDetailedFindingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableDetailedFindingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
=======
>>>>>>> main
