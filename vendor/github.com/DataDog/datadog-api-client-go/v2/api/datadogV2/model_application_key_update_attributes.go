// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationKeyUpdateAttributes Attributes used to update an application Key.
type ApplicationKeyUpdateAttributes struct {
	// Name of the application key.
	Name *string `json:"name,omitempty"`
	// Array of scopes to grant the application key.
	Scopes datadog.NullableList[string] `json:"scopes,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewApplicationKeyUpdateAttributes instantiates a new ApplicationKeyUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApplicationKeyUpdateAttributes() *ApplicationKeyUpdateAttributes {
	this := ApplicationKeyUpdateAttributes{}
	return &this
}

// NewApplicationKeyUpdateAttributesWithDefaults instantiates a new ApplicationKeyUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApplicationKeyUpdateAttributesWithDefaults() *ApplicationKeyUpdateAttributes {
	this := ApplicationKeyUpdateAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApplicationKeyUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationKeyUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApplicationKeyUpdateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApplicationKeyUpdateAttributes) SetName(v string) {
	o.Name = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationKeyUpdateAttributes) GetScopes() []string {
	if o == nil || o.Scopes.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Scopes.Get()
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ApplicationKeyUpdateAttributes) GetScopesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scopes.Get(), o.Scopes.IsSet()
}

// HasScopes returns a boolean if a field has been set.
func (o *ApplicationKeyUpdateAttributes) HasScopes() bool {
	return o != nil && o.Scopes.IsSet()
}

// SetScopes gets a reference to the given datadog.NullableList[string] and assigns it to the Scopes field.
func (o *ApplicationKeyUpdateAttributes) SetScopes(v []string) {
	o.Scopes.Set(&v)
}

// SetScopesNil sets the value for Scopes to be an explicit nil.
func (o *ApplicationKeyUpdateAttributes) SetScopesNil() {
	o.Scopes.Set(nil)
}

// UnsetScopes ensures that no value is present for Scopes, not even an explicit nil.
func (o *ApplicationKeyUpdateAttributes) UnsetScopes() {
	o.Scopes.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ApplicationKeyUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Scopes.IsSet() {
		toSerialize["scopes"] = o.Scopes.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApplicationKeyUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name   *string                      `json:"name,omitempty"`
		Scopes datadog.NullableList[string] `json:"scopes,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "scopes"})
	} else {
		return err
	}
	o.Name = all.Name
	o.Scopes = all.Scopes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
