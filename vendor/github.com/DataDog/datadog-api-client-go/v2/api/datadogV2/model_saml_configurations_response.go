// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SAMLConfigurationsResponse Response containing a list of SAML configurations.
type SAMLConfigurationsResponse struct {
	// Array of SAML configurations. An organization has at most one SAML configuration.
	Data []SAMLConfiguration `json:"data,omitempty"`
	// Resources related to the SAML configurations, such as the default roles.
	Included []Role `json:"included,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSAMLConfigurationsResponse instantiates a new SAMLConfigurationsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSAMLConfigurationsResponse() *SAMLConfigurationsResponse {
	this := SAMLConfigurationsResponse{}
	return &this
}

// NewSAMLConfigurationsResponseWithDefaults instantiates a new SAMLConfigurationsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSAMLConfigurationsResponseWithDefaults() *SAMLConfigurationsResponse {
	this := SAMLConfigurationsResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SAMLConfigurationsResponse) GetData() []SAMLConfiguration {
	if o == nil || o.Data == nil {
		var ret []SAMLConfiguration
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLConfigurationsResponse) GetDataOk() (*[]SAMLConfiguration, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SAMLConfigurationsResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given []SAMLConfiguration and assigns it to the Data field.
func (o *SAMLConfigurationsResponse) SetData(v []SAMLConfiguration) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *SAMLConfigurationsResponse) GetIncluded() []Role {
	if o == nil || o.Included == nil {
		var ret []Role
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLConfigurationsResponse) GetIncludedOk() (*[]Role, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return &o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *SAMLConfigurationsResponse) HasIncluded() bool {
	return o != nil && o.Included != nil
}

// SetIncluded gets a reference to the given []Role and assigns it to the Included field.
func (o *SAMLConfigurationsResponse) SetIncluded(v []Role) {
	o.Included = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SAMLConfigurationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SAMLConfigurationsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data     []SAMLConfiguration `json:"data,omitempty"`
		Included []Role              `json:"included,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "included"})
	} else {
		return err
	}
	o.Data = all.Data
	o.Included = all.Included

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
