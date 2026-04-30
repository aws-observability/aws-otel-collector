// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentGatesEvaluationRequestDataType JSON:API type for a deployment gate evaluation request.
type DeploymentGatesEvaluationRequestDataType string

// List of DeploymentGatesEvaluationRequestDataType.
const (
	DEPLOYMENTGATESEVALUATIONREQUESTDATATYPE_DEPLOYMENT_GATES_EVALUATION_REQUEST DeploymentGatesEvaluationRequestDataType = "deployment_gates_evaluation_request"
)

var allowedDeploymentGatesEvaluationRequestDataTypeEnumValues = []DeploymentGatesEvaluationRequestDataType{
	DEPLOYMENTGATESEVALUATIONREQUESTDATATYPE_DEPLOYMENT_GATES_EVALUATION_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DeploymentGatesEvaluationRequestDataType) GetAllowedValues() []DeploymentGatesEvaluationRequestDataType {
	return allowedDeploymentGatesEvaluationRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DeploymentGatesEvaluationRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DeploymentGatesEvaluationRequestDataType(value)
	return nil
}

// NewDeploymentGatesEvaluationRequestDataTypeFromValue returns a pointer to a valid DeploymentGatesEvaluationRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDeploymentGatesEvaluationRequestDataTypeFromValue(v string) (*DeploymentGatesEvaluationRequestDataType, error) {
	ev := DeploymentGatesEvaluationRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DeploymentGatesEvaluationRequestDataType: valid values are %v", v, allowedDeploymentGatesEvaluationRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DeploymentGatesEvaluationRequestDataType) IsValid() bool {
	for _, existing := range allowedDeploymentGatesEvaluationRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeploymentGatesEvaluationRequestDataType value.
func (v DeploymentGatesEvaluationRequestDataType) Ptr() *DeploymentGatesEvaluationRequestDataType {
	return &v
}
