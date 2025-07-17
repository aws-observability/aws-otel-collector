// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpsertCatalogKindRequest - Create or update kind request.
type UpsertCatalogKindRequest struct {
	KindObj *KindObj
	KindRaw *string

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// KindObjAsUpsertCatalogKindRequest is a convenience function that returns KindObj wrapped in UpsertCatalogKindRequest.
func KindObjAsUpsertCatalogKindRequest(v *KindObj) UpsertCatalogKindRequest {
	return UpsertCatalogKindRequest{KindObj: v}
}

// KindRawAsUpsertCatalogKindRequest is a convenience function that returns string wrapped in UpsertCatalogKindRequest.
func KindRawAsUpsertCatalogKindRequest(v *string) UpsertCatalogKindRequest {
	return UpsertCatalogKindRequest{KindRaw: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *UpsertCatalogKindRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into KindObj
	err = datadog.Unmarshal(data, &obj.KindObj)
	if err == nil {
		if obj.KindObj != nil && obj.KindObj.UnparsedObject == nil {
			jsonKindObj, _ := datadog.Marshal(obj.KindObj)
			if string(jsonKindObj) == "{}" { // empty struct
				obj.KindObj = nil
			} else {
				match++
			}
		} else {
			obj.KindObj = nil
		}
	} else {
		obj.KindObj = nil
	}

	// try to unmarshal data into KindRaw
	err = datadog.Unmarshal(data, &obj.KindRaw)
	if err == nil {
		if obj.KindRaw != nil {
			jsonKindRaw, _ := datadog.Marshal(obj.KindRaw)
			if string(jsonKindRaw) == "{}" { // empty struct
				obj.KindRaw = nil
			} else {
				match++
			}
		} else {
			obj.KindRaw = nil
		}
	} else {
		obj.KindRaw = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.KindObj = nil
		obj.KindRaw = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj UpsertCatalogKindRequest) MarshalJSON() ([]byte, error) {
	if obj.KindObj != nil {
		return datadog.Marshal(&obj.KindObj)
	}

	if obj.KindRaw != nil {
		return datadog.Marshal(&obj.KindRaw)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *UpsertCatalogKindRequest) GetActualInstance() interface{} {
	if obj.KindObj != nil {
		return obj.KindObj
	}

	if obj.KindRaw != nil {
		return obj.KindRaw
	}

	// all schemas are nil
	return nil
}
