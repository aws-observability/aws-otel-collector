// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringSignalsSort The sort parameters used for querying security signals.
type SecurityMonitoringSignalsSort string

// List of SecurityMonitoringSignalsSort.
const (
	SECURITYMONITORINGSIGNALSSORT_TIMESTAMP_ASCENDING  SecurityMonitoringSignalsSort = "timestamp"
	SECURITYMONITORINGSIGNALSSORT_TIMESTAMP_DESCENDING SecurityMonitoringSignalsSort = "-timestamp"
)

var allowedSecurityMonitoringSignalsSortEnumValues = []SecurityMonitoringSignalsSort{
	SECURITYMONITORINGSIGNALSSORT_TIMESTAMP_ASCENDING,
	SECURITYMONITORINGSIGNALSSORT_TIMESTAMP_DESCENDING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringSignalsSort) GetAllowedValues() []SecurityMonitoringSignalsSort {
	return allowedSecurityMonitoringSignalsSortEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringSignalsSort) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringSignalsSort(value)
	return nil
}

// NewSecurityMonitoringSignalsSortFromValue returns a pointer to a valid SecurityMonitoringSignalsSort
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringSignalsSortFromValue(v string) (*SecurityMonitoringSignalsSort, error) {
	ev := SecurityMonitoringSignalsSort(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringSignalsSort: valid values are %v", v, allowedSecurityMonitoringSignalsSortEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringSignalsSort) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringSignalsSortEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringSignalsSort value.
func (v SecurityMonitoringSignalsSort) Ptr() *SecurityMonitoringSignalsSort {
	return &v
}
