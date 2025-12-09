// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GCPCredentialsUpdate - The definition of the `GCPCredentialsUpdate` object.
type GCPCredentialsUpdate struct {
	GCPServiceAccountUpdate *GCPServiceAccountUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// GCPServiceAccountUpdateAsGCPCredentialsUpdate is a convenience function that returns GCPServiceAccountUpdate wrapped in GCPCredentialsUpdate.
func GCPServiceAccountUpdateAsGCPCredentialsUpdate(v *GCPServiceAccountUpdate) GCPCredentialsUpdate {
	return GCPCredentialsUpdate{GCPServiceAccountUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *GCPCredentialsUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GCPServiceAccountUpdate
	err = datadog.Unmarshal(data, &obj.GCPServiceAccountUpdate)
	if err == nil {
		if obj.GCPServiceAccountUpdate != nil && obj.GCPServiceAccountUpdate.UnparsedObject == nil {
			jsonGCPServiceAccountUpdate, _ := datadog.Marshal(obj.GCPServiceAccountUpdate)
			if string(jsonGCPServiceAccountUpdate) == "{}" { // empty struct
				obj.GCPServiceAccountUpdate = nil
			} else {
				match++
			}
		} else {
			obj.GCPServiceAccountUpdate = nil
		}
	} else {
		obj.GCPServiceAccountUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.GCPServiceAccountUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj GCPCredentialsUpdate) MarshalJSON() ([]byte, error) {
	if obj.GCPServiceAccountUpdate != nil {
		return datadog.Marshal(&obj.GCPServiceAccountUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *GCPCredentialsUpdate) GetActualInstance() interface{} {
	if obj.GCPServiceAccountUpdate != nil {
		return obj.GCPServiceAccountUpdate
	}

	// all schemas are nil
	return nil
}
