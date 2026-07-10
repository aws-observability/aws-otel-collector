// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringIntegrationConfigState The state of the credentials configured on the entity context sync.
type SecurityMonitoringIntegrationConfigState string

// List of SecurityMonitoringIntegrationConfigState.
const (
	SECURITYMONITORINGINTEGRATIONCONFIGSTATE_VALID        SecurityMonitoringIntegrationConfigState = "valid"
	SECURITYMONITORINGINTEGRATIONCONFIGSTATE_INVALID      SecurityMonitoringIntegrationConfigState = "invalid"
	SECURITYMONITORINGINTEGRATIONCONFIGSTATE_INITIALIZING SecurityMonitoringIntegrationConfigState = "initializing"
)

var allowedSecurityMonitoringIntegrationConfigStateEnumValues = []SecurityMonitoringIntegrationConfigState{
	SECURITYMONITORINGINTEGRATIONCONFIGSTATE_VALID,
	SECURITYMONITORINGINTEGRATIONCONFIGSTATE_INVALID,
	SECURITYMONITORINGINTEGRATIONCONFIGSTATE_INITIALIZING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringIntegrationConfigState) GetAllowedValues() []SecurityMonitoringIntegrationConfigState {
	return allowedSecurityMonitoringIntegrationConfigStateEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringIntegrationConfigState) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringIntegrationConfigState(value)
	return nil
}

// NewSecurityMonitoringIntegrationConfigStateFromValue returns a pointer to a valid SecurityMonitoringIntegrationConfigState
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringIntegrationConfigStateFromValue(v string) (*SecurityMonitoringIntegrationConfigState, error) {
	ev := SecurityMonitoringIntegrationConfigState(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringIntegrationConfigState: valid values are %v", v, allowedSecurityMonitoringIntegrationConfigStateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringIntegrationConfigState) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringIntegrationConfigStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringIntegrationConfigState value.
func (v SecurityMonitoringIntegrationConfigState) Ptr() *SecurityMonitoringIntegrationConfigState {
	return &v
}
