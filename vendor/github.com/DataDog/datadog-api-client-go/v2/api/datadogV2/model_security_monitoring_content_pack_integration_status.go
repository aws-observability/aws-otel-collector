// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringContentPackIntegrationStatus The installation status of the related integration.
type SecurityMonitoringContentPackIntegrationStatus string

// List of SecurityMonitoringContentPackIntegrationStatus.
const (
	SECURITYMONITORINGCONTENTPACKINTEGRATIONSTATUS_INSTALLED           SecurityMonitoringContentPackIntegrationStatus = "installed"
	SECURITYMONITORINGCONTENTPACKINTEGRATIONSTATUS_AVAILABLE           SecurityMonitoringContentPackIntegrationStatus = "available"
	SECURITYMONITORINGCONTENTPACKINTEGRATIONSTATUS_PARTIALLY_INSTALLED SecurityMonitoringContentPackIntegrationStatus = "partially_installed"
	SECURITYMONITORINGCONTENTPACKINTEGRATIONSTATUS_DETECTED            SecurityMonitoringContentPackIntegrationStatus = "detected"
	SECURITYMONITORINGCONTENTPACKINTEGRATIONSTATUS_ERROR               SecurityMonitoringContentPackIntegrationStatus = "error"
)

var allowedSecurityMonitoringContentPackIntegrationStatusEnumValues = []SecurityMonitoringContentPackIntegrationStatus{
	SECURITYMONITORINGCONTENTPACKINTEGRATIONSTATUS_INSTALLED,
	SECURITYMONITORINGCONTENTPACKINTEGRATIONSTATUS_AVAILABLE,
	SECURITYMONITORINGCONTENTPACKINTEGRATIONSTATUS_PARTIALLY_INSTALLED,
	SECURITYMONITORINGCONTENTPACKINTEGRATIONSTATUS_DETECTED,
	SECURITYMONITORINGCONTENTPACKINTEGRATIONSTATUS_ERROR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringContentPackIntegrationStatus) GetAllowedValues() []SecurityMonitoringContentPackIntegrationStatus {
	return allowedSecurityMonitoringContentPackIntegrationStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringContentPackIntegrationStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringContentPackIntegrationStatus(value)
	return nil
}

// NewSecurityMonitoringContentPackIntegrationStatusFromValue returns a pointer to a valid SecurityMonitoringContentPackIntegrationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringContentPackIntegrationStatusFromValue(v string) (*SecurityMonitoringContentPackIntegrationStatus, error) {
	ev := SecurityMonitoringContentPackIntegrationStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringContentPackIntegrationStatus: valid values are %v", v, allowedSecurityMonitoringContentPackIntegrationStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringContentPackIntegrationStatus) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringContentPackIntegrationStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringContentPackIntegrationStatus value.
func (v SecurityMonitoringContentPackIntegrationStatus) Ptr() *SecurityMonitoringContentPackIntegrationStatus {
	return &v
}
