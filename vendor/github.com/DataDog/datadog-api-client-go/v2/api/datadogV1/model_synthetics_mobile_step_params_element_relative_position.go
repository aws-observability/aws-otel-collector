// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsMobileStepParamsElementRelativePosition Position of the action relative to the element.
type SyntheticsMobileStepParamsElementRelativePosition struct {
	// The `relativePosition` on the `x` axis for the element.
	X *float64 `json:"x,omitempty"`
	// The `relativePosition` on the `y` axis for the element.
	Y *float64 `json:"y,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsMobileStepParamsElementRelativePosition instantiates a new SyntheticsMobileStepParamsElementRelativePosition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsMobileStepParamsElementRelativePosition() *SyntheticsMobileStepParamsElementRelativePosition {
	this := SyntheticsMobileStepParamsElementRelativePosition{}
	return &this
}

// NewSyntheticsMobileStepParamsElementRelativePositionWithDefaults instantiates a new SyntheticsMobileStepParamsElementRelativePosition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsMobileStepParamsElementRelativePositionWithDefaults() *SyntheticsMobileStepParamsElementRelativePosition {
	this := SyntheticsMobileStepParamsElementRelativePosition{}
	return &this
}

// GetX returns the X field value if set, zero value otherwise.
func (o *SyntheticsMobileStepParamsElementRelativePosition) GetX() float64 {
	if o == nil || o.X == nil {
		var ret float64
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParamsElementRelativePosition) GetXOk() (*float64, bool) {
	if o == nil || o.X == nil {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *SyntheticsMobileStepParamsElementRelativePosition) HasX() bool {
	return o != nil && o.X != nil
}

// SetX gets a reference to the given float64 and assigns it to the X field.
func (o *SyntheticsMobileStepParamsElementRelativePosition) SetX(v float64) {
	o.X = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *SyntheticsMobileStepParamsElementRelativePosition) GetY() float64 {
	if o == nil || o.Y == nil {
		var ret float64
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParamsElementRelativePosition) GetYOk() (*float64, bool) {
	if o == nil || o.Y == nil {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *SyntheticsMobileStepParamsElementRelativePosition) HasY() bool {
	return o != nil && o.Y != nil
}

// SetY gets a reference to the given float64 and assigns it to the Y field.
func (o *SyntheticsMobileStepParamsElementRelativePosition) SetY(v float64) {
	o.Y = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsMobileStepParamsElementRelativePosition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
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
func (o *SyntheticsMobileStepParamsElementRelativePosition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		X *float64 `json:"x,omitempty"`
		Y *float64 `json:"y,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"x", "y"})
	} else {
		return err
	}
	o.X = all.X
	o.Y = all.Y

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
