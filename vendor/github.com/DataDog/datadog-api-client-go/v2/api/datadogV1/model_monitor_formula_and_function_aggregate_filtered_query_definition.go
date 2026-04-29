// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorFormulaAndFunctionAggregateFilteredQueryDefinition A formula and functions aggregate filtered query. Used to filter base query results using data from another source.
type MonitorFormulaAndFunctionAggregateFilteredQueryDefinition struct {
	// Base query for aggregate queries. Can be an events query or a metrics query.
	BaseQuery MonitorFormulaAndFunctionAggregateBaseQuery `json:"base_query"`
	// Compute options for the query.
	Compute []MonitorFormulaAndFunctionEventQueryDefinitionCompute `json:"compute,omitempty"`
	// Data source for aggregate filtered queries.
	DataSource MonitorFormulaAndFunctionAggregateFilteredDataSource `json:"data_source"`
	// Filter query for aggregate filtered queries. Can be an events query or a reference table query.
	FilterQuery MonitorFormulaAndFunctionAggregateFilterQuery `json:"filter_query"`
	// Filter conditions for the query.
	Filters []MonitorFormulaAndFunctionAggregateQueryFilter `json:"filters"`
	// Group by options for the query.
	GroupBy []MonitorFormulaAndFunctionEventQueryGroupBy `json:"group_by,omitempty"`
	// Name of the query for use in formulas.
	Name *string `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewMonitorFormulaAndFunctionAggregateFilteredQueryDefinition instantiates a new MonitorFormulaAndFunctionAggregateFilteredQueryDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorFormulaAndFunctionAggregateFilteredQueryDefinition(baseQuery MonitorFormulaAndFunctionAggregateBaseQuery, dataSource MonitorFormulaAndFunctionAggregateFilteredDataSource, filterQuery MonitorFormulaAndFunctionAggregateFilterQuery, filters []MonitorFormulaAndFunctionAggregateQueryFilter) *MonitorFormulaAndFunctionAggregateFilteredQueryDefinition {
	this := MonitorFormulaAndFunctionAggregateFilteredQueryDefinition{}
	this.BaseQuery = baseQuery
	this.DataSource = dataSource
	this.FilterQuery = filterQuery
	this.Filters = filters
	return &this
}

// NewMonitorFormulaAndFunctionAggregateFilteredQueryDefinitionWithDefaults instantiates a new MonitorFormulaAndFunctionAggregateFilteredQueryDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorFormulaAndFunctionAggregateFilteredQueryDefinitionWithDefaults() *MonitorFormulaAndFunctionAggregateFilteredQueryDefinition {
	this := MonitorFormulaAndFunctionAggregateFilteredQueryDefinition{}
	return &this
}

// GetBaseQuery returns the BaseQuery field value.
func (o *MonitorFormulaAndFunctionAggregateFilteredQueryDefinition) GetBaseQuery() MonitorFormulaAndFunctionAggregateBaseQuery {
	if o == nil {
		var ret MonitorFormulaAndFunctionAggregateBaseQuery
		return ret
	}
	return o.BaseQuery
}

// GetBaseQueryOk returns a tuple with the BaseQuery field value
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionAggregateFilteredQueryDefinition) GetBaseQueryOk() (*MonitorFormulaAndFunctionAggregateBaseQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseQuery, true
}

// SetBaseQuery sets field value.
func (o *MonitorFormulaAndFunctionAggregateFilteredQueryDefinition) SetBaseQuery(v MonitorFormulaAndFunctionAggregateBaseQuery) {
	o.BaseQuery = v
}

// GetCompute returns the Compute field value if set, zero value otherwise.
func (o *MonitorFormulaAndFunctionAggregateFilteredQueryDefinition) GetCompute() []MonitorFormulaAndFunctionEventQueryDefinitionCompute {
	if o == nil || o.Compute == nil {
		var ret []MonitorFormulaAndFunctionEventQueryDefinitionCompute
		return ret
	}
	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionAggregateFilteredQueryDefinition) GetComputeOk() (*[]MonitorFormulaAndFunctionEventQueryDefinitionCompute, bool) {
	if o == nil || o.Compute == nil {
		return nil, false
	}
	return &o.Compute, true
}

// HasCompute returns a boolean if a field has been set.
func (o *MonitorFormulaAndFunctionAggregateFilteredQueryDefinition) HasCompute() bool {
	return o != nil && o.Compute != nil
}

// SetCompute gets a reference to the given []MonitorFormulaAndFunctionEventQueryDefinitionCompute and assigns it to the Compute field.
func (o *MonitorFormulaAndFunctionAggregateFilteredQueryDefinition) SetCompute(v []MonitorFormulaAndFunctionEventQueryDefinitionCompute) {
	o.Compute = v
}

// GetDataSource returns the DataSource field value.
func (o *MonitorFormulaAndFunctionAggregateFilteredQueryDefinition) GetDataSource() MonitorFormulaAndFunctionAggregateFilteredDataSource {
	if o == nil {
		var ret MonitorFormulaAndFunctionAggregateFilteredDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionAggregateFilteredQueryDefinition) GetDataSourceOk() (*MonitorFormulaAndFunctionAggregateFilteredDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *MonitorFormulaAndFunctionAggregateFilteredQueryDefinition) SetDataSource(v MonitorFormulaAndFunctionAggregateFilteredDataSource) {
	o.DataSource = v
}

// GetFilterQuery returns the FilterQuery field value.
func (o *MonitorFormulaAndFunctionAggregateFilteredQueryDefinition) GetFilterQuery() MonitorFormulaAndFunctionAggregateFilterQuery {
	if o == nil {
		var ret MonitorFormulaAndFunctionAggregateFilterQuery
		return ret
	}
	return o.FilterQuery
}

// GetFilterQueryOk returns a tuple with the FilterQuery field value
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionAggregateFilteredQueryDefinition) GetFilterQueryOk() (*MonitorFormulaAndFunctionAggregateFilterQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilterQuery, true
}

// SetFilterQuery sets field value.
func (o *MonitorFormulaAndFunctionAggregateFilteredQueryDefinition) SetFilterQuery(v MonitorFormulaAndFunctionAggregateFilterQuery) {
	o.FilterQuery = v
}

// GetFilters returns the Filters field value.
func (o *MonitorFormulaAndFunctionAggregateFilteredQueryDefinition) GetFilters() []MonitorFormulaAndFunctionAggregateQueryFilter {
	if o == nil {
		var ret []MonitorFormulaAndFunctionAggregateQueryFilter
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionAggregateFilteredQueryDefinition) GetFiltersOk() (*[]MonitorFormulaAndFunctionAggregateQueryFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filters, true
}

// SetFilters sets field value.
func (o *MonitorFormulaAndFunctionAggregateFilteredQueryDefinition) SetFilters(v []MonitorFormulaAndFunctionAggregateQueryFilter) {
	o.Filters = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *MonitorFormulaAndFunctionAggregateFilteredQueryDefinition) GetGroupBy() []MonitorFormulaAndFunctionEventQueryGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []MonitorFormulaAndFunctionEventQueryGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionAggregateFilteredQueryDefinition) GetGroupByOk() (*[]MonitorFormulaAndFunctionEventQueryGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *MonitorFormulaAndFunctionAggregateFilteredQueryDefinition) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []MonitorFormulaAndFunctionEventQueryGroupBy and assigns it to the GroupBy field.
func (o *MonitorFormulaAndFunctionAggregateFilteredQueryDefinition) SetGroupBy(v []MonitorFormulaAndFunctionEventQueryGroupBy) {
	o.GroupBy = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MonitorFormulaAndFunctionAggregateFilteredQueryDefinition) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionAggregateFilteredQueryDefinition) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MonitorFormulaAndFunctionAggregateFilteredQueryDefinition) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MonitorFormulaAndFunctionAggregateFilteredQueryDefinition) SetName(v string) {
	o.Name = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorFormulaAndFunctionAggregateFilteredQueryDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["base_query"] = o.BaseQuery
	if o.Compute != nil {
		toSerialize["compute"] = o.Compute
	}
	toSerialize["data_source"] = o.DataSource
	toSerialize["filter_query"] = o.FilterQuery
	toSerialize["filters"] = o.Filters
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorFormulaAndFunctionAggregateFilteredQueryDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BaseQuery   *MonitorFormulaAndFunctionAggregateBaseQuery           `json:"base_query"`
		Compute     []MonitorFormulaAndFunctionEventQueryDefinitionCompute `json:"compute,omitempty"`
		DataSource  *MonitorFormulaAndFunctionAggregateFilteredDataSource  `json:"data_source"`
		FilterQuery *MonitorFormulaAndFunctionAggregateFilterQuery         `json:"filter_query"`
		Filters     *[]MonitorFormulaAndFunctionAggregateQueryFilter       `json:"filters"`
		GroupBy     []MonitorFormulaAndFunctionEventQueryGroupBy           `json:"group_by,omitempty"`
		Name        *string                                                `json:"name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.BaseQuery == nil {
		return fmt.Errorf("required field base_query missing")
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.FilterQuery == nil {
		return fmt.Errorf("required field filter_query missing")
	}
	if all.Filters == nil {
		return fmt.Errorf("required field filters missing")
	}

	hasInvalidField := false
	o.BaseQuery = *all.BaseQuery
	o.Compute = all.Compute
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.FilterQuery = *all.FilterQuery
	o.Filters = *all.Filters
	o.GroupBy = all.GroupBy
	o.Name = all.Name

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
