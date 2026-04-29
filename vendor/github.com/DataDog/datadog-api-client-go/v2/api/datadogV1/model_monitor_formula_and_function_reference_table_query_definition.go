// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorFormulaAndFunctionReferenceTableQueryDefinition A reference table query for use in aggregate queries.
type MonitorFormulaAndFunctionReferenceTableQueryDefinition struct {
	// List of columns to retrieve from the reference table.
	Columns []MonitorFormulaAndFunctionReferenceTableColumn `json:"columns,omitempty"`
	// Data source for reference table queries.
	DataSource MonitorFormulaAndFunctionReferenceTableDataSource `json:"data_source"`
	// Name of the query.
	Name *string `json:"name,omitempty"`
	// Optional filter expression for the reference table query.
	QueryFilter *string `json:"query_filter,omitempty"`
	// Name of the reference table.
	TableName string `json:"table_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewMonitorFormulaAndFunctionReferenceTableQueryDefinition instantiates a new MonitorFormulaAndFunctionReferenceTableQueryDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorFormulaAndFunctionReferenceTableQueryDefinition(dataSource MonitorFormulaAndFunctionReferenceTableDataSource, tableName string) *MonitorFormulaAndFunctionReferenceTableQueryDefinition {
	this := MonitorFormulaAndFunctionReferenceTableQueryDefinition{}
	this.DataSource = dataSource
	this.TableName = tableName
	return &this
}

// NewMonitorFormulaAndFunctionReferenceTableQueryDefinitionWithDefaults instantiates a new MonitorFormulaAndFunctionReferenceTableQueryDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorFormulaAndFunctionReferenceTableQueryDefinitionWithDefaults() *MonitorFormulaAndFunctionReferenceTableQueryDefinition {
	this := MonitorFormulaAndFunctionReferenceTableQueryDefinition{}
	return &this
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *MonitorFormulaAndFunctionReferenceTableQueryDefinition) GetColumns() []MonitorFormulaAndFunctionReferenceTableColumn {
	if o == nil || o.Columns == nil {
		var ret []MonitorFormulaAndFunctionReferenceTableColumn
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionReferenceTableQueryDefinition) GetColumnsOk() (*[]MonitorFormulaAndFunctionReferenceTableColumn, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return &o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *MonitorFormulaAndFunctionReferenceTableQueryDefinition) HasColumns() bool {
	return o != nil && o.Columns != nil
}

// SetColumns gets a reference to the given []MonitorFormulaAndFunctionReferenceTableColumn and assigns it to the Columns field.
func (o *MonitorFormulaAndFunctionReferenceTableQueryDefinition) SetColumns(v []MonitorFormulaAndFunctionReferenceTableColumn) {
	o.Columns = v
}

// GetDataSource returns the DataSource field value.
func (o *MonitorFormulaAndFunctionReferenceTableQueryDefinition) GetDataSource() MonitorFormulaAndFunctionReferenceTableDataSource {
	if o == nil {
		var ret MonitorFormulaAndFunctionReferenceTableDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionReferenceTableQueryDefinition) GetDataSourceOk() (*MonitorFormulaAndFunctionReferenceTableDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *MonitorFormulaAndFunctionReferenceTableQueryDefinition) SetDataSource(v MonitorFormulaAndFunctionReferenceTableDataSource) {
	o.DataSource = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MonitorFormulaAndFunctionReferenceTableQueryDefinition) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionReferenceTableQueryDefinition) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MonitorFormulaAndFunctionReferenceTableQueryDefinition) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MonitorFormulaAndFunctionReferenceTableQueryDefinition) SetName(v string) {
	o.Name = &v
}

// GetQueryFilter returns the QueryFilter field value if set, zero value otherwise.
func (o *MonitorFormulaAndFunctionReferenceTableQueryDefinition) GetQueryFilter() string {
	if o == nil || o.QueryFilter == nil {
		var ret string
		return ret
	}
	return *o.QueryFilter
}

// GetQueryFilterOk returns a tuple with the QueryFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionReferenceTableQueryDefinition) GetQueryFilterOk() (*string, bool) {
	if o == nil || o.QueryFilter == nil {
		return nil, false
	}
	return o.QueryFilter, true
}

// HasQueryFilter returns a boolean if a field has been set.
func (o *MonitorFormulaAndFunctionReferenceTableQueryDefinition) HasQueryFilter() bool {
	return o != nil && o.QueryFilter != nil
}

// SetQueryFilter gets a reference to the given string and assigns it to the QueryFilter field.
func (o *MonitorFormulaAndFunctionReferenceTableQueryDefinition) SetQueryFilter(v string) {
	o.QueryFilter = &v
}

// GetTableName returns the TableName field value.
func (o *MonitorFormulaAndFunctionReferenceTableQueryDefinition) GetTableName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TableName
}

// GetTableNameOk returns a tuple with the TableName field value
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionReferenceTableQueryDefinition) GetTableNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TableName, true
}

// SetTableName sets field value.
func (o *MonitorFormulaAndFunctionReferenceTableQueryDefinition) SetTableName(v string) {
	o.TableName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorFormulaAndFunctionReferenceTableQueryDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	toSerialize["data_source"] = o.DataSource
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.QueryFilter != nil {
		toSerialize["query_filter"] = o.QueryFilter
	}
	toSerialize["table_name"] = o.TableName
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorFormulaAndFunctionReferenceTableQueryDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Columns     []MonitorFormulaAndFunctionReferenceTableColumn    `json:"columns,omitempty"`
		DataSource  *MonitorFormulaAndFunctionReferenceTableDataSource `json:"data_source"`
		Name        *string                                            `json:"name,omitempty"`
		QueryFilter *string                                            `json:"query_filter,omitempty"`
		TableName   *string                                            `json:"table_name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.TableName == nil {
		return fmt.Errorf("required field table_name missing")
	}

	hasInvalidField := false
	o.Columns = all.Columns
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.Name = all.Name
	o.QueryFilter = all.QueryFilter
	o.TableName = *all.TableName

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
