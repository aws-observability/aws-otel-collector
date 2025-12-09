// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatadogCredentialsUpdate - The definition of the `DatadogCredentialsUpdate` object.
type DatadogCredentialsUpdate struct {
	DatadogAPIKeyUpdate *DatadogAPIKeyUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// DatadogAPIKeyUpdateAsDatadogCredentialsUpdate is a convenience function that returns DatadogAPIKeyUpdate wrapped in DatadogCredentialsUpdate.
func DatadogAPIKeyUpdateAsDatadogCredentialsUpdate(v *DatadogAPIKeyUpdate) DatadogCredentialsUpdate {
	return DatadogCredentialsUpdate{DatadogAPIKeyUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *DatadogCredentialsUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DatadogAPIKeyUpdate
	err = datadog.Unmarshal(data, &obj.DatadogAPIKeyUpdate)
	if err == nil {
		if obj.DatadogAPIKeyUpdate != nil && obj.DatadogAPIKeyUpdate.UnparsedObject == nil {
			jsonDatadogAPIKeyUpdate, _ := datadog.Marshal(obj.DatadogAPIKeyUpdate)
			if string(jsonDatadogAPIKeyUpdate) == "{}" { // empty struct
				obj.DatadogAPIKeyUpdate = nil
			} else {
				match++
			}
		} else {
			obj.DatadogAPIKeyUpdate = nil
		}
	} else {
		obj.DatadogAPIKeyUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.DatadogAPIKeyUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj DatadogCredentialsUpdate) MarshalJSON() ([]byte, error) {
	if obj.DatadogAPIKeyUpdate != nil {
		return datadog.Marshal(&obj.DatadogAPIKeyUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *DatadogCredentialsUpdate) GetActualInstance() interface{} {
	if obj.DatadogAPIKeyUpdate != nil {
		return obj.DatadogAPIKeyUpdate
	}

	// all schemas are nil
	return nil
}
