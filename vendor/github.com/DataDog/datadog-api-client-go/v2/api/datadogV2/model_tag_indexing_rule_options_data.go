// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagIndexingRuleOptionsData Data payload for tag indexing rule options.
type TagIndexingRuleOptionsData struct {
	// Configuration for including dynamically queried tags.
	DynamicTags *TagIndexingRuleDynamicTags `json:"dynamic_tags,omitempty"`
	// When true, the rule applies to metrics that were ingested before the rule was created.
	ManagePreexistingMetrics *bool `json:"manage_preexisting_metrics,omitempty"`
	// Criteria for matching metrics based on query state.
	MetricMatch *TagIndexingRuleMetricMatch `json:"metric_match,omitempty"`
	// When true, this rule's tag list overrides tags configured by earlier rules for the same metric. When false (default), tags from all matching rules are combined.
	OverridePreviousRules *bool `json:"override_previous_rules,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTagIndexingRuleOptionsData instantiates a new TagIndexingRuleOptionsData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTagIndexingRuleOptionsData() *TagIndexingRuleOptionsData {
	this := TagIndexingRuleOptionsData{}
	return &this
}

// NewTagIndexingRuleOptionsDataWithDefaults instantiates a new TagIndexingRuleOptionsData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTagIndexingRuleOptionsDataWithDefaults() *TagIndexingRuleOptionsData {
	this := TagIndexingRuleOptionsData{}
	return &this
}

// GetDynamicTags returns the DynamicTags field value if set, zero value otherwise.
func (o *TagIndexingRuleOptionsData) GetDynamicTags() TagIndexingRuleDynamicTags {
	if o == nil || o.DynamicTags == nil {
		var ret TagIndexingRuleDynamicTags
		return ret
	}
	return *o.DynamicTags
}

// GetDynamicTagsOk returns a tuple with the DynamicTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleOptionsData) GetDynamicTagsOk() (*TagIndexingRuleDynamicTags, bool) {
	if o == nil || o.DynamicTags == nil {
		return nil, false
	}
	return o.DynamicTags, true
}

// HasDynamicTags returns a boolean if a field has been set.
func (o *TagIndexingRuleOptionsData) HasDynamicTags() bool {
	return o != nil && o.DynamicTags != nil
}

// SetDynamicTags gets a reference to the given TagIndexingRuleDynamicTags and assigns it to the DynamicTags field.
func (o *TagIndexingRuleOptionsData) SetDynamicTags(v TagIndexingRuleDynamicTags) {
	o.DynamicTags = &v
}

// GetManagePreexistingMetrics returns the ManagePreexistingMetrics field value if set, zero value otherwise.
func (o *TagIndexingRuleOptionsData) GetManagePreexistingMetrics() bool {
	if o == nil || o.ManagePreexistingMetrics == nil {
		var ret bool
		return ret
	}
	return *o.ManagePreexistingMetrics
}

// GetManagePreexistingMetricsOk returns a tuple with the ManagePreexistingMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleOptionsData) GetManagePreexistingMetricsOk() (*bool, bool) {
	if o == nil || o.ManagePreexistingMetrics == nil {
		return nil, false
	}
	return o.ManagePreexistingMetrics, true
}

// HasManagePreexistingMetrics returns a boolean if a field has been set.
func (o *TagIndexingRuleOptionsData) HasManagePreexistingMetrics() bool {
	return o != nil && o.ManagePreexistingMetrics != nil
}

// SetManagePreexistingMetrics gets a reference to the given bool and assigns it to the ManagePreexistingMetrics field.
func (o *TagIndexingRuleOptionsData) SetManagePreexistingMetrics(v bool) {
	o.ManagePreexistingMetrics = &v
}

// GetMetricMatch returns the MetricMatch field value if set, zero value otherwise.
func (o *TagIndexingRuleOptionsData) GetMetricMatch() TagIndexingRuleMetricMatch {
	if o == nil || o.MetricMatch == nil {
		var ret TagIndexingRuleMetricMatch
		return ret
	}
	return *o.MetricMatch
}

// GetMetricMatchOk returns a tuple with the MetricMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleOptionsData) GetMetricMatchOk() (*TagIndexingRuleMetricMatch, bool) {
	if o == nil || o.MetricMatch == nil {
		return nil, false
	}
	return o.MetricMatch, true
}

// HasMetricMatch returns a boolean if a field has been set.
func (o *TagIndexingRuleOptionsData) HasMetricMatch() bool {
	return o != nil && o.MetricMatch != nil
}

// SetMetricMatch gets a reference to the given TagIndexingRuleMetricMatch and assigns it to the MetricMatch field.
func (o *TagIndexingRuleOptionsData) SetMetricMatch(v TagIndexingRuleMetricMatch) {
	o.MetricMatch = &v
}

// GetOverridePreviousRules returns the OverridePreviousRules field value if set, zero value otherwise.
func (o *TagIndexingRuleOptionsData) GetOverridePreviousRules() bool {
	if o == nil || o.OverridePreviousRules == nil {
		var ret bool
		return ret
	}
	return *o.OverridePreviousRules
}

// GetOverridePreviousRulesOk returns a tuple with the OverridePreviousRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleOptionsData) GetOverridePreviousRulesOk() (*bool, bool) {
	if o == nil || o.OverridePreviousRules == nil {
		return nil, false
	}
	return o.OverridePreviousRules, true
}

// HasOverridePreviousRules returns a boolean if a field has been set.
func (o *TagIndexingRuleOptionsData) HasOverridePreviousRules() bool {
	return o != nil && o.OverridePreviousRules != nil
}

// SetOverridePreviousRules gets a reference to the given bool and assigns it to the OverridePreviousRules field.
func (o *TagIndexingRuleOptionsData) SetOverridePreviousRules(v bool) {
	o.OverridePreviousRules = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TagIndexingRuleOptionsData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DynamicTags != nil {
		toSerialize["dynamic_tags"] = o.DynamicTags
	}
	if o.ManagePreexistingMetrics != nil {
		toSerialize["manage_preexisting_metrics"] = o.ManagePreexistingMetrics
	}
	if o.MetricMatch != nil {
		toSerialize["metric_match"] = o.MetricMatch
	}
	if o.OverridePreviousRules != nil {
		toSerialize["override_previous_rules"] = o.OverridePreviousRules
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TagIndexingRuleOptionsData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DynamicTags              *TagIndexingRuleDynamicTags `json:"dynamic_tags,omitempty"`
		ManagePreexistingMetrics *bool                       `json:"manage_preexisting_metrics,omitempty"`
		MetricMatch              *TagIndexingRuleMetricMatch `json:"metric_match,omitempty"`
		OverridePreviousRules    *bool                       `json:"override_previous_rules,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"dynamic_tags", "manage_preexisting_metrics", "metric_match", "override_previous_rules"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.DynamicTags != nil && all.DynamicTags.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DynamicTags = all.DynamicTags
	o.ManagePreexistingMetrics = all.ManagePreexistingMetrics
	if all.MetricMatch != nil && all.MetricMatch.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.MetricMatch = all.MetricMatch
	o.OverridePreviousRules = all.OverridePreviousRules

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
