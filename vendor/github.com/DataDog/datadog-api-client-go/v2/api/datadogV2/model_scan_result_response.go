// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScanResultResponse - The raw scan result document produced by the SCA processor.
// The contents reflect the vulnerabilities and metadata produced for the libraries
// submitted in the original scan request.
type ScanResultResponse struct {
	AnyValueObject map[string]interface{}

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// AnyValueObjectAsScanResultResponse is a convenience function that returns map[string]interface{} wrapped in ScanResultResponse.
func AnyValueObjectAsScanResultResponse(v map[string]interface{}) ScanResultResponse {
	return ScanResultResponse{AnyValueObject: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ScanResultResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AnyValueObject
	err = datadog.Unmarshal(data, &obj.AnyValueObject)
	if err == nil {
		if obj.AnyValueObject != nil {
			jsonAnyValueObject, _ := datadog.Marshal(obj.AnyValueObject)
			if string(jsonAnyValueObject) == "{}" && string(data) != "{}" { // empty struct
				obj.AnyValueObject = nil
			} else {
				match++
			}
		} else {
			obj.AnyValueObject = nil
		}
	} else {
		obj.AnyValueObject = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.AnyValueObject = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ScanResultResponse) MarshalJSON() ([]byte, error) {
	if obj.AnyValueObject != nil {
		return datadog.Marshal(&obj.AnyValueObject)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ScanResultResponse) GetActualInstance() interface{} {
	if obj.AnyValueObject != nil {
		return obj.AnyValueObject
	}

	// all schemas are nil
	return nil
}
