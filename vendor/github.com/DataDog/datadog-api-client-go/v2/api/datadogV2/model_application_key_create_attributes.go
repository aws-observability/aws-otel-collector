// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationKeyCreateAttributes Attributes used to create an application Key.
type ApplicationKeyCreateAttributes struct {
	// Name of the application key.
	Name string `json:"name"`
	// Array of scopes to grant the application key.
	Scopes datadog.NullableList[string] `json:"scopes,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewApplicationKeyCreateAttributes instantiates a new ApplicationKeyCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApplicationKeyCreateAttributes(name string) *ApplicationKeyCreateAttributes {
	this := ApplicationKeyCreateAttributes{}
	this.Name = name
	return &this
}

// NewApplicationKeyCreateAttributesWithDefaults instantiates a new ApplicationKeyCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApplicationKeyCreateAttributesWithDefaults() *ApplicationKeyCreateAttributes {
	this := ApplicationKeyCreateAttributes{}
	return &this
}

// GetName returns the Name field value.
func (o *ApplicationKeyCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplicationKeyCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ApplicationKeyCreateAttributes) SetName(v string) {
	o.Name = v
}

// GetScopes returns the Scopes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationKeyCreateAttributes) GetScopes() []string {
	if o == nil || o.Scopes.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Scopes.Get()
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ApplicationKeyCreateAttributes) GetScopesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scopes.Get(), o.Scopes.IsSet()
}

// HasScopes returns a boolean if a field has been set.
func (o *ApplicationKeyCreateAttributes) HasScopes() bool {
	return o != nil && o.Scopes.IsSet()
}

// SetScopes gets a reference to the given datadog.NullableList[string] and assigns it to the Scopes field.
func (o *ApplicationKeyCreateAttributes) SetScopes(v []string) {
	o.Scopes.Set(&v)
}

// SetScopesNil sets the value for Scopes to be an explicit nil.
func (o *ApplicationKeyCreateAttributes) SetScopesNil() {
	o.Scopes.Set(nil)
}

// UnsetScopes ensures that no value is present for Scopes, not even an explicit nil.
func (o *ApplicationKeyCreateAttributes) UnsetScopes() {
	o.Scopes.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ApplicationKeyCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	if o.Scopes.IsSet() {
		toSerialize["scopes"] = o.Scopes.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApplicationKeyCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name   *string                      `json:"name"`
		Scopes datadog.NullableList[string] `json:"scopes,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "scopes"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Scopes = all.Scopes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
