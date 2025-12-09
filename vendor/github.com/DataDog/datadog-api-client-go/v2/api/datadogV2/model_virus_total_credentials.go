// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// VirusTotalCredentials - The definition of the `VirusTotalCredentials` object.
type VirusTotalCredentials struct {
	VirusTotalAPIKey *VirusTotalAPIKey

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// VirusTotalAPIKeyAsVirusTotalCredentials is a convenience function that returns VirusTotalAPIKey wrapped in VirusTotalCredentials.
func VirusTotalAPIKeyAsVirusTotalCredentials(v *VirusTotalAPIKey) VirusTotalCredentials {
	return VirusTotalCredentials{VirusTotalAPIKey: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *VirusTotalCredentials) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into VirusTotalAPIKey
	err = datadog.Unmarshal(data, &obj.VirusTotalAPIKey)
	if err == nil {
		if obj.VirusTotalAPIKey != nil && obj.VirusTotalAPIKey.UnparsedObject == nil {
			jsonVirusTotalAPIKey, _ := datadog.Marshal(obj.VirusTotalAPIKey)
			if string(jsonVirusTotalAPIKey) == "{}" { // empty struct
				obj.VirusTotalAPIKey = nil
			} else {
				match++
			}
		} else {
			obj.VirusTotalAPIKey = nil
		}
	} else {
		obj.VirusTotalAPIKey = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.VirusTotalAPIKey = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj VirusTotalCredentials) MarshalJSON() ([]byte, error) {
	if obj.VirusTotalAPIKey != nil {
		return datadog.Marshal(&obj.VirusTotalAPIKey)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *VirusTotalCredentials) GetActualInstance() interface{} {
	if obj.VirusTotalAPIKey != nil {
		return obj.VirusTotalAPIKey
	}

	// all schemas are nil
	return nil
}
