// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ComponentGridType The grid component type.
type ComponentGridType string

// List of ComponentGridType.
const (
	COMPONENTGRIDTYPE_GRID ComponentGridType = "grid"
)

var allowedComponentGridTypeEnumValues = []ComponentGridType{
	COMPONENTGRIDTYPE_GRID,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ComponentGridType) GetAllowedValues() []ComponentGridType {
	return allowedComponentGridTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ComponentGridType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ComponentGridType(value)
	return nil
}

// NewComponentGridTypeFromValue returns a pointer to a valid ComponentGridType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewComponentGridTypeFromValue(v string) (*ComponentGridType, error) {
	ev := ComponentGridType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ComponentGridType: valid values are %v", v, allowedComponentGridTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ComponentGridType) IsValid() bool {
	for _, existing := range allowedComponentGridTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ComponentGridType value.
func (v ComponentGridType) Ptr() *ComponentGridType {
	return &v
}
