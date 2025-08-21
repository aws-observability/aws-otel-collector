// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SBOMComponentLicense The software license of the component of the SBOM.
type SBOMComponentLicense struct {
	// The software license of the component of the SBOM.
	License SBOMComponentLicenseLicense `json:"license"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSBOMComponentLicense instantiates a new SBOMComponentLicense object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSBOMComponentLicense(license SBOMComponentLicenseLicense) *SBOMComponentLicense {
	this := SBOMComponentLicense{}
	this.License = license
	return &this
}

// NewSBOMComponentLicenseWithDefaults instantiates a new SBOMComponentLicense object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSBOMComponentLicenseWithDefaults() *SBOMComponentLicense {
	this := SBOMComponentLicense{}
	return &this
}

// GetLicense returns the License field value.
func (o *SBOMComponentLicense) GetLicense() SBOMComponentLicenseLicense {
	if o == nil {
		var ret SBOMComponentLicenseLicense
		return ret
	}
	return o.License
}

// GetLicenseOk returns a tuple with the License field value
// and a boolean to check if the value has been set.
func (o *SBOMComponentLicense) GetLicenseOk() (*SBOMComponentLicenseLicense, bool) {
	if o == nil {
		return nil, false
	}
	return &o.License, true
}

// SetLicense sets field value.
func (o *SBOMComponentLicense) SetLicense(v SBOMComponentLicenseLicense) {
	o.License = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SBOMComponentLicense) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["license"] = o.License

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SBOMComponentLicense) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		License *SBOMComponentLicenseLicense `json:"license"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.License == nil {
		return fmt.Errorf("required field license missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"license"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.License.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.License = *all.License

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
