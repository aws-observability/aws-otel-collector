// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PowerpackTemplateVariable Powerpack template variables.
type PowerpackTemplateVariable struct {
	// The list of values that the template variable drop-down is limited to.
	AvailableValues datadog.NullableList[string] `json:"available_values,omitempty"`
	// One or many template variable default values within the saved view, which are unioned together using `OR` if more than one is specified.
	Defaults []string `json:"defaults,omitempty"`
	// The name of the variable.
	Name string `json:"name"`
	// The tag prefix associated with the variable. Only tags with this prefix appear in the variable drop-down.
	Prefix datadog.NullableString `json:"prefix,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPowerpackTemplateVariable instantiates a new PowerpackTemplateVariable object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPowerpackTemplateVariable(name string) *PowerpackTemplateVariable {
	this := PowerpackTemplateVariable{}
	this.Name = name
	return &this
}

// NewPowerpackTemplateVariableWithDefaults instantiates a new PowerpackTemplateVariable object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPowerpackTemplateVariableWithDefaults() *PowerpackTemplateVariable {
	this := PowerpackTemplateVariable{}
	return &this
}

// GetAvailableValues returns the AvailableValues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PowerpackTemplateVariable) GetAvailableValues() []string {
	if o == nil || o.AvailableValues.Get() == nil {
		var ret []string
		return ret
	}
	return *o.AvailableValues.Get()
}

// GetAvailableValuesOk returns a tuple with the AvailableValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PowerpackTemplateVariable) GetAvailableValuesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailableValues.Get(), o.AvailableValues.IsSet()
}

// HasAvailableValues returns a boolean if a field has been set.
func (o *PowerpackTemplateVariable) HasAvailableValues() bool {
	return o != nil && o.AvailableValues.IsSet()
}

// SetAvailableValues gets a reference to the given datadog.NullableList[string] and assigns it to the AvailableValues field.
func (o *PowerpackTemplateVariable) SetAvailableValues(v []string) {
	o.AvailableValues.Set(&v)
}

// SetAvailableValuesNil sets the value for AvailableValues to be an explicit nil.
func (o *PowerpackTemplateVariable) SetAvailableValuesNil() {
	o.AvailableValues.Set(nil)
}

// UnsetAvailableValues ensures that no value is present for AvailableValues, not even an explicit nil.
func (o *PowerpackTemplateVariable) UnsetAvailableValues() {
	o.AvailableValues.Unset()
}

// GetDefaults returns the Defaults field value if set, zero value otherwise.
func (o *PowerpackTemplateVariable) GetDefaults() []string {
	if o == nil || o.Defaults == nil {
		var ret []string
		return ret
	}
	return o.Defaults
}

// GetDefaultsOk returns a tuple with the Defaults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerpackTemplateVariable) GetDefaultsOk() (*[]string, bool) {
	if o == nil || o.Defaults == nil {
		return nil, false
	}
	return &o.Defaults, true
}

// HasDefaults returns a boolean if a field has been set.
func (o *PowerpackTemplateVariable) HasDefaults() bool {
	return o != nil && o.Defaults != nil
}

// SetDefaults gets a reference to the given []string and assigns it to the Defaults field.
func (o *PowerpackTemplateVariable) SetDefaults(v []string) {
	o.Defaults = v
}

// GetName returns the Name field value.
func (o *PowerpackTemplateVariable) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PowerpackTemplateVariable) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *PowerpackTemplateVariable) SetName(v string) {
	o.Name = v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PowerpackTemplateVariable) GetPrefix() string {
	if o == nil || o.Prefix.Get() == nil {
		var ret string
		return ret
	}
	return *o.Prefix.Get()
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PowerpackTemplateVariable) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prefix.Get(), o.Prefix.IsSet()
}

// HasPrefix returns a boolean if a field has been set.
func (o *PowerpackTemplateVariable) HasPrefix() bool {
	return o != nil && o.Prefix.IsSet()
}

// SetPrefix gets a reference to the given datadog.NullableString and assigns it to the Prefix field.
func (o *PowerpackTemplateVariable) SetPrefix(v string) {
	o.Prefix.Set(&v)
}

// SetPrefixNil sets the value for Prefix to be an explicit nil.
func (o *PowerpackTemplateVariable) SetPrefixNil() {
	o.Prefix.Set(nil)
}

// UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil.
func (o *PowerpackTemplateVariable) UnsetPrefix() {
	o.Prefix.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o PowerpackTemplateVariable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AvailableValues.IsSet() {
		toSerialize["available_values"] = o.AvailableValues.Get()
	}
	if o.Defaults != nil {
		toSerialize["defaults"] = o.Defaults
	}
	toSerialize["name"] = o.Name
	if o.Prefix.IsSet() {
		toSerialize["prefix"] = o.Prefix.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PowerpackTemplateVariable) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AvailableValues datadog.NullableList[string] `json:"available_values,omitempty"`
		Defaults        []string                     `json:"defaults,omitempty"`
		Name            *string                      `json:"name"`
		Prefix          datadog.NullableString       `json:"prefix,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"available_values", "defaults", "name", "prefix"})
	} else {
		return err
	}
	o.AvailableValues = all.AvailableValues
	o.Defaults = all.Defaults
	o.Name = *all.Name
	o.Prefix = all.Prefix

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
