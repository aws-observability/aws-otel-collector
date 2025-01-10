// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ContainerItem - Possible Container models.
type ContainerItem struct {
	Container      *Container
	ContainerGroup *ContainerGroup

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ContainerAsContainerItem is a convenience function that returns Container wrapped in ContainerItem.
func ContainerAsContainerItem(v *Container) ContainerItem {
	return ContainerItem{Container: v}
}

// ContainerGroupAsContainerItem is a convenience function that returns ContainerGroup wrapped in ContainerItem.
func ContainerGroupAsContainerItem(v *ContainerGroup) ContainerItem {
	return ContainerItem{ContainerGroup: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ContainerItem) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Container
	err = datadog.Unmarshal(data, &obj.Container)
	if err == nil {
		if obj.Container != nil && obj.Container.UnparsedObject == nil {
			jsonContainer, _ := datadog.Marshal(obj.Container)
			if string(jsonContainer) == "{}" && string(data) != "{}" { // empty struct
				obj.Container = nil
			} else {
				match++
			}
		} else {
			obj.Container = nil
		}
	} else {
		obj.Container = nil
	}

	// try to unmarshal data into ContainerGroup
	err = datadog.Unmarshal(data, &obj.ContainerGroup)
	if err == nil {
		if obj.ContainerGroup != nil && obj.ContainerGroup.UnparsedObject == nil {
			jsonContainerGroup, _ := datadog.Marshal(obj.ContainerGroup)
			if string(jsonContainerGroup) == "{}" && string(data) != "{}" { // empty struct
				obj.ContainerGroup = nil
			} else {
				match++
			}
		} else {
			obj.ContainerGroup = nil
		}
	} else {
		obj.ContainerGroup = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.Container = nil
		obj.ContainerGroup = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ContainerItem) MarshalJSON() ([]byte, error) {
	if obj.Container != nil {
		return datadog.Marshal(&obj.Container)
	}

	if obj.ContainerGroup != nil {
		return datadog.Marshal(&obj.ContainerGroup)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ContainerItem) GetActualInstance() interface{} {
	if obj.Container != nil {
		return obj.Container
	}

	if obj.ContainerGroup != nil {
		return obj.ContainerGroup
	}

	// all schemas are nil
	return nil
}
