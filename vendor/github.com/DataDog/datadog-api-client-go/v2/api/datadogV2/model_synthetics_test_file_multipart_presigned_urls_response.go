// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestFileMultipartPresignedUrlsResponse Response containing presigned URLs for multipart file upload and the bucket key.
type SyntheticsTestFileMultipartPresignedUrlsResponse struct {
	// The bucket key that references the uploaded file after completion.
	BucketKey *string `json:"bucketKey,omitempty"`
	// Presigned URL parameters returned for a multipart upload.
	MultipartPresignedUrlsParams *SyntheticsTestFileMultipartPresignedUrlsParams `json:"multipart_presigned_urls_params,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestFileMultipartPresignedUrlsResponse instantiates a new SyntheticsTestFileMultipartPresignedUrlsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestFileMultipartPresignedUrlsResponse() *SyntheticsTestFileMultipartPresignedUrlsResponse {
	this := SyntheticsTestFileMultipartPresignedUrlsResponse{}
	return &this
}

// NewSyntheticsTestFileMultipartPresignedUrlsResponseWithDefaults instantiates a new SyntheticsTestFileMultipartPresignedUrlsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestFileMultipartPresignedUrlsResponseWithDefaults() *SyntheticsTestFileMultipartPresignedUrlsResponse {
	this := SyntheticsTestFileMultipartPresignedUrlsResponse{}
	return &this
}

// GetBucketKey returns the BucketKey field value if set, zero value otherwise.
func (o *SyntheticsTestFileMultipartPresignedUrlsResponse) GetBucketKey() string {
	if o == nil || o.BucketKey == nil {
		var ret string
		return ret
	}
	return *o.BucketKey
}

// GetBucketKeyOk returns a tuple with the BucketKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestFileMultipartPresignedUrlsResponse) GetBucketKeyOk() (*string, bool) {
	if o == nil || o.BucketKey == nil {
		return nil, false
	}
	return o.BucketKey, true
}

// HasBucketKey returns a boolean if a field has been set.
func (o *SyntheticsTestFileMultipartPresignedUrlsResponse) HasBucketKey() bool {
	return o != nil && o.BucketKey != nil
}

// SetBucketKey gets a reference to the given string and assigns it to the BucketKey field.
func (o *SyntheticsTestFileMultipartPresignedUrlsResponse) SetBucketKey(v string) {
	o.BucketKey = &v
}

// GetMultipartPresignedUrlsParams returns the MultipartPresignedUrlsParams field value if set, zero value otherwise.
func (o *SyntheticsTestFileMultipartPresignedUrlsResponse) GetMultipartPresignedUrlsParams() SyntheticsTestFileMultipartPresignedUrlsParams {
	if o == nil || o.MultipartPresignedUrlsParams == nil {
		var ret SyntheticsTestFileMultipartPresignedUrlsParams
		return ret
	}
	return *o.MultipartPresignedUrlsParams
}

// GetMultipartPresignedUrlsParamsOk returns a tuple with the MultipartPresignedUrlsParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestFileMultipartPresignedUrlsResponse) GetMultipartPresignedUrlsParamsOk() (*SyntheticsTestFileMultipartPresignedUrlsParams, bool) {
	if o == nil || o.MultipartPresignedUrlsParams == nil {
		return nil, false
	}
	return o.MultipartPresignedUrlsParams, true
}

// HasMultipartPresignedUrlsParams returns a boolean if a field has been set.
func (o *SyntheticsTestFileMultipartPresignedUrlsResponse) HasMultipartPresignedUrlsParams() bool {
	return o != nil && o.MultipartPresignedUrlsParams != nil
}

// SetMultipartPresignedUrlsParams gets a reference to the given SyntheticsTestFileMultipartPresignedUrlsParams and assigns it to the MultipartPresignedUrlsParams field.
func (o *SyntheticsTestFileMultipartPresignedUrlsResponse) SetMultipartPresignedUrlsParams(v SyntheticsTestFileMultipartPresignedUrlsParams) {
	o.MultipartPresignedUrlsParams = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestFileMultipartPresignedUrlsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.BucketKey != nil {
		toSerialize["bucketKey"] = o.BucketKey
	}
	if o.MultipartPresignedUrlsParams != nil {
		toSerialize["multipart_presigned_urls_params"] = o.MultipartPresignedUrlsParams
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestFileMultipartPresignedUrlsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BucketKey                    *string                                         `json:"bucketKey,omitempty"`
		MultipartPresignedUrlsParams *SyntheticsTestFileMultipartPresignedUrlsParams `json:"multipart_presigned_urls_params,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"bucketKey", "multipart_presigned_urls_params"})
	} else {
		return err
	}

	hasInvalidField := false
	o.BucketKey = all.BucketKey
	if all.MultipartPresignedUrlsParams != nil && all.MultipartPresignedUrlsParams.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.MultipartPresignedUrlsParams = all.MultipartPresignedUrlsParams

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
