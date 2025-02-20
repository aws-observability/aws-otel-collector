// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableWidgetTextFormatReplaceAllType Table widget text format replace all type.
type TableWidgetTextFormatReplaceAllType string

// List of TableWidgetTextFormatReplaceAllType.
const (
	TABLEWIDGETTEXTFORMATREPLACEALLTYPE_ALL TableWidgetTextFormatReplaceAllType = "all"
)

var allowedTableWidgetTextFormatReplaceAllTypeEnumValues = []TableWidgetTextFormatReplaceAllType{
	TABLEWIDGETTEXTFORMATREPLACEALLTYPE_ALL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TableWidgetTextFormatReplaceAllType) GetAllowedValues() []TableWidgetTextFormatReplaceAllType {
	return allowedTableWidgetTextFormatReplaceAllTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TableWidgetTextFormatReplaceAllType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TableWidgetTextFormatReplaceAllType(value)
	return nil
}

// NewTableWidgetTextFormatReplaceAllTypeFromValue returns a pointer to a valid TableWidgetTextFormatReplaceAllType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTableWidgetTextFormatReplaceAllTypeFromValue(v string) (*TableWidgetTextFormatReplaceAllType, error) {
	ev := TableWidgetTextFormatReplaceAllType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TableWidgetTextFormatReplaceAllType: valid values are %v", v, allowedTableWidgetTextFormatReplaceAllTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TableWidgetTextFormatReplaceAllType) IsValid() bool {
	for _, existing := range allowedTableWidgetTextFormatReplaceAllTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TableWidgetTextFormatReplaceAllType value.
func (v TableWidgetTextFormatReplaceAllType) Ptr() *TableWidgetTextFormatReplaceAllType {
	return &v
}
