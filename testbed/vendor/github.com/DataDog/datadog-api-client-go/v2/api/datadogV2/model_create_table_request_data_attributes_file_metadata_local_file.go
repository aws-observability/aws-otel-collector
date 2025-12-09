// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateTableRequestDataAttributesFileMetadataLocalFile Local file metadata for create requests using the upload ID.
type CreateTableRequestDataAttributesFileMetadataLocalFile struct {
	// The upload ID.
	UploadId string `json:"upload_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewCreateTableRequestDataAttributesFileMetadataLocalFile instantiates a new CreateTableRequestDataAttributesFileMetadataLocalFile object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateTableRequestDataAttributesFileMetadataLocalFile(uploadId string) *CreateTableRequestDataAttributesFileMetadataLocalFile {
	this := CreateTableRequestDataAttributesFileMetadataLocalFile{}
	this.UploadId = uploadId
	return &this
}

// NewCreateTableRequestDataAttributesFileMetadataLocalFileWithDefaults instantiates a new CreateTableRequestDataAttributesFileMetadataLocalFile object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateTableRequestDataAttributesFileMetadataLocalFileWithDefaults() *CreateTableRequestDataAttributesFileMetadataLocalFile {
	this := CreateTableRequestDataAttributesFileMetadataLocalFile{}
	return &this
}

// GetUploadId returns the UploadId field value.
func (o *CreateTableRequestDataAttributesFileMetadataLocalFile) GetUploadId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UploadId
}

// GetUploadIdOk returns a tuple with the UploadId field value
// and a boolean to check if the value has been set.
func (o *CreateTableRequestDataAttributesFileMetadataLocalFile) GetUploadIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UploadId, true
}

// SetUploadId sets field value.
func (o *CreateTableRequestDataAttributesFileMetadataLocalFile) SetUploadId(v string) {
	o.UploadId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateTableRequestDataAttributesFileMetadataLocalFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["upload_id"] = o.UploadId
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateTableRequestDataAttributesFileMetadataLocalFile) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		UploadId *string `json:"upload_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.UploadId == nil {
		return fmt.Errorf("required field upload_id missing")
	}
	o.UploadId = *all.UploadId

	return nil
}
