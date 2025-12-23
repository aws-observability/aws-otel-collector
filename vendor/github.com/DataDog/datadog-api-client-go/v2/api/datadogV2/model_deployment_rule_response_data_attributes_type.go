// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentRuleResponseDataAttributesType The type of the deployment rule.
type DeploymentRuleResponseDataAttributesType string

// List of DeploymentRuleResponseDataAttributesType.
const (
	DEPLOYMENTRULERESPONSEDATAATTRIBUTESTYPE_FAULTY_DEPLOYMENT_DETECTION DeploymentRuleResponseDataAttributesType = "faulty_deployment_detection"
	DEPLOYMENTRULERESPONSEDATAATTRIBUTESTYPE_MONITOR                     DeploymentRuleResponseDataAttributesType = "monitor"
)

var allowedDeploymentRuleResponseDataAttributesTypeEnumValues = []DeploymentRuleResponseDataAttributesType{
	DEPLOYMENTRULERESPONSEDATAATTRIBUTESTYPE_FAULTY_DEPLOYMENT_DETECTION,
	DEPLOYMENTRULERESPONSEDATAATTRIBUTESTYPE_MONITOR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DeploymentRuleResponseDataAttributesType) GetAllowedValues() []DeploymentRuleResponseDataAttributesType {
	return allowedDeploymentRuleResponseDataAttributesTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DeploymentRuleResponseDataAttributesType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DeploymentRuleResponseDataAttributesType(value)
	return nil
}

// NewDeploymentRuleResponseDataAttributesTypeFromValue returns a pointer to a valid DeploymentRuleResponseDataAttributesType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDeploymentRuleResponseDataAttributesTypeFromValue(v string) (*DeploymentRuleResponseDataAttributesType, error) {
	ev := DeploymentRuleResponseDataAttributesType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DeploymentRuleResponseDataAttributesType: valid values are %v", v, allowedDeploymentRuleResponseDataAttributesTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DeploymentRuleResponseDataAttributesType) IsValid() bool {
	for _, existing := range allowedDeploymentRuleResponseDataAttributesTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeploymentRuleResponseDataAttributesType value.
func (v DeploymentRuleResponseDataAttributesType) Ptr() *DeploymentRuleResponseDataAttributesType {
	return &v
}
