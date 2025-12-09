// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PatchTableRequestDataAttributesFileMetadataLocalFile Local file metadata for patch requests using upload ID.
type PatchTableRequestDataAttributesFileMetadataLocalFile struct {
	// The upload ID.
	UploadId string `json:"upload_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewPatchTableRequestDataAttributesFileMetadataLocalFile instantiates a new PatchTableRequestDataAttributesFileMetadataLocalFile object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPatchTableRequestDataAttributesFileMetadataLocalFile(uploadId string) *PatchTableRequestDataAttributesFileMetadataLocalFile {
	this := PatchTableRequestDataAttributesFileMetadataLocalFile{}
	this.UploadId = uploadId
	return &this
}

// NewPatchTableRequestDataAttributesFileMetadataLocalFileWithDefaults instantiates a new PatchTableRequestDataAttributesFileMetadataLocalFile object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPatchTableRequestDataAttributesFileMetadataLocalFileWithDefaults() *PatchTableRequestDataAttributesFileMetadataLocalFile {
	this := PatchTableRequestDataAttributesFileMetadataLocalFile{}
	return &this
}

// GetUploadId returns the UploadId field value.
func (o *PatchTableRequestDataAttributesFileMetadataLocalFile) GetUploadId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UploadId
}

// GetUploadIdOk returns a tuple with the UploadId field value
// and a boolean to check if the value has been set.
func (o *PatchTableRequestDataAttributesFileMetadataLocalFile) GetUploadIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UploadId, true
}

// SetUploadId sets field value.
func (o *PatchTableRequestDataAttributesFileMetadataLocalFile) SetUploadId(v string) {
	o.UploadId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PatchTableRequestDataAttributesFileMetadataLocalFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["upload_id"] = o.UploadId
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PatchTableRequestDataAttributesFileMetadataLocalFile) UnmarshalJSON(bytes []byte) (err error) {
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
