// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsGroupBySort Sort configuration for group-by results.
type ProductAnalyticsGroupBySort struct {
	// The aggregation function to sort by.
	Aggregation *string `json:"aggregation,omitempty"`
	// The metric to sort by.
	Metric *string `json:"metric,omitempty"`
	// Direction of sort.
	Order *QuerySortOrder `json:"order,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsGroupBySort instantiates a new ProductAnalyticsGroupBySort object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsGroupBySort() *ProductAnalyticsGroupBySort {
	this := ProductAnalyticsGroupBySort{}
	var order QuerySortOrder = QUERYSORTORDER_DESC
	this.Order = &order
	return &this
}

// NewProductAnalyticsGroupBySortWithDefaults instantiates a new ProductAnalyticsGroupBySort object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsGroupBySortWithDefaults() *ProductAnalyticsGroupBySort {
	this := ProductAnalyticsGroupBySort{}
	var order QuerySortOrder = QUERYSORTORDER_DESC
	this.Order = &order
	return &this
}

// GetAggregation returns the Aggregation field value if set, zero value otherwise.
func (o *ProductAnalyticsGroupBySort) GetAggregation() string {
	if o == nil || o.Aggregation == nil {
		var ret string
		return ret
	}
	return *o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsGroupBySort) GetAggregationOk() (*string, bool) {
	if o == nil || o.Aggregation == nil {
		return nil, false
	}
	return o.Aggregation, true
}

// HasAggregation returns a boolean if a field has been set.
func (o *ProductAnalyticsGroupBySort) HasAggregation() bool {
	return o != nil && o.Aggregation != nil
}

// SetAggregation gets a reference to the given string and assigns it to the Aggregation field.
func (o *ProductAnalyticsGroupBySort) SetAggregation(v string) {
	o.Aggregation = &v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *ProductAnalyticsGroupBySort) GetMetric() string {
	if o == nil || o.Metric == nil {
		var ret string
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsGroupBySort) GetMetricOk() (*string, bool) {
	if o == nil || o.Metric == nil {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *ProductAnalyticsGroupBySort) HasMetric() bool {
	return o != nil && o.Metric != nil
}

// SetMetric gets a reference to the given string and assigns it to the Metric field.
func (o *ProductAnalyticsGroupBySort) SetMetric(v string) {
	o.Metric = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *ProductAnalyticsGroupBySort) GetOrder() QuerySortOrder {
	if o == nil || o.Order == nil {
		var ret QuerySortOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsGroupBySort) GetOrderOk() (*QuerySortOrder, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *ProductAnalyticsGroupBySort) HasOrder() bool {
	return o != nil && o.Order != nil
}

// SetOrder gets a reference to the given QuerySortOrder and assigns it to the Order field.
func (o *ProductAnalyticsGroupBySort) SetOrder(v QuerySortOrder) {
	o.Order = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsGroupBySort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Aggregation != nil {
		toSerialize["aggregation"] = o.Aggregation
	}
	if o.Metric != nil {
		toSerialize["metric"] = o.Metric
	}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsGroupBySort) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregation *string         `json:"aggregation,omitempty"`
		Metric      *string         `json:"metric,omitempty"`
		Order       *QuerySortOrder `json:"order,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aggregation", "metric", "order"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Aggregation = all.Aggregation
	o.Metric = all.Metric
	if all.Order != nil && !all.Order.IsValid() {
		hasInvalidField = true
	} else {
		o.Order = all.Order
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
