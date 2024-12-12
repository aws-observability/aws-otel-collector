// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ContainerImageItem - Possible Container Image models.
type ContainerImageItem struct {
	ContainerImage      *ContainerImage
	ContainerImageGroup *ContainerImageGroup

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ContainerImageAsContainerImageItem is a convenience function that returns ContainerImage wrapped in ContainerImageItem.
func ContainerImageAsContainerImageItem(v *ContainerImage) ContainerImageItem {
	return ContainerImageItem{ContainerImage: v}
}

// ContainerImageGroupAsContainerImageItem is a convenience function that returns ContainerImageGroup wrapped in ContainerImageItem.
func ContainerImageGroupAsContainerImageItem(v *ContainerImageGroup) ContainerImageItem {
	return ContainerImageItem{ContainerImageGroup: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ContainerImageItem) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ContainerImage
	err = datadog.Unmarshal(data, &obj.ContainerImage)
	if err == nil {
		if obj.ContainerImage != nil && obj.ContainerImage.UnparsedObject == nil {
			jsonContainerImage, _ := datadog.Marshal(obj.ContainerImage)
			if string(jsonContainerImage) == "{}" && string(data) != "{}" { // empty struct
				obj.ContainerImage = nil
			} else {
				match++
			}
		} else {
			obj.ContainerImage = nil
		}
	} else {
		obj.ContainerImage = nil
	}

	// try to unmarshal data into ContainerImageGroup
	err = datadog.Unmarshal(data, &obj.ContainerImageGroup)
	if err == nil {
		if obj.ContainerImageGroup != nil && obj.ContainerImageGroup.UnparsedObject == nil {
			jsonContainerImageGroup, _ := datadog.Marshal(obj.ContainerImageGroup)
			if string(jsonContainerImageGroup) == "{}" && string(data) != "{}" { // empty struct
				obj.ContainerImageGroup = nil
			} else {
				match++
			}
		} else {
			obj.ContainerImageGroup = nil
		}
	} else {
		obj.ContainerImageGroup = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ContainerImage = nil
		obj.ContainerImageGroup = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ContainerImageItem) MarshalJSON() ([]byte, error) {
	if obj.ContainerImage != nil {
		return datadog.Marshal(&obj.ContainerImage)
	}

	if obj.ContainerImageGroup != nil {
		return datadog.Marshal(&obj.ContainerImageGroup)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ContainerImageItem) GetActualInstance() interface{} {
	if obj.ContainerImage != nil {
		return obj.ContainerImage
	}

	if obj.ContainerImageGroup != nil {
		return obj.ContainerImageGroup
	}

	// all schemas are nil
	return nil
}
