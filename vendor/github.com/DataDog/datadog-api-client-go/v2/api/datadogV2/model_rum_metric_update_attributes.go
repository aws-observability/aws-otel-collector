// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumMetricUpdateAttributes The rum-based metric properties that will be updated.
type RumMetricUpdateAttributes struct {
	// The compute rule to compute the rum-based metric.
	Compute *RumMetricUpdateCompute `json:"compute,omitempty"`
	// The rum-based metric filter. Events matching this filter will be aggregated in this metric.
	Filter *RumMetricFilter `json:"filter,omitempty"`
	// The rules for the group by.
	GroupBy []RumMetricGroupBy `json:"group_by,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumMetricUpdateAttributes instantiates a new RumMetricUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumMetricUpdateAttributes() *RumMetricUpdateAttributes {
	this := RumMetricUpdateAttributes{}
	return &this
}

// NewRumMetricUpdateAttributesWithDefaults instantiates a new RumMetricUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumMetricUpdateAttributesWithDefaults() *RumMetricUpdateAttributes {
	this := RumMetricUpdateAttributes{}
	return &this
}

// GetCompute returns the Compute field value if set, zero value otherwise.
func (o *RumMetricUpdateAttributes) GetCompute() RumMetricUpdateCompute {
	if o == nil || o.Compute == nil {
		var ret RumMetricUpdateCompute
		return ret
	}
	return *o.Compute
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumMetricUpdateAttributes) GetComputeOk() (*RumMetricUpdateCompute, bool) {
	if o == nil || o.Compute == nil {
		return nil, false
	}
	return o.Compute, true
}

// HasCompute returns a boolean if a field has been set.
func (o *RumMetricUpdateAttributes) HasCompute() bool {
	return o != nil && o.Compute != nil
}

// SetCompute gets a reference to the given RumMetricUpdateCompute and assigns it to the Compute field.
func (o *RumMetricUpdateAttributes) SetCompute(v RumMetricUpdateCompute) {
	o.Compute = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *RumMetricUpdateAttributes) GetFilter() RumMetricFilter {
	if o == nil || o.Filter == nil {
		var ret RumMetricFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumMetricUpdateAttributes) GetFilterOk() (*RumMetricFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *RumMetricUpdateAttributes) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given RumMetricFilter and assigns it to the Filter field.
func (o *RumMetricUpdateAttributes) SetFilter(v RumMetricFilter) {
	o.Filter = &v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *RumMetricUpdateAttributes) GetGroupBy() []RumMetricGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []RumMetricGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumMetricUpdateAttributes) GetGroupByOk() (*[]RumMetricGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *RumMetricUpdateAttributes) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []RumMetricGroupBy and assigns it to the GroupBy field.
func (o *RumMetricUpdateAttributes) SetGroupBy(v []RumMetricGroupBy) {
	o.GroupBy = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumMetricUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Compute != nil {
		toSerialize["compute"] = o.Compute
	}
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
func (o *RumMetricUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Compute *RumMetricUpdateCompute `json:"compute,omitempty"`
		Filter  *RumMetricFilter        `json:"filter,omitempty"`
		GroupBy []RumMetricGroupBy      `json:"group_by,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"compute", "filter", "group_by"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Compute != nil && all.Compute.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Compute = all.Compute
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
