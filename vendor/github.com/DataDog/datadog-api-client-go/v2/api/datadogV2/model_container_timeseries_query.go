// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ContainerTimeseriesQuery A query for container-level metrics such as CPU and memory usage.
type ContainerTimeseriesQuery struct {
	// A data source for container-level infrastructure metrics.
	DataSource ContainerDataSource `json:"data_source"`
	// Whether CPU metrics should be normalized by core count.
	IsNormalizedCpu *bool `json:"is_normalized_cpu,omitempty"`
	// Maximum number of results to return.
	Limit *int64 `json:"limit,omitempty"`
	// The container metric to query.
	Metric string `json:"metric"`
	// The variable name for use in formulas.
	Name string `json:"name"`
	// Direction of sort.
	Sort *QuerySortOrder `json:"sort,omitempty"`
	// Tag filters to narrow down containers.
	TagFilters []string `json:"tag_filters,omitempty"`
	// A full-text search filter to match container names.
	TextFilter *string `json:"text_filter,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewContainerTimeseriesQuery instantiates a new ContainerTimeseriesQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewContainerTimeseriesQuery(dataSource ContainerDataSource, metric string, name string) *ContainerTimeseriesQuery {
	this := ContainerTimeseriesQuery{}
	this.DataSource = dataSource
	this.Metric = metric
	this.Name = name
	var sort QuerySortOrder = QUERYSORTORDER_DESC
	this.Sort = &sort
	return &this
}

// NewContainerTimeseriesQueryWithDefaults instantiates a new ContainerTimeseriesQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewContainerTimeseriesQueryWithDefaults() *ContainerTimeseriesQuery {
	this := ContainerTimeseriesQuery{}
	var dataSource ContainerDataSource = CONTAINERDATASOURCE_CONTAINER
	this.DataSource = dataSource
	var sort QuerySortOrder = QUERYSORTORDER_DESC
	this.Sort = &sort
	return &this
}

// GetDataSource returns the DataSource field value.
func (o *ContainerTimeseriesQuery) GetDataSource() ContainerDataSource {
	if o == nil {
		var ret ContainerDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *ContainerTimeseriesQuery) GetDataSourceOk() (*ContainerDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *ContainerTimeseriesQuery) SetDataSource(v ContainerDataSource) {
	o.DataSource = v
}

// GetIsNormalizedCpu returns the IsNormalizedCpu field value if set, zero value otherwise.
func (o *ContainerTimeseriesQuery) GetIsNormalizedCpu() bool {
	if o == nil || o.IsNormalizedCpu == nil {
		var ret bool
		return ret
	}
	return *o.IsNormalizedCpu
}

// GetIsNormalizedCpuOk returns a tuple with the IsNormalizedCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerTimeseriesQuery) GetIsNormalizedCpuOk() (*bool, bool) {
	if o == nil || o.IsNormalizedCpu == nil {
		return nil, false
	}
	return o.IsNormalizedCpu, true
}

// HasIsNormalizedCpu returns a boolean if a field has been set.
func (o *ContainerTimeseriesQuery) HasIsNormalizedCpu() bool {
	return o != nil && o.IsNormalizedCpu != nil
}

// SetIsNormalizedCpu gets a reference to the given bool and assigns it to the IsNormalizedCpu field.
func (o *ContainerTimeseriesQuery) SetIsNormalizedCpu(v bool) {
	o.IsNormalizedCpu = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ContainerTimeseriesQuery) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerTimeseriesQuery) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ContainerTimeseriesQuery) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *ContainerTimeseriesQuery) SetLimit(v int64) {
	o.Limit = &v
}

// GetMetric returns the Metric field value.
func (o *ContainerTimeseriesQuery) GetMetric() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Metric
}

// GetMetricOk returns a tuple with the Metric field value
// and a boolean to check if the value has been set.
func (o *ContainerTimeseriesQuery) GetMetricOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metric, true
}

// SetMetric sets field value.
func (o *ContainerTimeseriesQuery) SetMetric(v string) {
	o.Metric = v
}

// GetName returns the Name field value.
func (o *ContainerTimeseriesQuery) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContainerTimeseriesQuery) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ContainerTimeseriesQuery) SetName(v string) {
	o.Name = v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *ContainerTimeseriesQuery) GetSort() QuerySortOrder {
	if o == nil || o.Sort == nil {
		var ret QuerySortOrder
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerTimeseriesQuery) GetSortOk() (*QuerySortOrder, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *ContainerTimeseriesQuery) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given QuerySortOrder and assigns it to the Sort field.
func (o *ContainerTimeseriesQuery) SetSort(v QuerySortOrder) {
	o.Sort = &v
}

// GetTagFilters returns the TagFilters field value if set, zero value otherwise.
func (o *ContainerTimeseriesQuery) GetTagFilters() []string {
	if o == nil || o.TagFilters == nil {
		var ret []string
		return ret
	}
	return o.TagFilters
}

// GetTagFiltersOk returns a tuple with the TagFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerTimeseriesQuery) GetTagFiltersOk() (*[]string, bool) {
	if o == nil || o.TagFilters == nil {
		return nil, false
	}
	return &o.TagFilters, true
}

// HasTagFilters returns a boolean if a field has been set.
func (o *ContainerTimeseriesQuery) HasTagFilters() bool {
	return o != nil && o.TagFilters != nil
}

// SetTagFilters gets a reference to the given []string and assigns it to the TagFilters field.
func (o *ContainerTimeseriesQuery) SetTagFilters(v []string) {
	o.TagFilters = v
}

// GetTextFilter returns the TextFilter field value if set, zero value otherwise.
func (o *ContainerTimeseriesQuery) GetTextFilter() string {
	if o == nil || o.TextFilter == nil {
		var ret string
		return ret
	}
	return *o.TextFilter
}

// GetTextFilterOk returns a tuple with the TextFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerTimeseriesQuery) GetTextFilterOk() (*string, bool) {
	if o == nil || o.TextFilter == nil {
		return nil, false
	}
	return o.TextFilter, true
}

// HasTextFilter returns a boolean if a field has been set.
func (o *ContainerTimeseriesQuery) HasTextFilter() bool {
	return o != nil && o.TextFilter != nil
}

// SetTextFilter gets a reference to the given string and assigns it to the TextFilter field.
func (o *ContainerTimeseriesQuery) SetTextFilter(v string) {
	o.TextFilter = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ContainerTimeseriesQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data_source"] = o.DataSource
	if o.IsNormalizedCpu != nil {
		toSerialize["is_normalized_cpu"] = o.IsNormalizedCpu
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	toSerialize["metric"] = o.Metric
	toSerialize["name"] = o.Name
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	if o.TagFilters != nil {
		toSerialize["tag_filters"] = o.TagFilters
	}
	if o.TextFilter != nil {
		toSerialize["text_filter"] = o.TextFilter
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ContainerTimeseriesQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataSource      *ContainerDataSource `json:"data_source"`
		IsNormalizedCpu *bool                `json:"is_normalized_cpu,omitempty"`
		Limit           *int64               `json:"limit,omitempty"`
		Metric          *string              `json:"metric"`
		Name            *string              `json:"name"`
		Sort            *QuerySortOrder      `json:"sort,omitempty"`
		TagFilters      []string             `json:"tag_filters,omitempty"`
		TextFilter      *string              `json:"text_filter,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.Metric == nil {
		return fmt.Errorf("required field metric missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data_source", "is_normalized_cpu", "limit", "metric", "name", "sort", "tag_filters", "text_filter"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.IsNormalizedCpu = all.IsNormalizedCpu
	o.Limit = all.Limit
	o.Metric = *all.Metric
	o.Name = *all.Name
	if all.Sort != nil && !all.Sort.IsValid() {
		hasInvalidField = true
	} else {
		o.Sort = all.Sort
	}
	o.TagFilters = all.TagFilters
	o.TextFilter = all.TextFilter

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
