// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// VirusTotalCredentialsUpdate - The definition of the `VirusTotalCredentialsUpdate` object.
type VirusTotalCredentialsUpdate struct {
	VirusTotalAPIKeyUpdate *VirusTotalAPIKeyUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// VirusTotalAPIKeyUpdateAsVirusTotalCredentialsUpdate is a convenience function that returns VirusTotalAPIKeyUpdate wrapped in VirusTotalCredentialsUpdate.
func VirusTotalAPIKeyUpdateAsVirusTotalCredentialsUpdate(v *VirusTotalAPIKeyUpdate) VirusTotalCredentialsUpdate {
	return VirusTotalCredentialsUpdate{VirusTotalAPIKeyUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *VirusTotalCredentialsUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into VirusTotalAPIKeyUpdate
	err = datadog.Unmarshal(data, &obj.VirusTotalAPIKeyUpdate)
	if err == nil {
		if obj.VirusTotalAPIKeyUpdate != nil && obj.VirusTotalAPIKeyUpdate.UnparsedObject == nil {
			jsonVirusTotalAPIKeyUpdate, _ := datadog.Marshal(obj.VirusTotalAPIKeyUpdate)
			if string(jsonVirusTotalAPIKeyUpdate) == "{}" { // empty struct
				obj.VirusTotalAPIKeyUpdate = nil
			} else {
				match++
			}
		} else {
			obj.VirusTotalAPIKeyUpdate = nil
		}
	} else {
		obj.VirusTotalAPIKeyUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.VirusTotalAPIKeyUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj VirusTotalCredentialsUpdate) MarshalJSON() ([]byte, error) {
	if obj.VirusTotalAPIKeyUpdate != nil {
		return datadog.Marshal(&obj.VirusTotalAPIKeyUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *VirusTotalCredentialsUpdate) GetActualInstance() interface{} {
	if obj.VirusTotalAPIKeyUpdate != nil {
		return obj.VirusTotalAPIKeyUpdate
	}

	// all schemas are nil
	return nil
}
