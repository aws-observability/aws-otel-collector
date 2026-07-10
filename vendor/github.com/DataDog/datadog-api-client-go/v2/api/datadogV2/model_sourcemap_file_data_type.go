// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SourcemapFileDataType The resource type for source map file objects.
type SourcemapFileDataType string

// List of SourcemapFileDataType.
const (
	SOURCEMAPFILEDATATYPE_SOURCEMAP_FILES SourcemapFileDataType = "sourcemap_files"
)

var allowedSourcemapFileDataTypeEnumValues = []SourcemapFileDataType{
	SOURCEMAPFILEDATATYPE_SOURCEMAP_FILES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SourcemapFileDataType) GetAllowedValues() []SourcemapFileDataType {
	return allowedSourcemapFileDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SourcemapFileDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SourcemapFileDataType(value)
	return nil
}

// NewSourcemapFileDataTypeFromValue returns a pointer to a valid SourcemapFileDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSourcemapFileDataTypeFromValue(v string) (*SourcemapFileDataType, error) {
	ev := SourcemapFileDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SourcemapFileDataType: valid values are %v", v, allowedSourcemapFileDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SourcemapFileDataType) IsValid() bool {
	for _, existing := range allowedSourcemapFileDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SourcemapFileDataType value.
func (v SourcemapFileDataType) Ptr() *SourcemapFileDataType {
	return &v
}
