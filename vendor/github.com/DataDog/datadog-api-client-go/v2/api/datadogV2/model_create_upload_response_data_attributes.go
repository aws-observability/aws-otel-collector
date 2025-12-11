// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateUploadResponseDataAttributes Pre-signed URLs for uploading parts of the file.
type CreateUploadResponseDataAttributes struct {
	// The pre-signed URLs for uploading parts. These URLs expire after 5 minutes.
	PartUrls []string `json:"part_urls,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateUploadResponseDataAttributes instantiates a new CreateUploadResponseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateUploadResponseDataAttributes() *CreateUploadResponseDataAttributes {
	this := CreateUploadResponseDataAttributes{}
	return &this
}

// NewCreateUploadResponseDataAttributesWithDefaults instantiates a new CreateUploadResponseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateUploadResponseDataAttributesWithDefaults() *CreateUploadResponseDataAttributes {
	this := CreateUploadResponseDataAttributes{}
	return &this
}

// GetPartUrls returns the PartUrls field value if set, zero value otherwise.
func (o *CreateUploadResponseDataAttributes) GetPartUrls() []string {
	if o == nil || o.PartUrls == nil {
		var ret []string
		return ret
	}
	return o.PartUrls
}

// GetPartUrlsOk returns a tuple with the PartUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUploadResponseDataAttributes) GetPartUrlsOk() (*[]string, bool) {
	if o == nil || o.PartUrls == nil {
		return nil, false
	}
	return &o.PartUrls, true
}

// HasPartUrls returns a boolean if a field has been set.
func (o *CreateUploadResponseDataAttributes) HasPartUrls() bool {
	return o != nil && o.PartUrls != nil
}

// SetPartUrls gets a reference to the given []string and assigns it to the PartUrls field.
func (o *CreateUploadResponseDataAttributes) SetPartUrls(v []string) {
	o.PartUrls = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateUploadResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.PartUrls != nil {
		toSerialize["part_urls"] = o.PartUrls
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateUploadResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		PartUrls []string `json:"part_urls,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"part_urls"})
	} else {
		return err
	}
	o.PartUrls = all.PartUrls

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
