// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsMetricResponseAttributes The object describing a Datadog log-based metric.
type LogsMetricResponseAttributes struct {
	// The compute rule to compute the log-based metric.
	Compute *LogsMetricResponseCompute `json:"compute,omitempty"`
	// The log-based metric filter. Logs matching this filter will be aggregated in this metric.
	Filter *LogsMetricResponseFilter `json:"filter,omitempty"`
	// The rules for the group by.
	GroupBy []LogsMetricResponseGroupBy `json:"group_by,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLogsMetricResponseAttributes instantiates a new LogsMetricResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsMetricResponseAttributes() *LogsMetricResponseAttributes {
	this := LogsMetricResponseAttributes{}
	return &this
}

// NewLogsMetricResponseAttributesWithDefaults instantiates a new LogsMetricResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsMetricResponseAttributesWithDefaults() *LogsMetricResponseAttributes {
	this := LogsMetricResponseAttributes{}
	return &this
}

// GetCompute returns the Compute field value if set, zero value otherwise.
func (o *LogsMetricResponseAttributes) GetCompute() LogsMetricResponseCompute {
	if o == nil || o.Compute == nil {
		var ret LogsMetricResponseCompute
		return ret
	}
	return *o.Compute
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsMetricResponseAttributes) GetComputeOk() (*LogsMetricResponseCompute, bool) {
	if o == nil || o.Compute == nil {
		return nil, false
	}
	return o.Compute, true
}

// HasCompute returns a boolean if a field has been set.
func (o *LogsMetricResponseAttributes) HasCompute() bool {
	return o != nil && o.Compute != nil
}

// SetCompute gets a reference to the given LogsMetricResponseCompute and assigns it to the Compute field.
func (o *LogsMetricResponseAttributes) SetCompute(v LogsMetricResponseCompute) {
	o.Compute = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *LogsMetricResponseAttributes) GetFilter() LogsMetricResponseFilter {
	if o == nil || o.Filter == nil {
		var ret LogsMetricResponseFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsMetricResponseAttributes) GetFilterOk() (*LogsMetricResponseFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *LogsMetricResponseAttributes) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given LogsMetricResponseFilter and assigns it to the Filter field.
func (o *LogsMetricResponseAttributes) SetFilter(v LogsMetricResponseFilter) {
	o.Filter = &v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *LogsMetricResponseAttributes) GetGroupBy() []LogsMetricResponseGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []LogsMetricResponseGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsMetricResponseAttributes) GetGroupByOk() (*[]LogsMetricResponseGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *LogsMetricResponseAttributes) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []LogsMetricResponseGroupBy and assigns it to the GroupBy field.
func (o *LogsMetricResponseAttributes) SetGroupBy(v []LogsMetricResponseGroupBy) {
	o.GroupBy = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsMetricResponseAttributes) MarshalJSON() ([]byte, error) {
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
func (o *LogsMetricResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Compute *LogsMetricResponseCompute  `json:"compute,omitempty"`
		Filter  *LogsMetricResponseFilter   `json:"filter,omitempty"`
		GroupBy []LogsMetricResponseGroupBy `json:"group_by,omitempty"`
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
