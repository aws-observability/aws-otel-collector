// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ActionConnectionIntegration - The definition of `ActionConnectionIntegration` object.
type ActionConnectionIntegration struct {
	AWSIntegration  *AWSIntegration
	HTTPIntegration *HTTPIntegration

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// AWSIntegrationAsActionConnectionIntegration is a convenience function that returns AWSIntegration wrapped in ActionConnectionIntegration.
func AWSIntegrationAsActionConnectionIntegration(v *AWSIntegration) ActionConnectionIntegration {
	return ActionConnectionIntegration{AWSIntegration: v}
}

// HTTPIntegrationAsActionConnectionIntegration is a convenience function that returns HTTPIntegration wrapped in ActionConnectionIntegration.
func HTTPIntegrationAsActionConnectionIntegration(v *HTTPIntegration) ActionConnectionIntegration {
	return ActionConnectionIntegration{HTTPIntegration: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ActionConnectionIntegration) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AWSIntegration
	err = datadog.Unmarshal(data, &obj.AWSIntegration)
	if err == nil {
		if obj.AWSIntegration != nil && obj.AWSIntegration.UnparsedObject == nil {
			jsonAWSIntegration, _ := datadog.Marshal(obj.AWSIntegration)
			if string(jsonAWSIntegration) == "{}" { // empty struct
				obj.AWSIntegration = nil
			} else {
				match++
			}
		} else {
			obj.AWSIntegration = nil
		}
	} else {
		obj.AWSIntegration = nil
	}

	// try to unmarshal data into HTTPIntegration
	err = datadog.Unmarshal(data, &obj.HTTPIntegration)
	if err == nil {
		if obj.HTTPIntegration != nil && obj.HTTPIntegration.UnparsedObject == nil {
			jsonHTTPIntegration, _ := datadog.Marshal(obj.HTTPIntegration)
			if string(jsonHTTPIntegration) == "{}" { // empty struct
				obj.HTTPIntegration = nil
			} else {
				match++
			}
		} else {
			obj.HTTPIntegration = nil
		}
	} else {
		obj.HTTPIntegration = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.AWSIntegration = nil
		obj.HTTPIntegration = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ActionConnectionIntegration) MarshalJSON() ([]byte, error) {
	if obj.AWSIntegration != nil {
		return datadog.Marshal(&obj.AWSIntegration)
	}

	if obj.HTTPIntegration != nil {
		return datadog.Marshal(&obj.HTTPIntegration)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ActionConnectionIntegration) GetActualInstance() interface{} {
	if obj.AWSIntegration != nil {
		return obj.AWSIntegration
	}

	if obj.HTTPIntegration != nil {
		return obj.HTTPIntegration
	}

	// all schemas are nil
	return nil
}
