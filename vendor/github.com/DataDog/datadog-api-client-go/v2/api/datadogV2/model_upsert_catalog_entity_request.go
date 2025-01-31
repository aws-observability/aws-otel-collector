// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpsertCatalogEntityRequest - Create or update entity request.
type UpsertCatalogEntityRequest struct {
	EntityV3  *EntityV3
	EntityRaw *string

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// EntityV3AsUpsertCatalogEntityRequest is a convenience function that returns EntityV3 wrapped in UpsertCatalogEntityRequest.
func EntityV3AsUpsertCatalogEntityRequest(v *EntityV3) UpsertCatalogEntityRequest {
	return UpsertCatalogEntityRequest{EntityV3: v}
}

// EntityRawAsUpsertCatalogEntityRequest is a convenience function that returns string wrapped in UpsertCatalogEntityRequest.
func EntityRawAsUpsertCatalogEntityRequest(v *string) UpsertCatalogEntityRequest {
	return UpsertCatalogEntityRequest{EntityRaw: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *UpsertCatalogEntityRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into EntityV3
	err = datadog.Unmarshal(data, &obj.EntityV3)
	if err == nil {
		if obj.EntityV3 != nil && obj.EntityV3.UnparsedObject == nil {
			jsonEntityV3, _ := datadog.Marshal(obj.EntityV3)
			if string(jsonEntityV3) == "{}" && string(data) != "{}" { // empty struct
				obj.EntityV3 = nil
			} else {
				match++
			}
		} else {
			obj.EntityV3 = nil
		}
	} else {
		obj.EntityV3 = nil
	}

	// try to unmarshal data into EntityRaw
	err = datadog.Unmarshal(data, &obj.EntityRaw)
	if err == nil {
		if obj.EntityRaw != nil {
			jsonEntityRaw, _ := datadog.Marshal(obj.EntityRaw)
			if string(jsonEntityRaw) == "{}" { // empty struct
				obj.EntityRaw = nil
			} else {
				match++
			}
		} else {
			obj.EntityRaw = nil
		}
	} else {
		obj.EntityRaw = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.EntityV3 = nil
		obj.EntityRaw = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj UpsertCatalogEntityRequest) MarshalJSON() ([]byte, error) {
	if obj.EntityV3 != nil {
		return datadog.Marshal(&obj.EntityV3)
	}

	if obj.EntityRaw != nil {
		return datadog.Marshal(&obj.EntityRaw)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *UpsertCatalogEntityRequest) GetActualInstance() interface{} {
	if obj.EntityV3 != nil {
		return obj.EntityV3
	}

	if obj.EntityRaw != nil {
		return obj.EntityRaw
	}

	// all schemas are nil
	return nil
}
