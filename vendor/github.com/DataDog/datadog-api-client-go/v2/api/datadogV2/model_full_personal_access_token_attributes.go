// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FullPersonalAccessTokenAttributes Attributes of a full personal access token, including the token key.
type FullPersonalAccessTokenAttributes struct {
	// Creation date of the personal access token.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Expiration date of the personal access token.
	ExpiresAt datadog.NullableTime `json:"expires_at,omitempty"`
	// The personal access token key. Only returned upon creation.
	Key *string `json:"key,omitempty"`
	// Name of the personal access token.
	Name *string `json:"name,omitempty"`
	// The public portion of the personal access token.
	PublicPortion *string `json:"public_portion,omitempty"`
	// Array of scopes granted to the personal access token.
	Scopes []string `json:"scopes,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFullPersonalAccessTokenAttributes instantiates a new FullPersonalAccessTokenAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFullPersonalAccessTokenAttributes() *FullPersonalAccessTokenAttributes {
	this := FullPersonalAccessTokenAttributes{}
	return &this
}

// NewFullPersonalAccessTokenAttributesWithDefaults instantiates a new FullPersonalAccessTokenAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFullPersonalAccessTokenAttributesWithDefaults() *FullPersonalAccessTokenAttributes {
	this := FullPersonalAccessTokenAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *FullPersonalAccessTokenAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullPersonalAccessTokenAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *FullPersonalAccessTokenAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *FullPersonalAccessTokenAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FullPersonalAccessTokenAttributes) GetExpiresAt() time.Time {
	if o == nil || o.ExpiresAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FullPersonalAccessTokenAttributes) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *FullPersonalAccessTokenAttributes) HasExpiresAt() bool {
	return o != nil && o.ExpiresAt.IsSet()
}

// SetExpiresAt gets a reference to the given datadog.NullableTime and assigns it to the ExpiresAt field.
func (o *FullPersonalAccessTokenAttributes) SetExpiresAt(v time.Time) {
	o.ExpiresAt.Set(&v)
}

// SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil.
func (o *FullPersonalAccessTokenAttributes) SetExpiresAtNil() {
	o.ExpiresAt.Set(nil)
}

// UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil.
func (o *FullPersonalAccessTokenAttributes) UnsetExpiresAt() {
	o.ExpiresAt.Unset()
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *FullPersonalAccessTokenAttributes) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullPersonalAccessTokenAttributes) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *FullPersonalAccessTokenAttributes) HasKey() bool {
	return o != nil && o.Key != nil
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *FullPersonalAccessTokenAttributes) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FullPersonalAccessTokenAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullPersonalAccessTokenAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FullPersonalAccessTokenAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FullPersonalAccessTokenAttributes) SetName(v string) {
	o.Name = &v
}

// GetPublicPortion returns the PublicPortion field value if set, zero value otherwise.
func (o *FullPersonalAccessTokenAttributes) GetPublicPortion() string {
	if o == nil || o.PublicPortion == nil {
		var ret string
		return ret
	}
	return *o.PublicPortion
}

// GetPublicPortionOk returns a tuple with the PublicPortion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullPersonalAccessTokenAttributes) GetPublicPortionOk() (*string, bool) {
	if o == nil || o.PublicPortion == nil {
		return nil, false
	}
	return o.PublicPortion, true
}

// HasPublicPortion returns a boolean if a field has been set.
func (o *FullPersonalAccessTokenAttributes) HasPublicPortion() bool {
	return o != nil && o.PublicPortion != nil
}

// SetPublicPortion gets a reference to the given string and assigns it to the PublicPortion field.
func (o *FullPersonalAccessTokenAttributes) SetPublicPortion(v string) {
	o.PublicPortion = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *FullPersonalAccessTokenAttributes) GetScopes() []string {
	if o == nil || o.Scopes == nil {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullPersonalAccessTokenAttributes) GetScopesOk() (*[]string, bool) {
	if o == nil || o.Scopes == nil {
		return nil, false
	}
	return &o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *FullPersonalAccessTokenAttributes) HasScopes() bool {
	return o != nil && o.Scopes != nil
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *FullPersonalAccessTokenAttributes) SetScopes(v []string) {
	o.Scopes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FullPersonalAccessTokenAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.ExpiresAt.IsSet() {
		toSerialize["expires_at"] = o.ExpiresAt.Get()
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PublicPortion != nil {
		toSerialize["public_portion"] = o.PublicPortion
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
func (o *FullPersonalAccessTokenAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt     *time.Time           `json:"created_at,omitempty"`
		ExpiresAt     datadog.NullableTime `json:"expires_at,omitempty"`
		Key           *string              `json:"key,omitempty"`
		Name          *string              `json:"name,omitempty"`
		PublicPortion *string              `json:"public_portion,omitempty"`
		Scopes        []string             `json:"scopes,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "expires_at", "key", "name", "public_portion", "scopes"})
	} else {
		return err
	}
	o.CreatedAt = all.CreatedAt
	o.ExpiresAt = all.ExpiresAt
	o.Key = all.Key
	o.Name = all.Name
	o.PublicPortion = all.PublicPortion
	o.Scopes = all.Scopes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
