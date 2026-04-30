// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecureEmbedViewingPreferences Display settings for the secure embed shared dashboard.
type SecureEmbedViewingPreferences struct {
	// Whether widgets are displayed in high density mode.
	HighDensity *bool `json:"high_density,omitempty"`
	// The theme of the shared dashboard view. `system` follows the viewer's system default.
	Theme *SecureEmbedViewingPreferencesTheme `json:"theme,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecureEmbedViewingPreferences instantiates a new SecureEmbedViewingPreferences object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecureEmbedViewingPreferences() *SecureEmbedViewingPreferences {
	this := SecureEmbedViewingPreferences{}
	return &this
}

// NewSecureEmbedViewingPreferencesWithDefaults instantiates a new SecureEmbedViewingPreferences object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecureEmbedViewingPreferencesWithDefaults() *SecureEmbedViewingPreferences {
	this := SecureEmbedViewingPreferences{}
	return &this
}

// GetHighDensity returns the HighDensity field value if set, zero value otherwise.
func (o *SecureEmbedViewingPreferences) GetHighDensity() bool {
	if o == nil || o.HighDensity == nil {
		var ret bool
		return ret
	}
	return *o.HighDensity
}

// GetHighDensityOk returns a tuple with the HighDensity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureEmbedViewingPreferences) GetHighDensityOk() (*bool, bool) {
	if o == nil || o.HighDensity == nil {
		return nil, false
	}
	return o.HighDensity, true
}

// HasHighDensity returns a boolean if a field has been set.
func (o *SecureEmbedViewingPreferences) HasHighDensity() bool {
	return o != nil && o.HighDensity != nil
}

// SetHighDensity gets a reference to the given bool and assigns it to the HighDensity field.
func (o *SecureEmbedViewingPreferences) SetHighDensity(v bool) {
	o.HighDensity = &v
}

// GetTheme returns the Theme field value if set, zero value otherwise.
func (o *SecureEmbedViewingPreferences) GetTheme() SecureEmbedViewingPreferencesTheme {
	if o == nil || o.Theme == nil {
		var ret SecureEmbedViewingPreferencesTheme
		return ret
	}
	return *o.Theme
}

// GetThemeOk returns a tuple with the Theme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureEmbedViewingPreferences) GetThemeOk() (*SecureEmbedViewingPreferencesTheme, bool) {
	if o == nil || o.Theme == nil {
		return nil, false
	}
	return o.Theme, true
}

// HasTheme returns a boolean if a field has been set.
func (o *SecureEmbedViewingPreferences) HasTheme() bool {
	return o != nil && o.Theme != nil
}

// SetTheme gets a reference to the given SecureEmbedViewingPreferencesTheme and assigns it to the Theme field.
func (o *SecureEmbedViewingPreferences) SetTheme(v SecureEmbedViewingPreferencesTheme) {
	o.Theme = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecureEmbedViewingPreferences) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.HighDensity != nil {
		toSerialize["high_density"] = o.HighDensity
	}
	if o.Theme != nil {
		toSerialize["theme"] = o.Theme
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecureEmbedViewingPreferences) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		HighDensity *bool                               `json:"high_density,omitempty"`
		Theme       *SecureEmbedViewingPreferencesTheme `json:"theme,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"high_density", "theme"})
	} else {
		return err
	}

	hasInvalidField := false
	o.HighDensity = all.HighDensity
	if all.Theme != nil && !all.Theme.IsValid() {
		hasInvalidField = true
	} else {
		o.Theme = all.Theme
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
