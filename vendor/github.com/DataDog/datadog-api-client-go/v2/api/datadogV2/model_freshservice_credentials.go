// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FreshserviceCredentials - The definition of the `FreshserviceCredentials` object.
type FreshserviceCredentials struct {
	FreshserviceAPIKey *FreshserviceAPIKey

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// FreshserviceAPIKeyAsFreshserviceCredentials is a convenience function that returns FreshserviceAPIKey wrapped in FreshserviceCredentials.
func FreshserviceAPIKeyAsFreshserviceCredentials(v *FreshserviceAPIKey) FreshserviceCredentials {
	return FreshserviceCredentials{FreshserviceAPIKey: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *FreshserviceCredentials) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into FreshserviceAPIKey
	err = datadog.Unmarshal(data, &obj.FreshserviceAPIKey)
	if err == nil {
		if obj.FreshserviceAPIKey != nil && obj.FreshserviceAPIKey.UnparsedObject == nil {
			jsonFreshserviceAPIKey, _ := datadog.Marshal(obj.FreshserviceAPIKey)
			if string(jsonFreshserviceAPIKey) == "{}" { // empty struct
				obj.FreshserviceAPIKey = nil
			} else {
				match++
			}
		} else {
			obj.FreshserviceAPIKey = nil
		}
	} else {
		obj.FreshserviceAPIKey = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.FreshserviceAPIKey = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj FreshserviceCredentials) MarshalJSON() ([]byte, error) {
	if obj.FreshserviceAPIKey != nil {
		return datadog.Marshal(&obj.FreshserviceAPIKey)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *FreshserviceCredentials) GetActualInstance() interface{} {
	if obj.FreshserviceAPIKey != nil {
		return obj.FreshserviceAPIKey
	}

	// all schemas are nil
	return nil
}
