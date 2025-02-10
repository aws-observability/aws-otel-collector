// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSMetricsConfig AWS Metrics Collection config.
type AWSMetricsConfig struct {
	// Enable EC2 automute for AWS metrics. Defaults to `true`.
	AutomuteEnabled *bool `json:"automute_enabled,omitempty"`
	// Enable CloudWatch alarms collection. Defaults to `false`.
	CollectCloudwatchAlarms *bool `json:"collect_cloudwatch_alarms,omitempty"`
	// Enable custom metrics collection. Defaults to `false`.
	CollectCustomMetrics *bool `json:"collect_custom_metrics,omitempty"`
	// Enable AWS metrics collection. Defaults to `true`.
	Enabled *bool `json:"enabled,omitempty"`
	// AWS Metrics namespace filters. Defaults to `exclude_only`.
	NamespaceFilters *AWSNamespaceFilters `json:"namespace_filters,omitempty"`
	// AWS Metrics collection tag filters list. Defaults to `[]`.
	TagFilters []AWSNamespaceTagFilter `json:"tag_filters,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSMetricsConfig instantiates a new AWSMetricsConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSMetricsConfig() *AWSMetricsConfig {
	this := AWSMetricsConfig{}
	return &this
}

// NewAWSMetricsConfigWithDefaults instantiates a new AWSMetricsConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSMetricsConfigWithDefaults() *AWSMetricsConfig {
	this := AWSMetricsConfig{}
	return &this
}

// GetAutomuteEnabled returns the AutomuteEnabled field value if set, zero value otherwise.
func (o *AWSMetricsConfig) GetAutomuteEnabled() bool {
	if o == nil || o.AutomuteEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AutomuteEnabled
}

// GetAutomuteEnabledOk returns a tuple with the AutomuteEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSMetricsConfig) GetAutomuteEnabledOk() (*bool, bool) {
	if o == nil || o.AutomuteEnabled == nil {
		return nil, false
	}
	return o.AutomuteEnabled, true
}

// HasAutomuteEnabled returns a boolean if a field has been set.
func (o *AWSMetricsConfig) HasAutomuteEnabled() bool {
	return o != nil && o.AutomuteEnabled != nil
}

// SetAutomuteEnabled gets a reference to the given bool and assigns it to the AutomuteEnabled field.
func (o *AWSMetricsConfig) SetAutomuteEnabled(v bool) {
	o.AutomuteEnabled = &v
}

// GetCollectCloudwatchAlarms returns the CollectCloudwatchAlarms field value if set, zero value otherwise.
func (o *AWSMetricsConfig) GetCollectCloudwatchAlarms() bool {
	if o == nil || o.CollectCloudwatchAlarms == nil {
		var ret bool
		return ret
	}
	return *o.CollectCloudwatchAlarms
}

// GetCollectCloudwatchAlarmsOk returns a tuple with the CollectCloudwatchAlarms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSMetricsConfig) GetCollectCloudwatchAlarmsOk() (*bool, bool) {
	if o == nil || o.CollectCloudwatchAlarms == nil {
		return nil, false
	}
	return o.CollectCloudwatchAlarms, true
}

// HasCollectCloudwatchAlarms returns a boolean if a field has been set.
func (o *AWSMetricsConfig) HasCollectCloudwatchAlarms() bool {
	return o != nil && o.CollectCloudwatchAlarms != nil
}

// SetCollectCloudwatchAlarms gets a reference to the given bool and assigns it to the CollectCloudwatchAlarms field.
func (o *AWSMetricsConfig) SetCollectCloudwatchAlarms(v bool) {
	o.CollectCloudwatchAlarms = &v
}

// GetCollectCustomMetrics returns the CollectCustomMetrics field value if set, zero value otherwise.
func (o *AWSMetricsConfig) GetCollectCustomMetrics() bool {
	if o == nil || o.CollectCustomMetrics == nil {
		var ret bool
		return ret
	}
	return *o.CollectCustomMetrics
}

// GetCollectCustomMetricsOk returns a tuple with the CollectCustomMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSMetricsConfig) GetCollectCustomMetricsOk() (*bool, bool) {
	if o == nil || o.CollectCustomMetrics == nil {
		return nil, false
	}
	return o.CollectCustomMetrics, true
}

// HasCollectCustomMetrics returns a boolean if a field has been set.
func (o *AWSMetricsConfig) HasCollectCustomMetrics() bool {
	return o != nil && o.CollectCustomMetrics != nil
}

// SetCollectCustomMetrics gets a reference to the given bool and assigns it to the CollectCustomMetrics field.
func (o *AWSMetricsConfig) SetCollectCustomMetrics(v bool) {
	o.CollectCustomMetrics = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AWSMetricsConfig) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSMetricsConfig) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AWSMetricsConfig) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AWSMetricsConfig) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetNamespaceFilters returns the NamespaceFilters field value if set, zero value otherwise.
func (o *AWSMetricsConfig) GetNamespaceFilters() AWSNamespaceFilters {
	if o == nil || o.NamespaceFilters == nil {
		var ret AWSNamespaceFilters
		return ret
	}
	return *o.NamespaceFilters
}

// GetNamespaceFiltersOk returns a tuple with the NamespaceFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSMetricsConfig) GetNamespaceFiltersOk() (*AWSNamespaceFilters, bool) {
	if o == nil || o.NamespaceFilters == nil {
		return nil, false
	}
	return o.NamespaceFilters, true
}

// HasNamespaceFilters returns a boolean if a field has been set.
func (o *AWSMetricsConfig) HasNamespaceFilters() bool {
	return o != nil && o.NamespaceFilters != nil
}

// SetNamespaceFilters gets a reference to the given AWSNamespaceFilters and assigns it to the NamespaceFilters field.
func (o *AWSMetricsConfig) SetNamespaceFilters(v AWSNamespaceFilters) {
	o.NamespaceFilters = &v
}

// GetTagFilters returns the TagFilters field value if set, zero value otherwise.
func (o *AWSMetricsConfig) GetTagFilters() []AWSNamespaceTagFilter {
	if o == nil || o.TagFilters == nil {
		var ret []AWSNamespaceTagFilter
		return ret
	}
	return o.TagFilters
}

// GetTagFiltersOk returns a tuple with the TagFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSMetricsConfig) GetTagFiltersOk() (*[]AWSNamespaceTagFilter, bool) {
	if o == nil || o.TagFilters == nil {
		return nil, false
	}
	return &o.TagFilters, true
}

// HasTagFilters returns a boolean if a field has been set.
func (o *AWSMetricsConfig) HasTagFilters() bool {
	return o != nil && o.TagFilters != nil
}

// SetTagFilters gets a reference to the given []AWSNamespaceTagFilter and assigns it to the TagFilters field.
func (o *AWSMetricsConfig) SetTagFilters(v []AWSNamespaceTagFilter) {
	o.TagFilters = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSMetricsConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AutomuteEnabled != nil {
		toSerialize["automute_enabled"] = o.AutomuteEnabled
	}
	if o.CollectCloudwatchAlarms != nil {
		toSerialize["collect_cloudwatch_alarms"] = o.CollectCloudwatchAlarms
	}
	if o.CollectCustomMetrics != nil {
		toSerialize["collect_custom_metrics"] = o.CollectCustomMetrics
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.NamespaceFilters != nil {
		toSerialize["namespace_filters"] = o.NamespaceFilters
	}
	if o.TagFilters != nil {
		toSerialize["tag_filters"] = o.TagFilters
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSMetricsConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AutomuteEnabled         *bool                   `json:"automute_enabled,omitempty"`
		CollectCloudwatchAlarms *bool                   `json:"collect_cloudwatch_alarms,omitempty"`
		CollectCustomMetrics    *bool                   `json:"collect_custom_metrics,omitempty"`
		Enabled                 *bool                   `json:"enabled,omitempty"`
		NamespaceFilters        *AWSNamespaceFilters    `json:"namespace_filters,omitempty"`
		TagFilters              []AWSNamespaceTagFilter `json:"tag_filters,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"automute_enabled", "collect_cloudwatch_alarms", "collect_custom_metrics", "enabled", "namespace_filters", "tag_filters"})
	} else {
		return err
	}
	o.AutomuteEnabled = all.AutomuteEnabled
	o.CollectCloudwatchAlarms = all.CollectCloudwatchAlarms
	o.CollectCustomMetrics = all.CollectCustomMetrics
	o.Enabled = all.Enabled
	o.NamespaceFilters = all.NamespaceFilters
	o.TagFilters = all.TagFilters

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
