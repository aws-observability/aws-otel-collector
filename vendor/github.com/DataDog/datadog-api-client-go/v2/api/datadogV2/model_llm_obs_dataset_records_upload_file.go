// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	_io "io"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsDatasetRecordsUploadFile Multipart payload for uploading dataset records from a file.
type LLMObsDatasetRecordsUploadFile struct {
	// The records file to upload. Currently only CSV is supported. The file must include an `input` column. Optional columns include `id`, `expected_output`, `metadata`, and `tags`.
	File *_io.Reader `json:"file,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsDatasetRecordsUploadFile instantiates a new LLMObsDatasetRecordsUploadFile object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsDatasetRecordsUploadFile() *LLMObsDatasetRecordsUploadFile {
	this := LLMObsDatasetRecordsUploadFile{}
	return &this
}

// NewLLMObsDatasetRecordsUploadFileWithDefaults instantiates a new LLMObsDatasetRecordsUploadFile object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsDatasetRecordsUploadFileWithDefaults() *LLMObsDatasetRecordsUploadFile {
	this := LLMObsDatasetRecordsUploadFile{}
	return &this
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *LLMObsDatasetRecordsUploadFile) GetFile() _io.Reader {
	if o == nil || o.File == nil {
		var ret _io.Reader
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsDatasetRecordsUploadFile) GetFileOk() (*_io.Reader, bool) {
	if o == nil || o.File == nil {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *LLMObsDatasetRecordsUploadFile) HasFile() bool {
	return o != nil && o.File != nil
}

// SetFile gets a reference to the given _io.Reader and assigns it to the File field.
func (o *LLMObsDatasetRecordsUploadFile) SetFile(v _io.Reader) {
	o.File = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsDatasetRecordsUploadFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.File != nil {
		toSerialize["file"] = o.File
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsDatasetRecordsUploadFile) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		File *_io.Reader `json:"file,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"file"})
	} else {
		return err
	}
	o.File = all.File

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
