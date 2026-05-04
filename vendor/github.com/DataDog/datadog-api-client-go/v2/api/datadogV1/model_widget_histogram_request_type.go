// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetHistogramRequestType Request type for distribution of point values for distribution metrics. Query space aggregator must be `histogram:<metric name>` for points distributions.
type WidgetHistogramRequestType string

// List of WidgetHistogramRequestType.
const (
	WIDGETHISTOGRAMREQUESTTYPE_HISTOGRAM WidgetHistogramRequestType = "histogram"
)

var allowedWidgetHistogramRequestTypeEnumValues = []WidgetHistogramRequestType{
	WIDGETHISTOGRAMREQUESTTYPE_HISTOGRAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WidgetHistogramRequestType) GetAllowedValues() []WidgetHistogramRequestType {
	return allowedWidgetHistogramRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WidgetHistogramRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WidgetHistogramRequestType(value)
	return nil
}

// NewWidgetHistogramRequestTypeFromValue returns a pointer to a valid WidgetHistogramRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWidgetHistogramRequestTypeFromValue(v string) (*WidgetHistogramRequestType, error) {
	ev := WidgetHistogramRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WidgetHistogramRequestType: valid values are %v", v, allowedWidgetHistogramRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WidgetHistogramRequestType) IsValid() bool {
	for _, existing := range allowedWidgetHistogramRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WidgetHistogramRequestType value.
func (v WidgetHistogramRequestType) Ptr() *WidgetHistogramRequestType {
	return &v
}
