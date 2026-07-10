// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SourcemapDataType The resource type for source map objects.
type SourcemapDataType string

// List of SourcemapDataType.
const (
	SOURCEMAPDATATYPE_SOURCEMAPS SourcemapDataType = "sourcemaps"
)

var allowedSourcemapDataTypeEnumValues = []SourcemapDataType{
	SOURCEMAPDATATYPE_SOURCEMAPS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SourcemapDataType) GetAllowedValues() []SourcemapDataType {
	return allowedSourcemapDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SourcemapDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SourcemapDataType(value)
	return nil
}

// NewSourcemapDataTypeFromValue returns a pointer to a valid SourcemapDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSourcemapDataTypeFromValue(v string) (*SourcemapDataType, error) {
	ev := SourcemapDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SourcemapDataType: valid values are %v", v, allowedSourcemapDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SourcemapDataType) IsValid() bool {
	for _, existing := range allowedSourcemapDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SourcemapDataType value.
func (v SourcemapDataType) Ptr() *SourcemapDataType {
	return &v
}
