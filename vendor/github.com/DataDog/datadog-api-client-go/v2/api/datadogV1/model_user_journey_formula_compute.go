// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UserJourneyFormulaCompute Compute configuration for User Journey formula queries.
type UserJourneyFormulaCompute struct {
	// Aggregation methods for event platform queries.
	Aggregation FormulaAndFunctionEventAggregation `json:"aggregation"`
	// Time bucket interval in milliseconds for time series queries.
	Interval *float64 `json:"interval,omitempty"`
	// Metric for User Journey formula compute. `__dd.conversion` and `__dd.conversion_rate` accept `count` and `cardinality` as aggregations. `__dd.time_to_convert` accepts `avg`, `median`, `pc75`, `pc95`, `pc98`, `pc99`, `min`, and `max`.
	Metric *UserJourneyFormulaComputeMetric `json:"metric,omitempty"`
	// Target for user journey search.
	Target *UserJourneySearchTarget `json:"target,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewUserJourneyFormulaCompute instantiates a new UserJourneyFormulaCompute object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUserJourneyFormulaCompute(aggregation FormulaAndFunctionEventAggregation) *UserJourneyFormulaCompute {
	this := UserJourneyFormulaCompute{}
	this.Aggregation = aggregation
	return &this
}

// NewUserJourneyFormulaComputeWithDefaults instantiates a new UserJourneyFormulaCompute object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUserJourneyFormulaComputeWithDefaults() *UserJourneyFormulaCompute {
	this := UserJourneyFormulaCompute{}
	return &this
}

// GetAggregation returns the Aggregation field value.
func (o *UserJourneyFormulaCompute) GetAggregation() FormulaAndFunctionEventAggregation {
	if o == nil {
		var ret FormulaAndFunctionEventAggregation
		return ret
	}
	return o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value
// and a boolean to check if the value has been set.
func (o *UserJourneyFormulaCompute) GetAggregationOk() (*FormulaAndFunctionEventAggregation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregation, true
}

// SetAggregation sets field value.
func (o *UserJourneyFormulaCompute) SetAggregation(v FormulaAndFunctionEventAggregation) {
	o.Aggregation = v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *UserJourneyFormulaCompute) GetInterval() float64 {
	if o == nil || o.Interval == nil {
		var ret float64
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserJourneyFormulaCompute) GetIntervalOk() (*float64, bool) {
	if o == nil || o.Interval == nil {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *UserJourneyFormulaCompute) HasInterval() bool {
	return o != nil && o.Interval != nil
}

// SetInterval gets a reference to the given float64 and assigns it to the Interval field.
func (o *UserJourneyFormulaCompute) SetInterval(v float64) {
	o.Interval = &v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *UserJourneyFormulaCompute) GetMetric() UserJourneyFormulaComputeMetric {
	if o == nil || o.Metric == nil {
		var ret UserJourneyFormulaComputeMetric
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserJourneyFormulaCompute) GetMetricOk() (*UserJourneyFormulaComputeMetric, bool) {
	if o == nil || o.Metric == nil {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *UserJourneyFormulaCompute) HasMetric() bool {
	return o != nil && o.Metric != nil
}

// SetMetric gets a reference to the given UserJourneyFormulaComputeMetric and assigns it to the Metric field.
func (o *UserJourneyFormulaCompute) SetMetric(v UserJourneyFormulaComputeMetric) {
	o.Metric = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *UserJourneyFormulaCompute) GetTarget() UserJourneySearchTarget {
	if o == nil || o.Target == nil {
		var ret UserJourneySearchTarget
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserJourneyFormulaCompute) GetTargetOk() (*UserJourneySearchTarget, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *UserJourneyFormulaCompute) HasTarget() bool {
	return o != nil && o.Target != nil
}

// SetTarget gets a reference to the given UserJourneySearchTarget and assigns it to the Target field.
func (o *UserJourneyFormulaCompute) SetTarget(v UserJourneySearchTarget) {
	o.Target = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UserJourneyFormulaCompute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["aggregation"] = o.Aggregation
	if o.Interval != nil {
		toSerialize["interval"] = o.Interval
	}
	if o.Metric != nil {
		toSerialize["metric"] = o.Metric
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UserJourneyFormulaCompute) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregation *FormulaAndFunctionEventAggregation `json:"aggregation"`
		Interval    *float64                            `json:"interval,omitempty"`
		Metric      *UserJourneyFormulaComputeMetric    `json:"metric,omitempty"`
		Target      *UserJourneySearchTarget            `json:"target,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Aggregation == nil {
		return fmt.Errorf("required field aggregation missing")
	}

	hasInvalidField := false
	if !all.Aggregation.IsValid() {
		hasInvalidField = true
	} else {
		o.Aggregation = *all.Aggregation
	}
	o.Interval = all.Interval
	if all.Metric != nil && !all.Metric.IsValid() {
		hasInvalidField = true
	} else {
		o.Metric = all.Metric
	}
	if all.Target != nil && all.Target.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Target = all.Target

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
