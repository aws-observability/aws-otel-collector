// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"encoding/json"
	"fmt"
)

// FormulaAndFunctionSLODataSource Data source for SLO measures queries.
type FormulaAndFunctionSLODataSource string

// List of FormulaAndFunctionSLODataSource.
const (
	FORMULAANDFUNCTIONSLODATASOURCE_SLO FormulaAndFunctionSLODataSource = "slo"
)

var allowedFormulaAndFunctionSLODataSourceEnumValues = []FormulaAndFunctionSLODataSource{
	FORMULAANDFUNCTIONSLODATASOURCE_SLO,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaAndFunctionSLODataSource) GetAllowedValues() []FormulaAndFunctionSLODataSource {
	return allowedFormulaAndFunctionSLODataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaAndFunctionSLODataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaAndFunctionSLODataSource(value)
	return nil
}

// NewFormulaAndFunctionSLODataSourceFromValue returns a pointer to a valid FormulaAndFunctionSLODataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaAndFunctionSLODataSourceFromValue(v string) (*FormulaAndFunctionSLODataSource, error) {
	ev := FormulaAndFunctionSLODataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaAndFunctionSLODataSource: valid values are %v", v, allowedFormulaAndFunctionSLODataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaAndFunctionSLODataSource) IsValid() bool {
	for _, existing := range allowedFormulaAndFunctionSLODataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaAndFunctionSLODataSource value.
func (v FormulaAndFunctionSLODataSource) Ptr() *FormulaAndFunctionSLODataSource {
	return &v
}

// NullableFormulaAndFunctionSLODataSource handles when a null is used for FormulaAndFunctionSLODataSource.
type NullableFormulaAndFunctionSLODataSource struct {
	value *FormulaAndFunctionSLODataSource
	isSet bool
}

// Get returns the associated value.
func (v NullableFormulaAndFunctionSLODataSource) Get() *FormulaAndFunctionSLODataSource {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableFormulaAndFunctionSLODataSource) Set(val *FormulaAndFunctionSLODataSource) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableFormulaAndFunctionSLODataSource) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableFormulaAndFunctionSLODataSource) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableFormulaAndFunctionSLODataSource initializes the struct as if Set has been called.
func NewNullableFormulaAndFunctionSLODataSource(val *FormulaAndFunctionSLODataSource) *NullableFormulaAndFunctionSLODataSource {
	return &NullableFormulaAndFunctionSLODataSource{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableFormulaAndFunctionSLODataSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableFormulaAndFunctionSLODataSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
