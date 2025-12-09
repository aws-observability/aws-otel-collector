// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GreyNoiseCredentialsUpdate - The definition of the `GreyNoiseCredentialsUpdate` object.
type GreyNoiseCredentialsUpdate struct {
	GreyNoiseAPIKeyUpdate *GreyNoiseAPIKeyUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// GreyNoiseAPIKeyUpdateAsGreyNoiseCredentialsUpdate is a convenience function that returns GreyNoiseAPIKeyUpdate wrapped in GreyNoiseCredentialsUpdate.
func GreyNoiseAPIKeyUpdateAsGreyNoiseCredentialsUpdate(v *GreyNoiseAPIKeyUpdate) GreyNoiseCredentialsUpdate {
	return GreyNoiseCredentialsUpdate{GreyNoiseAPIKeyUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *GreyNoiseCredentialsUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GreyNoiseAPIKeyUpdate
	err = datadog.Unmarshal(data, &obj.GreyNoiseAPIKeyUpdate)
	if err == nil {
		if obj.GreyNoiseAPIKeyUpdate != nil && obj.GreyNoiseAPIKeyUpdate.UnparsedObject == nil {
			jsonGreyNoiseAPIKeyUpdate, _ := datadog.Marshal(obj.GreyNoiseAPIKeyUpdate)
			if string(jsonGreyNoiseAPIKeyUpdate) == "{}" { // empty struct
				obj.GreyNoiseAPIKeyUpdate = nil
			} else {
				match++
			}
		} else {
			obj.GreyNoiseAPIKeyUpdate = nil
		}
	} else {
		obj.GreyNoiseAPIKeyUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.GreyNoiseAPIKeyUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj GreyNoiseCredentialsUpdate) MarshalJSON() ([]byte, error) {
	if obj.GreyNoiseAPIKeyUpdate != nil {
		return datadog.Marshal(&obj.GreyNoiseAPIKeyUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *GreyNoiseCredentialsUpdate) GetActualInstance() interface{} {
	if obj.GreyNoiseAPIKeyUpdate != nil {
		return obj.GreyNoiseAPIKeyUpdate
	}

	// all schemas are nil
	return nil
}
