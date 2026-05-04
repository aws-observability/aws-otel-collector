// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaAndFunctionApmMetricStatName APM metric stat name.
type FormulaAndFunctionApmMetricStatName string

// List of FormulaAndFunctionApmMetricStatName.
const (
	FORMULAANDFUNCTIONAPMMETRICSTATNAME_ERRORS               FormulaAndFunctionApmMetricStatName = "errors"
	FORMULAANDFUNCTIONAPMMETRICSTATNAME_ERROR_RATE           FormulaAndFunctionApmMetricStatName = "error_rate"
	FORMULAANDFUNCTIONAPMMETRICSTATNAME_ERRORS_PER_SECOND    FormulaAndFunctionApmMetricStatName = "errors_per_second"
	FORMULAANDFUNCTIONAPMMETRICSTATNAME_LATENCY_AVG          FormulaAndFunctionApmMetricStatName = "latency_avg"
	FORMULAANDFUNCTIONAPMMETRICSTATNAME_LATENCY_MAX          FormulaAndFunctionApmMetricStatName = "latency_max"
	FORMULAANDFUNCTIONAPMMETRICSTATNAME_LATENCY_P50          FormulaAndFunctionApmMetricStatName = "latency_p50"
	FORMULAANDFUNCTIONAPMMETRICSTATNAME_LATENCY_P75          FormulaAndFunctionApmMetricStatName = "latency_p75"
	FORMULAANDFUNCTIONAPMMETRICSTATNAME_LATENCY_P90          FormulaAndFunctionApmMetricStatName = "latency_p90"
	FORMULAANDFUNCTIONAPMMETRICSTATNAME_LATENCY_P95          FormulaAndFunctionApmMetricStatName = "latency_p95"
	FORMULAANDFUNCTIONAPMMETRICSTATNAME_LATENCY_P99          FormulaAndFunctionApmMetricStatName = "latency_p99"
	FORMULAANDFUNCTIONAPMMETRICSTATNAME_LATENCY_P999         FormulaAndFunctionApmMetricStatName = "latency_p999"
	FORMULAANDFUNCTIONAPMMETRICSTATNAME_LATENCY_DISTRIBUTION FormulaAndFunctionApmMetricStatName = "latency_distribution"
	FORMULAANDFUNCTIONAPMMETRICSTATNAME_HITS                 FormulaAndFunctionApmMetricStatName = "hits"
	FORMULAANDFUNCTIONAPMMETRICSTATNAME_HITS_PER_SECOND      FormulaAndFunctionApmMetricStatName = "hits_per_second"
	FORMULAANDFUNCTIONAPMMETRICSTATNAME_TOTAL_TIME           FormulaAndFunctionApmMetricStatName = "total_time"
	FORMULAANDFUNCTIONAPMMETRICSTATNAME_APDEX                FormulaAndFunctionApmMetricStatName = "apdex"
)

var allowedFormulaAndFunctionApmMetricStatNameEnumValues = []FormulaAndFunctionApmMetricStatName{
	FORMULAANDFUNCTIONAPMMETRICSTATNAME_ERRORS,
	FORMULAANDFUNCTIONAPMMETRICSTATNAME_ERROR_RATE,
	FORMULAANDFUNCTIONAPMMETRICSTATNAME_ERRORS_PER_SECOND,
	FORMULAANDFUNCTIONAPMMETRICSTATNAME_LATENCY_AVG,
	FORMULAANDFUNCTIONAPMMETRICSTATNAME_LATENCY_MAX,
	FORMULAANDFUNCTIONAPMMETRICSTATNAME_LATENCY_P50,
	FORMULAANDFUNCTIONAPMMETRICSTATNAME_LATENCY_P75,
	FORMULAANDFUNCTIONAPMMETRICSTATNAME_LATENCY_P90,
	FORMULAANDFUNCTIONAPMMETRICSTATNAME_LATENCY_P95,
	FORMULAANDFUNCTIONAPMMETRICSTATNAME_LATENCY_P99,
	FORMULAANDFUNCTIONAPMMETRICSTATNAME_LATENCY_P999,
	FORMULAANDFUNCTIONAPMMETRICSTATNAME_LATENCY_DISTRIBUTION,
	FORMULAANDFUNCTIONAPMMETRICSTATNAME_HITS,
	FORMULAANDFUNCTIONAPMMETRICSTATNAME_HITS_PER_SECOND,
	FORMULAANDFUNCTIONAPMMETRICSTATNAME_TOTAL_TIME,
	FORMULAANDFUNCTIONAPMMETRICSTATNAME_APDEX,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaAndFunctionApmMetricStatName) GetAllowedValues() []FormulaAndFunctionApmMetricStatName {
	return allowedFormulaAndFunctionApmMetricStatNameEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaAndFunctionApmMetricStatName) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaAndFunctionApmMetricStatName(value)
	return nil
}

// NewFormulaAndFunctionApmMetricStatNameFromValue returns a pointer to a valid FormulaAndFunctionApmMetricStatName
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaAndFunctionApmMetricStatNameFromValue(v string) (*FormulaAndFunctionApmMetricStatName, error) {
	ev := FormulaAndFunctionApmMetricStatName(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaAndFunctionApmMetricStatName: valid values are %v", v, allowedFormulaAndFunctionApmMetricStatNameEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaAndFunctionApmMetricStatName) IsValid() bool {
	for _, existing := range allowedFormulaAndFunctionApmMetricStatNameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaAndFunctionApmMetricStatName value.
func (v FormulaAndFunctionApmMetricStatName) Ptr() *FormulaAndFunctionApmMetricStatName {
	return &v
}
