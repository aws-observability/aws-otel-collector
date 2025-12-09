// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ContainerMetaPageType Type of Container pagination.
type ContainerMetaPageType string

// List of ContainerMetaPageType.
const (
	CONTAINERMETAPAGETYPE_CURSOR_LIMIT ContainerMetaPageType = "cursor_limit"
)

var allowedContainerMetaPageTypeEnumValues = []ContainerMetaPageType{
	CONTAINERMETAPAGETYPE_CURSOR_LIMIT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ContainerMetaPageType) GetAllowedValues() []ContainerMetaPageType {
	return allowedContainerMetaPageTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ContainerMetaPageType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ContainerMetaPageType(value)
	return nil
}

// NewContainerMetaPageTypeFromValue returns a pointer to a valid ContainerMetaPageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewContainerMetaPageTypeFromValue(v string) (*ContainerMetaPageType, error) {
	ev := ContainerMetaPageType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ContainerMetaPageType: valid values are %v", v, allowedContainerMetaPageTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ContainerMetaPageType) IsValid() bool {
	for _, existing := range allowedContainerMetaPageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ContainerMetaPageType value.
func (v ContainerMetaPageType) Ptr() *ContainerMetaPageType {
	return &v
}
