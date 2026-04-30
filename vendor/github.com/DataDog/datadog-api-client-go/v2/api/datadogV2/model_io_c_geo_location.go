// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IoCGeoLocation Geographic location information for an IP indicator.
type IoCGeoLocation struct {
	// City name.
	City *string `json:"city,omitempty"`
	// ISO country code.
	CountryCode *string `json:"country_code,omitempty"`
	// Full country name.
	CountryName *string `json:"country_name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIoCGeoLocation instantiates a new IoCGeoLocation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIoCGeoLocation() *IoCGeoLocation {
	this := IoCGeoLocation{}
	return &this
}

// NewIoCGeoLocationWithDefaults instantiates a new IoCGeoLocation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIoCGeoLocationWithDefaults() *IoCGeoLocation {
	this := IoCGeoLocation{}
	return &this
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *IoCGeoLocation) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCGeoLocation) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *IoCGeoLocation) HasCity() bool {
	return o != nil && o.City != nil
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *IoCGeoLocation) SetCity(v string) {
	o.City = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *IoCGeoLocation) GetCountryCode() string {
	if o == nil || o.CountryCode == nil {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCGeoLocation) GetCountryCodeOk() (*string, bool) {
	if o == nil || o.CountryCode == nil {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *IoCGeoLocation) HasCountryCode() bool {
	return o != nil && o.CountryCode != nil
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *IoCGeoLocation) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetCountryName returns the CountryName field value if set, zero value otherwise.
func (o *IoCGeoLocation) GetCountryName() string {
	if o == nil || o.CountryName == nil {
		var ret string
		return ret
	}
	return *o.CountryName
}

// GetCountryNameOk returns a tuple with the CountryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCGeoLocation) GetCountryNameOk() (*string, bool) {
	if o == nil || o.CountryName == nil {
		return nil, false
	}
	return o.CountryName, true
}

// HasCountryName returns a boolean if a field has been set.
func (o *IoCGeoLocation) HasCountryName() bool {
	return o != nil && o.CountryName != nil
}

// SetCountryName gets a reference to the given string and assigns it to the CountryName field.
func (o *IoCGeoLocation) SetCountryName(v string) {
	o.CountryName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IoCGeoLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.CountryCode != nil {
		toSerialize["country_code"] = o.CountryCode
	}
	if o.CountryName != nil {
		toSerialize["country_name"] = o.CountryName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IoCGeoLocation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		City        *string `json:"city,omitempty"`
		CountryCode *string `json:"country_code,omitempty"`
		CountryName *string `json:"country_name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"city", "country_code", "country_name"})
	} else {
		return err
	}
	o.City = all.City
	o.CountryCode = all.CountryCode
	o.CountryName = all.CountryName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
