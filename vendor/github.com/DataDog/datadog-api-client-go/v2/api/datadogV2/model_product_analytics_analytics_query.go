// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsAnalyticsQuery The analytics query definition containing a base query, compute rule, and optional grouping.
type ProductAnalyticsAnalyticsQuery struct {
	// Audience filter definitions for targeting specific user segments.
	AudienceFilters *ProductAnalyticsAudienceFilters `json:"audience_filters,omitempty"`
	// A compute rule for aggregating data.
	Compute ProductAnalyticsCompute `json:"compute"`
	// Group-by rules for segmenting results.
	GroupBy []ProductAnalyticsGroupBy `json:"group_by,omitempty"`
	// Restrict the query to specific indexes. Max 1 entry.
	Indexes []string `json:"indexes,omitempty"`
	// A query definition discriminated by the `data_source` field.
	// Use `product_analytics` for standard event queries, or
	// `product_analytics_occurrence` for occurrence-filtered queries.
	Query ProductAnalyticsBaseQuery `json:"query"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsAnalyticsQuery instantiates a new ProductAnalyticsAnalyticsQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsAnalyticsQuery(compute ProductAnalyticsCompute, query ProductAnalyticsBaseQuery) *ProductAnalyticsAnalyticsQuery {
	this := ProductAnalyticsAnalyticsQuery{}
	this.Compute = compute
	this.Query = query
	return &this
}

// NewProductAnalyticsAnalyticsQueryWithDefaults instantiates a new ProductAnalyticsAnalyticsQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsAnalyticsQueryWithDefaults() *ProductAnalyticsAnalyticsQuery {
	this := ProductAnalyticsAnalyticsQuery{}
	return &this
}

// GetAudienceFilters returns the AudienceFilters field value if set, zero value otherwise.
func (o *ProductAnalyticsAnalyticsQuery) GetAudienceFilters() ProductAnalyticsAudienceFilters {
	if o == nil || o.AudienceFilters == nil {
		var ret ProductAnalyticsAudienceFilters
		return ret
	}
	return *o.AudienceFilters
}

// GetAudienceFiltersOk returns a tuple with the AudienceFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsAnalyticsQuery) GetAudienceFiltersOk() (*ProductAnalyticsAudienceFilters, bool) {
	if o == nil || o.AudienceFilters == nil {
		return nil, false
	}
	return o.AudienceFilters, true
}

// HasAudienceFilters returns a boolean if a field has been set.
func (o *ProductAnalyticsAnalyticsQuery) HasAudienceFilters() bool {
	return o != nil && o.AudienceFilters != nil
}

// SetAudienceFilters gets a reference to the given ProductAnalyticsAudienceFilters and assigns it to the AudienceFilters field.
func (o *ProductAnalyticsAnalyticsQuery) SetAudienceFilters(v ProductAnalyticsAudienceFilters) {
	o.AudienceFilters = &v
}

// GetCompute returns the Compute field value.
func (o *ProductAnalyticsAnalyticsQuery) GetCompute() ProductAnalyticsCompute {
	if o == nil {
		var ret ProductAnalyticsCompute
		return ret
	}
	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsAnalyticsQuery) GetComputeOk() (*ProductAnalyticsCompute, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Compute, true
}

// SetCompute sets field value.
func (o *ProductAnalyticsAnalyticsQuery) SetCompute(v ProductAnalyticsCompute) {
	o.Compute = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *ProductAnalyticsAnalyticsQuery) GetGroupBy() []ProductAnalyticsGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []ProductAnalyticsGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsAnalyticsQuery) GetGroupByOk() (*[]ProductAnalyticsGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *ProductAnalyticsAnalyticsQuery) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []ProductAnalyticsGroupBy and assigns it to the GroupBy field.
func (o *ProductAnalyticsAnalyticsQuery) SetGroupBy(v []ProductAnalyticsGroupBy) {
	o.GroupBy = v
}

// GetIndexes returns the Indexes field value if set, zero value otherwise.
func (o *ProductAnalyticsAnalyticsQuery) GetIndexes() []string {
	if o == nil || o.Indexes == nil {
		var ret []string
		return ret
	}
	return o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsAnalyticsQuery) GetIndexesOk() (*[]string, bool) {
	if o == nil || o.Indexes == nil {
		return nil, false
	}
	return &o.Indexes, true
}

// HasIndexes returns a boolean if a field has been set.
func (o *ProductAnalyticsAnalyticsQuery) HasIndexes() bool {
	return o != nil && o.Indexes != nil
}

// SetIndexes gets a reference to the given []string and assigns it to the Indexes field.
func (o *ProductAnalyticsAnalyticsQuery) SetIndexes(v []string) {
	o.Indexes = v
}

// GetQuery returns the Query field value.
func (o *ProductAnalyticsAnalyticsQuery) GetQuery() ProductAnalyticsBaseQuery {
	if o == nil {
		var ret ProductAnalyticsBaseQuery
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsAnalyticsQuery) GetQueryOk() (*ProductAnalyticsBaseQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *ProductAnalyticsAnalyticsQuery) SetQuery(v ProductAnalyticsBaseQuery) {
	o.Query = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsAnalyticsQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AudienceFilters != nil {
		toSerialize["audience_filters"] = o.AudienceFilters
	}
	toSerialize["compute"] = o.Compute
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	if o.Indexes != nil {
		toSerialize["indexes"] = o.Indexes
	}
	toSerialize["query"] = o.Query

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsAnalyticsQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AudienceFilters *ProductAnalyticsAudienceFilters `json:"audience_filters,omitempty"`
		Compute         *ProductAnalyticsCompute         `json:"compute"`
		GroupBy         []ProductAnalyticsGroupBy        `json:"group_by,omitempty"`
		Indexes         []string                         `json:"indexes,omitempty"`
		Query           *ProductAnalyticsBaseQuery       `json:"query"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Compute == nil {
		return fmt.Errorf("required field compute missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"audience_filters", "compute", "group_by", "indexes", "query"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AudienceFilters != nil && all.AudienceFilters.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AudienceFilters = all.AudienceFilters
	if all.Compute.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Compute = *all.Compute
	o.GroupBy = all.GroupBy
	o.Indexes = all.Indexes
	o.Query = *all.Query

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
