// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ResolveVulnerableSymbolsRequestDataAttributes
type ResolveVulnerableSymbolsRequestDataAttributes struct {
	//
	Purls []string `json:"purls,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewResolveVulnerableSymbolsRequestDataAttributes instantiates a new ResolveVulnerableSymbolsRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewResolveVulnerableSymbolsRequestDataAttributes() *ResolveVulnerableSymbolsRequestDataAttributes {
	this := ResolveVulnerableSymbolsRequestDataAttributes{}
	return &this
}

// NewResolveVulnerableSymbolsRequestDataAttributesWithDefaults instantiates a new ResolveVulnerableSymbolsRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewResolveVulnerableSymbolsRequestDataAttributesWithDefaults() *ResolveVulnerableSymbolsRequestDataAttributes {
	this := ResolveVulnerableSymbolsRequestDataAttributes{}
	return &this
}

// GetPurls returns the Purls field value if set, zero value otherwise.
func (o *ResolveVulnerableSymbolsRequestDataAttributes) GetPurls() []string {
	if o == nil || o.Purls == nil {
		var ret []string
		return ret
	}
	return o.Purls
}

// GetPurlsOk returns a tuple with the Purls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResolveVulnerableSymbolsRequestDataAttributes) GetPurlsOk() (*[]string, bool) {
	if o == nil || o.Purls == nil {
		return nil, false
	}
	return &o.Purls, true
}

// HasPurls returns a boolean if a field has been set.
func (o *ResolveVulnerableSymbolsRequestDataAttributes) HasPurls() bool {
	return o != nil && o.Purls != nil
}

// SetPurls gets a reference to the given []string and assigns it to the Purls field.
func (o *ResolveVulnerableSymbolsRequestDataAttributes) SetPurls(v []string) {
	o.Purls = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ResolveVulnerableSymbolsRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Purls != nil {
		toSerialize["purls"] = o.Purls
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ResolveVulnerableSymbolsRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Purls []string `json:"purls,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"purls"})
	} else {
		return err
	}
	o.Purls = all.Purls

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
