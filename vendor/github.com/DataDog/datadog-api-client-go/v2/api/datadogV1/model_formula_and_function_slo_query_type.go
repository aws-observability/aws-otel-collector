// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaAndFunctionSLOQueryType Name of the query for use in formulas.
type FormulaAndFunctionSLOQueryType string

// List of FormulaAndFunctionSLOQueryType.
const (
	FORMULAANDFUNCTIONSLOQUERYTYPE_METRIC     FormulaAndFunctionSLOQueryType = "metric"
	FORMULAANDFUNCTIONSLOQUERYTYPE_MONITOR    FormulaAndFunctionSLOQueryType = "monitor"
	FORMULAANDFUNCTIONSLOQUERYTYPE_TIME_SLICE FormulaAndFunctionSLOQueryType = "time_slice"
)

var allowedFormulaAndFunctionSLOQueryTypeEnumValues = []FormulaAndFunctionSLOQueryType{
	FORMULAANDFUNCTIONSLOQUERYTYPE_METRIC,
	FORMULAANDFUNCTIONSLOQUERYTYPE_MONITOR,
	FORMULAANDFUNCTIONSLOQUERYTYPE_TIME_SLICE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaAndFunctionSLOQueryType) GetAllowedValues() []FormulaAndFunctionSLOQueryType {
	return allowedFormulaAndFunctionSLOQueryTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaAndFunctionSLOQueryType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
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
