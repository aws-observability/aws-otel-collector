// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SankeyWidgetDefinitionType Type of the Sankey widget.
type SankeyWidgetDefinitionType string

// List of SankeyWidgetDefinitionType.
const (
	SANKEYWIDGETDEFINITIONTYPE_SANKEY SankeyWidgetDefinitionType = "sankey"
)

var allowedSankeyWidgetDefinitionTypeEnumValues = []SankeyWidgetDefinitionType{
	SANKEYWIDGETDEFINITIONTYPE_SANKEY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SankeyWidgetDefinitionType) GetAllowedValues() []SankeyWidgetDefinitionType {
	return allowedSankeyWidgetDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SankeyWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SankeyWidgetDefinitionType(value)
	return nil
}

// NewSankeyWidgetDefinitionTypeFromValue returns a pointer to a valid SankeyWidgetDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSankeyWidgetDefinitionTypeFromValue(v string) (*SankeyWidgetDefinitionType, error) {
	ev := SankeyWidgetDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SankeyWidgetDefinitionType: valid values are %v", v, allowedSankeyWidgetDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SankeyWidgetDefinitionType) IsValid() bool {
	for _, existing := range allowedSankeyWidgetDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SankeyWidgetDefinitionType value.
func (v SankeyWidgetDefinitionType) Ptr() *SankeyWidgetDefinitionType {
	return &v
}
