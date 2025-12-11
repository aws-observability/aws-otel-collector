// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GreyNoiseCredentials - The definition of the `GreyNoiseCredentials` object.
type GreyNoiseCredentials struct {
	GreyNoiseAPIKey *GreyNoiseAPIKey

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// GreyNoiseAPIKeyAsGreyNoiseCredentials is a convenience function that returns GreyNoiseAPIKey wrapped in GreyNoiseCredentials.
func GreyNoiseAPIKeyAsGreyNoiseCredentials(v *GreyNoiseAPIKey) GreyNoiseCredentials {
	return GreyNoiseCredentials{GreyNoiseAPIKey: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *GreyNoiseCredentials) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GreyNoiseAPIKey
	err = datadog.Unmarshal(data, &obj.GreyNoiseAPIKey)
	if err == nil {
		if obj.GreyNoiseAPIKey != nil && obj.GreyNoiseAPIKey.UnparsedObject == nil {
			jsonGreyNoiseAPIKey, _ := datadog.Marshal(obj.GreyNoiseAPIKey)
			if string(jsonGreyNoiseAPIKey) == "{}" { // empty struct
				obj.GreyNoiseAPIKey = nil
			} else {
				match++
			}
		} else {
			obj.GreyNoiseAPIKey = nil
		}
	} else {
		obj.GreyNoiseAPIKey = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.GreyNoiseAPIKey = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj GreyNoiseCredentials) MarshalJSON() ([]byte, error) {
	if obj.GreyNoiseAPIKey != nil {
		return datadog.Marshal(&obj.GreyNoiseAPIKey)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *GreyNoiseCredentials) GetActualInstance() interface{} {
	if obj.GreyNoiseAPIKey != nil {
		return obj.GreyNoiseAPIKey
	}

	// all schemas are nil
	return nil
}
