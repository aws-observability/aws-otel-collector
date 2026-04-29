// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorFormulaAndFunctionReferenceTableDataSource Data source for reference table queries.
type MonitorFormulaAndFunctionReferenceTableDataSource string

// List of MonitorFormulaAndFunctionReferenceTableDataSource.
const (
	MONITORFORMULAANDFUNCTIONREFERENCETABLEDATASOURCE_REFERENCE_TABLE MonitorFormulaAndFunctionReferenceTableDataSource = "reference_table"
)

var allowedMonitorFormulaAndFunctionReferenceTableDataSourceEnumValues = []MonitorFormulaAndFunctionReferenceTableDataSource{
	MONITORFORMULAANDFUNCTIONREFERENCETABLEDATASOURCE_REFERENCE_TABLE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MonitorFormulaAndFunctionReferenceTableDataSource) GetAllowedValues() []MonitorFormulaAndFunctionReferenceTableDataSource {
	return allowedMonitorFormulaAndFunctionReferenceTableDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MonitorFormulaAndFunctionReferenceTableDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MonitorFormulaAndFunctionReferenceTableDataSource(value)
	return nil
}

// NewMonitorFormulaAndFunctionReferenceTableDataSourceFromValue returns a pointer to a valid MonitorFormulaAndFunctionReferenceTableDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMonitorFormulaAndFunctionReferenceTableDataSourceFromValue(v string) (*MonitorFormulaAndFunctionReferenceTableDataSource, error) {
	ev := MonitorFormulaAndFunctionReferenceTableDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MonitorFormulaAndFunctionReferenceTableDataSource: valid values are %v", v, allowedMonitorFormulaAndFunctionReferenceTableDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MonitorFormulaAndFunctionReferenceTableDataSource) IsValid() bool {
	for _, existing := range allowedMonitorFormulaAndFunctionReferenceTableDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MonitorFormulaAndFunctionReferenceTableDataSource value.
func (v MonitorFormulaAndFunctionReferenceTableDataSource) Ptr() *MonitorFormulaAndFunctionReferenceTableDataSource {
	return &v
}
