// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RetentionQuery Retention query definition.
type RetentionQuery struct {
	// Compute configuration for retention queries.
	Compute RetentionCompute `json:"compute"`
	// Data source for retention queries.
	DataSource RetentionDataSource `json:"data_source"`
	// Filters for retention queries.
	Filters *RetentionFilters `json:"filters,omitempty"`
	// Group by configuration.
	GroupBy []RetentionGroupBy `json:"group_by,omitempty"`
	// Name of the query.
	Name *string `json:"name,omitempty"`
	// Search configuration for retention queries.
	Search RetentionSearch `json:"search"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewRetentionQuery instantiates a new RetentionQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRetentionQuery(compute RetentionCompute, dataSource RetentionDataSource, search RetentionSearch) *RetentionQuery {
	this := RetentionQuery{}
	this.Compute = compute
	this.DataSource = dataSource
	this.Search = search
	return &this
}

// NewRetentionQueryWithDefaults instantiates a new RetentionQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRetentionQueryWithDefaults() *RetentionQuery {
	this := RetentionQuery{}
	return &this
}

// GetCompute returns the Compute field value.
func (o *RetentionQuery) GetCompute() RetentionCompute {
	if o == nil {
		var ret RetentionCompute
		return ret
	}
	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value
// and a boolean to check if the value has been set.
func (o *RetentionQuery) GetComputeOk() (*RetentionCompute, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Compute, true
}

// SetCompute sets field value.
func (o *RetentionQuery) SetCompute(v RetentionCompute) {
	o.Compute = v
}

// GetDataSource returns the DataSource field value.
func (o *RetentionQuery) GetDataSource() RetentionDataSource {
	if o == nil {
		var ret RetentionDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *RetentionQuery) GetDataSourceOk() (*RetentionDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *RetentionQuery) SetDataSource(v RetentionDataSource) {
	o.DataSource = v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *RetentionQuery) GetFilters() RetentionFilters {
	if o == nil || o.Filters == nil {
		var ret RetentionFilters
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionQuery) GetFiltersOk() (*RetentionFilters, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *RetentionQuery) HasFilters() bool {
	return o != nil && o.Filters != nil
}

// SetFilters gets a reference to the given RetentionFilters and assigns it to the Filters field.
func (o *RetentionQuery) SetFilters(v RetentionFilters) {
	o.Filters = &v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *RetentionQuery) GetGroupBy() []RetentionGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []RetentionGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionQuery) GetGroupByOk() (*[]RetentionGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *RetentionQuery) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []RetentionGroupBy and assigns it to the GroupBy field.
func (o *RetentionQuery) SetGroupBy(v []RetentionGroupBy) {
	o.GroupBy = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RetentionQuery) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionQuery) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RetentionQuery) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RetentionQuery) SetName(v string) {
	o.Name = &v
}

// GetSearch returns the Search field value.
func (o *RetentionQuery) GetSearch() RetentionSearch {
	if o == nil {
		var ret RetentionSearch
		return ret
	}
	return o.Search
}

// GetSearchOk returns a tuple with the Search field value
// and a boolean to check if the value has been set.
func (o *RetentionQuery) GetSearchOk() (*RetentionSearch, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Search, true
}

// SetSearch sets field value.
func (o *RetentionQuery) SetSearch(v RetentionSearch) {
	o.Search = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RetentionQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["compute"] = o.Compute
	toSerialize["data_source"] = o.DataSource
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	toSerialize["search"] = o.Search
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RetentionQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Compute    *RetentionCompute    `json:"compute"`
		DataSource *RetentionDataSource `json:"data_source"`
		Filters    *RetentionFilters    `json:"filters,omitempty"`
		GroupBy    []RetentionGroupBy   `json:"group_by,omitempty"`
		Name       *string              `json:"name,omitempty"`
		Search     *RetentionSearch     `json:"search"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Compute == nil {
		return fmt.Errorf("required field compute missing")
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.Search == nil {
		return fmt.Errorf("required field search missing")
	}

	hasInvalidField := false
	if all.Compute.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Compute = *all.Compute
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	if all.Filters != nil && all.Filters.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Filters = all.Filters
	o.GroupBy = all.GroupBy
	o.Name = all.Name
	if all.Search.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Search = *all.Search

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
