// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SplitCredentialsUpdate - The definition of the `SplitCredentialsUpdate` object.
type SplitCredentialsUpdate struct {
	SplitAPIKeyUpdate *SplitAPIKeyUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// SplitAPIKeyUpdateAsSplitCredentialsUpdate is a convenience function that returns SplitAPIKeyUpdate wrapped in SplitCredentialsUpdate.
func SplitAPIKeyUpdateAsSplitCredentialsUpdate(v *SplitAPIKeyUpdate) SplitCredentialsUpdate {
	return SplitCredentialsUpdate{SplitAPIKeyUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *SplitCredentialsUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SplitAPIKeyUpdate
	err = datadog.Unmarshal(data, &obj.SplitAPIKeyUpdate)
	if err == nil {
		if obj.SplitAPIKeyUpdate != nil && obj.SplitAPIKeyUpdate.UnparsedObject == nil {
			jsonSplitAPIKeyUpdate, _ := datadog.Marshal(obj.SplitAPIKeyUpdate)
			if string(jsonSplitAPIKeyUpdate) == "{}" { // empty struct
				obj.SplitAPIKeyUpdate = nil
			} else {
				match++
			}
		} else {
			obj.SplitAPIKeyUpdate = nil
		}
	} else {
		obj.SplitAPIKeyUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.SplitAPIKeyUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj SplitCredentialsUpdate) MarshalJSON() ([]byte, error) {
	if obj.SplitAPIKeyUpdate != nil {
		return datadog.Marshal(&obj.SplitAPIKeyUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *SplitCredentialsUpdate) GetActualInstance() interface{} {
	if obj.SplitAPIKeyUpdate != nil {
		return obj.SplitAPIKeyUpdate
	}

	// all schemas are nil
	return nil
}
