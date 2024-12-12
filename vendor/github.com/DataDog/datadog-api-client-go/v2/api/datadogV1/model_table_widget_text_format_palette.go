// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableWidgetTextFormatPalette Color-on-color palette to highlight replaced text.
type TableWidgetTextFormatPalette string

// List of TableWidgetTextFormatPalette.
const (
	TABLEWIDGETTEXTFORMATPALETTE_WHITE_ON_RED          TableWidgetTextFormatPalette = "white_on_red"
	TABLEWIDGETTEXTFORMATPALETTE_WHITE_ON_YELLOW       TableWidgetTextFormatPalette = "white_on_yellow"
	TABLEWIDGETTEXTFORMATPALETTE_WHITE_ON_GREEN        TableWidgetTextFormatPalette = "white_on_green"
	TABLEWIDGETTEXTFORMATPALETTE_BLACK_ON_LIGHT_RED    TableWidgetTextFormatPalette = "black_on_light_red"
	TABLEWIDGETTEXTFORMATPALETTE_BLACK_ON_LIGHT_YELLOW TableWidgetTextFormatPalette = "black_on_light_yellow"
	TABLEWIDGETTEXTFORMATPALETTE_BLACK_ON_LIGHT_GREEN  TableWidgetTextFormatPalette = "black_on_light_green"
	TABLEWIDGETTEXTFORMATPALETTE_RED_ON_WHITE          TableWidgetTextFormatPalette = "red_on_white"
	TABLEWIDGETTEXTFORMATPALETTE_YELLOW_ON_WHITE       TableWidgetTextFormatPalette = "yellow_on_white"
	TABLEWIDGETTEXTFORMATPALETTE_GREEN_ON_WHITE        TableWidgetTextFormatPalette = "green_on_white"
	TABLEWIDGETTEXTFORMATPALETTE_CUSTOM_BG             TableWidgetTextFormatPalette = "custom_bg"
	TABLEWIDGETTEXTFORMATPALETTE_CUSTOM_TEXT           TableWidgetTextFormatPalette = "custom_text"
)

var allowedTableWidgetTextFormatPaletteEnumValues = []TableWidgetTextFormatPalette{
	TABLEWIDGETTEXTFORMATPALETTE_WHITE_ON_RED,
	TABLEWIDGETTEXTFORMATPALETTE_WHITE_ON_YELLOW,
	TABLEWIDGETTEXTFORMATPALETTE_WHITE_ON_GREEN,
	TABLEWIDGETTEXTFORMATPALETTE_BLACK_ON_LIGHT_RED,
	TABLEWIDGETTEXTFORMATPALETTE_BLACK_ON_LIGHT_YELLOW,
	TABLEWIDGETTEXTFORMATPALETTE_BLACK_ON_LIGHT_GREEN,
	TABLEWIDGETTEXTFORMATPALETTE_RED_ON_WHITE,
	TABLEWIDGETTEXTFORMATPALETTE_YELLOW_ON_WHITE,
	TABLEWIDGETTEXTFORMATPALETTE_GREEN_ON_WHITE,
	TABLEWIDGETTEXTFORMATPALETTE_CUSTOM_BG,
	TABLEWIDGETTEXTFORMATPALETTE_CUSTOM_TEXT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TableWidgetTextFormatPalette) GetAllowedValues() []TableWidgetTextFormatPalette {
	return allowedTableWidgetTextFormatPaletteEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TableWidgetTextFormatPalette) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TableWidgetTextFormatPalette(value)
	return nil
}

// NewTableWidgetTextFormatPaletteFromValue returns a pointer to a valid TableWidgetTextFormatPalette
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTableWidgetTextFormatPaletteFromValue(v string) (*TableWidgetTextFormatPalette, error) {
	ev := TableWidgetTextFormatPalette(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TableWidgetTextFormatPalette: valid values are %v", v, allowedTableWidgetTextFormatPaletteEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TableWidgetTextFormatPalette) IsValid() bool {
	for _, existing := range allowedTableWidgetTextFormatPaletteEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TableWidgetTextFormatPalette value.
func (v TableWidgetTextFormatPalette) Ptr() *TableWidgetTextFormatPalette {
	return &v
}
