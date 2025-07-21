// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorUserTemplateTemplateVariablesItems List of objects representing template variables on the monitor which can have selectable values.
type MonitorUserTemplateTemplateVariablesItems struct {
	// Available values for the variable.
	AvailableValues []string `json:"available_values,omitempty"`
	// Default values of the template variable.
	Defaults []string `json:"defaults,omitempty"`
	// The name of the template variable.
	Name string `json:"name"`
	// The tag key associated with the variable. This works the same as dashboard template variables.
	TagKey *string `json:"tag_key,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewMonitorUserTemplateTemplateVariablesItems instantiates a new MonitorUserTemplateTemplateVariablesItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorUserTemplateTemplateVariablesItems(name string) *MonitorUserTemplateTemplateVariablesItems {
	this := MonitorUserTemplateTemplateVariablesItems{}
	this.Name = name
	return &this
}

// NewMonitorUserTemplateTemplateVariablesItemsWithDefaults instantiates a new MonitorUserTemplateTemplateVariablesItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorUserTemplateTemplateVariablesItemsWithDefaults() *MonitorUserTemplateTemplateVariablesItems {
	this := MonitorUserTemplateTemplateVariablesItems{}
	return &this
}

// GetAvailableValues returns the AvailableValues field value if set, zero value otherwise.
func (o *MonitorUserTemplateTemplateVariablesItems) GetAvailableValues() []string {
	if o == nil || o.AvailableValues == nil {
		var ret []string
		return ret
	}
	return o.AvailableValues
}

// GetAvailableValuesOk returns a tuple with the AvailableValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorUserTemplateTemplateVariablesItems) GetAvailableValuesOk() (*[]string, bool) {
	if o == nil || o.AvailableValues == nil {
		return nil, false
	}
	return &o.AvailableValues, true
}

// HasAvailableValues returns a boolean if a field has been set.
func (o *MonitorUserTemplateTemplateVariablesItems) HasAvailableValues() bool {
	return o != nil && o.AvailableValues != nil
}

// SetAvailableValues gets a reference to the given []string and assigns it to the AvailableValues field.
func (o *MonitorUserTemplateTemplateVariablesItems) SetAvailableValues(v []string) {
	o.AvailableValues = v
}

// GetDefaults returns the Defaults field value if set, zero value otherwise.
func (o *MonitorUserTemplateTemplateVariablesItems) GetDefaults() []string {
	if o == nil || o.Defaults == nil {
		var ret []string
		return ret
	}
	return o.Defaults
}

// GetDefaultsOk returns a tuple with the Defaults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorUserTemplateTemplateVariablesItems) GetDefaultsOk() (*[]string, bool) {
	if o == nil || o.Defaults == nil {
		return nil, false
	}
	return &o.Defaults, true
}

// HasDefaults returns a boolean if a field has been set.
func (o *MonitorUserTemplateTemplateVariablesItems) HasDefaults() bool {
	return o != nil && o.Defaults != nil
}

// SetDefaults gets a reference to the given []string and assigns it to the Defaults field.
func (o *MonitorUserTemplateTemplateVariablesItems) SetDefaults(v []string) {
	o.Defaults = v
}

// GetName returns the Name field value.
func (o *MonitorUserTemplateTemplateVariablesItems) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MonitorUserTemplateTemplateVariablesItems) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *MonitorUserTemplateTemplateVariablesItems) SetName(v string) {
	o.Name = v
}

// GetTagKey returns the TagKey field value if set, zero value otherwise.
func (o *MonitorUserTemplateTemplateVariablesItems) GetTagKey() string {
	if o == nil || o.TagKey == nil {
		var ret string
		return ret
	}
	return *o.TagKey
}

// GetTagKeyOk returns a tuple with the TagKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorUserTemplateTemplateVariablesItems) GetTagKeyOk() (*string, bool) {
	if o == nil || o.TagKey == nil {
		return nil, false
	}
	return o.TagKey, true
}

// HasTagKey returns a boolean if a field has been set.
func (o *MonitorUserTemplateTemplateVariablesItems) HasTagKey() bool {
	return o != nil && o.TagKey != nil
}

// SetTagKey gets a reference to the given string and assigns it to the TagKey field.
func (o *MonitorUserTemplateTemplateVariablesItems) SetTagKey(v string) {
	o.TagKey = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorUserTemplateTemplateVariablesItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AvailableValues != nil {
		toSerialize["available_values"] = o.AvailableValues
	}
	if o.Defaults != nil {
		toSerialize["defaults"] = o.Defaults
	}
	toSerialize["name"] = o.Name
	if o.TagKey != nil {
		toSerialize["tag_key"] = o.TagKey
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorUserTemplateTemplateVariablesItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AvailableValues []string `json:"available_values,omitempty"`
		Defaults        []string `json:"defaults,omitempty"`
		Name            *string  `json:"name"`
		TagKey          *string  `json:"tag_key,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	o.AvailableValues = all.AvailableValues
	o.Defaults = all.Defaults
	o.Name = *all.Name
	o.TagKey = all.TagKey

	return nil
}
