// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PartialApplicationKeyAttributes Attributes of a partial application key.
type PartialApplicationKeyAttributes struct {
	// Creation date of the application key.
	CreatedAt *string `json:"created_at,omitempty"`
	// The last four characters of the application key.
	Last4 *string `json:"last4,omitempty"`
	// Last usage timestamp of the application key.
	LastUsedAt datadog.NullableString `json:"last_used_at,omitempty"`
	// Name of the application key.
	Name *string `json:"name,omitempty"`
	// Array of scopes to grant the application key.
	Scopes datadog.NullableList[string] `json:"scopes,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPartialApplicationKeyAttributes instantiates a new PartialApplicationKeyAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPartialApplicationKeyAttributes() *PartialApplicationKeyAttributes {
	this := PartialApplicationKeyAttributes{}
	return &this
}

// NewPartialApplicationKeyAttributesWithDefaults instantiates a new PartialApplicationKeyAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPartialApplicationKeyAttributesWithDefaults() *PartialApplicationKeyAttributes {
	this := PartialApplicationKeyAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PartialApplicationKeyAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialApplicationKeyAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PartialApplicationKeyAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *PartialApplicationKeyAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetLast4 returns the Last4 field value if set, zero value otherwise.
func (o *PartialApplicationKeyAttributes) GetLast4() string {
	if o == nil || o.Last4 == nil {
		var ret string
		return ret
	}
	return *o.Last4
}

// GetLast4Ok returns a tuple with the Last4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialApplicationKeyAttributes) GetLast4Ok() (*string, bool) {
	if o == nil || o.Last4 == nil {
		return nil, false
	}
	return o.Last4, true
}

// HasLast4 returns a boolean if a field has been set.
func (o *PartialApplicationKeyAttributes) HasLast4() bool {
	return o != nil && o.Last4 != nil
}

// SetLast4 gets a reference to the given string and assigns it to the Last4 field.
func (o *PartialApplicationKeyAttributes) SetLast4(v string) {
	o.Last4 = &v
}

// GetLastUsedAt returns the LastUsedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartialApplicationKeyAttributes) GetLastUsedAt() string {
	if o == nil || o.LastUsedAt.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastUsedAt.Get()
}

// GetLastUsedAtOk returns a tuple with the LastUsedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PartialApplicationKeyAttributes) GetLastUsedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUsedAt.Get(), o.LastUsedAt.IsSet()
}

// HasLastUsedAt returns a boolean if a field has been set.
func (o *PartialApplicationKeyAttributes) HasLastUsedAt() bool {
	return o != nil && o.LastUsedAt.IsSet()
}

// SetLastUsedAt gets a reference to the given datadog.NullableString and assigns it to the LastUsedAt field.
func (o *PartialApplicationKeyAttributes) SetLastUsedAt(v string) {
	o.LastUsedAt.Set(&v)
}

// SetLastUsedAtNil sets the value for LastUsedAt to be an explicit nil.
func (o *PartialApplicationKeyAttributes) SetLastUsedAtNil() {
	o.LastUsedAt.Set(nil)
}

// UnsetLastUsedAt ensures that no value is present for LastUsedAt, not even an explicit nil.
func (o *PartialApplicationKeyAttributes) UnsetLastUsedAt() {
	o.LastUsedAt.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PartialApplicationKeyAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialApplicationKeyAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PartialApplicationKeyAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PartialApplicationKeyAttributes) SetName(v string) {
	o.Name = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartialApplicationKeyAttributes) GetScopes() []string {
	if o == nil || o.Scopes.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Scopes.Get()
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PartialApplicationKeyAttributes) GetScopesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scopes.Get(), o.Scopes.IsSet()
}

// HasScopes returns a boolean if a field has been set.
func (o *PartialApplicationKeyAttributes) HasScopes() bool {
	return o != nil && o.Scopes.IsSet()
}

// SetScopes gets a reference to the given datadog.NullableList[string] and assigns it to the Scopes field.
func (o *PartialApplicationKeyAttributes) SetScopes(v []string) {
	o.Scopes.Set(&v)
}

// SetScopesNil sets the value for Scopes to be an explicit nil.
func (o *PartialApplicationKeyAttributes) SetScopesNil() {
	o.Scopes.Set(nil)
}

// UnsetScopes ensures that no value is present for Scopes, not even an explicit nil.
func (o *PartialApplicationKeyAttributes) UnsetScopes() {
	o.Scopes.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o PartialApplicationKeyAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Last4 != nil {
		toSerialize["last4"] = o.Last4
	}
	if o.LastUsedAt.IsSet() {
		toSerialize["last_used_at"] = o.LastUsedAt.Get()
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
func (o *PartialApplicationKeyAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt  *string                      `json:"created_at,omitempty"`
		Last4      *string                      `json:"last4,omitempty"`
		LastUsedAt datadog.NullableString       `json:"last_used_at,omitempty"`
		Name       *string                      `json:"name,omitempty"`
		Scopes     datadog.NullableList[string] `json:"scopes,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "last4", "last_used_at", "name", "scopes"})
	} else {
		return err
	}
	o.CreatedAt = all.CreatedAt
	o.Last4 = all.Last4
	o.LastUsedAt = all.LastUsedAt
	o.Name = all.Name
	o.Scopes = all.Scopes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
