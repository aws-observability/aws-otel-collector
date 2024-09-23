// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorFormulaAndFunctionEventsDataSource Data source for event platform-based queries.
type MonitorFormulaAndFunctionEventsDataSource string

// List of MonitorFormulaAndFunctionEventsDataSource.
const (
	MONITORFORMULAANDFUNCTIONEVENTSDATASOURCE_RUM              MonitorFormulaAndFunctionEventsDataSource = "rum"
	MONITORFORMULAANDFUNCTIONEVENTSDATASOURCE_CI_PIPELINES     MonitorFormulaAndFunctionEventsDataSource = "ci_pipelines"
	MONITORFORMULAANDFUNCTIONEVENTSDATASOURCE_CI_TESTS         MonitorFormulaAndFunctionEventsDataSource = "ci_tests"
	MONITORFORMULAANDFUNCTIONEVENTSDATASOURCE_AUDIT            MonitorFormulaAndFunctionEventsDataSource = "audit"
	MONITORFORMULAANDFUNCTIONEVENTSDATASOURCE_EVENTS           MonitorFormulaAndFunctionEventsDataSource = "events"
	MONITORFORMULAANDFUNCTIONEVENTSDATASOURCE_LOGS             MonitorFormulaAndFunctionEventsDataSource = "logs"
	MONITORFORMULAANDFUNCTIONEVENTSDATASOURCE_SPANS            MonitorFormulaAndFunctionEventsDataSource = "spans"
	MONITORFORMULAANDFUNCTIONEVENTSDATASOURCE_DATABASE_QUERIES MonitorFormulaAndFunctionEventsDataSource = "database_queries"
	MONITORFORMULAANDFUNCTIONEVENTSDATASOURCE_NETWORK          MonitorFormulaAndFunctionEventsDataSource = "network"
)

var allowedMonitorFormulaAndFunctionEventsDataSourceEnumValues = []MonitorFormulaAndFunctionEventsDataSource{
	MONITORFORMULAANDFUNCTIONEVENTSDATASOURCE_RUM,
	MONITORFORMULAANDFUNCTIONEVENTSDATASOURCE_CI_PIPELINES,
	MONITORFORMULAANDFUNCTIONEVENTSDATASOURCE_CI_TESTS,
	MONITORFORMULAANDFUNCTIONEVENTSDATASOURCE_AUDIT,
	MONITORFORMULAANDFUNCTIONEVENTSDATASOURCE_EVENTS,
	MONITORFORMULAANDFUNCTIONEVENTSDATASOURCE_LOGS,
	MONITORFORMULAANDFUNCTIONEVENTSDATASOURCE_SPANS,
	MONITORFORMULAANDFUNCTIONEVENTSDATASOURCE_DATABASE_QUERIES,
	MONITORFORMULAANDFUNCTIONEVENTSDATASOURCE_NETWORK,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MonitorFormulaAndFunctionEventsDataSource) GetAllowedValues() []MonitorFormulaAndFunctionEventsDataSource {
	return allowedMonitorFormulaAndFunctionEventsDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MonitorFormulaAndFunctionEventsDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MonitorFormulaAndFunctionEventsDataSource(value)
	return nil
}

// NewMonitorFormulaAndFunctionEventsDataSourceFromValue returns a pointer to a valid MonitorFormulaAndFunctionEventsDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMonitorFormulaAndFunctionEventsDataSourceFromValue(v string) (*MonitorFormulaAndFunctionEventsDataSource, error) {
	ev := MonitorFormulaAndFunctionEventsDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MonitorFormulaAndFunctionEventsDataSource: valid values are %v", v, allowedMonitorFormulaAndFunctionEventsDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MonitorFormulaAndFunctionEventsDataSource) IsValid() bool {
	for _, existing := range allowedMonitorFormulaAndFunctionEventsDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MonitorFormulaAndFunctionEventsDataSource value.
func (v MonitorFormulaAndFunctionEventsDataSource) Ptr() *MonitorFormulaAndFunctionEventsDataSource {
	return &v
}
