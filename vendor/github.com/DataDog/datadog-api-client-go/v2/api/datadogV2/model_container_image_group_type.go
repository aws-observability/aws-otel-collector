// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ContainerImageGroupType Type of Container Image Group.
type ContainerImageGroupType string

// List of ContainerImageGroupType.
const (
	CONTAINERIMAGEGROUPTYPE_CONTAINER_IMAGE_GROUP ContainerImageGroupType = "container_image_group"
)

var allowedContainerImageGroupTypeEnumValues = []ContainerImageGroupType{
	CONTAINERIMAGEGROUPTYPE_CONTAINER_IMAGE_GROUP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ContainerImageGroupType) GetAllowedValues() []ContainerImageGroupType {
	return allowedContainerImageGroupTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ContainerImageGroupType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ContainerImageGroupType(value)
	return nil
}

// NewContainerImageGroupTypeFromValue returns a pointer to a valid ContainerImageGroupType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewContainerImageGroupTypeFromValue(v string) (*ContainerImageGroupType, error) {
	ev := ContainerImageGroupType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ContainerImageGroupType: valid values are %v", v, allowedContainerImageGroupTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ContainerImageGroupType) IsValid() bool {
	for _, existing := range allowedContainerImageGroupTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ContainerImageGroupType value.
func (v ContainerImageGroupType) Ptr() *ContainerImageGroupType {
	return &v
}
