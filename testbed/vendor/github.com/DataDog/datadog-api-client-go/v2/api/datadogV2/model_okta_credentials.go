// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OktaCredentials - The definition of the `OktaCredentials` object.
type OktaCredentials struct {
	OktaAPIToken *OktaAPIToken

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// OktaAPITokenAsOktaCredentials is a convenience function that returns OktaAPIToken wrapped in OktaCredentials.
func OktaAPITokenAsOktaCredentials(v *OktaAPIToken) OktaCredentials {
	return OktaCredentials{OktaAPIToken: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *OktaCredentials) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into OktaAPIToken
	err = datadog.Unmarshal(data, &obj.OktaAPIToken)
	if err == nil {
		if obj.OktaAPIToken != nil && obj.OktaAPIToken.UnparsedObject == nil {
			jsonOktaAPIToken, _ := datadog.Marshal(obj.OktaAPIToken)
			if string(jsonOktaAPIToken) == "{}" { // empty struct
				obj.OktaAPIToken = nil
			} else {
				match++
			}
		} else {
			obj.OktaAPIToken = nil
		}
	} else {
		obj.OktaAPIToken = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.OktaAPIToken = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj OktaCredentials) MarshalJSON() ([]byte, error) {
	if obj.OktaAPIToken != nil {
		return datadog.Marshal(&obj.OktaAPIToken)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *OktaCredentials) GetActualInstance() interface{} {
	if obj.OktaAPIToken != nil {
		return obj.OktaAPIToken
	}

	// all schemas are nil
	return nil
}
