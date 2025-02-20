// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GetAppResponseDataType The definition of `GetAppResponseDataType` object.
type GetAppResponseDataType string

// List of GetAppResponseDataType.
const (
	GETAPPRESPONSEDATATYPE_APPDEFINITIONS GetAppResponseDataType = "appDefinitions"
)

var allowedGetAppResponseDataTypeEnumValues = []GetAppResponseDataType{
	GETAPPRESPONSEDATATYPE_APPDEFINITIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GetAppResponseDataType) GetAllowedValues() []GetAppResponseDataType {
	return allowedGetAppResponseDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GetAppResponseDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GetAppResponseDataType(value)
	return nil
}

// NewGetAppResponseDataTypeFromValue returns a pointer to a valid GetAppResponseDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGetAppResponseDataTypeFromValue(v string) (*GetAppResponseDataType, error) {
	ev := GetAppResponseDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GetAppResponseDataType: valid values are %v", v, allowedGetAppResponseDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GetAppResponseDataType) IsValid() bool {
	for _, existing := range allowedGetAppResponseDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GetAppResponseDataType value.
func (v GetAppResponseDataType) Ptr() *GetAppResponseDataType {
	return &v
}
