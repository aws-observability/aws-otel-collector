// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestFileCompleteMultipartUploadRequest Request body for completing a multipart file upload.
type SyntheticsTestFileCompleteMultipartUploadRequest struct {
	// The full storage path for the uploaded file.
	Key string `json:"key"`
	// Array of completed parts with their ETags.
	Parts []SyntheticsTestFileCompleteMultipartUploadPart `json:"parts"`
	// The upload ID returned when the multipart upload was initiated.
	UploadId string `json:"uploadId"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestFileCompleteMultipartUploadRequest instantiates a new SyntheticsTestFileCompleteMultipartUploadRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestFileCompleteMultipartUploadRequest(key string, parts []SyntheticsTestFileCompleteMultipartUploadPart, uploadId string) *SyntheticsTestFileCompleteMultipartUploadRequest {
	this := SyntheticsTestFileCompleteMultipartUploadRequest{}
	this.Key = key
	this.Parts = parts
	this.UploadId = uploadId
	return &this
}

// NewSyntheticsTestFileCompleteMultipartUploadRequestWithDefaults instantiates a new SyntheticsTestFileCompleteMultipartUploadRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestFileCompleteMultipartUploadRequestWithDefaults() *SyntheticsTestFileCompleteMultipartUploadRequest {
	this := SyntheticsTestFileCompleteMultipartUploadRequest{}
	return &this
}

// GetKey returns the Key field value.
func (o *SyntheticsTestFileCompleteMultipartUploadRequest) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *SyntheticsTestFileCompleteMultipartUploadRequest) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value.
func (o *SyntheticsTestFileCompleteMultipartUploadRequest) SetKey(v string) {
	o.Key = v
}

// GetParts returns the Parts field value.
func (o *SyntheticsTestFileCompleteMultipartUploadRequest) GetParts() []SyntheticsTestFileCompleteMultipartUploadPart {
	if o == nil {
		var ret []SyntheticsTestFileCompleteMultipartUploadPart
		return ret
	}
	return o.Parts
}

// GetPartsOk returns a tuple with the Parts field value
// and a boolean to check if the value has been set.
func (o *SyntheticsTestFileCompleteMultipartUploadRequest) GetPartsOk() (*[]SyntheticsTestFileCompleteMultipartUploadPart, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parts, true
}

// SetParts sets field value.
func (o *SyntheticsTestFileCompleteMultipartUploadRequest) SetParts(v []SyntheticsTestFileCompleteMultipartUploadPart) {
	o.Parts = v
}

// GetUploadId returns the UploadId field value.
func (o *SyntheticsTestFileCompleteMultipartUploadRequest) GetUploadId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UploadId
}

// GetUploadIdOk returns a tuple with the UploadId field value
// and a boolean to check if the value has been set.
func (o *SyntheticsTestFileCompleteMultipartUploadRequest) GetUploadIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UploadId, true
}

// SetUploadId sets field value.
func (o *SyntheticsTestFileCompleteMultipartUploadRequest) SetUploadId(v string) {
	o.UploadId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestFileCompleteMultipartUploadRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["key"] = o.Key
	toSerialize["parts"] = o.Parts
	toSerialize["uploadId"] = o.UploadId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestFileCompleteMultipartUploadRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Key      *string                                          `json:"key"`
		Parts    *[]SyntheticsTestFileCompleteMultipartUploadPart `json:"parts"`
		UploadId *string                                          `json:"uploadId"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Key == nil {
		return fmt.Errorf("required field key missing")
	}
	if all.Parts == nil {
		return fmt.Errorf("required field parts missing")
	}
	if all.UploadId == nil {
		return fmt.Errorf("required field uploadId missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"key", "parts", "uploadId"})
	} else {
		return err
	}
	o.Key = *all.Key
	o.Parts = *all.Parts
	o.UploadId = *all.UploadId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
