// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaAndFunctionEventsDataSource Data source for event platform-based queries.
type FormulaAndFunctionEventsDataSource string

// List of FormulaAndFunctionEventsDataSource.
const (
	FORMULAANDFUNCTIONEVENTSDATASOURCE_LOGS               FormulaAndFunctionEventsDataSource = "logs"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_SPANS              FormulaAndFunctionEventsDataSource = "spans"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_NETWORK            FormulaAndFunctionEventsDataSource = "network"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_RUM                FormulaAndFunctionEventsDataSource = "rum"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_SECURITY_SIGNALS   FormulaAndFunctionEventsDataSource = "security_signals"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_PROFILES           FormulaAndFunctionEventsDataSource = "profiles"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_AUDIT              FormulaAndFunctionEventsDataSource = "audit"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_EVENTS             FormulaAndFunctionEventsDataSource = "events"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_CI_TESTS           FormulaAndFunctionEventsDataSource = "ci_tests"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_CI_PIPELINES       FormulaAndFunctionEventsDataSource = "ci_pipelines"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_INCIDENT_ANALYTICS FormulaAndFunctionEventsDataSource = "incident_analytics"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_PRODUCT_ANALYTICS  FormulaAndFunctionEventsDataSource = "product_analytics"
	FORMULAANDFUNCTIONEVENTSDATASOURCE_ON_CALL_EVENTS     FormulaAndFunctionEventsDataSource = "on_call_events"
)

var allowedFormulaAndFunctionEventsDataSourceEnumValues = []FormulaAndFunctionEventsDataSource{
	FORMULAANDFUNCTIONEVENTSDATASOURCE_LOGS,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_SPANS,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_NETWORK,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_RUM,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_SECURITY_SIGNALS,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_PROFILES,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_AUDIT,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_EVENTS,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_CI_TESTS,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_CI_PIPELINES,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_INCIDENT_ANALYTICS,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_PRODUCT_ANALYTICS,
	FORMULAANDFUNCTIONEVENTSDATASOURCE_ON_CALL_EVENTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaAndFunctionEventsDataSource) GetAllowedValues() []FormulaAndFunctionEventsDataSource {
	return allowedFormulaAndFunctionEventsDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaAndFunctionEventsDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaAndFunctionEventsDataSource(value)
	return nil
}

// NewFormulaAndFunctionEventsDataSourceFromValue returns a pointer to a valid FormulaAndFunctionEventsDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaAndFunctionEventsDataSourceFromValue(v string) (*FormulaAndFunctionEventsDataSource, error) {
	ev := FormulaAndFunctionEventsDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaAndFunctionEventsDataSource: valid values are %v", v, allowedFormulaAndFunctionEventsDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaAndFunctionEventsDataSource) IsValid() bool {
	for _, existing := range allowedFormulaAndFunctionEventsDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaAndFunctionEventsDataSource value.
func (v FormulaAndFunctionEventsDataSource) Ptr() *FormulaAndFunctionEventsDataSource {
	return &v
}
