// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListEntityCatalogResponseIncludedItem - List entity response included item.
type ListEntityCatalogResponseIncludedItem struct {
	EntityResponseIncludedSchema        *EntityResponseIncludedSchema
	EntityResponseIncludedRawSchema     *EntityResponseIncludedRawSchema
	EntityResponseIncludedRelatedEntity *EntityResponseIncludedRelatedEntity
	EntityResponseIncludedOncall        *EntityResponseIncludedOncall
	EntityResponseIncludedIncident      *EntityResponseIncludedIncident

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// EntityResponseIncludedSchemaAsListEntityCatalogResponseIncludedItem is a convenience function that returns EntityResponseIncludedSchema wrapped in ListEntityCatalogResponseIncludedItem.
func EntityResponseIncludedSchemaAsListEntityCatalogResponseIncludedItem(v *EntityResponseIncludedSchema) ListEntityCatalogResponseIncludedItem {
	return ListEntityCatalogResponseIncludedItem{EntityResponseIncludedSchema: v}
}

// EntityResponseIncludedRawSchemaAsListEntityCatalogResponseIncludedItem is a convenience function that returns EntityResponseIncludedRawSchema wrapped in ListEntityCatalogResponseIncludedItem.
func EntityResponseIncludedRawSchemaAsListEntityCatalogResponseIncludedItem(v *EntityResponseIncludedRawSchema) ListEntityCatalogResponseIncludedItem {
	return ListEntityCatalogResponseIncludedItem{EntityResponseIncludedRawSchema: v}
}

// EntityResponseIncludedRelatedEntityAsListEntityCatalogResponseIncludedItem is a convenience function that returns EntityResponseIncludedRelatedEntity wrapped in ListEntityCatalogResponseIncludedItem.
func EntityResponseIncludedRelatedEntityAsListEntityCatalogResponseIncludedItem(v *EntityResponseIncludedRelatedEntity) ListEntityCatalogResponseIncludedItem {
	return ListEntityCatalogResponseIncludedItem{EntityResponseIncludedRelatedEntity: v}
}

// EntityResponseIncludedOncallAsListEntityCatalogResponseIncludedItem is a convenience function that returns EntityResponseIncludedOncall wrapped in ListEntityCatalogResponseIncludedItem.
func EntityResponseIncludedOncallAsListEntityCatalogResponseIncludedItem(v *EntityResponseIncludedOncall) ListEntityCatalogResponseIncludedItem {
	return ListEntityCatalogResponseIncludedItem{EntityResponseIncludedOncall: v}
}

// EntityResponseIncludedIncidentAsListEntityCatalogResponseIncludedItem is a convenience function that returns EntityResponseIncludedIncident wrapped in ListEntityCatalogResponseIncludedItem.
func EntityResponseIncludedIncidentAsListEntityCatalogResponseIncludedItem(v *EntityResponseIncludedIncident) ListEntityCatalogResponseIncludedItem {
	return ListEntityCatalogResponseIncludedItem{EntityResponseIncludedIncident: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ListEntityCatalogResponseIncludedItem) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into EntityResponseIncludedRawSchema
	err = datadog.Unmarshal(data, &obj.EntityResponseIncludedRawSchema)
	if err == nil {
		if obj.EntityResponseIncludedRawSchema != nil && obj.EntityResponseIncludedRawSchema.UnparsedObject == nil {
			jsonEntityResponseIncludedRawSchema, _ := datadog.Marshal(obj.EntityResponseIncludedRawSchema)
			if string(jsonEntityResponseIncludedRawSchema) == "{}" && string(data) != "{}" { // empty struct
				obj.EntityResponseIncludedRawSchema = nil
			} else {
				match++
			}
		} else {
			obj.EntityResponseIncludedRawSchema = nil
		}
	} else {
		obj.EntityResponseIncludedRawSchema = nil
	}

	// try to unmarshal data into EntityResponseIncludedRelatedEntity
	err = datadog.Unmarshal(data, &obj.EntityResponseIncludedRelatedEntity)
	if err == nil {
		if obj.EntityResponseIncludedRelatedEntity != nil && obj.EntityResponseIncludedRelatedEntity.UnparsedObject == nil {
			jsonEntityResponseIncludedRelatedEntity, _ := datadog.Marshal(obj.EntityResponseIncludedRelatedEntity)
			if string(jsonEntityResponseIncludedRelatedEntity) == "{}" && string(data) != "{}" { // empty struct
				obj.EntityResponseIncludedRelatedEntity = nil
			} else {
				match++
			}
		} else {
			obj.EntityResponseIncludedRelatedEntity = nil
		}
	} else {
		obj.EntityResponseIncludedRelatedEntity = nil
	}

	// try to unmarshal data into EntityResponseIncludedOncall
	err = datadog.Unmarshal(data, &obj.EntityResponseIncludedOncall)
	if err == nil {
		if obj.EntityResponseIncludedOncall != nil && obj.EntityResponseIncludedOncall.UnparsedObject == nil {
			jsonEntityResponseIncludedOncall, _ := datadog.Marshal(obj.EntityResponseIncludedOncall)
			if string(jsonEntityResponseIncludedOncall) == "{}" && string(data) != "{}" { // empty struct
				obj.EntityResponseIncludedOncall = nil
			} else {
				match++
			}
		} else {
			obj.EntityResponseIncludedOncall = nil
		}
	} else {
		obj.EntityResponseIncludedOncall = nil
	}

	// try to unmarshal data into EntityResponseIncludedIncident
	err = datadog.Unmarshal(data, &obj.EntityResponseIncludedIncident)
	if err == nil {
		if obj.EntityResponseIncludedIncident != nil && obj.EntityResponseIncludedIncident.UnparsedObject == nil {
			jsonEntityResponseIncludedIncident, _ := datadog.Marshal(obj.EntityResponseIncludedIncident)
			if string(jsonEntityResponseIncludedIncident) == "{}" && string(data) != "{}" { // empty struct
				obj.EntityResponseIncludedIncident = nil
			} else {
				match++
			}
		} else {
			obj.EntityResponseIncludedIncident = nil
		}
	} else {
		obj.EntityResponseIncludedIncident = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.EntityResponseIncludedSchema = nil
		obj.EntityResponseIncludedRawSchema = nil
		obj.EntityResponseIncludedRelatedEntity = nil
		obj.EntityResponseIncludedOncall = nil
		obj.EntityResponseIncludedIncident = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ListEntityCatalogResponseIncludedItem) MarshalJSON() ([]byte, error) {
	if obj.EntityResponseIncludedSchema != nil {
		return datadog.Marshal(&obj.EntityResponseIncludedSchema)
	}

	if obj.EntityResponseIncludedRawSchema != nil {
		return datadog.Marshal(&obj.EntityResponseIncludedRawSchema)
	}

	if obj.EntityResponseIncludedRelatedEntity != nil {
		return datadog.Marshal(&obj.EntityResponseIncludedRelatedEntity)
	}

	if obj.EntityResponseIncludedOncall != nil {
		return datadog.Marshal(&obj.EntityResponseIncludedOncall)
	}

	if obj.EntityResponseIncludedIncident != nil {
		return datadog.Marshal(&obj.EntityResponseIncludedIncident)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ListEntityCatalogResponseIncludedItem) GetActualInstance() interface{} {
	if obj.EntityResponseIncludedSchema != nil {
		return obj.EntityResponseIncludedSchema
	}

	if obj.EntityResponseIncludedRawSchema != nil {
		return obj.EntityResponseIncludedRawSchema
	}

	if obj.EntityResponseIncludedRelatedEntity != nil {
		return obj.EntityResponseIncludedRelatedEntity
	}

	if obj.EntityResponseIncludedOncall != nil {
		return obj.EntityResponseIncludedOncall
	}

	if obj.EntityResponseIncludedIncident != nil {
		return obj.EntityResponseIncludedIncident
	}

	// all schemas are nil
	return nil
}
