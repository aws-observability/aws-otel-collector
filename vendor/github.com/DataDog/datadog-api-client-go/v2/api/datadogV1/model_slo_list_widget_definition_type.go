// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SLOListWidgetDefinitionType Type of the SLO List widget.
type SLOListWidgetDefinitionType string

// List of SLOListWidgetDefinitionType.
const (
	SLOLISTWIDGETDEFINITIONTYPE_SLO_LIST SLOListWidgetDefinitionType = "slo_list"
)

var allowedSLOListWidgetDefinitionTypeEnumValues = []SLOListWidgetDefinitionType{
	SLOLISTWIDGETDEFINITIONTYPE_SLO_LIST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SLOListWidgetDefinitionType) GetAllowedValues() []SLOListWidgetDefinitionType {
	return allowedSLOListWidgetDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SLOListWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SLOListWidgetDefinitionType(value)
	return nil
}

// NewSLOListWidgetDefinitionTypeFromValue returns a pointer to a valid SLOListWidgetDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSLOListWidgetDefinitionTypeFromValue(v string) (*SLOListWidgetDefinitionType, error) {
	ev := SLOListWidgetDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SLOListWidgetDefinitionType: valid values are %v", v, allowedSLOListWidgetDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SLOListWidgetDefinitionType) IsValid() bool {
	for _, existing := range allowedSLOListWidgetDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SLOListWidgetDefinitionType value.
func (v SLOListWidgetDefinitionType) Ptr() *SLOListWidgetDefinitionType {
	return &v
}
