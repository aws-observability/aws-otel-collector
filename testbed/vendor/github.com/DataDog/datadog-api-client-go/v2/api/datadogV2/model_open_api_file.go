// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	_io "io"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAPIFile Object for API data in an `OpenAPI` format as a file.
type OpenAPIFile struct {
	// Binary `OpenAPI` spec file
	OpenapiSpecFile *_io.Reader `json:"openapi_spec_file,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOpenAPIFile instantiates a new OpenAPIFile object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpenAPIFile() *OpenAPIFile {
	this := OpenAPIFile{}
	return &this
}

// NewOpenAPIFileWithDefaults instantiates a new OpenAPIFile object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpenAPIFileWithDefaults() *OpenAPIFile {
	this := OpenAPIFile{}
	return &this
}

// GetOpenapiSpecFile returns the OpenapiSpecFile field value if set, zero value otherwise.
func (o *OpenAPIFile) GetOpenapiSpecFile() _io.Reader {
	if o == nil || o.OpenapiSpecFile == nil {
		var ret _io.Reader
		return ret
	}
	return *o.OpenapiSpecFile
}

// GetOpenapiSpecFileOk returns a tuple with the OpenapiSpecFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAPIFile) GetOpenapiSpecFileOk() (*_io.Reader, bool) {
	if o == nil || o.OpenapiSpecFile == nil {
		return nil, false
	}
	return o.OpenapiSpecFile, true
}

// HasOpenapiSpecFile returns a boolean if a field has been set.
func (o *OpenAPIFile) HasOpenapiSpecFile() bool {
	return o != nil && o.OpenapiSpecFile != nil
}

// SetOpenapiSpecFile gets a reference to the given _io.Reader and assigns it to the OpenapiSpecFile field.
func (o *OpenAPIFile) SetOpenapiSpecFile(v _io.Reader) {
	o.OpenapiSpecFile = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpenAPIFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.OpenapiSpecFile != nil {
		toSerialize["openapi_spec_file"] = o.OpenapiSpecFile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpenAPIFile) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OpenapiSpecFile *_io.Reader `json:"openapi_spec_file,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"openapi_spec_file"})
	} else {
		return err
	}
	o.OpenapiSpecFile = all.OpenapiSpecFile

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
