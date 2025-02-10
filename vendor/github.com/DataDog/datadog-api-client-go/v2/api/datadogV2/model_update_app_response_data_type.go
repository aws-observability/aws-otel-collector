// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateAppResponseDataType The definition of `UpdateAppResponseDataType` object.
type UpdateAppResponseDataType string

// List of UpdateAppResponseDataType.
const (
	UPDATEAPPRESPONSEDATATYPE_APPDEFINITIONS UpdateAppResponseDataType = "appDefinitions"
)

var allowedUpdateAppResponseDataTypeEnumValues = []UpdateAppResponseDataType{
	UPDATEAPPRESPONSEDATATYPE_APPDEFINITIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UpdateAppResponseDataType) GetAllowedValues() []UpdateAppResponseDataType {
	return allowedUpdateAppResponseDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UpdateAppResponseDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UpdateAppResponseDataType(value)
	return nil
}

// NewUpdateAppResponseDataTypeFromValue returns a pointer to a valid UpdateAppResponseDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUpdateAppResponseDataTypeFromValue(v string) (*UpdateAppResponseDataType, error) {
	ev := UpdateAppResponseDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UpdateAppResponseDataType: valid values are %v", v, allowedUpdateAppResponseDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UpdateAppResponseDataType) IsValid() bool {
	for _, existing := range allowedUpdateAppResponseDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UpdateAppResponseDataType value.
func (v UpdateAppResponseDataType) Ptr() *UpdateAppResponseDataType {
	return &v
}
