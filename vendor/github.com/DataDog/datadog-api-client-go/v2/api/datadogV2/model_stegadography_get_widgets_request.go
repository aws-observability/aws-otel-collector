// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	_io "io"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StegadographyGetWidgetsRequest Multipart form data containing the PNG image to scan for watermarks.
type StegadographyGetWidgetsRequest struct {
	// PNG image file to scan for embedded watermarks.
	Image _io.Reader `json:"image"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStegadographyGetWidgetsRequest instantiates a new StegadographyGetWidgetsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStegadographyGetWidgetsRequest(image _io.Reader) *StegadographyGetWidgetsRequest {
	this := StegadographyGetWidgetsRequest{}
	this.Image = image
	return &this
}

// NewStegadographyGetWidgetsRequestWithDefaults instantiates a new StegadographyGetWidgetsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStegadographyGetWidgetsRequestWithDefaults() *StegadographyGetWidgetsRequest {
	this := StegadographyGetWidgetsRequest{}
	return &this
}

// GetImage returns the Image field value.
func (o *StegadographyGetWidgetsRequest) GetImage() _io.Reader {
	if o == nil {
		var ret _io.Reader
		return ret
	}
	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *StegadographyGetWidgetsRequest) GetImageOk() (*_io.Reader, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value.
func (o *StegadographyGetWidgetsRequest) SetImage(v _io.Reader) {
	o.Image = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StegadographyGetWidgetsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["image"] = o.Image

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StegadographyGetWidgetsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Image *_io.Reader `json:"image"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Image == nil {
		return fmt.Errorf("required field image missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"image"})
	} else {
		return err
	}
	o.Image = *all.Image

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
