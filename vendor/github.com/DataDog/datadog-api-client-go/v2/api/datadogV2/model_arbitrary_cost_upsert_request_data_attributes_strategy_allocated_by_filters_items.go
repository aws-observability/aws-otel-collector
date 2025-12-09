// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByFiltersItems The definition of `ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByFiltersItems` object.
type ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByFiltersItems struct {
	// The `items` `condition`.
	Condition string `json:"condition"`
	// The `items` `tag`.
	Tag string `json:"tag"`
	// The `items` `value`.
	Value *string `json:"value,omitempty"`
	// The `items` `values`.
	Values datadog.NullableList[string] `json:"values,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByFiltersItems instantiates a new ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByFiltersItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByFiltersItems(condition string, tag string) *ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByFiltersItems {
	this := ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByFiltersItems{}
	this.Condition = condition
	this.Tag = tag
	return &this
}

// NewArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByFiltersItemsWithDefaults instantiates a new ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByFiltersItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByFiltersItemsWithDefaults() *ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByFiltersItems {
	this := ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByFiltersItems{}
	return &this
}

// GetCondition returns the Condition field value.
func (o *ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByFiltersItems) GetCondition() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Condition
}

// GetConditionOk returns a tuple with the Condition field value
// and a boolean to check if the value has been set.
func (o *ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByFiltersItems) GetConditionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Condition, true
}

// SetCondition sets field value.
func (o *ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByFiltersItems) SetCondition(v string) {
	o.Condition = v
}

// GetTag returns the Tag field value.
func (o *ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByFiltersItems) GetTag() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
func (o *ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByFiltersItems) GetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tag, true
}

// SetTag sets field value.
func (o *ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByFiltersItems) SetTag(v string) {
	o.Tag = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByFiltersItems) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByFiltersItems) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByFiltersItems) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByFiltersItems) SetValue(v string) {
	o.Value = &v
}

// GetValues returns the Values field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByFiltersItems) GetValues() []string {
	if o == nil || o.Values.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Values.Get()
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByFiltersItems) GetValuesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values.Get(), o.Values.IsSet()
}

// HasValues returns a boolean if a field has been set.
func (o *ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByFiltersItems) HasValues() bool {
	return o != nil && o.Values.IsSet()
}

// SetValues gets a reference to the given datadog.NullableList[string] and assigns it to the Values field.
func (o *ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByFiltersItems) SetValues(v []string) {
	o.Values.Set(&v)
}

// SetValuesNil sets the value for Values to be an explicit nil.
func (o *ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByFiltersItems) SetValuesNil() {
	o.Values.Set(nil)
}

// UnsetValues ensures that no value is present for Values, not even an explicit nil.
func (o *ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByFiltersItems) UnsetValues() {
	o.Values.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByFiltersItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["condition"] = o.Condition
	toSerialize["tag"] = o.Tag
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Values.IsSet() {
		toSerialize["values"] = o.Values.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ArbitraryCostUpsertRequestDataAttributesStrategyAllocatedByFiltersItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Condition *string                      `json:"condition"`
		Tag       *string                      `json:"tag"`
		Value     *string                      `json:"value,omitempty"`
		Values    datadog.NullableList[string] `json:"values,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Condition == nil {
		return fmt.Errorf("required field condition missing")
	}
	if all.Tag == nil {
		return fmt.Errorf("required field tag missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"condition", "tag", "value", "values"})
	} else {
		return err
	}
	o.Condition = *all.Condition
	o.Tag = *all.Tag
	o.Value = all.Value
	o.Values = all.Values

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
