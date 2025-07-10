// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ViewingPreferences The viewing preferences for a shared dashboard.
type ViewingPreferences struct {
	// Whether the widgets on the shared dashboard should be displayed with high density.
	HighDensity *bool `json:"high_density,omitempty"`
	// The theme of the shared dashboard view. "system" follows your system's default viewing theme.
	Theme *ViewingPreferencesTheme `json:"theme,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewViewingPreferences instantiates a new ViewingPreferences object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewViewingPreferences() *ViewingPreferences {
	this := ViewingPreferences{}
	return &this
}

// NewViewingPreferencesWithDefaults instantiates a new ViewingPreferences object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewViewingPreferencesWithDefaults() *ViewingPreferences {
	this := ViewingPreferences{}
	return &this
}

// GetHighDensity returns the HighDensity field value if set, zero value otherwise.
func (o *ViewingPreferences) GetHighDensity() bool {
	if o == nil || o.HighDensity == nil {
		var ret bool
		return ret
	}
	return *o.HighDensity
}

// GetHighDensityOk returns a tuple with the HighDensity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewingPreferences) GetHighDensityOk() (*bool, bool) {
	if o == nil || o.HighDensity == nil {
		return nil, false
	}
	return o.HighDensity, true
}

// HasHighDensity returns a boolean if a field has been set.
func (o *ViewingPreferences) HasHighDensity() bool {
	return o != nil && o.HighDensity != nil
}

// SetHighDensity gets a reference to the given bool and assigns it to the HighDensity field.
func (o *ViewingPreferences) SetHighDensity(v bool) {
	o.HighDensity = &v
}

// GetTheme returns the Theme field value if set, zero value otherwise.
func (o *ViewingPreferences) GetTheme() ViewingPreferencesTheme {
	if o == nil || o.Theme == nil {
		var ret ViewingPreferencesTheme
		return ret
	}
	return *o.Theme
}

// GetThemeOk returns a tuple with the Theme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewingPreferences) GetThemeOk() (*ViewingPreferencesTheme, bool) {
	if o == nil || o.Theme == nil {
		return nil, false
	}
	return o.Theme, true
}

// HasTheme returns a boolean if a field has been set.
func (o *ViewingPreferences) HasTheme() bool {
	return o != nil && o.Theme != nil
}

// SetTheme gets a reference to the given ViewingPreferencesTheme and assigns it to the Theme field.
func (o *ViewingPreferences) SetTheme(v ViewingPreferencesTheme) {
	o.Theme = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ViewingPreferences) MarshalJSON() ([]byte, error) {
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
func (o *ViewingPreferences) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		HighDensity *bool                    `json:"high_density,omitempty"`
		Theme       *ViewingPreferencesTheme `json:"theme,omitempty"`
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
