// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SunburstWidgetLegendInlineAutomaticType Whether to show the legend inline or let it be automatically generated.
type SunburstWidgetLegendInlineAutomaticType string

// List of SunburstWidgetLegendInlineAutomaticType.
const (
	SUNBURSTWIDGETLEGENDINLINEAUTOMATICTYPE_INLINE    SunburstWidgetLegendInlineAutomaticType = "inline"
	SUNBURSTWIDGETLEGENDINLINEAUTOMATICTYPE_AUTOMATIC SunburstWidgetLegendInlineAutomaticType = "automatic"
)

var allowedSunburstWidgetLegendInlineAutomaticTypeEnumValues = []SunburstWidgetLegendInlineAutomaticType{
	SUNBURSTWIDGETLEGENDINLINEAUTOMATICTYPE_INLINE,
	SUNBURSTWIDGETLEGENDINLINEAUTOMATICTYPE_AUTOMATIC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SunburstWidgetLegendInlineAutomaticType) GetAllowedValues() []SunburstWidgetLegendInlineAutomaticType {
	return allowedSunburstWidgetLegendInlineAutomaticTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SunburstWidgetLegendInlineAutomaticType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SunburstWidgetLegendInlineAutomaticType(value)
	return nil
}

// NewSunburstWidgetLegendInlineAutomaticTypeFromValue returns a pointer to a valid SunburstWidgetLegendInlineAutomaticType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSunburstWidgetLegendInlineAutomaticTypeFromValue(v string) (*SunburstWidgetLegendInlineAutomaticType, error) {
	ev := SunburstWidgetLegendInlineAutomaticType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SunburstWidgetLegendInlineAutomaticType: valid values are %v", v, allowedSunburstWidgetLegendInlineAutomaticTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SunburstWidgetLegendInlineAutomaticType) IsValid() bool {
	for _, existing := range allowedSunburstWidgetLegendInlineAutomaticTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SunburstWidgetLegendInlineAutomaticType value.
func (v SunburstWidgetLegendInlineAutomaticType) Ptr() *SunburstWidgetLegendInlineAutomaticType {
	return &v
}
