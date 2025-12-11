// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudflareCredentials - The definition of the `CloudflareCredentials` object.
type CloudflareCredentials struct {
	CloudflareAPIToken       *CloudflareAPIToken
	CloudflareGlobalAPIToken *CloudflareGlobalAPIToken

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// CloudflareAPITokenAsCloudflareCredentials is a convenience function that returns CloudflareAPIToken wrapped in CloudflareCredentials.
func CloudflareAPITokenAsCloudflareCredentials(v *CloudflareAPIToken) CloudflareCredentials {
	return CloudflareCredentials{CloudflareAPIToken: v}
}

// CloudflareGlobalAPITokenAsCloudflareCredentials is a convenience function that returns CloudflareGlobalAPIToken wrapped in CloudflareCredentials.
func CloudflareGlobalAPITokenAsCloudflareCredentials(v *CloudflareGlobalAPIToken) CloudflareCredentials {
	return CloudflareCredentials{CloudflareGlobalAPIToken: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *CloudflareCredentials) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CloudflareAPIToken
	err = datadog.Unmarshal(data, &obj.CloudflareAPIToken)
	if err == nil {
		if obj.CloudflareAPIToken != nil && obj.CloudflareAPIToken.UnparsedObject == nil {
			jsonCloudflareAPIToken, _ := datadog.Marshal(obj.CloudflareAPIToken)
			if string(jsonCloudflareAPIToken) == "{}" { // empty struct
				obj.CloudflareAPIToken = nil
			} else {
				match++
			}
		} else {
			obj.CloudflareAPIToken = nil
		}
	} else {
		obj.CloudflareAPIToken = nil
	}

	// try to unmarshal data into CloudflareGlobalAPIToken
	err = datadog.Unmarshal(data, &obj.CloudflareGlobalAPIToken)
	if err == nil {
		if obj.CloudflareGlobalAPIToken != nil && obj.CloudflareGlobalAPIToken.UnparsedObject == nil {
			jsonCloudflareGlobalAPIToken, _ := datadog.Marshal(obj.CloudflareGlobalAPIToken)
			if string(jsonCloudflareGlobalAPIToken) == "{}" { // empty struct
				obj.CloudflareGlobalAPIToken = nil
			} else {
				match++
			}
		} else {
			obj.CloudflareGlobalAPIToken = nil
		}
	} else {
		obj.CloudflareGlobalAPIToken = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.CloudflareAPIToken = nil
		obj.CloudflareGlobalAPIToken = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj CloudflareCredentials) MarshalJSON() ([]byte, error) {
	if obj.CloudflareAPIToken != nil {
		return datadog.Marshal(&obj.CloudflareAPIToken)
	}

	if obj.CloudflareGlobalAPIToken != nil {
		return datadog.Marshal(&obj.CloudflareGlobalAPIToken)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *CloudflareCredentials) GetActualInstance() interface{} {
	if obj.CloudflareAPIToken != nil {
		return obj.CloudflareAPIToken
	}

	if obj.CloudflareGlobalAPIToken != nil {
		return obj.CloudflareGlobalAPIToken
	}

	// all schemas are nil
	return nil
}
