// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LaunchDarklyCredentials - The definition of the `LaunchDarklyCredentials` object.
type LaunchDarklyCredentials struct {
	LaunchDarklyAPIKey *LaunchDarklyAPIKey

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// LaunchDarklyAPIKeyAsLaunchDarklyCredentials is a convenience function that returns LaunchDarklyAPIKey wrapped in LaunchDarklyCredentials.
func LaunchDarklyAPIKeyAsLaunchDarklyCredentials(v *LaunchDarklyAPIKey) LaunchDarklyCredentials {
	return LaunchDarklyCredentials{LaunchDarklyAPIKey: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *LaunchDarklyCredentials) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into LaunchDarklyAPIKey
	err = datadog.Unmarshal(data, &obj.LaunchDarklyAPIKey)
	if err == nil {
		if obj.LaunchDarklyAPIKey != nil && obj.LaunchDarklyAPIKey.UnparsedObject == nil {
			jsonLaunchDarklyAPIKey, _ := datadog.Marshal(obj.LaunchDarklyAPIKey)
			if string(jsonLaunchDarklyAPIKey) == "{}" { // empty struct
				obj.LaunchDarklyAPIKey = nil
			} else {
				match++
			}
		} else {
			obj.LaunchDarklyAPIKey = nil
		}
	} else {
		obj.LaunchDarklyAPIKey = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.LaunchDarklyAPIKey = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj LaunchDarklyCredentials) MarshalJSON() ([]byte, error) {
	if obj.LaunchDarklyAPIKey != nil {
		return datadog.Marshal(&obj.LaunchDarklyAPIKey)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *LaunchDarklyCredentials) GetActualInstance() interface{} {
	if obj.LaunchDarklyAPIKey != nil {
		return obj.LaunchDarklyAPIKey
	}

	// all schemas are nil
	return nil
}
