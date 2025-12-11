// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleDetectionMethod The detection method.
type SecurityMonitoringRuleDetectionMethod string

// List of SecurityMonitoringRuleDetectionMethod.
const (
	SECURITYMONITORINGRULEDETECTIONMETHOD_THRESHOLD          SecurityMonitoringRuleDetectionMethod = "threshold"
	SECURITYMONITORINGRULEDETECTIONMETHOD_NEW_VALUE          SecurityMonitoringRuleDetectionMethod = "new_value"
	SECURITYMONITORINGRULEDETECTIONMETHOD_ANOMALY_DETECTION  SecurityMonitoringRuleDetectionMethod = "anomaly_detection"
	SECURITYMONITORINGRULEDETECTIONMETHOD_IMPOSSIBLE_TRAVEL  SecurityMonitoringRuleDetectionMethod = "impossible_travel"
	SECURITYMONITORINGRULEDETECTIONMETHOD_HARDCODED          SecurityMonitoringRuleDetectionMethod = "hardcoded"
	SECURITYMONITORINGRULEDETECTIONMETHOD_THIRD_PARTY        SecurityMonitoringRuleDetectionMethod = "third_party"
	SECURITYMONITORINGRULEDETECTIONMETHOD_ANOMALY_THRESHOLD  SecurityMonitoringRuleDetectionMethod = "anomaly_threshold"
	SECURITYMONITORINGRULEDETECTIONMETHOD_SEQUENCE_DETECTION SecurityMonitoringRuleDetectionMethod = "sequence_detection"
)

var allowedSecurityMonitoringRuleDetectionMethodEnumValues = []SecurityMonitoringRuleDetectionMethod{
	SECURITYMONITORINGRULEDETECTIONMETHOD_THRESHOLD,
	SECURITYMONITORINGRULEDETECTIONMETHOD_NEW_VALUE,
	SECURITYMONITORINGRULEDETECTIONMETHOD_ANOMALY_DETECTION,
	SECURITYMONITORINGRULEDETECTIONMETHOD_IMPOSSIBLE_TRAVEL,
	SECURITYMONITORINGRULEDETECTIONMETHOD_HARDCODED,
	SECURITYMONITORINGRULEDETECTIONMETHOD_THIRD_PARTY,
	SECURITYMONITORINGRULEDETECTIONMETHOD_ANOMALY_THRESHOLD,
	SECURITYMONITORINGRULEDETECTIONMETHOD_SEQUENCE_DETECTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringRuleDetectionMethod) GetAllowedValues() []SecurityMonitoringRuleDetectionMethod {
	return allowedSecurityMonitoringRuleDetectionMethodEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringRuleDetectionMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringRuleDetectionMethod(value)
	return nil
}

// NewSecurityMonitoringRuleDetectionMethodFromValue returns a pointer to a valid SecurityMonitoringRuleDetectionMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringRuleDetectionMethodFromValue(v string) (*SecurityMonitoringRuleDetectionMethod, error) {
	ev := SecurityMonitoringRuleDetectionMethod(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringRuleDetectionMethod: valid values are %v", v, allowedSecurityMonitoringRuleDetectionMethodEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringRuleDetectionMethod) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringRuleDetectionMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringRuleDetectionMethod value.
func (v SecurityMonitoringRuleDetectionMethod) Ptr() *SecurityMonitoringRuleDetectionMethod {
	return &v
}
