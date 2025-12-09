// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AsanaCredentials - The definition of the `AsanaCredentials` object.
type AsanaCredentials struct {
	AsanaAccessToken *AsanaAccessToken

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// AsanaAccessTokenAsAsanaCredentials is a convenience function that returns AsanaAccessToken wrapped in AsanaCredentials.
func AsanaAccessTokenAsAsanaCredentials(v *AsanaAccessToken) AsanaCredentials {
	return AsanaCredentials{AsanaAccessToken: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *AsanaCredentials) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AsanaAccessToken
	err = datadog.Unmarshal(data, &obj.AsanaAccessToken)
	if err == nil {
		if obj.AsanaAccessToken != nil && obj.AsanaAccessToken.UnparsedObject == nil {
			jsonAsanaAccessToken, _ := datadog.Marshal(obj.AsanaAccessToken)
			if string(jsonAsanaAccessToken) == "{}" { // empty struct
				obj.AsanaAccessToken = nil
			} else {
				match++
			}
		} else {
			obj.AsanaAccessToken = nil
		}
	} else {
		obj.AsanaAccessToken = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.AsanaAccessToken = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj AsanaCredentials) MarshalJSON() ([]byte, error) {
	if obj.AsanaAccessToken != nil {
		return datadog.Marshal(&obj.AsanaAccessToken)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *AsanaCredentials) GetActualInstance() interface{} {
	if obj.AsanaAccessToken != nil {
		return obj.AsanaAccessToken
	}

	// all schemas are nil
	return nil
}
