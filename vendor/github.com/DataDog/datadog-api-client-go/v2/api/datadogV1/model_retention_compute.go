// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RetentionCompute Compute configuration for retention queries.
type RetentionCompute struct {
	// The type of aggregation that can be performed on events-based queries.
	Aggregation EventsAggregation `json:"aggregation"`
	// Metric for retention compute.
	Metric RetentionComputeMetric `json:"metric"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewRetentionCompute instantiates a new RetentionCompute object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRetentionCompute(aggregation EventsAggregation, metric RetentionComputeMetric) *RetentionCompute {
	this := RetentionCompute{}
	this.Aggregation = aggregation
	this.Metric = metric
	return &this
}

// NewRetentionComputeWithDefaults instantiates a new RetentionCompute object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRetentionComputeWithDefaults() *RetentionCompute {
	this := RetentionCompute{}
	return &this
}

// GetAggregation returns the Aggregation field value.
func (o *RetentionCompute) GetAggregation() EventsAggregation {
	if o == nil {
		var ret EventsAggregation
		return ret
	}
	return o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value
// and a boolean to check if the value has been set.
func (o *RetentionCompute) GetAggregationOk() (*EventsAggregation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregation, true
}

// SetAggregation sets field value.
func (o *RetentionCompute) SetAggregation(v EventsAggregation) {
	o.Aggregation = v
}

// GetMetric returns the Metric field value.
func (o *RetentionCompute) GetMetric() RetentionComputeMetric {
	if o == nil {
		var ret RetentionComputeMetric
		return ret
	}
	return o.Metric
}

// GetMetricOk returns a tuple with the Metric field value
// and a boolean to check if the value has been set.
func (o *RetentionCompute) GetMetricOk() (*RetentionComputeMetric, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metric, true
}

// SetMetric sets field value.
func (o *RetentionCompute) SetMetric(v RetentionComputeMetric) {
	o.Metric = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RetentionCompute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["aggregation"] = o.Aggregation
	toSerialize["metric"] = o.Metric
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RetentionCompute) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregation *EventsAggregation      `json:"aggregation"`
		Metric      *RetentionComputeMetric `json:"metric"`
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
	o.Aggregation = *all.Aggregation
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
