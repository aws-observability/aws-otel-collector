// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SLOWidgetDefinitionType Type of the SLO widget.
type SLOWidgetDefinitionType string

// List of SLOWidgetDefinitionType.
const (
	SLOWIDGETDEFINITIONTYPE_SLO SLOWidgetDefinitionType = "slo"
)

var allowedSLOWidgetDefinitionTypeEnumValues = []SLOWidgetDefinitionType{
	SLOWIDGETDEFINITIONTYPE_SLO,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SLOWidgetDefinitionType) GetAllowedValues() []SLOWidgetDefinitionType {
	return allowedSLOWidgetDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SLOWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SLOWidgetDefinitionType(value)
	return nil
}

// NewSLOWidgetDefinitionTypeFromValue returns a pointer to a valid SLOWidgetDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSLOWidgetDefinitionTypeFromValue(v string) (*SLOWidgetDefinitionType, error) {
	ev := SLOWidgetDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SLOWidgetDefinitionType: valid values are %v", v, allowedSLOWidgetDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SLOWidgetDefinitionType) IsValid() bool {
	for _, existing := range allowedSLOWidgetDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SLOWidgetDefinitionType value.
func (v SLOWidgetDefinitionType) Ptr() *SLOWidgetDefinitionType {
	return &v
}
