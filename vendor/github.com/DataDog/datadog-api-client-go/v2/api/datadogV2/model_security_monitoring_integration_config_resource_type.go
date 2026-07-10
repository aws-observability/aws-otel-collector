// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringIntegrationConfigResourceType The type of the resource. The value should always be `integration_config`.
type SecurityMonitoringIntegrationConfigResourceType string

// List of SecurityMonitoringIntegrationConfigResourceType.
const (
	SECURITYMONITORINGINTEGRATIONCONFIGRESOURCETYPE_INTEGRATION_CONFIG SecurityMonitoringIntegrationConfigResourceType = "integration_config"
)

var allowedSecurityMonitoringIntegrationConfigResourceTypeEnumValues = []SecurityMonitoringIntegrationConfigResourceType{
	SECURITYMONITORINGINTEGRATIONCONFIGRESOURCETYPE_INTEGRATION_CONFIG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringIntegrationConfigResourceType) GetAllowedValues() []SecurityMonitoringIntegrationConfigResourceType {
	return allowedSecurityMonitoringIntegrationConfigResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringIntegrationConfigResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringIntegrationConfigResourceType(value)
	return nil
}

// NewSecurityMonitoringIntegrationConfigResourceTypeFromValue returns a pointer to a valid SecurityMonitoringIntegrationConfigResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringIntegrationConfigResourceTypeFromValue(v string) (*SecurityMonitoringIntegrationConfigResourceType, error) {
	ev := SecurityMonitoringIntegrationConfigResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringIntegrationConfigResourceType: valid values are %v", v, allowedSecurityMonitoringIntegrationConfigResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringIntegrationConfigResourceType) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringIntegrationConfigResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringIntegrationConfigResourceType value.
func (v SecurityMonitoringIntegrationConfigResourceType) Ptr() *SecurityMonitoringIntegrationConfigResourceType {
	return &v
}
