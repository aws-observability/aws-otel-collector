// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleTypeTest The rule type.
type SecurityMonitoringRuleTypeTest string

// List of SecurityMonitoringRuleTypeTest.
const (
	SECURITYMONITORINGRULETYPETEST_LOG_DETECTION SecurityMonitoringRuleTypeTest = "log_detection"
)

var allowedSecurityMonitoringRuleTypeTestEnumValues = []SecurityMonitoringRuleTypeTest{
	SECURITYMONITORINGRULETYPETEST_LOG_DETECTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringRuleTypeTest) GetAllowedValues() []SecurityMonitoringRuleTypeTest {
	return allowedSecurityMonitoringRuleTypeTestEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringRuleTypeTest) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringRuleTypeTest(value)
	return nil
}

// NewSecurityMonitoringRuleTypeTestFromValue returns a pointer to a valid SecurityMonitoringRuleTypeTest
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringRuleTypeTestFromValue(v string) (*SecurityMonitoringRuleTypeTest, error) {
	ev := SecurityMonitoringRuleTypeTest(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringRuleTypeTest: valid values are %v", v, allowedSecurityMonitoringRuleTypeTestEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringRuleTypeTest) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringRuleTypeTestEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringRuleTypeTest value.
func (v SecurityMonitoringRuleTypeTest) Ptr() *SecurityMonitoringRuleTypeTest {
	return &v
}
