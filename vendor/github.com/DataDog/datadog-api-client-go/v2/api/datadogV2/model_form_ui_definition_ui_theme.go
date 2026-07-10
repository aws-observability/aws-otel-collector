// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormUiDefinitionUiTheme The visual theme applied to the form.
type FormUiDefinitionUiTheme struct {
	// The primary color of the form theme.
	PrimaryColor *FormUiDefinitionUiThemePrimaryColor `json:"primaryColor,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFormUiDefinitionUiTheme instantiates a new FormUiDefinitionUiTheme object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormUiDefinitionUiTheme() *FormUiDefinitionUiTheme {
	this := FormUiDefinitionUiTheme{}
	return &this
}

// NewFormUiDefinitionUiThemeWithDefaults instantiates a new FormUiDefinitionUiTheme object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormUiDefinitionUiThemeWithDefaults() *FormUiDefinitionUiTheme {
	this := FormUiDefinitionUiTheme{}
	return &this
}

// GetPrimaryColor returns the PrimaryColor field value if set, zero value otherwise.
func (o *FormUiDefinitionUiTheme) GetPrimaryColor() FormUiDefinitionUiThemePrimaryColor {
	if o == nil || o.PrimaryColor == nil {
		var ret FormUiDefinitionUiThemePrimaryColor
		return ret
	}
	return *o.PrimaryColor
}

// GetPrimaryColorOk returns a tuple with the PrimaryColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormUiDefinitionUiTheme) GetPrimaryColorOk() (*FormUiDefinitionUiThemePrimaryColor, bool) {
	if o == nil || o.PrimaryColor == nil {
		return nil, false
	}
	return o.PrimaryColor, true
}

// HasPrimaryColor returns a boolean if a field has been set.
func (o *FormUiDefinitionUiTheme) HasPrimaryColor() bool {
	return o != nil && o.PrimaryColor != nil
}

// SetPrimaryColor gets a reference to the given FormUiDefinitionUiThemePrimaryColor and assigns it to the PrimaryColor field.
func (o *FormUiDefinitionUiTheme) SetPrimaryColor(v FormUiDefinitionUiThemePrimaryColor) {
	o.PrimaryColor = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormUiDefinitionUiTheme) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.PrimaryColor != nil {
		toSerialize["primaryColor"] = o.PrimaryColor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormUiDefinitionUiTheme) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		PrimaryColor *FormUiDefinitionUiThemePrimaryColor `json:"primaryColor,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"primaryColor"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.PrimaryColor != nil && !all.PrimaryColor.IsValid() {
		hasInvalidField = true
	} else {
		o.PrimaryColor = all.PrimaryColor
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
