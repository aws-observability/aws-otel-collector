// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetFormulaCellDisplayModeOptions Cell display mode options for the widget formula. (only if `cell_display_mode` is set to `trend`).
type WidgetFormulaCellDisplayModeOptions struct {
	// Trend type for the cell display mode options.
	TrendType *WidgetFormulaCellDisplayModeOptionsTrendType `json:"trend_type,omitempty"`
	// Y scale for the cell display mode options.
	YScale *WidgetFormulaCellDisplayModeOptionsYScale `json:"y_scale,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWidgetFormulaCellDisplayModeOptions instantiates a new WidgetFormulaCellDisplayModeOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWidgetFormulaCellDisplayModeOptions() *WidgetFormulaCellDisplayModeOptions {
	this := WidgetFormulaCellDisplayModeOptions{}
	return &this
}

// NewWidgetFormulaCellDisplayModeOptionsWithDefaults instantiates a new WidgetFormulaCellDisplayModeOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWidgetFormulaCellDisplayModeOptionsWithDefaults() *WidgetFormulaCellDisplayModeOptions {
	this := WidgetFormulaCellDisplayModeOptions{}
	return &this
}

// GetTrendType returns the TrendType field value if set, zero value otherwise.
func (o *WidgetFormulaCellDisplayModeOptions) GetTrendType() WidgetFormulaCellDisplayModeOptionsTrendType {
	if o == nil || o.TrendType == nil {
		var ret WidgetFormulaCellDisplayModeOptionsTrendType
		return ret
	}
	return *o.TrendType
}

// GetTrendTypeOk returns a tuple with the TrendType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetFormulaCellDisplayModeOptions) GetTrendTypeOk() (*WidgetFormulaCellDisplayModeOptionsTrendType, bool) {
	if o == nil || o.TrendType == nil {
		return nil, false
	}
	return o.TrendType, true
}

// HasTrendType returns a boolean if a field has been set.
func (o *WidgetFormulaCellDisplayModeOptions) HasTrendType() bool {
	return o != nil && o.TrendType != nil
}

// SetTrendType gets a reference to the given WidgetFormulaCellDisplayModeOptionsTrendType and assigns it to the TrendType field.
func (o *WidgetFormulaCellDisplayModeOptions) SetTrendType(v WidgetFormulaCellDisplayModeOptionsTrendType) {
	o.TrendType = &v
}

// GetYScale returns the YScale field value if set, zero value otherwise.
func (o *WidgetFormulaCellDisplayModeOptions) GetYScale() WidgetFormulaCellDisplayModeOptionsYScale {
	if o == nil || o.YScale == nil {
		var ret WidgetFormulaCellDisplayModeOptionsYScale
		return ret
	}
	return *o.YScale
}

// GetYScaleOk returns a tuple with the YScale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetFormulaCellDisplayModeOptions) GetYScaleOk() (*WidgetFormulaCellDisplayModeOptionsYScale, bool) {
	if o == nil || o.YScale == nil {
		return nil, false
	}
	return o.YScale, true
}

// HasYScale returns a boolean if a field has been set.
func (o *WidgetFormulaCellDisplayModeOptions) HasYScale() bool {
	return o != nil && o.YScale != nil
}

// SetYScale gets a reference to the given WidgetFormulaCellDisplayModeOptionsYScale and assigns it to the YScale field.
func (o *WidgetFormulaCellDisplayModeOptions) SetYScale(v WidgetFormulaCellDisplayModeOptionsYScale) {
	o.YScale = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o WidgetFormulaCellDisplayModeOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.TrendType != nil {
		toSerialize["trend_type"] = o.TrendType
	}
	if o.YScale != nil {
		toSerialize["y_scale"] = o.YScale
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WidgetFormulaCellDisplayModeOptions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TrendType *WidgetFormulaCellDisplayModeOptionsTrendType `json:"trend_type,omitempty"`
		YScale    *WidgetFormulaCellDisplayModeOptionsYScale    `json:"y_scale,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"trend_type", "y_scale"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.TrendType != nil && !all.TrendType.IsValid() {
		hasInvalidField = true
	} else {
		o.TrendType = all.TrendType
	}
	if all.YScale != nil && !all.YScale.IsValid() {
		hasInvalidField = true
	} else {
		o.YScale = all.YScale
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
