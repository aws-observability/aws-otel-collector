// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ClickupCredentials - The definition of the `ClickupCredentials` object.
type ClickupCredentials struct {
	ClickupAPIKey *ClickupAPIKey

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ClickupAPIKeyAsClickupCredentials is a convenience function that returns ClickupAPIKey wrapped in ClickupCredentials.
func ClickupAPIKeyAsClickupCredentials(v *ClickupAPIKey) ClickupCredentials {
	return ClickupCredentials{ClickupAPIKey: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ClickupCredentials) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ClickupAPIKey
	err = datadog.Unmarshal(data, &obj.ClickupAPIKey)
	if err == nil {
		if obj.ClickupAPIKey != nil && obj.ClickupAPIKey.UnparsedObject == nil {
			jsonClickupAPIKey, _ := datadog.Marshal(obj.ClickupAPIKey)
			if string(jsonClickupAPIKey) == "{}" { // empty struct
				obj.ClickupAPIKey = nil
			} else {
				match++
			}
		} else {
			obj.ClickupAPIKey = nil
		}
	} else {
		obj.ClickupAPIKey = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ClickupAPIKey = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ClickupCredentials) MarshalJSON() ([]byte, error) {
	if obj.ClickupAPIKey != nil {
		return datadog.Marshal(&obj.ClickupAPIKey)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ClickupCredentials) GetActualInstance() interface{} {
	if obj.ClickupAPIKey != nil {
		return obj.ClickupAPIKey
	}

	// all schemas are nil
	return nil
}
