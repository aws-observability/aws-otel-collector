// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ArbitraryRuleResponseDataAttributesStrategy The definition of `ArbitraryRuleResponseDataAttributesStrategy` object.
type ArbitraryRuleResponseDataAttributesStrategy struct {
	// The `strategy` `allocated_by`.
	AllocatedBy []ArbitraryRuleResponseDataAttributesStrategyAllocatedByItems `json:"allocated_by,omitempty"`
	// The `strategy` `allocated_by_filters`.
	AllocatedByFilters []ArbitraryRuleResponseDataAttributesStrategyAllocatedByFiltersItems `json:"allocated_by_filters,omitempty"`
	// The `strategy` `allocated_by_tag_keys`.
	AllocatedByTagKeys []string `json:"allocated_by_tag_keys,omitempty"`
	// The `strategy` `based_on_costs`.
	BasedOnCosts []ArbitraryRuleResponseDataAttributesStrategyBasedOnCostsItems `json:"based_on_costs,omitempty"`
	// The rule `strategy` `based_on_timeseries`.
	BasedOnTimeseries map[string]interface{} `json:"based_on_timeseries,omitempty"`
	// The `strategy` `evaluate_grouped_by_filters`.
	EvaluateGroupedByFilters []ArbitraryRuleResponseDataAttributesStrategyEvaluateGroupedByFiltersItems `json:"evaluate_grouped_by_filters,omitempty"`
	// The `strategy` `evaluate_grouped_by_tag_keys`.
	EvaluateGroupedByTagKeys []string `json:"evaluate_grouped_by_tag_keys,omitempty"`
	// The `strategy` `granularity`.
	Granularity *string `json:"granularity,omitempty"`
	// The `strategy` `method`.
	Method string `json:"method"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewArbitraryRuleResponseDataAttributesStrategy instantiates a new ArbitraryRuleResponseDataAttributesStrategy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewArbitraryRuleResponseDataAttributesStrategy(method string) *ArbitraryRuleResponseDataAttributesStrategy {
	this := ArbitraryRuleResponseDataAttributesStrategy{}
	this.Method = method
	return &this
}

// NewArbitraryRuleResponseDataAttributesStrategyWithDefaults instantiates a new ArbitraryRuleResponseDataAttributesStrategy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewArbitraryRuleResponseDataAttributesStrategyWithDefaults() *ArbitraryRuleResponseDataAttributesStrategy {
	this := ArbitraryRuleResponseDataAttributesStrategy{}
	return &this
}

// GetAllocatedBy returns the AllocatedBy field value if set, zero value otherwise.
func (o *ArbitraryRuleResponseDataAttributesStrategy) GetAllocatedBy() []ArbitraryRuleResponseDataAttributesStrategyAllocatedByItems {
	if o == nil || o.AllocatedBy == nil {
		var ret []ArbitraryRuleResponseDataAttributesStrategyAllocatedByItems
		return ret
	}
	return o.AllocatedBy
}

// GetAllocatedByOk returns a tuple with the AllocatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArbitraryRuleResponseDataAttributesStrategy) GetAllocatedByOk() (*[]ArbitraryRuleResponseDataAttributesStrategyAllocatedByItems, bool) {
	if o == nil || o.AllocatedBy == nil {
		return nil, false
	}
	return &o.AllocatedBy, true
}

// HasAllocatedBy returns a boolean if a field has been set.
func (o *ArbitraryRuleResponseDataAttributesStrategy) HasAllocatedBy() bool {
	return o != nil && o.AllocatedBy != nil
}

// SetAllocatedBy gets a reference to the given []ArbitraryRuleResponseDataAttributesStrategyAllocatedByItems and assigns it to the AllocatedBy field.
func (o *ArbitraryRuleResponseDataAttributesStrategy) SetAllocatedBy(v []ArbitraryRuleResponseDataAttributesStrategyAllocatedByItems) {
	o.AllocatedBy = v
}

// GetAllocatedByFilters returns the AllocatedByFilters field value if set, zero value otherwise.
func (o *ArbitraryRuleResponseDataAttributesStrategy) GetAllocatedByFilters() []ArbitraryRuleResponseDataAttributesStrategyAllocatedByFiltersItems {
	if o == nil || o.AllocatedByFilters == nil {
		var ret []ArbitraryRuleResponseDataAttributesStrategyAllocatedByFiltersItems
		return ret
	}
	return o.AllocatedByFilters
}

// GetAllocatedByFiltersOk returns a tuple with the AllocatedByFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArbitraryRuleResponseDataAttributesStrategy) GetAllocatedByFiltersOk() (*[]ArbitraryRuleResponseDataAttributesStrategyAllocatedByFiltersItems, bool) {
	if o == nil || o.AllocatedByFilters == nil {
		return nil, false
	}
	return &o.AllocatedByFilters, true
}

// HasAllocatedByFilters returns a boolean if a field has been set.
func (o *ArbitraryRuleResponseDataAttributesStrategy) HasAllocatedByFilters() bool {
	return o != nil && o.AllocatedByFilters != nil
}

// SetAllocatedByFilters gets a reference to the given []ArbitraryRuleResponseDataAttributesStrategyAllocatedByFiltersItems and assigns it to the AllocatedByFilters field.
func (o *ArbitraryRuleResponseDataAttributesStrategy) SetAllocatedByFilters(v []ArbitraryRuleResponseDataAttributesStrategyAllocatedByFiltersItems) {
	o.AllocatedByFilters = v
}

// GetAllocatedByTagKeys returns the AllocatedByTagKeys field value if set, zero value otherwise.
func (o *ArbitraryRuleResponseDataAttributesStrategy) GetAllocatedByTagKeys() []string {
	if o == nil || o.AllocatedByTagKeys == nil {
		var ret []string
		return ret
	}
	return o.AllocatedByTagKeys
}

// GetAllocatedByTagKeysOk returns a tuple with the AllocatedByTagKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArbitraryRuleResponseDataAttributesStrategy) GetAllocatedByTagKeysOk() (*[]string, bool) {
	if o == nil || o.AllocatedByTagKeys == nil {
		return nil, false
	}
	return &o.AllocatedByTagKeys, true
}

// HasAllocatedByTagKeys returns a boolean if a field has been set.
func (o *ArbitraryRuleResponseDataAttributesStrategy) HasAllocatedByTagKeys() bool {
	return o != nil && o.AllocatedByTagKeys != nil
}

// SetAllocatedByTagKeys gets a reference to the given []string and assigns it to the AllocatedByTagKeys field.
func (o *ArbitraryRuleResponseDataAttributesStrategy) SetAllocatedByTagKeys(v []string) {
	o.AllocatedByTagKeys = v
}

// GetBasedOnCosts returns the BasedOnCosts field value if set, zero value otherwise.
func (o *ArbitraryRuleResponseDataAttributesStrategy) GetBasedOnCosts() []ArbitraryRuleResponseDataAttributesStrategyBasedOnCostsItems {
	if o == nil || o.BasedOnCosts == nil {
		var ret []ArbitraryRuleResponseDataAttributesStrategyBasedOnCostsItems
		return ret
	}
	return o.BasedOnCosts
}

// GetBasedOnCostsOk returns a tuple with the BasedOnCosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArbitraryRuleResponseDataAttributesStrategy) GetBasedOnCostsOk() (*[]ArbitraryRuleResponseDataAttributesStrategyBasedOnCostsItems, bool) {
	if o == nil || o.BasedOnCosts == nil {
		return nil, false
	}
	return &o.BasedOnCosts, true
}

// HasBasedOnCosts returns a boolean if a field has been set.
func (o *ArbitraryRuleResponseDataAttributesStrategy) HasBasedOnCosts() bool {
	return o != nil && o.BasedOnCosts != nil
}

// SetBasedOnCosts gets a reference to the given []ArbitraryRuleResponseDataAttributesStrategyBasedOnCostsItems and assigns it to the BasedOnCosts field.
func (o *ArbitraryRuleResponseDataAttributesStrategy) SetBasedOnCosts(v []ArbitraryRuleResponseDataAttributesStrategyBasedOnCostsItems) {
	o.BasedOnCosts = v
}

// GetBasedOnTimeseries returns the BasedOnTimeseries field value if set, zero value otherwise.
func (o *ArbitraryRuleResponseDataAttributesStrategy) GetBasedOnTimeseries() map[string]interface{} {
	if o == nil || o.BasedOnTimeseries == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.BasedOnTimeseries
}

// GetBasedOnTimeseriesOk returns a tuple with the BasedOnTimeseries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArbitraryRuleResponseDataAttributesStrategy) GetBasedOnTimeseriesOk() (*map[string]interface{}, bool) {
	if o == nil || o.BasedOnTimeseries == nil {
		return nil, false
	}
	return &o.BasedOnTimeseries, true
}

// HasBasedOnTimeseries returns a boolean if a field has been set.
func (o *ArbitraryRuleResponseDataAttributesStrategy) HasBasedOnTimeseries() bool {
	return o != nil && o.BasedOnTimeseries != nil
}

// SetBasedOnTimeseries gets a reference to the given map[string]interface{} and assigns it to the BasedOnTimeseries field.
func (o *ArbitraryRuleResponseDataAttributesStrategy) SetBasedOnTimeseries(v map[string]interface{}) {
	o.BasedOnTimeseries = v
}

// GetEvaluateGroupedByFilters returns the EvaluateGroupedByFilters field value if set, zero value otherwise.
func (o *ArbitraryRuleResponseDataAttributesStrategy) GetEvaluateGroupedByFilters() []ArbitraryRuleResponseDataAttributesStrategyEvaluateGroupedByFiltersItems {
	if o == nil || o.EvaluateGroupedByFilters == nil {
		var ret []ArbitraryRuleResponseDataAttributesStrategyEvaluateGroupedByFiltersItems
		return ret
	}
	return o.EvaluateGroupedByFilters
}

// GetEvaluateGroupedByFiltersOk returns a tuple with the EvaluateGroupedByFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArbitraryRuleResponseDataAttributesStrategy) GetEvaluateGroupedByFiltersOk() (*[]ArbitraryRuleResponseDataAttributesStrategyEvaluateGroupedByFiltersItems, bool) {
	if o == nil || o.EvaluateGroupedByFilters == nil {
		return nil, false
	}
	return &o.EvaluateGroupedByFilters, true
}

// HasEvaluateGroupedByFilters returns a boolean if a field has been set.
func (o *ArbitraryRuleResponseDataAttributesStrategy) HasEvaluateGroupedByFilters() bool {
	return o != nil && o.EvaluateGroupedByFilters != nil
}

// SetEvaluateGroupedByFilters gets a reference to the given []ArbitraryRuleResponseDataAttributesStrategyEvaluateGroupedByFiltersItems and assigns it to the EvaluateGroupedByFilters field.
func (o *ArbitraryRuleResponseDataAttributesStrategy) SetEvaluateGroupedByFilters(v []ArbitraryRuleResponseDataAttributesStrategyEvaluateGroupedByFiltersItems) {
	o.EvaluateGroupedByFilters = v
}

// GetEvaluateGroupedByTagKeys returns the EvaluateGroupedByTagKeys field value if set, zero value otherwise.
func (o *ArbitraryRuleResponseDataAttributesStrategy) GetEvaluateGroupedByTagKeys() []string {
	if o == nil || o.EvaluateGroupedByTagKeys == nil {
		var ret []string
		return ret
	}
	return o.EvaluateGroupedByTagKeys
}

// GetEvaluateGroupedByTagKeysOk returns a tuple with the EvaluateGroupedByTagKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArbitraryRuleResponseDataAttributesStrategy) GetEvaluateGroupedByTagKeysOk() (*[]string, bool) {
	if o == nil || o.EvaluateGroupedByTagKeys == nil {
		return nil, false
	}
	return &o.EvaluateGroupedByTagKeys, true
}

// HasEvaluateGroupedByTagKeys returns a boolean if a field has been set.
func (o *ArbitraryRuleResponseDataAttributesStrategy) HasEvaluateGroupedByTagKeys() bool {
	return o != nil && o.EvaluateGroupedByTagKeys != nil
}

// SetEvaluateGroupedByTagKeys gets a reference to the given []string and assigns it to the EvaluateGroupedByTagKeys field.
func (o *ArbitraryRuleResponseDataAttributesStrategy) SetEvaluateGroupedByTagKeys(v []string) {
	o.EvaluateGroupedByTagKeys = v
}

// GetGranularity returns the Granularity field value if set, zero value otherwise.
func (o *ArbitraryRuleResponseDataAttributesStrategy) GetGranularity() string {
	if o == nil || o.Granularity == nil {
		var ret string
		return ret
	}
	return *o.Granularity
}

// GetGranularityOk returns a tuple with the Granularity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArbitraryRuleResponseDataAttributesStrategy) GetGranularityOk() (*string, bool) {
	if o == nil || o.Granularity == nil {
		return nil, false
	}
	return o.Granularity, true
}

// HasGranularity returns a boolean if a field has been set.
func (o *ArbitraryRuleResponseDataAttributesStrategy) HasGranularity() bool {
	return o != nil && o.Granularity != nil
}

// SetGranularity gets a reference to the given string and assigns it to the Granularity field.
func (o *ArbitraryRuleResponseDataAttributesStrategy) SetGranularity(v string) {
	o.Granularity = &v
}

// GetMethod returns the Method field value.
func (o *ArbitraryRuleResponseDataAttributesStrategy) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *ArbitraryRuleResponseDataAttributesStrategy) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value.
func (o *ArbitraryRuleResponseDataAttributesStrategy) SetMethod(v string) {
	o.Method = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ArbitraryRuleResponseDataAttributesStrategy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AllocatedBy != nil {
		toSerialize["allocated_by"] = o.AllocatedBy
	}
	if o.AllocatedByFilters != nil {
		toSerialize["allocated_by_filters"] = o.AllocatedByFilters
	}
	if o.AllocatedByTagKeys != nil {
		toSerialize["allocated_by_tag_keys"] = o.AllocatedByTagKeys
	}
	if o.BasedOnCosts != nil {
		toSerialize["based_on_costs"] = o.BasedOnCosts
	}
	if o.BasedOnTimeseries != nil {
		toSerialize["based_on_timeseries"] = o.BasedOnTimeseries
	}
	if o.EvaluateGroupedByFilters != nil {
		toSerialize["evaluate_grouped_by_filters"] = o.EvaluateGroupedByFilters
	}
	if o.EvaluateGroupedByTagKeys != nil {
		toSerialize["evaluate_grouped_by_tag_keys"] = o.EvaluateGroupedByTagKeys
	}
	if o.Granularity != nil {
		toSerialize["granularity"] = o.Granularity
	}
	toSerialize["method"] = o.Method

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ArbitraryRuleResponseDataAttributesStrategy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AllocatedBy              []ArbitraryRuleResponseDataAttributesStrategyAllocatedByItems              `json:"allocated_by,omitempty"`
		AllocatedByFilters       []ArbitraryRuleResponseDataAttributesStrategyAllocatedByFiltersItems       `json:"allocated_by_filters,omitempty"`
		AllocatedByTagKeys       []string                                                                   `json:"allocated_by_tag_keys,omitempty"`
		BasedOnCosts             []ArbitraryRuleResponseDataAttributesStrategyBasedOnCostsItems             `json:"based_on_costs,omitempty"`
		BasedOnTimeseries        map[string]interface{}                                                     `json:"based_on_timeseries,omitempty"`
		EvaluateGroupedByFilters []ArbitraryRuleResponseDataAttributesStrategyEvaluateGroupedByFiltersItems `json:"evaluate_grouped_by_filters,omitempty"`
		EvaluateGroupedByTagKeys []string                                                                   `json:"evaluate_grouped_by_tag_keys,omitempty"`
		Granularity              *string                                                                    `json:"granularity,omitempty"`
		Method                   *string                                                                    `json:"method"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Method == nil {
		return fmt.Errorf("required field method missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"allocated_by", "allocated_by_filters", "allocated_by_tag_keys", "based_on_costs", "based_on_timeseries", "evaluate_grouped_by_filters", "evaluate_grouped_by_tag_keys", "granularity", "method"})
	} else {
		return err
	}
	o.AllocatedBy = all.AllocatedBy
	o.AllocatedByFilters = all.AllocatedByFilters
	o.AllocatedByTagKeys = all.AllocatedByTagKeys
	o.BasedOnCosts = all.BasedOnCosts
	o.BasedOnTimeseries = all.BasedOnTimeseries
	o.EvaluateGroupedByFilters = all.EvaluateGroupedByFilters
	o.EvaluateGroupedByTagKeys = all.EvaluateGroupedByTagKeys
	o.Granularity = all.Granularity
	o.Method = *all.Method

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
