// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableWidgetCellDisplayMode Define a display mode for the table cell.
type TableWidgetCellDisplayMode string

// List of TableWidgetCellDisplayMode.
const (
	TABLEWIDGETCELLDISPLAYMODE_NUMBER TableWidgetCellDisplayMode = "number"
	TABLEWIDGETCELLDISPLAYMODE_BAR    TableWidgetCellDisplayMode = "bar"
	TABLEWIDGETCELLDISPLAYMODE_TREND  TableWidgetCellDisplayMode = "trend"
)

var allowedTableWidgetCellDisplayModeEnumValues = []TableWidgetCellDisplayMode{
	TABLEWIDGETCELLDISPLAYMODE_NUMBER,
	TABLEWIDGETCELLDISPLAYMODE_BAR,
	TABLEWIDGETCELLDISPLAYMODE_TREND,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TableWidgetCellDisplayMode) GetAllowedValues() []TableWidgetCellDisplayMode {
	return allowedTableWidgetCellDisplayModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TableWidgetCellDisplayMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TableWidgetCellDisplayMode(value)
	return nil
}

// NewTableWidgetCellDisplayModeFromValue returns a pointer to a valid TableWidgetCellDisplayMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTableWidgetCellDisplayModeFromValue(v string) (*TableWidgetCellDisplayMode, error) {
	ev := TableWidgetCellDisplayMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TableWidgetCellDisplayMode: valid values are %v", v, allowedTableWidgetCellDisplayModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TableWidgetCellDisplayMode) IsValid() bool {
	for _, existing := range allowedTableWidgetCellDisplayModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TableWidgetCellDisplayMode value.
func (v TableWidgetCellDisplayMode) Ptr() *TableWidgetCellDisplayMode {
	return &v
}
