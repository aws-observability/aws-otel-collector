// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentGatesEvaluationResponseDataType JSON:API type for a deployment gate evaluation response.
type DeploymentGatesEvaluationResponseDataType string

// List of DeploymentGatesEvaluationResponseDataType.
const (
	DEPLOYMENTGATESEVALUATIONRESPONSEDATATYPE_DEPLOYMENT_GATES_EVALUATION_RESPONSE DeploymentGatesEvaluationResponseDataType = "deployment_gates_evaluation_response"
)

var allowedDeploymentGatesEvaluationResponseDataTypeEnumValues = []DeploymentGatesEvaluationResponseDataType{
	DEPLOYMENTGATESEVALUATIONRESPONSEDATATYPE_DEPLOYMENT_GATES_EVALUATION_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DeploymentGatesEvaluationResponseDataType) GetAllowedValues() []DeploymentGatesEvaluationResponseDataType {
	return allowedDeploymentGatesEvaluationResponseDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DeploymentGatesEvaluationResponseDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DeploymentGatesEvaluationResponseDataType(value)
	return nil
}

// NewDeploymentGatesEvaluationResponseDataTypeFromValue returns a pointer to a valid DeploymentGatesEvaluationResponseDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDeploymentGatesEvaluationResponseDataTypeFromValue(v string) (*DeploymentGatesEvaluationResponseDataType, error) {
	ev := DeploymentGatesEvaluationResponseDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DeploymentGatesEvaluationResponseDataType: valid values are %v", v, allowedDeploymentGatesEvaluationResponseDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DeploymentGatesEvaluationResponseDataType) IsValid() bool {
	for _, existing := range allowedDeploymentGatesEvaluationResponseDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeploymentGatesEvaluationResponseDataType value.
func (v DeploymentGatesEvaluationResponseDataType) Ptr() *DeploymentGatesEvaluationResponseDataType {
	return &v
}
