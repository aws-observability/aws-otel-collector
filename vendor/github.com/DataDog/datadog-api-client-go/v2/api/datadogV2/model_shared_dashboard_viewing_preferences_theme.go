// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SharedDashboardViewingPreferencesTheme The theme of the shared dashboard view. `system` follows the viewer's system default.
type SharedDashboardViewingPreferencesTheme string

// List of SharedDashboardViewingPreferencesTheme.
const (
	SHAREDDASHBOARDVIEWINGPREFERENCESTHEME_SYSTEM SharedDashboardViewingPreferencesTheme = "system"
	SHAREDDASHBOARDVIEWINGPREFERENCESTHEME_LIGHT  SharedDashboardViewingPreferencesTheme = "light"
	SHAREDDASHBOARDVIEWINGPREFERENCESTHEME_DARK   SharedDashboardViewingPreferencesTheme = "dark"
)

var allowedSharedDashboardViewingPreferencesThemeEnumValues = []SharedDashboardViewingPreferencesTheme{
	SHAREDDASHBOARDVIEWINGPREFERENCESTHEME_SYSTEM,
	SHAREDDASHBOARDVIEWINGPREFERENCESTHEME_LIGHT,
	SHAREDDASHBOARDVIEWINGPREFERENCESTHEME_DARK,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SharedDashboardViewingPreferencesTheme) GetAllowedValues() []SharedDashboardViewingPreferencesTheme {
	return allowedSharedDashboardViewingPreferencesThemeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SharedDashboardViewingPreferencesTheme) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SharedDashboardViewingPreferencesTheme(value)
	return nil
}

// NewSharedDashboardViewingPreferencesThemeFromValue returns a pointer to a valid SharedDashboardViewingPreferencesTheme
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSharedDashboardViewingPreferencesThemeFromValue(v string) (*SharedDashboardViewingPreferencesTheme, error) {
	ev := SharedDashboardViewingPreferencesTheme(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SharedDashboardViewingPreferencesTheme: valid values are %v", v, allowedSharedDashboardViewingPreferencesThemeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SharedDashboardViewingPreferencesTheme) IsValid() bool {
	for _, existing := range allowedSharedDashboardViewingPreferencesThemeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SharedDashboardViewingPreferencesTheme value.
func (v SharedDashboardViewingPreferencesTheme) Ptr() *SharedDashboardViewingPreferencesTheme {
	return &v
}
