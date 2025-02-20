// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaAndFunctionSLOQueryDefinition A formula and functions metrics query.
type FormulaAndFunctionSLOQueryDefinition struct {
	// Additional filters applied to the SLO query.
	AdditionalQueryFilters *string `json:"additional_query_filters,omitempty"`
	// The source organization UUID for cross organization queries. Feature in Private Beta.
	CrossOrgUuids []string `json:"cross_org_uuids,omitempty"`
	// Data source for SLO measures queries.
	DataSource FormulaAndFunctionSLODataSource `json:"data_source"`
	// Group mode to query measures.
	GroupMode *FormulaAndFunctionSLOGroupMode `json:"group_mode,omitempty"`
	// SLO measures queries.
	Measure FormulaAndFunctionSLOMeasure `json:"measure"`
	// Name of the query for use in formulas.
	Name *string `json:"name,omitempty"`
	// ID of an SLO to query measures.
	SloId string `json:"slo_id"`
	// Name of the query for use in formulas.
	SloQueryType *FormulaAndFunctionSLOQueryType `json:"slo_query_type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFormulaAndFunctionSLOQueryDefinition instantiates a new FormulaAndFunctionSLOQueryDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormulaAndFunctionSLOQueryDefinition(dataSource FormulaAndFunctionSLODataSource, measure FormulaAndFunctionSLOMeasure, sloId string) *FormulaAndFunctionSLOQueryDefinition {
	this := FormulaAndFunctionSLOQueryDefinition{}
	this.DataSource = dataSource
	this.Measure = measure
	this.SloId = sloId
	return &this
}

// NewFormulaAndFunctionSLOQueryDefinitionWithDefaults instantiates a new FormulaAndFunctionSLOQueryDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormulaAndFunctionSLOQueryDefinitionWithDefaults() *FormulaAndFunctionSLOQueryDefinition {
	this := FormulaAndFunctionSLOQueryDefinition{}
	return &this
}

// GetAdditionalQueryFilters returns the AdditionalQueryFilters field value if set, zero value otherwise.
func (o *FormulaAndFunctionSLOQueryDefinition) GetAdditionalQueryFilters() string {
	if o == nil || o.AdditionalQueryFilters == nil {
		var ret string
		return ret
	}
	return *o.AdditionalQueryFilters
}

// GetAdditionalQueryFiltersOk returns a tuple with the AdditionalQueryFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionSLOQueryDefinition) GetAdditionalQueryFiltersOk() (*string, bool) {
	if o == nil || o.AdditionalQueryFilters == nil {
		return nil, false
	}
	return o.AdditionalQueryFilters, true
}

// HasAdditionalQueryFilters returns a boolean if a field has been set.
func (o *FormulaAndFunctionSLOQueryDefinition) HasAdditionalQueryFilters() bool {
	return o != nil && o.AdditionalQueryFilters != nil
}

// SetAdditionalQueryFilters gets a reference to the given string and assigns it to the AdditionalQueryFilters field.
func (o *FormulaAndFunctionSLOQueryDefinition) SetAdditionalQueryFilters(v string) {
	o.AdditionalQueryFilters = &v
}

// GetCrossOrgUuids returns the CrossOrgUuids field value if set, zero value otherwise.
func (o *FormulaAndFunctionSLOQueryDefinition) GetCrossOrgUuids() []string {
	if o == nil || o.CrossOrgUuids == nil {
		var ret []string
		return ret
	}
	return o.CrossOrgUuids
}

// GetCrossOrgUuidsOk returns a tuple with the CrossOrgUuids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionSLOQueryDefinition) GetCrossOrgUuidsOk() (*[]string, bool) {
	if o == nil || o.CrossOrgUuids == nil {
		return nil, false
	}
	return &o.CrossOrgUuids, true
}

// HasCrossOrgUuids returns a boolean if a field has been set.
func (o *FormulaAndFunctionSLOQueryDefinition) HasCrossOrgUuids() bool {
	return o != nil && o.CrossOrgUuids != nil
}

// SetCrossOrgUuids gets a reference to the given []string and assigns it to the CrossOrgUuids field.
func (o *FormulaAndFunctionSLOQueryDefinition) SetCrossOrgUuids(v []string) {
	o.CrossOrgUuids = v
}

// GetDataSource returns the DataSource field value.
func (o *FormulaAndFunctionSLOQueryDefinition) GetDataSource() FormulaAndFunctionSLODataSource {
	if o == nil {
		var ret FormulaAndFunctionSLODataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionSLOQueryDefinition) GetDataSourceOk() (*FormulaAndFunctionSLODataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *FormulaAndFunctionSLOQueryDefinition) SetDataSource(v FormulaAndFunctionSLODataSource) {
	o.DataSource = v
}

// GetGroupMode returns the GroupMode field value if set, zero value otherwise.
func (o *FormulaAndFunctionSLOQueryDefinition) GetGroupMode() FormulaAndFunctionSLOGroupMode {
	if o == nil || o.GroupMode == nil {
		var ret FormulaAndFunctionSLOGroupMode
		return ret
	}
	return *o.GroupMode
}

// GetGroupModeOk returns a tuple with the GroupMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionSLOQueryDefinition) GetGroupModeOk() (*FormulaAndFunctionSLOGroupMode, bool) {
	if o == nil || o.GroupMode == nil {
		return nil, false
	}
	return o.GroupMode, true
}

// HasGroupMode returns a boolean if a field has been set.
func (o *FormulaAndFunctionSLOQueryDefinition) HasGroupMode() bool {
	return o != nil && o.GroupMode != nil
}

// SetGroupMode gets a reference to the given FormulaAndFunctionSLOGroupMode and assigns it to the GroupMode field.
func (o *FormulaAndFunctionSLOQueryDefinition) SetGroupMode(v FormulaAndFunctionSLOGroupMode) {
	o.GroupMode = &v
}

// GetMeasure returns the Measure field value.
func (o *FormulaAndFunctionSLOQueryDefinition) GetMeasure() FormulaAndFunctionSLOMeasure {
	if o == nil {
		var ret FormulaAndFunctionSLOMeasure
		return ret
	}
	return o.Measure
}

// GetMeasureOk returns a tuple with the Measure field value
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionSLOQueryDefinition) GetMeasureOk() (*FormulaAndFunctionSLOMeasure, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Measure, true
}

// SetMeasure sets field value.
func (o *FormulaAndFunctionSLOQueryDefinition) SetMeasure(v FormulaAndFunctionSLOMeasure) {
	o.Measure = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FormulaAndFunctionSLOQueryDefinition) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionSLOQueryDefinition) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FormulaAndFunctionSLOQueryDefinition) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FormulaAndFunctionSLOQueryDefinition) SetName(v string) {
	o.Name = &v
}

// GetSloId returns the SloId field value.
func (o *FormulaAndFunctionSLOQueryDefinition) GetSloId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SloId
}

// GetSloIdOk returns a tuple with the SloId field value
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionSLOQueryDefinition) GetSloIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SloId, true
}

// SetSloId sets field value.
func (o *FormulaAndFunctionSLOQueryDefinition) SetSloId(v string) {
	o.SloId = v
}

// GetSloQueryType returns the SloQueryType field value if set, zero value otherwise.
func (o *FormulaAndFunctionSLOQueryDefinition) GetSloQueryType() FormulaAndFunctionSLOQueryType {
	if o == nil || o.SloQueryType == nil {
		var ret FormulaAndFunctionSLOQueryType
		return ret
	}
	return *o.SloQueryType
}

// GetSloQueryTypeOk returns a tuple with the SloQueryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionSLOQueryDefinition) GetSloQueryTypeOk() (*FormulaAndFunctionSLOQueryType, bool) {
	if o == nil || o.SloQueryType == nil {
		return nil, false
	}
	return o.SloQueryType, true
}

// HasSloQueryType returns a boolean if a field has been set.
func (o *FormulaAndFunctionSLOQueryDefinition) HasSloQueryType() bool {
	return o != nil && o.SloQueryType != nil
}

// SetSloQueryType gets a reference to the given FormulaAndFunctionSLOQueryType and assigns it to the SloQueryType field.
func (o *FormulaAndFunctionSLOQueryDefinition) SetSloQueryType(v FormulaAndFunctionSLOQueryType) {
	o.SloQueryType = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormulaAndFunctionSLOQueryDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AdditionalQueryFilters != nil {
		toSerialize["additional_query_filters"] = o.AdditionalQueryFilters
	}
	if o.CrossOrgUuids != nil {
		toSerialize["cross_org_uuids"] = o.CrossOrgUuids
	}
	toSerialize["data_source"] = o.DataSource
	if o.GroupMode != nil {
		toSerialize["group_mode"] = o.GroupMode
	}
	toSerialize["measure"] = o.Measure
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	toSerialize["slo_id"] = o.SloId
	if o.SloQueryType != nil {
		toSerialize["slo_query_type"] = o.SloQueryType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormulaAndFunctionSLOQueryDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AdditionalQueryFilters *string                          `json:"additional_query_filters,omitempty"`
		CrossOrgUuids          []string                         `json:"cross_org_uuids,omitempty"`
		DataSource             *FormulaAndFunctionSLODataSource `json:"data_source"`
		GroupMode              *FormulaAndFunctionSLOGroupMode  `json:"group_mode,omitempty"`
		Measure                *FormulaAndFunctionSLOMeasure    `json:"measure"`
		Name                   *string                          `json:"name,omitempty"`
		SloId                  *string                          `json:"slo_id"`
		SloQueryType           *FormulaAndFunctionSLOQueryType  `json:"slo_query_type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.Measure == nil {
		return fmt.Errorf("required field measure missing")
	}
	if all.SloId == nil {
		return fmt.Errorf("required field slo_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"additional_query_filters", "cross_org_uuids", "data_source", "group_mode", "measure", "name", "slo_id", "slo_query_type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AdditionalQueryFilters = all.AdditionalQueryFilters
	o.CrossOrgUuids = all.CrossOrgUuids
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	if all.GroupMode != nil && !all.GroupMode.IsValid() {
		hasInvalidField = true
	} else {
		o.GroupMode = all.GroupMode
	}
	if !all.Measure.IsValid() {
		hasInvalidField = true
	} else {
		o.Measure = *all.Measure
	}
	o.Name = all.Name
	o.SloId = *all.SloId
	if all.SloQueryType != nil && !all.SloQueryType.IsValid() {
		hasInvalidField = true
	} else {
		o.SloQueryType = all.SloQueryType
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
