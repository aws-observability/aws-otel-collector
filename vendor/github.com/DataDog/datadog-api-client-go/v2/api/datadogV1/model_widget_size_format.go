// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetSizeFormat Size of the widget.
type WidgetSizeFormat string

// List of WidgetSizeFormat.
const (
	WIDGETSIZEFORMAT_SMALL  WidgetSizeFormat = "small"
	WIDGETSIZEFORMAT_MEDIUM WidgetSizeFormat = "medium"
	WIDGETSIZEFORMAT_LARGE  WidgetSizeFormat = "large"
)

var allowedWidgetSizeFormatEnumValues = []WidgetSizeFormat{
	WIDGETSIZEFORMAT_SMALL,
	WIDGETSIZEFORMAT_MEDIUM,
	WIDGETSIZEFORMAT_LARGE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WidgetSizeFormat) GetAllowedValues() []WidgetSizeFormat {
	return allowedWidgetSizeFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WidgetSizeFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WidgetSizeFormat(value)
	return nil
}

// NewWidgetSizeFormatFromValue returns a pointer to a valid WidgetSizeFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWidgetSizeFormatFromValue(v string) (*WidgetSizeFormat, error) {
	ev := WidgetSizeFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WidgetSizeFormat: valid values are %v", v, allowedWidgetSizeFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WidgetSizeFormat) IsValid() bool {
	for _, existing := range allowedWidgetSizeFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WidgetSizeFormat value.
func (v WidgetSizeFormat) Ptr() *WidgetSizeFormat {
	return &v
}
