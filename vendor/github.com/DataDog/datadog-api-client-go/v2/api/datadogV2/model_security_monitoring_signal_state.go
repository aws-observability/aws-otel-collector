// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringSignalState The new triage state of the signal.
type SecurityMonitoringSignalState string

// List of SecurityMonitoringSignalState.
const (
	SECURITYMONITORINGSIGNALSTATE_OPEN         SecurityMonitoringSignalState = "open"
	SECURITYMONITORINGSIGNALSTATE_ARCHIVED     SecurityMonitoringSignalState = "archived"
	SECURITYMONITORINGSIGNALSTATE_UNDER_REVIEW SecurityMonitoringSignalState = "under_review"
)

var allowedSecurityMonitoringSignalStateEnumValues = []SecurityMonitoringSignalState{
	SECURITYMONITORINGSIGNALSTATE_OPEN,
	SECURITYMONITORINGSIGNALSTATE_ARCHIVED,
	SECURITYMONITORINGSIGNALSTATE_UNDER_REVIEW,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringSignalState) GetAllowedValues() []SecurityMonitoringSignalState {
	return allowedSecurityMonitoringSignalStateEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringSignalState) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringSignalState(value)
	return nil
}

// NewSecurityMonitoringSignalStateFromValue returns a pointer to a valid SecurityMonitoringSignalState
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringSignalStateFromValue(v string) (*SecurityMonitoringSignalState, error) {
	ev := SecurityMonitoringSignalState(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringSignalState: valid values are %v", v, allowedSecurityMonitoringSignalStateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringSignalState) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringSignalStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringSignalState value.
func (v SecurityMonitoringSignalState) Ptr() *SecurityMonitoringSignalState {
	return &v
}
