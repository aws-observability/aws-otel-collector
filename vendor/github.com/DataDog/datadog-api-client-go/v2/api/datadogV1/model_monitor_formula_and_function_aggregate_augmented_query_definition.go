// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition A formula and functions aggregate augmented query. Used to enrich base query results with data from a reference table.
type MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition struct {
	// Augment query for aggregate augmented queries. Can be an events query or a reference table query.
	AugmentQuery MonitorFormulaAndFunctionAggregateAugmentQuery `json:"augment_query"`
	// Base query for aggregate queries. Can be an events query or a metrics query.
	BaseQuery MonitorFormulaAndFunctionAggregateBaseQuery `json:"base_query"`
	// Compute options for the query.
	Compute []MonitorFormulaAndFunctionEventQueryDefinitionCompute `json:"compute"`
	// Data source for aggregate augmented queries.
	DataSource MonitorFormulaAndFunctionAggregateAugmentedDataSource `json:"data_source"`
	// Group by options for the query.
	GroupBy []MonitorFormulaAndFunctionEventQueryGroupBy `json:"group_by"`
	// Join condition for aggregate augmented queries.
	JoinCondition MonitorFormulaAndFunctionAggregateQueryJoinCondition `json:"join_condition"`
	// Name of the query for use in formulas.
	Name *string `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewMonitorFormulaAndFunctionAggregateAugmentedQueryDefinition instantiates a new MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorFormulaAndFunctionAggregateAugmentedQueryDefinition(augmentQuery MonitorFormulaAndFunctionAggregateAugmentQuery, baseQuery MonitorFormulaAndFunctionAggregateBaseQuery, compute []MonitorFormulaAndFunctionEventQueryDefinitionCompute, dataSource MonitorFormulaAndFunctionAggregateAugmentedDataSource, groupBy []MonitorFormulaAndFunctionEventQueryGroupBy, joinCondition MonitorFormulaAndFunctionAggregateQueryJoinCondition) *MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition {
	this := MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition{}
	this.AugmentQuery = augmentQuery
	this.BaseQuery = baseQuery
	this.Compute = compute
	this.DataSource = dataSource
	this.GroupBy = groupBy
	this.JoinCondition = joinCondition
	return &this
}

// NewMonitorFormulaAndFunctionAggregateAugmentedQueryDefinitionWithDefaults instantiates a new MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorFormulaAndFunctionAggregateAugmentedQueryDefinitionWithDefaults() *MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition {
	this := MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition{}
	return &this
}

// GetAugmentQuery returns the AugmentQuery field value.
func (o *MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition) GetAugmentQuery() MonitorFormulaAndFunctionAggregateAugmentQuery {
	if o == nil {
		var ret MonitorFormulaAndFunctionAggregateAugmentQuery
		return ret
	}
	return o.AugmentQuery
}

// GetAugmentQueryOk returns a tuple with the AugmentQuery field value
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition) GetAugmentQueryOk() (*MonitorFormulaAndFunctionAggregateAugmentQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AugmentQuery, true
}

// SetAugmentQuery sets field value.
func (o *MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition) SetAugmentQuery(v MonitorFormulaAndFunctionAggregateAugmentQuery) {
	o.AugmentQuery = v
}

// GetBaseQuery returns the BaseQuery field value.
func (o *MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition) GetBaseQuery() MonitorFormulaAndFunctionAggregateBaseQuery {
	if o == nil {
		var ret MonitorFormulaAndFunctionAggregateBaseQuery
		return ret
	}
	return o.BaseQuery
}

// GetBaseQueryOk returns a tuple with the BaseQuery field value
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition) GetBaseQueryOk() (*MonitorFormulaAndFunctionAggregateBaseQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseQuery, true
}

// SetBaseQuery sets field value.
func (o *MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition) SetBaseQuery(v MonitorFormulaAndFunctionAggregateBaseQuery) {
	o.BaseQuery = v
}

// GetCompute returns the Compute field value.
func (o *MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition) GetCompute() []MonitorFormulaAndFunctionEventQueryDefinitionCompute {
	if o == nil {
		var ret []MonitorFormulaAndFunctionEventQueryDefinitionCompute
		return ret
	}
	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition) GetComputeOk() (*[]MonitorFormulaAndFunctionEventQueryDefinitionCompute, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Compute, true
}

// SetCompute sets field value.
func (o *MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition) SetCompute(v []MonitorFormulaAndFunctionEventQueryDefinitionCompute) {
	o.Compute = v
}

// GetDataSource returns the DataSource field value.
func (o *MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition) GetDataSource() MonitorFormulaAndFunctionAggregateAugmentedDataSource {
	if o == nil {
		var ret MonitorFormulaAndFunctionAggregateAugmentedDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition) GetDataSourceOk() (*MonitorFormulaAndFunctionAggregateAugmentedDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition) SetDataSource(v MonitorFormulaAndFunctionAggregateAugmentedDataSource) {
	o.DataSource = v
}

// GetGroupBy returns the GroupBy field value.
func (o *MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition) GetGroupBy() []MonitorFormulaAndFunctionEventQueryGroupBy {
	if o == nil {
		var ret []MonitorFormulaAndFunctionEventQueryGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition) GetGroupByOk() (*[]MonitorFormulaAndFunctionEventQueryGroupBy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// SetGroupBy sets field value.
func (o *MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition) SetGroupBy(v []MonitorFormulaAndFunctionEventQueryGroupBy) {
	o.GroupBy = v
}

// GetJoinCondition returns the JoinCondition field value.
func (o *MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition) GetJoinCondition() MonitorFormulaAndFunctionAggregateQueryJoinCondition {
	if o == nil {
		var ret MonitorFormulaAndFunctionAggregateQueryJoinCondition
		return ret
	}
	return o.JoinCondition
}

// GetJoinConditionOk returns a tuple with the JoinCondition field value
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition) GetJoinConditionOk() (*MonitorFormulaAndFunctionAggregateQueryJoinCondition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JoinCondition, true
}

// SetJoinCondition sets field value.
func (o *MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition) SetJoinCondition(v MonitorFormulaAndFunctionAggregateQueryJoinCondition) {
	o.JoinCondition = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition) SetName(v string) {
	o.Name = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["augment_query"] = o.AugmentQuery
	toSerialize["base_query"] = o.BaseQuery
	toSerialize["compute"] = o.Compute
	toSerialize["data_source"] = o.DataSource
	toSerialize["group_by"] = o.GroupBy
	toSerialize["join_condition"] = o.JoinCondition
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorFormulaAndFunctionAggregateAugmentedQueryDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AugmentQuery  *MonitorFormulaAndFunctionAggregateAugmentQuery         `json:"augment_query"`
		BaseQuery     *MonitorFormulaAndFunctionAggregateBaseQuery            `json:"base_query"`
		Compute       *[]MonitorFormulaAndFunctionEventQueryDefinitionCompute `json:"compute"`
		DataSource    *MonitorFormulaAndFunctionAggregateAugmentedDataSource  `json:"data_source"`
		GroupBy       *[]MonitorFormulaAndFunctionEventQueryGroupBy           `json:"group_by"`
		JoinCondition *MonitorFormulaAndFunctionAggregateQueryJoinCondition   `json:"join_condition"`
		Name          *string                                                 `json:"name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AugmentQuery == nil {
		return fmt.Errorf("required field augment_query missing")
	}
	if all.BaseQuery == nil {
		return fmt.Errorf("required field base_query missing")
	}
	if all.Compute == nil {
		return fmt.Errorf("required field compute missing")
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.GroupBy == nil {
		return fmt.Errorf("required field group_by missing")
	}
	if all.JoinCondition == nil {
		return fmt.Errorf("required field join_condition missing")
	}

	hasInvalidField := false
	o.AugmentQuery = *all.AugmentQuery
	o.BaseQuery = *all.BaseQuery
	o.Compute = *all.Compute
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.GroupBy = *all.GroupBy
	if all.JoinCondition.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.JoinCondition = *all.JoinCondition
	o.Name = all.Name

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
