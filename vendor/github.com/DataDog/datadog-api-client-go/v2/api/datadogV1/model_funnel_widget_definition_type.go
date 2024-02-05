// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FunnelWidgetDefinitionType Type of funnel widget.
type FunnelWidgetDefinitionType string

// List of FunnelWidgetDefinitionType.
const (
	FUNNELWIDGETDEFINITIONTYPE_FUNNEL FunnelWidgetDefinitionType = "funnel"
)

var allowedFunnelWidgetDefinitionTypeEnumValues = []FunnelWidgetDefinitionType{
	FUNNELWIDGETDEFINITIONTYPE_FUNNEL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FunnelWidgetDefinitionType) GetAllowedValues() []FunnelWidgetDefinitionType {
	return allowedFunnelWidgetDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FunnelWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FunnelWidgetDefinitionType(value)
	return nil
}

// NewFunnelWidgetDefinitionTypeFromValue returns a pointer to a valid FunnelWidgetDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFunnelWidgetDefinitionTypeFromValue(v string) (*FunnelWidgetDefinitionType, error) {
	ev := FunnelWidgetDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FunnelWidgetDefinitionType: valid values are %v", v, allowedFunnelWidgetDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FunnelWidgetDefinitionType) IsValid() bool {
	for _, existing := range allowedFunnelWidgetDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FunnelWidgetDefinitionType value.
func (v FunnelWidgetDefinitionType) Ptr() *FunnelWidgetDefinitionType {
	return &v
}
