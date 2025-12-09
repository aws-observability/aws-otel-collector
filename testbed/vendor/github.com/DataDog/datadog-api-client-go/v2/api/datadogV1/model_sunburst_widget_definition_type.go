// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SunburstWidgetDefinitionType Type of the Sunburst widget.
type SunburstWidgetDefinitionType string

// List of SunburstWidgetDefinitionType.
const (
	SUNBURSTWIDGETDEFINITIONTYPE_SUNBURST SunburstWidgetDefinitionType = "sunburst"
)

var allowedSunburstWidgetDefinitionTypeEnumValues = []SunburstWidgetDefinitionType{
	SUNBURSTWIDGETDEFINITIONTYPE_SUNBURST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SunburstWidgetDefinitionType) GetAllowedValues() []SunburstWidgetDefinitionType {
	return allowedSunburstWidgetDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SunburstWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SunburstWidgetDefinitionType(value)
	return nil
}

// NewSunburstWidgetDefinitionTypeFromValue returns a pointer to a valid SunburstWidgetDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSunburstWidgetDefinitionTypeFromValue(v string) (*SunburstWidgetDefinitionType, error) {
	ev := SunburstWidgetDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SunburstWidgetDefinitionType: valid values are %v", v, allowedSunburstWidgetDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SunburstWidgetDefinitionType) IsValid() bool {
	for _, existing := range allowedSunburstWidgetDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SunburstWidgetDefinitionType value.
func (v SunburstWidgetDefinitionType) Ptr() *SunburstWidgetDefinitionType {
	return &v
}
