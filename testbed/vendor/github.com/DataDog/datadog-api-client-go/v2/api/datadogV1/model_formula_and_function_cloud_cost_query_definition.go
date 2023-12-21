// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaAndFunctionCloudCostQueryDefinition A formula and functions Cloud Cost query.
type FormulaAndFunctionCloudCostQueryDefinition struct {
	// Aggregator used for the request.
	Aggregator *WidgetAggregator `json:"aggregator,omitempty"`
	// Data source for Cloud Cost queries.
	DataSource FormulaAndFunctionCloudCostDataSource `json:"data_source"`
	// Name of the query for use in formulas.
	Name string `json:"name"`
	// Query for Cloud Cost data.
	Query string `json:"query"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewFormulaAndFunctionCloudCostQueryDefinition instantiates a new FormulaAndFunctionCloudCostQueryDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormulaAndFunctionCloudCostQueryDefinition(dataSource FormulaAndFunctionCloudCostDataSource, name string, query string) *FormulaAndFunctionCloudCostQueryDefinition {
	this := FormulaAndFunctionCloudCostQueryDefinition{}
	this.DataSource = dataSource
	this.Name = name
	this.Query = query
	return &this
}

// NewFormulaAndFunctionCloudCostQueryDefinitionWithDefaults instantiates a new FormulaAndFunctionCloudCostQueryDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormulaAndFunctionCloudCostQueryDefinitionWithDefaults() *FormulaAndFunctionCloudCostQueryDefinition {
	this := FormulaAndFunctionCloudCostQueryDefinition{}
	return &this
}

// GetAggregator returns the Aggregator field value if set, zero value otherwise.
func (o *FormulaAndFunctionCloudCostQueryDefinition) GetAggregator() WidgetAggregator {
	if o == nil || o.Aggregator == nil {
		var ret WidgetAggregator
		return ret
	}
	return *o.Aggregator
}

// GetAggregatorOk returns a tuple with the Aggregator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionCloudCostQueryDefinition) GetAggregatorOk() (*WidgetAggregator, bool) {
	if o == nil || o.Aggregator == nil {
		return nil, false
	}
	return o.Aggregator, true
}

// HasAggregator returns a boolean if a field has been set.
func (o *FormulaAndFunctionCloudCostQueryDefinition) HasAggregator() bool {
	return o != nil && o.Aggregator != nil
}

// SetAggregator gets a reference to the given WidgetAggregator and assigns it to the Aggregator field.
func (o *FormulaAndFunctionCloudCostQueryDefinition) SetAggregator(v WidgetAggregator) {
	o.Aggregator = &v
}

// GetDataSource returns the DataSource field value.
func (o *FormulaAndFunctionCloudCostQueryDefinition) GetDataSource() FormulaAndFunctionCloudCostDataSource {
	if o == nil {
		var ret FormulaAndFunctionCloudCostDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionCloudCostQueryDefinition) GetDataSourceOk() (*FormulaAndFunctionCloudCostDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *FormulaAndFunctionCloudCostQueryDefinition) SetDataSource(v FormulaAndFunctionCloudCostDataSource) {
	o.DataSource = v
}

// GetName returns the Name field value.
func (o *FormulaAndFunctionCloudCostQueryDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionCloudCostQueryDefinition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *FormulaAndFunctionCloudCostQueryDefinition) SetName(v string) {
	o.Name = v
}

// GetQuery returns the Query field value.
func (o *FormulaAndFunctionCloudCostQueryDefinition) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionCloudCostQueryDefinition) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *FormulaAndFunctionCloudCostQueryDefinition) SetQuery(v string) {
	o.Query = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormulaAndFunctionCloudCostQueryDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Aggregator != nil {
		toSerialize["aggregator"] = o.Aggregator
	}
	toSerialize["data_source"] = o.DataSource
	toSerialize["name"] = o.Name
	toSerialize["query"] = o.Query

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormulaAndFunctionCloudCostQueryDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregator *WidgetAggregator                      `json:"aggregator,omitempty"`
		DataSource *FormulaAndFunctionCloudCostDataSource `json:"data_source"`
		Name       *string                                `json:"name"`
		Query      *string                                `json:"query"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
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
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aggregator", "data_source", "name", "query"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Aggregator != nil && !all.Aggregator.IsValid() {
		hasInvalidField = true
	} else {
		o.Aggregator = all.Aggregator
	}
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.Name = *all.Name
	o.Query = *all.Query

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
