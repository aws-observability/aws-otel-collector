// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityV3APISpecInterface - The API definition.
type EntityV3APISpecInterface struct {
	EntityV3APISpecInterfaceFileRef    *EntityV3APISpecInterfaceFileRef
	EntityV3APISpecInterfaceDefinition *EntityV3APISpecInterfaceDefinition

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// EntityV3APISpecInterfaceFileRefAsEntityV3APISpecInterface is a convenience function that returns EntityV3APISpecInterfaceFileRef wrapped in EntityV3APISpecInterface.
func EntityV3APISpecInterfaceFileRefAsEntityV3APISpecInterface(v *EntityV3APISpecInterfaceFileRef) EntityV3APISpecInterface {
	return EntityV3APISpecInterface{EntityV3APISpecInterfaceFileRef: v}
}

// EntityV3APISpecInterfaceDefinitionAsEntityV3APISpecInterface is a convenience function that returns EntityV3APISpecInterfaceDefinition wrapped in EntityV3APISpecInterface.
func EntityV3APISpecInterfaceDefinitionAsEntityV3APISpecInterface(v *EntityV3APISpecInterfaceDefinition) EntityV3APISpecInterface {
	return EntityV3APISpecInterface{EntityV3APISpecInterfaceDefinition: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *EntityV3APISpecInterface) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into EntityV3APISpecInterfaceFileRef
	err = datadog.Unmarshal(data, &obj.EntityV3APISpecInterfaceFileRef)
	if err == nil {
		if obj.EntityV3APISpecInterfaceFileRef != nil && obj.EntityV3APISpecInterfaceFileRef.UnparsedObject == nil {
			jsonEntityV3APISpecInterfaceFileRef, _ := datadog.Marshal(obj.EntityV3APISpecInterfaceFileRef)
			if string(jsonEntityV3APISpecInterfaceFileRef) == "{}" && string(data) != "{}" { // empty struct
				obj.EntityV3APISpecInterfaceFileRef = nil
			} else {
				match++
			}
		} else {
			obj.EntityV3APISpecInterfaceFileRef = nil
		}
	} else {
		obj.EntityV3APISpecInterfaceFileRef = nil
	}

	// try to unmarshal data into EntityV3APISpecInterfaceDefinition
	err = datadog.Unmarshal(data, &obj.EntityV3APISpecInterfaceDefinition)
	if err == nil {
		if obj.EntityV3APISpecInterfaceDefinition != nil && obj.EntityV3APISpecInterfaceDefinition.UnparsedObject == nil {
			jsonEntityV3APISpecInterfaceDefinition, _ := datadog.Marshal(obj.EntityV3APISpecInterfaceDefinition)
			if string(jsonEntityV3APISpecInterfaceDefinition) == "{}" && string(data) != "{}" { // empty struct
				obj.EntityV3APISpecInterfaceDefinition = nil
			} else {
				match++
			}
		} else {
			obj.EntityV3APISpecInterfaceDefinition = nil
		}
	} else {
		obj.EntityV3APISpecInterfaceDefinition = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.EntityV3APISpecInterfaceFileRef = nil
		obj.EntityV3APISpecInterfaceDefinition = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj EntityV3APISpecInterface) MarshalJSON() ([]byte, error) {
	if obj.EntityV3APISpecInterfaceFileRef != nil {
		return datadog.Marshal(&obj.EntityV3APISpecInterfaceFileRef)
	}

	if obj.EntityV3APISpecInterfaceDefinition != nil {
		return datadog.Marshal(&obj.EntityV3APISpecInterfaceDefinition)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *EntityV3APISpecInterface) GetActualInstance() interface{} {
	if obj.EntityV3APISpecInterfaceFileRef != nil {
		return obj.EntityV3APISpecInterfaceFileRef
	}

	if obj.EntityV3APISpecInterfaceDefinition != nil {
		return obj.EntityV3APISpecInterfaceDefinition
	}

	// all schemas are nil
	return nil
}
