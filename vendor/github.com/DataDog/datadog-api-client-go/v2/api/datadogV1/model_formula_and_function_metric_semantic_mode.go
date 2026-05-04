// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaAndFunctionMetricSemanticMode Semantic mode for metrics queries. This determines how metrics from different sources are combined or displayed.
type FormulaAndFunctionMetricSemanticMode string

// List of FormulaAndFunctionMetricSemanticMode.
const (
	FORMULAANDFUNCTIONMETRICSEMANTICMODE_COMBINED FormulaAndFunctionMetricSemanticMode = "combined"
	FORMULAANDFUNCTIONMETRICSEMANTICMODE_NATIVE   FormulaAndFunctionMetricSemanticMode = "native"
)

var allowedFormulaAndFunctionMetricSemanticModeEnumValues = []FormulaAndFunctionMetricSemanticMode{
	FORMULAANDFUNCTIONMETRICSEMANTICMODE_COMBINED,
	FORMULAANDFUNCTIONMETRICSEMANTICMODE_NATIVE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaAndFunctionMetricSemanticMode) GetAllowedValues() []FormulaAndFunctionMetricSemanticMode {
	return allowedFormulaAndFunctionMetricSemanticModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaAndFunctionMetricSemanticMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaAndFunctionMetricSemanticMode(value)
	return nil
}

// NewFormulaAndFunctionMetricSemanticModeFromValue returns a pointer to a valid FormulaAndFunctionMetricSemanticMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaAndFunctionMetricSemanticModeFromValue(v string) (*FormulaAndFunctionMetricSemanticMode, error) {
	ev := FormulaAndFunctionMetricSemanticMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaAndFunctionMetricSemanticMode: valid values are %v", v, allowedFormulaAndFunctionMetricSemanticModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaAndFunctionMetricSemanticMode) IsValid() bool {
	for _, existing := range allowedFormulaAndFunctionMetricSemanticModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaAndFunctionMetricSemanticMode value.
func (v FormulaAndFunctionMetricSemanticMode) Ptr() *FormulaAndFunctionMetricSemanticMode {
	return &v
}
