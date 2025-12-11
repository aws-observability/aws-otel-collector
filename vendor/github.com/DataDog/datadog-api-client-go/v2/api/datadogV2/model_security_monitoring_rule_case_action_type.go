// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleCaseActionType The action type.
type SecurityMonitoringRuleCaseActionType string

// List of SecurityMonitoringRuleCaseActionType.
const (
	SECURITYMONITORINGRULECASEACTIONTYPE_BLOCK_IP      SecurityMonitoringRuleCaseActionType = "block_ip"
	SECURITYMONITORINGRULECASEACTIONTYPE_BLOCK_USER    SecurityMonitoringRuleCaseActionType = "block_user"
	SECURITYMONITORINGRULECASEACTIONTYPE_USER_BEHAVIOR SecurityMonitoringRuleCaseActionType = "user_behavior"
	SECURITYMONITORINGRULECASEACTIONTYPE_FLAG_IP       SecurityMonitoringRuleCaseActionType = "flag_ip"
)

var allowedSecurityMonitoringRuleCaseActionTypeEnumValues = []SecurityMonitoringRuleCaseActionType{
	SECURITYMONITORINGRULECASEACTIONTYPE_BLOCK_IP,
	SECURITYMONITORINGRULECASEACTIONTYPE_BLOCK_USER,
	SECURITYMONITORINGRULECASEACTIONTYPE_USER_BEHAVIOR,
	SECURITYMONITORINGRULECASEACTIONTYPE_FLAG_IP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringRuleCaseActionType) GetAllowedValues() []SecurityMonitoringRuleCaseActionType {
	return allowedSecurityMonitoringRuleCaseActionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringRuleCaseActionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringRuleCaseActionType(value)
	return nil
}

// NewSecurityMonitoringRuleCaseActionTypeFromValue returns a pointer to a valid SecurityMonitoringRuleCaseActionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringRuleCaseActionTypeFromValue(v string) (*SecurityMonitoringRuleCaseActionType, error) {
	ev := SecurityMonitoringRuleCaseActionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringRuleCaseActionType: valid values are %v", v, allowedSecurityMonitoringRuleCaseActionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringRuleCaseActionType) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringRuleCaseActionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringRuleCaseActionType value.
func (v SecurityMonitoringRuleCaseActionType) Ptr() *SecurityMonitoringRuleCaseActionType {
	return &v
}
