// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SAMLConfigurationUpdateAttributes Attributes for updating a SAML configuration.
type SAMLConfigurationUpdateAttributes struct {
	// Whether identity-provider-initiated login is enabled for the organization.
	IdpInitiated *bool `json:"idp_initiated,omitempty"`
	// Email domains for which users are automatically provisioned on first SAML login
	// (just-in-time provisioning). A default role is required to enable just-in-time provisioning.
	JitDomains []string `json:"jit_domains,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSAMLConfigurationUpdateAttributes instantiates a new SAMLConfigurationUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSAMLConfigurationUpdateAttributes() *SAMLConfigurationUpdateAttributes {
	this := SAMLConfigurationUpdateAttributes{}
	return &this
}

// NewSAMLConfigurationUpdateAttributesWithDefaults instantiates a new SAMLConfigurationUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSAMLConfigurationUpdateAttributesWithDefaults() *SAMLConfigurationUpdateAttributes {
	this := SAMLConfigurationUpdateAttributes{}
	return &this
}

// GetIdpInitiated returns the IdpInitiated field value if set, zero value otherwise.
func (o *SAMLConfigurationUpdateAttributes) GetIdpInitiated() bool {
	if o == nil || o.IdpInitiated == nil {
		var ret bool
		return ret
	}
	return *o.IdpInitiated
}

// GetIdpInitiatedOk returns a tuple with the IdpInitiated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLConfigurationUpdateAttributes) GetIdpInitiatedOk() (*bool, bool) {
	if o == nil || o.IdpInitiated == nil {
		return nil, false
	}
	return o.IdpInitiated, true
}

// HasIdpInitiated returns a boolean if a field has been set.
func (o *SAMLConfigurationUpdateAttributes) HasIdpInitiated() bool {
	return o != nil && o.IdpInitiated != nil
}

// SetIdpInitiated gets a reference to the given bool and assigns it to the IdpInitiated field.
func (o *SAMLConfigurationUpdateAttributes) SetIdpInitiated(v bool) {
	o.IdpInitiated = &v
}

// GetJitDomains returns the JitDomains field value if set, zero value otherwise.
func (o *SAMLConfigurationUpdateAttributes) GetJitDomains() []string {
	if o == nil || o.JitDomains == nil {
		var ret []string
		return ret
	}
	return o.JitDomains
}

// GetJitDomainsOk returns a tuple with the JitDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLConfigurationUpdateAttributes) GetJitDomainsOk() (*[]string, bool) {
	if o == nil || o.JitDomains == nil {
		return nil, false
	}
	return &o.JitDomains, true
}

// HasJitDomains returns a boolean if a field has been set.
func (o *SAMLConfigurationUpdateAttributes) HasJitDomains() bool {
	return o != nil && o.JitDomains != nil
}

// SetJitDomains gets a reference to the given []string and assigns it to the JitDomains field.
func (o *SAMLConfigurationUpdateAttributes) SetJitDomains(v []string) {
	o.JitDomains = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SAMLConfigurationUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.IdpInitiated != nil {
		toSerialize["idp_initiated"] = o.IdpInitiated
	}
	if o.JitDomains != nil {
		toSerialize["jit_domains"] = o.JitDomains
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SAMLConfigurationUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IdpInitiated *bool    `json:"idp_initiated,omitempty"`
		JitDomains   []string `json:"jit_domains,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"idp_initiated", "jit_domains"})
	} else {
		return err
	}
	o.IdpInitiated = all.IdpInitiated
	o.JitDomains = all.JitDomains

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
