// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecureEmbedViewingPreferencesTheme The theme of the shared dashboard view. `system` follows the viewer's system default.
type SecureEmbedViewingPreferencesTheme string

// List of SecureEmbedViewingPreferencesTheme.
const (
	SECUREEMBEDVIEWINGPREFERENCESTHEME_SYSTEM SecureEmbedViewingPreferencesTheme = "system"
	SECUREEMBEDVIEWINGPREFERENCESTHEME_LIGHT  SecureEmbedViewingPreferencesTheme = "light"
	SECUREEMBEDVIEWINGPREFERENCESTHEME_DARK   SecureEmbedViewingPreferencesTheme = "dark"
)

var allowedSecureEmbedViewingPreferencesThemeEnumValues = []SecureEmbedViewingPreferencesTheme{
	SECUREEMBEDVIEWINGPREFERENCESTHEME_SYSTEM,
	SECUREEMBEDVIEWINGPREFERENCESTHEME_LIGHT,
	SECUREEMBEDVIEWINGPREFERENCESTHEME_DARK,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecureEmbedViewingPreferencesTheme) GetAllowedValues() []SecureEmbedViewingPreferencesTheme {
	return allowedSecureEmbedViewingPreferencesThemeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecureEmbedViewingPreferencesTheme) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecureEmbedViewingPreferencesTheme(value)
	return nil
}

// NewSecureEmbedViewingPreferencesThemeFromValue returns a pointer to a valid SecureEmbedViewingPreferencesTheme
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecureEmbedViewingPreferencesThemeFromValue(v string) (*SecureEmbedViewingPreferencesTheme, error) {
	ev := SecureEmbedViewingPreferencesTheme(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecureEmbedViewingPreferencesTheme: valid values are %v", v, allowedSecureEmbedViewingPreferencesThemeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecureEmbedViewingPreferencesTheme) IsValid() bool {
	for _, existing := range allowedSecureEmbedViewingPreferencesThemeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecureEmbedViewingPreferencesTheme value.
func (v SecureEmbedViewingPreferencesTheme) Ptr() *SecureEmbedViewingPreferencesTheme {
	return &v
}
