// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LaunchDarklyCredentialsUpdate - The definition of the `LaunchDarklyCredentialsUpdate` object.
type LaunchDarklyCredentialsUpdate struct {
	LaunchDarklyAPIKeyUpdate *LaunchDarklyAPIKeyUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// LaunchDarklyAPIKeyUpdateAsLaunchDarklyCredentialsUpdate is a convenience function that returns LaunchDarklyAPIKeyUpdate wrapped in LaunchDarklyCredentialsUpdate.
func LaunchDarklyAPIKeyUpdateAsLaunchDarklyCredentialsUpdate(v *LaunchDarklyAPIKeyUpdate) LaunchDarklyCredentialsUpdate {
	return LaunchDarklyCredentialsUpdate{LaunchDarklyAPIKeyUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *LaunchDarklyCredentialsUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into LaunchDarklyAPIKeyUpdate
	err = datadog.Unmarshal(data, &obj.LaunchDarklyAPIKeyUpdate)
	if err == nil {
		if obj.LaunchDarklyAPIKeyUpdate != nil && obj.LaunchDarklyAPIKeyUpdate.UnparsedObject == nil {
			jsonLaunchDarklyAPIKeyUpdate, _ := datadog.Marshal(obj.LaunchDarklyAPIKeyUpdate)
			if string(jsonLaunchDarklyAPIKeyUpdate) == "{}" { // empty struct
				obj.LaunchDarklyAPIKeyUpdate = nil
			} else {
				match++
			}
		} else {
			obj.LaunchDarklyAPIKeyUpdate = nil
		}
	} else {
		obj.LaunchDarklyAPIKeyUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.LaunchDarklyAPIKeyUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj LaunchDarklyCredentialsUpdate) MarshalJSON() ([]byte, error) {
	if obj.LaunchDarklyAPIKeyUpdate != nil {
		return datadog.Marshal(&obj.LaunchDarklyAPIKeyUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *LaunchDarklyCredentialsUpdate) GetActualInstance() interface{} {
	if obj.LaunchDarklyAPIKeyUpdate != nil {
		return obj.LaunchDarklyAPIKeyUpdate
	}

	// all schemas are nil
	return nil
}
