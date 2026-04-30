// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestFileMultipartPresignedUrlsPart A part descriptor for initiating a multipart upload.
type SyntheticsTestFileMultipartPresignedUrlsPart struct {
	// Base64-encoded MD5 digest of the part content.
	Md5 string `json:"md5"`
	// The 1-indexed part number for the multipart upload.
	PartNumber int64 `json:"partNumber"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestFileMultipartPresignedUrlsPart instantiates a new SyntheticsTestFileMultipartPresignedUrlsPart object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestFileMultipartPresignedUrlsPart(md5 string, partNumber int64) *SyntheticsTestFileMultipartPresignedUrlsPart {
	this := SyntheticsTestFileMultipartPresignedUrlsPart{}
	this.Md5 = md5
	this.PartNumber = partNumber
	return &this
}

// NewSyntheticsTestFileMultipartPresignedUrlsPartWithDefaults instantiates a new SyntheticsTestFileMultipartPresignedUrlsPart object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestFileMultipartPresignedUrlsPartWithDefaults() *SyntheticsTestFileMultipartPresignedUrlsPart {
	this := SyntheticsTestFileMultipartPresignedUrlsPart{}
	return &this
}

// GetMd5 returns the Md5 field value.
func (o *SyntheticsTestFileMultipartPresignedUrlsPart) GetMd5() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Md5
}

// GetMd5Ok returns a tuple with the Md5 field value
// and a boolean to check if the value has been set.
func (o *SyntheticsTestFileMultipartPresignedUrlsPart) GetMd5Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Md5, true
}

// SetMd5 sets field value.
func (o *SyntheticsTestFileMultipartPresignedUrlsPart) SetMd5(v string) {
	o.Md5 = v
}

// GetPartNumber returns the PartNumber field value.
func (o *SyntheticsTestFileMultipartPresignedUrlsPart) GetPartNumber() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value
// and a boolean to check if the value has been set.
func (o *SyntheticsTestFileMultipartPresignedUrlsPart) GetPartNumberOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartNumber, true
}

// SetPartNumber sets field value.
func (o *SyntheticsTestFileMultipartPresignedUrlsPart) SetPartNumber(v int64) {
	o.PartNumber = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestFileMultipartPresignedUrlsPart) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["md5"] = o.Md5
	toSerialize["partNumber"] = o.PartNumber

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestFileMultipartPresignedUrlsPart) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Md5        *string `json:"md5"`
		PartNumber *int64  `json:"partNumber"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Md5 == nil {
		return fmt.Errorf("required field md5 missing")
	}
	if all.PartNumber == nil {
		return fmt.Errorf("required field partNumber missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"md5", "partNumber"})
	} else {
		return err
	}
	o.Md5 = *all.Md5
	o.PartNumber = *all.PartNumber

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
