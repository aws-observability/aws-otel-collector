// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SBOMComponentLicenseLicense The software license of the component of the SBOM.
type SBOMComponentLicenseLicense struct {
	// The name of the software license of the component of the SBOM.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSBOMComponentLicenseLicense instantiates a new SBOMComponentLicenseLicense object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSBOMComponentLicenseLicense(name string) *SBOMComponentLicenseLicense {
	this := SBOMComponentLicenseLicense{}
	this.Name = name
	return &this
}

// NewSBOMComponentLicenseLicenseWithDefaults instantiates a new SBOMComponentLicenseLicense object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSBOMComponentLicenseLicenseWithDefaults() *SBOMComponentLicenseLicense {
	this := SBOMComponentLicenseLicense{}
	return &this
}

// GetName returns the Name field value.
func (o *SBOMComponentLicenseLicense) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SBOMComponentLicenseLicense) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SBOMComponentLicenseLicense) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SBOMComponentLicenseLicense) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SBOMComponentLicenseLicense) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name *string `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name"})
	} else {
		return err
	}
	o.Name = *all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
