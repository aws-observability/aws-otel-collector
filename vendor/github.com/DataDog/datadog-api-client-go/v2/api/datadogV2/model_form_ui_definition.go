// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormUiDefinition UI configuration for rendering form fields, including widget overrides, field ordering, and themes.
type FormUiDefinition struct {
	// The order in which form fields are displayed.
	UiOrder []string `json:"ui:order,omitempty"`
	// The visual theme applied to the form.
	UiTheme *FormUiDefinitionUiTheme `json:"ui:theme,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFormUiDefinition instantiates a new FormUiDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormUiDefinition() *FormUiDefinition {
	this := FormUiDefinition{}
	return &this
}

// NewFormUiDefinitionWithDefaults instantiates a new FormUiDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormUiDefinitionWithDefaults() *FormUiDefinition {
	this := FormUiDefinition{}
	return &this
}

// GetUiOrder returns the UiOrder field value if set, zero value otherwise.
func (o *FormUiDefinition) GetUiOrder() []string {
	if o == nil || o.UiOrder == nil {
		var ret []string
		return ret
	}
	return o.UiOrder
}

// GetUiOrderOk returns a tuple with the UiOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormUiDefinition) GetUiOrderOk() (*[]string, bool) {
	if o == nil || o.UiOrder == nil {
		return nil, false
	}
	return &o.UiOrder, true
}

// HasUiOrder returns a boolean if a field has been set.
func (o *FormUiDefinition) HasUiOrder() bool {
	return o != nil && o.UiOrder != nil
}

// SetUiOrder gets a reference to the given []string and assigns it to the UiOrder field.
func (o *FormUiDefinition) SetUiOrder(v []string) {
	o.UiOrder = v
}

// GetUiTheme returns the UiTheme field value if set, zero value otherwise.
func (o *FormUiDefinition) GetUiTheme() FormUiDefinitionUiTheme {
	if o == nil || o.UiTheme == nil {
		var ret FormUiDefinitionUiTheme
		return ret
	}
	return *o.UiTheme
}

// GetUiThemeOk returns a tuple with the UiTheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormUiDefinition) GetUiThemeOk() (*FormUiDefinitionUiTheme, bool) {
	if o == nil || o.UiTheme == nil {
		return nil, false
	}
	return o.UiTheme, true
}

// HasUiTheme returns a boolean if a field has been set.
func (o *FormUiDefinition) HasUiTheme() bool {
	return o != nil && o.UiTheme != nil
}

// SetUiTheme gets a reference to the given FormUiDefinitionUiTheme and assigns it to the UiTheme field.
func (o *FormUiDefinition) SetUiTheme(v FormUiDefinitionUiTheme) {
	o.UiTheme = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormUiDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.UiOrder != nil {
		toSerialize["ui:order"] = o.UiOrder
	}
	if o.UiTheme != nil {
		toSerialize["ui:theme"] = o.UiTheme
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormUiDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		UiOrder []string                 `json:"ui:order,omitempty"`
		UiTheme *FormUiDefinitionUiTheme `json:"ui:theme,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"ui:order", "ui:theme"})
	} else {
		return err
	}

	hasInvalidField := false
	o.UiOrder = all.UiOrder
	if all.UiTheme != nil && all.UiTheme.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.UiTheme = all.UiTheme

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
