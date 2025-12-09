// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSAuthConfig - AWS Authentication config.
type AWSAuthConfig struct {
	AWSAuthConfigKeys *AWSAuthConfigKeys
	AWSAuthConfigRole *AWSAuthConfigRole

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// AWSAuthConfigKeysAsAWSAuthConfig is a convenience function that returns AWSAuthConfigKeys wrapped in AWSAuthConfig.
func AWSAuthConfigKeysAsAWSAuthConfig(v *AWSAuthConfigKeys) AWSAuthConfig {
	return AWSAuthConfig{AWSAuthConfigKeys: v}
}

// AWSAuthConfigRoleAsAWSAuthConfig is a convenience function that returns AWSAuthConfigRole wrapped in AWSAuthConfig.
func AWSAuthConfigRoleAsAWSAuthConfig(v *AWSAuthConfigRole) AWSAuthConfig {
	return AWSAuthConfig{AWSAuthConfigRole: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *AWSAuthConfig) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AWSAuthConfigKeys
	err = datadog.Unmarshal(data, &obj.AWSAuthConfigKeys)
	if err == nil {
		if obj.AWSAuthConfigKeys != nil && obj.AWSAuthConfigKeys.UnparsedObject == nil {
			jsonAWSAuthConfigKeys, _ := datadog.Marshal(obj.AWSAuthConfigKeys)
			if string(jsonAWSAuthConfigKeys) == "{}" { // empty struct
				obj.AWSAuthConfigKeys = nil
			} else {
				match++
			}
		} else {
			obj.AWSAuthConfigKeys = nil
		}
	} else {
		obj.AWSAuthConfigKeys = nil
	}

	// try to unmarshal data into AWSAuthConfigRole
	err = datadog.Unmarshal(data, &obj.AWSAuthConfigRole)
	if err == nil {
		if obj.AWSAuthConfigRole != nil && obj.AWSAuthConfigRole.UnparsedObject == nil {
			jsonAWSAuthConfigRole, _ := datadog.Marshal(obj.AWSAuthConfigRole)
			if string(jsonAWSAuthConfigRole) == "{}" { // empty struct
				obj.AWSAuthConfigRole = nil
			} else {
				match++
			}
		} else {
			obj.AWSAuthConfigRole = nil
		}
	} else {
		obj.AWSAuthConfigRole = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.AWSAuthConfigKeys = nil
		obj.AWSAuthConfigRole = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj AWSAuthConfig) MarshalJSON() ([]byte, error) {
	if obj.AWSAuthConfigKeys != nil {
		return datadog.Marshal(&obj.AWSAuthConfigKeys)
	}

	if obj.AWSAuthConfigRole != nil {
		return datadog.Marshal(&obj.AWSAuthConfigRole)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *AWSAuthConfig) GetActualInstance() interface{} {
	if obj.AWSAuthConfigKeys != nil {
		return obj.AWSAuthConfigKeys
	}

	if obj.AWSAuthConfigRole != nil {
		return obj.AWSAuthConfigRole
	}

	// all schemas are nil
	return nil
}
