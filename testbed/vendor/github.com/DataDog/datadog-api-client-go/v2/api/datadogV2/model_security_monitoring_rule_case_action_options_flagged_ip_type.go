// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleCaseActionOptionsFlaggedIPType Used with the case action of type 'flag_ip'. The value specified in this field is applied as a flag to the IP addresses.
type SecurityMonitoringRuleCaseActionOptionsFlaggedIPType string

// List of SecurityMonitoringRuleCaseActionOptionsFlaggedIPType.
const (
	SECURITYMONITORINGRULECASEACTIONOPTIONSFLAGGEDIPTYPE_SUSPICIOUS SecurityMonitoringRuleCaseActionOptionsFlaggedIPType = "SUSPICIOUS"
	SECURITYMONITORINGRULECASEACTIONOPTIONSFLAGGEDIPTYPE_FLAGGED    SecurityMonitoringRuleCaseActionOptionsFlaggedIPType = "FLAGGED"
)

var allowedSecurityMonitoringRuleCaseActionOptionsFlaggedIPTypeEnumValues = []SecurityMonitoringRuleCaseActionOptionsFlaggedIPType{
	SECURITYMONITORINGRULECASEACTIONOPTIONSFLAGGEDIPTYPE_SUSPICIOUS,
	SECURITYMONITORINGRULECASEACTIONOPTIONSFLAGGEDIPTYPE_FLAGGED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringRuleCaseActionOptionsFlaggedIPType) GetAllowedValues() []SecurityMonitoringRuleCaseActionOptionsFlaggedIPType {
	return allowedSecurityMonitoringRuleCaseActionOptionsFlaggedIPTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringRuleCaseActionOptionsFlaggedIPType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringRuleCaseActionOptionsFlaggedIPType(value)
	return nil
}

// NewSecurityMonitoringRuleCaseActionOptionsFlaggedIPTypeFromValue returns a pointer to a valid SecurityMonitoringRuleCaseActionOptionsFlaggedIPType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringRuleCaseActionOptionsFlaggedIPTypeFromValue(v string) (*SecurityMonitoringRuleCaseActionOptionsFlaggedIPType, error) {
	ev := SecurityMonitoringRuleCaseActionOptionsFlaggedIPType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringRuleCaseActionOptionsFlaggedIPType: valid values are %v", v, allowedSecurityMonitoringRuleCaseActionOptionsFlaggedIPTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringRuleCaseActionOptionsFlaggedIPType) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringRuleCaseActionOptionsFlaggedIPTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringRuleCaseActionOptionsFlaggedIPType value.
func (v SecurityMonitoringRuleCaseActionOptionsFlaggedIPType) Ptr() *SecurityMonitoringRuleCaseActionOptionsFlaggedIPType {
	return &v
}
