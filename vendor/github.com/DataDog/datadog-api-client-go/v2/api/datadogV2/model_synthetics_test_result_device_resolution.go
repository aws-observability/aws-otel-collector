// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultDeviceResolution Screen resolution of the device used to run the test.
type SyntheticsTestResultDeviceResolution struct {
	// Viewport height in pixels.
	Height *int64 `json:"height,omitempty"`
	// Device pixel ratio.
	PixelRatio *float64 `json:"pixel_ratio,omitempty"`
	// Viewport width in pixels.
	Width *int64 `json:"width,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultDeviceResolution instantiates a new SyntheticsTestResultDeviceResolution object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultDeviceResolution() *SyntheticsTestResultDeviceResolution {
	this := SyntheticsTestResultDeviceResolution{}
	return &this
}

// NewSyntheticsTestResultDeviceResolutionWithDefaults instantiates a new SyntheticsTestResultDeviceResolution object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultDeviceResolutionWithDefaults() *SyntheticsTestResultDeviceResolution {
	this := SyntheticsTestResultDeviceResolution{}
	return &this
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *SyntheticsTestResultDeviceResolution) GetHeight() int64 {
	if o == nil || o.Height == nil {
		var ret int64
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDeviceResolution) GetHeightOk() (*int64, bool) {
	if o == nil || o.Height == nil {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *SyntheticsTestResultDeviceResolution) HasHeight() bool {
	return o != nil && o.Height != nil
}

// SetHeight gets a reference to the given int64 and assigns it to the Height field.
func (o *SyntheticsTestResultDeviceResolution) SetHeight(v int64) {
	o.Height = &v
}

// GetPixelRatio returns the PixelRatio field value if set, zero value otherwise.
func (o *SyntheticsTestResultDeviceResolution) GetPixelRatio() float64 {
	if o == nil || o.PixelRatio == nil {
		var ret float64
		return ret
	}
	return *o.PixelRatio
}

// GetPixelRatioOk returns a tuple with the PixelRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDeviceResolution) GetPixelRatioOk() (*float64, bool) {
	if o == nil || o.PixelRatio == nil {
		return nil, false
	}
	return o.PixelRatio, true
}

// HasPixelRatio returns a boolean if a field has been set.
func (o *SyntheticsTestResultDeviceResolution) HasPixelRatio() bool {
	return o != nil && o.PixelRatio != nil
}

// SetPixelRatio gets a reference to the given float64 and assigns it to the PixelRatio field.
func (o *SyntheticsTestResultDeviceResolution) SetPixelRatio(v float64) {
	o.PixelRatio = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *SyntheticsTestResultDeviceResolution) GetWidth() int64 {
	if o == nil || o.Width == nil {
		var ret int64
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDeviceResolution) GetWidthOk() (*int64, bool) {
	if o == nil || o.Width == nil {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *SyntheticsTestResultDeviceResolution) HasWidth() bool {
	return o != nil && o.Width != nil
}

// SetWidth gets a reference to the given int64 and assigns it to the Width field.
func (o *SyntheticsTestResultDeviceResolution) SetWidth(v int64) {
	o.Width = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultDeviceResolution) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Height != nil {
		toSerialize["height"] = o.Height
	}
	if o.PixelRatio != nil {
		toSerialize["pixel_ratio"] = o.PixelRatio
	}
	if o.Width != nil {
		toSerialize["width"] = o.Width
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultDeviceResolution) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Height     *int64   `json:"height,omitempty"`
		PixelRatio *float64 `json:"pixel_ratio,omitempty"`
		Width      *int64   `json:"width,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"height", "pixel_ratio", "width"})
	} else {
		return err
	}
	o.Height = all.Height
	o.PixelRatio = all.PixelRatio
	o.Width = all.Width

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
