// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CheckStatusWidgetDefinitionType Type of the check status widget.
type CheckStatusWidgetDefinitionType string

// List of CheckStatusWidgetDefinitionType.
const (
	CHECKSTATUSWIDGETDEFINITIONTYPE_CHECK_STATUS CheckStatusWidgetDefinitionType = "check_status"
)

var allowedCheckStatusWidgetDefinitionTypeEnumValues = []CheckStatusWidgetDefinitionType{
	CHECKSTATUSWIDGETDEFINITIONTYPE_CHECK_STATUS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CheckStatusWidgetDefinitionType) GetAllowedValues() []CheckStatusWidgetDefinitionType {
	return allowedCheckStatusWidgetDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CheckStatusWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CheckStatusWidgetDefinitionType(value)
	return nil
}

// NewCheckStatusWidgetDefinitionTypeFromValue returns a pointer to a valid CheckStatusWidgetDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCheckStatusWidgetDefinitionTypeFromValue(v string) (*CheckStatusWidgetDefinitionType, error) {
	ev := CheckStatusWidgetDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CheckStatusWidgetDefinitionType: valid values are %v", v, allowedCheckStatusWidgetDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CheckStatusWidgetDefinitionType) IsValid() bool {
	for _, existing := range allowedCheckStatusWidgetDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CheckStatusWidgetDefinitionType value.
func (v CheckStatusWidgetDefinitionType) Ptr() *CheckStatusWidgetDefinitionType {
	return &v
}
