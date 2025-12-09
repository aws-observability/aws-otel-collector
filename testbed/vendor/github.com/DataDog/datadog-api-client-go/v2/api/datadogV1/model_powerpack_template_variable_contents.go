// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PowerpackTemplateVariableContents Powerpack template variable contents.
type PowerpackTemplateVariableContents struct {
	// The name of the variable.
	Name string `json:"name"`
	// The tag prefix associated with the variable.
	Prefix *string `json:"prefix,omitempty"`
	// One or many template variable values within the saved view, which will be unioned together using `OR` if more than one is specified.
	Values []string `json:"values"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPowerpackTemplateVariableContents instantiates a new PowerpackTemplateVariableContents object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPowerpackTemplateVariableContents(name string, values []string) *PowerpackTemplateVariableContents {
	this := PowerpackTemplateVariableContents{}
	this.Name = name
	this.Values = values
	return &this
}

// NewPowerpackTemplateVariableContentsWithDefaults instantiates a new PowerpackTemplateVariableContents object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPowerpackTemplateVariableContentsWithDefaults() *PowerpackTemplateVariableContents {
	this := PowerpackTemplateVariableContents{}
	return &this
}

// GetName returns the Name field value.
func (o *PowerpackTemplateVariableContents) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PowerpackTemplateVariableContents) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *PowerpackTemplateVariableContents) SetName(v string) {
	o.Name = v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *PowerpackTemplateVariableContents) GetPrefix() string {
	if o == nil || o.Prefix == nil {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerpackTemplateVariableContents) GetPrefixOk() (*string, bool) {
	if o == nil || o.Prefix == nil {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *PowerpackTemplateVariableContents) HasPrefix() bool {
	return o != nil && o.Prefix != nil
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *PowerpackTemplateVariableContents) SetPrefix(v string) {
	o.Prefix = &v
}

// GetValues returns the Values field value.
func (o *PowerpackTemplateVariableContents) GetValues() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *PowerpackTemplateVariableContents) GetValuesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Values, true
}

// SetValues sets field value.
func (o *PowerpackTemplateVariableContents) SetValues(v []string) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PowerpackTemplateVariableContents) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	if o.Prefix != nil {
		toSerialize["prefix"] = o.Prefix
	}
	toSerialize["values"] = o.Values

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PowerpackTemplateVariableContents) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name   *string   `json:"name"`
		Prefix *string   `json:"prefix,omitempty"`
		Values *[]string `json:"values"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Values == nil {
		return fmt.Errorf("required field values missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "prefix", "values"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Prefix = all.Prefix
	o.Values = *all.Values

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
