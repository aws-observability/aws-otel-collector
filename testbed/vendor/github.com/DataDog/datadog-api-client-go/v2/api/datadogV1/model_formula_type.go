// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaType Set the sort type to formula.
type FormulaType string

// List of FormulaType.
const (
	FORMULATYPE_FORMULA FormulaType = "formula"
)

var allowedFormulaTypeEnumValues = []FormulaType{
	FORMULATYPE_FORMULA,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaType) GetAllowedValues() []FormulaType {
	return allowedFormulaTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaType(value)
	return nil
}

// NewFormulaTypeFromValue returns a pointer to a valid FormulaType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaTypeFromValue(v string) (*FormulaType, error) {
	ev := FormulaType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaType: valid values are %v", v, allowedFormulaTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaType) IsValid() bool {
	for _, existing := range allowedFormulaTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaType value.
func (v FormulaType) Ptr() *FormulaType {
	return &v
}
