// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BarChartWidgetStyle Style customization for a bar chart widget.
type BarChartWidgetStyle struct {
	// Bar chart widget display options.
	Display *BarChartWidgetDisplay `json:"display,omitempty"`
	// Color palette to apply to the widget.
	Palette *string `json:"palette,omitempty"`
	// Bar chart widget scaling definition.
	Scaling *BarChartWidgetScaling `json:"scaling,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBarChartWidgetStyle instantiates a new BarChartWidgetStyle object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBarChartWidgetStyle() *BarChartWidgetStyle {
	this := BarChartWidgetStyle{}
	return &this
}

// NewBarChartWidgetStyleWithDefaults instantiates a new BarChartWidgetStyle object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBarChartWidgetStyleWithDefaults() *BarChartWidgetStyle {
	this := BarChartWidgetStyle{}
	return &this
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *BarChartWidgetStyle) GetDisplay() BarChartWidgetDisplay {
	if o == nil || o.Display == nil {
		var ret BarChartWidgetDisplay
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BarChartWidgetStyle) GetDisplayOk() (*BarChartWidgetDisplay, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *BarChartWidgetStyle) HasDisplay() bool {
	return o != nil && o.Display != nil
}

// SetDisplay gets a reference to the given BarChartWidgetDisplay and assigns it to the Display field.
func (o *BarChartWidgetStyle) SetDisplay(v BarChartWidgetDisplay) {
	o.Display = &v
}

// GetPalette returns the Palette field value if set, zero value otherwise.
func (o *BarChartWidgetStyle) GetPalette() string {
	if o == nil || o.Palette == nil {
		var ret string
		return ret
	}
	return *o.Palette
}

// GetPaletteOk returns a tuple with the Palette field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BarChartWidgetStyle) GetPaletteOk() (*string, bool) {
	if o == nil || o.Palette == nil {
		return nil, false
	}
	return o.Palette, true
}

// HasPalette returns a boolean if a field has been set.
func (o *BarChartWidgetStyle) HasPalette() bool {
	return o != nil && o.Palette != nil
}

// SetPalette gets a reference to the given string and assigns it to the Palette field.
func (o *BarChartWidgetStyle) SetPalette(v string) {
	o.Palette = &v
}

// GetScaling returns the Scaling field value if set, zero value otherwise.
func (o *BarChartWidgetStyle) GetScaling() BarChartWidgetScaling {
	if o == nil || o.Scaling == nil {
		var ret BarChartWidgetScaling
		return ret
	}
	return *o.Scaling
}

// GetScalingOk returns a tuple with the Scaling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BarChartWidgetStyle) GetScalingOk() (*BarChartWidgetScaling, bool) {
	if o == nil || o.Scaling == nil {
		return nil, false
	}
	return o.Scaling, true
}

// HasScaling returns a boolean if a field has been set.
func (o *BarChartWidgetStyle) HasScaling() bool {
	return o != nil && o.Scaling != nil
}

// SetScaling gets a reference to the given BarChartWidgetScaling and assigns it to the Scaling field.
func (o *BarChartWidgetStyle) SetScaling(v BarChartWidgetScaling) {
	o.Scaling = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o BarChartWidgetStyle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Display != nil {
		toSerialize["display"] = o.Display
	}
	if o.Palette != nil {
		toSerialize["palette"] = o.Palette
	}
	if o.Scaling != nil {
		toSerialize["scaling"] = o.Scaling
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BarChartWidgetStyle) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Display *BarChartWidgetDisplay `json:"display,omitempty"`
		Palette *string                `json:"palette,omitempty"`
		Scaling *BarChartWidgetScaling `json:"scaling,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"display", "palette", "scaling"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Display = all.Display
	o.Palette = all.Palette
	if all.Scaling != nil && !all.Scaling.IsValid() {
		hasInvalidField = true
	} else {
		o.Scaling = all.Scaling
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
