// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RulesetRespDataAttributesRulesItemsQueryAddition The definition of `RulesetRespDataAttributesRulesItemsQueryAddition` object.
type RulesetRespDataAttributesRulesItemsQueryAddition struct {
	// The `addition` `key`.
	Key string `json:"key"`
	// The `addition` `value`.
	Value string `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRulesetRespDataAttributesRulesItemsQueryAddition instantiates a new RulesetRespDataAttributesRulesItemsQueryAddition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRulesetRespDataAttributesRulesItemsQueryAddition(key string, value string) *RulesetRespDataAttributesRulesItemsQueryAddition {
	this := RulesetRespDataAttributesRulesItemsQueryAddition{}
	this.Key = key
	this.Value = value
	return &this
}

// NewRulesetRespDataAttributesRulesItemsQueryAdditionWithDefaults instantiates a new RulesetRespDataAttributesRulesItemsQueryAddition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRulesetRespDataAttributesRulesItemsQueryAdditionWithDefaults() *RulesetRespDataAttributesRulesItemsQueryAddition {
	this := RulesetRespDataAttributesRulesItemsQueryAddition{}
	return &this
}

// GetKey returns the Key field value.
func (o *RulesetRespDataAttributesRulesItemsQueryAddition) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *RulesetRespDataAttributesRulesItemsQueryAddition) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value.
func (o *RulesetRespDataAttributesRulesItemsQueryAddition) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value.
func (o *RulesetRespDataAttributesRulesItemsQueryAddition) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *RulesetRespDataAttributesRulesItemsQueryAddition) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *RulesetRespDataAttributesRulesItemsQueryAddition) SetValue(v string) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RulesetRespDataAttributesRulesItemsQueryAddition) MarshalJSON() ([]byte, error) {
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
func (o *RulesetRespDataAttributesRulesItemsQueryAddition) UnmarshalJSON(bytes []byte) (err error) {
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

// NullableRulesetRespDataAttributesRulesItemsQueryAddition handles when a null is used for RulesetRespDataAttributesRulesItemsQueryAddition.
type NullableRulesetRespDataAttributesRulesItemsQueryAddition struct {
	value *RulesetRespDataAttributesRulesItemsQueryAddition
	isSet bool
}

// Get returns the associated value.
func (v NullableRulesetRespDataAttributesRulesItemsQueryAddition) Get() *RulesetRespDataAttributesRulesItemsQueryAddition {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableRulesetRespDataAttributesRulesItemsQueryAddition) Set(val *RulesetRespDataAttributesRulesItemsQueryAddition) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableRulesetRespDataAttributesRulesItemsQueryAddition) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableRulesetRespDataAttributesRulesItemsQueryAddition) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRulesetRespDataAttributesRulesItemsQueryAddition initializes the struct as if Set has been called.
func NewNullableRulesetRespDataAttributesRulesItemsQueryAddition(val *RulesetRespDataAttributesRulesItemsQueryAddition) *NullableRulesetRespDataAttributesRulesItemsQueryAddition {
	return &NullableRulesetRespDataAttributesRulesItemsQueryAddition{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableRulesetRespDataAttributesRulesItemsQueryAddition) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableRulesetRespDataAttributesRulesItemsQueryAddition) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
