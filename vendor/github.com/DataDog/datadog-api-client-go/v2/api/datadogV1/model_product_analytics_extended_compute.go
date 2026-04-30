// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsExtendedCompute Compute configuration for Product Analytics Extended queries.
type ProductAnalyticsExtendedCompute struct {
	// Aggregation methods for event platform queries.
	Aggregation FormulaAndFunctionEventAggregation `json:"aggregation"`
	// Fixed-width time bucket interval in milliseconds for time series queries. Mutually exclusive with `rollup`.
	Interval *float64 `json:"interval,omitempty"`
	// Measurable attribute to compute.
	Metric *string `json:"metric,omitempty"`
	// Name of the compute for use in formulas.
	Name *string `json:"name,omitempty"`
	// Calendar interval definition.
	Rollup *CalendarInterval `json:"rollup,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewProductAnalyticsExtendedCompute instantiates a new ProductAnalyticsExtendedCompute object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsExtendedCompute(aggregation FormulaAndFunctionEventAggregation) *ProductAnalyticsExtendedCompute {
	this := ProductAnalyticsExtendedCompute{}
	this.Aggregation = aggregation
	return &this
}

// NewProductAnalyticsExtendedComputeWithDefaults instantiates a new ProductAnalyticsExtendedCompute object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsExtendedComputeWithDefaults() *ProductAnalyticsExtendedCompute {
	this := ProductAnalyticsExtendedCompute{}
	return &this
}

// GetAggregation returns the Aggregation field value.
func (o *ProductAnalyticsExtendedCompute) GetAggregation() FormulaAndFunctionEventAggregation {
	if o == nil {
		var ret FormulaAndFunctionEventAggregation
		return ret
	}
	return o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsExtendedCompute) GetAggregationOk() (*FormulaAndFunctionEventAggregation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregation, true
}

// SetAggregation sets field value.
func (o *ProductAnalyticsExtendedCompute) SetAggregation(v FormulaAndFunctionEventAggregation) {
	o.Aggregation = v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *ProductAnalyticsExtendedCompute) GetInterval() float64 {
	if o == nil || o.Interval == nil {
		var ret float64
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsExtendedCompute) GetIntervalOk() (*float64, bool) {
	if o == nil || o.Interval == nil {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *ProductAnalyticsExtendedCompute) HasInterval() bool {
	return o != nil && o.Interval != nil
}

// SetInterval gets a reference to the given float64 and assigns it to the Interval field.
func (o *ProductAnalyticsExtendedCompute) SetInterval(v float64) {
	o.Interval = &v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *ProductAnalyticsExtendedCompute) GetMetric() string {
	if o == nil || o.Metric == nil {
		var ret string
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsExtendedCompute) GetMetricOk() (*string, bool) {
	if o == nil || o.Metric == nil {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *ProductAnalyticsExtendedCompute) HasMetric() bool {
	return o != nil && o.Metric != nil
}

// SetMetric gets a reference to the given string and assigns it to the Metric field.
func (o *ProductAnalyticsExtendedCompute) SetMetric(v string) {
	o.Metric = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProductAnalyticsExtendedCompute) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsExtendedCompute) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProductAnalyticsExtendedCompute) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProductAnalyticsExtendedCompute) SetName(v string) {
	o.Name = &v
}

// GetRollup returns the Rollup field value if set, zero value otherwise.
func (o *ProductAnalyticsExtendedCompute) GetRollup() CalendarInterval {
	if o == nil || o.Rollup == nil {
		var ret CalendarInterval
		return ret
	}
	return *o.Rollup
}

// GetRollupOk returns a tuple with the Rollup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsExtendedCompute) GetRollupOk() (*CalendarInterval, bool) {
	if o == nil || o.Rollup == nil {
		return nil, false
	}
	return o.Rollup, true
}

// HasRollup returns a boolean if a field has been set.
func (o *ProductAnalyticsExtendedCompute) HasRollup() bool {
	return o != nil && o.Rollup != nil
}

// SetRollup gets a reference to the given CalendarInterval and assigns it to the Rollup field.
func (o *ProductAnalyticsExtendedCompute) SetRollup(v CalendarInterval) {
	o.Rollup = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsExtendedCompute) MarshalJSON() ([]byte, error) {
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
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Rollup != nil {
		toSerialize["rollup"] = o.Rollup
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsExtendedCompute) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregation *FormulaAndFunctionEventAggregation `json:"aggregation"`
		Interval    *float64                            `json:"interval,omitempty"`
		Metric      *string                             `json:"metric,omitempty"`
		Name        *string                             `json:"name,omitempty"`
		Rollup      *CalendarInterval                   `json:"rollup,omitempty"`
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
	o.Metric = all.Metric
	o.Name = all.Name
	if all.Rollup != nil && all.Rollup.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Rollup = all.Rollup

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
