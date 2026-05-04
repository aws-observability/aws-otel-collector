// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetStyleOrderBy How to order series in timeseries visualizations.
// - `tags`: Order series alphabetically by tag name (default behavior)
// - `values`: Order series by their current metric values (typically descending)
type WidgetStyleOrderBy string

// List of WidgetStyleOrderBy.
const (
	WIDGETSTYLEORDERBY_TAGS   WidgetStyleOrderBy = "tags"
	WIDGETSTYLEORDERBY_VALUES WidgetStyleOrderBy = "values"
)

var allowedWidgetStyleOrderByEnumValues = []WidgetStyleOrderBy{
	WIDGETSTYLEORDERBY_TAGS,
	WIDGETSTYLEORDERBY_VALUES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WidgetStyleOrderBy) GetAllowedValues() []WidgetStyleOrderBy {
	return allowedWidgetStyleOrderByEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WidgetStyleOrderBy) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WidgetStyleOrderBy(value)
	return nil
}

// NewWidgetStyleOrderByFromValue returns a pointer to a valid WidgetStyleOrderBy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWidgetStyleOrderByFromValue(v string) (*WidgetStyleOrderBy, error) {
	ev := WidgetStyleOrderBy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WidgetStyleOrderBy: valid values are %v", v, allowedWidgetStyleOrderByEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WidgetStyleOrderBy) IsValid() bool {
	for _, existing := range allowedWidgetStyleOrderByEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WidgetStyleOrderBy value.
func (v WidgetStyleOrderBy) Ptr() *WidgetStyleOrderBy {
	return &v
}
