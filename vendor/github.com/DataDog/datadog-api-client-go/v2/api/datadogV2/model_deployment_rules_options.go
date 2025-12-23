// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentRulesOptions - Options for deployment rule response representing either faulty deployment detection or monitor options.
type DeploymentRulesOptions struct {
	DeploymentRuleOptionsFaultyDeploymentDetection *DeploymentRuleOptionsFaultyDeploymentDetection
	DeploymentRuleOptionsMonitor                   *DeploymentRuleOptionsMonitor

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// DeploymentRuleOptionsFaultyDeploymentDetectionAsDeploymentRulesOptions is a convenience function that returns DeploymentRuleOptionsFaultyDeploymentDetection wrapped in DeploymentRulesOptions.
func DeploymentRuleOptionsFaultyDeploymentDetectionAsDeploymentRulesOptions(v *DeploymentRuleOptionsFaultyDeploymentDetection) DeploymentRulesOptions {
	return DeploymentRulesOptions{DeploymentRuleOptionsFaultyDeploymentDetection: v}
}

// DeploymentRuleOptionsMonitorAsDeploymentRulesOptions is a convenience function that returns DeploymentRuleOptionsMonitor wrapped in DeploymentRulesOptions.
func DeploymentRuleOptionsMonitorAsDeploymentRulesOptions(v *DeploymentRuleOptionsMonitor) DeploymentRulesOptions {
	return DeploymentRulesOptions{DeploymentRuleOptionsMonitor: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *DeploymentRulesOptions) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DeploymentRuleOptionsFaultyDeploymentDetection
	err = datadog.Unmarshal(data, &obj.DeploymentRuleOptionsFaultyDeploymentDetection)
	if err == nil {
		if obj.DeploymentRuleOptionsFaultyDeploymentDetection != nil && obj.DeploymentRuleOptionsFaultyDeploymentDetection.UnparsedObject == nil {
			jsonDeploymentRuleOptionsFaultyDeploymentDetection, _ := datadog.Marshal(obj.DeploymentRuleOptionsFaultyDeploymentDetection)
			if string(jsonDeploymentRuleOptionsFaultyDeploymentDetection) == "{}" && string(data) != "{}" { // empty struct
				obj.DeploymentRuleOptionsFaultyDeploymentDetection = nil
			} else {
				match++
			}
		} else {
			obj.DeploymentRuleOptionsFaultyDeploymentDetection = nil
		}
	} else {
		obj.DeploymentRuleOptionsFaultyDeploymentDetection = nil
	}

	// try to unmarshal data into DeploymentRuleOptionsMonitor
	err = datadog.Unmarshal(data, &obj.DeploymentRuleOptionsMonitor)
	if err == nil {
		if obj.DeploymentRuleOptionsMonitor != nil && obj.DeploymentRuleOptionsMonitor.UnparsedObject == nil {
			jsonDeploymentRuleOptionsMonitor, _ := datadog.Marshal(obj.DeploymentRuleOptionsMonitor)
			if string(jsonDeploymentRuleOptionsMonitor) == "{}" { // empty struct
				obj.DeploymentRuleOptionsMonitor = nil
			} else {
				match++
			}
		} else {
			obj.DeploymentRuleOptionsMonitor = nil
		}
	} else {
		obj.DeploymentRuleOptionsMonitor = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.DeploymentRuleOptionsFaultyDeploymentDetection = nil
		obj.DeploymentRuleOptionsMonitor = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj DeploymentRulesOptions) MarshalJSON() ([]byte, error) {
	if obj.DeploymentRuleOptionsFaultyDeploymentDetection != nil {
		return datadog.Marshal(&obj.DeploymentRuleOptionsFaultyDeploymentDetection)
	}

	if obj.DeploymentRuleOptionsMonitor != nil {
		return datadog.Marshal(&obj.DeploymentRuleOptionsMonitor)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *DeploymentRulesOptions) GetActualInstance() interface{} {
	if obj.DeploymentRuleOptionsFaultyDeploymentDetection != nil {
		return obj.DeploymentRuleOptionsFaultyDeploymentDetection
	}

	if obj.DeploymentRuleOptionsMonitor != nil {
		return obj.DeploymentRuleOptionsMonitor
	}

	// all schemas are nil
	return nil
}
