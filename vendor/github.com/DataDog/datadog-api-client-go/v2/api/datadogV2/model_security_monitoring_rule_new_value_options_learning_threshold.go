// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleNewValueOptionsLearningThreshold A number of occurrences after which signals will be generated for values that weren't learned.
type SecurityMonitoringRuleNewValueOptionsLearningThreshold int32

// List of SecurityMonitoringRuleNewValueOptionsLearningThreshold.
const (
	SECURITYMONITORINGRULENEWVALUEOPTIONSLEARNINGTHRESHOLD_ZERO_OCCURRENCES SecurityMonitoringRuleNewValueOptionsLearningThreshold = 0
	SECURITYMONITORINGRULENEWVALUEOPTIONSLEARNINGTHRESHOLD_ONE_OCCURRENCE   SecurityMonitoringRuleNewValueOptionsLearningThreshold = 1
)

var allowedSecurityMonitoringRuleNewValueOptionsLearningThresholdEnumValues = []SecurityMonitoringRuleNewValueOptionsLearningThreshold{
	SECURITYMONITORINGRULENEWVALUEOPTIONSLEARNINGTHRESHOLD_ZERO_OCCURRENCES,
	SECURITYMONITORINGRULENEWVALUEOPTIONSLEARNINGTHRESHOLD_ONE_OCCURRENCE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringRuleNewValueOptionsLearningThreshold) GetAllowedValues() []SecurityMonitoringRuleNewValueOptionsLearningThreshold {
	return allowedSecurityMonitoringRuleNewValueOptionsLearningThresholdEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringRuleNewValueOptionsLearningThreshold) UnmarshalJSON(src []byte) error {
	var value int32
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringRuleNewValueOptionsLearningThreshold(value)
	return nil
}

// NewSecurityMonitoringRuleNewValueOptionsLearningThresholdFromValue returns a pointer to a valid SecurityMonitoringRuleNewValueOptionsLearningThreshold
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringRuleNewValueOptionsLearningThresholdFromValue(v int32) (*SecurityMonitoringRuleNewValueOptionsLearningThreshold, error) {
	ev := SecurityMonitoringRuleNewValueOptionsLearningThreshold(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringRuleNewValueOptionsLearningThreshold: valid values are %v", v, allowedSecurityMonitoringRuleNewValueOptionsLearningThresholdEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringRuleNewValueOptionsLearningThreshold) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringRuleNewValueOptionsLearningThresholdEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringRuleNewValueOptionsLearningThreshold value.
func (v SecurityMonitoringRuleNewValueOptionsLearningThreshold) Ptr() *SecurityMonitoringRuleNewValueOptionsLearningThreshold {
	return &v
}
