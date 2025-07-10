// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DataTransformType The data transform type.
type DataTransformType string

// List of DataTransformType.
const (
	DATATRANSFORMTYPE_DATATRANSFORM DataTransformType = "dataTransform"
)

var allowedDataTransformTypeEnumValues = []DataTransformType{
	DATATRANSFORMTYPE_DATATRANSFORM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DataTransformType) GetAllowedValues() []DataTransformType {
	return allowedDataTransformTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DataTransformType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DataTransformType(value)
	return nil
}

// NewDataTransformTypeFromValue returns a pointer to a valid DataTransformType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDataTransformTypeFromValue(v string) (*DataTransformType, error) {
	ev := DataTransformType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DataTransformType: valid values are %v", v, allowedDataTransformTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DataTransformType) IsValid() bool {
	for _, existing := range allowedDataTransformTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataTransformType value.
func (v DataTransformType) Ptr() *DataTransformType {
	return &v
}
