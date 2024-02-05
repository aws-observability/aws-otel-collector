// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleHardcodedEvaluatorType Hardcoded evaluator type.
type SecurityMonitoringRuleHardcodedEvaluatorType string

// List of SecurityMonitoringRuleHardcodedEvaluatorType.
const (
	SECURITYMONITORINGRULEHARDCODEDEVALUATORTYPE_LOG4SHELL SecurityMonitoringRuleHardcodedEvaluatorType = "log4shell"
)

var allowedSecurityMonitoringRuleHardcodedEvaluatorTypeEnumValues = []SecurityMonitoringRuleHardcodedEvaluatorType{
	SECURITYMONITORINGRULEHARDCODEDEVALUATORTYPE_LOG4SHELL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringRuleHardcodedEvaluatorType) GetAllowedValues() []SecurityMonitoringRuleHardcodedEvaluatorType {
	return allowedSecurityMonitoringRuleHardcodedEvaluatorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringRuleHardcodedEvaluatorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringRuleHardcodedEvaluatorType(value)
	return nil
}

// NewSecurityMonitoringRuleHardcodedEvaluatorTypeFromValue returns a pointer to a valid SecurityMonitoringRuleHardcodedEvaluatorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringRuleHardcodedEvaluatorTypeFromValue(v string) (*SecurityMonitoringRuleHardcodedEvaluatorType, error) {
	ev := SecurityMonitoringRuleHardcodedEvaluatorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringRuleHardcodedEvaluatorType: valid values are %v", v, allowedSecurityMonitoringRuleHardcodedEvaluatorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringRuleHardcodedEvaluatorType) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringRuleHardcodedEvaluatorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringRuleHardcodedEvaluatorType value.
func (v SecurityMonitoringRuleHardcodedEvaluatorType) Ptr() *SecurityMonitoringRuleHardcodedEvaluatorType {
	return &v
}
