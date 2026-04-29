// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WildcardWidgetDefinitionType Type of the wildcard widget.
type WildcardWidgetDefinitionType string

// List of WildcardWidgetDefinitionType.
const (
	WILDCARDWIDGETDEFINITIONTYPE_WILDCARD WildcardWidgetDefinitionType = "wildcard"
)

var allowedWildcardWidgetDefinitionTypeEnumValues = []WildcardWidgetDefinitionType{
	WILDCARDWIDGETDEFINITIONTYPE_WILDCARD,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WildcardWidgetDefinitionType) GetAllowedValues() []WildcardWidgetDefinitionType {
	return allowedWildcardWidgetDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WildcardWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WildcardWidgetDefinitionType(value)
	return nil
}

// NewWildcardWidgetDefinitionTypeFromValue returns a pointer to a valid WildcardWidgetDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWildcardWidgetDefinitionTypeFromValue(v string) (*WildcardWidgetDefinitionType, error) {
	ev := WildcardWidgetDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WildcardWidgetDefinitionType: valid values are %v", v, allowedWildcardWidgetDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WildcardWidgetDefinitionType) IsValid() bool {
	for _, existing := range allowedWildcardWidgetDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WildcardWidgetDefinitionType value.
func (v WildcardWidgetDefinitionType) Ptr() *WildcardWidgetDefinitionType {
	return &v
}
