// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaAndFunctionProductAnalyticsExtendedQueryDefinition A formula and functions Product Analytics Extended query for advanced analytics features.
type FormulaAndFunctionProductAnalyticsExtendedQueryDefinition struct {
	// Product Analytics/RUM audience filters.
	AudienceFilters *ProductAnalyticsAudienceFilters `json:"audience_filters,omitempty"`
	// Compute configuration for Product Analytics Extended queries.
	Compute ProductAnalyticsExtendedCompute `json:"compute"`
	// Data source for Product Analytics Extended queries.
	DataSource FormulaAndFunctionProductAnalyticsExtendedDataSource `json:"data_source"`
	// Group by configuration.
	GroupBy []ProductAnalyticsExtendedGroupBy `json:"group_by,omitempty"`
	// Event indexes to query.
	Indexes []FormulaAndFunctionProductAnalyticsExtendedQueryDefinitionIndexesItems `json:"indexes,omitempty"`
	// Name of the query for use in formulas.
	Name string `json:"name"`
	// Product Analytics event query.
	Query ProductAnalyticsBaseQuery `json:"query"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFormulaAndFunctionProductAnalyticsExtendedQueryDefinition instantiates a new FormulaAndFunctionProductAnalyticsExtendedQueryDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormulaAndFunctionProductAnalyticsExtendedQueryDefinition(compute ProductAnalyticsExtendedCompute, dataSource FormulaAndFunctionProductAnalyticsExtendedDataSource, name string, query ProductAnalyticsBaseQuery) *FormulaAndFunctionProductAnalyticsExtendedQueryDefinition {
	this := FormulaAndFunctionProductAnalyticsExtendedQueryDefinition{}
	this.Compute = compute
	this.DataSource = dataSource
	this.Name = name
	this.Query = query
	return &this
}

// NewFormulaAndFunctionProductAnalyticsExtendedQueryDefinitionWithDefaults instantiates a new FormulaAndFunctionProductAnalyticsExtendedQueryDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormulaAndFunctionProductAnalyticsExtendedQueryDefinitionWithDefaults() *FormulaAndFunctionProductAnalyticsExtendedQueryDefinition {
	this := FormulaAndFunctionProductAnalyticsExtendedQueryDefinition{}
	return &this
}

// GetAudienceFilters returns the AudienceFilters field value if set, zero value otherwise.
func (o *FormulaAndFunctionProductAnalyticsExtendedQueryDefinition) GetAudienceFilters() ProductAnalyticsAudienceFilters {
	if o == nil || o.AudienceFilters == nil {
		var ret ProductAnalyticsAudienceFilters
		return ret
	}
	return *o.AudienceFilters
}

// GetAudienceFiltersOk returns a tuple with the AudienceFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionProductAnalyticsExtendedQueryDefinition) GetAudienceFiltersOk() (*ProductAnalyticsAudienceFilters, bool) {
	if o == nil || o.AudienceFilters == nil {
		return nil, false
	}
	return o.AudienceFilters, true
}

// HasAudienceFilters returns a boolean if a field has been set.
func (o *FormulaAndFunctionProductAnalyticsExtendedQueryDefinition) HasAudienceFilters() bool {
	return o != nil && o.AudienceFilters != nil
}

// SetAudienceFilters gets a reference to the given ProductAnalyticsAudienceFilters and assigns it to the AudienceFilters field.
func (o *FormulaAndFunctionProductAnalyticsExtendedQueryDefinition) SetAudienceFilters(v ProductAnalyticsAudienceFilters) {
	o.AudienceFilters = &v
}

// GetCompute returns the Compute field value.
func (o *FormulaAndFunctionProductAnalyticsExtendedQueryDefinition) GetCompute() ProductAnalyticsExtendedCompute {
	if o == nil {
		var ret ProductAnalyticsExtendedCompute
		return ret
	}
	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionProductAnalyticsExtendedQueryDefinition) GetComputeOk() (*ProductAnalyticsExtendedCompute, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Compute, true
}

// SetCompute sets field value.
func (o *FormulaAndFunctionProductAnalyticsExtendedQueryDefinition) SetCompute(v ProductAnalyticsExtendedCompute) {
	o.Compute = v
}

// GetDataSource returns the DataSource field value.
func (o *FormulaAndFunctionProductAnalyticsExtendedQueryDefinition) GetDataSource() FormulaAndFunctionProductAnalyticsExtendedDataSource {
	if o == nil {
		var ret FormulaAndFunctionProductAnalyticsExtendedDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionProductAnalyticsExtendedQueryDefinition) GetDataSourceOk() (*FormulaAndFunctionProductAnalyticsExtendedDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *FormulaAndFunctionProductAnalyticsExtendedQueryDefinition) SetDataSource(v FormulaAndFunctionProductAnalyticsExtendedDataSource) {
	o.DataSource = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *FormulaAndFunctionProductAnalyticsExtendedQueryDefinition) GetGroupBy() []ProductAnalyticsExtendedGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []ProductAnalyticsExtendedGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionProductAnalyticsExtendedQueryDefinition) GetGroupByOk() (*[]ProductAnalyticsExtendedGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *FormulaAndFunctionProductAnalyticsExtendedQueryDefinition) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []ProductAnalyticsExtendedGroupBy and assigns it to the GroupBy field.
func (o *FormulaAndFunctionProductAnalyticsExtendedQueryDefinition) SetGroupBy(v []ProductAnalyticsExtendedGroupBy) {
	o.GroupBy = v
}

// GetIndexes returns the Indexes field value if set, zero value otherwise.
func (o *FormulaAndFunctionProductAnalyticsExtendedQueryDefinition) GetIndexes() []FormulaAndFunctionProductAnalyticsExtendedQueryDefinitionIndexesItems {
	if o == nil || o.Indexes == nil {
		var ret []FormulaAndFunctionProductAnalyticsExtendedQueryDefinitionIndexesItems
		return ret
	}
	return o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionProductAnalyticsExtendedQueryDefinition) GetIndexesOk() (*[]FormulaAndFunctionProductAnalyticsExtendedQueryDefinitionIndexesItems, bool) {
	if o == nil || o.Indexes == nil {
		return nil, false
	}
	return &o.Indexes, true
}

// HasIndexes returns a boolean if a field has been set.
func (o *FormulaAndFunctionProductAnalyticsExtendedQueryDefinition) HasIndexes() bool {
	return o != nil && o.Indexes != nil
}

// SetIndexes gets a reference to the given []FormulaAndFunctionProductAnalyticsExtendedQueryDefinitionIndexesItems and assigns it to the Indexes field.
func (o *FormulaAndFunctionProductAnalyticsExtendedQueryDefinition) SetIndexes(v []FormulaAndFunctionProductAnalyticsExtendedQueryDefinitionIndexesItems) {
	o.Indexes = v
}

// GetName returns the Name field value.
func (o *FormulaAndFunctionProductAnalyticsExtendedQueryDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionProductAnalyticsExtendedQueryDefinition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *FormulaAndFunctionProductAnalyticsExtendedQueryDefinition) SetName(v string) {
	o.Name = v
}

// GetQuery returns the Query field value.
func (o *FormulaAndFunctionProductAnalyticsExtendedQueryDefinition) GetQuery() ProductAnalyticsBaseQuery {
	if o == nil {
		var ret ProductAnalyticsBaseQuery
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionProductAnalyticsExtendedQueryDefinition) GetQueryOk() (*ProductAnalyticsBaseQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *FormulaAndFunctionProductAnalyticsExtendedQueryDefinition) SetQuery(v ProductAnalyticsBaseQuery) {
	o.Query = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormulaAndFunctionProductAnalyticsExtendedQueryDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AudienceFilters != nil {
		toSerialize["audience_filters"] = o.AudienceFilters
	}
	toSerialize["compute"] = o.Compute
	toSerialize["data_source"] = o.DataSource
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	if o.Indexes != nil {
		toSerialize["indexes"] = o.Indexes
	}
	toSerialize["name"] = o.Name
	toSerialize["query"] = o.Query

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormulaAndFunctionProductAnalyticsExtendedQueryDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AudienceFilters *ProductAnalyticsAudienceFilters                                        `json:"audience_filters,omitempty"`
		Compute         *ProductAnalyticsExtendedCompute                                        `json:"compute"`
		DataSource      *FormulaAndFunctionProductAnalyticsExtendedDataSource                   `json:"data_source"`
		GroupBy         []ProductAnalyticsExtendedGroupBy                                       `json:"group_by,omitempty"`
		Indexes         []FormulaAndFunctionProductAnalyticsExtendedQueryDefinitionIndexesItems `json:"indexes,omitempty"`
		Name            *string                                                                 `json:"name"`
		Query           *ProductAnalyticsBaseQuery                                              `json:"query"`
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
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"audience_filters", "compute", "data_source", "group_by", "indexes", "name", "query"})
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
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.GroupBy = all.GroupBy
	o.Indexes = all.Indexes
	o.Name = *all.Name
	if all.Query.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Query = *all.Query

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
