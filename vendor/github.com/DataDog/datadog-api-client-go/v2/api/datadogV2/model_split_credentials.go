// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SplitCredentials - The definition of the `SplitCredentials` object.
type SplitCredentials struct {
	SplitAPIKey *SplitAPIKey

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SplitAPIKeyAsSplitCredentials is a convenience function that returns SplitAPIKey wrapped in SplitCredentials.
func SplitAPIKeyAsSplitCredentials(v *SplitAPIKey) SplitCredentials {
	return SplitCredentials{SplitAPIKey: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *SplitCredentials) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SplitAPIKey
	err = datadog.Unmarshal(data, &obj.SplitAPIKey)
	if err == nil {
		if obj.SplitAPIKey != nil && obj.SplitAPIKey.UnparsedObject == nil {
			jsonSplitAPIKey, _ := datadog.Marshal(obj.SplitAPIKey)
			if string(jsonSplitAPIKey) == "{}" { // empty struct
				obj.SplitAPIKey = nil
			} else {
				match++
			}
		} else {
			obj.SplitAPIKey = nil
		}
	} else {
		obj.SplitAPIKey = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.SplitAPIKey = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj SplitCredentials) MarshalJSON() ([]byte, error) {
	if obj.SplitAPIKey != nil {
		return datadog.Marshal(&obj.SplitAPIKey)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *SplitCredentials) GetActualInstance() interface{} {
	if obj.SplitAPIKey != nil {
		return obj.SplitAPIKey
	}

	// all schemas are nil
	return nil
}
