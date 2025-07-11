// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ViewingPreferencesTheme The theme of the shared dashboard view. "system" follows your system's default viewing theme.
type ViewingPreferencesTheme string

// List of ViewingPreferencesTheme.
const (
	VIEWINGPREFERENCESTHEME_SYSTEM ViewingPreferencesTheme = "system"
	VIEWINGPREFERENCESTHEME_LIGHT  ViewingPreferencesTheme = "light"
	VIEWINGPREFERENCESTHEME_DARK   ViewingPreferencesTheme = "dark"
)

var allowedViewingPreferencesThemeEnumValues = []ViewingPreferencesTheme{
	VIEWINGPREFERENCESTHEME_SYSTEM,
	VIEWINGPREFERENCESTHEME_LIGHT,
	VIEWINGPREFERENCESTHEME_DARK,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ViewingPreferencesTheme) GetAllowedValues() []ViewingPreferencesTheme {
	return allowedViewingPreferencesThemeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ViewingPreferencesTheme) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ViewingPreferencesTheme(value)
	return nil
}

// NewViewingPreferencesThemeFromValue returns a pointer to a valid ViewingPreferencesTheme
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewViewingPreferencesThemeFromValue(v string) (*ViewingPreferencesTheme, error) {
	ev := ViewingPreferencesTheme(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ViewingPreferencesTheme: valid values are %v", v, allowedViewingPreferencesThemeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ViewingPreferencesTheme) IsValid() bool {
	for _, existing := range allowedViewingPreferencesThemeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ViewingPreferencesTheme value.
func (v ViewingPreferencesTheme) Ptr() *ViewingPreferencesTheme {
	return &v
}
