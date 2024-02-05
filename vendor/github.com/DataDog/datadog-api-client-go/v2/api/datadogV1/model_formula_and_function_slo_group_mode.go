// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaAndFunctionSLOGroupMode Group mode to query measures.
type FormulaAndFunctionSLOGroupMode string

// List of FormulaAndFunctionSLOGroupMode.
const (
	FORMULAANDFUNCTIONSLOGROUPMODE_OVERALL    FormulaAndFunctionSLOGroupMode = "overall"
	FORMULAANDFUNCTIONSLOGROUPMODE_COMPONENTS FormulaAndFunctionSLOGroupMode = "components"
)

var allowedFormulaAndFunctionSLOGroupModeEnumValues = []FormulaAndFunctionSLOGroupMode{
	FORMULAANDFUNCTIONSLOGROUPMODE_OVERALL,
	FORMULAANDFUNCTIONSLOGROUPMODE_COMPONENTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaAndFunctionSLOGroupMode) GetAllowedValues() []FormulaAndFunctionSLOGroupMode {
	return allowedFormulaAndFunctionSLOGroupModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaAndFunctionSLOGroupMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaAndFunctionSLOGroupMode(value)
	return nil
}

// NewFormulaAndFunctionSLOGroupModeFromValue returns a pointer to a valid FormulaAndFunctionSLOGroupMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaAndFunctionSLOGroupModeFromValue(v string) (*FormulaAndFunctionSLOGroupMode, error) {
	ev := FormulaAndFunctionSLOGroupMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaAndFunctionSLOGroupMode: valid values are %v", v, allowedFormulaAndFunctionSLOGroupModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaAndFunctionSLOGroupMode) IsValid() bool {
	for _, existing := range allowedFormulaAndFunctionSLOGroupModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaAndFunctionSLOGroupMode value.
func (v FormulaAndFunctionSLOGroupMode) Ptr() *FormulaAndFunctionSLOGroupMode {
	return &v
}
