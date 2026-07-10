// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SharedDashboardViewingPreferences Display settings for the shared dashboard.
type SharedDashboardViewingPreferences struct {
	// Whether widgets are displayed in high-density mode.
	HighDensity bool `json:"high_density"`
	// The theme of the shared dashboard view. `system` follows the viewer's system default.
	Theme SharedDashboardViewingPreferencesTheme `json:"theme"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSharedDashboardViewingPreferences instantiates a new SharedDashboardViewingPreferences object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSharedDashboardViewingPreferences(highDensity bool, theme SharedDashboardViewingPreferencesTheme) *SharedDashboardViewingPreferences {
	this := SharedDashboardViewingPreferences{}
	this.HighDensity = highDensity
	this.Theme = theme
	return &this
}

// NewSharedDashboardViewingPreferencesWithDefaults instantiates a new SharedDashboardViewingPreferences object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSharedDashboardViewingPreferencesWithDefaults() *SharedDashboardViewingPreferences {
	this := SharedDashboardViewingPreferences{}
	return &this
}

// GetHighDensity returns the HighDensity field value.
func (o *SharedDashboardViewingPreferences) GetHighDensity() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.HighDensity
}

// GetHighDensityOk returns a tuple with the HighDensity field value
// and a boolean to check if the value has been set.
func (o *SharedDashboardViewingPreferences) GetHighDensityOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HighDensity, true
}

// SetHighDensity sets field value.
func (o *SharedDashboardViewingPreferences) SetHighDensity(v bool) {
	o.HighDensity = v
}

// GetTheme returns the Theme field value.
func (o *SharedDashboardViewingPreferences) GetTheme() SharedDashboardViewingPreferencesTheme {
	if o == nil {
		var ret SharedDashboardViewingPreferencesTheme
		return ret
	}
	return o.Theme
}

// GetThemeOk returns a tuple with the Theme field value
// and a boolean to check if the value has been set.
func (o *SharedDashboardViewingPreferences) GetThemeOk() (*SharedDashboardViewingPreferencesTheme, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Theme, true
}

// SetTheme sets field value.
func (o *SharedDashboardViewingPreferences) SetTheme(v SharedDashboardViewingPreferencesTheme) {
	o.Theme = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SharedDashboardViewingPreferences) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["high_density"] = o.HighDensity
	toSerialize["theme"] = o.Theme

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SharedDashboardViewingPreferences) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		HighDensity *bool                                   `json:"high_density"`
		Theme       *SharedDashboardViewingPreferencesTheme `json:"theme"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.HighDensity == nil {
		return fmt.Errorf("required field high_density missing")
	}
	if all.Theme == nil {
		return fmt.Errorf("required field theme missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"high_density", "theme"})
	} else {
		return err
	}

	hasInvalidField := false
	o.HighDensity = *all.HighDensity
	if !all.Theme.IsValid() {
		hasInvalidField = true
	} else {
		o.Theme = *all.Theme
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
