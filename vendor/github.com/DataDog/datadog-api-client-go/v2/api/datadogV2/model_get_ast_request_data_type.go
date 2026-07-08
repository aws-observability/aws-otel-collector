// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GetAstRequestDataType Get AST request resource type.
type GetAstRequestDataType string

// List of GetAstRequestDataType.
const (
	GETASTREQUESTDATATYPE_GET_AST_REQUEST GetAstRequestDataType = "get_ast_request"
)

var allowedGetAstRequestDataTypeEnumValues = []GetAstRequestDataType{
	GETASTREQUESTDATATYPE_GET_AST_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GetAstRequestDataType) GetAllowedValues() []GetAstRequestDataType {
	return allowedGetAstRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GetAstRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GetAstRequestDataType(value)
	return nil
}

// NewGetAstRequestDataTypeFromValue returns a pointer to a valid GetAstRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGetAstRequestDataTypeFromValue(v string) (*GetAstRequestDataType, error) {
	ev := GetAstRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GetAstRequestDataType: valid values are %v", v, allowedGetAstRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GetAstRequestDataType) IsValid() bool {
	for _, existing := range allowedGetAstRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GetAstRequestDataType value.
func (v GetAstRequestDataType) Ptr() *GetAstRequestDataType {
	return &v
}
