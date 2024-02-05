// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorConfigPolicyPolicyCreateRequest - Configuration for the policy.
type MonitorConfigPolicyPolicyCreateRequest struct {
	MonitorConfigPolicyTagPolicyCreateRequest *MonitorConfigPolicyTagPolicyCreateRequest

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// MonitorConfigPolicyTagPolicyCreateRequestAsMonitorConfigPolicyPolicyCreateRequest is a convenience function that returns MonitorConfigPolicyTagPolicyCreateRequest wrapped in MonitorConfigPolicyPolicyCreateRequest.
func MonitorConfigPolicyTagPolicyCreateRequestAsMonitorConfigPolicyPolicyCreateRequest(v *MonitorConfigPolicyTagPolicyCreateRequest) MonitorConfigPolicyPolicyCreateRequest {
	return MonitorConfigPolicyPolicyCreateRequest{MonitorConfigPolicyTagPolicyCreateRequest: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *MonitorConfigPolicyPolicyCreateRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into MonitorConfigPolicyTagPolicyCreateRequest
	err = datadog.Unmarshal(data, &obj.MonitorConfigPolicyTagPolicyCreateRequest)
	if err == nil {
		if obj.MonitorConfigPolicyTagPolicyCreateRequest != nil && obj.MonitorConfigPolicyTagPolicyCreateRequest.UnparsedObject == nil {
			jsonMonitorConfigPolicyTagPolicyCreateRequest, _ := datadog.Marshal(obj.MonitorConfigPolicyTagPolicyCreateRequest)
			if string(jsonMonitorConfigPolicyTagPolicyCreateRequest) == "{}" { // empty struct
				obj.MonitorConfigPolicyTagPolicyCreateRequest = nil
			} else {
				match++
			}
		} else {
			obj.MonitorConfigPolicyTagPolicyCreateRequest = nil
		}
	} else {
		obj.MonitorConfigPolicyTagPolicyCreateRequest = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.MonitorConfigPolicyTagPolicyCreateRequest = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj MonitorConfigPolicyPolicyCreateRequest) MarshalJSON() ([]byte, error) {
	if obj.MonitorConfigPolicyTagPolicyCreateRequest != nil {
		return datadog.Marshal(&obj.MonitorConfigPolicyTagPolicyCreateRequest)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *MonitorConfigPolicyPolicyCreateRequest) GetActualInstance() interface{} {
	if obj.MonitorConfigPolicyTagPolicyCreateRequest != nil {
		return obj.MonitorConfigPolicyTagPolicyCreateRequest
	}

	// all schemas are nil
	return nil
}
