// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OAuth2WellKnownSitesAttributes Attributes containing the list of public OAuth2 sites.
type OAuth2WellKnownSitesAttributes struct {
	// Array of public OAuth2 site URLs for the environment.
	Sites []string `json:"sites"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOAuth2WellKnownSitesAttributes instantiates a new OAuth2WellKnownSitesAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOAuth2WellKnownSitesAttributes(sites []string) *OAuth2WellKnownSitesAttributes {
	this := OAuth2WellKnownSitesAttributes{}
	this.Sites = sites
	return &this
}

// NewOAuth2WellKnownSitesAttributesWithDefaults instantiates a new OAuth2WellKnownSitesAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOAuth2WellKnownSitesAttributesWithDefaults() *OAuth2WellKnownSitesAttributes {
	this := OAuth2WellKnownSitesAttributes{}
	return &this
}

// GetSites returns the Sites field value.
func (o *OAuth2WellKnownSitesAttributes) GetSites() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Sites
}

// GetSitesOk returns a tuple with the Sites field value
// and a boolean to check if the value has been set.
func (o *OAuth2WellKnownSitesAttributes) GetSitesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sites, true
}

// SetSites sets field value.
func (o *OAuth2WellKnownSitesAttributes) SetSites(v []string) {
	o.Sites = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OAuth2WellKnownSitesAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["sites"] = o.Sites

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OAuth2WellKnownSitesAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Sites *[]string `json:"sites"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Sites == nil {
		return fmt.Errorf("required field sites missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"sites"})
	} else {
		return err
	}
	o.Sites = *all.Sites

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
