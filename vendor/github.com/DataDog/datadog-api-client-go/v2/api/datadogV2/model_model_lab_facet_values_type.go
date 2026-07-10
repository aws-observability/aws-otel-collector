// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ModelLabFacetValuesType The JSON:API type for a facet values resource.
type ModelLabFacetValuesType string

// List of ModelLabFacetValuesType.
const (
	MODELLABFACETVALUESTYPE_FACET_VALUES ModelLabFacetValuesType = "facet_values"
)

var allowedModelLabFacetValuesTypeEnumValues = []ModelLabFacetValuesType{
	MODELLABFACETVALUESTYPE_FACET_VALUES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ModelLabFacetValuesType) GetAllowedValues() []ModelLabFacetValuesType {
	return allowedModelLabFacetValuesTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ModelLabFacetValuesType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ModelLabFacetValuesType(value)
	return nil
}

// NewModelLabFacetValuesTypeFromValue returns a pointer to a valid ModelLabFacetValuesType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewModelLabFacetValuesTypeFromValue(v string) (*ModelLabFacetValuesType, error) {
	ev := ModelLabFacetValuesType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ModelLabFacetValuesType: valid values are %v", v, allowedModelLabFacetValuesTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ModelLabFacetValuesType) IsValid() bool {
	for _, existing := range allowedModelLabFacetValuesTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ModelLabFacetValuesType value.
func (v ModelLabFacetValuesType) Ptr() *ModelLabFacetValuesType {
	return &v
}
