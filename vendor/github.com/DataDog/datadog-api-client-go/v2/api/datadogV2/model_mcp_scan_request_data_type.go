// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// McpScanRequestDataType The type identifier for MCP SCA scan requests.
type McpScanRequestDataType string

// List of McpScanRequestDataType.
const (
	MCPSCANREQUESTDATATYPE_MCPSCANREQUEST McpScanRequestDataType = "mcpscanrequest"
)

var allowedMcpScanRequestDataTypeEnumValues = []McpScanRequestDataType{
	MCPSCANREQUESTDATATYPE_MCPSCANREQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *McpScanRequestDataType) GetAllowedValues() []McpScanRequestDataType {
	return allowedMcpScanRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *McpScanRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = McpScanRequestDataType(value)
	return nil
}

// NewMcpScanRequestDataTypeFromValue returns a pointer to a valid McpScanRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMcpScanRequestDataTypeFromValue(v string) (*McpScanRequestDataType, error) {
	ev := McpScanRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for McpScanRequestDataType: valid values are %v", v, allowedMcpScanRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v McpScanRequestDataType) IsValid() bool {
	for _, existing := range allowedMcpScanRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to McpScanRequestDataType value.
func (v McpScanRequestDataType) Ptr() *McpScanRequestDataType {
	return &v
}
