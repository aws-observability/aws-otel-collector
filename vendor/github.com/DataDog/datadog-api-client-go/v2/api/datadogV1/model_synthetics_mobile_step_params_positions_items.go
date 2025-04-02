// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsMobileStepParamsPositionsItems A description of a single position for a `flick` step type.
type SyntheticsMobileStepParamsPositionsItems struct {
	// The `x` position for the flick.
	X *float64 `json:"x,omitempty"`
	// The `y` position for the flick.
	Y *float64 `json:"y,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsMobileStepParamsPositionsItems instantiates a new SyntheticsMobileStepParamsPositionsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsMobileStepParamsPositionsItems() *SyntheticsMobileStepParamsPositionsItems {
	this := SyntheticsMobileStepParamsPositionsItems{}
	return &this
}

// NewSyntheticsMobileStepParamsPositionsItemsWithDefaults instantiates a new SyntheticsMobileStepParamsPositionsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsMobileStepParamsPositionsItemsWithDefaults() *SyntheticsMobileStepParamsPositionsItems {
	this := SyntheticsMobileStepParamsPositionsItems{}
	return &this
}

// GetX returns the X field value if set, zero value otherwise.
func (o *SyntheticsMobileStepParamsPositionsItems) GetX() float64 {
	if o == nil || o.X == nil {
		var ret float64
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParamsPositionsItems) GetXOk() (*float64, bool) {
	if o == nil || o.X == nil {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *SyntheticsMobileStepParamsPositionsItems) HasX() bool {
	return o != nil && o.X != nil
}

// SetX gets a reference to the given float64 and assigns it to the X field.
func (o *SyntheticsMobileStepParamsPositionsItems) SetX(v float64) {
	o.X = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *SyntheticsMobileStepParamsPositionsItems) GetY() float64 {
	if o == nil || o.Y == nil {
		var ret float64
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParamsPositionsItems) GetYOk() (*float64, bool) {
	if o == nil || o.Y == nil {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *SyntheticsMobileStepParamsPositionsItems) HasY() bool {
	return o != nil && o.Y != nil
}

// SetY gets a reference to the given float64 and assigns it to the Y field.
func (o *SyntheticsMobileStepParamsPositionsItems) SetY(v float64) {
	o.Y = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsMobileStepParamsPositionsItems) MarshalJSON() ([]byte, error) {
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
func (o *SyntheticsMobileStepParamsPositionsItems) UnmarshalJSON(bytes []byte) (err error) {
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
