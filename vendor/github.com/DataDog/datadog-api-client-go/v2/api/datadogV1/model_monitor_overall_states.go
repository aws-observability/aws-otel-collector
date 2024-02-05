// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorOverallStates The different states your monitor can be in.
type MonitorOverallStates string

// List of MonitorOverallStates.
const (
	MONITOROVERALLSTATES_ALERT   MonitorOverallStates = "Alert"
	MONITOROVERALLSTATES_IGNORED MonitorOverallStates = "Ignored"
	MONITOROVERALLSTATES_NO_DATA MonitorOverallStates = "No Data"
	MONITOROVERALLSTATES_OK      MonitorOverallStates = "OK"
	MONITOROVERALLSTATES_SKIPPED MonitorOverallStates = "Skipped"
	MONITOROVERALLSTATES_UNKNOWN MonitorOverallStates = "Unknown"
	MONITOROVERALLSTATES_WARN    MonitorOverallStates = "Warn"
)

var allowedMonitorOverallStatesEnumValues = []MonitorOverallStates{
	MONITOROVERALLSTATES_ALERT,
	MONITOROVERALLSTATES_IGNORED,
	MONITOROVERALLSTATES_NO_DATA,
	MONITOROVERALLSTATES_OK,
	MONITOROVERALLSTATES_SKIPPED,
	MONITOROVERALLSTATES_UNKNOWN,
	MONITOROVERALLSTATES_WARN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MonitorOverallStates) GetAllowedValues() []MonitorOverallStates {
	return allowedMonitorOverallStatesEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MonitorOverallStates) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MonitorOverallStates(value)
	return nil
}

// NewMonitorOverallStatesFromValue returns a pointer to a valid MonitorOverallStates
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMonitorOverallStatesFromValue(v string) (*MonitorOverallStates, error) {
	ev := MonitorOverallStates(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MonitorOverallStates: valid values are %v", v, allowedMonitorOverallStatesEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MonitorOverallStates) IsValid() bool {
	for _, existing := range allowedMonitorOverallStatesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MonitorOverallStates value.
func (v MonitorOverallStates) Ptr() *MonitorOverallStates {
	return &v
}
