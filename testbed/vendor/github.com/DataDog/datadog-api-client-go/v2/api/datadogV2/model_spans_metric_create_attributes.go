// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SpansMetricCreateAttributes The object describing the Datadog span-based metric to create.
type SpansMetricCreateAttributes struct {
	// The compute rule to compute the span-based metric.
	Compute SpansMetricCompute `json:"compute"`
	// The span-based metric filter. Spans matching this filter will be aggregated in this metric.
	Filter *SpansMetricFilter `json:"filter,omitempty"`
	// The rules for the group by.
	GroupBy []SpansMetricGroupBy `json:"group_by,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSpansMetricCreateAttributes instantiates a new SpansMetricCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSpansMetricCreateAttributes(compute SpansMetricCompute) *SpansMetricCreateAttributes {
	this := SpansMetricCreateAttributes{}
	this.Compute = compute
	return &this
}

// NewSpansMetricCreateAttributesWithDefaults instantiates a new SpansMetricCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSpansMetricCreateAttributesWithDefaults() *SpansMetricCreateAttributes {
	this := SpansMetricCreateAttributes{}
	return &this
}

// GetCompute returns the Compute field value.
func (o *SpansMetricCreateAttributes) GetCompute() SpansMetricCompute {
	if o == nil {
		var ret SpansMetricCompute
		return ret
	}
	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value
// and a boolean to check if the value has been set.
func (o *SpansMetricCreateAttributes) GetComputeOk() (*SpansMetricCompute, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Compute, true
}

// SetCompute sets field value.
func (o *SpansMetricCreateAttributes) SetCompute(v SpansMetricCompute) {
	o.Compute = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *SpansMetricCreateAttributes) GetFilter() SpansMetricFilter {
	if o == nil || o.Filter == nil {
		var ret SpansMetricFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansMetricCreateAttributes) GetFilterOk() (*SpansMetricFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *SpansMetricCreateAttributes) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given SpansMetricFilter and assigns it to the Filter field.
func (o *SpansMetricCreateAttributes) SetFilter(v SpansMetricFilter) {
	o.Filter = &v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *SpansMetricCreateAttributes) GetGroupBy() []SpansMetricGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []SpansMetricGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansMetricCreateAttributes) GetGroupByOk() (*[]SpansMetricGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *SpansMetricCreateAttributes) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []SpansMetricGroupBy and assigns it to the GroupBy field.
func (o *SpansMetricCreateAttributes) SetGroupBy(v []SpansMetricGroupBy) {
	o.GroupBy = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SpansMetricCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["compute"] = o.Compute
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SpansMetricCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Compute *SpansMetricCompute  `json:"compute"`
		Filter  *SpansMetricFilter   `json:"filter,omitempty"`
		GroupBy []SpansMetricGroupBy `json:"group_by,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Compute == nil {
		return fmt.Errorf("required field compute missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"compute", "filter", "group_by"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Compute.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Compute = *all.Compute
	if all.Filter != nil && all.Filter.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Filter = all.Filter
	o.GroupBy = all.GroupBy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
