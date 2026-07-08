// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReportScheduleIndexTemplateVariable Template variable metadata from a dashboard index.
type ReportScheduleIndexTemplateVariable struct {
	// Available values for the template variable.
	AvailableValues datadog.NullableList[string] `json:"available_values,omitempty"`
	// Default values for the template variable.
	Defaults datadog.NullableList[string] `json:"defaults,omitempty"`
	// The template variable name.
	Name datadog.NullableString `json:"name,omitempty"`
	// The tag prefix for the template variable, when available.
	Prefix datadog.NullableString `json:"prefix,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewReportScheduleIndexTemplateVariable instantiates a new ReportScheduleIndexTemplateVariable object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReportScheduleIndexTemplateVariable() *ReportScheduleIndexTemplateVariable {
	this := ReportScheduleIndexTemplateVariable{}
	return &this
}

// NewReportScheduleIndexTemplateVariableWithDefaults instantiates a new ReportScheduleIndexTemplateVariable object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReportScheduleIndexTemplateVariableWithDefaults() *ReportScheduleIndexTemplateVariable {
	this := ReportScheduleIndexTemplateVariable{}
	return &this
}

// GetAvailableValues returns the AvailableValues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportScheduleIndexTemplateVariable) GetAvailableValues() []string {
	if o == nil || o.AvailableValues.Get() == nil {
		var ret []string
		return ret
	}
	return *o.AvailableValues.Get()
}

// GetAvailableValuesOk returns a tuple with the AvailableValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportScheduleIndexTemplateVariable) GetAvailableValuesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailableValues.Get(), o.AvailableValues.IsSet()
}

// HasAvailableValues returns a boolean if a field has been set.
func (o *ReportScheduleIndexTemplateVariable) HasAvailableValues() bool {
	return o != nil && o.AvailableValues.IsSet()
}

// SetAvailableValues gets a reference to the given datadog.NullableList[string] and assigns it to the AvailableValues field.
func (o *ReportScheduleIndexTemplateVariable) SetAvailableValues(v []string) {
	o.AvailableValues.Set(&v)
}

// SetAvailableValuesNil sets the value for AvailableValues to be an explicit nil.
func (o *ReportScheduleIndexTemplateVariable) SetAvailableValuesNil() {
	o.AvailableValues.Set(nil)
}

// UnsetAvailableValues ensures that no value is present for AvailableValues, not even an explicit nil.
func (o *ReportScheduleIndexTemplateVariable) UnsetAvailableValues() {
	o.AvailableValues.Unset()
}

// GetDefaults returns the Defaults field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportScheduleIndexTemplateVariable) GetDefaults() []string {
	if o == nil || o.Defaults.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Defaults.Get()
}

// GetDefaultsOk returns a tuple with the Defaults field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportScheduleIndexTemplateVariable) GetDefaultsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Defaults.Get(), o.Defaults.IsSet()
}

// HasDefaults returns a boolean if a field has been set.
func (o *ReportScheduleIndexTemplateVariable) HasDefaults() bool {
	return o != nil && o.Defaults.IsSet()
}

// SetDefaults gets a reference to the given datadog.NullableList[string] and assigns it to the Defaults field.
func (o *ReportScheduleIndexTemplateVariable) SetDefaults(v []string) {
	o.Defaults.Set(&v)
}

// SetDefaultsNil sets the value for Defaults to be an explicit nil.
func (o *ReportScheduleIndexTemplateVariable) SetDefaultsNil() {
	o.Defaults.Set(nil)
}

// UnsetDefaults ensures that no value is present for Defaults, not even an explicit nil.
func (o *ReportScheduleIndexTemplateVariable) UnsetDefaults() {
	o.Defaults.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportScheduleIndexTemplateVariable) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportScheduleIndexTemplateVariable) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ReportScheduleIndexTemplateVariable) HasName() bool {
	return o != nil && o.Name.IsSet()
}

// SetName gets a reference to the given datadog.NullableString and assigns it to the Name field.
func (o *ReportScheduleIndexTemplateVariable) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil.
func (o *ReportScheduleIndexTemplateVariable) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil.
func (o *ReportScheduleIndexTemplateVariable) UnsetName() {
	o.Name.Unset()
}

// GetPrefix returns the Prefix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportScheduleIndexTemplateVariable) GetPrefix() string {
	if o == nil || o.Prefix.Get() == nil {
		var ret string
		return ret
	}
	return *o.Prefix.Get()
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportScheduleIndexTemplateVariable) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prefix.Get(), o.Prefix.IsSet()
}

// HasPrefix returns a boolean if a field has been set.
func (o *ReportScheduleIndexTemplateVariable) HasPrefix() bool {
	return o != nil && o.Prefix.IsSet()
}

// SetPrefix gets a reference to the given datadog.NullableString and assigns it to the Prefix field.
func (o *ReportScheduleIndexTemplateVariable) SetPrefix(v string) {
	o.Prefix.Set(&v)
}

// SetPrefixNil sets the value for Prefix to be an explicit nil.
func (o *ReportScheduleIndexTemplateVariable) SetPrefixNil() {
	o.Prefix.Set(nil)
}

// UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil.
func (o *ReportScheduleIndexTemplateVariable) UnsetPrefix() {
	o.Prefix.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ReportScheduleIndexTemplateVariable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AvailableValues.IsSet() {
		toSerialize["available_values"] = o.AvailableValues.Get()
	}
	if o.Defaults.IsSet() {
		toSerialize["defaults"] = o.Defaults.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Prefix.IsSet() {
		toSerialize["prefix"] = o.Prefix.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ReportScheduleIndexTemplateVariable) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AvailableValues datadog.NullableList[string] `json:"available_values,omitempty"`
		Defaults        datadog.NullableList[string] `json:"defaults,omitempty"`
		Name            datadog.NullableString       `json:"name,omitempty"`
		Prefix          datadog.NullableString       `json:"prefix,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"available_values", "defaults", "name", "prefix"})
	} else {
		return err
	}
	o.AvailableValues = all.AvailableValues
	o.Defaults = all.Defaults
	o.Name = all.Name
	o.Prefix = all.Prefix

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
