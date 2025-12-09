// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScaRequestDataType
type ScaRequestDataType string

// List of ScaRequestDataType.
const (
	SCAREQUESTDATATYPE_SCAREQUESTS ScaRequestDataType = "scarequests"
)

var allowedScaRequestDataTypeEnumValues = []ScaRequestDataType{
	SCAREQUESTDATATYPE_SCAREQUESTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScaRequestDataType) GetAllowedValues() []ScaRequestDataType {
	return allowedScaRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScaRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScaRequestDataType(value)
	return nil
}

// NewScaRequestDataTypeFromValue returns a pointer to a valid ScaRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScaRequestDataTypeFromValue(v string) (*ScaRequestDataType, error) {
	ev := ScaRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScaRequestDataType: valid values are %v", v, allowedScaRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScaRequestDataType) IsValid() bool {
	for _, existing := range allowedScaRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScaRequestDataType value.
func (v ScaRequestDataType) Ptr() *ScaRequestDataType {
	return &v
}
