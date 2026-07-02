// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentGatesMonitorRuleType The type identifier for a monitor rule.
type DeploymentGatesMonitorRuleType string

// List of DeploymentGatesMonitorRuleType.
const (
	DEPLOYMENTGATESMONITORRULETYPE_MONITOR DeploymentGatesMonitorRuleType = "monitor"
)

var allowedDeploymentGatesMonitorRuleTypeEnumValues = []DeploymentGatesMonitorRuleType{
	DEPLOYMENTGATESMONITORRULETYPE_MONITOR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DeploymentGatesMonitorRuleType) GetAllowedValues() []DeploymentGatesMonitorRuleType {
	return allowedDeploymentGatesMonitorRuleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DeploymentGatesMonitorRuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DeploymentGatesMonitorRuleType(value)
	return nil
}

// NewDeploymentGatesMonitorRuleTypeFromValue returns a pointer to a valid DeploymentGatesMonitorRuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDeploymentGatesMonitorRuleTypeFromValue(v string) (*DeploymentGatesMonitorRuleType, error) {
	ev := DeploymentGatesMonitorRuleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DeploymentGatesMonitorRuleType: valid values are %v", v, allowedDeploymentGatesMonitorRuleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DeploymentGatesMonitorRuleType) IsValid() bool {
	for _, existing := range allowedDeploymentGatesMonitorRuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeploymentGatesMonitorRuleType value.
func (v DeploymentGatesMonitorRuleType) Ptr() *DeploymentGatesMonitorRuleType {
	return &v
}
