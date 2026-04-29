// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaAndFunctionRetentionQueryDefinition A formula and functions Retention query for defining timeseries and scalar visualizations.
type FormulaAndFunctionRetentionQueryDefinition struct {
	// Compute configuration for retention queries.
	Compute RetentionCompute `json:"compute"`
	// Data source for retention queries.
	DataSource RetentionDataSource `json:"data_source"`
	// Group by configuration.
	GroupBy []RetentionGroupBy `json:"group_by,omitempty"`
	// Name of the query for use in formulas.
	Name string `json:"name"`
	// Search configuration for retention queries.
	Search RetentionSearch `json:"search"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFormulaAndFunctionRetentionQueryDefinition instantiates a new FormulaAndFunctionRetentionQueryDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormulaAndFunctionRetentionQueryDefinition(compute RetentionCompute, dataSource RetentionDataSource, name string, search RetentionSearch) *FormulaAndFunctionRetentionQueryDefinition {
	this := FormulaAndFunctionRetentionQueryDefinition{}
	this.Compute = compute
	this.DataSource = dataSource
	this.Name = name
	this.Search = search
	return &this
}

// NewFormulaAndFunctionRetentionQueryDefinitionWithDefaults instantiates a new FormulaAndFunctionRetentionQueryDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormulaAndFunctionRetentionQueryDefinitionWithDefaults() *FormulaAndFunctionRetentionQueryDefinition {
	this := FormulaAndFunctionRetentionQueryDefinition{}
	return &this
}

// GetCompute returns the Compute field value.
func (o *FormulaAndFunctionRetentionQueryDefinition) GetCompute() RetentionCompute {
	if o == nil {
		var ret RetentionCompute
		return ret
	}
	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionRetentionQueryDefinition) GetComputeOk() (*RetentionCompute, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Compute, true
}

// SetCompute sets field value.
func (o *FormulaAndFunctionRetentionQueryDefinition) SetCompute(v RetentionCompute) {
	o.Compute = v
}

// GetDataSource returns the DataSource field value.
func (o *FormulaAndFunctionRetentionQueryDefinition) GetDataSource() RetentionDataSource {
	if o == nil {
		var ret RetentionDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionRetentionQueryDefinition) GetDataSourceOk() (*RetentionDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *FormulaAndFunctionRetentionQueryDefinition) SetDataSource(v RetentionDataSource) {
	o.DataSource = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *FormulaAndFunctionRetentionQueryDefinition) GetGroupBy() []RetentionGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []RetentionGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionRetentionQueryDefinition) GetGroupByOk() (*[]RetentionGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *FormulaAndFunctionRetentionQueryDefinition) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []RetentionGroupBy and assigns it to the GroupBy field.
func (o *FormulaAndFunctionRetentionQueryDefinition) SetGroupBy(v []RetentionGroupBy) {
	o.GroupBy = v
}

// GetName returns the Name field value.
func (o *FormulaAndFunctionRetentionQueryDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionRetentionQueryDefinition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *FormulaAndFunctionRetentionQueryDefinition) SetName(v string) {
	o.Name = v
}

// GetSearch returns the Search field value.
func (o *FormulaAndFunctionRetentionQueryDefinition) GetSearch() RetentionSearch {
	if o == nil {
		var ret RetentionSearch
		return ret
	}
	return o.Search
}

// GetSearchOk returns a tuple with the Search field value
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionRetentionQueryDefinition) GetSearchOk() (*RetentionSearch, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Search, true
}

// SetSearch sets field value.
func (o *FormulaAndFunctionRetentionQueryDefinition) SetSearch(v RetentionSearch) {
	o.Search = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormulaAndFunctionRetentionQueryDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["compute"] = o.Compute
	toSerialize["data_source"] = o.DataSource
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	toSerialize["name"] = o.Name
	toSerialize["search"] = o.Search

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormulaAndFunctionRetentionQueryDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Compute    *RetentionCompute    `json:"compute"`
		DataSource *RetentionDataSource `json:"data_source"`
		GroupBy    []RetentionGroupBy   `json:"group_by,omitempty"`
		Name       *string              `json:"name"`
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
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Search == nil {
		return fmt.Errorf("required field search missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"compute", "data_source", "group_by", "name", "search"})
	} else {
		return err
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
	o.GroupBy = all.GroupBy
	o.Name = *all.Name
	if all.Search.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Search = *all.Search

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
