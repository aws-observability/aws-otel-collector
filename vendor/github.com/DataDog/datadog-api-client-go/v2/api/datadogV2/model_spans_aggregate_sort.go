// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SpansAggregateSort A sort rule.
type SpansAggregateSort struct {
	// An aggregation function.
	Aggregation *SpansAggregationFunction `json:"aggregation,omitempty"`
	// The metric to sort by (only used for `type=measure`).
	Metric *string `json:"metric,omitempty"`
	// The order to use, ascending or descending.
	Order *SpansSortOrder `json:"order,omitempty"`
	// The type of sorting algorithm.
	Type *SpansAggregateSortType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSpansAggregateSort instantiates a new SpansAggregateSort object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSpansAggregateSort() *SpansAggregateSort {
	this := SpansAggregateSort{}
	var typeVar SpansAggregateSortType = SPANSAGGREGATESORTTYPE_ALPHABETICAL
	this.Type = &typeVar
	return &this
}

// NewSpansAggregateSortWithDefaults instantiates a new SpansAggregateSort object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSpansAggregateSortWithDefaults() *SpansAggregateSort {
	this := SpansAggregateSort{}
	var typeVar SpansAggregateSortType = SPANSAGGREGATESORTTYPE_ALPHABETICAL
	this.Type = &typeVar
	return &this
}

// GetAggregation returns the Aggregation field value if set, zero value otherwise.
func (o *SpansAggregateSort) GetAggregation() SpansAggregationFunction {
	if o == nil || o.Aggregation == nil {
		var ret SpansAggregationFunction
		return ret
	}
	return *o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansAggregateSort) GetAggregationOk() (*SpansAggregationFunction, bool) {
	if o == nil || o.Aggregation == nil {
		return nil, false
	}
	return o.Aggregation, true
}

// HasAggregation returns a boolean if a field has been set.
func (o *SpansAggregateSort) HasAggregation() bool {
	return o != nil && o.Aggregation != nil
}

// SetAggregation gets a reference to the given SpansAggregationFunction and assigns it to the Aggregation field.
func (o *SpansAggregateSort) SetAggregation(v SpansAggregationFunction) {
	o.Aggregation = &v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *SpansAggregateSort) GetMetric() string {
	if o == nil || o.Metric == nil {
		var ret string
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansAggregateSort) GetMetricOk() (*string, bool) {
	if o == nil || o.Metric == nil {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *SpansAggregateSort) HasMetric() bool {
	return o != nil && o.Metric != nil
}

// SetMetric gets a reference to the given string and assigns it to the Metric field.
func (o *SpansAggregateSort) SetMetric(v string) {
	o.Metric = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *SpansAggregateSort) GetOrder() SpansSortOrder {
	if o == nil || o.Order == nil {
		var ret SpansSortOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansAggregateSort) GetOrderOk() (*SpansSortOrder, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *SpansAggregateSort) HasOrder() bool {
	return o != nil && o.Order != nil
}

// SetOrder gets a reference to the given SpansSortOrder and assigns it to the Order field.
func (o *SpansAggregateSort) SetOrder(v SpansSortOrder) {
	o.Order = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SpansAggregateSort) GetType() SpansAggregateSortType {
	if o == nil || o.Type == nil {
		var ret SpansAggregateSortType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansAggregateSort) GetTypeOk() (*SpansAggregateSortType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SpansAggregateSort) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given SpansAggregateSortType and assigns it to the Type field.
func (o *SpansAggregateSort) SetType(v SpansAggregateSortType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SpansAggregateSort) MarshalJSON() ([]byte, error) {
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
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SpansAggregateSort) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregation *SpansAggregationFunction `json:"aggregation,omitempty"`
		Metric      *string                   `json:"metric,omitempty"`
		Order       *SpansSortOrder           `json:"order,omitempty"`
		Type        *SpansAggregateSortType   `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aggregation", "metric", "order", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Aggregation != nil && !all.Aggregation.IsValid() {
		hasInvalidField = true
	} else {
		o.Aggregation = all.Aggregation
	}
	o.Metric = all.Metric
	if all.Order != nil && !all.Order.IsValid() {
		hasInvalidField = true
	} else {
		o.Order = all.Order
	}
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
