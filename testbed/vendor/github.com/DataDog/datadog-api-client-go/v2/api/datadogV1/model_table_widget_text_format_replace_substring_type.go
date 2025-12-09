// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableWidgetTextFormatReplaceSubstringType Table widget text format replace sub-string type.
type TableWidgetTextFormatReplaceSubstringType string

// List of TableWidgetTextFormatReplaceSubstringType.
const (
	TABLEWIDGETTEXTFORMATREPLACESUBSTRINGTYPE_SUBSTRING TableWidgetTextFormatReplaceSubstringType = "substring"
)

var allowedTableWidgetTextFormatReplaceSubstringTypeEnumValues = []TableWidgetTextFormatReplaceSubstringType{
	TABLEWIDGETTEXTFORMATREPLACESUBSTRINGTYPE_SUBSTRING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TableWidgetTextFormatReplaceSubstringType) GetAllowedValues() []TableWidgetTextFormatReplaceSubstringType {
	return allowedTableWidgetTextFormatReplaceSubstringTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TableWidgetTextFormatReplaceSubstringType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TableWidgetTextFormatReplaceSubstringType(value)
	return nil
}

// NewTableWidgetTextFormatReplaceSubstringTypeFromValue returns a pointer to a valid TableWidgetTextFormatReplaceSubstringType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTableWidgetTextFormatReplaceSubstringTypeFromValue(v string) (*TableWidgetTextFormatReplaceSubstringType, error) {
	ev := TableWidgetTextFormatReplaceSubstringType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TableWidgetTextFormatReplaceSubstringType: valid values are %v", v, allowedTableWidgetTextFormatReplaceSubstringTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TableWidgetTextFormatReplaceSubstringType) IsValid() bool {
	for _, existing := range allowedTableWidgetTextFormatReplaceSubstringTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TableWidgetTextFormatReplaceSubstringType value.
func (v TableWidgetTextFormatReplaceSubstringType) Ptr() *TableWidgetTextFormatReplaceSubstringType {
	return &v
}
