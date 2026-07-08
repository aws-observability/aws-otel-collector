// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceAccessTokenAttributes Attributes of an access token.
type ServiceAccessTokenAttributes struct {
	// Creation date of the access token.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Expiration date of the access token.
	ExpiresAt datadog.NullableTime `json:"expires_at,omitempty"`
	// Date the access token was last used.
	LastUsedAt datadog.NullableTime `json:"last_used_at,omitempty"`
	// Date of last modification of the access token.
	ModifiedAt datadog.NullableTime `json:"modified_at,omitempty"`
	// Name of the access token.
	Name *string `json:"name,omitempty"`
	// The public portion of the access token.
	PublicPortion *string `json:"public_portion,omitempty"`
	// Array of scopes granted to the access token.
	Scopes []string `json:"scopes,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewServiceAccessTokenAttributes instantiates a new ServiceAccessTokenAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewServiceAccessTokenAttributes() *ServiceAccessTokenAttributes {
	this := ServiceAccessTokenAttributes{}
	return &this
}

// NewServiceAccessTokenAttributesWithDefaults instantiates a new ServiceAccessTokenAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewServiceAccessTokenAttributesWithDefaults() *ServiceAccessTokenAttributes {
	this := ServiceAccessTokenAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ServiceAccessTokenAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccessTokenAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ServiceAccessTokenAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ServiceAccessTokenAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceAccessTokenAttributes) GetExpiresAt() time.Time {
	if o == nil || o.ExpiresAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ServiceAccessTokenAttributes) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *ServiceAccessTokenAttributes) HasExpiresAt() bool {
	return o != nil && o.ExpiresAt.IsSet()
}

// SetExpiresAt gets a reference to the given datadog.NullableTime and assigns it to the ExpiresAt field.
func (o *ServiceAccessTokenAttributes) SetExpiresAt(v time.Time) {
	o.ExpiresAt.Set(&v)
}

// SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil.
func (o *ServiceAccessTokenAttributes) SetExpiresAtNil() {
	o.ExpiresAt.Set(nil)
}

// UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil.
func (o *ServiceAccessTokenAttributes) UnsetExpiresAt() {
	o.ExpiresAt.Unset()
}

// GetLastUsedAt returns the LastUsedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceAccessTokenAttributes) GetLastUsedAt() time.Time {
	if o == nil || o.LastUsedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUsedAt.Get()
}

// GetLastUsedAtOk returns a tuple with the LastUsedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ServiceAccessTokenAttributes) GetLastUsedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUsedAt.Get(), o.LastUsedAt.IsSet()
}

// HasLastUsedAt returns a boolean if a field has been set.
func (o *ServiceAccessTokenAttributes) HasLastUsedAt() bool {
	return o != nil && o.LastUsedAt.IsSet()
}

// SetLastUsedAt gets a reference to the given datadog.NullableTime and assigns it to the LastUsedAt field.
func (o *ServiceAccessTokenAttributes) SetLastUsedAt(v time.Time) {
	o.LastUsedAt.Set(&v)
}

// SetLastUsedAtNil sets the value for LastUsedAt to be an explicit nil.
func (o *ServiceAccessTokenAttributes) SetLastUsedAtNil() {
	o.LastUsedAt.Set(nil)
}

// UnsetLastUsedAt ensures that no value is present for LastUsedAt, not even an explicit nil.
func (o *ServiceAccessTokenAttributes) UnsetLastUsedAt() {
	o.LastUsedAt.Unset()
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceAccessTokenAttributes) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt.Get()
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ServiceAccessTokenAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedAt.Get(), o.ModifiedAt.IsSet()
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *ServiceAccessTokenAttributes) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt.IsSet()
}

// SetModifiedAt gets a reference to the given datadog.NullableTime and assigns it to the ModifiedAt field.
func (o *ServiceAccessTokenAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt.Set(&v)
}

// SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil.
func (o *ServiceAccessTokenAttributes) SetModifiedAtNil() {
	o.ModifiedAt.Set(nil)
}

// UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil.
func (o *ServiceAccessTokenAttributes) UnsetModifiedAt() {
	o.ModifiedAt.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServiceAccessTokenAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccessTokenAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceAccessTokenAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceAccessTokenAttributes) SetName(v string) {
	o.Name = &v
}

// GetPublicPortion returns the PublicPortion field value if set, zero value otherwise.
func (o *ServiceAccessTokenAttributes) GetPublicPortion() string {
	if o == nil || o.PublicPortion == nil {
		var ret string
		return ret
	}
	return *o.PublicPortion
}

// GetPublicPortionOk returns a tuple with the PublicPortion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccessTokenAttributes) GetPublicPortionOk() (*string, bool) {
	if o == nil || o.PublicPortion == nil {
		return nil, false
	}
	return o.PublicPortion, true
}

// HasPublicPortion returns a boolean if a field has been set.
func (o *ServiceAccessTokenAttributes) HasPublicPortion() bool {
	return o != nil && o.PublicPortion != nil
}

// SetPublicPortion gets a reference to the given string and assigns it to the PublicPortion field.
func (o *ServiceAccessTokenAttributes) SetPublicPortion(v string) {
	o.PublicPortion = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *ServiceAccessTokenAttributes) GetScopes() []string {
	if o == nil || o.Scopes == nil {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccessTokenAttributes) GetScopesOk() (*[]string, bool) {
	if o == nil || o.Scopes == nil {
		return nil, false
	}
	return &o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *ServiceAccessTokenAttributes) HasScopes() bool {
	return o != nil && o.Scopes != nil
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *ServiceAccessTokenAttributes) SetScopes(v []string) {
	o.Scopes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ServiceAccessTokenAttributes) MarshalJSON() ([]byte, error) {
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
	if o.LastUsedAt.IsSet() {
		toSerialize["last_used_at"] = o.LastUsedAt.Get()
	}
	if o.ModifiedAt.IsSet() {
		toSerialize["modified_at"] = o.ModifiedAt.Get()
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
func (o *ServiceAccessTokenAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt     *time.Time           `json:"created_at,omitempty"`
		ExpiresAt     datadog.NullableTime `json:"expires_at,omitempty"`
		LastUsedAt    datadog.NullableTime `json:"last_used_at,omitempty"`
		ModifiedAt    datadog.NullableTime `json:"modified_at,omitempty"`
		Name          *string              `json:"name,omitempty"`
		PublicPortion *string              `json:"public_portion,omitempty"`
		Scopes        []string             `json:"scopes,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "expires_at", "last_used_at", "modified_at", "name", "public_portion", "scopes"})
	} else {
		return err
	}
	o.CreatedAt = all.CreatedAt
	o.ExpiresAt = all.ExpiresAt
	o.LastUsedAt = all.LastUsedAt
	o.ModifiedAt = all.ModifiedAt
	o.Name = all.Name
	o.PublicPortion = all.PublicPortion
	o.Scopes = all.Scopes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
