// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaAndFunctionEventQueryDefinition A formula and functions events query.
type FormulaAndFunctionEventQueryDefinition struct {
	// Compute options.
	Compute FormulaAndFunctionEventQueryDefinitionCompute `json:"compute"`
	// The source organization UUID for cross organization queries. Feature in Private Beta.
	CrossOrgUuids []string `json:"cross_org_uuids,omitempty"`
	// Data source for event platform-based queries.
	DataSource FormulaAndFunctionEventsDataSource `json:"data_source"`
	// Group by options.
	GroupBy []FormulaAndFunctionEventQueryGroupBy `json:"group_by,omitempty"`
	// An array of index names to query in the stream. Omit or use `[]` to query all indexes at once.
	Indexes []string `json:"indexes,omitempty"`
	// Name of the query for use in formulas.
	Name string `json:"name"`
	// Search options.
	Search *FormulaAndFunctionEventQueryDefinitionSearch `json:"search,omitempty"`
	// Option for storage location. Feature in Private Beta.
	Storage *string `json:"storage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFormulaAndFunctionEventQueryDefinition instantiates a new FormulaAndFunctionEventQueryDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormulaAndFunctionEventQueryDefinition(compute FormulaAndFunctionEventQueryDefinitionCompute, dataSource FormulaAndFunctionEventsDataSource, name string) *FormulaAndFunctionEventQueryDefinition {
	this := FormulaAndFunctionEventQueryDefinition{}
	this.Compute = compute
	this.DataSource = dataSource
	this.Name = name
	return &this
}

// NewFormulaAndFunctionEventQueryDefinitionWithDefaults instantiates a new FormulaAndFunctionEventQueryDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormulaAndFunctionEventQueryDefinitionWithDefaults() *FormulaAndFunctionEventQueryDefinition {
	this := FormulaAndFunctionEventQueryDefinition{}
	return &this
}

// GetCompute returns the Compute field value.
func (o *FormulaAndFunctionEventQueryDefinition) GetCompute() FormulaAndFunctionEventQueryDefinitionCompute {
	if o == nil {
		var ret FormulaAndFunctionEventQueryDefinitionCompute
		return ret
	}
	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionEventQueryDefinition) GetComputeOk() (*FormulaAndFunctionEventQueryDefinitionCompute, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Compute, true
}

// SetCompute sets field value.
func (o *FormulaAndFunctionEventQueryDefinition) SetCompute(v FormulaAndFunctionEventQueryDefinitionCompute) {
	o.Compute = v
}

// GetCrossOrgUuids returns the CrossOrgUuids field value if set, zero value otherwise.
func (o *FormulaAndFunctionEventQueryDefinition) GetCrossOrgUuids() []string {
	if o == nil || o.CrossOrgUuids == nil {
		var ret []string
		return ret
	}
	return o.CrossOrgUuids
}

// GetCrossOrgUuidsOk returns a tuple with the CrossOrgUuids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionEventQueryDefinition) GetCrossOrgUuidsOk() (*[]string, bool) {
	if o == nil || o.CrossOrgUuids == nil {
		return nil, false
	}
	return &o.CrossOrgUuids, true
}

// HasCrossOrgUuids returns a boolean if a field has been set.
func (o *FormulaAndFunctionEventQueryDefinition) HasCrossOrgUuids() bool {
	return o != nil && o.CrossOrgUuids != nil
}

// SetCrossOrgUuids gets a reference to the given []string and assigns it to the CrossOrgUuids field.
func (o *FormulaAndFunctionEventQueryDefinition) SetCrossOrgUuids(v []string) {
	o.CrossOrgUuids = v
}

// GetDataSource returns the DataSource field value.
func (o *FormulaAndFunctionEventQueryDefinition) GetDataSource() FormulaAndFunctionEventsDataSource {
	if o == nil {
		var ret FormulaAndFunctionEventsDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionEventQueryDefinition) GetDataSourceOk() (*FormulaAndFunctionEventsDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *FormulaAndFunctionEventQueryDefinition) SetDataSource(v FormulaAndFunctionEventsDataSource) {
	o.DataSource = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *FormulaAndFunctionEventQueryDefinition) GetGroupBy() []FormulaAndFunctionEventQueryGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []FormulaAndFunctionEventQueryGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionEventQueryDefinition) GetGroupByOk() (*[]FormulaAndFunctionEventQueryGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *FormulaAndFunctionEventQueryDefinition) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []FormulaAndFunctionEventQueryGroupBy and assigns it to the GroupBy field.
func (o *FormulaAndFunctionEventQueryDefinition) SetGroupBy(v []FormulaAndFunctionEventQueryGroupBy) {
	o.GroupBy = v
}

// GetIndexes returns the Indexes field value if set, zero value otherwise.
func (o *FormulaAndFunctionEventQueryDefinition) GetIndexes() []string {
	if o == nil || o.Indexes == nil {
		var ret []string
		return ret
	}
	return o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionEventQueryDefinition) GetIndexesOk() (*[]string, bool) {
	if o == nil || o.Indexes == nil {
		return nil, false
	}
	return &o.Indexes, true
}

// HasIndexes returns a boolean if a field has been set.
func (o *FormulaAndFunctionEventQueryDefinition) HasIndexes() bool {
	return o != nil && o.Indexes != nil
}

// SetIndexes gets a reference to the given []string and assigns it to the Indexes field.
func (o *FormulaAndFunctionEventQueryDefinition) SetIndexes(v []string) {
	o.Indexes = v
}

// GetName returns the Name field value.
func (o *FormulaAndFunctionEventQueryDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionEventQueryDefinition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *FormulaAndFunctionEventQueryDefinition) SetName(v string) {
	o.Name = v
}

// GetSearch returns the Search field value if set, zero value otherwise.
func (o *FormulaAndFunctionEventQueryDefinition) GetSearch() FormulaAndFunctionEventQueryDefinitionSearch {
	if o == nil || o.Search == nil {
		var ret FormulaAndFunctionEventQueryDefinitionSearch
		return ret
	}
	return *o.Search
}

// GetSearchOk returns a tuple with the Search field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionEventQueryDefinition) GetSearchOk() (*FormulaAndFunctionEventQueryDefinitionSearch, bool) {
	if o == nil || o.Search == nil {
		return nil, false
	}
	return o.Search, true
}

// HasSearch returns a boolean if a field has been set.
func (o *FormulaAndFunctionEventQueryDefinition) HasSearch() bool {
	return o != nil && o.Search != nil
}

// SetSearch gets a reference to the given FormulaAndFunctionEventQueryDefinitionSearch and assigns it to the Search field.
func (o *FormulaAndFunctionEventQueryDefinition) SetSearch(v FormulaAndFunctionEventQueryDefinitionSearch) {
	o.Search = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *FormulaAndFunctionEventQueryDefinition) GetStorage() string {
	if o == nil || o.Storage == nil {
		var ret string
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionEventQueryDefinition) GetStorageOk() (*string, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *FormulaAndFunctionEventQueryDefinition) HasStorage() bool {
	return o != nil && o.Storage != nil
}

// SetStorage gets a reference to the given string and assigns it to the Storage field.
func (o *FormulaAndFunctionEventQueryDefinition) SetStorage(v string) {
	o.Storage = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormulaAndFunctionEventQueryDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["compute"] = o.Compute
	if o.CrossOrgUuids != nil {
		toSerialize["cross_org_uuids"] = o.CrossOrgUuids
	}
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
	if o.Storage != nil {
		toSerialize["storage"] = o.Storage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormulaAndFunctionEventQueryDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Compute       *FormulaAndFunctionEventQueryDefinitionCompute `json:"compute"`
		CrossOrgUuids []string                                       `json:"cross_org_uuids,omitempty"`
		DataSource    *FormulaAndFunctionEventsDataSource            `json:"data_source"`
		GroupBy       []FormulaAndFunctionEventQueryGroupBy          `json:"group_by,omitempty"`
		Indexes       []string                                       `json:"indexes,omitempty"`
		Name          *string                                        `json:"name"`
		Search        *FormulaAndFunctionEventQueryDefinitionSearch  `json:"search,omitempty"`
		Storage       *string                                        `json:"storage,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"compute", "cross_org_uuids", "data_source", "group_by", "indexes", "name", "search", "storage"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Compute.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Compute = *all.Compute
	o.CrossOrgUuids = all.CrossOrgUuids
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
	o.Storage = all.Storage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
