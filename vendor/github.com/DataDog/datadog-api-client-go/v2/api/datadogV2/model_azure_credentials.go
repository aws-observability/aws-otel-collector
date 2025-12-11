// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AzureCredentials - The definition of the `AzureCredentials` object.
type AzureCredentials struct {
	AzureTenant *AzureTenant

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// AzureTenantAsAzureCredentials is a convenience function that returns AzureTenant wrapped in AzureCredentials.
func AzureTenantAsAzureCredentials(v *AzureTenant) AzureCredentials {
	return AzureCredentials{AzureTenant: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *AzureCredentials) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AzureTenant
	err = datadog.Unmarshal(data, &obj.AzureTenant)
	if err == nil {
		if obj.AzureTenant != nil && obj.AzureTenant.UnparsedObject == nil {
			jsonAzureTenant, _ := datadog.Marshal(obj.AzureTenant)
			if string(jsonAzureTenant) == "{}" { // empty struct
				obj.AzureTenant = nil
			} else {
				match++
			}
		} else {
			obj.AzureTenant = nil
		}
	} else {
		obj.AzureTenant = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.AzureTenant = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj AzureCredentials) MarshalJSON() ([]byte, error) {
	if obj.AzureTenant != nil {
		return datadog.Marshal(&obj.AzureTenant)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *AzureCredentials) GetActualInstance() interface{} {
	if obj.AzureTenant != nil {
		return obj.AzureTenant
	}

	// all schemas are nil
	return nil
}
