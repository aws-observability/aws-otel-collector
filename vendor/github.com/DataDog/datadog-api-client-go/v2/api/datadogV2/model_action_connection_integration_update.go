// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ActionConnectionIntegrationUpdate - The definition of `ActionConnectionIntegrationUpdate` object.
type ActionConnectionIntegrationUpdate struct {
	AWSIntegrationUpdate  *AWSIntegrationUpdate
	HTTPIntegrationUpdate *HTTPIntegrationUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// AWSIntegrationUpdateAsActionConnectionIntegrationUpdate is a convenience function that returns AWSIntegrationUpdate wrapped in ActionConnectionIntegrationUpdate.
func AWSIntegrationUpdateAsActionConnectionIntegrationUpdate(v *AWSIntegrationUpdate) ActionConnectionIntegrationUpdate {
	return ActionConnectionIntegrationUpdate{AWSIntegrationUpdate: v}
}

// HTTPIntegrationUpdateAsActionConnectionIntegrationUpdate is a convenience function that returns HTTPIntegrationUpdate wrapped in ActionConnectionIntegrationUpdate.
func HTTPIntegrationUpdateAsActionConnectionIntegrationUpdate(v *HTTPIntegrationUpdate) ActionConnectionIntegrationUpdate {
	return ActionConnectionIntegrationUpdate{HTTPIntegrationUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ActionConnectionIntegrationUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AWSIntegrationUpdate
	err = datadog.Unmarshal(data, &obj.AWSIntegrationUpdate)
	if err == nil {
		if obj.AWSIntegrationUpdate != nil && obj.AWSIntegrationUpdate.UnparsedObject == nil {
			jsonAWSIntegrationUpdate, _ := datadog.Marshal(obj.AWSIntegrationUpdate)
			if string(jsonAWSIntegrationUpdate) == "{}" { // empty struct
				obj.AWSIntegrationUpdate = nil
			} else {
				match++
			}
		} else {
			obj.AWSIntegrationUpdate = nil
		}
	} else {
		obj.AWSIntegrationUpdate = nil
	}

	// try to unmarshal data into HTTPIntegrationUpdate
	err = datadog.Unmarshal(data, &obj.HTTPIntegrationUpdate)
	if err == nil {
		if obj.HTTPIntegrationUpdate != nil && obj.HTTPIntegrationUpdate.UnparsedObject == nil {
			jsonHTTPIntegrationUpdate, _ := datadog.Marshal(obj.HTTPIntegrationUpdate)
			if string(jsonHTTPIntegrationUpdate) == "{}" { // empty struct
				obj.HTTPIntegrationUpdate = nil
			} else {
				match++
			}
		} else {
			obj.HTTPIntegrationUpdate = nil
		}
	} else {
		obj.HTTPIntegrationUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.AWSIntegrationUpdate = nil
		obj.HTTPIntegrationUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ActionConnectionIntegrationUpdate) MarshalJSON() ([]byte, error) {
	if obj.AWSIntegrationUpdate != nil {
		return datadog.Marshal(&obj.AWSIntegrationUpdate)
	}

	if obj.HTTPIntegrationUpdate != nil {
		return datadog.Marshal(&obj.HTTPIntegrationUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ActionConnectionIntegrationUpdate) GetActualInstance() interface{} {
	if obj.AWSIntegrationUpdate != nil {
		return obj.AWSIntegrationUpdate
	}

	if obj.HTTPIntegrationUpdate != nil {
		return obj.HTTPIntegrationUpdate
	}

	// all schemas are nil
	return nil
}
