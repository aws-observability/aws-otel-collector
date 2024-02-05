// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ContainerGroupType Type of container group.
type ContainerGroupType string

// List of ContainerGroupType.
const (
	CONTAINERGROUPTYPE_CONTAINER_GROUP ContainerGroupType = "container_group"
)

var allowedContainerGroupTypeEnumValues = []ContainerGroupType{
	CONTAINERGROUPTYPE_CONTAINER_GROUP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ContainerGroupType) GetAllowedValues() []ContainerGroupType {
	return allowedContainerGroupTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ContainerGroupType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ContainerGroupType(value)
	return nil
}

// NewContainerGroupTypeFromValue returns a pointer to a valid ContainerGroupType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewContainerGroupTypeFromValue(v string) (*ContainerGroupType, error) {
	ev := ContainerGroupType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ContainerGroupType: valid values are %v", v, allowedContainerGroupTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ContainerGroupType) IsValid() bool {
	for _, existing := range allowedContainerGroupTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ContainerGroupType value.
func (v ContainerGroupType) Ptr() *ContainerGroupType {
	return &v
}
