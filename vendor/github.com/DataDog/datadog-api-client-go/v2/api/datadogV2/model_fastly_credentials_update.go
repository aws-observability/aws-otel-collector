// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FastlyCredentialsUpdate - The definition of the `FastlyCredentialsUpdate` object.
type FastlyCredentialsUpdate struct {
	FastlyAPIKeyUpdate *FastlyAPIKeyUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// FastlyAPIKeyUpdateAsFastlyCredentialsUpdate is a convenience function that returns FastlyAPIKeyUpdate wrapped in FastlyCredentialsUpdate.
func FastlyAPIKeyUpdateAsFastlyCredentialsUpdate(v *FastlyAPIKeyUpdate) FastlyCredentialsUpdate {
	return FastlyCredentialsUpdate{FastlyAPIKeyUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *FastlyCredentialsUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into FastlyAPIKeyUpdate
	err = datadog.Unmarshal(data, &obj.FastlyAPIKeyUpdate)
	if err == nil {
		if obj.FastlyAPIKeyUpdate != nil && obj.FastlyAPIKeyUpdate.UnparsedObject == nil {
			jsonFastlyAPIKeyUpdate, _ := datadog.Marshal(obj.FastlyAPIKeyUpdate)
			if string(jsonFastlyAPIKeyUpdate) == "{}" { // empty struct
				obj.FastlyAPIKeyUpdate = nil
			} else {
				match++
			}
		} else {
			obj.FastlyAPIKeyUpdate = nil
		}
	} else {
		obj.FastlyAPIKeyUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.FastlyAPIKeyUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj FastlyCredentialsUpdate) MarshalJSON() ([]byte, error) {
	if obj.FastlyAPIKeyUpdate != nil {
		return datadog.Marshal(&obj.FastlyAPIKeyUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *FastlyCredentialsUpdate) GetActualInstance() interface{} {
	if obj.FastlyAPIKeyUpdate != nil {
		return obj.FastlyAPIKeyUpdate
	}

	// all schemas are nil
	return nil
}
