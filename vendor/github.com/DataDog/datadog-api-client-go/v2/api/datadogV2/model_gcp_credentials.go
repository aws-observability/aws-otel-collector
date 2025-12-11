// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GCPCredentials - The definition of the `GCPCredentials` object.
type GCPCredentials struct {
	GCPServiceAccount *GCPServiceAccount

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// GCPServiceAccountAsGCPCredentials is a convenience function that returns GCPServiceAccount wrapped in GCPCredentials.
func GCPServiceAccountAsGCPCredentials(v *GCPServiceAccount) GCPCredentials {
	return GCPCredentials{GCPServiceAccount: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *GCPCredentials) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GCPServiceAccount
	err = datadog.Unmarshal(data, &obj.GCPServiceAccount)
	if err == nil {
		if obj.GCPServiceAccount != nil && obj.GCPServiceAccount.UnparsedObject == nil {
			jsonGCPServiceAccount, _ := datadog.Marshal(obj.GCPServiceAccount)
			if string(jsonGCPServiceAccount) == "{}" { // empty struct
				obj.GCPServiceAccount = nil
			} else {
				match++
			}
		} else {
			obj.GCPServiceAccount = nil
		}
	} else {
		obj.GCPServiceAccount = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.GCPServiceAccount = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj GCPCredentials) MarshalJSON() ([]byte, error) {
	if obj.GCPServiceAccount != nil {
		return datadog.Marshal(&obj.GCPServiceAccount)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *GCPCredentials) GetActualInstance() interface{} {
	if obj.GCPServiceAccount != nil {
		return obj.GCPServiceAccount
	}

	// all schemas are nil
	return nil
}
