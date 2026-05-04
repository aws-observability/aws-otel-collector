// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorFormulaAndFunctionAggregateFilteredDataSource Data source for aggregate filtered queries.
type MonitorFormulaAndFunctionAggregateFilteredDataSource string

// List of MonitorFormulaAndFunctionAggregateFilteredDataSource.
const (
	MONITORFORMULAANDFUNCTIONAGGREGATEFILTEREDDATASOURCE_AGGREGATE_FILTERED_QUERY MonitorFormulaAndFunctionAggregateFilteredDataSource = "aggregate_filtered_query"
)

var allowedMonitorFormulaAndFunctionAggregateFilteredDataSourceEnumValues = []MonitorFormulaAndFunctionAggregateFilteredDataSource{
	MONITORFORMULAANDFUNCTIONAGGREGATEFILTEREDDATASOURCE_AGGREGATE_FILTERED_QUERY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MonitorFormulaAndFunctionAggregateFilteredDataSource) GetAllowedValues() []MonitorFormulaAndFunctionAggregateFilteredDataSource {
	return allowedMonitorFormulaAndFunctionAggregateFilteredDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MonitorFormulaAndFunctionAggregateFilteredDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MonitorFormulaAndFunctionAggregateFilteredDataSource(value)
	return nil
}

// NewMonitorFormulaAndFunctionAggregateFilteredDataSourceFromValue returns a pointer to a valid MonitorFormulaAndFunctionAggregateFilteredDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMonitorFormulaAndFunctionAggregateFilteredDataSourceFromValue(v string) (*MonitorFormulaAndFunctionAggregateFilteredDataSource, error) {
	ev := MonitorFormulaAndFunctionAggregateFilteredDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MonitorFormulaAndFunctionAggregateFilteredDataSource: valid values are %v", v, allowedMonitorFormulaAndFunctionAggregateFilteredDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MonitorFormulaAndFunctionAggregateFilteredDataSource) IsValid() bool {
	for _, existing := range allowedMonitorFormulaAndFunctionAggregateFilteredDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MonitorFormulaAndFunctionAggregateFilteredDataSource value.
func (v MonitorFormulaAndFunctionAggregateFilteredDataSource) Ptr() *MonitorFormulaAndFunctionAggregateFilteredDataSource {
	return &v
}
