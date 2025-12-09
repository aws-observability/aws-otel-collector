// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceNowCredentialsUpdate - The definition of the `ServiceNowCredentialsUpdate` object.
type ServiceNowCredentialsUpdate struct {
	ServiceNowBasicAuthUpdate *ServiceNowBasicAuthUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ServiceNowBasicAuthUpdateAsServiceNowCredentialsUpdate is a convenience function that returns ServiceNowBasicAuthUpdate wrapped in ServiceNowCredentialsUpdate.
func ServiceNowBasicAuthUpdateAsServiceNowCredentialsUpdate(v *ServiceNowBasicAuthUpdate) ServiceNowCredentialsUpdate {
	return ServiceNowCredentialsUpdate{ServiceNowBasicAuthUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ServiceNowCredentialsUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ServiceNowBasicAuthUpdate
	err = datadog.Unmarshal(data, &obj.ServiceNowBasicAuthUpdate)
	if err == nil {
		if obj.ServiceNowBasicAuthUpdate != nil && obj.ServiceNowBasicAuthUpdate.UnparsedObject == nil {
			jsonServiceNowBasicAuthUpdate, _ := datadog.Marshal(obj.ServiceNowBasicAuthUpdate)
			if string(jsonServiceNowBasicAuthUpdate) == "{}" { // empty struct
				obj.ServiceNowBasicAuthUpdate = nil
			} else {
				match++
			}
		} else {
			obj.ServiceNowBasicAuthUpdate = nil
		}
	} else {
		obj.ServiceNowBasicAuthUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ServiceNowBasicAuthUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ServiceNowCredentialsUpdate) MarshalJSON() ([]byte, error) {
	if obj.ServiceNowBasicAuthUpdate != nil {
		return datadog.Marshal(&obj.ServiceNowBasicAuthUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ServiceNowCredentialsUpdate) GetActualInstance() interface{} {
	if obj.ServiceNowBasicAuthUpdate != nil {
		return obj.ServiceNowBasicAuthUpdate
	}

	// all schemas are nil
	return nil
}
