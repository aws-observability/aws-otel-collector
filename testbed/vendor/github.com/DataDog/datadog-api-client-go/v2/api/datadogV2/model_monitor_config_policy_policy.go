// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorConfigPolicyPolicy - Configuration for the policy.
type MonitorConfigPolicyPolicy struct {
	MonitorConfigPolicyTagPolicy *MonitorConfigPolicyTagPolicy

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// MonitorConfigPolicyTagPolicyAsMonitorConfigPolicyPolicy is a convenience function that returns MonitorConfigPolicyTagPolicy wrapped in MonitorConfigPolicyPolicy.
func MonitorConfigPolicyTagPolicyAsMonitorConfigPolicyPolicy(v *MonitorConfigPolicyTagPolicy) MonitorConfigPolicyPolicy {
	return MonitorConfigPolicyPolicy{MonitorConfigPolicyTagPolicy: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *MonitorConfigPolicyPolicy) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into MonitorConfigPolicyTagPolicy
	err = datadog.Unmarshal(data, &obj.MonitorConfigPolicyTagPolicy)
	if err == nil {
		if obj.MonitorConfigPolicyTagPolicy != nil && obj.MonitorConfigPolicyTagPolicy.UnparsedObject == nil {
			jsonMonitorConfigPolicyTagPolicy, _ := datadog.Marshal(obj.MonitorConfigPolicyTagPolicy)
			if string(jsonMonitorConfigPolicyTagPolicy) == "{}" && string(data) != "{}" { // empty struct
				obj.MonitorConfigPolicyTagPolicy = nil
			} else {
				match++
			}
		} else {
			obj.MonitorConfigPolicyTagPolicy = nil
		}
	} else {
		obj.MonitorConfigPolicyTagPolicy = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.MonitorConfigPolicyTagPolicy = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj MonitorConfigPolicyPolicy) MarshalJSON() ([]byte, error) {
	if obj.MonitorConfigPolicyTagPolicy != nil {
		return datadog.Marshal(&obj.MonitorConfigPolicyTagPolicy)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *MonitorConfigPolicyPolicy) GetActualInstance() interface{} {
	if obj.MonitorConfigPolicyTagPolicy != nil {
		return obj.MonitorConfigPolicyTagPolicy
	}

	// all schemas are nil
	return nil
}
