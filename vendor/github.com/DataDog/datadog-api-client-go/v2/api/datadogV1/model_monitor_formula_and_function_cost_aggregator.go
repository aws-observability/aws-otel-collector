// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorFormulaAndFunctionCostAggregator Aggregation methods for metric queries.
type MonitorFormulaAndFunctionCostAggregator string

// List of MonitorFormulaAndFunctionCostAggregator.
const (
	MONITORFORMULAANDFUNCTIONCOSTAGGREGATOR_AVG        MonitorFormulaAndFunctionCostAggregator = "avg"
	MONITORFORMULAANDFUNCTIONCOSTAGGREGATOR_SUM        MonitorFormulaAndFunctionCostAggregator = "sum"
	MONITORFORMULAANDFUNCTIONCOSTAGGREGATOR_MAX        MonitorFormulaAndFunctionCostAggregator = "max"
	MONITORFORMULAANDFUNCTIONCOSTAGGREGATOR_MIN        MonitorFormulaAndFunctionCostAggregator = "min"
	MONITORFORMULAANDFUNCTIONCOSTAGGREGATOR_LAST       MonitorFormulaAndFunctionCostAggregator = "last"
	MONITORFORMULAANDFUNCTIONCOSTAGGREGATOR_AREA       MonitorFormulaAndFunctionCostAggregator = "area"
	MONITORFORMULAANDFUNCTIONCOSTAGGREGATOR_L2NORM     MonitorFormulaAndFunctionCostAggregator = "l2norm"
	MONITORFORMULAANDFUNCTIONCOSTAGGREGATOR_PERCENTILE MonitorFormulaAndFunctionCostAggregator = "percentile"
	MONITORFORMULAANDFUNCTIONCOSTAGGREGATOR_STDDEV     MonitorFormulaAndFunctionCostAggregator = "stddev"
)

var allowedMonitorFormulaAndFunctionCostAggregatorEnumValues = []MonitorFormulaAndFunctionCostAggregator{
	MONITORFORMULAANDFUNCTIONCOSTAGGREGATOR_AVG,
	MONITORFORMULAANDFUNCTIONCOSTAGGREGATOR_SUM,
	MONITORFORMULAANDFUNCTIONCOSTAGGREGATOR_MAX,
	MONITORFORMULAANDFUNCTIONCOSTAGGREGATOR_MIN,
	MONITORFORMULAANDFUNCTIONCOSTAGGREGATOR_LAST,
	MONITORFORMULAANDFUNCTIONCOSTAGGREGATOR_AREA,
	MONITORFORMULAANDFUNCTIONCOSTAGGREGATOR_L2NORM,
	MONITORFORMULAANDFUNCTIONCOSTAGGREGATOR_PERCENTILE,
	MONITORFORMULAANDFUNCTIONCOSTAGGREGATOR_STDDEV,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MonitorFormulaAndFunctionCostAggregator) GetAllowedValues() []MonitorFormulaAndFunctionCostAggregator {
	return allowedMonitorFormulaAndFunctionCostAggregatorEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MonitorFormulaAndFunctionCostAggregator) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MonitorFormulaAndFunctionCostAggregator(value)
	return nil
}

// NewMonitorFormulaAndFunctionCostAggregatorFromValue returns a pointer to a valid MonitorFormulaAndFunctionCostAggregator
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMonitorFormulaAndFunctionCostAggregatorFromValue(v string) (*MonitorFormulaAndFunctionCostAggregator, error) {
	ev := MonitorFormulaAndFunctionCostAggregator(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MonitorFormulaAndFunctionCostAggregator: valid values are %v", v, allowedMonitorFormulaAndFunctionCostAggregatorEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MonitorFormulaAndFunctionCostAggregator) IsValid() bool {
	for _, existing := range allowedMonitorFormulaAndFunctionCostAggregatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MonitorFormulaAndFunctionCostAggregator value.
func (v MonitorFormulaAndFunctionCostAggregator) Ptr() *MonitorFormulaAndFunctionCostAggregator {
	return &v
}
