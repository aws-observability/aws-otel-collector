// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagIndexingRuleMetricMatch Criteria for matching metrics based on query state.
type TagIndexingRuleMetricMatch struct {
	// Match metrics that are being queried.
	IsQueried *bool `json:"is_queried,omitempty"`
	// Match metrics that are not being queried.
	NotQueried *bool `json:"not_queried,omitempty"`
	// Match metrics not used in any dashboards or monitors.
	NotUsedInAssets *bool `json:"not_used_in_assets,omitempty"`
	// Window in seconds for evaluating query state.
	QueriedWindowSeconds *int64 `json:"queried_window_seconds,omitempty"`
	// Match metrics used in dashboards or monitors.
	UsedInAssets *bool `json:"used_in_assets,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTagIndexingRuleMetricMatch instantiates a new TagIndexingRuleMetricMatch object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTagIndexingRuleMetricMatch() *TagIndexingRuleMetricMatch {
	this := TagIndexingRuleMetricMatch{}
	return &this
}

// NewTagIndexingRuleMetricMatchWithDefaults instantiates a new TagIndexingRuleMetricMatch object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTagIndexingRuleMetricMatchWithDefaults() *TagIndexingRuleMetricMatch {
	this := TagIndexingRuleMetricMatch{}
	return &this
}

// GetIsQueried returns the IsQueried field value if set, zero value otherwise.
func (o *TagIndexingRuleMetricMatch) GetIsQueried() bool {
	if o == nil || o.IsQueried == nil {
		var ret bool
		return ret
	}
	return *o.IsQueried
}

// GetIsQueriedOk returns a tuple with the IsQueried field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleMetricMatch) GetIsQueriedOk() (*bool, bool) {
	if o == nil || o.IsQueried == nil {
		return nil, false
	}
	return o.IsQueried, true
}

// HasIsQueried returns a boolean if a field has been set.
func (o *TagIndexingRuleMetricMatch) HasIsQueried() bool {
	return o != nil && o.IsQueried != nil
}

// SetIsQueried gets a reference to the given bool and assigns it to the IsQueried field.
func (o *TagIndexingRuleMetricMatch) SetIsQueried(v bool) {
	o.IsQueried = &v
}

// GetNotQueried returns the NotQueried field value if set, zero value otherwise.
func (o *TagIndexingRuleMetricMatch) GetNotQueried() bool {
	if o == nil || o.NotQueried == nil {
		var ret bool
		return ret
	}
	return *o.NotQueried
}

// GetNotQueriedOk returns a tuple with the NotQueried field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleMetricMatch) GetNotQueriedOk() (*bool, bool) {
	if o == nil || o.NotQueried == nil {
		return nil, false
	}
	return o.NotQueried, true
}

// HasNotQueried returns a boolean if a field has been set.
func (o *TagIndexingRuleMetricMatch) HasNotQueried() bool {
	return o != nil && o.NotQueried != nil
}

// SetNotQueried gets a reference to the given bool and assigns it to the NotQueried field.
func (o *TagIndexingRuleMetricMatch) SetNotQueried(v bool) {
	o.NotQueried = &v
}

// GetNotUsedInAssets returns the NotUsedInAssets field value if set, zero value otherwise.
func (o *TagIndexingRuleMetricMatch) GetNotUsedInAssets() bool {
	if o == nil || o.NotUsedInAssets == nil {
		var ret bool
		return ret
	}
	return *o.NotUsedInAssets
}

// GetNotUsedInAssetsOk returns a tuple with the NotUsedInAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleMetricMatch) GetNotUsedInAssetsOk() (*bool, bool) {
	if o == nil || o.NotUsedInAssets == nil {
		return nil, false
	}
	return o.NotUsedInAssets, true
}

// HasNotUsedInAssets returns a boolean if a field has been set.
func (o *TagIndexingRuleMetricMatch) HasNotUsedInAssets() bool {
	return o != nil && o.NotUsedInAssets != nil
}

// SetNotUsedInAssets gets a reference to the given bool and assigns it to the NotUsedInAssets field.
func (o *TagIndexingRuleMetricMatch) SetNotUsedInAssets(v bool) {
	o.NotUsedInAssets = &v
}

// GetQueriedWindowSeconds returns the QueriedWindowSeconds field value if set, zero value otherwise.
func (o *TagIndexingRuleMetricMatch) GetQueriedWindowSeconds() int64 {
	if o == nil || o.QueriedWindowSeconds == nil {
		var ret int64
		return ret
	}
	return *o.QueriedWindowSeconds
}

// GetQueriedWindowSecondsOk returns a tuple with the QueriedWindowSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleMetricMatch) GetQueriedWindowSecondsOk() (*int64, bool) {
	if o == nil || o.QueriedWindowSeconds == nil {
		return nil, false
	}
	return o.QueriedWindowSeconds, true
}

// HasQueriedWindowSeconds returns a boolean if a field has been set.
func (o *TagIndexingRuleMetricMatch) HasQueriedWindowSeconds() bool {
	return o != nil && o.QueriedWindowSeconds != nil
}

// SetQueriedWindowSeconds gets a reference to the given int64 and assigns it to the QueriedWindowSeconds field.
func (o *TagIndexingRuleMetricMatch) SetQueriedWindowSeconds(v int64) {
	o.QueriedWindowSeconds = &v
}

// GetUsedInAssets returns the UsedInAssets field value if set, zero value otherwise.
func (o *TagIndexingRuleMetricMatch) GetUsedInAssets() bool {
	if o == nil || o.UsedInAssets == nil {
		var ret bool
		return ret
	}
	return *o.UsedInAssets
}

// GetUsedInAssetsOk returns a tuple with the UsedInAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleMetricMatch) GetUsedInAssetsOk() (*bool, bool) {
	if o == nil || o.UsedInAssets == nil {
		return nil, false
	}
	return o.UsedInAssets, true
}

// HasUsedInAssets returns a boolean if a field has been set.
func (o *TagIndexingRuleMetricMatch) HasUsedInAssets() bool {
	return o != nil && o.UsedInAssets != nil
}

// SetUsedInAssets gets a reference to the given bool and assigns it to the UsedInAssets field.
func (o *TagIndexingRuleMetricMatch) SetUsedInAssets(v bool) {
	o.UsedInAssets = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TagIndexingRuleMetricMatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.IsQueried != nil {
		toSerialize["is_queried"] = o.IsQueried
	}
	if o.NotQueried != nil {
		toSerialize["not_queried"] = o.NotQueried
	}
	if o.NotUsedInAssets != nil {
		toSerialize["not_used_in_assets"] = o.NotUsedInAssets
	}
	if o.QueriedWindowSeconds != nil {
		toSerialize["queried_window_seconds"] = o.QueriedWindowSeconds
	}
	if o.UsedInAssets != nil {
		toSerialize["used_in_assets"] = o.UsedInAssets
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TagIndexingRuleMetricMatch) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IsQueried            *bool  `json:"is_queried,omitempty"`
		NotQueried           *bool  `json:"not_queried,omitempty"`
		NotUsedInAssets      *bool  `json:"not_used_in_assets,omitempty"`
		QueriedWindowSeconds *int64 `json:"queried_window_seconds,omitempty"`
		UsedInAssets         *bool  `json:"used_in_assets,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"is_queried", "not_queried", "not_used_in_assets", "queried_window_seconds", "used_in_assets"})
	} else {
		return err
	}
	o.IsQueried = all.IsQueried
	o.NotQueried = all.NotQueried
	o.NotUsedInAssets = all.NotUsedInAssets
	o.QueriedWindowSeconds = all.QueriedWindowSeconds
	o.UsedInAssets = all.UsedInAssets

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
