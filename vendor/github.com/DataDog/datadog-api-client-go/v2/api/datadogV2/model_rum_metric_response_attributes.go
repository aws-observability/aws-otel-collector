// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumMetricResponseAttributes The object describing a Datadog rum-based metric.
type RumMetricResponseAttributes struct {
	// The compute rule to compute the rum-based metric.
	Compute *RumMetricResponseCompute `json:"compute,omitempty"`
	// The type of RUM events to filter on.
	EventType *RumMetricEventType `json:"event_type,omitempty"`
	// The rum-based metric filter. RUM events matching this filter will be aggregated in this metric.
	Filter *RumMetricResponseFilter `json:"filter,omitempty"`
	// The rules for the group by.
	GroupBy []RumMetricResponseGroupBy `json:"group_by,omitempty"`
	// The rule to count updatable events. Is only set if `event_type` is `session` or `view`.
	Uniqueness *RumMetricResponseUniqueness `json:"uniqueness,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumMetricResponseAttributes instantiates a new RumMetricResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumMetricResponseAttributes() *RumMetricResponseAttributes {
	this := RumMetricResponseAttributes{}
	return &this
}

// NewRumMetricResponseAttributesWithDefaults instantiates a new RumMetricResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumMetricResponseAttributesWithDefaults() *RumMetricResponseAttributes {
	this := RumMetricResponseAttributes{}
	return &this
}

// GetCompute returns the Compute field value if set, zero value otherwise.
func (o *RumMetricResponseAttributes) GetCompute() RumMetricResponseCompute {
	if o == nil || o.Compute == nil {
		var ret RumMetricResponseCompute
		return ret
	}
	return *o.Compute
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumMetricResponseAttributes) GetComputeOk() (*RumMetricResponseCompute, bool) {
	if o == nil || o.Compute == nil {
		return nil, false
	}
	return o.Compute, true
}

// HasCompute returns a boolean if a field has been set.
func (o *RumMetricResponseAttributes) HasCompute() bool {
	return o != nil && o.Compute != nil
}

// SetCompute gets a reference to the given RumMetricResponseCompute and assigns it to the Compute field.
func (o *RumMetricResponseAttributes) SetCompute(v RumMetricResponseCompute) {
	o.Compute = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *RumMetricResponseAttributes) GetEventType() RumMetricEventType {
	if o == nil || o.EventType == nil {
		var ret RumMetricEventType
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumMetricResponseAttributes) GetEventTypeOk() (*RumMetricEventType, bool) {
	if o == nil || o.EventType == nil {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *RumMetricResponseAttributes) HasEventType() bool {
	return o != nil && o.EventType != nil
}

// SetEventType gets a reference to the given RumMetricEventType and assigns it to the EventType field.
func (o *RumMetricResponseAttributes) SetEventType(v RumMetricEventType) {
	o.EventType = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *RumMetricResponseAttributes) GetFilter() RumMetricResponseFilter {
	if o == nil || o.Filter == nil {
		var ret RumMetricResponseFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumMetricResponseAttributes) GetFilterOk() (*RumMetricResponseFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *RumMetricResponseAttributes) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given RumMetricResponseFilter and assigns it to the Filter field.
func (o *RumMetricResponseAttributes) SetFilter(v RumMetricResponseFilter) {
	o.Filter = &v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *RumMetricResponseAttributes) GetGroupBy() []RumMetricResponseGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []RumMetricResponseGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumMetricResponseAttributes) GetGroupByOk() (*[]RumMetricResponseGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *RumMetricResponseAttributes) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []RumMetricResponseGroupBy and assigns it to the GroupBy field.
func (o *RumMetricResponseAttributes) SetGroupBy(v []RumMetricResponseGroupBy) {
	o.GroupBy = v
}

// GetUniqueness returns the Uniqueness field value if set, zero value otherwise.
func (o *RumMetricResponseAttributes) GetUniqueness() RumMetricResponseUniqueness {
	if o == nil || o.Uniqueness == nil {
		var ret RumMetricResponseUniqueness
		return ret
	}
	return *o.Uniqueness
}

// GetUniquenessOk returns a tuple with the Uniqueness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumMetricResponseAttributes) GetUniquenessOk() (*RumMetricResponseUniqueness, bool) {
	if o == nil || o.Uniqueness == nil {
		return nil, false
	}
	return o.Uniqueness, true
}

// HasUniqueness returns a boolean if a field has been set.
func (o *RumMetricResponseAttributes) HasUniqueness() bool {
	return o != nil && o.Uniqueness != nil
}

// SetUniqueness gets a reference to the given RumMetricResponseUniqueness and assigns it to the Uniqueness field.
func (o *RumMetricResponseAttributes) SetUniqueness(v RumMetricResponseUniqueness) {
	o.Uniqueness = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumMetricResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Compute != nil {
		toSerialize["compute"] = o.Compute
	}
	if o.EventType != nil {
		toSerialize["event_type"] = o.EventType
	}
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
func (o *RumMetricResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Compute    *RumMetricResponseCompute    `json:"compute,omitempty"`
		EventType  *RumMetricEventType          `json:"event_type,omitempty"`
		Filter     *RumMetricResponseFilter     `json:"filter,omitempty"`
		GroupBy    []RumMetricResponseGroupBy   `json:"group_by,omitempty"`
		Uniqueness *RumMetricResponseUniqueness `json:"uniqueness,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"compute", "event_type", "filter", "group_by", "uniqueness"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Compute != nil && all.Compute.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Compute = all.Compute
	if all.EventType != nil && !all.EventType.IsValid() {
		hasInvalidField = true
	} else {
		o.EventType = all.EventType
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
