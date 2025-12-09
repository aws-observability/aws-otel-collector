// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DomainAllowlistResponseDataAttributes The details of the email domain allowlist.
type DomainAllowlistResponseDataAttributes struct {
	// The list of domains in the email domain allowlist.
	Domains []string `json:"domains,omitempty"`
	// Whether the email domain allowlist is enabled for the org.
	Enabled *bool `json:"enabled,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDomainAllowlistResponseDataAttributes instantiates a new DomainAllowlistResponseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDomainAllowlistResponseDataAttributes() *DomainAllowlistResponseDataAttributes {
	this := DomainAllowlistResponseDataAttributes{}
	return &this
}

// NewDomainAllowlistResponseDataAttributesWithDefaults instantiates a new DomainAllowlistResponseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDomainAllowlistResponseDataAttributesWithDefaults() *DomainAllowlistResponseDataAttributes {
	this := DomainAllowlistResponseDataAttributes{}
	return &this
}

// GetDomains returns the Domains field value if set, zero value otherwise.
func (o *DomainAllowlistResponseDataAttributes) GetDomains() []string {
	if o == nil || o.Domains == nil {
		var ret []string
		return ret
	}
	return o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainAllowlistResponseDataAttributes) GetDomainsOk() (*[]string, bool) {
	if o == nil || o.Domains == nil {
		return nil, false
	}
	return &o.Domains, true
}

// HasDomains returns a boolean if a field has been set.
func (o *DomainAllowlistResponseDataAttributes) HasDomains() bool {
	return o != nil && o.Domains != nil
}

// SetDomains gets a reference to the given []string and assigns it to the Domains field.
func (o *DomainAllowlistResponseDataAttributes) SetDomains(v []string) {
	o.Domains = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *DomainAllowlistResponseDataAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainAllowlistResponseDataAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *DomainAllowlistResponseDataAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *DomainAllowlistResponseDataAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DomainAllowlistResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Domains != nil {
		toSerialize["domains"] = o.Domains
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DomainAllowlistResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Domains []string `json:"domains,omitempty"`
		Enabled *bool    `json:"enabled,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"domains", "enabled"})
	} else {
		return err
	}
	o.Domains = all.Domains
	o.Enabled = all.Enabled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
