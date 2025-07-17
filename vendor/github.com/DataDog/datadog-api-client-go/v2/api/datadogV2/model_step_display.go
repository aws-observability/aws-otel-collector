// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StepDisplay The definition of `StepDisplay` object.
type StepDisplay struct {
	// The definition of `StepDisplayBounds` object.
	Bounds *StepDisplayBounds `json:"bounds,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStepDisplay instantiates a new StepDisplay object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStepDisplay() *StepDisplay {
	this := StepDisplay{}
	return &this
}

// NewStepDisplayWithDefaults instantiates a new StepDisplay object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStepDisplayWithDefaults() *StepDisplay {
	this := StepDisplay{}
	return &this
}

// GetBounds returns the Bounds field value if set, zero value otherwise.
func (o *StepDisplay) GetBounds() StepDisplayBounds {
	if o == nil || o.Bounds == nil {
		var ret StepDisplayBounds
		return ret
	}
	return *o.Bounds
}

// GetBoundsOk returns a tuple with the Bounds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StepDisplay) GetBoundsOk() (*StepDisplayBounds, bool) {
	if o == nil || o.Bounds == nil {
		return nil, false
	}
	return o.Bounds, true
}

// HasBounds returns a boolean if a field has been set.
func (o *StepDisplay) HasBounds() bool {
	return o != nil && o.Bounds != nil
}

// SetBounds gets a reference to the given StepDisplayBounds and assigns it to the Bounds field.
func (o *StepDisplay) SetBounds(v StepDisplayBounds) {
	o.Bounds = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o StepDisplay) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Bounds != nil {
		toSerialize["bounds"] = o.Bounds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StepDisplay) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Bounds *StepDisplayBounds `json:"bounds,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"bounds"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Bounds != nil && all.Bounds.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Bounds = all.Bounds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
