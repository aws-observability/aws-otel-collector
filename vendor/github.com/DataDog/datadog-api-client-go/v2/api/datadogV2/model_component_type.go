// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ComponentType The UI component type.
type ComponentType string

// List of ComponentType.
const (
	COMPONENTTYPE_TABLE           ComponentType = "table"
	COMPONENTTYPE_TEXTINPUT       ComponentType = "textInput"
	COMPONENTTYPE_TEXTAREA        ComponentType = "textArea"
	COMPONENTTYPE_BUTTON          ComponentType = "button"
	COMPONENTTYPE_TEXT            ComponentType = "text"
	COMPONENTTYPE_SELECT          ComponentType = "select"
	COMPONENTTYPE_MODAL           ComponentType = "modal"
	COMPONENTTYPE_SCHEMAFORM      ComponentType = "schemaForm"
	COMPONENTTYPE_CHECKBOX        ComponentType = "checkbox"
	COMPONENTTYPE_TABS            ComponentType = "tabs"
	COMPONENTTYPE_VEGACHART       ComponentType = "vegaChart"
	COMPONENTTYPE_RADIOBUTTONS    ComponentType = "radioButtons"
	COMPONENTTYPE_NUMBERINPUT     ComponentType = "numberInput"
	COMPONENTTYPE_FILEINPUT       ComponentType = "fileInput"
	COMPONENTTYPE_JSONINPUT       ComponentType = "jsonInput"
	COMPONENTTYPE_GRIDCELL        ComponentType = "gridCell"
	COMPONENTTYPE_DATERANGEPICKER ComponentType = "dateRangePicker"
	COMPONENTTYPE_SEARCH          ComponentType = "search"
	COMPONENTTYPE_CONTAINER       ComponentType = "container"
	COMPONENTTYPE_CALLOUTVALUE    ComponentType = "calloutValue"
)

var allowedComponentTypeEnumValues = []ComponentType{
	COMPONENTTYPE_TABLE,
	COMPONENTTYPE_TEXTINPUT,
	COMPONENTTYPE_TEXTAREA,
	COMPONENTTYPE_BUTTON,
	COMPONENTTYPE_TEXT,
	COMPONENTTYPE_SELECT,
	COMPONENTTYPE_MODAL,
	COMPONENTTYPE_SCHEMAFORM,
	COMPONENTTYPE_CHECKBOX,
	COMPONENTTYPE_TABS,
	COMPONENTTYPE_VEGACHART,
	COMPONENTTYPE_RADIOBUTTONS,
	COMPONENTTYPE_NUMBERINPUT,
	COMPONENTTYPE_FILEINPUT,
	COMPONENTTYPE_JSONINPUT,
	COMPONENTTYPE_GRIDCELL,
	COMPONENTTYPE_DATERANGEPICKER,
	COMPONENTTYPE_SEARCH,
	COMPONENTTYPE_CONTAINER,
	COMPONENTTYPE_CALLOUTVALUE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ComponentType) GetAllowedValues() []ComponentType {
	return allowedComponentTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ComponentType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ComponentType(value)
	return nil
}

// NewComponentTypeFromValue returns a pointer to a valid ComponentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewComponentTypeFromValue(v string) (*ComponentType, error) {
	ev := ComponentType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ComponentType: valid values are %v", v, allowedComponentTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ComponentType) IsValid() bool {
	for _, existing := range allowedComponentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ComponentType value.
func (v ComponentType) Ptr() *ComponentType {
	return &v
}
