// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentGatesEvaluationRule - A rule to evaluate as part of a deployment gate evaluation.
type DeploymentGatesEvaluationRule struct {
	DeploymentGatesMonitorRule *DeploymentGatesMonitorRule
	DeploymentGatesFDDRule     *DeploymentGatesFDDRule

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// DeploymentGatesMonitorRuleAsDeploymentGatesEvaluationRule is a convenience function that returns DeploymentGatesMonitorRule wrapped in DeploymentGatesEvaluationRule.
func DeploymentGatesMonitorRuleAsDeploymentGatesEvaluationRule(v *DeploymentGatesMonitorRule) DeploymentGatesEvaluationRule {
	return DeploymentGatesEvaluationRule{DeploymentGatesMonitorRule: v}
}

// DeploymentGatesFDDRuleAsDeploymentGatesEvaluationRule is a convenience function that returns DeploymentGatesFDDRule wrapped in DeploymentGatesEvaluationRule.
func DeploymentGatesFDDRuleAsDeploymentGatesEvaluationRule(v *DeploymentGatesFDDRule) DeploymentGatesEvaluationRule {
	return DeploymentGatesEvaluationRule{DeploymentGatesFDDRule: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *DeploymentGatesEvaluationRule) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = datadog.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup.")
	}
	// check if the discriminator value is 'faulty_deployment_detection'
	if jsonDict["type"] == "faulty_deployment_detection" {
		// try to unmarshal JSON data into DeploymentGatesFDDRule
		err = datadog.Unmarshal(data, &obj.DeploymentGatesFDDRule)
		if err == nil {
			return nil // data stored in obj.DeploymentGatesFDDRule, return on the first match
		} else {
			obj.DeploymentGatesFDDRule = nil
			return fmt.Errorf("failed to unmarshal DeploymentGatesEvaluationRule as DeploymentGatesFDDRule: %s", err.Error())
		}
	}
	// check if the discriminator value is 'monitor'
	if jsonDict["type"] == "monitor" {
		// try to unmarshal JSON data into DeploymentGatesMonitorRule
		err = datadog.Unmarshal(data, &obj.DeploymentGatesMonitorRule)
		if err == nil {
			return nil // data stored in obj.DeploymentGatesMonitorRule, return on the first match
		} else {
			obj.DeploymentGatesMonitorRule = nil
			return fmt.Errorf("failed to unmarshal DeploymentGatesEvaluationRule as DeploymentGatesMonitorRule: %s", err.Error())
		}
	}
	return nil
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj DeploymentGatesEvaluationRule) MarshalJSON() ([]byte, error) {
	if obj.DeploymentGatesMonitorRule != nil {
		return datadog.Marshal(&obj.DeploymentGatesMonitorRule)
	}

	if obj.DeploymentGatesFDDRule != nil {
		return datadog.Marshal(&obj.DeploymentGatesFDDRule)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *DeploymentGatesEvaluationRule) GetActualInstance() interface{} {
	if obj.DeploymentGatesMonitorRule != nil {
		return obj.DeploymentGatesMonitorRule
	}

	if obj.DeploymentGatesFDDRule != nil {
		return obj.DeploymentGatesFDDRule
	}

	// all schemas are nil
	return nil
}
