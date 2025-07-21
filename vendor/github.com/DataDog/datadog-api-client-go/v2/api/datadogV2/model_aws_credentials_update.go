// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSCredentialsUpdate - The definition of `AWSCredentialsUpdate` object.
type AWSCredentialsUpdate struct {
	AWSAssumeRoleUpdate *AWSAssumeRoleUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// AWSAssumeRoleUpdateAsAWSCredentialsUpdate is a convenience function that returns AWSAssumeRoleUpdate wrapped in AWSCredentialsUpdate.
func AWSAssumeRoleUpdateAsAWSCredentialsUpdate(v *AWSAssumeRoleUpdate) AWSCredentialsUpdate {
	return AWSCredentialsUpdate{AWSAssumeRoleUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *AWSCredentialsUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AWSAssumeRoleUpdate
	err = datadog.Unmarshal(data, &obj.AWSAssumeRoleUpdate)
	if err == nil {
		if obj.AWSAssumeRoleUpdate != nil && obj.AWSAssumeRoleUpdate.UnparsedObject == nil {
			jsonAWSAssumeRoleUpdate, _ := datadog.Marshal(obj.AWSAssumeRoleUpdate)
			if string(jsonAWSAssumeRoleUpdate) == "{}" { // empty struct
				obj.AWSAssumeRoleUpdate = nil
			} else {
				match++
			}
		} else {
			obj.AWSAssumeRoleUpdate = nil
		}
	} else {
		obj.AWSAssumeRoleUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.AWSAssumeRoleUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj AWSCredentialsUpdate) MarshalJSON() ([]byte, error) {
	if obj.AWSAssumeRoleUpdate != nil {
		return datadog.Marshal(&obj.AWSAssumeRoleUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *AWSCredentialsUpdate) GetActualInstance() interface{} {
	if obj.AWSAssumeRoleUpdate != nil {
		return obj.AWSAssumeRoleUpdate
	}

	// all schemas are nil
	return nil
}
