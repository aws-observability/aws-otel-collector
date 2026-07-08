// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GoogleChatOrganizationAttributes Google Chat organization attributes.
type GoogleChatOrganizationAttributes struct {
	// The Google Chat organization domain ID.
	DomainId *string `json:"domain_id,omitempty"`
	// The Google Chat organization domain name.
	DomainName *string `json:"domain_name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGoogleChatOrganizationAttributes instantiates a new GoogleChatOrganizationAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGoogleChatOrganizationAttributes() *GoogleChatOrganizationAttributes {
	this := GoogleChatOrganizationAttributes{}
	return &this
}

// NewGoogleChatOrganizationAttributesWithDefaults instantiates a new GoogleChatOrganizationAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGoogleChatOrganizationAttributesWithDefaults() *GoogleChatOrganizationAttributes {
	this := GoogleChatOrganizationAttributes{}
	return &this
}

// GetDomainId returns the DomainId field value if set, zero value otherwise.
func (o *GoogleChatOrganizationAttributes) GetDomainId() string {
	if o == nil || o.DomainId == nil {
		var ret string
		return ret
	}
	return *o.DomainId
}

// GetDomainIdOk returns a tuple with the DomainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleChatOrganizationAttributes) GetDomainIdOk() (*string, bool) {
	if o == nil || o.DomainId == nil {
		return nil, false
	}
	return o.DomainId, true
}

// HasDomainId returns a boolean if a field has been set.
func (o *GoogleChatOrganizationAttributes) HasDomainId() bool {
	return o != nil && o.DomainId != nil
}

// SetDomainId gets a reference to the given string and assigns it to the DomainId field.
func (o *GoogleChatOrganizationAttributes) SetDomainId(v string) {
	o.DomainId = &v
}

// GetDomainName returns the DomainName field value if set, zero value otherwise.
func (o *GoogleChatOrganizationAttributes) GetDomainName() string {
	if o == nil || o.DomainName == nil {
		var ret string
		return ret
	}
	return *o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleChatOrganizationAttributes) GetDomainNameOk() (*string, bool) {
	if o == nil || o.DomainName == nil {
		return nil, false
	}
	return o.DomainName, true
}

// HasDomainName returns a boolean if a field has been set.
func (o *GoogleChatOrganizationAttributes) HasDomainName() bool {
	return o != nil && o.DomainName != nil
}

// SetDomainName gets a reference to the given string and assigns it to the DomainName field.
func (o *GoogleChatOrganizationAttributes) SetDomainName(v string) {
	o.DomainName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GoogleChatOrganizationAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DomainId != nil {
		toSerialize["domain_id"] = o.DomainId
	}
	if o.DomainName != nil {
		toSerialize["domain_name"] = o.DomainName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GoogleChatOrganizationAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DomainId   *string `json:"domain_id,omitempty"`
		DomainName *string `json:"domain_name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"domain_id", "domain_name"})
	} else {
		return err
	}
	o.DomainId = all.DomainId
	o.DomainName = all.DomainName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
