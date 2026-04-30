// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentGatesEvaluationResultResponseAttributesGateStatus The overall status of the gate evaluation.
// - `in_progress`: The evaluation is still running.
// - `pass`: All rules passed successfully and the deployment is allowed to proceed.
// - `fail`: One or more rules did not pass; the deployment should not proceed.
type DeploymentGatesEvaluationResultResponseAttributesGateStatus string

// List of DeploymentGatesEvaluationResultResponseAttributesGateStatus.
const (
	DEPLOYMENTGATESEVALUATIONRESULTRESPONSEATTRIBUTESGATESTATUS_IN_PROGRESS DeploymentGatesEvaluationResultResponseAttributesGateStatus = "in_progress"
	DEPLOYMENTGATESEVALUATIONRESULTRESPONSEATTRIBUTESGATESTATUS_PASS        DeploymentGatesEvaluationResultResponseAttributesGateStatus = "pass"
	DEPLOYMENTGATESEVALUATIONRESULTRESPONSEATTRIBUTESGATESTATUS_FAIL        DeploymentGatesEvaluationResultResponseAttributesGateStatus = "fail"
)

var allowedDeploymentGatesEvaluationResultResponseAttributesGateStatusEnumValues = []DeploymentGatesEvaluationResultResponseAttributesGateStatus{
	DEPLOYMENTGATESEVALUATIONRESULTRESPONSEATTRIBUTESGATESTATUS_IN_PROGRESS,
	DEPLOYMENTGATESEVALUATIONRESULTRESPONSEATTRIBUTESGATESTATUS_PASS,
	DEPLOYMENTGATESEVALUATIONRESULTRESPONSEATTRIBUTESGATESTATUS_FAIL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DeploymentGatesEvaluationResultResponseAttributesGateStatus) GetAllowedValues() []DeploymentGatesEvaluationResultResponseAttributesGateStatus {
	return allowedDeploymentGatesEvaluationResultResponseAttributesGateStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DeploymentGatesEvaluationResultResponseAttributesGateStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DeploymentGatesEvaluationResultResponseAttributesGateStatus(value)
	return nil
}

// NewDeploymentGatesEvaluationResultResponseAttributesGateStatusFromValue returns a pointer to a valid DeploymentGatesEvaluationResultResponseAttributesGateStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDeploymentGatesEvaluationResultResponseAttributesGateStatusFromValue(v string) (*DeploymentGatesEvaluationResultResponseAttributesGateStatus, error) {
	ev := DeploymentGatesEvaluationResultResponseAttributesGateStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DeploymentGatesEvaluationResultResponseAttributesGateStatus: valid values are %v", v, allowedDeploymentGatesEvaluationResultResponseAttributesGateStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DeploymentGatesEvaluationResultResponseAttributesGateStatus) IsValid() bool {
	for _, existing := range allowedDeploymentGatesEvaluationResultResponseAttributesGateStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeploymentGatesEvaluationResultResponseAttributesGateStatus value.
func (v DeploymentGatesEvaluationResultResponseAttributesGateStatus) Ptr() *DeploymentGatesEvaluationResultResponseAttributesGateStatus {
	return &v
}
