// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ContainerType Type of container.
type ContainerType string

// List of ContainerType.
const (
	CONTAINERTYPE_CONTAINER ContainerType = "container"
)

var allowedContainerTypeEnumValues = []ContainerType{
	CONTAINERTYPE_CONTAINER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ContainerType) GetAllowedValues() []ContainerType {
	return allowedContainerTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ContainerType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ContainerType(value)
	return nil
}

// NewContainerTypeFromValue returns a pointer to a valid ContainerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewContainerTypeFromValue(v string) (*ContainerType, error) {
	ev := ContainerType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ContainerType: valid values are %v", v, allowedContainerTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ContainerType) IsValid() bool {
	for _, existing := range allowedContainerTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ContainerType value.
func (v ContainerType) Ptr() *ContainerType {
	return &v
}
