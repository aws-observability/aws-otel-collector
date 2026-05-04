// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsFunnelCompute Compute configuration for user journey funnel.
type ProductAnalyticsFunnelCompute struct {
	// Aggregation type for user journey funnel compute.
	Aggregation ProductAnalyticsFunnelComputeAggregation `json:"aggregation"`
	// Metric for user journey funnel compute. `__dd.conversion` and `__dd.conversion_rate` accept `count` (unique users/sessions) and `cardinality` (total users/sessions) as aggregations.
	Metric ProductAnalyticsFunnelComputeMetric `json:"metric"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewProductAnalyticsFunnelCompute instantiates a new ProductAnalyticsFunnelCompute object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsFunnelCompute(aggregation ProductAnalyticsFunnelComputeAggregation, metric ProductAnalyticsFunnelComputeMetric) *ProductAnalyticsFunnelCompute {
	this := ProductAnalyticsFunnelCompute{}
	this.Aggregation = aggregation
	this.Metric = metric
	return &this
}

// NewProductAnalyticsFunnelComputeWithDefaults instantiates a new ProductAnalyticsFunnelCompute object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsFunnelComputeWithDefaults() *ProductAnalyticsFunnelCompute {
	this := ProductAnalyticsFunnelCompute{}
	return &this
}

// GetAggregation returns the Aggregation field value.
func (o *ProductAnalyticsFunnelCompute) GetAggregation() ProductAnalyticsFunnelComputeAggregation {
	if o == nil {
		var ret ProductAnalyticsFunnelComputeAggregation
		return ret
	}
	return o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFunnelCompute) GetAggregationOk() (*ProductAnalyticsFunnelComputeAggregation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregation, true
}

// SetAggregation sets field value.
func (o *ProductAnalyticsFunnelCompute) SetAggregation(v ProductAnalyticsFunnelComputeAggregation) {
	o.Aggregation = v
}

// GetMetric returns the Metric field value.
func (o *ProductAnalyticsFunnelCompute) GetMetric() ProductAnalyticsFunnelComputeMetric {
	if o == nil {
		var ret ProductAnalyticsFunnelComputeMetric
		return ret
	}
	return o.Metric
}

// GetMetricOk returns a tuple with the Metric field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFunnelCompute) GetMetricOk() (*ProductAnalyticsFunnelComputeMetric, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metric, true
}

// SetMetric sets field value.
func (o *ProductAnalyticsFunnelCompute) SetMetric(v ProductAnalyticsFunnelComputeMetric) {
	o.Metric = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsFunnelCompute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["aggregation"] = o.Aggregation
	toSerialize["metric"] = o.Metric
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsFunnelCompute) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregation *ProductAnalyticsFunnelComputeAggregation `json:"aggregation"`
		Metric      *ProductAnalyticsFunnelComputeMetric      `json:"metric"`
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

	hasInvalidField := false
	if !all.Aggregation.IsValid() {
		hasInvalidField = true
	} else {
		o.Aggregation = *all.Aggregation
	}
	if !all.Metric.IsValid() {
		hasInvalidField = true
	} else {
		o.Metric = *all.Metric
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
