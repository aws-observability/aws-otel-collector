// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ModelLabFacetType The type of facet for filtering Model Lab runs.
type ModelLabFacetType string

// List of ModelLabFacetType.
const (
	MODELLABFACETTYPE_PARAMETER ModelLabFacetType = "parameter"
	MODELLABFACETTYPE_ATTRIBUTE ModelLabFacetType = "attribute"
	MODELLABFACETTYPE_TAG       ModelLabFacetType = "tag"
	MODELLABFACETTYPE_METRIC    ModelLabFacetType = "metric"
)

var allowedModelLabFacetTypeEnumValues = []ModelLabFacetType{
	MODELLABFACETTYPE_PARAMETER,
	MODELLABFACETTYPE_ATTRIBUTE,
	MODELLABFACETTYPE_TAG,
	MODELLABFACETTYPE_METRIC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ModelLabFacetType) GetAllowedValues() []ModelLabFacetType {
	return allowedModelLabFacetTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ModelLabFacetType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ModelLabFacetType(value)
	return nil
}

// NewModelLabFacetTypeFromValue returns a pointer to a valid ModelLabFacetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewModelLabFacetTypeFromValue(v string) (*ModelLabFacetType, error) {
	ev := ModelLabFacetType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ModelLabFacetType: valid values are %v", v, allowedModelLabFacetTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ModelLabFacetType) IsValid() bool {
	for _, existing := range allowedModelLabFacetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ModelLabFacetType value.
func (v ModelLabFacetType) Ptr() *ModelLabFacetType {
	return &v
}
