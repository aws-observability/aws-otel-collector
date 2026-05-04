// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleAnomalyDetectionOptionsDetectionTolerance An optional parameter that sets how permissive anomaly detection is.
// Higher values require higher deviations before triggering a signal.
type SecurityMonitoringRuleAnomalyDetectionOptionsDetectionTolerance int32

// List of SecurityMonitoringRuleAnomalyDetectionOptionsDetectionTolerance.
const (
	SECURITYMONITORINGRULEANOMALYDETECTIONOPTIONSDETECTIONTOLERANCE_ONE   SecurityMonitoringRuleAnomalyDetectionOptionsDetectionTolerance = 1
	SECURITYMONITORINGRULEANOMALYDETECTIONOPTIONSDETECTIONTOLERANCE_TWO   SecurityMonitoringRuleAnomalyDetectionOptionsDetectionTolerance = 2
	SECURITYMONITORINGRULEANOMALYDETECTIONOPTIONSDETECTIONTOLERANCE_THREE SecurityMonitoringRuleAnomalyDetectionOptionsDetectionTolerance = 3
	SECURITYMONITORINGRULEANOMALYDETECTIONOPTIONSDETECTIONTOLERANCE_FOUR  SecurityMonitoringRuleAnomalyDetectionOptionsDetectionTolerance = 4
	SECURITYMONITORINGRULEANOMALYDETECTIONOPTIONSDETECTIONTOLERANCE_FIVE  SecurityMonitoringRuleAnomalyDetectionOptionsDetectionTolerance = 5
)

var allowedSecurityMonitoringRuleAnomalyDetectionOptionsDetectionToleranceEnumValues = []SecurityMonitoringRuleAnomalyDetectionOptionsDetectionTolerance{
	SECURITYMONITORINGRULEANOMALYDETECTIONOPTIONSDETECTIONTOLERANCE_ONE,
	SECURITYMONITORINGRULEANOMALYDETECTIONOPTIONSDETECTIONTOLERANCE_TWO,
	SECURITYMONITORINGRULEANOMALYDETECTIONOPTIONSDETECTIONTOLERANCE_THREE,
	SECURITYMONITORINGRULEANOMALYDETECTIONOPTIONSDETECTIONTOLERANCE_FOUR,
	SECURITYMONITORINGRULEANOMALYDETECTIONOPTIONSDETECTIONTOLERANCE_FIVE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringRuleAnomalyDetectionOptionsDetectionTolerance) GetAllowedValues() []SecurityMonitoringRuleAnomalyDetectionOptionsDetectionTolerance {
	return allowedSecurityMonitoringRuleAnomalyDetectionOptionsDetectionToleranceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringRuleAnomalyDetectionOptionsDetectionTolerance) UnmarshalJSON(src []byte) error {
	var value int32
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringRuleAnomalyDetectionOptionsDetectionTolerance(value)
	return nil
}

// NewSecurityMonitoringRuleAnomalyDetectionOptionsDetectionToleranceFromValue returns a pointer to a valid SecurityMonitoringRuleAnomalyDetectionOptionsDetectionTolerance
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringRuleAnomalyDetectionOptionsDetectionToleranceFromValue(v int32) (*SecurityMonitoringRuleAnomalyDetectionOptionsDetectionTolerance, error) {
	ev := SecurityMonitoringRuleAnomalyDetectionOptionsDetectionTolerance(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringRuleAnomalyDetectionOptionsDetectionTolerance: valid values are %v", v, allowedSecurityMonitoringRuleAnomalyDetectionOptionsDetectionToleranceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringRuleAnomalyDetectionOptionsDetectionTolerance) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringRuleAnomalyDetectionOptionsDetectionToleranceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringRuleAnomalyDetectionOptionsDetectionTolerance value.
func (v SecurityMonitoringRuleAnomalyDetectionOptionsDetectionTolerance) Ptr() *SecurityMonitoringRuleAnomalyDetectionOptionsDetectionTolerance {
	return &v
}
