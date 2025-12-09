// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SingleAggregatedDnsResponseDataAttributesGroupByItems Attributes associated with a group by
type SingleAggregatedDnsResponseDataAttributesGroupByItems struct {
	// The group by key.
	Key *string `json:"key,omitempty"`
	// The group by value.
	Value *string `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSingleAggregatedDnsResponseDataAttributesGroupByItems instantiates a new SingleAggregatedDnsResponseDataAttributesGroupByItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSingleAggregatedDnsResponseDataAttributesGroupByItems() *SingleAggregatedDnsResponseDataAttributesGroupByItems {
	this := SingleAggregatedDnsResponseDataAttributesGroupByItems{}
	return &this
}

// NewSingleAggregatedDnsResponseDataAttributesGroupByItemsWithDefaults instantiates a new SingleAggregatedDnsResponseDataAttributesGroupByItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSingleAggregatedDnsResponseDataAttributesGroupByItemsWithDefaults() *SingleAggregatedDnsResponseDataAttributesGroupByItems {
	this := SingleAggregatedDnsResponseDataAttributesGroupByItems{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *SingleAggregatedDnsResponseDataAttributesGroupByItems) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleAggregatedDnsResponseDataAttributesGroupByItems) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *SingleAggregatedDnsResponseDataAttributesGroupByItems) HasKey() bool {
	return o != nil && o.Key != nil
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *SingleAggregatedDnsResponseDataAttributesGroupByItems) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SingleAggregatedDnsResponseDataAttributesGroupByItems) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleAggregatedDnsResponseDataAttributesGroupByItems) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SingleAggregatedDnsResponseDataAttributesGroupByItems) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SingleAggregatedDnsResponseDataAttributesGroupByItems) SetValue(v string) {
	o.Value = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SingleAggregatedDnsResponseDataAttributesGroupByItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
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
func (o *SingleAggregatedDnsResponseDataAttributesGroupByItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Key   *string `json:"key,omitempty"`
		Value *string `json:"value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"key", "value"})
	} else {
		return err
	}
	o.Key = all.Key
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
