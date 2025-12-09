// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SplitConfigSortCompute Defines the metric and aggregation used as the sort value.
type SplitConfigSortCompute struct {
	// How to aggregate the sort metric for the purposes of ordering.
	Aggregation string `json:"aggregation"`
	// The metric to use for sorting graphs.
	Metric string `json:"metric"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSplitConfigSortCompute instantiates a new SplitConfigSortCompute object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSplitConfigSortCompute(aggregation string, metric string) *SplitConfigSortCompute {
	this := SplitConfigSortCompute{}
	this.Aggregation = aggregation
	this.Metric = metric
	return &this
}

// NewSplitConfigSortComputeWithDefaults instantiates a new SplitConfigSortCompute object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSplitConfigSortComputeWithDefaults() *SplitConfigSortCompute {
	this := SplitConfigSortCompute{}
	return &this
}

// GetAggregation returns the Aggregation field value.
func (o *SplitConfigSortCompute) GetAggregation() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value
// and a boolean to check if the value has been set.
func (o *SplitConfigSortCompute) GetAggregationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregation, true
}

// SetAggregation sets field value.
func (o *SplitConfigSortCompute) SetAggregation(v string) {
	o.Aggregation = v
}

// GetMetric returns the Metric field value.
func (o *SplitConfigSortCompute) GetMetric() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Metric
}

// GetMetricOk returns a tuple with the Metric field value
// and a boolean to check if the value has been set.
func (o *SplitConfigSortCompute) GetMetricOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metric, true
}

// SetMetric sets field value.
func (o *SplitConfigSortCompute) SetMetric(v string) {
	o.Metric = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SplitConfigSortCompute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["aggregation"] = o.Aggregation
	toSerialize["metric"] = o.Metric

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SplitConfigSortCompute) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregation *string `json:"aggregation"`
		Metric      *string `json:"metric"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Aggregation == nil {
		return fmt.Errorf("required field aggregation missing")
	}
	if all.Metric == nil {
		return fmt.Errorf("required field metric missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aggregation", "metric"})
	} else {
		return err
	}
	o.Aggregation = *all.Aggregation
	o.Metric = *all.Metric

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
