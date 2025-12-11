// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AzureCredentialsUpdate - The definition of the `AzureCredentialsUpdate` object.
type AzureCredentialsUpdate struct {
	AzureTenantUpdate *AzureTenantUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// AzureTenantUpdateAsAzureCredentialsUpdate is a convenience function that returns AzureTenantUpdate wrapped in AzureCredentialsUpdate.
func AzureTenantUpdateAsAzureCredentialsUpdate(v *AzureTenantUpdate) AzureCredentialsUpdate {
	return AzureCredentialsUpdate{AzureTenantUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *AzureCredentialsUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AzureTenantUpdate
	err = datadog.Unmarshal(data, &obj.AzureTenantUpdate)
	if err == nil {
		if obj.AzureTenantUpdate != nil && obj.AzureTenantUpdate.UnparsedObject == nil {
			jsonAzureTenantUpdate, _ := datadog.Marshal(obj.AzureTenantUpdate)
			if string(jsonAzureTenantUpdate) == "{}" { // empty struct
				obj.AzureTenantUpdate = nil
			} else {
				match++
			}
		} else {
			obj.AzureTenantUpdate = nil
		}
	} else {
		obj.AzureTenantUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.AzureTenantUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj AzureCredentialsUpdate) MarshalJSON() ([]byte, error) {
	if obj.AzureTenantUpdate != nil {
		return datadog.Marshal(&obj.AzureTenantUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *AzureCredentialsUpdate) GetActualInstance() interface{} {
	if obj.AzureTenantUpdate != nil {
		return obj.AzureTenantUpdate
	}

	// all schemas are nil
	return nil
}
