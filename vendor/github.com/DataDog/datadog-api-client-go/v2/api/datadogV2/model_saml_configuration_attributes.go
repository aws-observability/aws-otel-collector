// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SAMLConfigurationAttributes Attributes of a SAML configuration.
type SAMLConfigurationAttributes struct {
	// The assertion consumer service (ACS) URLs that the identity provider posts SAML responses to.
	AssertionConsumerService []string `json:"assertion_consumer_service,omitempty"`
	// Creation time of the SAML configuration.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The service provider entity ID Datadog presents to the identity provider.
	EntityId *string `json:"entity_id,omitempty"`
	// Expiration time of the uploaded identity provider metadata.
	ExpiresAt datadog.NullableTime `json:"expires_at,omitempty"`
	// Whether identity-provider-initiated login is enabled for the organization.
	IdpInitiated *bool `json:"idp_initiated,omitempty"`
	// Email domains for which users are automatically provisioned on first SAML login
	// (just-in-time provisioning).
	JitDomains []string `json:"jit_domains,omitempty"`
	// Time of the last SAML configuration modification.
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// The single sign-on URL users can visit to start a SAML login.
	// Returns `null` when the organization is identity-provider-initiated and has no subdomain.
	SsoUrl datadog.NullableString `json:"sso_url,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSAMLConfigurationAttributes instantiates a new SAMLConfigurationAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSAMLConfigurationAttributes() *SAMLConfigurationAttributes {
	this := SAMLConfigurationAttributes{}
	return &this
}

// NewSAMLConfigurationAttributesWithDefaults instantiates a new SAMLConfigurationAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSAMLConfigurationAttributesWithDefaults() *SAMLConfigurationAttributes {
	this := SAMLConfigurationAttributes{}
	return &this
}

// GetAssertionConsumerService returns the AssertionConsumerService field value if set, zero value otherwise.
func (o *SAMLConfigurationAttributes) GetAssertionConsumerService() []string {
	if o == nil || o.AssertionConsumerService == nil {
		var ret []string
		return ret
	}
	return o.AssertionConsumerService
}

// GetAssertionConsumerServiceOk returns a tuple with the AssertionConsumerService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLConfigurationAttributes) GetAssertionConsumerServiceOk() (*[]string, bool) {
	if o == nil || o.AssertionConsumerService == nil {
		return nil, false
	}
	return &o.AssertionConsumerService, true
}

// HasAssertionConsumerService returns a boolean if a field has been set.
func (o *SAMLConfigurationAttributes) HasAssertionConsumerService() bool {
	return o != nil && o.AssertionConsumerService != nil
}

// SetAssertionConsumerService gets a reference to the given []string and assigns it to the AssertionConsumerService field.
func (o *SAMLConfigurationAttributes) SetAssertionConsumerService(v []string) {
	o.AssertionConsumerService = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SAMLConfigurationAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLConfigurationAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SAMLConfigurationAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *SAMLConfigurationAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *SAMLConfigurationAttributes) GetEntityId() string {
	if o == nil || o.EntityId == nil {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLConfigurationAttributes) GetEntityIdOk() (*string, bool) {
	if o == nil || o.EntityId == nil {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *SAMLConfigurationAttributes) HasEntityId() bool {
	return o != nil && o.EntityId != nil
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *SAMLConfigurationAttributes) SetEntityId(v string) {
	o.EntityId = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SAMLConfigurationAttributes) GetExpiresAt() time.Time {
	if o == nil || o.ExpiresAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SAMLConfigurationAttributes) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *SAMLConfigurationAttributes) HasExpiresAt() bool {
	return o != nil && o.ExpiresAt.IsSet()
}

// SetExpiresAt gets a reference to the given datadog.NullableTime and assigns it to the ExpiresAt field.
func (o *SAMLConfigurationAttributes) SetExpiresAt(v time.Time) {
	o.ExpiresAt.Set(&v)
}

// SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil.
func (o *SAMLConfigurationAttributes) SetExpiresAtNil() {
	o.ExpiresAt.Set(nil)
}

// UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil.
func (o *SAMLConfigurationAttributes) UnsetExpiresAt() {
	o.ExpiresAt.Unset()
}

// GetIdpInitiated returns the IdpInitiated field value if set, zero value otherwise.
func (o *SAMLConfigurationAttributes) GetIdpInitiated() bool {
	if o == nil || o.IdpInitiated == nil {
		var ret bool
		return ret
	}
	return *o.IdpInitiated
}

// GetIdpInitiatedOk returns a tuple with the IdpInitiated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLConfigurationAttributes) GetIdpInitiatedOk() (*bool, bool) {
	if o == nil || o.IdpInitiated == nil {
		return nil, false
	}
	return o.IdpInitiated, true
}

// HasIdpInitiated returns a boolean if a field has been set.
func (o *SAMLConfigurationAttributes) HasIdpInitiated() bool {
	return o != nil && o.IdpInitiated != nil
}

// SetIdpInitiated gets a reference to the given bool and assigns it to the IdpInitiated field.
func (o *SAMLConfigurationAttributes) SetIdpInitiated(v bool) {
	o.IdpInitiated = &v
}

// GetJitDomains returns the JitDomains field value if set, zero value otherwise.
func (o *SAMLConfigurationAttributes) GetJitDomains() []string {
	if o == nil || o.JitDomains == nil {
		var ret []string
		return ret
	}
	return o.JitDomains
}

// GetJitDomainsOk returns a tuple with the JitDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLConfigurationAttributes) GetJitDomainsOk() (*[]string, bool) {
	if o == nil || o.JitDomains == nil {
		return nil, false
	}
	return &o.JitDomains, true
}

// HasJitDomains returns a boolean if a field has been set.
func (o *SAMLConfigurationAttributes) HasJitDomains() bool {
	return o != nil && o.JitDomains != nil
}

// SetJitDomains gets a reference to the given []string and assigns it to the JitDomains field.
func (o *SAMLConfigurationAttributes) SetJitDomains(v []string) {
	o.JitDomains = v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *SAMLConfigurationAttributes) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLConfigurationAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *SAMLConfigurationAttributes) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt != nil
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *SAMLConfigurationAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetSsoUrl returns the SsoUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SAMLConfigurationAttributes) GetSsoUrl() string {
	if o == nil || o.SsoUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.SsoUrl.Get()
}

// GetSsoUrlOk returns a tuple with the SsoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SAMLConfigurationAttributes) GetSsoUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SsoUrl.Get(), o.SsoUrl.IsSet()
}

// HasSsoUrl returns a boolean if a field has been set.
func (o *SAMLConfigurationAttributes) HasSsoUrl() bool {
	return o != nil && o.SsoUrl.IsSet()
}

// SetSsoUrl gets a reference to the given datadog.NullableString and assigns it to the SsoUrl field.
func (o *SAMLConfigurationAttributes) SetSsoUrl(v string) {
	o.SsoUrl.Set(&v)
}

// SetSsoUrlNil sets the value for SsoUrl to be an explicit nil.
func (o *SAMLConfigurationAttributes) SetSsoUrlNil() {
	o.SsoUrl.Set(nil)
}

// UnsetSsoUrl ensures that no value is present for SsoUrl, not even an explicit nil.
func (o *SAMLConfigurationAttributes) UnsetSsoUrl() {
	o.SsoUrl.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o SAMLConfigurationAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AssertionConsumerService != nil {
		toSerialize["assertion_consumer_service"] = o.AssertionConsumerService
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.EntityId != nil {
		toSerialize["entity_id"] = o.EntityId
	}
	if o.ExpiresAt.IsSet() {
		toSerialize["expires_at"] = o.ExpiresAt.Get()
	}
	if o.IdpInitiated != nil {
		toSerialize["idp_initiated"] = o.IdpInitiated
	}
	if o.JitDomains != nil {
		toSerialize["jit_domains"] = o.JitDomains
	}
	if o.ModifiedAt != nil {
		if o.ModifiedAt.Nanosecond() == 0 {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.SsoUrl.IsSet() {
		toSerialize["sso_url"] = o.SsoUrl.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SAMLConfigurationAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AssertionConsumerService []string               `json:"assertion_consumer_service,omitempty"`
		CreatedAt                *time.Time             `json:"created_at,omitempty"`
		EntityId                 *string                `json:"entity_id,omitempty"`
		ExpiresAt                datadog.NullableTime   `json:"expires_at,omitempty"`
		IdpInitiated             *bool                  `json:"idp_initiated,omitempty"`
		JitDomains               []string               `json:"jit_domains,omitempty"`
		ModifiedAt               *time.Time             `json:"modified_at,omitempty"`
		SsoUrl                   datadog.NullableString `json:"sso_url,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assertion_consumer_service", "created_at", "entity_id", "expires_at", "idp_initiated", "jit_domains", "modified_at", "sso_url"})
	} else {
		return err
	}
	o.AssertionConsumerService = all.AssertionConsumerService
	o.CreatedAt = all.CreatedAt
	o.EntityId = all.EntityId
	o.ExpiresAt = all.ExpiresAt
	o.IdpInitiated = all.IdpInitiated
	o.JitDomains = all.JitDomains
	o.ModifiedAt = all.ModifiedAt
	o.SsoUrl = all.SsoUrl

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
