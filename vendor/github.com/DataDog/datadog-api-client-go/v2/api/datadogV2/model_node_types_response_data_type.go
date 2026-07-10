// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NodeTypesResponseDataType Get node types response resource type.
type NodeTypesResponseDataType string

// List of NodeTypesResponseDataType.
const (
	NODETYPESRESPONSEDATATYPE_GET_NODE_TYPES_RESPONSE NodeTypesResponseDataType = "get_node_types_response"
)

var allowedNodeTypesResponseDataTypeEnumValues = []NodeTypesResponseDataType{
	NODETYPESRESPONSEDATATYPE_GET_NODE_TYPES_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *NodeTypesResponseDataType) GetAllowedValues() []NodeTypesResponseDataType {
	return allowedNodeTypesResponseDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NodeTypesResponseDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NodeTypesResponseDataType(value)
	return nil
}

// NewNodeTypesResponseDataTypeFromValue returns a pointer to a valid NodeTypesResponseDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNodeTypesResponseDataTypeFromValue(v string) (*NodeTypesResponseDataType, error) {
	ev := NodeTypesResponseDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NodeTypesResponseDataType: valid values are %v", v, allowedNodeTypesResponseDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NodeTypesResponseDataType) IsValid() bool {
	for _, existing := range allowedNodeTypesResponseDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NodeTypesResponseDataType value.
func (v NodeTypesResponseDataType) Ptr() *NodeTypesResponseDataType {
	return &v
}
