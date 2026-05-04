// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestFileAbortMultipartUploadRequest Request body for aborting a multipart file upload.
type SyntheticsTestFileAbortMultipartUploadRequest struct {
	// The full storage path of the file whose upload should be aborted.
	Key string `json:"key"`
	// The upload ID of the multipart upload to abort.
	UploadId string `json:"uploadId"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestFileAbortMultipartUploadRequest instantiates a new SyntheticsTestFileAbortMultipartUploadRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestFileAbortMultipartUploadRequest(key string, uploadId string) *SyntheticsTestFileAbortMultipartUploadRequest {
	this := SyntheticsTestFileAbortMultipartUploadRequest{}
	this.Key = key
	this.UploadId = uploadId
	return &this
}

// NewSyntheticsTestFileAbortMultipartUploadRequestWithDefaults instantiates a new SyntheticsTestFileAbortMultipartUploadRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestFileAbortMultipartUploadRequestWithDefaults() *SyntheticsTestFileAbortMultipartUploadRequest {
	this := SyntheticsTestFileAbortMultipartUploadRequest{}
	return &this
}

// GetKey returns the Key field value.
func (o *SyntheticsTestFileAbortMultipartUploadRequest) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *SyntheticsTestFileAbortMultipartUploadRequest) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value.
func (o *SyntheticsTestFileAbortMultipartUploadRequest) SetKey(v string) {
	o.Key = v
}

// GetUploadId returns the UploadId field value.
func (o *SyntheticsTestFileAbortMultipartUploadRequest) GetUploadId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UploadId
}

// GetUploadIdOk returns a tuple with the UploadId field value
// and a boolean to check if the value has been set.
func (o *SyntheticsTestFileAbortMultipartUploadRequest) GetUploadIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UploadId, true
}

// SetUploadId sets field value.
func (o *SyntheticsTestFileAbortMultipartUploadRequest) SetUploadId(v string) {
	o.UploadId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestFileAbortMultipartUploadRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["key"] = o.Key
	toSerialize["uploadId"] = o.UploadId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestFileAbortMultipartUploadRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Key      *string `json:"key"`
		UploadId *string `json:"uploadId"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Key == nil {
		return fmt.Errorf("required field key missing")
	}
	if all.UploadId == nil {
		return fmt.Errorf("required field uploadId missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"key", "uploadId"})
	} else {
		return err
	}
	o.Key = *all.Key
	o.UploadId = *all.UploadId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
