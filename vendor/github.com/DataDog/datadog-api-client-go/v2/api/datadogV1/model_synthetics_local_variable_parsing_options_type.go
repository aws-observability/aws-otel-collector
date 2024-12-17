// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsLocalVariableParsingOptionsType Property of the Synthetic Test Response to extract into a local variable.
type SyntheticsLocalVariableParsingOptionsType string

// List of SyntheticsLocalVariableParsingOptionsType.
const (
	SYNTHETICSLOCALVARIABLEPARSINGOPTIONSTYPE_GRPC_MESSAGE     SyntheticsLocalVariableParsingOptionsType = "grpc_message"
	SYNTHETICSLOCALVARIABLEPARSINGOPTIONSTYPE_GRPC_METADATA    SyntheticsLocalVariableParsingOptionsType = "grpc_metadata"
	SYNTHETICSLOCALVARIABLEPARSINGOPTIONSTYPE_HTTP_BODY        SyntheticsLocalVariableParsingOptionsType = "http_body"
	SYNTHETICSLOCALVARIABLEPARSINGOPTIONSTYPE_HTTP_HEADER      SyntheticsLocalVariableParsingOptionsType = "http_header"
	SYNTHETICSLOCALVARIABLEPARSINGOPTIONSTYPE_HTTP_STATUS_CODE SyntheticsLocalVariableParsingOptionsType = "http_status_code"
)

var allowedSyntheticsLocalVariableParsingOptionsTypeEnumValues = []SyntheticsLocalVariableParsingOptionsType{
	SYNTHETICSLOCALVARIABLEPARSINGOPTIONSTYPE_GRPC_MESSAGE,
	SYNTHETICSLOCALVARIABLEPARSINGOPTIONSTYPE_GRPC_METADATA,
	SYNTHETICSLOCALVARIABLEPARSINGOPTIONSTYPE_HTTP_BODY,
	SYNTHETICSLOCALVARIABLEPARSINGOPTIONSTYPE_HTTP_HEADER,
	SYNTHETICSLOCALVARIABLEPARSINGOPTIONSTYPE_HTTP_STATUS_CODE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsLocalVariableParsingOptionsType) GetAllowedValues() []SyntheticsLocalVariableParsingOptionsType {
	return allowedSyntheticsLocalVariableParsingOptionsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsLocalVariableParsingOptionsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsLocalVariableParsingOptionsType(value)
	return nil
}

// NewSyntheticsLocalVariableParsingOptionsTypeFromValue returns a pointer to a valid SyntheticsLocalVariableParsingOptionsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsLocalVariableParsingOptionsTypeFromValue(v string) (*SyntheticsLocalVariableParsingOptionsType, error) {
	ev := SyntheticsLocalVariableParsingOptionsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsLocalVariableParsingOptionsType: valid values are %v", v, allowedSyntheticsLocalVariableParsingOptionsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsLocalVariableParsingOptionsType) IsValid() bool {
	for _, existing := range allowedSyntheticsLocalVariableParsingOptionsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsLocalVariableParsingOptionsType value.
func (v SyntheticsLocalVariableParsingOptionsType) Ptr() *SyntheticsLocalVariableParsingOptionsType {
	return &v
}
