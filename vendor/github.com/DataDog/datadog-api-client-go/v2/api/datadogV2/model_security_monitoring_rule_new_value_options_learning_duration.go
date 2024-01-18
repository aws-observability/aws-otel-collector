// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleNewValueOptionsLearningDuration The duration in days during which values are learned, and after which signals will be generated for values that
// weren't learned. If set to 0, a signal will be generated for all new values after the first value is learned.
type SecurityMonitoringRuleNewValueOptionsLearningDuration int32

// List of SecurityMonitoringRuleNewValueOptionsLearningDuration.
const (
	SECURITYMONITORINGRULENEWVALUEOPTIONSLEARNINGDURATION_ZERO_DAYS  SecurityMonitoringRuleNewValueOptionsLearningDuration = 0
	SECURITYMONITORINGRULENEWVALUEOPTIONSLEARNINGDURATION_ONE_DAY    SecurityMonitoringRuleNewValueOptionsLearningDuration = 1
	SECURITYMONITORINGRULENEWVALUEOPTIONSLEARNINGDURATION_SEVEN_DAYS SecurityMonitoringRuleNewValueOptionsLearningDuration = 7
)

var allowedSecurityMonitoringRuleNewValueOptionsLearningDurationEnumValues = []SecurityMonitoringRuleNewValueOptionsLearningDuration{
	SECURITYMONITORINGRULENEWVALUEOPTIONSLEARNINGDURATION_ZERO_DAYS,
	SECURITYMONITORINGRULENEWVALUEOPTIONSLEARNINGDURATION_ONE_DAY,
	SECURITYMONITORINGRULENEWVALUEOPTIONSLEARNINGDURATION_SEVEN_DAYS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringRuleNewValueOptionsLearningDuration) GetAllowedValues() []SecurityMonitoringRuleNewValueOptionsLearningDuration {
	return allowedSecurityMonitoringRuleNewValueOptionsLearningDurationEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringRuleNewValueOptionsLearningDuration) UnmarshalJSON(src []byte) error {
	var value int32
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringRuleNewValueOptionsLearningDuration(value)
	return nil
}

// NewSecurityMonitoringRuleNewValueOptionsLearningDurationFromValue returns a pointer to a valid SecurityMonitoringRuleNewValueOptionsLearningDuration
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringRuleNewValueOptionsLearningDurationFromValue(v int32) (*SecurityMonitoringRuleNewValueOptionsLearningDuration, error) {
	ev := SecurityMonitoringRuleNewValueOptionsLearningDuration(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringRuleNewValueOptionsLearningDuration: valid values are %v", v, allowedSecurityMonitoringRuleNewValueOptionsLearningDurationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringRuleNewValueOptionsLearningDuration) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringRuleNewValueOptionsLearningDurationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringRuleNewValueOptionsLearningDuration value.
func (v SecurityMonitoringRuleNewValueOptionsLearningDuration) Ptr() *SecurityMonitoringRuleNewValueOptionsLearningDuration {
	return &v
}
