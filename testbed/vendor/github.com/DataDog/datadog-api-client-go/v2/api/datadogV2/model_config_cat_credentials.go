// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ConfigCatCredentials - The definition of the `ConfigCatCredentials` object.
type ConfigCatCredentials struct {
	ConfigCatSDKKey *ConfigCatSDKKey

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ConfigCatSDKKeyAsConfigCatCredentials is a convenience function that returns ConfigCatSDKKey wrapped in ConfigCatCredentials.
func ConfigCatSDKKeyAsConfigCatCredentials(v *ConfigCatSDKKey) ConfigCatCredentials {
	return ConfigCatCredentials{ConfigCatSDKKey: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ConfigCatCredentials) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ConfigCatSDKKey
	err = datadog.Unmarshal(data, &obj.ConfigCatSDKKey)
	if err == nil {
		if obj.ConfigCatSDKKey != nil && obj.ConfigCatSDKKey.UnparsedObject == nil {
			jsonConfigCatSDKKey, _ := datadog.Marshal(obj.ConfigCatSDKKey)
			if string(jsonConfigCatSDKKey) == "{}" { // empty struct
				obj.ConfigCatSDKKey = nil
			} else {
				match++
			}
		} else {
			obj.ConfigCatSDKKey = nil
		}
	} else {
		obj.ConfigCatSDKKey = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ConfigCatSDKKey = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ConfigCatCredentials) MarshalJSON() ([]byte, error) {
	if obj.ConfigCatSDKKey != nil {
		return datadog.Marshal(&obj.ConfigCatSDKKey)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ConfigCatCredentials) GetActualInstance() interface{} {
	if obj.ConfigCatSDKKey != nil {
		return obj.ConfigCatSDKKey
	}

	// all schemas are nil
	return nil
}
