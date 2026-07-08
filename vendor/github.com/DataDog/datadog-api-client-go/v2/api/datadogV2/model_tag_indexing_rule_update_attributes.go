// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagIndexingRuleUpdateAttributes Attributes for updating a tag indexing rule. All fields are optional; omitted fields are unchanged.
type TagIndexingRuleUpdateAttributes struct {
	// When true, the rule excludes the listed tags and indexes all others.
	ExcludeTagsMode *bool `json:"exclude_tags_mode,omitempty"`
	// Metric name prefixes excluded from the rule's scope.
	IgnoredMetricNameMatches []string `json:"ignored_metric_name_matches,omitempty"`
	// Metric name prefixes (glob patterns) this rule applies to.
	MetricNameMatches []string `json:"metric_name_matches,omitempty"`
	// Human-readable name for the rule.
	Name *string `json:"name,omitempty"`
	// Versioned configuration options for a tag indexing rule.
	Options *TagIndexingRuleOptions `json:"options,omitempty"`
	// Desired evaluation order. Returns 409 if the value conflicts with another rule; use POST /api/v2/metrics/tag-indexing-rules/order for atomic re-sequencing.
	RuleOrder *int64 `json:"rule_order,omitempty"`
	// Tag keys managed by this rule.
	Tags []string `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTagIndexingRuleUpdateAttributes instantiates a new TagIndexingRuleUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTagIndexingRuleUpdateAttributes() *TagIndexingRuleUpdateAttributes {
	this := TagIndexingRuleUpdateAttributes{}
	return &this
}

// NewTagIndexingRuleUpdateAttributesWithDefaults instantiates a new TagIndexingRuleUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTagIndexingRuleUpdateAttributesWithDefaults() *TagIndexingRuleUpdateAttributes {
	this := TagIndexingRuleUpdateAttributes{}
	return &this
}

// GetExcludeTagsMode returns the ExcludeTagsMode field value if set, zero value otherwise.
func (o *TagIndexingRuleUpdateAttributes) GetExcludeTagsMode() bool {
	if o == nil || o.ExcludeTagsMode == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeTagsMode
}

// GetExcludeTagsModeOk returns a tuple with the ExcludeTagsMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleUpdateAttributes) GetExcludeTagsModeOk() (*bool, bool) {
	if o == nil || o.ExcludeTagsMode == nil {
		return nil, false
	}
	return o.ExcludeTagsMode, true
}

// HasExcludeTagsMode returns a boolean if a field has been set.
func (o *TagIndexingRuleUpdateAttributes) HasExcludeTagsMode() bool {
	return o != nil && o.ExcludeTagsMode != nil
}

// SetExcludeTagsMode gets a reference to the given bool and assigns it to the ExcludeTagsMode field.
func (o *TagIndexingRuleUpdateAttributes) SetExcludeTagsMode(v bool) {
	o.ExcludeTagsMode = &v
}

// GetIgnoredMetricNameMatches returns the IgnoredMetricNameMatches field value if set, zero value otherwise.
func (o *TagIndexingRuleUpdateAttributes) GetIgnoredMetricNameMatches() []string {
	if o == nil || o.IgnoredMetricNameMatches == nil {
		var ret []string
		return ret
	}
	return o.IgnoredMetricNameMatches
}

// GetIgnoredMetricNameMatchesOk returns a tuple with the IgnoredMetricNameMatches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleUpdateAttributes) GetIgnoredMetricNameMatchesOk() (*[]string, bool) {
	if o == nil || o.IgnoredMetricNameMatches == nil {
		return nil, false
	}
	return &o.IgnoredMetricNameMatches, true
}

// HasIgnoredMetricNameMatches returns a boolean if a field has been set.
func (o *TagIndexingRuleUpdateAttributes) HasIgnoredMetricNameMatches() bool {
	return o != nil && o.IgnoredMetricNameMatches != nil
}

// SetIgnoredMetricNameMatches gets a reference to the given []string and assigns it to the IgnoredMetricNameMatches field.
func (o *TagIndexingRuleUpdateAttributes) SetIgnoredMetricNameMatches(v []string) {
	o.IgnoredMetricNameMatches = v
}

// GetMetricNameMatches returns the MetricNameMatches field value if set, zero value otherwise.
func (o *TagIndexingRuleUpdateAttributes) GetMetricNameMatches() []string {
	if o == nil || o.MetricNameMatches == nil {
		var ret []string
		return ret
	}
	return o.MetricNameMatches
}

// GetMetricNameMatchesOk returns a tuple with the MetricNameMatches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleUpdateAttributes) GetMetricNameMatchesOk() (*[]string, bool) {
	if o == nil || o.MetricNameMatches == nil {
		return nil, false
	}
	return &o.MetricNameMatches, true
}

// HasMetricNameMatches returns a boolean if a field has been set.
func (o *TagIndexingRuleUpdateAttributes) HasMetricNameMatches() bool {
	return o != nil && o.MetricNameMatches != nil
}

// SetMetricNameMatches gets a reference to the given []string and assigns it to the MetricNameMatches field.
func (o *TagIndexingRuleUpdateAttributes) SetMetricNameMatches(v []string) {
	o.MetricNameMatches = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TagIndexingRuleUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TagIndexingRuleUpdateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TagIndexingRuleUpdateAttributes) SetName(v string) {
	o.Name = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *TagIndexingRuleUpdateAttributes) GetOptions() TagIndexingRuleOptions {
	if o == nil || o.Options == nil {
		var ret TagIndexingRuleOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleUpdateAttributes) GetOptionsOk() (*TagIndexingRuleOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *TagIndexingRuleUpdateAttributes) HasOptions() bool {
	return o != nil && o.Options != nil
}

// SetOptions gets a reference to the given TagIndexingRuleOptions and assigns it to the Options field.
func (o *TagIndexingRuleUpdateAttributes) SetOptions(v TagIndexingRuleOptions) {
	o.Options = &v
}

// GetRuleOrder returns the RuleOrder field value if set, zero value otherwise.
func (o *TagIndexingRuleUpdateAttributes) GetRuleOrder() int64 {
	if o == nil || o.RuleOrder == nil {
		var ret int64
		return ret
	}
	return *o.RuleOrder
}

// GetRuleOrderOk returns a tuple with the RuleOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleUpdateAttributes) GetRuleOrderOk() (*int64, bool) {
	if o == nil || o.RuleOrder == nil {
		return nil, false
	}
	return o.RuleOrder, true
}

// HasRuleOrder returns a boolean if a field has been set.
func (o *TagIndexingRuleUpdateAttributes) HasRuleOrder() bool {
	return o != nil && o.RuleOrder != nil
}

// SetRuleOrder gets a reference to the given int64 and assigns it to the RuleOrder field.
func (o *TagIndexingRuleUpdateAttributes) SetRuleOrder(v int64) {
	o.RuleOrder = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *TagIndexingRuleUpdateAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleUpdateAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *TagIndexingRuleUpdateAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *TagIndexingRuleUpdateAttributes) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TagIndexingRuleUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ExcludeTagsMode != nil {
		toSerialize["exclude_tags_mode"] = o.ExcludeTagsMode
	}
	if o.IgnoredMetricNameMatches != nil {
		toSerialize["ignored_metric_name_matches"] = o.IgnoredMetricNameMatches
	}
	if o.MetricNameMatches != nil {
		toSerialize["metric_name_matches"] = o.MetricNameMatches
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.RuleOrder != nil {
		toSerialize["rule_order"] = o.RuleOrder
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TagIndexingRuleUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ExcludeTagsMode          *bool                   `json:"exclude_tags_mode,omitempty"`
		IgnoredMetricNameMatches []string                `json:"ignored_metric_name_matches,omitempty"`
		MetricNameMatches        []string                `json:"metric_name_matches,omitempty"`
		Name                     *string                 `json:"name,omitempty"`
		Options                  *TagIndexingRuleOptions `json:"options,omitempty"`
		RuleOrder                *int64                  `json:"rule_order,omitempty"`
		Tags                     []string                `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"exclude_tags_mode", "ignored_metric_name_matches", "metric_name_matches", "name", "options", "rule_order", "tags"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ExcludeTagsMode = all.ExcludeTagsMode
	o.IgnoredMetricNameMatches = all.IgnoredMetricNameMatches
	o.MetricNameMatches = all.MetricNameMatches
	o.Name = all.Name
	if all.Options != nil && all.Options.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Options = all.Options
	o.RuleOrder = all.RuleOrder
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
