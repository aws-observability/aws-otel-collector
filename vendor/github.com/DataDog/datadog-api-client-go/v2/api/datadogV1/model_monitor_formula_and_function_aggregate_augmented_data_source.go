// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorFormulaAndFunctionAggregateAugmentedDataSource Data source for aggregate augmented queries.
type MonitorFormulaAndFunctionAggregateAugmentedDataSource string

// List of MonitorFormulaAndFunctionAggregateAugmentedDataSource.
const (
	MONITORFORMULAANDFUNCTIONAGGREGATEAUGMENTEDDATASOURCE_AGGREGATE_AUGMENTED_QUERY MonitorFormulaAndFunctionAggregateAugmentedDataSource = "aggregate_augmented_query"
)

var allowedMonitorFormulaAndFunctionAggregateAugmentedDataSourceEnumValues = []MonitorFormulaAndFunctionAggregateAugmentedDataSource{
	MONITORFORMULAANDFUNCTIONAGGREGATEAUGMENTEDDATASOURCE_AGGREGATE_AUGMENTED_QUERY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MonitorFormulaAndFunctionAggregateAugmentedDataSource) GetAllowedValues() []MonitorFormulaAndFunctionAggregateAugmentedDataSource {
	return allowedMonitorFormulaAndFunctionAggregateAugmentedDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MonitorFormulaAndFunctionAggregateAugmentedDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MonitorFormulaAndFunctionAggregateAugmentedDataSource(value)
	return nil
}

// NewMonitorFormulaAndFunctionAggregateAugmentedDataSourceFromValue returns a pointer to a valid MonitorFormulaAndFunctionAggregateAugmentedDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMonitorFormulaAndFunctionAggregateAugmentedDataSourceFromValue(v string) (*MonitorFormulaAndFunctionAggregateAugmentedDataSource, error) {
	ev := MonitorFormulaAndFunctionAggregateAugmentedDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MonitorFormulaAndFunctionAggregateAugmentedDataSource: valid values are %v", v, allowedMonitorFormulaAndFunctionAggregateAugmentedDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MonitorFormulaAndFunctionAggregateAugmentedDataSource) IsValid() bool {
	for _, existing := range allowedMonitorFormulaAndFunctionAggregateAugmentedDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MonitorFormulaAndFunctionAggregateAugmentedDataSource value.
func (v MonitorFormulaAndFunctionAggregateAugmentedDataSource) Ptr() *MonitorFormulaAndFunctionAggregateAugmentedDataSource {
	return &v
}
