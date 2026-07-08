// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagIndexingRuleAttributes Attributes of a tag indexing rule.
type TagIndexingRuleAttributes struct {
	// Timestamp when the rule was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Handle of the user who created the rule.
	CreatedByHandle *string `json:"created_by_handle,omitempty"`
	// When true, the rule excludes the listed tags and indexes all others. When false (default), the rule includes only the listed tags.
	ExcludeTagsMode *bool `json:"exclude_tags_mode,omitempty"`
	// Metric name prefixes excluded from the rule's scope.
	IgnoredMetricNameMatches []string `json:"ignored_metric_name_matches,omitempty"`
	// Metric name prefixes (glob patterns) this rule applies to.
	MetricNameMatches []string `json:"metric_name_matches,omitempty"`
	// Timestamp when the rule was last modified.
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// Handle of the user who last modified the rule.
	ModifiedByHandle *string `json:"modified_by_handle,omitempty"`
	// Human-readable name for the rule.
	Name *string `json:"name,omitempty"`
	// Versioned configuration options for a tag indexing rule.
	Options *TagIndexingRuleOptions `json:"options,omitempty"`
	// Evaluation order within the org. Lower values are evaluated first. Assigned server-side on create (max+1); pass on update to change the rule's position.
	RuleOrder *int64 `json:"rule_order,omitempty"`
	// Tag keys managed by this rule.
	Tags []string `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTagIndexingRuleAttributes instantiates a new TagIndexingRuleAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTagIndexingRuleAttributes() *TagIndexingRuleAttributes {
	this := TagIndexingRuleAttributes{}
	return &this
}

// NewTagIndexingRuleAttributesWithDefaults instantiates a new TagIndexingRuleAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTagIndexingRuleAttributesWithDefaults() *TagIndexingRuleAttributes {
	this := TagIndexingRuleAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TagIndexingRuleAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TagIndexingRuleAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *TagIndexingRuleAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedByHandle returns the CreatedByHandle field value if set, zero value otherwise.
func (o *TagIndexingRuleAttributes) GetCreatedByHandle() string {
	if o == nil || o.CreatedByHandle == nil {
		var ret string
		return ret
	}
	return *o.CreatedByHandle
}

// GetCreatedByHandleOk returns a tuple with the CreatedByHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleAttributes) GetCreatedByHandleOk() (*string, bool) {
	if o == nil || o.CreatedByHandle == nil {
		return nil, false
	}
	return o.CreatedByHandle, true
}

// HasCreatedByHandle returns a boolean if a field has been set.
func (o *TagIndexingRuleAttributes) HasCreatedByHandle() bool {
	return o != nil && o.CreatedByHandle != nil
}

// SetCreatedByHandle gets a reference to the given string and assigns it to the CreatedByHandle field.
func (o *TagIndexingRuleAttributes) SetCreatedByHandle(v string) {
	o.CreatedByHandle = &v
}

// GetExcludeTagsMode returns the ExcludeTagsMode field value if set, zero value otherwise.
func (o *TagIndexingRuleAttributes) GetExcludeTagsMode() bool {
	if o == nil || o.ExcludeTagsMode == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeTagsMode
}

// GetExcludeTagsModeOk returns a tuple with the ExcludeTagsMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleAttributes) GetExcludeTagsModeOk() (*bool, bool) {
	if o == nil || o.ExcludeTagsMode == nil {
		return nil, false
	}
	return o.ExcludeTagsMode, true
}

// HasExcludeTagsMode returns a boolean if a field has been set.
func (o *TagIndexingRuleAttributes) HasExcludeTagsMode() bool {
	return o != nil && o.ExcludeTagsMode != nil
}

// SetExcludeTagsMode gets a reference to the given bool and assigns it to the ExcludeTagsMode field.
func (o *TagIndexingRuleAttributes) SetExcludeTagsMode(v bool) {
	o.ExcludeTagsMode = &v
}

// GetIgnoredMetricNameMatches returns the IgnoredMetricNameMatches field value if set, zero value otherwise.
func (o *TagIndexingRuleAttributes) GetIgnoredMetricNameMatches() []string {
	if o == nil || o.IgnoredMetricNameMatches == nil {
		var ret []string
		return ret
	}
	return o.IgnoredMetricNameMatches
}

// GetIgnoredMetricNameMatchesOk returns a tuple with the IgnoredMetricNameMatches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleAttributes) GetIgnoredMetricNameMatchesOk() (*[]string, bool) {
	if o == nil || o.IgnoredMetricNameMatches == nil {
		return nil, false
	}
	return &o.IgnoredMetricNameMatches, true
}

// HasIgnoredMetricNameMatches returns a boolean if a field has been set.
func (o *TagIndexingRuleAttributes) HasIgnoredMetricNameMatches() bool {
	return o != nil && o.IgnoredMetricNameMatches != nil
}

// SetIgnoredMetricNameMatches gets a reference to the given []string and assigns it to the IgnoredMetricNameMatches field.
func (o *TagIndexingRuleAttributes) SetIgnoredMetricNameMatches(v []string) {
	o.IgnoredMetricNameMatches = v
}

// GetMetricNameMatches returns the MetricNameMatches field value if set, zero value otherwise.
func (o *TagIndexingRuleAttributes) GetMetricNameMatches() []string {
	if o == nil || o.MetricNameMatches == nil {
		var ret []string
		return ret
	}
	return o.MetricNameMatches
}

// GetMetricNameMatchesOk returns a tuple with the MetricNameMatches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleAttributes) GetMetricNameMatchesOk() (*[]string, bool) {
	if o == nil || o.MetricNameMatches == nil {
		return nil, false
	}
	return &o.MetricNameMatches, true
}

// HasMetricNameMatches returns a boolean if a field has been set.
func (o *TagIndexingRuleAttributes) HasMetricNameMatches() bool {
	return o != nil && o.MetricNameMatches != nil
}

// SetMetricNameMatches gets a reference to the given []string and assigns it to the MetricNameMatches field.
func (o *TagIndexingRuleAttributes) SetMetricNameMatches(v []string) {
	o.MetricNameMatches = v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *TagIndexingRuleAttributes) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *TagIndexingRuleAttributes) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt != nil
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *TagIndexingRuleAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetModifiedByHandle returns the ModifiedByHandle field value if set, zero value otherwise.
func (o *TagIndexingRuleAttributes) GetModifiedByHandle() string {
	if o == nil || o.ModifiedByHandle == nil {
		var ret string
		return ret
	}
	return *o.ModifiedByHandle
}

// GetModifiedByHandleOk returns a tuple with the ModifiedByHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleAttributes) GetModifiedByHandleOk() (*string, bool) {
	if o == nil || o.ModifiedByHandle == nil {
		return nil, false
	}
	return o.ModifiedByHandle, true
}

// HasModifiedByHandle returns a boolean if a field has been set.
func (o *TagIndexingRuleAttributes) HasModifiedByHandle() bool {
	return o != nil && o.ModifiedByHandle != nil
}

// SetModifiedByHandle gets a reference to the given string and assigns it to the ModifiedByHandle field.
func (o *TagIndexingRuleAttributes) SetModifiedByHandle(v string) {
	o.ModifiedByHandle = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TagIndexingRuleAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TagIndexingRuleAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TagIndexingRuleAttributes) SetName(v string) {
	o.Name = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *TagIndexingRuleAttributes) GetOptions() TagIndexingRuleOptions {
	if o == nil || o.Options == nil {
		var ret TagIndexingRuleOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleAttributes) GetOptionsOk() (*TagIndexingRuleOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *TagIndexingRuleAttributes) HasOptions() bool {
	return o != nil && o.Options != nil
}

// SetOptions gets a reference to the given TagIndexingRuleOptions and assigns it to the Options field.
func (o *TagIndexingRuleAttributes) SetOptions(v TagIndexingRuleOptions) {
	o.Options = &v
}

// GetRuleOrder returns the RuleOrder field value if set, zero value otherwise.
func (o *TagIndexingRuleAttributes) GetRuleOrder() int64 {
	if o == nil || o.RuleOrder == nil {
		var ret int64
		return ret
	}
	return *o.RuleOrder
}

// GetRuleOrderOk returns a tuple with the RuleOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleAttributes) GetRuleOrderOk() (*int64, bool) {
	if o == nil || o.RuleOrder == nil {
		return nil, false
	}
	return o.RuleOrder, true
}

// HasRuleOrder returns a boolean if a field has been set.
func (o *TagIndexingRuleAttributes) HasRuleOrder() bool {
	return o != nil && o.RuleOrder != nil
}

// SetRuleOrder gets a reference to the given int64 and assigns it to the RuleOrder field.
func (o *TagIndexingRuleAttributes) SetRuleOrder(v int64) {
	o.RuleOrder = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *TagIndexingRuleAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *TagIndexingRuleAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *TagIndexingRuleAttributes) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TagIndexingRuleAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.CreatedByHandle != nil {
		toSerialize["created_by_handle"] = o.CreatedByHandle
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
	if o.ModifiedAt != nil {
		if o.ModifiedAt.Nanosecond() == 0 {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.ModifiedByHandle != nil {
		toSerialize["modified_by_handle"] = o.ModifiedByHandle
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
func (o *TagIndexingRuleAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt                *time.Time              `json:"created_at,omitempty"`
		CreatedByHandle          *string                 `json:"created_by_handle,omitempty"`
		ExcludeTagsMode          *bool                   `json:"exclude_tags_mode,omitempty"`
		IgnoredMetricNameMatches []string                `json:"ignored_metric_name_matches,omitempty"`
		MetricNameMatches        []string                `json:"metric_name_matches,omitempty"`
		ModifiedAt               *time.Time              `json:"modified_at,omitempty"`
		ModifiedByHandle         *string                 `json:"modified_by_handle,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "created_by_handle", "exclude_tags_mode", "ignored_metric_name_matches", "metric_name_matches", "modified_at", "modified_by_handle", "name", "options", "rule_order", "tags"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = all.CreatedAt
	o.CreatedByHandle = all.CreatedByHandle
	o.ExcludeTagsMode = all.ExcludeTagsMode
	o.IgnoredMetricNameMatches = all.IgnoredMetricNameMatches
	o.MetricNameMatches = all.MetricNameMatches
	o.ModifiedAt = all.ModifiedAt
	o.ModifiedByHandle = all.ModifiedByHandle
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
