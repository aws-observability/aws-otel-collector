// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PowerpackInnerWidgetLayout Powerpack inner widget layout.
type PowerpackInnerWidgetLayout struct {
	// The height of the widget. Should be a non-negative integer.
	Height int64 `json:"height"`
	// The width of the widget. Should be a non-negative integer.
	Width int64 `json:"width"`
	// The position of the widget on the x (horizontal) axis. Should be a non-negative integer.
	X int64 `json:"x"`
	// The position of the widget on the y (vertical) axis. Should be a non-negative integer.
	Y int64 `json:"y"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPowerpackInnerWidgetLayout instantiates a new PowerpackInnerWidgetLayout object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPowerpackInnerWidgetLayout(height int64, width int64, x int64, y int64) *PowerpackInnerWidgetLayout {
	this := PowerpackInnerWidgetLayout{}
	this.Height = height
	this.Width = width
	this.X = x
	this.Y = y
	return &this
}

// NewPowerpackInnerWidgetLayoutWithDefaults instantiates a new PowerpackInnerWidgetLayout object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPowerpackInnerWidgetLayoutWithDefaults() *PowerpackInnerWidgetLayout {
	this := PowerpackInnerWidgetLayout{}
	return &this
}

// GetHeight returns the Height field value.
func (o *PowerpackInnerWidgetLayout) GetHeight() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *PowerpackInnerWidgetLayout) GetHeightOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value.
func (o *PowerpackInnerWidgetLayout) SetHeight(v int64) {
	o.Height = v
}

// GetWidth returns the Width field value.
func (o *PowerpackInnerWidgetLayout) GetWidth() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *PowerpackInnerWidgetLayout) GetWidthOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value.
func (o *PowerpackInnerWidgetLayout) SetWidth(v int64) {
	o.Width = v
}

// GetX returns the X field value.
func (o *PowerpackInnerWidgetLayout) GetX() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.X
}

// GetXOk returns a tuple with the X field value
// and a boolean to check if the value has been set.
func (o *PowerpackInnerWidgetLayout) GetXOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X, true
}

// SetX sets field value.
func (o *PowerpackInnerWidgetLayout) SetX(v int64) {
	o.X = v
}

// GetY returns the Y field value.
func (o *PowerpackInnerWidgetLayout) GetY() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Y
}

// GetYOk returns a tuple with the Y field value
// and a boolean to check if the value has been set.
func (o *PowerpackInnerWidgetLayout) GetYOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Y, true
}

// SetY sets field value.
func (o *PowerpackInnerWidgetLayout) SetY(v int64) {
	o.Y = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PowerpackInnerWidgetLayout) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["height"] = o.Height
	toSerialize["width"] = o.Width
	toSerialize["x"] = o.X
	toSerialize["y"] = o.Y

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PowerpackInnerWidgetLayout) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Height *int64 `json:"height"`
		Width  *int64 `json:"width"`
		X      *int64 `json:"x"`
		Y      *int64 `json:"y"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Height == nil {
		return fmt.Errorf("required field height missing")
	}
	if all.Width == nil {
		return fmt.Errorf("required field width missing")
	}
	if all.X == nil {
		return fmt.Errorf("required field x missing")
	}
	if all.Y == nil {
		return fmt.Errorf("required field y missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"height", "width", "x", "y"})
	} else {
		return err
	}
	o.Height = *all.Height
	o.Width = *all.Width
	o.X = *all.X
	o.Y = *all.Y

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
