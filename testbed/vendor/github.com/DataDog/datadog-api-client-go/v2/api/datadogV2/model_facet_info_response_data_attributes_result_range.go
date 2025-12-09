// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FacetInfoResponseDataAttributesResultRange
type FacetInfoResponseDataAttributesResultRange struct {
	//
	Max interface{} `json:"max,omitempty"`
	//
	Min interface{} `json:"min,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFacetInfoResponseDataAttributesResultRange instantiates a new FacetInfoResponseDataAttributesResultRange object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFacetInfoResponseDataAttributesResultRange() *FacetInfoResponseDataAttributesResultRange {
	this := FacetInfoResponseDataAttributesResultRange{}
	return &this
}

// NewFacetInfoResponseDataAttributesResultRangeWithDefaults instantiates a new FacetInfoResponseDataAttributesResultRange object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFacetInfoResponseDataAttributesResultRangeWithDefaults() *FacetInfoResponseDataAttributesResultRange {
	this := FacetInfoResponseDataAttributesResultRange{}
	return &this
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *FacetInfoResponseDataAttributesResultRange) GetMax() interface{} {
	if o == nil || o.Max == nil {
		var ret interface{}
		return ret
	}
	return o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacetInfoResponseDataAttributesResultRange) GetMaxOk() (*interface{}, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return &o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *FacetInfoResponseDataAttributesResultRange) HasMax() bool {
	return o != nil && o.Max != nil
}

// SetMax gets a reference to the given interface{} and assigns it to the Max field.
func (o *FacetInfoResponseDataAttributesResultRange) SetMax(v interface{}) {
	o.Max = v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *FacetInfoResponseDataAttributesResultRange) GetMin() interface{} {
	if o == nil || o.Min == nil {
		var ret interface{}
		return ret
	}
	return o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacetInfoResponseDataAttributesResultRange) GetMinOk() (*interface{}, bool) {
	if o == nil || o.Min == nil {
		return nil, false
	}
	return &o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *FacetInfoResponseDataAttributesResultRange) HasMin() bool {
	return o != nil && o.Min != nil
}

// SetMin gets a reference to the given interface{} and assigns it to the Min field.
func (o *FacetInfoResponseDataAttributesResultRange) SetMin(v interface{}) {
	o.Min = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FacetInfoResponseDataAttributesResultRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Max != nil {
		toSerialize["max"] = o.Max
	}
	if o.Min != nil {
		toSerialize["min"] = o.Min
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FacetInfoResponseDataAttributesResultRange) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Max interface{} `json:"max,omitempty"`
		Min interface{} `json:"min,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"max", "min"})
	} else {
		return err
	}
	o.Max = all.Max
	o.Min = all.Min

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
