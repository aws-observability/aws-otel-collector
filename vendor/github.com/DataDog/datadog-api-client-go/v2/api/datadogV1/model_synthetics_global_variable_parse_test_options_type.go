// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsGlobalVariableParseTestOptionsType Type of value to extract from a test for a Synthetic global variable.
type SyntheticsGlobalVariableParseTestOptionsType string

// List of SyntheticsGlobalVariableParseTestOptionsType.
const (
	SYNTHETICSGLOBALVARIABLEPARSETESTOPTIONSTYPE_HTTP_BODY        SyntheticsGlobalVariableParseTestOptionsType = "http_body"
	SYNTHETICSGLOBALVARIABLEPARSETESTOPTIONSTYPE_HTTP_HEADER      SyntheticsGlobalVariableParseTestOptionsType = "http_header"
	SYNTHETICSGLOBALVARIABLEPARSETESTOPTIONSTYPE_HTTP_STATUS_CODE SyntheticsGlobalVariableParseTestOptionsType = "http_status_code"
	SYNTHETICSGLOBALVARIABLEPARSETESTOPTIONSTYPE_LOCAL_VARIABLE   SyntheticsGlobalVariableParseTestOptionsType = "local_variable"
)

var allowedSyntheticsGlobalVariableParseTestOptionsTypeEnumValues = []SyntheticsGlobalVariableParseTestOptionsType{
	SYNTHETICSGLOBALVARIABLEPARSETESTOPTIONSTYPE_HTTP_BODY,
	SYNTHETICSGLOBALVARIABLEPARSETESTOPTIONSTYPE_HTTP_HEADER,
	SYNTHETICSGLOBALVARIABLEPARSETESTOPTIONSTYPE_HTTP_STATUS_CODE,
	SYNTHETICSGLOBALVARIABLEPARSETESTOPTIONSTYPE_LOCAL_VARIABLE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsGlobalVariableParseTestOptionsType) GetAllowedValues() []SyntheticsGlobalVariableParseTestOptionsType {
	return allowedSyntheticsGlobalVariableParseTestOptionsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsGlobalVariableParseTestOptionsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsGlobalVariableParseTestOptionsType(value)
	return nil
}

// NewSyntheticsGlobalVariableParseTestOptionsTypeFromValue returns a pointer to a valid SyntheticsGlobalVariableParseTestOptionsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsGlobalVariableParseTestOptionsTypeFromValue(v string) (*SyntheticsGlobalVariableParseTestOptionsType, error) {
	ev := SyntheticsGlobalVariableParseTestOptionsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsGlobalVariableParseTestOptionsType: valid values are %v", v, allowedSyntheticsGlobalVariableParseTestOptionsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsGlobalVariableParseTestOptionsType) IsValid() bool {
	for _, existing := range allowedSyntheticsGlobalVariableParseTestOptionsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsGlobalVariableParseTestOptionsType value.
func (v SyntheticsGlobalVariableParseTestOptionsType) Ptr() *SyntheticsGlobalVariableParseTestOptionsType {
	return &v
}
