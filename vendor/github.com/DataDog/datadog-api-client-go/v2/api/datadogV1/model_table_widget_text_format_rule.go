// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableWidgetTextFormatRule Text format rules.
type TableWidgetTextFormatRule struct {
	// Hex representation of the custom background color. Used with custom background palette option.
	CustomBgColor *string `json:"custom_bg_color,omitempty"`
	// Hex representation of the custom text color. Used with custom text palette option.
	CustomFgColor *string `json:"custom_fg_color,omitempty"`
	// Match rule for the table widget text format.
	Match TableWidgetTextFormatMatch `json:"match"`
	// Color-on-color palette to highlight replaced text.
	Palette *TableWidgetTextFormatPalette `json:"palette,omitempty"`
	// Replace rule for the table widget text format.
	Replace *TableWidgetTextFormatReplace `json:"replace,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTableWidgetTextFormatRule instantiates a new TableWidgetTextFormatRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTableWidgetTextFormatRule(match TableWidgetTextFormatMatch) *TableWidgetTextFormatRule {
	this := TableWidgetTextFormatRule{}
	this.Match = match
	var palette TableWidgetTextFormatPalette = TABLEWIDGETTEXTFORMATPALETTE_WHITE_ON_GREEN
	this.Palette = &palette
	return &this
}

// NewTableWidgetTextFormatRuleWithDefaults instantiates a new TableWidgetTextFormatRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTableWidgetTextFormatRuleWithDefaults() *TableWidgetTextFormatRule {
	this := TableWidgetTextFormatRule{}
	var palette TableWidgetTextFormatPalette = TABLEWIDGETTEXTFORMATPALETTE_WHITE_ON_GREEN
	this.Palette = &palette
	return &this
}

// GetCustomBgColor returns the CustomBgColor field value if set, zero value otherwise.
func (o *TableWidgetTextFormatRule) GetCustomBgColor() string {
	if o == nil || o.CustomBgColor == nil {
		var ret string
		return ret
	}
	return *o.CustomBgColor
}

// GetCustomBgColorOk returns a tuple with the CustomBgColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableWidgetTextFormatRule) GetCustomBgColorOk() (*string, bool) {
	if o == nil || o.CustomBgColor == nil {
		return nil, false
	}
	return o.CustomBgColor, true
}

// HasCustomBgColor returns a boolean if a field has been set.
func (o *TableWidgetTextFormatRule) HasCustomBgColor() bool {
	return o != nil && o.CustomBgColor != nil
}

// SetCustomBgColor gets a reference to the given string and assigns it to the CustomBgColor field.
func (o *TableWidgetTextFormatRule) SetCustomBgColor(v string) {
	o.CustomBgColor = &v
}

// GetCustomFgColor returns the CustomFgColor field value if set, zero value otherwise.
func (o *TableWidgetTextFormatRule) GetCustomFgColor() string {
	if o == nil || o.CustomFgColor == nil {
		var ret string
		return ret
	}
	return *o.CustomFgColor
}

// GetCustomFgColorOk returns a tuple with the CustomFgColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableWidgetTextFormatRule) GetCustomFgColorOk() (*string, bool) {
	if o == nil || o.CustomFgColor == nil {
		return nil, false
	}
	return o.CustomFgColor, true
}

// HasCustomFgColor returns a boolean if a field has been set.
func (o *TableWidgetTextFormatRule) HasCustomFgColor() bool {
	return o != nil && o.CustomFgColor != nil
}

// SetCustomFgColor gets a reference to the given string and assigns it to the CustomFgColor field.
func (o *TableWidgetTextFormatRule) SetCustomFgColor(v string) {
	o.CustomFgColor = &v
}

// GetMatch returns the Match field value.
func (o *TableWidgetTextFormatRule) GetMatch() TableWidgetTextFormatMatch {
	if o == nil {
		var ret TableWidgetTextFormatMatch
		return ret
	}
	return o.Match
}

// GetMatchOk returns a tuple with the Match field value
// and a boolean to check if the value has been set.
func (o *TableWidgetTextFormatRule) GetMatchOk() (*TableWidgetTextFormatMatch, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Match, true
}

// SetMatch sets field value.
func (o *TableWidgetTextFormatRule) SetMatch(v TableWidgetTextFormatMatch) {
	o.Match = v
}

// GetPalette returns the Palette field value if set, zero value otherwise.
func (o *TableWidgetTextFormatRule) GetPalette() TableWidgetTextFormatPalette {
	if o == nil || o.Palette == nil {
		var ret TableWidgetTextFormatPalette
		return ret
	}
	return *o.Palette
}

// GetPaletteOk returns a tuple with the Palette field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableWidgetTextFormatRule) GetPaletteOk() (*TableWidgetTextFormatPalette, bool) {
	if o == nil || o.Palette == nil {
		return nil, false
	}
	return o.Palette, true
}

// HasPalette returns a boolean if a field has been set.
func (o *TableWidgetTextFormatRule) HasPalette() bool {
	return o != nil && o.Palette != nil
}

// SetPalette gets a reference to the given TableWidgetTextFormatPalette and assigns it to the Palette field.
func (o *TableWidgetTextFormatRule) SetPalette(v TableWidgetTextFormatPalette) {
	o.Palette = &v
}

// GetReplace returns the Replace field value if set, zero value otherwise.
func (o *TableWidgetTextFormatRule) GetReplace() TableWidgetTextFormatReplace {
	if o == nil || o.Replace == nil {
		var ret TableWidgetTextFormatReplace
		return ret
	}
	return *o.Replace
}

// GetReplaceOk returns a tuple with the Replace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableWidgetTextFormatRule) GetReplaceOk() (*TableWidgetTextFormatReplace, bool) {
	if o == nil || o.Replace == nil {
		return nil, false
	}
	return o.Replace, true
}

// HasReplace returns a boolean if a field has been set.
func (o *TableWidgetTextFormatRule) HasReplace() bool {
	return o != nil && o.Replace != nil
}

// SetReplace gets a reference to the given TableWidgetTextFormatReplace and assigns it to the Replace field.
func (o *TableWidgetTextFormatRule) SetReplace(v TableWidgetTextFormatReplace) {
	o.Replace = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TableWidgetTextFormatRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CustomBgColor != nil {
		toSerialize["custom_bg_color"] = o.CustomBgColor
	}
	if o.CustomFgColor != nil {
		toSerialize["custom_fg_color"] = o.CustomFgColor
	}
	toSerialize["match"] = o.Match
	if o.Palette != nil {
		toSerialize["palette"] = o.Palette
	}
	if o.Replace != nil {
		toSerialize["replace"] = o.Replace
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TableWidgetTextFormatRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CustomBgColor *string                       `json:"custom_bg_color,omitempty"`
		CustomFgColor *string                       `json:"custom_fg_color,omitempty"`
		Match         *TableWidgetTextFormatMatch   `json:"match"`
		Palette       *TableWidgetTextFormatPalette `json:"palette,omitempty"`
		Replace       *TableWidgetTextFormatReplace `json:"replace,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Match == nil {
		return fmt.Errorf("required field match missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"custom_bg_color", "custom_fg_color", "match", "palette", "replace"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CustomBgColor = all.CustomBgColor
	o.CustomFgColor = all.CustomFgColor
	if all.Match.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Match = *all.Match
	if all.Palette != nil && !all.Palette.IsValid() {
		hasInvalidField = true
	} else {
		o.Palette = all.Palette
	}
	o.Replace = all.Replace

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
