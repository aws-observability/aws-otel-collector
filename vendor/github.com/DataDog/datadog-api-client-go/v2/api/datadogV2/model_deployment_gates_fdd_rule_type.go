// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentGatesFDDRuleType The type identifier for a faulty deployment detection rule.
type DeploymentGatesFDDRuleType string

// List of DeploymentGatesFDDRuleType.
const (
	DEPLOYMENTGATESFDDRULETYPE_FAULTY_DEPLOYMENT_DETECTION DeploymentGatesFDDRuleType = "faulty_deployment_detection"
)

var allowedDeploymentGatesFDDRuleTypeEnumValues = []DeploymentGatesFDDRuleType{
	DEPLOYMENTGATESFDDRULETYPE_FAULTY_DEPLOYMENT_DETECTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DeploymentGatesFDDRuleType) GetAllowedValues() []DeploymentGatesFDDRuleType {
	return allowedDeploymentGatesFDDRuleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DeploymentGatesFDDRuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DeploymentGatesFDDRuleType(value)
	return nil
}

// NewDeploymentGatesFDDRuleTypeFromValue returns a pointer to a valid DeploymentGatesFDDRuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDeploymentGatesFDDRuleTypeFromValue(v string) (*DeploymentGatesFDDRuleType, error) {
	ev := DeploymentGatesFDDRuleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DeploymentGatesFDDRuleType: valid values are %v", v, allowedDeploymentGatesFDDRuleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DeploymentGatesFDDRuleType) IsValid() bool {
	for _, existing := range allowedDeploymentGatesFDDRuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeploymentGatesFDDRuleType value.
func (v DeploymentGatesFDDRuleType) Ptr() *DeploymentGatesFDDRuleType {
	return &v
}
