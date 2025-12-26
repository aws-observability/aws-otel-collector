// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentGateDataType Deployment gate resource type.
type DeploymentGateDataType string

// List of DeploymentGateDataType.
const (
	DEPLOYMENTGATEDATATYPE_DEPLOYMENT_GATE DeploymentGateDataType = "deployment_gate"
)

var allowedDeploymentGateDataTypeEnumValues = []DeploymentGateDataType{
	DEPLOYMENTGATEDATATYPE_DEPLOYMENT_GATE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DeploymentGateDataType) GetAllowedValues() []DeploymentGateDataType {
	return allowedDeploymentGateDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DeploymentGateDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DeploymentGateDataType(value)
	return nil
}

// NewDeploymentGateDataTypeFromValue returns a pointer to a valid DeploymentGateDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDeploymentGateDataTypeFromValue(v string) (*DeploymentGateDataType, error) {
	ev := DeploymentGateDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DeploymentGateDataType: valid values are %v", v, allowedDeploymentGateDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DeploymentGateDataType) IsValid() bool {
	for _, existing := range allowedDeploymentGateDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeploymentGateDataType value.
func (v DeploymentGateDataType) Ptr() *DeploymentGateDataType {
	return &v
}
