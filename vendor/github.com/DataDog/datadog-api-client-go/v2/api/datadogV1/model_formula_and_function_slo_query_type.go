// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
<<<<<<< HEAD
	"encoding/json"
	"fmt"
=======
	"fmt"

	"github.com/goccy/go-json"
>>>>>>> main
)

// FormulaAndFunctionSLOQueryType Name of the query for use in formulas.
type FormulaAndFunctionSLOQueryType string

// List of FormulaAndFunctionSLOQueryType.
const (
	FORMULAANDFUNCTIONSLOQUERYTYPE_METRIC FormulaAndFunctionSLOQueryType = "metric"
)

var allowedFormulaAndFunctionSLOQueryTypeEnumValues = []FormulaAndFunctionSLOQueryType{
	FORMULAANDFUNCTIONSLOQUERYTYPE_METRIC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaAndFunctionSLOQueryType) GetAllowedValues() []FormulaAndFunctionSLOQueryType {
	return allowedFormulaAndFunctionSLOQueryTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaAndFunctionSLOQueryType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaAndFunctionSLOQueryType(value)
	return nil
}

// NewFormulaAndFunctionSLOQueryTypeFromValue returns a pointer to a valid FormulaAndFunctionSLOQueryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaAndFunctionSLOQueryTypeFromValue(v string) (*FormulaAndFunctionSLOQueryType, error) {
	ev := FormulaAndFunctionSLOQueryType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaAndFunctionSLOQueryType: valid values are %v", v, allowedFormulaAndFunctionSLOQueryTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaAndFunctionSLOQueryType) IsValid() bool {
	for _, existing := range allowedFormulaAndFunctionSLOQueryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaAndFunctionSLOQueryType value.
func (v FormulaAndFunctionSLOQueryType) Ptr() *FormulaAndFunctionSLOQueryType {
	return &v
}
<<<<<<< HEAD

// NullableFormulaAndFunctionSLOQueryType handles when a null is used for FormulaAndFunctionSLOQueryType.
type NullableFormulaAndFunctionSLOQueryType struct {
	value *FormulaAndFunctionSLOQueryType
	isSet bool
}

// Get returns the associated value.
func (v NullableFormulaAndFunctionSLOQueryType) Get() *FormulaAndFunctionSLOQueryType {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableFormulaAndFunctionSLOQueryType) Set(val *FormulaAndFunctionSLOQueryType) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableFormulaAndFunctionSLOQueryType) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableFormulaAndFunctionSLOQueryType) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableFormulaAndFunctionSLOQueryType initializes the struct as if Set has been called.
func NewNullableFormulaAndFunctionSLOQueryType(val *FormulaAndFunctionSLOQueryType) *NullableFormulaAndFunctionSLOQueryType {
	return &NullableFormulaAndFunctionSLOQueryType{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableFormulaAndFunctionSLOQueryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableFormulaAndFunctionSLOQueryType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
=======
>>>>>>> main
