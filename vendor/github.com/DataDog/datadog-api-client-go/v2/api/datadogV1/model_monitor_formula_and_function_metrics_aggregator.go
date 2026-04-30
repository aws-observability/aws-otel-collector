// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorFormulaAndFunctionMetricsAggregator Aggregator for metrics queries.
type MonitorFormulaAndFunctionMetricsAggregator string

// List of MonitorFormulaAndFunctionMetricsAggregator.
const (
	MONITORFORMULAANDFUNCTIONMETRICSAGGREGATOR_AVG          MonitorFormulaAndFunctionMetricsAggregator = "avg"
	MONITORFORMULAANDFUNCTIONMETRICSAGGREGATOR_MIN          MonitorFormulaAndFunctionMetricsAggregator = "min"
	MONITORFORMULAANDFUNCTIONMETRICSAGGREGATOR_MAX          MonitorFormulaAndFunctionMetricsAggregator = "max"
	MONITORFORMULAANDFUNCTIONMETRICSAGGREGATOR_SUM          MonitorFormulaAndFunctionMetricsAggregator = "sum"
	MONITORFORMULAANDFUNCTIONMETRICSAGGREGATOR_LAST         MonitorFormulaAndFunctionMetricsAggregator = "last"
	MONITORFORMULAANDFUNCTIONMETRICSAGGREGATOR_MEAN         MonitorFormulaAndFunctionMetricsAggregator = "mean"
	MONITORFORMULAANDFUNCTIONMETRICSAGGREGATOR_AREA         MonitorFormulaAndFunctionMetricsAggregator = "area"
	MONITORFORMULAANDFUNCTIONMETRICSAGGREGATOR_L2NORM       MonitorFormulaAndFunctionMetricsAggregator = "l2norm"
	MONITORFORMULAANDFUNCTIONMETRICSAGGREGATOR_PERCENTILE   MonitorFormulaAndFunctionMetricsAggregator = "percentile"
	MONITORFORMULAANDFUNCTIONMETRICSAGGREGATOR_STDDEV       MonitorFormulaAndFunctionMetricsAggregator = "stddev"
	MONITORFORMULAANDFUNCTIONMETRICSAGGREGATOR_COUNT_UNIQUE MonitorFormulaAndFunctionMetricsAggregator = "count_unique"
)

var allowedMonitorFormulaAndFunctionMetricsAggregatorEnumValues = []MonitorFormulaAndFunctionMetricsAggregator{
	MONITORFORMULAANDFUNCTIONMETRICSAGGREGATOR_AVG,
	MONITORFORMULAANDFUNCTIONMETRICSAGGREGATOR_MIN,
	MONITORFORMULAANDFUNCTIONMETRICSAGGREGATOR_MAX,
	MONITORFORMULAANDFUNCTIONMETRICSAGGREGATOR_SUM,
	MONITORFORMULAANDFUNCTIONMETRICSAGGREGATOR_LAST,
	MONITORFORMULAANDFUNCTIONMETRICSAGGREGATOR_MEAN,
	MONITORFORMULAANDFUNCTIONMETRICSAGGREGATOR_AREA,
	MONITORFORMULAANDFUNCTIONMETRICSAGGREGATOR_L2NORM,
	MONITORFORMULAANDFUNCTIONMETRICSAGGREGATOR_PERCENTILE,
	MONITORFORMULAANDFUNCTIONMETRICSAGGREGATOR_STDDEV,
	MONITORFORMULAANDFUNCTIONMETRICSAGGREGATOR_COUNT_UNIQUE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MonitorFormulaAndFunctionMetricsAggregator) GetAllowedValues() []MonitorFormulaAndFunctionMetricsAggregator {
	return allowedMonitorFormulaAndFunctionMetricsAggregatorEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MonitorFormulaAndFunctionMetricsAggregator) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MonitorFormulaAndFunctionMetricsAggregator(value)
	return nil
}

// NewMonitorFormulaAndFunctionMetricsAggregatorFromValue returns a pointer to a valid MonitorFormulaAndFunctionMetricsAggregator
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMonitorFormulaAndFunctionMetricsAggregatorFromValue(v string) (*MonitorFormulaAndFunctionMetricsAggregator, error) {
	ev := MonitorFormulaAndFunctionMetricsAggregator(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MonitorFormulaAndFunctionMetricsAggregator: valid values are %v", v, allowedMonitorFormulaAndFunctionMetricsAggregatorEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MonitorFormulaAndFunctionMetricsAggregator) IsValid() bool {
	for _, existing := range allowedMonitorFormulaAndFunctionMetricsAggregatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MonitorFormulaAndFunctionMetricsAggregator value.
func (v MonitorFormulaAndFunctionMetricsAggregator) Ptr() *MonitorFormulaAndFunctionMetricsAggregator {
	return &v
}
