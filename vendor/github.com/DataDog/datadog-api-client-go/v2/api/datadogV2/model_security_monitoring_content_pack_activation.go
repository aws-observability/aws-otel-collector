// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringContentPackActivation The activation status of a content pack.
type SecurityMonitoringContentPackActivation string

// List of SecurityMonitoringContentPackActivation.
const (
	SECURITYMONITORINGCONTENTPACKACTIVATION_NEVER_ACTIVATED SecurityMonitoringContentPackActivation = "never_activated"
	SECURITYMONITORINGCONTENTPACKACTIVATION_ACTIVATED       SecurityMonitoringContentPackActivation = "activated"
	SECURITYMONITORINGCONTENTPACKACTIVATION_DEACTIVATED     SecurityMonitoringContentPackActivation = "deactivated"
)

var allowedSecurityMonitoringContentPackActivationEnumValues = []SecurityMonitoringContentPackActivation{
	SECURITYMONITORINGCONTENTPACKACTIVATION_NEVER_ACTIVATED,
	SECURITYMONITORINGCONTENTPACKACTIVATION_ACTIVATED,
	SECURITYMONITORINGCONTENTPACKACTIVATION_DEACTIVATED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringContentPackActivation) GetAllowedValues() []SecurityMonitoringContentPackActivation {
	return allowedSecurityMonitoringContentPackActivationEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringContentPackActivation) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringContentPackActivation(value)
	return nil
}

// NewSecurityMonitoringContentPackActivationFromValue returns a pointer to a valid SecurityMonitoringContentPackActivation
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringContentPackActivationFromValue(v string) (*SecurityMonitoringContentPackActivation, error) {
	ev := SecurityMonitoringContentPackActivation(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringContentPackActivation: valid values are %v", v, allowedSecurityMonitoringContentPackActivationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringContentPackActivation) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringContentPackActivationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringContentPackActivation value.
func (v SecurityMonitoringContentPackActivation) Ptr() *SecurityMonitoringContentPackActivation {
	return &v
}
