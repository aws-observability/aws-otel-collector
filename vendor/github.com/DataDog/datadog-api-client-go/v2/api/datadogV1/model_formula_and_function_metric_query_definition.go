// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaAndFunctionMetricQueryDefinition A formula and functions metrics query.
type FormulaAndFunctionMetricQueryDefinition struct {
	// The aggregation methods available for metrics queries.
	Aggregator *FormulaAndFunctionMetricAggregation `json:"aggregator,omitempty"`
	// The source organization UUID for cross organization queries. Feature in Private Beta.
	CrossOrgUuids []string `json:"cross_org_uuids,omitempty"`
	// Data source for metrics queries.
	DataSource FormulaAndFunctionMetricDataSource `json:"data_source"`
	// Name of the query for use in formulas.
	Name string `json:"name"`
	// Metrics query definition.
	Query string `json:"query"`
	// Semantic mode for metrics queries. This determines how metrics from different sources are combined or displayed.
	SemanticMode *FormulaAndFunctionMetricSemanticMode `json:"semantic_mode,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFormulaAndFunctionMetricQueryDefinition instantiates a new FormulaAndFunctionMetricQueryDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormulaAndFunctionMetricQueryDefinition(dataSource FormulaAndFunctionMetricDataSource, name string, query string) *FormulaAndFunctionMetricQueryDefinition {
	this := FormulaAndFunctionMetricQueryDefinition{}
	this.DataSource = dataSource
	this.Name = name
	this.Query = query
	return &this
}

// NewFormulaAndFunctionMetricQueryDefinitionWithDefaults instantiates a new FormulaAndFunctionMetricQueryDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormulaAndFunctionMetricQueryDefinitionWithDefaults() *FormulaAndFunctionMetricQueryDefinition {
	this := FormulaAndFunctionMetricQueryDefinition{}
	return &this
}

// GetAggregator returns the Aggregator field value if set, zero value otherwise.
func (o *FormulaAndFunctionMetricQueryDefinition) GetAggregator() FormulaAndFunctionMetricAggregation {
	if o == nil || o.Aggregator == nil {
		var ret FormulaAndFunctionMetricAggregation
		return ret
	}
	return *o.Aggregator
}

// GetAggregatorOk returns a tuple with the Aggregator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionMetricQueryDefinition) GetAggregatorOk() (*FormulaAndFunctionMetricAggregation, bool) {
	if o == nil || o.Aggregator == nil {
		return nil, false
	}
	return o.Aggregator, true
}

// HasAggregator returns a boolean if a field has been set.
func (o *FormulaAndFunctionMetricQueryDefinition) HasAggregator() bool {
	return o != nil && o.Aggregator != nil
}

// SetAggregator gets a reference to the given FormulaAndFunctionMetricAggregation and assigns it to the Aggregator field.
func (o *FormulaAndFunctionMetricQueryDefinition) SetAggregator(v FormulaAndFunctionMetricAggregation) {
	o.Aggregator = &v
}

// GetCrossOrgUuids returns the CrossOrgUuids field value if set, zero value otherwise.
func (o *FormulaAndFunctionMetricQueryDefinition) GetCrossOrgUuids() []string {
	if o == nil || o.CrossOrgUuids == nil {
		var ret []string
		return ret
	}
	return o.CrossOrgUuids
}

// GetCrossOrgUuidsOk returns a tuple with the CrossOrgUuids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionMetricQueryDefinition) GetCrossOrgUuidsOk() (*[]string, bool) {
	if o == nil || o.CrossOrgUuids == nil {
		return nil, false
	}
	return &o.CrossOrgUuids, true
}

// HasCrossOrgUuids returns a boolean if a field has been set.
func (o *FormulaAndFunctionMetricQueryDefinition) HasCrossOrgUuids() bool {
	return o != nil && o.CrossOrgUuids != nil
}

// SetCrossOrgUuids gets a reference to the given []string and assigns it to the CrossOrgUuids field.
func (o *FormulaAndFunctionMetricQueryDefinition) SetCrossOrgUuids(v []string) {
	o.CrossOrgUuids = v
}

// GetDataSource returns the DataSource field value.
func (o *FormulaAndFunctionMetricQueryDefinition) GetDataSource() FormulaAndFunctionMetricDataSource {
	if o == nil {
		var ret FormulaAndFunctionMetricDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionMetricQueryDefinition) GetDataSourceOk() (*FormulaAndFunctionMetricDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *FormulaAndFunctionMetricQueryDefinition) SetDataSource(v FormulaAndFunctionMetricDataSource) {
	o.DataSource = v
}

// GetName returns the Name field value.
func (o *FormulaAndFunctionMetricQueryDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionMetricQueryDefinition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *FormulaAndFunctionMetricQueryDefinition) SetName(v string) {
	o.Name = v
}

// GetQuery returns the Query field value.
func (o *FormulaAndFunctionMetricQueryDefinition) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionMetricQueryDefinition) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *FormulaAndFunctionMetricQueryDefinition) SetQuery(v string) {
	o.Query = v
}

// GetSemanticMode returns the SemanticMode field value if set, zero value otherwise.
func (o *FormulaAndFunctionMetricQueryDefinition) GetSemanticMode() FormulaAndFunctionMetricSemanticMode {
	if o == nil || o.SemanticMode == nil {
		var ret FormulaAndFunctionMetricSemanticMode
		return ret
	}
	return *o.SemanticMode
}

// GetSemanticModeOk returns a tuple with the SemanticMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionMetricQueryDefinition) GetSemanticModeOk() (*FormulaAndFunctionMetricSemanticMode, bool) {
	if o == nil || o.SemanticMode == nil {
		return nil, false
	}
	return o.SemanticMode, true
}

// HasSemanticMode returns a boolean if a field has been set.
func (o *FormulaAndFunctionMetricQueryDefinition) HasSemanticMode() bool {
	return o != nil && o.SemanticMode != nil
}

// SetSemanticMode gets a reference to the given FormulaAndFunctionMetricSemanticMode and assigns it to the SemanticMode field.
func (o *FormulaAndFunctionMetricQueryDefinition) SetSemanticMode(v FormulaAndFunctionMetricSemanticMode) {
	o.SemanticMode = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormulaAndFunctionMetricQueryDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Aggregator != nil {
		toSerialize["aggregator"] = o.Aggregator
	}
	if o.CrossOrgUuids != nil {
		toSerialize["cross_org_uuids"] = o.CrossOrgUuids
	}
	toSerialize["data_source"] = o.DataSource
	toSerialize["name"] = o.Name
	toSerialize["query"] = o.Query
	if o.SemanticMode != nil {
		toSerialize["semantic_mode"] = o.SemanticMode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormulaAndFunctionMetricQueryDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregator    *FormulaAndFunctionMetricAggregation  `json:"aggregator,omitempty"`
		CrossOrgUuids []string                              `json:"cross_org_uuids,omitempty"`
		DataSource    *FormulaAndFunctionMetricDataSource   `json:"data_source"`
		Name          *string                               `json:"name"`
		Query         *string                               `json:"query"`
		SemanticMode  *FormulaAndFunctionMetricSemanticMode `json:"semantic_mode,omitempty"`
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
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aggregator", "cross_org_uuids", "data_source", "name", "query", "semantic_mode"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Aggregator != nil && !all.Aggregator.IsValid() {
		hasInvalidField = true
	} else {
		o.Aggregator = all.Aggregator
	}
	o.CrossOrgUuids = all.CrossOrgUuids
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.Name = *all.Name
	o.Query = *all.Query
	if all.SemanticMode != nil && !all.SemanticMode.IsValid() {
		hasInvalidField = true
	} else {
		o.SemanticMode = all.SemanticMode
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
