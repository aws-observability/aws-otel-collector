// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ClickupCredentialsUpdate - The definition of the `ClickupCredentialsUpdate` object.
type ClickupCredentialsUpdate struct {
	ClickupAPIKeyUpdate *ClickupAPIKeyUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ClickupAPIKeyUpdateAsClickupCredentialsUpdate is a convenience function that returns ClickupAPIKeyUpdate wrapped in ClickupCredentialsUpdate.
func ClickupAPIKeyUpdateAsClickupCredentialsUpdate(v *ClickupAPIKeyUpdate) ClickupCredentialsUpdate {
	return ClickupCredentialsUpdate{ClickupAPIKeyUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ClickupCredentialsUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ClickupAPIKeyUpdate
	err = datadog.Unmarshal(data, &obj.ClickupAPIKeyUpdate)
	if err == nil {
		if obj.ClickupAPIKeyUpdate != nil && obj.ClickupAPIKeyUpdate.UnparsedObject == nil {
			jsonClickupAPIKeyUpdate, _ := datadog.Marshal(obj.ClickupAPIKeyUpdate)
			if string(jsonClickupAPIKeyUpdate) == "{}" { // empty struct
				obj.ClickupAPIKeyUpdate = nil
			} else {
				match++
			}
		} else {
			obj.ClickupAPIKeyUpdate = nil
		}
	} else {
		obj.ClickupAPIKeyUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ClickupAPIKeyUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ClickupCredentialsUpdate) MarshalJSON() ([]byte, error) {
	if obj.ClickupAPIKeyUpdate != nil {
		return datadog.Marshal(&obj.ClickupAPIKeyUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ClickupCredentialsUpdate) GetActualInstance() interface{} {
	if obj.ClickupAPIKeyUpdate != nil {
		return obj.ClickupAPIKeyUpdate
	}

	// all schemas are nil
	return nil
}
