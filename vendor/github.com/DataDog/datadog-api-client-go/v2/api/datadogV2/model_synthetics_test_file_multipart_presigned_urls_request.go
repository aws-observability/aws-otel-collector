// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestFileMultipartPresignedUrlsRequest Request body for getting presigned URLs for a multipart file upload.
type SyntheticsTestFileMultipartPresignedUrlsRequest struct {
	// The bucket key prefix indicating the type of file upload.
	BucketKeyPrefix SyntheticsTestFileMultipartPresignedUrlsRequestBucketKeyPrefix `json:"bucketKeyPrefix"`
	// Array of part descriptors for the multipart upload.
	Parts []SyntheticsTestFileMultipartPresignedUrlsPart `json:"parts"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestFileMultipartPresignedUrlsRequest instantiates a new SyntheticsTestFileMultipartPresignedUrlsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestFileMultipartPresignedUrlsRequest(bucketKeyPrefix SyntheticsTestFileMultipartPresignedUrlsRequestBucketKeyPrefix, parts []SyntheticsTestFileMultipartPresignedUrlsPart) *SyntheticsTestFileMultipartPresignedUrlsRequest {
	this := SyntheticsTestFileMultipartPresignedUrlsRequest{}
	this.BucketKeyPrefix = bucketKeyPrefix
	this.Parts = parts
	return &this
}

// NewSyntheticsTestFileMultipartPresignedUrlsRequestWithDefaults instantiates a new SyntheticsTestFileMultipartPresignedUrlsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestFileMultipartPresignedUrlsRequestWithDefaults() *SyntheticsTestFileMultipartPresignedUrlsRequest {
	this := SyntheticsTestFileMultipartPresignedUrlsRequest{}
	return &this
}

// GetBucketKeyPrefix returns the BucketKeyPrefix field value.
func (o *SyntheticsTestFileMultipartPresignedUrlsRequest) GetBucketKeyPrefix() SyntheticsTestFileMultipartPresignedUrlsRequestBucketKeyPrefix {
	if o == nil {
		var ret SyntheticsTestFileMultipartPresignedUrlsRequestBucketKeyPrefix
		return ret
	}
	return o.BucketKeyPrefix
}

// GetBucketKeyPrefixOk returns a tuple with the BucketKeyPrefix field value
// and a boolean to check if the value has been set.
func (o *SyntheticsTestFileMultipartPresignedUrlsRequest) GetBucketKeyPrefixOk() (*SyntheticsTestFileMultipartPresignedUrlsRequestBucketKeyPrefix, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BucketKeyPrefix, true
}

// SetBucketKeyPrefix sets field value.
func (o *SyntheticsTestFileMultipartPresignedUrlsRequest) SetBucketKeyPrefix(v SyntheticsTestFileMultipartPresignedUrlsRequestBucketKeyPrefix) {
	o.BucketKeyPrefix = v
}

// GetParts returns the Parts field value.
func (o *SyntheticsTestFileMultipartPresignedUrlsRequest) GetParts() []SyntheticsTestFileMultipartPresignedUrlsPart {
	if o == nil {
		var ret []SyntheticsTestFileMultipartPresignedUrlsPart
		return ret
	}
	return o.Parts
}

// GetPartsOk returns a tuple with the Parts field value
// and a boolean to check if the value has been set.
func (o *SyntheticsTestFileMultipartPresignedUrlsRequest) GetPartsOk() (*[]SyntheticsTestFileMultipartPresignedUrlsPart, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parts, true
}

// SetParts sets field value.
func (o *SyntheticsTestFileMultipartPresignedUrlsRequest) SetParts(v []SyntheticsTestFileMultipartPresignedUrlsPart) {
	o.Parts = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestFileMultipartPresignedUrlsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["bucketKeyPrefix"] = o.BucketKeyPrefix
	toSerialize["parts"] = o.Parts

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestFileMultipartPresignedUrlsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BucketKeyPrefix *SyntheticsTestFileMultipartPresignedUrlsRequestBucketKeyPrefix `json:"bucketKeyPrefix"`
		Parts           *[]SyntheticsTestFileMultipartPresignedUrlsPart                 `json:"parts"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.BucketKeyPrefix == nil {
		return fmt.Errorf("required field bucketKeyPrefix missing")
	}
	if all.Parts == nil {
		return fmt.Errorf("required field parts missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"bucketKeyPrefix", "parts"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.BucketKeyPrefix.IsValid() {
		hasInvalidField = true
	} else {
		o.BucketKeyPrefix = *all.BucketKeyPrefix
	}
	o.Parts = *all.Parts

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
