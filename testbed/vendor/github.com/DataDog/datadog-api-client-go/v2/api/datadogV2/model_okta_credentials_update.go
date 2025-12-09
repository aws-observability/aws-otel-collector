// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OktaCredentialsUpdate - The definition of the `OktaCredentialsUpdate` object.
type OktaCredentialsUpdate struct {
	OktaAPITokenUpdate *OktaAPITokenUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// OktaAPITokenUpdateAsOktaCredentialsUpdate is a convenience function that returns OktaAPITokenUpdate wrapped in OktaCredentialsUpdate.
func OktaAPITokenUpdateAsOktaCredentialsUpdate(v *OktaAPITokenUpdate) OktaCredentialsUpdate {
	return OktaCredentialsUpdate{OktaAPITokenUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *OktaCredentialsUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into OktaAPITokenUpdate
	err = datadog.Unmarshal(data, &obj.OktaAPITokenUpdate)
	if err == nil {
		if obj.OktaAPITokenUpdate != nil && obj.OktaAPITokenUpdate.UnparsedObject == nil {
			jsonOktaAPITokenUpdate, _ := datadog.Marshal(obj.OktaAPITokenUpdate)
			if string(jsonOktaAPITokenUpdate) == "{}" { // empty struct
				obj.OktaAPITokenUpdate = nil
			} else {
				match++
			}
		} else {
			obj.OktaAPITokenUpdate = nil
		}
	} else {
		obj.OktaAPITokenUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.OktaAPITokenUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj OktaCredentialsUpdate) MarshalJSON() ([]byte, error) {
	if obj.OktaAPITokenUpdate != nil {
		return datadog.Marshal(&obj.OktaAPITokenUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *OktaCredentialsUpdate) GetActualInstance() interface{} {
	if obj.OktaAPITokenUpdate != nil {
		return obj.OktaAPITokenUpdate
	}

	// all schemas are nil
	return nil
}
