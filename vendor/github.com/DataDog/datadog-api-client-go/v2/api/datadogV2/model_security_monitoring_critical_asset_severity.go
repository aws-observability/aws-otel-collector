// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringCriticalAssetSeverity Severity associated with this critical asset. Either an explicit severity can be set, or the severity can be increased or decreased, or the severity can be left unchanged (no-op).
type SecurityMonitoringCriticalAssetSeverity string

// List of SecurityMonitoringCriticalAssetSeverity.
const (
	SECURITYMONITORINGCRITICALASSETSEVERITY_INFO     SecurityMonitoringCriticalAssetSeverity = "info"
	SECURITYMONITORINGCRITICALASSETSEVERITY_LOW      SecurityMonitoringCriticalAssetSeverity = "low"
	SECURITYMONITORINGCRITICALASSETSEVERITY_MEDIUM   SecurityMonitoringCriticalAssetSeverity = "medium"
	SECURITYMONITORINGCRITICALASSETSEVERITY_HIGH     SecurityMonitoringCriticalAssetSeverity = "high"
	SECURITYMONITORINGCRITICALASSETSEVERITY_CRITICAL SecurityMonitoringCriticalAssetSeverity = "critical"
	SECURITYMONITORINGCRITICALASSETSEVERITY_INCREASE SecurityMonitoringCriticalAssetSeverity = "increase"
	SECURITYMONITORINGCRITICALASSETSEVERITY_DECREASE SecurityMonitoringCriticalAssetSeverity = "decrease"
	SECURITYMONITORINGCRITICALASSETSEVERITY_NO_OP    SecurityMonitoringCriticalAssetSeverity = "no-op"
)

var allowedSecurityMonitoringCriticalAssetSeverityEnumValues = []SecurityMonitoringCriticalAssetSeverity{
	SECURITYMONITORINGCRITICALASSETSEVERITY_INFO,
	SECURITYMONITORINGCRITICALASSETSEVERITY_LOW,
	SECURITYMONITORINGCRITICALASSETSEVERITY_MEDIUM,
	SECURITYMONITORINGCRITICALASSETSEVERITY_HIGH,
	SECURITYMONITORINGCRITICALASSETSEVERITY_CRITICAL,
	SECURITYMONITORINGCRITICALASSETSEVERITY_INCREASE,
	SECURITYMONITORINGCRITICALASSETSEVERITY_DECREASE,
	SECURITYMONITORINGCRITICALASSETSEVERITY_NO_OP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringCriticalAssetSeverity) GetAllowedValues() []SecurityMonitoringCriticalAssetSeverity {
	return allowedSecurityMonitoringCriticalAssetSeverityEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringCriticalAssetSeverity) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringCriticalAssetSeverity(value)
	return nil
}

// NewSecurityMonitoringCriticalAssetSeverityFromValue returns a pointer to a valid SecurityMonitoringCriticalAssetSeverity
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringCriticalAssetSeverityFromValue(v string) (*SecurityMonitoringCriticalAssetSeverity, error) {
	ev := SecurityMonitoringCriticalAssetSeverity(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringCriticalAssetSeverity: valid values are %v", v, allowedSecurityMonitoringCriticalAssetSeverityEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringCriticalAssetSeverity) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringCriticalAssetSeverityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringCriticalAssetSeverity value.
func (v SecurityMonitoringCriticalAssetSeverity) Ptr() *SecurityMonitoringCriticalAssetSeverity {
	return &v
}
