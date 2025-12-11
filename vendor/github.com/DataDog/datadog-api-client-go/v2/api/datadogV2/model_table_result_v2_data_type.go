// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableResultV2DataType Reference table resource type.
type TableResultV2DataType string

// List of TableResultV2DataType.
const (
	TABLERESULTV2DATATYPE_REFERENCE_TABLE TableResultV2DataType = "reference_table"
)

var allowedTableResultV2DataTypeEnumValues = []TableResultV2DataType{
	TABLERESULTV2DATATYPE_REFERENCE_TABLE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TableResultV2DataType) GetAllowedValues() []TableResultV2DataType {
	return allowedTableResultV2DataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TableResultV2DataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TableResultV2DataType(value)
	return nil
}

// NewTableResultV2DataTypeFromValue returns a pointer to a valid TableResultV2DataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTableResultV2DataTypeFromValue(v string) (*TableResultV2DataType, error) {
	ev := TableResultV2DataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TableResultV2DataType: valid values are %v", v, allowedTableResultV2DataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TableResultV2DataType) IsValid() bool {
	for _, existing := range allowedTableResultV2DataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TableResultV2DataType value.
func (v TableResultV2DataType) Ptr() *TableResultV2DataType {
	return &v
}
