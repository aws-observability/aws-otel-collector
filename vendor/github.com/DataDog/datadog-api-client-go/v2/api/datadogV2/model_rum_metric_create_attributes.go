// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumMetricCreateAttributes The object describing the Datadog rum-based metric to create.
type RumMetricCreateAttributes struct {
	// The compute rule to compute the rum-based metric.
	Compute RumMetricCompute `json:"compute"`
	// The type of RUM events to filter on.
	EventType RumMetricEventType `json:"event_type"`
	// The rum-based metric filter. Events matching this filter will be aggregated in this metric.
	Filter *RumMetricFilter `json:"filter,omitempty"`
	// The rules for the group by.
	GroupBy []RumMetricGroupBy `json:"group_by,omitempty"`
	// The rule to count updatable events. Is only set if `event_type` is `sessions` or `views`.
	Uniqueness *RumMetricUniqueness `json:"uniqueness,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumMetricCreateAttributes instantiates a new RumMetricCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumMetricCreateAttributes(compute RumMetricCompute, eventType RumMetricEventType) *RumMetricCreateAttributes {
	this := RumMetricCreateAttributes{}
	this.Compute = compute
	this.EventType = eventType
	return &this
}

// NewRumMetricCreateAttributesWithDefaults instantiates a new RumMetricCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumMetricCreateAttributesWithDefaults() *RumMetricCreateAttributes {
	this := RumMetricCreateAttributes{}
	return &this
}

// GetCompute returns the Compute field value.
func (o *RumMetricCreateAttributes) GetCompute() RumMetricCompute {
	if o == nil {
		var ret RumMetricCompute
		return ret
	}
	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value
// and a boolean to check if the value has been set.
func (o *RumMetricCreateAttributes) GetComputeOk() (*RumMetricCompute, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Compute, true
}

// SetCompute sets field value.
func (o *RumMetricCreateAttributes) SetCompute(v RumMetricCompute) {
	o.Compute = v
}

// GetEventType returns the EventType field value.
func (o *RumMetricCreateAttributes) GetEventType() RumMetricEventType {
	if o == nil {
		var ret RumMetricEventType
		return ret
	}
	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *RumMetricCreateAttributes) GetEventTypeOk() (*RumMetricEventType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value.
func (o *RumMetricCreateAttributes) SetEventType(v RumMetricEventType) {
	o.EventType = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *RumMetricCreateAttributes) GetFilter() RumMetricFilter {
	if o == nil || o.Filter == nil {
		var ret RumMetricFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumMetricCreateAttributes) GetFilterOk() (*RumMetricFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *RumMetricCreateAttributes) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given RumMetricFilter and assigns it to the Filter field.
func (o *RumMetricCreateAttributes) SetFilter(v RumMetricFilter) {
	o.Filter = &v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *RumMetricCreateAttributes) GetGroupBy() []RumMetricGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []RumMetricGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumMetricCreateAttributes) GetGroupByOk() (*[]RumMetricGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *RumMetricCreateAttributes) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []RumMetricGroupBy and assigns it to the GroupBy field.
func (o *RumMetricCreateAttributes) SetGroupBy(v []RumMetricGroupBy) {
	o.GroupBy = v
}

// GetUniqueness returns the Uniqueness field value if set, zero value otherwise.
func (o *RumMetricCreateAttributes) GetUniqueness() RumMetricUniqueness {
	if o == nil || o.Uniqueness == nil {
		var ret RumMetricUniqueness
		return ret
	}
	return *o.Uniqueness
}

// GetUniquenessOk returns a tuple with the Uniqueness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumMetricCreateAttributes) GetUniquenessOk() (*RumMetricUniqueness, bool) {
	if o == nil || o.Uniqueness == nil {
		return nil, false
	}
	return o.Uniqueness, true
}

// HasUniqueness returns a boolean if a field has been set.
func (o *RumMetricCreateAttributes) HasUniqueness() bool {
	return o != nil && o.Uniqueness != nil
}

// SetUniqueness gets a reference to the given RumMetricUniqueness and assigns it to the Uniqueness field.
func (o *RumMetricCreateAttributes) SetUniqueness(v RumMetricUniqueness) {
	o.Uniqueness = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumMetricCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["compute"] = o.Compute
	toSerialize["event_type"] = o.EventType
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	if o.Uniqueness != nil {
		toSerialize["uniqueness"] = o.Uniqueness
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumMetricCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Compute    *RumMetricCompute    `json:"compute"`
		EventType  *RumMetricEventType  `json:"event_type"`
		Filter     *RumMetricFilter     `json:"filter,omitempty"`
		GroupBy    []RumMetricGroupBy   `json:"group_by,omitempty"`
		Uniqueness *RumMetricUniqueness `json:"uniqueness,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Compute == nil {
		return fmt.Errorf("required field compute missing")
	}
	if all.EventType == nil {
		return fmt.Errorf("required field event_type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"compute", "event_type", "filter", "group_by", "uniqueness"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Compute.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Compute = *all.Compute
	if !all.EventType.IsValid() {
		hasInvalidField = true
	} else {
		o.EventType = *all.EventType
	}
	if all.Filter != nil && all.Filter.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Filter = all.Filter
	o.GroupBy = all.GroupBy
	if all.Uniqueness != nil && all.Uniqueness.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Uniqueness = all.Uniqueness

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
