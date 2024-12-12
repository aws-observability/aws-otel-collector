// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsMobileTestBinding Objects describing the binding used for a mobile test.
type SyntheticsMobileTestBinding struct {
	// Object describing the binding used for a mobile test.
	Items *SyntheticsMobileTestBindingItems `json:"items,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsMobileTestBinding instantiates a new SyntheticsMobileTestBinding object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsMobileTestBinding() *SyntheticsMobileTestBinding {
	this := SyntheticsMobileTestBinding{}
	return &this
}

// NewSyntheticsMobileTestBindingWithDefaults instantiates a new SyntheticsMobileTestBinding object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsMobileTestBindingWithDefaults() *SyntheticsMobileTestBinding {
	this := SyntheticsMobileTestBinding{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SyntheticsMobileTestBinding) GetItems() SyntheticsMobileTestBindingItems {
	if o == nil || o.Items == nil {
		var ret SyntheticsMobileTestBindingItems
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileTestBinding) GetItemsOk() (*SyntheticsMobileTestBindingItems, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SyntheticsMobileTestBinding) HasItems() bool {
	return o != nil && o.Items != nil
}

// SetItems gets a reference to the given SyntheticsMobileTestBindingItems and assigns it to the Items field.
func (o *SyntheticsMobileTestBinding) SetItems(v SyntheticsMobileTestBindingItems) {
	o.Items = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsMobileTestBinding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsMobileTestBinding) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Items *SyntheticsMobileTestBindingItems `json:"items,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"items"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Items != nil && all.Items.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Items = all.Items

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
