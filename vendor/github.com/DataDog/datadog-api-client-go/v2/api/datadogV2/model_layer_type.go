// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LayerType Layers resource type.
type LayerType string

// List of LayerType.
const (
	LAYERTYPE_LAYERS LayerType = "layers"
)

var allowedLayerTypeEnumValues = []LayerType{
	LAYERTYPE_LAYERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LayerType) GetAllowedValues() []LayerType {
	return allowedLayerTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LayerType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LayerType(value)
	return nil
}

// NewLayerTypeFromValue returns a pointer to a valid LayerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLayerTypeFromValue(v string) (*LayerType, error) {
	ev := LayerType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LayerType: valid values are %v", v, allowedLayerTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LayerType) IsValid() bool {
	for _, existing := range allowedLayerTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LayerType value.
func (v LayerType) Ptr() *LayerType {
	return &v
}
