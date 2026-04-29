// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultBounds Bounding box of an element on the page.
type SyntheticsTestResultBounds struct {
	// Height in pixels.
	Height *int64 `json:"height,omitempty"`
	// Width in pixels.
	Width *int64 `json:"width,omitempty"`
	// Horizontal position in pixels.
	X *int64 `json:"x,omitempty"`
	// Vertical position in pixels.
	Y *int64 `json:"y,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultBounds instantiates a new SyntheticsTestResultBounds object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultBounds() *SyntheticsTestResultBounds {
	this := SyntheticsTestResultBounds{}
	return &this
}

// NewSyntheticsTestResultBoundsWithDefaults instantiates a new SyntheticsTestResultBounds object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultBoundsWithDefaults() *SyntheticsTestResultBounds {
	this := SyntheticsTestResultBounds{}
	return &this
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *SyntheticsTestResultBounds) GetHeight() int64 {
	if o == nil || o.Height == nil {
		var ret int64
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultBounds) GetHeightOk() (*int64, bool) {
	if o == nil || o.Height == nil {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *SyntheticsTestResultBounds) HasHeight() bool {
	return o != nil && o.Height != nil
}

// SetHeight gets a reference to the given int64 and assigns it to the Height field.
func (o *SyntheticsTestResultBounds) SetHeight(v int64) {
	o.Height = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *SyntheticsTestResultBounds) GetWidth() int64 {
	if o == nil || o.Width == nil {
		var ret int64
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultBounds) GetWidthOk() (*int64, bool) {
	if o == nil || o.Width == nil {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *SyntheticsTestResultBounds) HasWidth() bool {
	return o != nil && o.Width != nil
}

// SetWidth gets a reference to the given int64 and assigns it to the Width field.
func (o *SyntheticsTestResultBounds) SetWidth(v int64) {
	o.Width = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *SyntheticsTestResultBounds) GetX() int64 {
	if o == nil || o.X == nil {
		var ret int64
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultBounds) GetXOk() (*int64, bool) {
	if o == nil || o.X == nil {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *SyntheticsTestResultBounds) HasX() bool {
	return o != nil && o.X != nil
}

// SetX gets a reference to the given int64 and assigns it to the X field.
func (o *SyntheticsTestResultBounds) SetX(v int64) {
	o.X = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *SyntheticsTestResultBounds) GetY() int64 {
	if o == nil || o.Y == nil {
		var ret int64
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultBounds) GetYOk() (*int64, bool) {
	if o == nil || o.Y == nil {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *SyntheticsTestResultBounds) HasY() bool {
	return o != nil && o.Y != nil
}

// SetY gets a reference to the given int64 and assigns it to the Y field.
func (o *SyntheticsTestResultBounds) SetY(v int64) {
	o.Y = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultBounds) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Height != nil {
		toSerialize["height"] = o.Height
	}
	if o.Width != nil {
		toSerialize["width"] = o.Width
	}
	if o.X != nil {
		toSerialize["x"] = o.X
	}
	if o.Y != nil {
		toSerialize["y"] = o.Y
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultBounds) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Height *int64 `json:"height,omitempty"`
		Width  *int64 `json:"width,omitempty"`
		X      *int64 `json:"x,omitempty"`
		Y      *int64 `json:"y,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"height", "width", "x", "y"})
	} else {
		return err
	}
	o.Height = all.Height
	o.Width = all.Width
	o.X = all.X
	o.Y = all.Y

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
