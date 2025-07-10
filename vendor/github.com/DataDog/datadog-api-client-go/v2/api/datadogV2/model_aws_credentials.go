// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSCredentials - The definition of `AWSCredentials` object.
type AWSCredentials struct {
	AWSAssumeRole *AWSAssumeRole

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// AWSAssumeRoleAsAWSCredentials is a convenience function that returns AWSAssumeRole wrapped in AWSCredentials.
func AWSAssumeRoleAsAWSCredentials(v *AWSAssumeRole) AWSCredentials {
	return AWSCredentials{AWSAssumeRole: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *AWSCredentials) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AWSAssumeRole
	err = datadog.Unmarshal(data, &obj.AWSAssumeRole)
	if err == nil {
		if obj.AWSAssumeRole != nil && obj.AWSAssumeRole.UnparsedObject == nil {
			jsonAWSAssumeRole, _ := datadog.Marshal(obj.AWSAssumeRole)
			if string(jsonAWSAssumeRole) == "{}" { // empty struct
				obj.AWSAssumeRole = nil
			} else {
				match++
			}
		} else {
			obj.AWSAssumeRole = nil
		}
	} else {
		obj.AWSAssumeRole = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.AWSAssumeRole = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj AWSCredentials) MarshalJSON() ([]byte, error) {
	if obj.AWSAssumeRole != nil {
		return datadog.Marshal(&obj.AWSAssumeRole)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *AWSCredentials) GetActualInstance() interface{} {
	if obj.AWSAssumeRole != nil {
		return obj.AWSAssumeRole
	}

	// all schemas are nil
	return nil
}
