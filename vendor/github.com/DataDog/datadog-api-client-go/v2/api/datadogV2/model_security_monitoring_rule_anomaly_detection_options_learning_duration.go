// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleAnomalyDetectionOptionsLearningDuration Learning duration in hours. Anomaly detection waits for at least this amount of historical data before it starts evaluating.
type SecurityMonitoringRuleAnomalyDetectionOptionsLearningDuration int32

// List of SecurityMonitoringRuleAnomalyDetectionOptionsLearningDuration.
const (
	SECURITYMONITORINGRULEANOMALYDETECTIONOPTIONSLEARNINGDURATION_ONE_HOUR     SecurityMonitoringRuleAnomalyDetectionOptionsLearningDuration = 1
	SECURITYMONITORINGRULEANOMALYDETECTIONOPTIONSLEARNINGDURATION_SIX_HOURS    SecurityMonitoringRuleAnomalyDetectionOptionsLearningDuration = 6
	SECURITYMONITORINGRULEANOMALYDETECTIONOPTIONSLEARNINGDURATION_TWELVE_HOURS SecurityMonitoringRuleAnomalyDetectionOptionsLearningDuration = 12
	SECURITYMONITORINGRULEANOMALYDETECTIONOPTIONSLEARNINGDURATION_ONE_DAY      SecurityMonitoringRuleAnomalyDetectionOptionsLearningDuration = 24
	SECURITYMONITORINGRULEANOMALYDETECTIONOPTIONSLEARNINGDURATION_TWO_DAYS     SecurityMonitoringRuleAnomalyDetectionOptionsLearningDuration = 48
	SECURITYMONITORINGRULEANOMALYDETECTIONOPTIONSLEARNINGDURATION_ONE_WEEK     SecurityMonitoringRuleAnomalyDetectionOptionsLearningDuration = 168
	SECURITYMONITORINGRULEANOMALYDETECTIONOPTIONSLEARNINGDURATION_TWO_WEEKS    SecurityMonitoringRuleAnomalyDetectionOptionsLearningDuration = 336
)

var allowedSecurityMonitoringRuleAnomalyDetectionOptionsLearningDurationEnumValues = []SecurityMonitoringRuleAnomalyDetectionOptionsLearningDuration{
	SECURITYMONITORINGRULEANOMALYDETECTIONOPTIONSLEARNINGDURATION_ONE_HOUR,
	SECURITYMONITORINGRULEANOMALYDETECTIONOPTIONSLEARNINGDURATION_SIX_HOURS,
	SECURITYMONITORINGRULEANOMALYDETECTIONOPTIONSLEARNINGDURATION_TWELVE_HOURS,
	SECURITYMONITORINGRULEANOMALYDETECTIONOPTIONSLEARNINGDURATION_ONE_DAY,
	SECURITYMONITORINGRULEANOMALYDETECTIONOPTIONSLEARNINGDURATION_TWO_DAYS,
	SECURITYMONITORINGRULEANOMALYDETECTIONOPTIONSLEARNINGDURATION_ONE_WEEK,
	SECURITYMONITORINGRULEANOMALYDETECTIONOPTIONSLEARNINGDURATION_TWO_WEEKS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringRuleAnomalyDetectionOptionsLearningDuration) GetAllowedValues() []SecurityMonitoringRuleAnomalyDetectionOptionsLearningDuration {
	return allowedSecurityMonitoringRuleAnomalyDetectionOptionsLearningDurationEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringRuleAnomalyDetectionOptionsLearningDuration) UnmarshalJSON(src []byte) error {
	var value int32
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringRuleAnomalyDetectionOptionsLearningDuration(value)
	return nil
}

// NewSecurityMonitoringRuleAnomalyDetectionOptionsLearningDurationFromValue returns a pointer to a valid SecurityMonitoringRuleAnomalyDetectionOptionsLearningDuration
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringRuleAnomalyDetectionOptionsLearningDurationFromValue(v int32) (*SecurityMonitoringRuleAnomalyDetectionOptionsLearningDuration, error) {
	ev := SecurityMonitoringRuleAnomalyDetectionOptionsLearningDuration(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringRuleAnomalyDetectionOptionsLearningDuration: valid values are %v", v, allowedSecurityMonitoringRuleAnomalyDetectionOptionsLearningDurationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringRuleAnomalyDetectionOptionsLearningDuration) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringRuleAnomalyDetectionOptionsLearningDurationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringRuleAnomalyDetectionOptionsLearningDuration value.
func (v SecurityMonitoringRuleAnomalyDetectionOptionsLearningDuration) Ptr() *SecurityMonitoringRuleAnomalyDetectionOptionsLearningDuration {
	return &v
}
