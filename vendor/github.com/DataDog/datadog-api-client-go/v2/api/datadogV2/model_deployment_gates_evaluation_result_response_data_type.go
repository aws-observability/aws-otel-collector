// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentGatesEvaluationResultResponseDataType JSON:API type for a deployment gate evaluation result response.
type DeploymentGatesEvaluationResultResponseDataType string

// List of DeploymentGatesEvaluationResultResponseDataType.
const (
	DEPLOYMENTGATESEVALUATIONRESULTRESPONSEDATATYPE_DEPLOYMENT_GATES_EVALUATION_RESULT_RESPONSE DeploymentGatesEvaluationResultResponseDataType = "deployment_gates_evaluation_result_response"
)

var allowedDeploymentGatesEvaluationResultResponseDataTypeEnumValues = []DeploymentGatesEvaluationResultResponseDataType{
	DEPLOYMENTGATESEVALUATIONRESULTRESPONSEDATATYPE_DEPLOYMENT_GATES_EVALUATION_RESULT_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DeploymentGatesEvaluationResultResponseDataType) GetAllowedValues() []DeploymentGatesEvaluationResultResponseDataType {
	return allowedDeploymentGatesEvaluationResultResponseDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DeploymentGatesEvaluationResultResponseDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DeploymentGatesEvaluationResultResponseDataType(value)
	return nil
}

// NewDeploymentGatesEvaluationResultResponseDataTypeFromValue returns a pointer to a valid DeploymentGatesEvaluationResultResponseDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDeploymentGatesEvaluationResultResponseDataTypeFromValue(v string) (*DeploymentGatesEvaluationResultResponseDataType, error) {
	ev := DeploymentGatesEvaluationResultResponseDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DeploymentGatesEvaluationResultResponseDataType: valid values are %v", v, allowedDeploymentGatesEvaluationResultResponseDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DeploymentGatesEvaluationResultResponseDataType) IsValid() bool {
	for _, existing := range allowedDeploymentGatesEvaluationResultResponseDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeploymentGatesEvaluationResultResponseDataType value.
func (v DeploymentGatesEvaluationResultResponseDataType) Ptr() *DeploymentGatesEvaluationResultResponseDataType {
	return &v
}
