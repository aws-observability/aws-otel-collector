// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsFunnelGroupBySort Sort configuration for user journey funnel group by.
type ProductAnalyticsFunnelGroupBySort struct {
	// Aggregation type.
	Aggregation string `json:"aggregation"`
	// Metric to sort by.
	Metric *string `json:"metric,omitempty"`
	// Widget sorting methods.
	Order *WidgetSort `json:"order,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewProductAnalyticsFunnelGroupBySort instantiates a new ProductAnalyticsFunnelGroupBySort object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsFunnelGroupBySort(aggregation string) *ProductAnalyticsFunnelGroupBySort {
	this := ProductAnalyticsFunnelGroupBySort{}
	this.Aggregation = aggregation
	return &this
}

// NewProductAnalyticsFunnelGroupBySortWithDefaults instantiates a new ProductAnalyticsFunnelGroupBySort object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsFunnelGroupBySortWithDefaults() *ProductAnalyticsFunnelGroupBySort {
	this := ProductAnalyticsFunnelGroupBySort{}
	return &this
}

// GetAggregation returns the Aggregation field value.
func (o *ProductAnalyticsFunnelGroupBySort) GetAggregation() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFunnelGroupBySort) GetAggregationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregation, true
}

// SetAggregation sets field value.
func (o *ProductAnalyticsFunnelGroupBySort) SetAggregation(v string) {
	o.Aggregation = v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *ProductAnalyticsFunnelGroupBySort) GetMetric() string {
	if o == nil || o.Metric == nil {
		var ret string
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFunnelGroupBySort) GetMetricOk() (*string, bool) {
	if o == nil || o.Metric == nil {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *ProductAnalyticsFunnelGroupBySort) HasMetric() bool {
	return o != nil && o.Metric != nil
}

// SetMetric gets a reference to the given string and assigns it to the Metric field.
func (o *ProductAnalyticsFunnelGroupBySort) SetMetric(v string) {
	o.Metric = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *ProductAnalyticsFunnelGroupBySort) GetOrder() WidgetSort {
	if o == nil || o.Order == nil {
		var ret WidgetSort
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFunnelGroupBySort) GetOrderOk() (*WidgetSort, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *ProductAnalyticsFunnelGroupBySort) HasOrder() bool {
	return o != nil && o.Order != nil
}

// SetOrder gets a reference to the given WidgetSort and assigns it to the Order field.
func (o *ProductAnalyticsFunnelGroupBySort) SetOrder(v WidgetSort) {
	o.Order = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsFunnelGroupBySort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["aggregation"] = o.Aggregation
	if o.Metric != nil {
		toSerialize["metric"] = o.Metric
	}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsFunnelGroupBySort) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregation *string     `json:"aggregation"`
		Metric      *string     `json:"metric,omitempty"`
		Order       *WidgetSort `json:"order,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Aggregation == nil {
		return fmt.Errorf("required field aggregation missing")
	}

	hasInvalidField := false
	o.Aggregation = *all.Aggregation
	o.Metric = all.Metric
	if all.Order != nil && !all.Order.IsValid() {
		hasInvalidField = true
	} else {
		o.Order = all.Order
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
