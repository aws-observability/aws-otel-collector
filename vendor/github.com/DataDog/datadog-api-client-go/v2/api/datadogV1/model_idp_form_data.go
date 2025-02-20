// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"
	_io "io"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IdpFormData Object describing the IdP configuration.
type IdpFormData struct {
	// The path to the XML metadata file you wish to upload.
	IdpFile _io.Reader `json:"idp_file"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIdpFormData instantiates a new IdpFormData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIdpFormData(idpFile _io.Reader) *IdpFormData {
	this := IdpFormData{}
	this.IdpFile = idpFile
	return &this
}

// NewIdpFormDataWithDefaults instantiates a new IdpFormData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIdpFormDataWithDefaults() *IdpFormData {
	this := IdpFormData{}
	return &this
}

// GetIdpFile returns the IdpFile field value.
func (o *IdpFormData) GetIdpFile() _io.Reader {
	if o == nil {
		var ret _io.Reader
		return ret
	}
	return o.IdpFile
}

// GetIdpFileOk returns a tuple with the IdpFile field value
// and a boolean to check if the value has been set.
func (o *IdpFormData) GetIdpFileOk() (*_io.Reader, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdpFile, true
}

// SetIdpFile sets field value.
func (o *IdpFormData) SetIdpFile(v _io.Reader) {
	o.IdpFile = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IdpFormData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["idp_file"] = o.IdpFile

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IdpFormData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IdpFile *_io.Reader `json:"idp_file"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.IdpFile == nil {
		return fmt.Errorf("required field idp_file missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"idp_file"})
	} else {
		return err
	}
	o.IdpFile = *all.IdpFile

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
