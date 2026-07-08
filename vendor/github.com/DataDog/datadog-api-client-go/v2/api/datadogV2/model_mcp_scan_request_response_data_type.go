// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// McpScanRequestResponseDataType The type identifier for MCP SCA scan request responses.
type McpScanRequestResponseDataType string

// List of McpScanRequestResponseDataType.
const (
	MCPSCANREQUESTRESPONSEDATATYPE_MCPSCANREQUESTRESPONSE McpScanRequestResponseDataType = "mcpscanrequestresponse"
)

var allowedMcpScanRequestResponseDataTypeEnumValues = []McpScanRequestResponseDataType{
	MCPSCANREQUESTRESPONSEDATATYPE_MCPSCANREQUESTRESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *McpScanRequestResponseDataType) GetAllowedValues() []McpScanRequestResponseDataType {
	return allowedMcpScanRequestResponseDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *McpScanRequestResponseDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = McpScanRequestResponseDataType(value)
	return nil
}

// NewMcpScanRequestResponseDataTypeFromValue returns a pointer to a valid McpScanRequestResponseDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMcpScanRequestResponseDataTypeFromValue(v string) (*McpScanRequestResponseDataType, error) {
	ev := McpScanRequestResponseDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for McpScanRequestResponseDataType: valid values are %v", v, allowedMcpScanRequestResponseDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v McpScanRequestResponseDataType) IsValid() bool {
	for _, existing := range allowedMcpScanRequestResponseDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to McpScanRequestResponseDataType value.
func (v McpScanRequestResponseDataType) Ptr() *McpScanRequestResponseDataType {
	return &v
}
