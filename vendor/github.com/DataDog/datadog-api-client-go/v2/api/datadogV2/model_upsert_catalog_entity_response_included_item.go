// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpsertCatalogEntityResponseIncludedItem - Upsert entity response included item.
type UpsertCatalogEntityResponseIncludedItem struct {
	EntityResponseIncludedSchema *EntityResponseIncludedSchema

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// EntityResponseIncludedSchemaAsUpsertCatalogEntityResponseIncludedItem is a convenience function that returns EntityResponseIncludedSchema wrapped in UpsertCatalogEntityResponseIncludedItem.
func EntityResponseIncludedSchemaAsUpsertCatalogEntityResponseIncludedItem(v *EntityResponseIncludedSchema) UpsertCatalogEntityResponseIncludedItem {
	return UpsertCatalogEntityResponseIncludedItem{EntityResponseIncludedSchema: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *UpsertCatalogEntityResponseIncludedItem) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into EntityResponseIncludedSchema
	err = datadog.Unmarshal(data, &obj.EntityResponseIncludedSchema)
	if err == nil {
		if obj.EntityResponseIncludedSchema != nil && obj.EntityResponseIncludedSchema.UnparsedObject == nil {
			jsonEntityResponseIncludedSchema, _ := datadog.Marshal(obj.EntityResponseIncludedSchema)
			if string(jsonEntityResponseIncludedSchema) == "{}" && string(data) != "{}" { // empty struct
				obj.EntityResponseIncludedSchema = nil
			} else {
				match++
			}
		} else {
			obj.EntityResponseIncludedSchema = nil
		}
	} else {
		obj.EntityResponseIncludedSchema = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.EntityResponseIncludedSchema = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj UpsertCatalogEntityResponseIncludedItem) MarshalJSON() ([]byte, error) {
	if obj.EntityResponseIncludedSchema != nil {
		return datadog.Marshal(&obj.EntityResponseIncludedSchema)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *UpsertCatalogEntityResponseIncludedItem) GetActualInstance() interface{} {
	if obj.EntityResponseIncludedSchema != nil {
		return obj.EntityResponseIncludedSchema
	}

	// all schemas are nil
	return nil
}
