// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormUiDefinitionUiThemePrimaryColor The primary color of the form theme.
type FormUiDefinitionUiThemePrimaryColor string

// List of FormUiDefinitionUiThemePrimaryColor.
const (
	FORMUIDEFINITIONUITHEMEPRIMARYCOLOR_GRAY       FormUiDefinitionUiThemePrimaryColor = "gray"
	FORMUIDEFINITIONUITHEMEPRIMARYCOLOR_RED        FormUiDefinitionUiThemePrimaryColor = "red"
	FORMUIDEFINITIONUITHEMEPRIMARYCOLOR_ORANGE     FormUiDefinitionUiThemePrimaryColor = "orange"
	FORMUIDEFINITIONUITHEMEPRIMARYCOLOR_YELLOW     FormUiDefinitionUiThemePrimaryColor = "yellow"
	FORMUIDEFINITIONUITHEMEPRIMARYCOLOR_GREEN      FormUiDefinitionUiThemePrimaryColor = "green"
	FORMUIDEFINITIONUITHEMEPRIMARYCOLOR_LIGHT_BLUE FormUiDefinitionUiThemePrimaryColor = "light-blue"
	FORMUIDEFINITIONUITHEMEPRIMARYCOLOR_DARK_BLUE  FormUiDefinitionUiThemePrimaryColor = "dark-blue"
	FORMUIDEFINITIONUITHEMEPRIMARYCOLOR_MAGENTA    FormUiDefinitionUiThemePrimaryColor = "magenta"
	FORMUIDEFINITIONUITHEMEPRIMARYCOLOR_INDIGO     FormUiDefinitionUiThemePrimaryColor = "indigo"
)

var allowedFormUiDefinitionUiThemePrimaryColorEnumValues = []FormUiDefinitionUiThemePrimaryColor{
	FORMUIDEFINITIONUITHEMEPRIMARYCOLOR_GRAY,
	FORMUIDEFINITIONUITHEMEPRIMARYCOLOR_RED,
	FORMUIDEFINITIONUITHEMEPRIMARYCOLOR_ORANGE,
	FORMUIDEFINITIONUITHEMEPRIMARYCOLOR_YELLOW,
	FORMUIDEFINITIONUITHEMEPRIMARYCOLOR_GREEN,
	FORMUIDEFINITIONUITHEMEPRIMARYCOLOR_LIGHT_BLUE,
	FORMUIDEFINITIONUITHEMEPRIMARYCOLOR_DARK_BLUE,
	FORMUIDEFINITIONUITHEMEPRIMARYCOLOR_MAGENTA,
	FORMUIDEFINITIONUITHEMEPRIMARYCOLOR_INDIGO,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormUiDefinitionUiThemePrimaryColor) GetAllowedValues() []FormUiDefinitionUiThemePrimaryColor {
	return allowedFormUiDefinitionUiThemePrimaryColorEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormUiDefinitionUiThemePrimaryColor) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormUiDefinitionUiThemePrimaryColor(value)
	return nil
}

// NewFormUiDefinitionUiThemePrimaryColorFromValue returns a pointer to a valid FormUiDefinitionUiThemePrimaryColor
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormUiDefinitionUiThemePrimaryColorFromValue(v string) (*FormUiDefinitionUiThemePrimaryColor, error) {
	ev := FormUiDefinitionUiThemePrimaryColor(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormUiDefinitionUiThemePrimaryColor: valid values are %v", v, allowedFormUiDefinitionUiThemePrimaryColorEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormUiDefinitionUiThemePrimaryColor) IsValid() bool {
	for _, existing := range allowedFormUiDefinitionUiThemePrimaryColorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormUiDefinitionUiThemePrimaryColor value.
func (v FormUiDefinitionUiThemePrimaryColor) Ptr() *FormUiDefinitionUiThemePrimaryColor {
	return &v
}
