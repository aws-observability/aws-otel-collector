// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudflareCredentialsUpdate - The definition of the `CloudflareCredentialsUpdate` object.
type CloudflareCredentialsUpdate struct {
	CloudflareAPITokenUpdate       *CloudflareAPITokenUpdate
	CloudflareGlobalAPITokenUpdate *CloudflareGlobalAPITokenUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// CloudflareAPITokenUpdateAsCloudflareCredentialsUpdate is a convenience function that returns CloudflareAPITokenUpdate wrapped in CloudflareCredentialsUpdate.
func CloudflareAPITokenUpdateAsCloudflareCredentialsUpdate(v *CloudflareAPITokenUpdate) CloudflareCredentialsUpdate {
	return CloudflareCredentialsUpdate{CloudflareAPITokenUpdate: v}
}

// CloudflareGlobalAPITokenUpdateAsCloudflareCredentialsUpdate is a convenience function that returns CloudflareGlobalAPITokenUpdate wrapped in CloudflareCredentialsUpdate.
func CloudflareGlobalAPITokenUpdateAsCloudflareCredentialsUpdate(v *CloudflareGlobalAPITokenUpdate) CloudflareCredentialsUpdate {
	return CloudflareCredentialsUpdate{CloudflareGlobalAPITokenUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *CloudflareCredentialsUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CloudflareAPITokenUpdate
	err = datadog.Unmarshal(data, &obj.CloudflareAPITokenUpdate)
	if err == nil {
		if obj.CloudflareAPITokenUpdate != nil && obj.CloudflareAPITokenUpdate.UnparsedObject == nil {
			jsonCloudflareAPITokenUpdate, _ := datadog.Marshal(obj.CloudflareAPITokenUpdate)
			if string(jsonCloudflareAPITokenUpdate) == "{}" { // empty struct
				obj.CloudflareAPITokenUpdate = nil
			} else {
				match++
			}
		} else {
			obj.CloudflareAPITokenUpdate = nil
		}
	} else {
		obj.CloudflareAPITokenUpdate = nil
	}

	// try to unmarshal data into CloudflareGlobalAPITokenUpdate
	err = datadog.Unmarshal(data, &obj.CloudflareGlobalAPITokenUpdate)
	if err == nil {
		if obj.CloudflareGlobalAPITokenUpdate != nil && obj.CloudflareGlobalAPITokenUpdate.UnparsedObject == nil {
			jsonCloudflareGlobalAPITokenUpdate, _ := datadog.Marshal(obj.CloudflareGlobalAPITokenUpdate)
			if string(jsonCloudflareGlobalAPITokenUpdate) == "{}" { // empty struct
				obj.CloudflareGlobalAPITokenUpdate = nil
			} else {
				match++
			}
		} else {
			obj.CloudflareGlobalAPITokenUpdate = nil
		}
	} else {
		obj.CloudflareGlobalAPITokenUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.CloudflareAPITokenUpdate = nil
		obj.CloudflareGlobalAPITokenUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj CloudflareCredentialsUpdate) MarshalJSON() ([]byte, error) {
	if obj.CloudflareAPITokenUpdate != nil {
		return datadog.Marshal(&obj.CloudflareAPITokenUpdate)
	}

	if obj.CloudflareGlobalAPITokenUpdate != nil {
		return datadog.Marshal(&obj.CloudflareGlobalAPITokenUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *CloudflareCredentialsUpdate) GetActualInstance() interface{} {
	if obj.CloudflareAPITokenUpdate != nil {
		return obj.CloudflareAPITokenUpdate
	}

	if obj.CloudflareGlobalAPITokenUpdate != nil {
		return obj.CloudflareGlobalAPITokenUpdate
	}

	// all schemas are nil
	return nil
}
