// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableWidgetTextFormatMatchType Match or compare option.
type TableWidgetTextFormatMatchType string

// List of TableWidgetTextFormatMatchType.
const (
	TABLEWIDGETTEXTFORMATMATCHTYPE_IS               TableWidgetTextFormatMatchType = "is"
	TABLEWIDGETTEXTFORMATMATCHTYPE_IS_NOT           TableWidgetTextFormatMatchType = "is_not"
	TABLEWIDGETTEXTFORMATMATCHTYPE_CONTAINS         TableWidgetTextFormatMatchType = "contains"
	TABLEWIDGETTEXTFORMATMATCHTYPE_DOES_NOT_CONTAIN TableWidgetTextFormatMatchType = "does_not_contain"
	TABLEWIDGETTEXTFORMATMATCHTYPE_STARTS_WITH      TableWidgetTextFormatMatchType = "starts_with"
	TABLEWIDGETTEXTFORMATMATCHTYPE_ENDS_WITH        TableWidgetTextFormatMatchType = "ends_with"
)

var allowedTableWidgetTextFormatMatchTypeEnumValues = []TableWidgetTextFormatMatchType{
	TABLEWIDGETTEXTFORMATMATCHTYPE_IS,
	TABLEWIDGETTEXTFORMATMATCHTYPE_IS_NOT,
	TABLEWIDGETTEXTFORMATMATCHTYPE_CONTAINS,
	TABLEWIDGETTEXTFORMATMATCHTYPE_DOES_NOT_CONTAIN,
	TABLEWIDGETTEXTFORMATMATCHTYPE_STARTS_WITH,
	TABLEWIDGETTEXTFORMATMATCHTYPE_ENDS_WITH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TableWidgetTextFormatMatchType) GetAllowedValues() []TableWidgetTextFormatMatchType {
	return allowedTableWidgetTextFormatMatchTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TableWidgetTextFormatMatchType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TableWidgetTextFormatMatchType(value)
	return nil
}

// NewTableWidgetTextFormatMatchTypeFromValue returns a pointer to a valid TableWidgetTextFormatMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTableWidgetTextFormatMatchTypeFromValue(v string) (*TableWidgetTextFormatMatchType, error) {
	ev := TableWidgetTextFormatMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TableWidgetTextFormatMatchType: valid values are %v", v, allowedTableWidgetTextFormatMatchTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TableWidgetTextFormatMatchType) IsValid() bool {
	for _, existing := range allowedTableWidgetTextFormatMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TableWidgetTextFormatMatchType value.
func (v TableWidgetTextFormatMatchType) Ptr() *TableWidgetTextFormatMatchType {
	return &v
}
