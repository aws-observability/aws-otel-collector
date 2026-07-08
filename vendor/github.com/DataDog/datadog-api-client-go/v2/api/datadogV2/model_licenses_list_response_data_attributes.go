// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LicensesListResponseDataAttributes The attributes of the licenses list response, containing the array of SPDX licenses.
type LicensesListResponseDataAttributes struct {
	// The list of SPDX licenses returned by the API.
	Licenses []LicensesListResponseDataAttributesLicensesItems `json:"licenses"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLicensesListResponseDataAttributes instantiates a new LicensesListResponseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLicensesListResponseDataAttributes(licenses []LicensesListResponseDataAttributesLicensesItems) *LicensesListResponseDataAttributes {
	this := LicensesListResponseDataAttributes{}
	this.Licenses = licenses
	return &this
}

// NewLicensesListResponseDataAttributesWithDefaults instantiates a new LicensesListResponseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLicensesListResponseDataAttributesWithDefaults() *LicensesListResponseDataAttributes {
	this := LicensesListResponseDataAttributes{}
	return &this
}

// GetLicenses returns the Licenses field value.
func (o *LicensesListResponseDataAttributes) GetLicenses() []LicensesListResponseDataAttributesLicensesItems {
	if o == nil {
		var ret []LicensesListResponseDataAttributesLicensesItems
		return ret
	}
	return o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value
// and a boolean to check if the value has been set.
func (o *LicensesListResponseDataAttributes) GetLicensesOk() (*[]LicensesListResponseDataAttributesLicensesItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Licenses, true
}

// SetLicenses sets field value.
func (o *LicensesListResponseDataAttributes) SetLicenses(v []LicensesListResponseDataAttributesLicensesItems) {
	o.Licenses = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LicensesListResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["licenses"] = o.Licenses

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LicensesListResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Licenses *[]LicensesListResponseDataAttributesLicensesItems `json:"licenses"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Licenses == nil {
		return fmt.Errorf("required field licenses missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"licenses"})
	} else {
		return err
	}
	o.Licenses = *all.Licenses

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
