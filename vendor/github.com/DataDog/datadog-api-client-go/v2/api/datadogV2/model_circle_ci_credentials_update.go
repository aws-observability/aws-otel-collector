// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CircleCICredentialsUpdate - The definition of the `CircleCICredentialsUpdate` object.
type CircleCICredentialsUpdate struct {
	CircleCIAPIKeyUpdate *CircleCIAPIKeyUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// CircleCIAPIKeyUpdateAsCircleCICredentialsUpdate is a convenience function that returns CircleCIAPIKeyUpdate wrapped in CircleCICredentialsUpdate.
func CircleCIAPIKeyUpdateAsCircleCICredentialsUpdate(v *CircleCIAPIKeyUpdate) CircleCICredentialsUpdate {
	return CircleCICredentialsUpdate{CircleCIAPIKeyUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *CircleCICredentialsUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CircleCIAPIKeyUpdate
	err = datadog.Unmarshal(data, &obj.CircleCIAPIKeyUpdate)
	if err == nil {
		if obj.CircleCIAPIKeyUpdate != nil && obj.CircleCIAPIKeyUpdate.UnparsedObject == nil {
			jsonCircleCIAPIKeyUpdate, _ := datadog.Marshal(obj.CircleCIAPIKeyUpdate)
			if string(jsonCircleCIAPIKeyUpdate) == "{}" { // empty struct
				obj.CircleCIAPIKeyUpdate = nil
			} else {
				match++
			}
		} else {
			obj.CircleCIAPIKeyUpdate = nil
		}
	} else {
		obj.CircleCIAPIKeyUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.CircleCIAPIKeyUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj CircleCICredentialsUpdate) MarshalJSON() ([]byte, error) {
	if obj.CircleCIAPIKeyUpdate != nil {
		return datadog.Marshal(&obj.CircleCIAPIKeyUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *CircleCICredentialsUpdate) GetActualInstance() interface{} {
	if obj.CircleCIAPIKeyUpdate != nil {
		return obj.CircleCIAPIKeyUpdate
	}

	// all schemas are nil
	return nil
}
