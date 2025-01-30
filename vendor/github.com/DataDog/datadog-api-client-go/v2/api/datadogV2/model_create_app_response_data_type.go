// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateAppResponseDataType The definition of `CreateAppResponseDataType` object.
type CreateAppResponseDataType string

// List of CreateAppResponseDataType.
const (
	CREATEAPPRESPONSEDATATYPE_APPDEFINITIONS CreateAppResponseDataType = "appDefinitions"
)

var allowedCreateAppResponseDataTypeEnumValues = []CreateAppResponseDataType{
	CREATEAPPRESPONSEDATATYPE_APPDEFINITIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CreateAppResponseDataType) GetAllowedValues() []CreateAppResponseDataType {
	return allowedCreateAppResponseDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CreateAppResponseDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CreateAppResponseDataType(value)
	return nil
}

// NewCreateAppResponseDataTypeFromValue returns a pointer to a valid CreateAppResponseDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCreateAppResponseDataTypeFromValue(v string) (*CreateAppResponseDataType, error) {
	ev := CreateAppResponseDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CreateAppResponseDataType: valid values are %v", v, allowedCreateAppResponseDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CreateAppResponseDataType) IsValid() bool {
	for _, existing := range allowedCreateAppResponseDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreateAppResponseDataType value.
func (v CreateAppResponseDataType) Ptr() *CreateAppResponseDataType {
	return &v
}
