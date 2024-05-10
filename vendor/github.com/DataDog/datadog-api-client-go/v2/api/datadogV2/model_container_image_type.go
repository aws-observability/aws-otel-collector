// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ContainerImageType Type of Container Image.
type ContainerImageType string

// List of ContainerImageType.
const (
	CONTAINERIMAGETYPE_CONTAINER_IMAGE ContainerImageType = "container_image"
)

var allowedContainerImageTypeEnumValues = []ContainerImageType{
	CONTAINERIMAGETYPE_CONTAINER_IMAGE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ContainerImageType) GetAllowedValues() []ContainerImageType {
	return allowedContainerImageTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ContainerImageType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ContainerImageType(value)
	return nil
}

// NewContainerImageTypeFromValue returns a pointer to a valid ContainerImageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewContainerImageTypeFromValue(v string) (*ContainerImageType, error) {
	ev := ContainerImageType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ContainerImageType: valid values are %v", v, allowedContainerImageTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ContainerImageType) IsValid() bool {
	for _, existing := range allowedContainerImageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ContainerImageType value.
func (v ContainerImageType) Ptr() *ContainerImageType {
	return &v
}
