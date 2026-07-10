// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringDatasetDefinition The definition of the dataset. The shape depends on the value of `data_source`.
// Use `reference_table` or `managed_resource` for a referential dataset, or one of the
// event platform sources (for example `logs`, `audit`, `events`, `spans`, `rum`) for
// an event platform dataset.
type SecurityMonitoringDatasetDefinition struct {
	// For event platform datasets, the list of columns exposed by the dataset.
	Columns []SecurityMonitoringDatasetColumn `json:"columns,omitempty"`
	// The data source backing this dataset definition.
	DataSource string `json:"data_source"`
	// For event platform datasets, the list of indexes to query.
	Indexes []string `json:"indexes,omitempty"`
	// The unique name of the dataset. Must start with a lowercase letter and contain only lowercase letters, digits, and underscores (max 255 characters).
	Name string `json:"name"`
	// For referential datasets, an optional filter expression applied to the table.
	QueryFilter *string `json:"query_filter,omitempty"`
	// The search clause applied to an event platform dataset.
	Search *SecurityMonitoringDatasetSearch `json:"search,omitempty"`
	// Storage tier the dataset reads from. Applies to event platform datasets.
	Storage *string `json:"storage,omitempty"`
	// For referential datasets, the name of the underlying table.
	TableName *string `json:"table_name,omitempty"`
	// An optional time window that overrides the default query time range.
	TimeWindow *SecurityMonitoringDatasetTimeWindow `json:"time_window,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringDatasetDefinition instantiates a new SecurityMonitoringDatasetDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringDatasetDefinition(dataSource string, name string) *SecurityMonitoringDatasetDefinition {
	this := SecurityMonitoringDatasetDefinition{}
	this.DataSource = dataSource
	this.Name = name
	return &this
}

// NewSecurityMonitoringDatasetDefinitionWithDefaults instantiates a new SecurityMonitoringDatasetDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringDatasetDefinitionWithDefaults() *SecurityMonitoringDatasetDefinition {
	this := SecurityMonitoringDatasetDefinition{}
	return &this
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *SecurityMonitoringDatasetDefinition) GetColumns() []SecurityMonitoringDatasetColumn {
	if o == nil || o.Columns == nil {
		var ret []SecurityMonitoringDatasetColumn
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetDefinition) GetColumnsOk() (*[]SecurityMonitoringDatasetColumn, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return &o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *SecurityMonitoringDatasetDefinition) HasColumns() bool {
	return o != nil && o.Columns != nil
}

// SetColumns gets a reference to the given []SecurityMonitoringDatasetColumn and assigns it to the Columns field.
func (o *SecurityMonitoringDatasetDefinition) SetColumns(v []SecurityMonitoringDatasetColumn) {
	o.Columns = v
}

// GetDataSource returns the DataSource field value.
func (o *SecurityMonitoringDatasetDefinition) GetDataSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetDefinition) GetDataSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *SecurityMonitoringDatasetDefinition) SetDataSource(v string) {
	o.DataSource = v
}

// GetIndexes returns the Indexes field value if set, zero value otherwise.
func (o *SecurityMonitoringDatasetDefinition) GetIndexes() []string {
	if o == nil || o.Indexes == nil {
		var ret []string
		return ret
	}
	return o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetDefinition) GetIndexesOk() (*[]string, bool) {
	if o == nil || o.Indexes == nil {
		return nil, false
	}
	return &o.Indexes, true
}

// HasIndexes returns a boolean if a field has been set.
func (o *SecurityMonitoringDatasetDefinition) HasIndexes() bool {
	return o != nil && o.Indexes != nil
}

// SetIndexes gets a reference to the given []string and assigns it to the Indexes field.
func (o *SecurityMonitoringDatasetDefinition) SetIndexes(v []string) {
	o.Indexes = v
}

// GetName returns the Name field value.
func (o *SecurityMonitoringDatasetDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetDefinition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SecurityMonitoringDatasetDefinition) SetName(v string) {
	o.Name = v
}

// GetQueryFilter returns the QueryFilter field value if set, zero value otherwise.
func (o *SecurityMonitoringDatasetDefinition) GetQueryFilter() string {
	if o == nil || o.QueryFilter == nil {
		var ret string
		return ret
	}
	return *o.QueryFilter
}

// GetQueryFilterOk returns a tuple with the QueryFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetDefinition) GetQueryFilterOk() (*string, bool) {
	if o == nil || o.QueryFilter == nil {
		return nil, false
	}
	return o.QueryFilter, true
}

// HasQueryFilter returns a boolean if a field has been set.
func (o *SecurityMonitoringDatasetDefinition) HasQueryFilter() bool {
	return o != nil && o.QueryFilter != nil
}

// SetQueryFilter gets a reference to the given string and assigns it to the QueryFilter field.
func (o *SecurityMonitoringDatasetDefinition) SetQueryFilter(v string) {
	o.QueryFilter = &v
}

// GetSearch returns the Search field value if set, zero value otherwise.
func (o *SecurityMonitoringDatasetDefinition) GetSearch() SecurityMonitoringDatasetSearch {
	if o == nil || o.Search == nil {
		var ret SecurityMonitoringDatasetSearch
		return ret
	}
	return *o.Search
}

// GetSearchOk returns a tuple with the Search field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetDefinition) GetSearchOk() (*SecurityMonitoringDatasetSearch, bool) {
	if o == nil || o.Search == nil {
		return nil, false
	}
	return o.Search, true
}

// HasSearch returns a boolean if a field has been set.
func (o *SecurityMonitoringDatasetDefinition) HasSearch() bool {
	return o != nil && o.Search != nil
}

// SetSearch gets a reference to the given SecurityMonitoringDatasetSearch and assigns it to the Search field.
func (o *SecurityMonitoringDatasetDefinition) SetSearch(v SecurityMonitoringDatasetSearch) {
	o.Search = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *SecurityMonitoringDatasetDefinition) GetStorage() string {
	if o == nil || o.Storage == nil {
		var ret string
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetDefinition) GetStorageOk() (*string, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *SecurityMonitoringDatasetDefinition) HasStorage() bool {
	return o != nil && o.Storage != nil
}

// SetStorage gets a reference to the given string and assigns it to the Storage field.
func (o *SecurityMonitoringDatasetDefinition) SetStorage(v string) {
	o.Storage = &v
}

// GetTableName returns the TableName field value if set, zero value otherwise.
func (o *SecurityMonitoringDatasetDefinition) GetTableName() string {
	if o == nil || o.TableName == nil {
		var ret string
		return ret
	}
	return *o.TableName
}

// GetTableNameOk returns a tuple with the TableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetDefinition) GetTableNameOk() (*string, bool) {
	if o == nil || o.TableName == nil {
		return nil, false
	}
	return o.TableName, true
}

// HasTableName returns a boolean if a field has been set.
func (o *SecurityMonitoringDatasetDefinition) HasTableName() bool {
	return o != nil && o.TableName != nil
}

// SetTableName gets a reference to the given string and assigns it to the TableName field.
func (o *SecurityMonitoringDatasetDefinition) SetTableName(v string) {
	o.TableName = &v
}

// GetTimeWindow returns the TimeWindow field value if set, zero value otherwise.
func (o *SecurityMonitoringDatasetDefinition) GetTimeWindow() SecurityMonitoringDatasetTimeWindow {
	if o == nil || o.TimeWindow == nil {
		var ret SecurityMonitoringDatasetTimeWindow
		return ret
	}
	return *o.TimeWindow
}

// GetTimeWindowOk returns a tuple with the TimeWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetDefinition) GetTimeWindowOk() (*SecurityMonitoringDatasetTimeWindow, bool) {
	if o == nil || o.TimeWindow == nil {
		return nil, false
	}
	return o.TimeWindow, true
}

// HasTimeWindow returns a boolean if a field has been set.
func (o *SecurityMonitoringDatasetDefinition) HasTimeWindow() bool {
	return o != nil && o.TimeWindow != nil
}

// SetTimeWindow gets a reference to the given SecurityMonitoringDatasetTimeWindow and assigns it to the TimeWindow field.
func (o *SecurityMonitoringDatasetDefinition) SetTimeWindow(v SecurityMonitoringDatasetTimeWindow) {
	o.TimeWindow = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringDatasetDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	toSerialize["data_source"] = o.DataSource
	if o.Indexes != nil {
		toSerialize["indexes"] = o.Indexes
	}
	toSerialize["name"] = o.Name
	if o.QueryFilter != nil {
		toSerialize["query_filter"] = o.QueryFilter
	}
	if o.Search != nil {
		toSerialize["search"] = o.Search
	}
	if o.Storage != nil {
		toSerialize["storage"] = o.Storage
	}
	if o.TableName != nil {
		toSerialize["table_name"] = o.TableName
	}
	if o.TimeWindow != nil {
		toSerialize["time_window"] = o.TimeWindow
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringDatasetDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Columns     []SecurityMonitoringDatasetColumn    `json:"columns,omitempty"`
		DataSource  *string                              `json:"data_source"`
		Indexes     []string                             `json:"indexes,omitempty"`
		Name        *string                              `json:"name"`
		QueryFilter *string                              `json:"query_filter,omitempty"`
		Search      *SecurityMonitoringDatasetSearch     `json:"search,omitempty"`
		Storage     *string                              `json:"storage,omitempty"`
		TableName   *string                              `json:"table_name,omitempty"`
		TimeWindow  *SecurityMonitoringDatasetTimeWindow `json:"time_window,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"columns", "data_source", "indexes", "name", "query_filter", "search", "storage", "table_name", "time_window"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Columns = all.Columns
	o.DataSource = *all.DataSource
	o.Indexes = all.Indexes
	o.Name = *all.Name
	o.QueryFilter = all.QueryFilter
	if all.Search != nil && all.Search.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Search = all.Search
	o.Storage = all.Storage
	o.TableName = all.TableName
	if all.TimeWindow != nil && all.TimeWindow.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TimeWindow = all.TimeWindow

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
