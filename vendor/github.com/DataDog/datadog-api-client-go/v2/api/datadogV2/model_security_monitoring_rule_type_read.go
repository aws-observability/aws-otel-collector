// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleTypeRead The rule type.
type SecurityMonitoringRuleTypeRead string

// List of SecurityMonitoringRuleTypeRead.
const (
	SECURITYMONITORINGRULETYPEREAD_LOG_DETECTION                SecurityMonitoringRuleTypeRead = "log_detection"
	SECURITYMONITORINGRULETYPEREAD_INFRASTRUCTURE_CONFIGURATION SecurityMonitoringRuleTypeRead = "infrastructure_configuration"
	SECURITYMONITORINGRULETYPEREAD_WORKLOAD_SECURITY            SecurityMonitoringRuleTypeRead = "workload_security"
	SECURITYMONITORINGRULETYPEREAD_CLOUD_CONFIGURATION          SecurityMonitoringRuleTypeRead = "cloud_configuration"
	SECURITYMONITORINGRULETYPEREAD_APPLICATION_SECURITY         SecurityMonitoringRuleTypeRead = "application_security"
	SECURITYMONITORINGRULETYPEREAD_API_SECURITY                 SecurityMonitoringRuleTypeRead = "api_security"
)

var allowedSecurityMonitoringRuleTypeReadEnumValues = []SecurityMonitoringRuleTypeRead{
	SECURITYMONITORINGRULETYPEREAD_LOG_DETECTION,
	SECURITYMONITORINGRULETYPEREAD_INFRASTRUCTURE_CONFIGURATION,
	SECURITYMONITORINGRULETYPEREAD_WORKLOAD_SECURITY,
	SECURITYMONITORINGRULETYPEREAD_CLOUD_CONFIGURATION,
	SECURITYMONITORINGRULETYPEREAD_APPLICATION_SECURITY,
	SECURITYMONITORINGRULETYPEREAD_API_SECURITY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringRuleTypeRead) GetAllowedValues() []SecurityMonitoringRuleTypeRead {
	return allowedSecurityMonitoringRuleTypeReadEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringRuleTypeRead) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringRuleTypeRead(value)
	return nil
}

// NewSecurityMonitoringRuleTypeReadFromValue returns a pointer to a valid SecurityMonitoringRuleTypeRead
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringRuleTypeReadFromValue(v string) (*SecurityMonitoringRuleTypeRead, error) {
	ev := SecurityMonitoringRuleTypeRead(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringRuleTypeRead: valid values are %v", v, allowedSecurityMonitoringRuleTypeReadEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringRuleTypeRead) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringRuleTypeReadEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringRuleTypeRead value.
func (v SecurityMonitoringRuleTypeRead) Ptr() *SecurityMonitoringRuleTypeRead {
	return &v
}
