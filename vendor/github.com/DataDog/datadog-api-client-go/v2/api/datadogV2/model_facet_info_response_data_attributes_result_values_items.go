// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FacetInfoResponseDataAttributesResultValuesItems
type FacetInfoResponseDataAttributesResultValuesItems struct {
	//
	Count *int64 `json:"count,omitempty"`
	//
	Value *string `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFacetInfoResponseDataAttributesResultValuesItems instantiates a new FacetInfoResponseDataAttributesResultValuesItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFacetInfoResponseDataAttributesResultValuesItems() *FacetInfoResponseDataAttributesResultValuesItems {
	this := FacetInfoResponseDataAttributesResultValuesItems{}
	return &this
}

// NewFacetInfoResponseDataAttributesResultValuesItemsWithDefaults instantiates a new FacetInfoResponseDataAttributesResultValuesItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFacetInfoResponseDataAttributesResultValuesItemsWithDefaults() *FacetInfoResponseDataAttributesResultValuesItems {
	this := FacetInfoResponseDataAttributesResultValuesItems{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *FacetInfoResponseDataAttributesResultValuesItems) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacetInfoResponseDataAttributesResultValuesItems) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *FacetInfoResponseDataAttributesResultValuesItems) HasCount() bool {
	return o != nil && o.Count != nil
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *FacetInfoResponseDataAttributesResultValuesItems) SetCount(v int64) {
	o.Count = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *FacetInfoResponseDataAttributesResultValuesItems) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacetInfoResponseDataAttributesResultValuesItems) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *FacetInfoResponseDataAttributesResultValuesItems) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *FacetInfoResponseDataAttributesResultValuesItems) SetValue(v string) {
	o.Value = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FacetInfoResponseDataAttributesResultValuesItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FacetInfoResponseDataAttributesResultValuesItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Count *int64  `json:"count,omitempty"`
		Value *string `json:"value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"count", "value"})
	} else {
		return err
	}
	o.Count = all.Count
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
