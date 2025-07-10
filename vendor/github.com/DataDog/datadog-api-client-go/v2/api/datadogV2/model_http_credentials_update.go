// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HTTPCredentialsUpdate - The definition of `HTTPCredentialsUpdate` object.
type HTTPCredentialsUpdate struct {
	HTTPTokenAuthUpdate *HTTPTokenAuthUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// HTTPTokenAuthUpdateAsHTTPCredentialsUpdate is a convenience function that returns HTTPTokenAuthUpdate wrapped in HTTPCredentialsUpdate.
func HTTPTokenAuthUpdateAsHTTPCredentialsUpdate(v *HTTPTokenAuthUpdate) HTTPCredentialsUpdate {
	return HTTPCredentialsUpdate{HTTPTokenAuthUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *HTTPCredentialsUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into HTTPTokenAuthUpdate
	err = datadog.Unmarshal(data, &obj.HTTPTokenAuthUpdate)
	if err == nil {
		if obj.HTTPTokenAuthUpdate != nil && obj.HTTPTokenAuthUpdate.UnparsedObject == nil {
			jsonHTTPTokenAuthUpdate, _ := datadog.Marshal(obj.HTTPTokenAuthUpdate)
			if string(jsonHTTPTokenAuthUpdate) == "{}" { // empty struct
				obj.HTTPTokenAuthUpdate = nil
			} else {
				match++
			}
		} else {
			obj.HTTPTokenAuthUpdate = nil
		}
	} else {
		obj.HTTPTokenAuthUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.HTTPTokenAuthUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj HTTPCredentialsUpdate) MarshalJSON() ([]byte, error) {
	if obj.HTTPTokenAuthUpdate != nil {
		return datadog.Marshal(&obj.HTTPTokenAuthUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *HTTPCredentialsUpdate) GetActualInstance() interface{} {
	if obj.HTTPTokenAuthUpdate != nil {
		return obj.HTTPTokenAuthUpdate
	}

	// all schemas are nil
	return nil
}
