// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GetAstResponseDataType Get AST response resource type.
type GetAstResponseDataType string

// List of GetAstResponseDataType.
const (
	GETASTRESPONSEDATATYPE_GET_AST_RESPONSE GetAstResponseDataType = "get_ast_response"
)

var allowedGetAstResponseDataTypeEnumValues = []GetAstResponseDataType{
	GETASTRESPONSEDATATYPE_GET_AST_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GetAstResponseDataType) GetAllowedValues() []GetAstResponseDataType {
	return allowedGetAstResponseDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GetAstResponseDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GetAstResponseDataType(value)
	return nil
}

// NewGetAstResponseDataTypeFromValue returns a pointer to a valid GetAstResponseDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGetAstResponseDataTypeFromValue(v string) (*GetAstResponseDataType, error) {
	ev := GetAstResponseDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GetAstResponseDataType: valid values are %v", v, allowedGetAstResponseDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GetAstResponseDataType) IsValid() bool {
	for _, existing := range allowedGetAstResponseDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GetAstResponseDataType value.
func (v GetAstResponseDataType) Ptr() *GetAstResponseDataType {
	return &v
}
