// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleAnomalyDetectionOptionsBucketDuration Duration in seconds of the time buckets used to aggregate events matched by the rule.
// Must be greater than or equal to 300.
type SecurityMonitoringRuleAnomalyDetectionOptionsBucketDuration int32

// List of SecurityMonitoringRuleAnomalyDetectionOptionsBucketDuration.
const (
	SECURITYMONITORINGRULEANOMALYDETECTIONOPTIONSBUCKETDURATION_FIVE_MINUTES    SecurityMonitoringRuleAnomalyDetectionOptionsBucketDuration = 300
	SECURITYMONITORINGRULEANOMALYDETECTIONOPTIONSBUCKETDURATION_TEN_MINUTES     SecurityMonitoringRuleAnomalyDetectionOptionsBucketDuration = 600
	SECURITYMONITORINGRULEANOMALYDETECTIONOPTIONSBUCKETDURATION_FIFTEEN_MINUTES SecurityMonitoringRuleAnomalyDetectionOptionsBucketDuration = 900
	SECURITYMONITORINGRULEANOMALYDETECTIONOPTIONSBUCKETDURATION_THIRTY_MINUTES  SecurityMonitoringRuleAnomalyDetectionOptionsBucketDuration = 1800
	SECURITYMONITORINGRULEANOMALYDETECTIONOPTIONSBUCKETDURATION_ONE_HOUR        SecurityMonitoringRuleAnomalyDetectionOptionsBucketDuration = 3600
	SECURITYMONITORINGRULEANOMALYDETECTIONOPTIONSBUCKETDURATION_THREE_HOURS     SecurityMonitoringRuleAnomalyDetectionOptionsBucketDuration = 10800
)

var allowedSecurityMonitoringRuleAnomalyDetectionOptionsBucketDurationEnumValues = []SecurityMonitoringRuleAnomalyDetectionOptionsBucketDuration{
	SECURITYMONITORINGRULEANOMALYDETECTIONOPTIONSBUCKETDURATION_FIVE_MINUTES,
	SECURITYMONITORINGRULEANOMALYDETECTIONOPTIONSBUCKETDURATION_TEN_MINUTES,
	SECURITYMONITORINGRULEANOMALYDETECTIONOPTIONSBUCKETDURATION_FIFTEEN_MINUTES,
	SECURITYMONITORINGRULEANOMALYDETECTIONOPTIONSBUCKETDURATION_THIRTY_MINUTES,
	SECURITYMONITORINGRULEANOMALYDETECTIONOPTIONSBUCKETDURATION_ONE_HOUR,
	SECURITYMONITORINGRULEANOMALYDETECTIONOPTIONSBUCKETDURATION_THREE_HOURS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringRuleAnomalyDetectionOptionsBucketDuration) GetAllowedValues() []SecurityMonitoringRuleAnomalyDetectionOptionsBucketDuration {
	return allowedSecurityMonitoringRuleAnomalyDetectionOptionsBucketDurationEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringRuleAnomalyDetectionOptionsBucketDuration) UnmarshalJSON(src []byte) error {
	var value int32
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringRuleAnomalyDetectionOptionsBucketDuration(value)
	return nil
}

// NewSecurityMonitoringRuleAnomalyDetectionOptionsBucketDurationFromValue returns a pointer to a valid SecurityMonitoringRuleAnomalyDetectionOptionsBucketDuration
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringRuleAnomalyDetectionOptionsBucketDurationFromValue(v int32) (*SecurityMonitoringRuleAnomalyDetectionOptionsBucketDuration, error) {
	ev := SecurityMonitoringRuleAnomalyDetectionOptionsBucketDuration(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringRuleAnomalyDetectionOptionsBucketDuration: valid values are %v", v, allowedSecurityMonitoringRuleAnomalyDetectionOptionsBucketDurationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringRuleAnomalyDetectionOptionsBucketDuration) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringRuleAnomalyDetectionOptionsBucketDurationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringRuleAnomalyDetectionOptionsBucketDuration value.
func (v SecurityMonitoringRuleAnomalyDetectionOptionsBucketDuration) Ptr() *SecurityMonitoringRuleAnomalyDetectionOptionsBucketDuration {
	return &v
}
