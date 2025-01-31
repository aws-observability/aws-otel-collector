// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeleteAppResponseDataType The definition of `DeleteAppResponseDataType` object.
type DeleteAppResponseDataType string

// List of DeleteAppResponseDataType.
const (
	DELETEAPPRESPONSEDATATYPE_APPDEFINITIONS DeleteAppResponseDataType = "appDefinitions"
)

var allowedDeleteAppResponseDataTypeEnumValues = []DeleteAppResponseDataType{
	DELETEAPPRESPONSEDATATYPE_APPDEFINITIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DeleteAppResponseDataType) GetAllowedValues() []DeleteAppResponseDataType {
	return allowedDeleteAppResponseDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DeleteAppResponseDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DeleteAppResponseDataType(value)
	return nil
}

// NewDeleteAppResponseDataTypeFromValue returns a pointer to a valid DeleteAppResponseDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDeleteAppResponseDataTypeFromValue(v string) (*DeleteAppResponseDataType, error) {
	ev := DeleteAppResponseDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DeleteAppResponseDataType: valid values are %v", v, allowedDeleteAppResponseDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DeleteAppResponseDataType) IsValid() bool {
	for _, existing := range allowedDeleteAppResponseDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeleteAppResponseDataType value.
func (v DeleteAppResponseDataType) Ptr() *DeleteAppResponseDataType {
	return &v
}
