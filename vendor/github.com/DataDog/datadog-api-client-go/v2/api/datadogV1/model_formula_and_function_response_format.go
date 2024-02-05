// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaAndFunctionResponseFormat Timeseries, scalar, or event list response. Event list response formats are supported by Geomap widgets.
type FormulaAndFunctionResponseFormat string

// List of FormulaAndFunctionResponseFormat.
const (
	FORMULAANDFUNCTIONRESPONSEFORMAT_TIMESERIES FormulaAndFunctionResponseFormat = "timeseries"
	FORMULAANDFUNCTIONRESPONSEFORMAT_SCALAR     FormulaAndFunctionResponseFormat = "scalar"
	FORMULAANDFUNCTIONRESPONSEFORMAT_EVENT_LIST FormulaAndFunctionResponseFormat = "event_list"
)

var allowedFormulaAndFunctionResponseFormatEnumValues = []FormulaAndFunctionResponseFormat{
	FORMULAANDFUNCTIONRESPONSEFORMAT_TIMESERIES,
	FORMULAANDFUNCTIONRESPONSEFORMAT_SCALAR,
	FORMULAANDFUNCTIONRESPONSEFORMAT_EVENT_LIST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaAndFunctionResponseFormat) GetAllowedValues() []FormulaAndFunctionResponseFormat {
	return allowedFormulaAndFunctionResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaAndFunctionResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaAndFunctionResponseFormat(value)
	return nil
}

// NewFormulaAndFunctionResponseFormatFromValue returns a pointer to a valid FormulaAndFunctionResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaAndFunctionResponseFormatFromValue(v string) (*FormulaAndFunctionResponseFormat, error) {
	ev := FormulaAndFunctionResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaAndFunctionResponseFormat: valid values are %v", v, allowedFormulaAndFunctionResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaAndFunctionResponseFormat) IsValid() bool {
	for _, existing := range allowedFormulaAndFunctionResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaAndFunctionResponseFormat value.
func (v FormulaAndFunctionResponseFormat) Ptr() *FormulaAndFunctionResponseFormat {
	return &v
}
