// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SpansAggregateRequestAttributes The object containing all the query parameters.
type SpansAggregateRequestAttributes struct {
	// The list of metrics or timeseries to compute for the retrieved buckets.
	Compute []SpansCompute `json:"compute,omitempty"`
	// The search and filter query settings.
	Filter *SpansQueryFilter `json:"filter,omitempty"`
	// The rules for the group by.
	GroupBy []SpansGroupBy `json:"group_by,omitempty"`
	// Global query options that are used during the query.
	// Note: You should only supply timezone or time offset but not both otherwise the query will fail.
	Options *SpansQueryOptions `json:"options,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSpansAggregateRequestAttributes instantiates a new SpansAggregateRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSpansAggregateRequestAttributes() *SpansAggregateRequestAttributes {
	this := SpansAggregateRequestAttributes{}
	return &this
}

// NewSpansAggregateRequestAttributesWithDefaults instantiates a new SpansAggregateRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSpansAggregateRequestAttributesWithDefaults() *SpansAggregateRequestAttributes {
	this := SpansAggregateRequestAttributes{}
	return &this
}

// GetCompute returns the Compute field value if set, zero value otherwise.
func (o *SpansAggregateRequestAttributes) GetCompute() []SpansCompute {
	if o == nil || o.Compute == nil {
		var ret []SpansCompute
		return ret
	}
	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansAggregateRequestAttributes) GetComputeOk() (*[]SpansCompute, bool) {
	if o == nil || o.Compute == nil {
		return nil, false
	}
	return &o.Compute, true
}

// HasCompute returns a boolean if a field has been set.
func (o *SpansAggregateRequestAttributes) HasCompute() bool {
	return o != nil && o.Compute != nil
}

// SetCompute gets a reference to the given []SpansCompute and assigns it to the Compute field.
func (o *SpansAggregateRequestAttributes) SetCompute(v []SpansCompute) {
	o.Compute = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *SpansAggregateRequestAttributes) GetFilter() SpansQueryFilter {
	if o == nil || o.Filter == nil {
		var ret SpansQueryFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansAggregateRequestAttributes) GetFilterOk() (*SpansQueryFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *SpansAggregateRequestAttributes) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given SpansQueryFilter and assigns it to the Filter field.
func (o *SpansAggregateRequestAttributes) SetFilter(v SpansQueryFilter) {
	o.Filter = &v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *SpansAggregateRequestAttributes) GetGroupBy() []SpansGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []SpansGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansAggregateRequestAttributes) GetGroupByOk() (*[]SpansGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *SpansAggregateRequestAttributes) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []SpansGroupBy and assigns it to the GroupBy field.
func (o *SpansAggregateRequestAttributes) SetGroupBy(v []SpansGroupBy) {
	o.GroupBy = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *SpansAggregateRequestAttributes) GetOptions() SpansQueryOptions {
	if o == nil || o.Options == nil {
		var ret SpansQueryOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansAggregateRequestAttributes) GetOptionsOk() (*SpansQueryOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *SpansAggregateRequestAttributes) HasOptions() bool {
	return o != nil && o.Options != nil
}

// SetOptions gets a reference to the given SpansQueryOptions and assigns it to the Options field.
func (o *SpansAggregateRequestAttributes) SetOptions(v SpansQueryOptions) {
	o.Options = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SpansAggregateRequestAttributes) MarshalJSON() ([]byte, error) {
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
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SpansAggregateRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Compute []SpansCompute     `json:"compute,omitempty"`
		Filter  *SpansQueryFilter  `json:"filter,omitempty"`
		GroupBy []SpansGroupBy     `json:"group_by,omitempty"`
		Options *SpansQueryOptions `json:"options,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"compute", "filter", "group_by", "options"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Compute = all.Compute
	if all.Filter != nil && all.Filter.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Filter = all.Filter
	o.GroupBy = all.GroupBy
	if all.Options != nil && all.Options.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Options = all.Options

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
