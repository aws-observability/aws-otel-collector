// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ModelLabFacetKeysType The JSON:API type for a facet keys resource.
type ModelLabFacetKeysType string

// List of ModelLabFacetKeysType.
const (
	MODELLABFACETKEYSTYPE_FACET_KEYS ModelLabFacetKeysType = "facet_keys"
)

var allowedModelLabFacetKeysTypeEnumValues = []ModelLabFacetKeysType{
	MODELLABFACETKEYSTYPE_FACET_KEYS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ModelLabFacetKeysType) GetAllowedValues() []ModelLabFacetKeysType {
	return allowedModelLabFacetKeysTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ModelLabFacetKeysType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ModelLabFacetKeysType(value)
	return nil
}

// NewModelLabFacetKeysTypeFromValue returns a pointer to a valid ModelLabFacetKeysType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewModelLabFacetKeysTypeFromValue(v string) (*ModelLabFacetKeysType, error) {
	ev := ModelLabFacetKeysType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ModelLabFacetKeysType: valid values are %v", v, allowedModelLabFacetKeysTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ModelLabFacetKeysType) IsValid() bool {
	for _, existing := range allowedModelLabFacetKeysTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ModelLabFacetKeysType value.
func (v ModelLabFacetKeysType) Ptr() *ModelLabFacetKeysType {
	return &v
}
