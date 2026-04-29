// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsFunnelQuery User journey funnel query definition.
type ProductAnalyticsFunnelQuery struct {
	// Compute configuration for user journey funnel.
	Compute *ProductAnalyticsFunnelCompute `json:"compute,omitempty"`
	// Data source for user journey funnel queries.
	DataSource ProductAnalyticsFunnelDataSource `json:"data_source"`
	// Group by configuration.
	GroupBy []ProductAnalyticsFunnelGroupBy `json:"group_by,omitempty"`
	// User journey search configuration.
	Search UserJourneySearch `json:"search"`
	// Subquery ID.
	SubqueryId *string `json:"subquery_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewProductAnalyticsFunnelQuery instantiates a new ProductAnalyticsFunnelQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsFunnelQuery(dataSource ProductAnalyticsFunnelDataSource, search UserJourneySearch) *ProductAnalyticsFunnelQuery {
	this := ProductAnalyticsFunnelQuery{}
	this.DataSource = dataSource
	this.Search = search
	return &this
}

// NewProductAnalyticsFunnelQueryWithDefaults instantiates a new ProductAnalyticsFunnelQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsFunnelQueryWithDefaults() *ProductAnalyticsFunnelQuery {
	this := ProductAnalyticsFunnelQuery{}
	return &this
}

// GetCompute returns the Compute field value if set, zero value otherwise.
func (o *ProductAnalyticsFunnelQuery) GetCompute() ProductAnalyticsFunnelCompute {
	if o == nil || o.Compute == nil {
		var ret ProductAnalyticsFunnelCompute
		return ret
	}
	return *o.Compute
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFunnelQuery) GetComputeOk() (*ProductAnalyticsFunnelCompute, bool) {
	if o == nil || o.Compute == nil {
		return nil, false
	}
	return o.Compute, true
}

// HasCompute returns a boolean if a field has been set.
func (o *ProductAnalyticsFunnelQuery) HasCompute() bool {
	return o != nil && o.Compute != nil
}

// SetCompute gets a reference to the given ProductAnalyticsFunnelCompute and assigns it to the Compute field.
func (o *ProductAnalyticsFunnelQuery) SetCompute(v ProductAnalyticsFunnelCompute) {
	o.Compute = &v
}

// GetDataSource returns the DataSource field value.
func (o *ProductAnalyticsFunnelQuery) GetDataSource() ProductAnalyticsFunnelDataSource {
	if o == nil {
		var ret ProductAnalyticsFunnelDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFunnelQuery) GetDataSourceOk() (*ProductAnalyticsFunnelDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *ProductAnalyticsFunnelQuery) SetDataSource(v ProductAnalyticsFunnelDataSource) {
	o.DataSource = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *ProductAnalyticsFunnelQuery) GetGroupBy() []ProductAnalyticsFunnelGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []ProductAnalyticsFunnelGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFunnelQuery) GetGroupByOk() (*[]ProductAnalyticsFunnelGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *ProductAnalyticsFunnelQuery) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []ProductAnalyticsFunnelGroupBy and assigns it to the GroupBy field.
func (o *ProductAnalyticsFunnelQuery) SetGroupBy(v []ProductAnalyticsFunnelGroupBy) {
	o.GroupBy = v
}

// GetSearch returns the Search field value.
func (o *ProductAnalyticsFunnelQuery) GetSearch() UserJourneySearch {
	if o == nil {
		var ret UserJourneySearch
		return ret
	}
	return o.Search
}

// GetSearchOk returns a tuple with the Search field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFunnelQuery) GetSearchOk() (*UserJourneySearch, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Search, true
}

// SetSearch sets field value.
func (o *ProductAnalyticsFunnelQuery) SetSearch(v UserJourneySearch) {
	o.Search = v
}

// GetSubqueryId returns the SubqueryId field value if set, zero value otherwise.
func (o *ProductAnalyticsFunnelQuery) GetSubqueryId() string {
	if o == nil || o.SubqueryId == nil {
		var ret string
		return ret
	}
	return *o.SubqueryId
}

// GetSubqueryIdOk returns a tuple with the SubqueryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFunnelQuery) GetSubqueryIdOk() (*string, bool) {
	if o == nil || o.SubqueryId == nil {
		return nil, false
	}
	return o.SubqueryId, true
}

// HasSubqueryId returns a boolean if a field has been set.
func (o *ProductAnalyticsFunnelQuery) HasSubqueryId() bool {
	return o != nil && o.SubqueryId != nil
}

// SetSubqueryId gets a reference to the given string and assigns it to the SubqueryId field.
func (o *ProductAnalyticsFunnelQuery) SetSubqueryId(v string) {
	o.SubqueryId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsFunnelQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Compute != nil {
		toSerialize["compute"] = o.Compute
	}
	toSerialize["data_source"] = o.DataSource
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	toSerialize["search"] = o.Search
	if o.SubqueryId != nil {
		toSerialize["subquery_id"] = o.SubqueryId
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsFunnelQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Compute    *ProductAnalyticsFunnelCompute    `json:"compute,omitempty"`
		DataSource *ProductAnalyticsFunnelDataSource `json:"data_source"`
		GroupBy    []ProductAnalyticsFunnelGroupBy   `json:"group_by,omitempty"`
		Search     *UserJourneySearch                `json:"search"`
		SubqueryId *string                           `json:"subquery_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.Search == nil {
		return fmt.Errorf("required field search missing")
	}

	hasInvalidField := false
	if all.Compute != nil && all.Compute.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Compute = all.Compute
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.GroupBy = all.GroupBy
	if all.Search.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Search = *all.Search
	o.SubqueryId = all.SubqueryId

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
