// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AggregatedWaterfallPerformanceCriteria Performance criteria to filter view instances by a metric threshold.
type AggregatedWaterfallPerformanceCriteria struct {
	// Maximum threshold in seconds (inclusive).
	Max *float64 `json:"max,omitempty"`
	// Performance metric used to filter view instances by threshold.
	Metric AggregatedWaterfallPerformanceCriteriaMetric `json:"metric"`
	// Minimum threshold in seconds (inclusive).
	Min *float64 `json:"min,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAggregatedWaterfallPerformanceCriteria instantiates a new AggregatedWaterfallPerformanceCriteria object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAggregatedWaterfallPerformanceCriteria(metric AggregatedWaterfallPerformanceCriteriaMetric) *AggregatedWaterfallPerformanceCriteria {
	this := AggregatedWaterfallPerformanceCriteria{}
	this.Metric = metric
	return &this
}

// NewAggregatedWaterfallPerformanceCriteriaWithDefaults instantiates a new AggregatedWaterfallPerformanceCriteria object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAggregatedWaterfallPerformanceCriteriaWithDefaults() *AggregatedWaterfallPerformanceCriteria {
	this := AggregatedWaterfallPerformanceCriteria{}
	return &this
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *AggregatedWaterfallPerformanceCriteria) GetMax() float64 {
	if o == nil || o.Max == nil {
		var ret float64
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedWaterfallPerformanceCriteria) GetMaxOk() (*float64, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *AggregatedWaterfallPerformanceCriteria) HasMax() bool {
	return o != nil && o.Max != nil
}

// SetMax gets a reference to the given float64 and assigns it to the Max field.
func (o *AggregatedWaterfallPerformanceCriteria) SetMax(v float64) {
	o.Max = &v
}

// GetMetric returns the Metric field value.
func (o *AggregatedWaterfallPerformanceCriteria) GetMetric() AggregatedWaterfallPerformanceCriteriaMetric {
	if o == nil {
		var ret AggregatedWaterfallPerformanceCriteriaMetric
		return ret
	}
	return o.Metric
}

// GetMetricOk returns a tuple with the Metric field value
// and a boolean to check if the value has been set.
func (o *AggregatedWaterfallPerformanceCriteria) GetMetricOk() (*AggregatedWaterfallPerformanceCriteriaMetric, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metric, true
}

// SetMetric sets field value.
func (o *AggregatedWaterfallPerformanceCriteria) SetMetric(v AggregatedWaterfallPerformanceCriteriaMetric) {
	o.Metric = v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *AggregatedWaterfallPerformanceCriteria) GetMin() float64 {
	if o == nil || o.Min == nil {
		var ret float64
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedWaterfallPerformanceCriteria) GetMinOk() (*float64, bool) {
	if o == nil || o.Min == nil {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *AggregatedWaterfallPerformanceCriteria) HasMin() bool {
	return o != nil && o.Min != nil
}

// SetMin gets a reference to the given float64 and assigns it to the Min field.
func (o *AggregatedWaterfallPerformanceCriteria) SetMin(v float64) {
	o.Min = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AggregatedWaterfallPerformanceCriteria) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Max != nil {
		toSerialize["max"] = o.Max
	}
	toSerialize["metric"] = o.Metric
	if o.Min != nil {
		toSerialize["min"] = o.Min
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AggregatedWaterfallPerformanceCriteria) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Max    *float64                                      `json:"max,omitempty"`
		Metric *AggregatedWaterfallPerformanceCriteriaMetric `json:"metric"`
		Min    *float64                                      `json:"min,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Metric == nil {
		return fmt.Errorf("required field metric missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"max", "metric", "min"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Max = all.Max
	if !all.Metric.IsValid() {
		hasInvalidField = true
	} else {
		o.Metric = *all.Metric
	}
	o.Min = all.Min

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
