// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatadogCredentials - The definition of the `DatadogCredentials` object.
type DatadogCredentials struct {
	DatadogAPIKey *DatadogAPIKey

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// DatadogAPIKeyAsDatadogCredentials is a convenience function that returns DatadogAPIKey wrapped in DatadogCredentials.
func DatadogAPIKeyAsDatadogCredentials(v *DatadogAPIKey) DatadogCredentials {
	return DatadogCredentials{DatadogAPIKey: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *DatadogCredentials) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DatadogAPIKey
	err = datadog.Unmarshal(data, &obj.DatadogAPIKey)
	if err == nil {
		if obj.DatadogAPIKey != nil && obj.DatadogAPIKey.UnparsedObject == nil {
			jsonDatadogAPIKey, _ := datadog.Marshal(obj.DatadogAPIKey)
			if string(jsonDatadogAPIKey) == "{}" { // empty struct
				obj.DatadogAPIKey = nil
			} else {
				match++
			}
		} else {
			obj.DatadogAPIKey = nil
		}
	} else {
		obj.DatadogAPIKey = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.DatadogAPIKey = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj DatadogCredentials) MarshalJSON() ([]byte, error) {
	if obj.DatadogAPIKey != nil {
		return datadog.Marshal(&obj.DatadogAPIKey)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *DatadogCredentials) GetActualInstance() interface{} {
	if obj.DatadogAPIKey != nil {
		return obj.DatadogAPIKey
	}

	// all schemas are nil
	return nil
}
