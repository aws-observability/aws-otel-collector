// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentRuleDataType Deployment rule resource type.
type DeploymentRuleDataType string

// List of DeploymentRuleDataType.
const (
	DEPLOYMENTRULEDATATYPE_DEPLOYMENT_RULE DeploymentRuleDataType = "deployment_rule"
)

var allowedDeploymentRuleDataTypeEnumValues = []DeploymentRuleDataType{
	DEPLOYMENTRULEDATATYPE_DEPLOYMENT_RULE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DeploymentRuleDataType) GetAllowedValues() []DeploymentRuleDataType {
	return allowedDeploymentRuleDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DeploymentRuleDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DeploymentRuleDataType(value)
	return nil
}

// NewDeploymentRuleDataTypeFromValue returns a pointer to a valid DeploymentRuleDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDeploymentRuleDataTypeFromValue(v string) (*DeploymentRuleDataType, error) {
	ev := DeploymentRuleDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DeploymentRuleDataType: valid values are %v", v, allowedDeploymentRuleDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DeploymentRuleDataType) IsValid() bool {
	for _, existing := range allowedDeploymentRuleDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeploymentRuleDataType value.
func (v DeploymentRuleDataType) Ptr() *DeploymentRuleDataType {
	return &v
}
