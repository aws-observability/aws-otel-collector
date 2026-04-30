// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PersonalAccessTokenUpdateAttributes Attributes used to update a personal access token.
type PersonalAccessTokenUpdateAttributes struct {
	// Name of the personal access token.
	Name *string `json:"name,omitempty"`
	// Array of scopes to grant the personal access token.
	Scopes []string `json:"scopes,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPersonalAccessTokenUpdateAttributes instantiates a new PersonalAccessTokenUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPersonalAccessTokenUpdateAttributes() *PersonalAccessTokenUpdateAttributes {
	this := PersonalAccessTokenUpdateAttributes{}
	return &this
}

// NewPersonalAccessTokenUpdateAttributesWithDefaults instantiates a new PersonalAccessTokenUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPersonalAccessTokenUpdateAttributesWithDefaults() *PersonalAccessTokenUpdateAttributes {
	this := PersonalAccessTokenUpdateAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PersonalAccessTokenUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonalAccessTokenUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PersonalAccessTokenUpdateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PersonalAccessTokenUpdateAttributes) SetName(v string) {
	o.Name = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *PersonalAccessTokenUpdateAttributes) GetScopes() []string {
	if o == nil || o.Scopes == nil {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonalAccessTokenUpdateAttributes) GetScopesOk() (*[]string, bool) {
	if o == nil || o.Scopes == nil {
		return nil, false
	}
	return &o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *PersonalAccessTokenUpdateAttributes) HasScopes() bool {
	return o != nil && o.Scopes != nil
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *PersonalAccessTokenUpdateAttributes) SetScopes(v []string) {
	o.Scopes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PersonalAccessTokenUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PersonalAccessTokenUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name   *string  `json:"name,omitempty"`
		Scopes []string `json:"scopes,omitempty"`
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
