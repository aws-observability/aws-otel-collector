// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorFormulaAndFunctionDataQualityQueryDefinition A formula and functions data quality query.
type MonitorFormulaAndFunctionDataQualityQueryDefinition struct {
	// Data source for data quality queries.
	DataSource MonitorFormulaAndFunctionDataQualityDataSource `json:"data_source"`
	// Filter expression used to match on data entities. Uses Aastra query syntax.
	Filter string `json:"filter"`
	// Optional grouping fields for aggregation.
	GroupBy []string `json:"group_by,omitempty"`
	// The data quality measure to query. Common values include:
	// `bytes`, `cardinality`, `custom`, `freshness`, `max`, `mean`, `min`,
	// `nullness`, `percent_negative`, `percent_zero`, `row_count`, `stddev`,
	// `sum`, `uniqueness`. Additional values may be supported.
	Measure string `json:"measure"`
	// Monitor configuration options for data quality queries.
	MonitorOptions *MonitorFormulaAndFunctionDataQualityMonitorOptions `json:"monitor_options,omitempty"`
	// Name of the query for use in formulas.
	Name string `json:"name"`
	// Schema version for the data quality query.
	SchemaVersion *string `json:"schema_version,omitempty"`
	// Optional scoping expression to further filter metrics. Uses metrics filter syntax.
	// This is useful when an entity has been configured to emit metrics with additional tags.
	Scope *string `json:"scope,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMonitorFormulaAndFunctionDataQualityQueryDefinition instantiates a new MonitorFormulaAndFunctionDataQualityQueryDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorFormulaAndFunctionDataQualityQueryDefinition(dataSource MonitorFormulaAndFunctionDataQualityDataSource, filter string, measure string, name string) *MonitorFormulaAndFunctionDataQualityQueryDefinition {
	this := MonitorFormulaAndFunctionDataQualityQueryDefinition{}
	this.DataSource = dataSource
	this.Filter = filter
	this.Measure = measure
	this.Name = name
	return &this
}

// NewMonitorFormulaAndFunctionDataQualityQueryDefinitionWithDefaults instantiates a new MonitorFormulaAndFunctionDataQualityQueryDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorFormulaAndFunctionDataQualityQueryDefinitionWithDefaults() *MonitorFormulaAndFunctionDataQualityQueryDefinition {
	this := MonitorFormulaAndFunctionDataQualityQueryDefinition{}
	return &this
}

// GetDataSource returns the DataSource field value.
func (o *MonitorFormulaAndFunctionDataQualityQueryDefinition) GetDataSource() MonitorFormulaAndFunctionDataQualityDataSource {
	if o == nil {
		var ret MonitorFormulaAndFunctionDataQualityDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionDataQualityQueryDefinition) GetDataSourceOk() (*MonitorFormulaAndFunctionDataQualityDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *MonitorFormulaAndFunctionDataQualityQueryDefinition) SetDataSource(v MonitorFormulaAndFunctionDataQualityDataSource) {
	o.DataSource = v
}

// GetFilter returns the Filter field value.
func (o *MonitorFormulaAndFunctionDataQualityQueryDefinition) GetFilter() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionDataQualityQueryDefinition) GetFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value.
func (o *MonitorFormulaAndFunctionDataQualityQueryDefinition) SetFilter(v string) {
	o.Filter = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *MonitorFormulaAndFunctionDataQualityQueryDefinition) GetGroupBy() []string {
	if o == nil || o.GroupBy == nil {
		var ret []string
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionDataQualityQueryDefinition) GetGroupByOk() (*[]string, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *MonitorFormulaAndFunctionDataQualityQueryDefinition) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []string and assigns it to the GroupBy field.
func (o *MonitorFormulaAndFunctionDataQualityQueryDefinition) SetGroupBy(v []string) {
	o.GroupBy = v
}

// GetMeasure returns the Measure field value.
func (o *MonitorFormulaAndFunctionDataQualityQueryDefinition) GetMeasure() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Measure
}

// GetMeasureOk returns a tuple with the Measure field value
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionDataQualityQueryDefinition) GetMeasureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Measure, true
}

// SetMeasure sets field value.
func (o *MonitorFormulaAndFunctionDataQualityQueryDefinition) SetMeasure(v string) {
	o.Measure = v
}

// GetMonitorOptions returns the MonitorOptions field value if set, zero value otherwise.
func (o *MonitorFormulaAndFunctionDataQualityQueryDefinition) GetMonitorOptions() MonitorFormulaAndFunctionDataQualityMonitorOptions {
	if o == nil || o.MonitorOptions == nil {
		var ret MonitorFormulaAndFunctionDataQualityMonitorOptions
		return ret
	}
	return *o.MonitorOptions
}

// GetMonitorOptionsOk returns a tuple with the MonitorOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionDataQualityQueryDefinition) GetMonitorOptionsOk() (*MonitorFormulaAndFunctionDataQualityMonitorOptions, bool) {
	if o == nil || o.MonitorOptions == nil {
		return nil, false
	}
	return o.MonitorOptions, true
}

// HasMonitorOptions returns a boolean if a field has been set.
func (o *MonitorFormulaAndFunctionDataQualityQueryDefinition) HasMonitorOptions() bool {
	return o != nil && o.MonitorOptions != nil
}

// SetMonitorOptions gets a reference to the given MonitorFormulaAndFunctionDataQualityMonitorOptions and assigns it to the MonitorOptions field.
func (o *MonitorFormulaAndFunctionDataQualityQueryDefinition) SetMonitorOptions(v MonitorFormulaAndFunctionDataQualityMonitorOptions) {
	o.MonitorOptions = &v
}

// GetName returns the Name field value.
func (o *MonitorFormulaAndFunctionDataQualityQueryDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionDataQualityQueryDefinition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *MonitorFormulaAndFunctionDataQualityQueryDefinition) SetName(v string) {
	o.Name = v
}

// GetSchemaVersion returns the SchemaVersion field value if set, zero value otherwise.
func (o *MonitorFormulaAndFunctionDataQualityQueryDefinition) GetSchemaVersion() string {
	if o == nil || o.SchemaVersion == nil {
		var ret string
		return ret
	}
	return *o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionDataQualityQueryDefinition) GetSchemaVersionOk() (*string, bool) {
	if o == nil || o.SchemaVersion == nil {
		return nil, false
	}
	return o.SchemaVersion, true
}

// HasSchemaVersion returns a boolean if a field has been set.
func (o *MonitorFormulaAndFunctionDataQualityQueryDefinition) HasSchemaVersion() bool {
	return o != nil && o.SchemaVersion != nil
}

// SetSchemaVersion gets a reference to the given string and assigns it to the SchemaVersion field.
func (o *MonitorFormulaAndFunctionDataQualityQueryDefinition) SetSchemaVersion(v string) {
	o.SchemaVersion = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *MonitorFormulaAndFunctionDataQualityQueryDefinition) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionDataQualityQueryDefinition) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *MonitorFormulaAndFunctionDataQualityQueryDefinition) HasScope() bool {
	return o != nil && o.Scope != nil
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *MonitorFormulaAndFunctionDataQualityQueryDefinition) SetScope(v string) {
	o.Scope = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorFormulaAndFunctionDataQualityQueryDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data_source"] = o.DataSource
	toSerialize["filter"] = o.Filter
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	toSerialize["measure"] = o.Measure
	if o.MonitorOptions != nil {
		toSerialize["monitor_options"] = o.MonitorOptions
	}
	toSerialize["name"] = o.Name
	if o.SchemaVersion != nil {
		toSerialize["schema_version"] = o.SchemaVersion
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorFormulaAndFunctionDataQualityQueryDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataSource     *MonitorFormulaAndFunctionDataQualityDataSource     `json:"data_source"`
		Filter         *string                                             `json:"filter"`
		GroupBy        []string                                            `json:"group_by,omitempty"`
		Measure        *string                                             `json:"measure"`
		MonitorOptions *MonitorFormulaAndFunctionDataQualityMonitorOptions `json:"monitor_options,omitempty"`
		Name           *string                                             `json:"name"`
		SchemaVersion  *string                                             `json:"schema_version,omitempty"`
		Scope          *string                                             `json:"scope,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.Filter == nil {
		return fmt.Errorf("required field filter missing")
	}
	if all.Measure == nil {
		return fmt.Errorf("required field measure missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data_source", "filter", "group_by", "measure", "monitor_options", "name", "schema_version", "scope"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.Filter = *all.Filter
	o.GroupBy = all.GroupBy
	o.Measure = *all.Measure
	if all.MonitorOptions != nil && all.MonitorOptions.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.MonitorOptions = all.MonitorOptions
	o.Name = *all.Name
	o.SchemaVersion = all.SchemaVersion
	o.Scope = all.Scope

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
