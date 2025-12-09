// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricSuggestedTagsAttributes Object containing the definition of a metric's actively queried tags and aggregations.
type MetricSuggestedTagsAttributes struct {
	// List of aggregation combinations that have been actively queried.
	ActiveAggregations []MetricCustomAggregation `json:"active_aggregations,omitempty"`
	// List of tag keys that have been actively queried.
	ActiveTags []string `json:"active_tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMetricSuggestedTagsAttributes instantiates a new MetricSuggestedTagsAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMetricSuggestedTagsAttributes() *MetricSuggestedTagsAttributes {
	this := MetricSuggestedTagsAttributes{}
	return &this
}

// NewMetricSuggestedTagsAttributesWithDefaults instantiates a new MetricSuggestedTagsAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMetricSuggestedTagsAttributesWithDefaults() *MetricSuggestedTagsAttributes {
	this := MetricSuggestedTagsAttributes{}
	return &this
}

// GetActiveAggregations returns the ActiveAggregations field value if set, zero value otherwise.
func (o *MetricSuggestedTagsAttributes) GetActiveAggregations() []MetricCustomAggregation {
	if o == nil || o.ActiveAggregations == nil {
		var ret []MetricCustomAggregation
		return ret
	}
	return o.ActiveAggregations
}

// GetActiveAggregationsOk returns a tuple with the ActiveAggregations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricSuggestedTagsAttributes) GetActiveAggregationsOk() (*[]MetricCustomAggregation, bool) {
	if o == nil || o.ActiveAggregations == nil {
		return nil, false
	}
	return &o.ActiveAggregations, true
}

// HasActiveAggregations returns a boolean if a field has been set.
func (o *MetricSuggestedTagsAttributes) HasActiveAggregations() bool {
	return o != nil && o.ActiveAggregations != nil
}

// SetActiveAggregations gets a reference to the given []MetricCustomAggregation and assigns it to the ActiveAggregations field.
func (o *MetricSuggestedTagsAttributes) SetActiveAggregations(v []MetricCustomAggregation) {
	o.ActiveAggregations = v
}

// GetActiveTags returns the ActiveTags field value if set, zero value otherwise.
func (o *MetricSuggestedTagsAttributes) GetActiveTags() []string {
	if o == nil || o.ActiveTags == nil {
		var ret []string
		return ret
	}
	return o.ActiveTags
}

// GetActiveTagsOk returns a tuple with the ActiveTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricSuggestedTagsAttributes) GetActiveTagsOk() (*[]string, bool) {
	if o == nil || o.ActiveTags == nil {
		return nil, false
	}
	return &o.ActiveTags, true
}

// HasActiveTags returns a boolean if a field has been set.
func (o *MetricSuggestedTagsAttributes) HasActiveTags() bool {
	return o != nil && o.ActiveTags != nil
}

// SetActiveTags gets a reference to the given []string and assigns it to the ActiveTags field.
func (o *MetricSuggestedTagsAttributes) SetActiveTags(v []string) {
	o.ActiveTags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MetricSuggestedTagsAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ActiveAggregations != nil {
		toSerialize["active_aggregations"] = o.ActiveAggregations
	}
	if o.ActiveTags != nil {
		toSerialize["active_tags"] = o.ActiveTags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MetricSuggestedTagsAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ActiveAggregations []MetricCustomAggregation `json:"active_aggregations,omitempty"`
		ActiveTags         []string                  `json:"active_tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"active_aggregations", "active_tags"})
	} else {
		return err
	}
	o.ActiveAggregations = all.ActiveAggregations
	o.ActiveTags = all.ActiveTags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
