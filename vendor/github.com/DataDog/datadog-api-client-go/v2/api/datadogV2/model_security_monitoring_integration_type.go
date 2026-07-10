// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringIntegrationType The type of external source that provides entities to Cloud SIEM.
type SecurityMonitoringIntegrationType string

// List of SecurityMonitoringIntegrationType.
const (
	SECURITYMONITORINGINTEGRATIONTYPE_GOOGLE_WORKSPACE SecurityMonitoringIntegrationType = "GOOGLE_WORKSPACE"
	SECURITYMONITORINGINTEGRATIONTYPE_OKTA             SecurityMonitoringIntegrationType = "OKTA"
	SECURITYMONITORINGINTEGRATIONTYPE_ENTRA_ID         SecurityMonitoringIntegrationType = "ENTRA_ID"
)

var allowedSecurityMonitoringIntegrationTypeEnumValues = []SecurityMonitoringIntegrationType{
	SECURITYMONITORINGINTEGRATIONTYPE_GOOGLE_WORKSPACE,
	SECURITYMONITORINGINTEGRATIONTYPE_OKTA,
	SECURITYMONITORINGINTEGRATIONTYPE_ENTRA_ID,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringIntegrationType) GetAllowedValues() []SecurityMonitoringIntegrationType {
	return allowedSecurityMonitoringIntegrationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringIntegrationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringIntegrationType(value)
	return nil
}

// NewSecurityMonitoringIntegrationTypeFromValue returns a pointer to a valid SecurityMonitoringIntegrationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringIntegrationTypeFromValue(v string) (*SecurityMonitoringIntegrationType, error) {
	ev := SecurityMonitoringIntegrationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringIntegrationType: valid values are %v", v, allowedSecurityMonitoringIntegrationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringIntegrationType) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringIntegrationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringIntegrationType value.
func (v SecurityMonitoringIntegrationType) Ptr() *SecurityMonitoringIntegrationType {
	return &v
}
