// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ModelLabProjectFacetType The type of facet for filtering Model Lab projects.
type ModelLabProjectFacetType string

// List of ModelLabProjectFacetType.
const (
	MODELLABPROJECTFACETTYPE_TAG ModelLabProjectFacetType = "tag"
)

var allowedModelLabProjectFacetTypeEnumValues = []ModelLabProjectFacetType{
	MODELLABPROJECTFACETTYPE_TAG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ModelLabProjectFacetType) GetAllowedValues() []ModelLabProjectFacetType {
	return allowedModelLabProjectFacetTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ModelLabProjectFacetType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ModelLabProjectFacetType(value)
	return nil
}

// NewModelLabProjectFacetTypeFromValue returns a pointer to a valid ModelLabProjectFacetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewModelLabProjectFacetTypeFromValue(v string) (*ModelLabProjectFacetType, error) {
	ev := ModelLabProjectFacetType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ModelLabProjectFacetType: valid values are %v", v, allowedModelLabProjectFacetTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ModelLabProjectFacetType) IsValid() bool {
	for _, existing := range allowedModelLabProjectFacetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ModelLabProjectFacetType value.
func (v ModelLabProjectFacetType) Ptr() *ModelLabProjectFacetType {
	return &v
}
