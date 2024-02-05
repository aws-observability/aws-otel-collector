// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetOrderBy What to order by.
type WidgetOrderBy string

// List of WidgetOrderBy.
const (
	WIDGETORDERBY_CHANGE  WidgetOrderBy = "change"
	WIDGETORDERBY_NAME    WidgetOrderBy = "name"
	WIDGETORDERBY_PRESENT WidgetOrderBy = "present"
	WIDGETORDERBY_PAST    WidgetOrderBy = "past"
)

var allowedWidgetOrderByEnumValues = []WidgetOrderBy{
	WIDGETORDERBY_CHANGE,
	WIDGETORDERBY_NAME,
	WIDGETORDERBY_PRESENT,
	WIDGETORDERBY_PAST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WidgetOrderBy) GetAllowedValues() []WidgetOrderBy {
	return allowedWidgetOrderByEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WidgetOrderBy) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WidgetOrderBy(value)
	return nil
}

// NewWidgetOrderByFromValue returns a pointer to a valid WidgetOrderBy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWidgetOrderByFromValue(v string) (*WidgetOrderBy, error) {
	ev := WidgetOrderBy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WidgetOrderBy: valid values are %v", v, allowedWidgetOrderByEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WidgetOrderBy) IsValid() bool {
	for _, existing := range allowedWidgetOrderByEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WidgetOrderBy value.
func (v WidgetOrderBy) Ptr() *WidgetOrderBy {
	return &v
}
