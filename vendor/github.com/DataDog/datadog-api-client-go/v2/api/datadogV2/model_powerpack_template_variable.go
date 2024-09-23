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
	// One or many template variable default values within the saved view, which are unioned together using `OR` if more than one is specified.
	Defaults []string `json:"defaults,omitempty"`
	// The name of the variable.
	Name string `json:"name"`
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

// MarshalJSON serializes the struct using spec logic.
func (o PowerpackTemplateVariable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Defaults != nil {
		toSerialize["defaults"] = o.Defaults
	}
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PowerpackTemplateVariable) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Defaults []string `json:"defaults,omitempty"`
		Name     *string  `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"defaults", "name"})
	} else {
		return err
	}
	o.Defaults = all.Defaults
	o.Name = *all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
