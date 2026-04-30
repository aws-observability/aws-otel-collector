// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorFormulaAndFunctionDataQualityDataSource Data source for data quality queries.
type MonitorFormulaAndFunctionDataQualityDataSource string

// List of MonitorFormulaAndFunctionDataQualityDataSource.
const (
	MONITORFORMULAANDFUNCTIONDATAQUALITYDATASOURCE_DATA_QUALITY_METRICS MonitorFormulaAndFunctionDataQualityDataSource = "data_quality_metrics"
)

var allowedMonitorFormulaAndFunctionDataQualityDataSourceEnumValues = []MonitorFormulaAndFunctionDataQualityDataSource{
	MONITORFORMULAANDFUNCTIONDATAQUALITYDATASOURCE_DATA_QUALITY_METRICS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MonitorFormulaAndFunctionDataQualityDataSource) GetAllowedValues() []MonitorFormulaAndFunctionDataQualityDataSource {
	return allowedMonitorFormulaAndFunctionDataQualityDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MonitorFormulaAndFunctionDataQualityDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MonitorFormulaAndFunctionDataQualityDataSource(value)
	return nil
}

// NewMonitorFormulaAndFunctionDataQualityDataSourceFromValue returns a pointer to a valid MonitorFormulaAndFunctionDataQualityDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMonitorFormulaAndFunctionDataQualityDataSourceFromValue(v string) (*MonitorFormulaAndFunctionDataQualityDataSource, error) {
	ev := MonitorFormulaAndFunctionDataQualityDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MonitorFormulaAndFunctionDataQualityDataSource: valid values are %v", v, allowedMonitorFormulaAndFunctionDataQualityDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MonitorFormulaAndFunctionDataQualityDataSource) IsValid() bool {
	for _, existing := range allowedMonitorFormulaAndFunctionDataQualityDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MonitorFormulaAndFunctionDataQualityDataSource value.
func (v MonitorFormulaAndFunctionDataQualityDataSource) Ptr() *MonitorFormulaAndFunctionDataQualityDataSource {
	return &v
}
