// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceNowCredentials - The definition of the `ServiceNowCredentials` object.
type ServiceNowCredentials struct {
	ServiceNowBasicAuth *ServiceNowBasicAuth

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ServiceNowBasicAuthAsServiceNowCredentials is a convenience function that returns ServiceNowBasicAuth wrapped in ServiceNowCredentials.
func ServiceNowBasicAuthAsServiceNowCredentials(v *ServiceNowBasicAuth) ServiceNowCredentials {
	return ServiceNowCredentials{ServiceNowBasicAuth: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ServiceNowCredentials) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ServiceNowBasicAuth
	err = datadog.Unmarshal(data, &obj.ServiceNowBasicAuth)
	if err == nil {
		if obj.ServiceNowBasicAuth != nil && obj.ServiceNowBasicAuth.UnparsedObject == nil {
			jsonServiceNowBasicAuth, _ := datadog.Marshal(obj.ServiceNowBasicAuth)
			if string(jsonServiceNowBasicAuth) == "{}" { // empty struct
				obj.ServiceNowBasicAuth = nil
			} else {
				match++
			}
		} else {
			obj.ServiceNowBasicAuth = nil
		}
	} else {
		obj.ServiceNowBasicAuth = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ServiceNowBasicAuth = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ServiceNowCredentials) MarshalJSON() ([]byte, error) {
	if obj.ServiceNowBasicAuth != nil {
		return datadog.Marshal(&obj.ServiceNowBasicAuth)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ServiceNowCredentials) GetActualInstance() interface{} {
	if obj.ServiceNowBasicAuth != nil {
		return obj.ServiceNowBasicAuth
	}

	// all schemas are nil
	return nil
}
