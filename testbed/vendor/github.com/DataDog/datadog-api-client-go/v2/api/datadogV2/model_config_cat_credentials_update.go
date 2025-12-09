// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ConfigCatCredentialsUpdate - The definition of the `ConfigCatCredentialsUpdate` object.
type ConfigCatCredentialsUpdate struct {
	ConfigCatSDKKeyUpdate *ConfigCatSDKKeyUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ConfigCatSDKKeyUpdateAsConfigCatCredentialsUpdate is a convenience function that returns ConfigCatSDKKeyUpdate wrapped in ConfigCatCredentialsUpdate.
func ConfigCatSDKKeyUpdateAsConfigCatCredentialsUpdate(v *ConfigCatSDKKeyUpdate) ConfigCatCredentialsUpdate {
	return ConfigCatCredentialsUpdate{ConfigCatSDKKeyUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ConfigCatCredentialsUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ConfigCatSDKKeyUpdate
	err = datadog.Unmarshal(data, &obj.ConfigCatSDKKeyUpdate)
	if err == nil {
		if obj.ConfigCatSDKKeyUpdate != nil && obj.ConfigCatSDKKeyUpdate.UnparsedObject == nil {
			jsonConfigCatSDKKeyUpdate, _ := datadog.Marshal(obj.ConfigCatSDKKeyUpdate)
			if string(jsonConfigCatSDKKeyUpdate) == "{}" { // empty struct
				obj.ConfigCatSDKKeyUpdate = nil
			} else {
				match++
			}
		} else {
			obj.ConfigCatSDKKeyUpdate = nil
		}
	} else {
		obj.ConfigCatSDKKeyUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ConfigCatSDKKeyUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ConfigCatCredentialsUpdate) MarshalJSON() ([]byte, error) {
	if obj.ConfigCatSDKKeyUpdate != nil {
		return datadog.Marshal(&obj.ConfigCatSDKKeyUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ConfigCatCredentialsUpdate) GetActualInstance() interface{} {
	if obj.ConfigCatSDKKeyUpdate != nil {
		return obj.ConfigCatSDKKeyUpdate
	}

	// all schemas are nil
	return nil
}
