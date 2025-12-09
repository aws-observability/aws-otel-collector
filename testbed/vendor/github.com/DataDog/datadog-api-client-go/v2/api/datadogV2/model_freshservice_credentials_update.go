// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FreshserviceCredentialsUpdate - The definition of the `FreshserviceCredentialsUpdate` object.
type FreshserviceCredentialsUpdate struct {
	FreshserviceAPIKeyUpdate *FreshserviceAPIKeyUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// FreshserviceAPIKeyUpdateAsFreshserviceCredentialsUpdate is a convenience function that returns FreshserviceAPIKeyUpdate wrapped in FreshserviceCredentialsUpdate.
func FreshserviceAPIKeyUpdateAsFreshserviceCredentialsUpdate(v *FreshserviceAPIKeyUpdate) FreshserviceCredentialsUpdate {
	return FreshserviceCredentialsUpdate{FreshserviceAPIKeyUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *FreshserviceCredentialsUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into FreshserviceAPIKeyUpdate
	err = datadog.Unmarshal(data, &obj.FreshserviceAPIKeyUpdate)
	if err == nil {
		if obj.FreshserviceAPIKeyUpdate != nil && obj.FreshserviceAPIKeyUpdate.UnparsedObject == nil {
			jsonFreshserviceAPIKeyUpdate, _ := datadog.Marshal(obj.FreshserviceAPIKeyUpdate)
			if string(jsonFreshserviceAPIKeyUpdate) == "{}" { // empty struct
				obj.FreshserviceAPIKeyUpdate = nil
			} else {
				match++
			}
		} else {
			obj.FreshserviceAPIKeyUpdate = nil
		}
	} else {
		obj.FreshserviceAPIKeyUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.FreshserviceAPIKeyUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj FreshserviceCredentialsUpdate) MarshalJSON() ([]byte, error) {
	if obj.FreshserviceAPIKeyUpdate != nil {
		return datadog.Marshal(&obj.FreshserviceAPIKeyUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *FreshserviceCredentialsUpdate) GetActualInstance() interface{} {
	if obj.FreshserviceAPIKeyUpdate != nil {
		return obj.FreshserviceAPIKeyUpdate
	}

	// all schemas are nil
	return nil
}
