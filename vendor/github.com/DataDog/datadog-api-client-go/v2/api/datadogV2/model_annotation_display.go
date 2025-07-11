// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnnotationDisplay The definition of `AnnotationDisplay` object.
type AnnotationDisplay struct {
	// The definition of `AnnotationDisplayBounds` object.
	Bounds *AnnotationDisplayBounds `json:"bounds,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAnnotationDisplay instantiates a new AnnotationDisplay object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAnnotationDisplay() *AnnotationDisplay {
	this := AnnotationDisplay{}
	return &this
}

// NewAnnotationDisplayWithDefaults instantiates a new AnnotationDisplay object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAnnotationDisplayWithDefaults() *AnnotationDisplay {
	this := AnnotationDisplay{}
	return &this
}

// GetBounds returns the Bounds field value if set, zero value otherwise.
func (o *AnnotationDisplay) GetBounds() AnnotationDisplayBounds {
	if o == nil || o.Bounds == nil {
		var ret AnnotationDisplayBounds
		return ret
	}
	return *o.Bounds
}

// GetBoundsOk returns a tuple with the Bounds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotationDisplay) GetBoundsOk() (*AnnotationDisplayBounds, bool) {
	if o == nil || o.Bounds == nil {
		return nil, false
	}
	return o.Bounds, true
}

// HasBounds returns a boolean if a field has been set.
func (o *AnnotationDisplay) HasBounds() bool {
	return o != nil && o.Bounds != nil
}

// SetBounds gets a reference to the given AnnotationDisplayBounds and assigns it to the Bounds field.
func (o *AnnotationDisplay) SetBounds(v AnnotationDisplayBounds) {
	o.Bounds = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AnnotationDisplay) MarshalJSON() ([]byte, error) {
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
func (o *AnnotationDisplay) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Bounds *AnnotationDisplayBounds `json:"bounds,omitempty"`
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
