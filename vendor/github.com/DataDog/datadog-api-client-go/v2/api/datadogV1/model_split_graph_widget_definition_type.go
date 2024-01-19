// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SplitGraphWidgetDefinitionType Type of the split graph widget
type SplitGraphWidgetDefinitionType string

// List of SplitGraphWidgetDefinitionType.
const (
	SPLITGRAPHWIDGETDEFINITIONTYPE_SPLIT_GROUP SplitGraphWidgetDefinitionType = "split_group"
)

var allowedSplitGraphWidgetDefinitionTypeEnumValues = []SplitGraphWidgetDefinitionType{
	SPLITGRAPHWIDGETDEFINITIONTYPE_SPLIT_GROUP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SplitGraphWidgetDefinitionType) GetAllowedValues() []SplitGraphWidgetDefinitionType {
	return allowedSplitGraphWidgetDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SplitGraphWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SplitGraphWidgetDefinitionType(value)
	return nil
}

// NewSplitGraphWidgetDefinitionTypeFromValue returns a pointer to a valid SplitGraphWidgetDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSplitGraphWidgetDefinitionTypeFromValue(v string) (*SplitGraphWidgetDefinitionType, error) {
	ev := SplitGraphWidgetDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SplitGraphWidgetDefinitionType: valid values are %v", v, allowedSplitGraphWidgetDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SplitGraphWidgetDefinitionType) IsValid() bool {
	for _, existing := range allowedSplitGraphWidgetDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SplitGraphWidgetDefinitionType value.
func (v SplitGraphWidgetDefinitionType) Ptr() *SplitGraphWidgetDefinitionType {
	return &v
}
