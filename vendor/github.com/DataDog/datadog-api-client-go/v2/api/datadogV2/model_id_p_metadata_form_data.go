// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	_io "io"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IdPMetadataFormData The form data submitted to upload IdP metadata
type IdPMetadataFormData struct {
	// The IdP metadata XML file
	IdpFile *_io.Reader `json:"idp_file,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIdPMetadataFormData instantiates a new IdPMetadataFormData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIdPMetadataFormData() *IdPMetadataFormData {
	this := IdPMetadataFormData{}
	return &this
}

// NewIdPMetadataFormDataWithDefaults instantiates a new IdPMetadataFormData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIdPMetadataFormDataWithDefaults() *IdPMetadataFormData {
	this := IdPMetadataFormData{}
	return &this
}

// GetIdpFile returns the IdpFile field value if set, zero value otherwise.
func (o *IdPMetadataFormData) GetIdpFile() _io.Reader {
	if o == nil || o.IdpFile == nil {
		var ret _io.Reader
		return ret
	}
	return *o.IdpFile
}

// GetIdpFileOk returns a tuple with the IdpFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdPMetadataFormData) GetIdpFileOk() (*_io.Reader, bool) {
	if o == nil || o.IdpFile == nil {
		return nil, false
	}
	return o.IdpFile, true
}

// HasIdpFile returns a boolean if a field has been set.
func (o *IdPMetadataFormData) HasIdpFile() bool {
	return o != nil && o.IdpFile != nil
}

// SetIdpFile gets a reference to the given _io.Reader and assigns it to the IdpFile field.
func (o *IdPMetadataFormData) SetIdpFile(v _io.Reader) {
	o.IdpFile = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IdPMetadataFormData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.IdpFile != nil {
		toSerialize["idp_file"] = o.IdpFile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IdPMetadataFormData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IdpFile *_io.Reader `json:"idp_file,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"idp_file"})
	} else {
		return err
	}
	o.IdpFile = all.IdpFile

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
