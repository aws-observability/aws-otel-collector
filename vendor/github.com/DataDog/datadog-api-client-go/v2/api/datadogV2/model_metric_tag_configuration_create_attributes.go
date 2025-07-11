// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricTagConfigurationCreateAttributes Object containing the definition of a metric tag configuration to be created.
type MetricTagConfigurationCreateAttributes struct {
	// Deprecated. You no longer need to configure specific time and space aggregations for Metrics Without Limits.
	Aggregations []MetricCustomAggregation `json:"aggregations,omitempty"`
	// When set to true, the configuration will exclude the configured tags and include any other submitted tags.
	// When set to false, the configuration will include the configured tags and exclude any other submitted tags.
	// Defaults to false. Requires `tags` property.
	ExcludeTagsMode *bool `json:"exclude_tags_mode,omitempty"`
	// Toggle to include/exclude percentiles for a distribution metric.
	// Defaults to false. Can only be applied to metrics that have a `metric_type` of `distribution`.
	IncludePercentiles *bool `json:"include_percentiles,omitempty"`
	// The metric's type.
	MetricType MetricTagConfigurationMetricTypes `json:"metric_type"`
	// A list of tag keys that will be queryable for your metric.
	Tags []string `json:"tags"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMetricTagConfigurationCreateAttributes instantiates a new MetricTagConfigurationCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMetricTagConfigurationCreateAttributes(metricType MetricTagConfigurationMetricTypes, tags []string) *MetricTagConfigurationCreateAttributes {
	this := MetricTagConfigurationCreateAttributes{}
	this.MetricType = metricType
	this.Tags = tags
	return &this
}

// NewMetricTagConfigurationCreateAttributesWithDefaults instantiates a new MetricTagConfigurationCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMetricTagConfigurationCreateAttributesWithDefaults() *MetricTagConfigurationCreateAttributes {
	this := MetricTagConfigurationCreateAttributes{}
	var metricType MetricTagConfigurationMetricTypes = METRICTAGCONFIGURATIONMETRICTYPES_GAUGE
	this.MetricType = metricType
	return &this
}

// GetAggregations returns the Aggregations field value if set, zero value otherwise.
func (o *MetricTagConfigurationCreateAttributes) GetAggregations() []MetricCustomAggregation {
	if o == nil || o.Aggregations == nil {
		var ret []MetricCustomAggregation
		return ret
	}
	return o.Aggregations
}

// GetAggregationsOk returns a tuple with the Aggregations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricTagConfigurationCreateAttributes) GetAggregationsOk() (*[]MetricCustomAggregation, bool) {
	if o == nil || o.Aggregations == nil {
		return nil, false
	}
	return &o.Aggregations, true
}

// HasAggregations returns a boolean if a field has been set.
func (o *MetricTagConfigurationCreateAttributes) HasAggregations() bool {
	return o != nil && o.Aggregations != nil
}

// SetAggregations gets a reference to the given []MetricCustomAggregation and assigns it to the Aggregations field.
func (o *MetricTagConfigurationCreateAttributes) SetAggregations(v []MetricCustomAggregation) {
	o.Aggregations = v
}

// GetExcludeTagsMode returns the ExcludeTagsMode field value if set, zero value otherwise.
func (o *MetricTagConfigurationCreateAttributes) GetExcludeTagsMode() bool {
	if o == nil || o.ExcludeTagsMode == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeTagsMode
}

// GetExcludeTagsModeOk returns a tuple with the ExcludeTagsMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricTagConfigurationCreateAttributes) GetExcludeTagsModeOk() (*bool, bool) {
	if o == nil || o.ExcludeTagsMode == nil {
		return nil, false
	}
	return o.ExcludeTagsMode, true
}

// HasExcludeTagsMode returns a boolean if a field has been set.
func (o *MetricTagConfigurationCreateAttributes) HasExcludeTagsMode() bool {
	return o != nil && o.ExcludeTagsMode != nil
}

// SetExcludeTagsMode gets a reference to the given bool and assigns it to the ExcludeTagsMode field.
func (o *MetricTagConfigurationCreateAttributes) SetExcludeTagsMode(v bool) {
	o.ExcludeTagsMode = &v
}

// GetIncludePercentiles returns the IncludePercentiles field value if set, zero value otherwise.
func (o *MetricTagConfigurationCreateAttributes) GetIncludePercentiles() bool {
	if o == nil || o.IncludePercentiles == nil {
		var ret bool
		return ret
	}
	return *o.IncludePercentiles
}

// GetIncludePercentilesOk returns a tuple with the IncludePercentiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricTagConfigurationCreateAttributes) GetIncludePercentilesOk() (*bool, bool) {
	if o == nil || o.IncludePercentiles == nil {
		return nil, false
	}
	return o.IncludePercentiles, true
}

// HasIncludePercentiles returns a boolean if a field has been set.
func (o *MetricTagConfigurationCreateAttributes) HasIncludePercentiles() bool {
	return o != nil && o.IncludePercentiles != nil
}

// SetIncludePercentiles gets a reference to the given bool and assigns it to the IncludePercentiles field.
func (o *MetricTagConfigurationCreateAttributes) SetIncludePercentiles(v bool) {
	o.IncludePercentiles = &v
}

// GetMetricType returns the MetricType field value.
func (o *MetricTagConfigurationCreateAttributes) GetMetricType() MetricTagConfigurationMetricTypes {
	if o == nil {
		var ret MetricTagConfigurationMetricTypes
		return ret
	}
	return o.MetricType
}

// GetMetricTypeOk returns a tuple with the MetricType field value
// and a boolean to check if the value has been set.
func (o *MetricTagConfigurationCreateAttributes) GetMetricTypeOk() (*MetricTagConfigurationMetricTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricType, true
}

// SetMetricType sets field value.
func (o *MetricTagConfigurationCreateAttributes) SetMetricType(v MetricTagConfigurationMetricTypes) {
	o.MetricType = v
}

// GetTags returns the Tags field value.
func (o *MetricTagConfigurationCreateAttributes) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *MetricTagConfigurationCreateAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value.
func (o *MetricTagConfigurationCreateAttributes) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MetricTagConfigurationCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Aggregations != nil {
		toSerialize["aggregations"] = o.Aggregations
	}
	if o.ExcludeTagsMode != nil {
		toSerialize["exclude_tags_mode"] = o.ExcludeTagsMode
	}
	if o.IncludePercentiles != nil {
		toSerialize["include_percentiles"] = o.IncludePercentiles
	}
	toSerialize["metric_type"] = o.MetricType
	toSerialize["tags"] = o.Tags

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MetricTagConfigurationCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregations       []MetricCustomAggregation          `json:"aggregations,omitempty"`
		ExcludeTagsMode    *bool                              `json:"exclude_tags_mode,omitempty"`
		IncludePercentiles *bool                              `json:"include_percentiles,omitempty"`
		MetricType         *MetricTagConfigurationMetricTypes `json:"metric_type"`
		Tags               *[]string                          `json:"tags"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.MetricType == nil {
		return fmt.Errorf("required field metric_type missing")
	}
	if all.Tags == nil {
		return fmt.Errorf("required field tags missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aggregations", "exclude_tags_mode", "include_percentiles", "metric_type", "tags"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Aggregations = all.Aggregations
	o.ExcludeTagsMode = all.ExcludeTagsMode
	o.IncludePercentiles = all.IncludePercentiles
	if !all.MetricType.IsValid() {
		hasInvalidField = true
	} else {
		o.MetricType = *all.MetricType
	}
	o.Tags = *all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
