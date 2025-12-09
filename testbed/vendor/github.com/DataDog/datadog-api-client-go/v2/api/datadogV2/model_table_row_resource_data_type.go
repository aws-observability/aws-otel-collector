// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableRowResourceDataType Row resource type.
type TableRowResourceDataType string

// List of TableRowResourceDataType.
const (
	TABLEROWRESOURCEDATATYPE_ROW TableRowResourceDataType = "row"
)

var allowedTableRowResourceDataTypeEnumValues = []TableRowResourceDataType{
	TABLEROWRESOURCEDATATYPE_ROW,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TableRowResourceDataType) GetAllowedValues() []TableRowResourceDataType {
	return allowedTableRowResourceDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TableRowResourceDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TableRowResourceDataType(value)
	return nil
}

// NewTableRowResourceDataTypeFromValue returns a pointer to a valid TableRowResourceDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTableRowResourceDataTypeFromValue(v string) (*TableRowResourceDataType, error) {
	ev := TableRowResourceDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TableRowResourceDataType: valid values are %v", v, allowedTableRowResourceDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TableRowResourceDataType) IsValid() bool {
	for _, existing := range allowedTableRowResourceDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TableRowResourceDataType value.
func (v TableRowResourceDataType) Ptr() *TableRowResourceDataType {
	return &v
}
