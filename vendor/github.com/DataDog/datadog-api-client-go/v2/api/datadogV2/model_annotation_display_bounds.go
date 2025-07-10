// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnnotationDisplayBounds The definition of `AnnotationDisplayBounds` object.
type AnnotationDisplayBounds struct {
	// The `bounds` `height`.
	Height *float64 `json:"height,omitempty"`
	// The `bounds` `width`.
	Width *float64 `json:"width,omitempty"`
	// The `bounds` `x`.
	X *float64 `json:"x,omitempty"`
	// The `bounds` `y`.
	Y *float64 `json:"y,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAnnotationDisplayBounds instantiates a new AnnotationDisplayBounds object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAnnotationDisplayBounds() *AnnotationDisplayBounds {
	this := AnnotationDisplayBounds{}
	return &this
}

// NewAnnotationDisplayBoundsWithDefaults instantiates a new AnnotationDisplayBounds object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAnnotationDisplayBoundsWithDefaults() *AnnotationDisplayBounds {
	this := AnnotationDisplayBounds{}
	return &this
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *AnnotationDisplayBounds) GetHeight() float64 {
	if o == nil || o.Height == nil {
		var ret float64
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotationDisplayBounds) GetHeightOk() (*float64, bool) {
	if o == nil || o.Height == nil {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *AnnotationDisplayBounds) HasHeight() bool {
	return o != nil && o.Height != nil
}

// SetHeight gets a reference to the given float64 and assigns it to the Height field.
func (o *AnnotationDisplayBounds) SetHeight(v float64) {
	o.Height = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *AnnotationDisplayBounds) GetWidth() float64 {
	if o == nil || o.Width == nil {
		var ret float64
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotationDisplayBounds) GetWidthOk() (*float64, bool) {
	if o == nil || o.Width == nil {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *AnnotationDisplayBounds) HasWidth() bool {
	return o != nil && o.Width != nil
}

// SetWidth gets a reference to the given float64 and assigns it to the Width field.
func (o *AnnotationDisplayBounds) SetWidth(v float64) {
	o.Width = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *AnnotationDisplayBounds) GetX() float64 {
	if o == nil || o.X == nil {
		var ret float64
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotationDisplayBounds) GetXOk() (*float64, bool) {
	if o == nil || o.X == nil {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *AnnotationDisplayBounds) HasX() bool {
	return o != nil && o.X != nil
}

// SetX gets a reference to the given float64 and assigns it to the X field.
func (o *AnnotationDisplayBounds) SetX(v float64) {
	o.X = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *AnnotationDisplayBounds) GetY() float64 {
	if o == nil || o.Y == nil {
		var ret float64
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotationDisplayBounds) GetYOk() (*float64, bool) {
	if o == nil || o.Y == nil {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *AnnotationDisplayBounds) HasY() bool {
	return o != nil && o.Y != nil
}

// SetY gets a reference to the given float64 and assigns it to the Y field.
func (o *AnnotationDisplayBounds) SetY(v float64) {
	o.Y = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AnnotationDisplayBounds) MarshalJSON() ([]byte, error) {
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
func (o *AnnotationDisplayBounds) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Height *float64 `json:"height,omitempty"`
		Width  *float64 `json:"width,omitempty"`
		X      *float64 `json:"x,omitempty"`
		Y      *float64 `json:"y,omitempty"`
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
