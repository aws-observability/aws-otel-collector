// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableWidgetDefinitionType Type of the table widget.
type TableWidgetDefinitionType string

// List of TableWidgetDefinitionType.
const (
	TABLEWIDGETDEFINITIONTYPE_QUERY_TABLE TableWidgetDefinitionType = "query_table"
)

var allowedTableWidgetDefinitionTypeEnumValues = []TableWidgetDefinitionType{
	TABLEWIDGETDEFINITIONTYPE_QUERY_TABLE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TableWidgetDefinitionType) GetAllowedValues() []TableWidgetDefinitionType {
	return allowedTableWidgetDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TableWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TableWidgetDefinitionType(value)
	return nil
}

// NewTableWidgetDefinitionTypeFromValue returns a pointer to a valid TableWidgetDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTableWidgetDefinitionTypeFromValue(v string) (*TableWidgetDefinitionType, error) {
	ev := TableWidgetDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TableWidgetDefinitionType: valid values are %v", v, allowedTableWidgetDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TableWidgetDefinitionType) IsValid() bool {
	for _, existing := range allowedTableWidgetDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TableWidgetDefinitionType value.
func (v TableWidgetDefinitionType) Ptr() *TableWidgetDefinitionType {
	return &v
}
