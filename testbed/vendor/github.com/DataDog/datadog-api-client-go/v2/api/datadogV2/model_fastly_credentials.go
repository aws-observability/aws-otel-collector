// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FastlyCredentials - The definition of the `FastlyCredentials` object.
type FastlyCredentials struct {
	FastlyAPIKey *FastlyAPIKey

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// FastlyAPIKeyAsFastlyCredentials is a convenience function that returns FastlyAPIKey wrapped in FastlyCredentials.
func FastlyAPIKeyAsFastlyCredentials(v *FastlyAPIKey) FastlyCredentials {
	return FastlyCredentials{FastlyAPIKey: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *FastlyCredentials) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into FastlyAPIKey
	err = datadog.Unmarshal(data, &obj.FastlyAPIKey)
	if err == nil {
		if obj.FastlyAPIKey != nil && obj.FastlyAPIKey.UnparsedObject == nil {
			jsonFastlyAPIKey, _ := datadog.Marshal(obj.FastlyAPIKey)
			if string(jsonFastlyAPIKey) == "{}" { // empty struct
				obj.FastlyAPIKey = nil
			} else {
				match++
			}
		} else {
			obj.FastlyAPIKey = nil
		}
	} else {
		obj.FastlyAPIKey = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.FastlyAPIKey = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj FastlyCredentials) MarshalJSON() ([]byte, error) {
	if obj.FastlyAPIKey != nil {
		return datadog.Marshal(&obj.FastlyAPIKey)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *FastlyCredentials) GetActualInstance() interface{} {
	if obj.FastlyAPIKey != nil {
		return obj.FastlyAPIKey
	}

	// all schemas are nil
	return nil
}
