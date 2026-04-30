// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaAndFunctionApmMetricsSpanKind Describes the relationship between the span, its parents, and its children in a trace.
type FormulaAndFunctionApmMetricsSpanKind string

// List of FormulaAndFunctionApmMetricsSpanKind.
const (
	FORMULAANDFUNCTIONAPMMETRICSSPANKIND_CONSUMER FormulaAndFunctionApmMetricsSpanKind = "consumer"
	FORMULAANDFUNCTIONAPMMETRICSSPANKIND_SERVER   FormulaAndFunctionApmMetricsSpanKind = "server"
	FORMULAANDFUNCTIONAPMMETRICSSPANKIND_CLIENT   FormulaAndFunctionApmMetricsSpanKind = "client"
	FORMULAANDFUNCTIONAPMMETRICSSPANKIND_PRODUCER FormulaAndFunctionApmMetricsSpanKind = "producer"
	FORMULAANDFUNCTIONAPMMETRICSSPANKIND_INTERNAL FormulaAndFunctionApmMetricsSpanKind = "internal"
)

var allowedFormulaAndFunctionApmMetricsSpanKindEnumValues = []FormulaAndFunctionApmMetricsSpanKind{
	FORMULAANDFUNCTIONAPMMETRICSSPANKIND_CONSUMER,
	FORMULAANDFUNCTIONAPMMETRICSSPANKIND_SERVER,
	FORMULAANDFUNCTIONAPMMETRICSSPANKIND_CLIENT,
	FORMULAANDFUNCTIONAPMMETRICSSPANKIND_PRODUCER,
	FORMULAANDFUNCTIONAPMMETRICSSPANKIND_INTERNAL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaAndFunctionApmMetricsSpanKind) GetAllowedValues() []FormulaAndFunctionApmMetricsSpanKind {
	return allowedFormulaAndFunctionApmMetricsSpanKindEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaAndFunctionApmMetricsSpanKind) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaAndFunctionApmMetricsSpanKind(value)
	return nil
}

// NewFormulaAndFunctionApmMetricsSpanKindFromValue returns a pointer to a valid FormulaAndFunctionApmMetricsSpanKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaAndFunctionApmMetricsSpanKindFromValue(v string) (*FormulaAndFunctionApmMetricsSpanKind, error) {
	ev := FormulaAndFunctionApmMetricsSpanKind(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaAndFunctionApmMetricsSpanKind: valid values are %v", v, allowedFormulaAndFunctionApmMetricsSpanKindEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaAndFunctionApmMetricsSpanKind) IsValid() bool {
	for _, existing := range allowedFormulaAndFunctionApmMetricsSpanKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaAndFunctionApmMetricsSpanKind value.
func (v FormulaAndFunctionApmMetricsSpanKind) Ptr() *FormulaAndFunctionApmMetricsSpanKind {
	return &v
}
