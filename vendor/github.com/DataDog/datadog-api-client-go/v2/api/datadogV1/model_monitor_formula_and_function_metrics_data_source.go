// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorFormulaAndFunctionMetricsDataSource Data source for metrics queries.
type MonitorFormulaAndFunctionMetricsDataSource string

// List of MonitorFormulaAndFunctionMetricsDataSource.
const (
	MONITORFORMULAANDFUNCTIONMETRICSDATASOURCE_METRICS       MonitorFormulaAndFunctionMetricsDataSource = "metrics"
	MONITORFORMULAANDFUNCTIONMETRICSDATASOURCE_CLOUD_COST    MonitorFormulaAndFunctionMetricsDataSource = "cloud_cost"
	MONITORFORMULAANDFUNCTIONMETRICSDATASOURCE_DATADOG_USAGE MonitorFormulaAndFunctionMetricsDataSource = "datadog_usage"
)

var allowedMonitorFormulaAndFunctionMetricsDataSourceEnumValues = []MonitorFormulaAndFunctionMetricsDataSource{
	MONITORFORMULAANDFUNCTIONMETRICSDATASOURCE_METRICS,
	MONITORFORMULAANDFUNCTIONMETRICSDATASOURCE_CLOUD_COST,
	MONITORFORMULAANDFUNCTIONMETRICSDATASOURCE_DATADOG_USAGE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MonitorFormulaAndFunctionMetricsDataSource) GetAllowedValues() []MonitorFormulaAndFunctionMetricsDataSource {
	return allowedMonitorFormulaAndFunctionMetricsDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MonitorFormulaAndFunctionMetricsDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MonitorFormulaAndFunctionMetricsDataSource(value)
	return nil
}

// NewMonitorFormulaAndFunctionMetricsDataSourceFromValue returns a pointer to a valid MonitorFormulaAndFunctionMetricsDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMonitorFormulaAndFunctionMetricsDataSourceFromValue(v string) (*MonitorFormulaAndFunctionMetricsDataSource, error) {
	ev := MonitorFormulaAndFunctionMetricsDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MonitorFormulaAndFunctionMetricsDataSource: valid values are %v", v, allowedMonitorFormulaAndFunctionMetricsDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MonitorFormulaAndFunctionMetricsDataSource) IsValid() bool {
	for _, existing := range allowedMonitorFormulaAndFunctionMetricsDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MonitorFormulaAndFunctionMetricsDataSource value.
func (v MonitorFormulaAndFunctionMetricsDataSource) Ptr() *MonitorFormulaAndFunctionMetricsDataSource {
	return &v
}
