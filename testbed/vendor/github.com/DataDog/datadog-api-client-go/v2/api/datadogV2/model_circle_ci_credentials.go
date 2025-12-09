// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CircleCICredentials - The definition of the `CircleCICredentials` object.
type CircleCICredentials struct {
	CircleCIAPIKey *CircleCIAPIKey

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// CircleCIAPIKeyAsCircleCICredentials is a convenience function that returns CircleCIAPIKey wrapped in CircleCICredentials.
func CircleCIAPIKeyAsCircleCICredentials(v *CircleCIAPIKey) CircleCICredentials {
	return CircleCICredentials{CircleCIAPIKey: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *CircleCICredentials) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CircleCIAPIKey
	err = datadog.Unmarshal(data, &obj.CircleCIAPIKey)
	if err == nil {
		if obj.CircleCIAPIKey != nil && obj.CircleCIAPIKey.UnparsedObject == nil {
			jsonCircleCIAPIKey, _ := datadog.Marshal(obj.CircleCIAPIKey)
			if string(jsonCircleCIAPIKey) == "{}" { // empty struct
				obj.CircleCIAPIKey = nil
			} else {
				match++
			}
		} else {
			obj.CircleCIAPIKey = nil
		}
	} else {
		obj.CircleCIAPIKey = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.CircleCIAPIKey = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj CircleCICredentials) MarshalJSON() ([]byte, error) {
	if obj.CircleCIAPIKey != nil {
		return datadog.Marshal(&obj.CircleCIAPIKey)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *CircleCICredentials) GetActualInstance() interface{} {
	if obj.CircleCIAPIKey != nil {
		return obj.CircleCIAPIKey
	}

	// all schemas are nil
	return nil
}
