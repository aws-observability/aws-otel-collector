// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestFileCompleteMultipartUploadPart A completed part of a multipart upload.
type SyntheticsTestFileCompleteMultipartUploadPart struct {
	// The ETag returned by the storage provider after uploading the part.
	ETag string `json:"ETag"`
	// The 1-indexed part number for the multipart upload.
	PartNumber int64 `json:"PartNumber"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestFileCompleteMultipartUploadPart instantiates a new SyntheticsTestFileCompleteMultipartUploadPart object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestFileCompleteMultipartUploadPart(eTag string, partNumber int64) *SyntheticsTestFileCompleteMultipartUploadPart {
	this := SyntheticsTestFileCompleteMultipartUploadPart{}
	this.ETag = eTag
	this.PartNumber = partNumber
	return &this
}

// NewSyntheticsTestFileCompleteMultipartUploadPartWithDefaults instantiates a new SyntheticsTestFileCompleteMultipartUploadPart object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestFileCompleteMultipartUploadPartWithDefaults() *SyntheticsTestFileCompleteMultipartUploadPart {
	this := SyntheticsTestFileCompleteMultipartUploadPart{}
	return &this
}

// GetETag returns the ETag field value.
func (o *SyntheticsTestFileCompleteMultipartUploadPart) GetETag() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ETag
}

// GetETagOk returns a tuple with the ETag field value
// and a boolean to check if the value has been set.
func (o *SyntheticsTestFileCompleteMultipartUploadPart) GetETagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ETag, true
}

// SetETag sets field value.
func (o *SyntheticsTestFileCompleteMultipartUploadPart) SetETag(v string) {
	o.ETag = v
}

// GetPartNumber returns the PartNumber field value.
func (o *SyntheticsTestFileCompleteMultipartUploadPart) GetPartNumber() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value
// and a boolean to check if the value has been set.
func (o *SyntheticsTestFileCompleteMultipartUploadPart) GetPartNumberOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartNumber, true
}

// SetPartNumber sets field value.
func (o *SyntheticsTestFileCompleteMultipartUploadPart) SetPartNumber(v int64) {
	o.PartNumber = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestFileCompleteMultipartUploadPart) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["ETag"] = o.ETag
	toSerialize["PartNumber"] = o.PartNumber

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestFileCompleteMultipartUploadPart) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ETag       *string `json:"ETag"`
		PartNumber *int64  `json:"PartNumber"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ETag == nil {
		return fmt.Errorf("required field ETag missing")
	}
	if all.PartNumber == nil {
		return fmt.Errorf("required field PartNumber missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"ETag", "PartNumber"})
	} else {
		return err
	}
	o.ETag = *all.ETag
	o.PartNumber = *all.PartNumber

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
