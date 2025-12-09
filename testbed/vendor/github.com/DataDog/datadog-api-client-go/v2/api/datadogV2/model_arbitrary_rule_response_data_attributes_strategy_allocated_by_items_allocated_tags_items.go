// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ArbitraryRuleResponseDataAttributesStrategyAllocatedByItemsAllocatedTagsItems The definition of `ArbitraryRuleResponseDataAttributesStrategyAllocatedByItemsAllocatedTagsItems` object.
type ArbitraryRuleResponseDataAttributesStrategyAllocatedByItemsAllocatedTagsItems struct {
	// The `items` `key`.
	Key string `json:"key"`
	// The `items` `value`.
	Value string `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewArbitraryRuleResponseDataAttributesStrategyAllocatedByItemsAllocatedTagsItems instantiates a new ArbitraryRuleResponseDataAttributesStrategyAllocatedByItemsAllocatedTagsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewArbitraryRuleResponseDataAttributesStrategyAllocatedByItemsAllocatedTagsItems(key string, value string) *ArbitraryRuleResponseDataAttributesStrategyAllocatedByItemsAllocatedTagsItems {
	this := ArbitraryRuleResponseDataAttributesStrategyAllocatedByItemsAllocatedTagsItems{}
	this.Key = key
	this.Value = value
	return &this
}

// NewArbitraryRuleResponseDataAttributesStrategyAllocatedByItemsAllocatedTagsItemsWithDefaults instantiates a new ArbitraryRuleResponseDataAttributesStrategyAllocatedByItemsAllocatedTagsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewArbitraryRuleResponseDataAttributesStrategyAllocatedByItemsAllocatedTagsItemsWithDefaults() *ArbitraryRuleResponseDataAttributesStrategyAllocatedByItemsAllocatedTagsItems {
	this := ArbitraryRuleResponseDataAttributesStrategyAllocatedByItemsAllocatedTagsItems{}
	return &this
}

// GetKey returns the Key field value.
func (o *ArbitraryRuleResponseDataAttributesStrategyAllocatedByItemsAllocatedTagsItems) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ArbitraryRuleResponseDataAttributesStrategyAllocatedByItemsAllocatedTagsItems) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value.
func (o *ArbitraryRuleResponseDataAttributesStrategyAllocatedByItemsAllocatedTagsItems) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value.
func (o *ArbitraryRuleResponseDataAttributesStrategyAllocatedByItemsAllocatedTagsItems) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ArbitraryRuleResponseDataAttributesStrategyAllocatedByItemsAllocatedTagsItems) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *ArbitraryRuleResponseDataAttributesStrategyAllocatedByItemsAllocatedTagsItems) SetValue(v string) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ArbitraryRuleResponseDataAttributesStrategyAllocatedByItemsAllocatedTagsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["key"] = o.Key
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ArbitraryRuleResponseDataAttributesStrategyAllocatedByItemsAllocatedTagsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Key   *string `json:"key"`
		Value *string `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Key == nil {
		return fmt.Errorf("required field key missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"key", "value"})
	} else {
		return err
	}
	o.Key = *all.Key
	o.Value = *all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
