// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SankeyNetworkQuery Query configuration for Sankey network widget.
type SankeyNetworkQuery struct {
	// Compute aggregation for network queries.
	Compute *SankeyNetworkQueryCompute `json:"compute,omitempty"`
	// Network data source type.
	DataSource SankeyNetworkDataSource `json:"data_source"`
	// Fields to group by.
	GroupBy []string `json:"group_by"`
	// Maximum number of results.
	Limit int64 `json:"limit"`
	// Sankey mode for network queries.
	Mode *SankeyNetworkQueryMode `json:"mode,omitempty"`
	// Query string for filtering network data.
	QueryString string `json:"query_string"`
	// Whether to exclude missing values.
	ShouldExcludeMissing *bool `json:"should_exclude_missing,omitempty"`
	// Sort configuration for network queries.
	Sort *SankeyNetworkQuerySort `json:"sort,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewSankeyNetworkQuery instantiates a new SankeyNetworkQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSankeyNetworkQuery(dataSource SankeyNetworkDataSource, groupBy []string, limit int64, queryString string) *SankeyNetworkQuery {
	this := SankeyNetworkQuery{}
	this.DataSource = dataSource
	this.GroupBy = groupBy
	this.Limit = limit
	var mode SankeyNetworkQueryMode = SANKEYNETWORKQUERYMODE_TARGET
	this.Mode = &mode
	this.QueryString = queryString
	return &this
}

// NewSankeyNetworkQueryWithDefaults instantiates a new SankeyNetworkQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSankeyNetworkQueryWithDefaults() *SankeyNetworkQuery {
	this := SankeyNetworkQuery{}
	var dataSource SankeyNetworkDataSource = SANKEYNETWORKDATASOURCE_NETWORK
	this.DataSource = dataSource
	var mode SankeyNetworkQueryMode = SANKEYNETWORKQUERYMODE_TARGET
	this.Mode = &mode
	return &this
}

// GetCompute returns the Compute field value if set, zero value otherwise.
func (o *SankeyNetworkQuery) GetCompute() SankeyNetworkQueryCompute {
	if o == nil || o.Compute == nil {
		var ret SankeyNetworkQueryCompute
		return ret
	}
	return *o.Compute
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyNetworkQuery) GetComputeOk() (*SankeyNetworkQueryCompute, bool) {
	if o == nil || o.Compute == nil {
		return nil, false
	}
	return o.Compute, true
}

// HasCompute returns a boolean if a field has been set.
func (o *SankeyNetworkQuery) HasCompute() bool {
	return o != nil && o.Compute != nil
}

// SetCompute gets a reference to the given SankeyNetworkQueryCompute and assigns it to the Compute field.
func (o *SankeyNetworkQuery) SetCompute(v SankeyNetworkQueryCompute) {
	o.Compute = &v
}

// GetDataSource returns the DataSource field value.
func (o *SankeyNetworkQuery) GetDataSource() SankeyNetworkDataSource {
	if o == nil {
		var ret SankeyNetworkDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *SankeyNetworkQuery) GetDataSourceOk() (*SankeyNetworkDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *SankeyNetworkQuery) SetDataSource(v SankeyNetworkDataSource) {
	o.DataSource = v
}

// GetGroupBy returns the GroupBy field value.
func (o *SankeyNetworkQuery) GetGroupBy() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value
// and a boolean to check if the value has been set.
func (o *SankeyNetworkQuery) GetGroupByOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// SetGroupBy sets field value.
func (o *SankeyNetworkQuery) SetGroupBy(v []string) {
	o.GroupBy = v
}

// GetLimit returns the Limit field value.
func (o *SankeyNetworkQuery) GetLimit() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *SankeyNetworkQuery) GetLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value.
func (o *SankeyNetworkQuery) SetLimit(v int64) {
	o.Limit = v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *SankeyNetworkQuery) GetMode() SankeyNetworkQueryMode {
	if o == nil || o.Mode == nil {
		var ret SankeyNetworkQueryMode
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyNetworkQuery) GetModeOk() (*SankeyNetworkQueryMode, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *SankeyNetworkQuery) HasMode() bool {
	return o != nil && o.Mode != nil
}

// SetMode gets a reference to the given SankeyNetworkQueryMode and assigns it to the Mode field.
func (o *SankeyNetworkQuery) SetMode(v SankeyNetworkQueryMode) {
	o.Mode = &v
}

// GetQueryString returns the QueryString field value.
func (o *SankeyNetworkQuery) GetQueryString() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value
// and a boolean to check if the value has been set.
func (o *SankeyNetworkQuery) GetQueryStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryString, true
}

// SetQueryString sets field value.
func (o *SankeyNetworkQuery) SetQueryString(v string) {
	o.QueryString = v
}

// GetShouldExcludeMissing returns the ShouldExcludeMissing field value if set, zero value otherwise.
func (o *SankeyNetworkQuery) GetShouldExcludeMissing() bool {
	if o == nil || o.ShouldExcludeMissing == nil {
		var ret bool
		return ret
	}
	return *o.ShouldExcludeMissing
}

// GetShouldExcludeMissingOk returns a tuple with the ShouldExcludeMissing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyNetworkQuery) GetShouldExcludeMissingOk() (*bool, bool) {
	if o == nil || o.ShouldExcludeMissing == nil {
		return nil, false
	}
	return o.ShouldExcludeMissing, true
}

// HasShouldExcludeMissing returns a boolean if a field has been set.
func (o *SankeyNetworkQuery) HasShouldExcludeMissing() bool {
	return o != nil && o.ShouldExcludeMissing != nil
}

// SetShouldExcludeMissing gets a reference to the given bool and assigns it to the ShouldExcludeMissing field.
func (o *SankeyNetworkQuery) SetShouldExcludeMissing(v bool) {
	o.ShouldExcludeMissing = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *SankeyNetworkQuery) GetSort() SankeyNetworkQuerySort {
	if o == nil || o.Sort == nil {
		var ret SankeyNetworkQuerySort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyNetworkQuery) GetSortOk() (*SankeyNetworkQuerySort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *SankeyNetworkQuery) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given SankeyNetworkQuerySort and assigns it to the Sort field.
func (o *SankeyNetworkQuery) SetSort(v SankeyNetworkQuerySort) {
	o.Sort = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SankeyNetworkQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Compute != nil {
		toSerialize["compute"] = o.Compute
	}
	toSerialize["data_source"] = o.DataSource
	toSerialize["group_by"] = o.GroupBy
	toSerialize["limit"] = o.Limit
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	toSerialize["query_string"] = o.QueryString
	if o.ShouldExcludeMissing != nil {
		toSerialize["should_exclude_missing"] = o.ShouldExcludeMissing
	}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SankeyNetworkQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Compute              *SankeyNetworkQueryCompute `json:"compute,omitempty"`
		DataSource           *SankeyNetworkDataSource   `json:"data_source"`
		GroupBy              *[]string                  `json:"group_by"`
		Limit                *int64                     `json:"limit"`
		Mode                 *SankeyNetworkQueryMode    `json:"mode,omitempty"`
		QueryString          *string                    `json:"query_string"`
		ShouldExcludeMissing *bool                      `json:"should_exclude_missing,omitempty"`
		Sort                 *SankeyNetworkQuerySort    `json:"sort,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.GroupBy == nil {
		return fmt.Errorf("required field group_by missing")
	}
	if all.Limit == nil {
		return fmt.Errorf("required field limit missing")
	}
	if all.QueryString == nil {
		return fmt.Errorf("required field query_string missing")
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
	o.GroupBy = *all.GroupBy
	o.Limit = *all.Limit
	if all.Mode != nil && !all.Mode.IsValid() {
		hasInvalidField = true
	} else {
		o.Mode = all.Mode
	}
	o.QueryString = *all.QueryString
	o.ShouldExcludeMissing = all.ShouldExcludeMissing
	if all.Sort != nil && all.Sort.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Sort = all.Sort

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
