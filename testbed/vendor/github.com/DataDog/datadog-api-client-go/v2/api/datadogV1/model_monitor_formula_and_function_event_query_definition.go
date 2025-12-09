// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorFormulaAndFunctionEventQueryDefinition A formula and functions events query.
type MonitorFormulaAndFunctionEventQueryDefinition struct {
	// Compute options.
	Compute MonitorFormulaAndFunctionEventQueryDefinitionCompute `json:"compute"`
	// Data source for event platform-based queries.
	DataSource MonitorFormulaAndFunctionEventsDataSource `json:"data_source"`
	// Group by options.
	GroupBy []MonitorFormulaAndFunctionEventQueryGroupBy `json:"group_by,omitempty"`
	// An array of index names to query in the stream. Omit or use `[]` to query all indexes at once.
	Indexes []string `json:"indexes,omitempty"`
	// Name of the query for use in formulas.
	Name string `json:"name"`
	// Search options.
	Search *MonitorFormulaAndFunctionEventQueryDefinitionSearch `json:"search,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMonitorFormulaAndFunctionEventQueryDefinition instantiates a new MonitorFormulaAndFunctionEventQueryDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorFormulaAndFunctionEventQueryDefinition(compute MonitorFormulaAndFunctionEventQueryDefinitionCompute, dataSource MonitorFormulaAndFunctionEventsDataSource, name string) *MonitorFormulaAndFunctionEventQueryDefinition {
	this := MonitorFormulaAndFunctionEventQueryDefinition{}
	this.Compute = compute
	this.DataSource = dataSource
	this.Name = name
	return &this
}

// NewMonitorFormulaAndFunctionEventQueryDefinitionWithDefaults instantiates a new MonitorFormulaAndFunctionEventQueryDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorFormulaAndFunctionEventQueryDefinitionWithDefaults() *MonitorFormulaAndFunctionEventQueryDefinition {
	this := MonitorFormulaAndFunctionEventQueryDefinition{}
	return &this
}

// GetCompute returns the Compute field value.
func (o *MonitorFormulaAndFunctionEventQueryDefinition) GetCompute() MonitorFormulaAndFunctionEventQueryDefinitionCompute {
	if o == nil {
		var ret MonitorFormulaAndFunctionEventQueryDefinitionCompute
		return ret
	}
	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionEventQueryDefinition) GetComputeOk() (*MonitorFormulaAndFunctionEventQueryDefinitionCompute, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Compute, true
}

// SetCompute sets field value.
func (o *MonitorFormulaAndFunctionEventQueryDefinition) SetCompute(v MonitorFormulaAndFunctionEventQueryDefinitionCompute) {
	o.Compute = v
}

// GetDataSource returns the DataSource field value.
func (o *MonitorFormulaAndFunctionEventQueryDefinition) GetDataSource() MonitorFormulaAndFunctionEventsDataSource {
	if o == nil {
		var ret MonitorFormulaAndFunctionEventsDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionEventQueryDefinition) GetDataSourceOk() (*MonitorFormulaAndFunctionEventsDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *MonitorFormulaAndFunctionEventQueryDefinition) SetDataSource(v MonitorFormulaAndFunctionEventsDataSource) {
	o.DataSource = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *MonitorFormulaAndFunctionEventQueryDefinition) GetGroupBy() []MonitorFormulaAndFunctionEventQueryGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []MonitorFormulaAndFunctionEventQueryGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionEventQueryDefinition) GetGroupByOk() (*[]MonitorFormulaAndFunctionEventQueryGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *MonitorFormulaAndFunctionEventQueryDefinition) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []MonitorFormulaAndFunctionEventQueryGroupBy and assigns it to the GroupBy field.
func (o *MonitorFormulaAndFunctionEventQueryDefinition) SetGroupBy(v []MonitorFormulaAndFunctionEventQueryGroupBy) {
	o.GroupBy = v
}

// GetIndexes returns the Indexes field value if set, zero value otherwise.
func (o *MonitorFormulaAndFunctionEventQueryDefinition) GetIndexes() []string {
	if o == nil || o.Indexes == nil {
		var ret []string
		return ret
	}
	return o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionEventQueryDefinition) GetIndexesOk() (*[]string, bool) {
	if o == nil || o.Indexes == nil {
		return nil, false
	}
	return &o.Indexes, true
}

// HasIndexes returns a boolean if a field has been set.
func (o *MonitorFormulaAndFunctionEventQueryDefinition) HasIndexes() bool {
	return o != nil && o.Indexes != nil
}

// SetIndexes gets a reference to the given []string and assigns it to the Indexes field.
func (o *MonitorFormulaAndFunctionEventQueryDefinition) SetIndexes(v []string) {
	o.Indexes = v
}

// GetName returns the Name field value.
func (o *MonitorFormulaAndFunctionEventQueryDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionEventQueryDefinition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *MonitorFormulaAndFunctionEventQueryDefinition) SetName(v string) {
	o.Name = v
}

// GetSearch returns the Search field value if set, zero value otherwise.
func (o *MonitorFormulaAndFunctionEventQueryDefinition) GetSearch() MonitorFormulaAndFunctionEventQueryDefinitionSearch {
	if o == nil || o.Search == nil {
		var ret MonitorFormulaAndFunctionEventQueryDefinitionSearch
		return ret
	}
	return *o.Search
}

// GetSearchOk returns a tuple with the Search field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionEventQueryDefinition) GetSearchOk() (*MonitorFormulaAndFunctionEventQueryDefinitionSearch, bool) {
	if o == nil || o.Search == nil {
		return nil, false
	}
	return o.Search, true
}

// HasSearch returns a boolean if a field has been set.
func (o *MonitorFormulaAndFunctionEventQueryDefinition) HasSearch() bool {
	return o != nil && o.Search != nil
}

// SetSearch gets a reference to the given MonitorFormulaAndFunctionEventQueryDefinitionSearch and assigns it to the Search field.
func (o *MonitorFormulaAndFunctionEventQueryDefinition) SetSearch(v MonitorFormulaAndFunctionEventQueryDefinitionSearch) {
	o.Search = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorFormulaAndFunctionEventQueryDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
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
	if o.Search != nil {
		toSerialize["search"] = o.Search
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorFormulaAndFunctionEventQueryDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Compute    *MonitorFormulaAndFunctionEventQueryDefinitionCompute `json:"compute"`
		DataSource *MonitorFormulaAndFunctionEventsDataSource            `json:"data_source"`
		GroupBy    []MonitorFormulaAndFunctionEventQueryGroupBy          `json:"group_by,omitempty"`
		Indexes    []string                                              `json:"indexes,omitempty"`
		Name       *string                                               `json:"name"`
		Search     *MonitorFormulaAndFunctionEventQueryDefinitionSearch  `json:"search,omitempty"`
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
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"compute", "data_source", "group_by", "indexes", "name", "search"})
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
	o.Indexes = all.Indexes
	o.Name = *all.Name
	if all.Search != nil && all.Search.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Search = all.Search

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
