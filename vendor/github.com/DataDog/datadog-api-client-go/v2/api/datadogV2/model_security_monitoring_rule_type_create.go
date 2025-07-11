// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleTypeCreate The rule type.
type SecurityMonitoringRuleTypeCreate string

// List of SecurityMonitoringRuleTypeCreate.
const (
	SECURITYMONITORINGRULETYPECREATE_API_SECURITY         SecurityMonitoringRuleTypeCreate = "api_security"
	SECURITYMONITORINGRULETYPECREATE_APPLICATION_SECURITY SecurityMonitoringRuleTypeCreate = "application_security"
	SECURITYMONITORINGRULETYPECREATE_LOG_DETECTION        SecurityMonitoringRuleTypeCreate = "log_detection"
	SECURITYMONITORINGRULETYPECREATE_WORKLOAD_SECURITY    SecurityMonitoringRuleTypeCreate = "workload_security"
)

var allowedSecurityMonitoringRuleTypeCreateEnumValues = []SecurityMonitoringRuleTypeCreate{
	SECURITYMONITORINGRULETYPECREATE_API_SECURITY,
	SECURITYMONITORINGRULETYPECREATE_APPLICATION_SECURITY,
	SECURITYMONITORINGRULETYPECREATE_LOG_DETECTION,
	SECURITYMONITORINGRULETYPECREATE_WORKLOAD_SECURITY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringRuleTypeCreate) GetAllowedValues() []SecurityMonitoringRuleTypeCreate {
	return allowedSecurityMonitoringRuleTypeCreateEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringRuleTypeCreate) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringRuleTypeCreate(value)
	return nil
}

// NewSecurityMonitoringRuleTypeCreateFromValue returns a pointer to a valid SecurityMonitoringRuleTypeCreate
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringRuleTypeCreateFromValue(v string) (*SecurityMonitoringRuleTypeCreate, error) {
	ev := SecurityMonitoringRuleTypeCreate(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringRuleTypeCreate: valid values are %v", v, allowedSecurityMonitoringRuleTypeCreateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringRuleTypeCreate) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringRuleTypeCreateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringRuleTypeCreate value.
func (v SecurityMonitoringRuleTypeCreate) Ptr() *SecurityMonitoringRuleTypeCreate {
	return &v
}
