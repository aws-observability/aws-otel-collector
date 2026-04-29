// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestFileMultipartPresignedUrlsParams Presigned URL parameters returned for a multipart upload.
type SyntheticsTestFileMultipartPresignedUrlsParams struct {
	// The full storage path for the file being uploaded.
	Key *string `json:"key,omitempty"`
	// The upload ID assigned by the storage provider for this multipart upload.
	UploadId *string `json:"upload_id,omitempty"`
	// A map of part numbers to presigned upload URLs.
	Urls map[string]string `json:"urls,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestFileMultipartPresignedUrlsParams instantiates a new SyntheticsTestFileMultipartPresignedUrlsParams object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestFileMultipartPresignedUrlsParams() *SyntheticsTestFileMultipartPresignedUrlsParams {
	this := SyntheticsTestFileMultipartPresignedUrlsParams{}
	return &this
}

// NewSyntheticsTestFileMultipartPresignedUrlsParamsWithDefaults instantiates a new SyntheticsTestFileMultipartPresignedUrlsParams object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestFileMultipartPresignedUrlsParamsWithDefaults() *SyntheticsTestFileMultipartPresignedUrlsParams {
	this := SyntheticsTestFileMultipartPresignedUrlsParams{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *SyntheticsTestFileMultipartPresignedUrlsParams) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestFileMultipartPresignedUrlsParams) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *SyntheticsTestFileMultipartPresignedUrlsParams) HasKey() bool {
	return o != nil && o.Key != nil
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *SyntheticsTestFileMultipartPresignedUrlsParams) SetKey(v string) {
	o.Key = &v
}

// GetUploadId returns the UploadId field value if set, zero value otherwise.
func (o *SyntheticsTestFileMultipartPresignedUrlsParams) GetUploadId() string {
	if o == nil || o.UploadId == nil {
		var ret string
		return ret
	}
	return *o.UploadId
}

// GetUploadIdOk returns a tuple with the UploadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestFileMultipartPresignedUrlsParams) GetUploadIdOk() (*string, bool) {
	if o == nil || o.UploadId == nil {
		return nil, false
	}
	return o.UploadId, true
}

// HasUploadId returns a boolean if a field has been set.
func (o *SyntheticsTestFileMultipartPresignedUrlsParams) HasUploadId() bool {
	return o != nil && o.UploadId != nil
}

// SetUploadId gets a reference to the given string and assigns it to the UploadId field.
func (o *SyntheticsTestFileMultipartPresignedUrlsParams) SetUploadId(v string) {
	o.UploadId = &v
}

// GetUrls returns the Urls field value if set, zero value otherwise.
func (o *SyntheticsTestFileMultipartPresignedUrlsParams) GetUrls() map[string]string {
	if o == nil || o.Urls == nil {
		var ret map[string]string
		return ret
	}
	return o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestFileMultipartPresignedUrlsParams) GetUrlsOk() (*map[string]string, bool) {
	if o == nil || o.Urls == nil {
		return nil, false
	}
	return &o.Urls, true
}

// HasUrls returns a boolean if a field has been set.
func (o *SyntheticsTestFileMultipartPresignedUrlsParams) HasUrls() bool {
	return o != nil && o.Urls != nil
}

// SetUrls gets a reference to the given map[string]string and assigns it to the Urls field.
func (o *SyntheticsTestFileMultipartPresignedUrlsParams) SetUrls(v map[string]string) {
	o.Urls = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestFileMultipartPresignedUrlsParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.UploadId != nil {
		toSerialize["upload_id"] = o.UploadId
	}
	if o.Urls != nil {
		toSerialize["urls"] = o.Urls
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestFileMultipartPresignedUrlsParams) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Key      *string           `json:"key,omitempty"`
		UploadId *string           `json:"upload_id,omitempty"`
		Urls     map[string]string `json:"urls,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"key", "upload_id", "urls"})
	} else {
		return err
	}
	o.Key = all.Key
	o.UploadId = all.UploadId
	o.Urls = all.Urls

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
