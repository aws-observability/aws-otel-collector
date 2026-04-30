// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SloQuery A query for SLO status, error budget, and burn rate metrics.
type SloQuery struct {
	// Additional filters applied to the SLO query.
	AdditionalQueryFilters *string `json:"additional_query_filters,omitempty"`
	// A data source for SLO queries.
	DataSource SloDataSource `json:"data_source"`
	// How SLO results are grouped in the response.
	GroupMode *SlosGroupMode `json:"group_mode,omitempty"`
	// The SLO measurement to retrieve.
	Measure SlosMeasure `json:"measure"`
	// The variable name for use in formulas.
	Name *string `json:"name,omitempty"`
	// The unique identifier of the SLO to query.
	SloId string `json:"slo_id"`
	// The type of SLO definition being queried.
	SloQueryType *SlosQueryType `json:"slo_query_type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSloQuery instantiates a new SloQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSloQuery(dataSource SloDataSource, measure SlosMeasure, sloId string) *SloQuery {
	this := SloQuery{}
	this.DataSource = dataSource
	this.Measure = measure
	this.SloId = sloId
	return &this
}

// NewSloQueryWithDefaults instantiates a new SloQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSloQueryWithDefaults() *SloQuery {
	this := SloQuery{}
	var dataSource SloDataSource = SLODATASOURCE_SLO
	this.DataSource = dataSource
	return &this
}

// GetAdditionalQueryFilters returns the AdditionalQueryFilters field value if set, zero value otherwise.
func (o *SloQuery) GetAdditionalQueryFilters() string {
	if o == nil || o.AdditionalQueryFilters == nil {
		var ret string
		return ret
	}
	return *o.AdditionalQueryFilters
}

// GetAdditionalQueryFiltersOk returns a tuple with the AdditionalQueryFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloQuery) GetAdditionalQueryFiltersOk() (*string, bool) {
	if o == nil || o.AdditionalQueryFilters == nil {
		return nil, false
	}
	return o.AdditionalQueryFilters, true
}

// HasAdditionalQueryFilters returns a boolean if a field has been set.
func (o *SloQuery) HasAdditionalQueryFilters() bool {
	return o != nil && o.AdditionalQueryFilters != nil
}

// SetAdditionalQueryFilters gets a reference to the given string and assigns it to the AdditionalQueryFilters field.
func (o *SloQuery) SetAdditionalQueryFilters(v string) {
	o.AdditionalQueryFilters = &v
}

// GetDataSource returns the DataSource field value.
func (o *SloQuery) GetDataSource() SloDataSource {
	if o == nil {
		var ret SloDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *SloQuery) GetDataSourceOk() (*SloDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *SloQuery) SetDataSource(v SloDataSource) {
	o.DataSource = v
}

// GetGroupMode returns the GroupMode field value if set, zero value otherwise.
func (o *SloQuery) GetGroupMode() SlosGroupMode {
	if o == nil || o.GroupMode == nil {
		var ret SlosGroupMode
		return ret
	}
	return *o.GroupMode
}

// GetGroupModeOk returns a tuple with the GroupMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloQuery) GetGroupModeOk() (*SlosGroupMode, bool) {
	if o == nil || o.GroupMode == nil {
		return nil, false
	}
	return o.GroupMode, true
}

// HasGroupMode returns a boolean if a field has been set.
func (o *SloQuery) HasGroupMode() bool {
	return o != nil && o.GroupMode != nil
}

// SetGroupMode gets a reference to the given SlosGroupMode and assigns it to the GroupMode field.
func (o *SloQuery) SetGroupMode(v SlosGroupMode) {
	o.GroupMode = &v
}

// GetMeasure returns the Measure field value.
func (o *SloQuery) GetMeasure() SlosMeasure {
	if o == nil {
		var ret SlosMeasure
		return ret
	}
	return o.Measure
}

// GetMeasureOk returns a tuple with the Measure field value
// and a boolean to check if the value has been set.
func (o *SloQuery) GetMeasureOk() (*SlosMeasure, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Measure, true
}

// SetMeasure sets field value.
func (o *SloQuery) SetMeasure(v SlosMeasure) {
	o.Measure = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SloQuery) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloQuery) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SloQuery) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SloQuery) SetName(v string) {
	o.Name = &v
}

// GetSloId returns the SloId field value.
func (o *SloQuery) GetSloId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SloId
}

// GetSloIdOk returns a tuple with the SloId field value
// and a boolean to check if the value has been set.
func (o *SloQuery) GetSloIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SloId, true
}

// SetSloId sets field value.
func (o *SloQuery) SetSloId(v string) {
	o.SloId = v
}

// GetSloQueryType returns the SloQueryType field value if set, zero value otherwise.
func (o *SloQuery) GetSloQueryType() SlosQueryType {
	if o == nil || o.SloQueryType == nil {
		var ret SlosQueryType
		return ret
	}
	return *o.SloQueryType
}

// GetSloQueryTypeOk returns a tuple with the SloQueryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloQuery) GetSloQueryTypeOk() (*SlosQueryType, bool) {
	if o == nil || o.SloQueryType == nil {
		return nil, false
	}
	return o.SloQueryType, true
}

// HasSloQueryType returns a boolean if a field has been set.
func (o *SloQuery) HasSloQueryType() bool {
	return o != nil && o.SloQueryType != nil
}

// SetSloQueryType gets a reference to the given SlosQueryType and assigns it to the SloQueryType field.
func (o *SloQuery) SetSloQueryType(v SlosQueryType) {
	o.SloQueryType = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SloQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AdditionalQueryFilters != nil {
		toSerialize["additional_query_filters"] = o.AdditionalQueryFilters
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
func (o *SloQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AdditionalQueryFilters *string        `json:"additional_query_filters,omitempty"`
		DataSource             *SloDataSource `json:"data_source"`
		GroupMode              *SlosGroupMode `json:"group_mode,omitempty"`
		Measure                *SlosMeasure   `json:"measure"`
		Name                   *string        `json:"name,omitempty"`
		SloId                  *string        `json:"slo_id"`
		SloQueryType           *SlosQueryType `json:"slo_query_type,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"additional_query_filters", "data_source", "group_mode", "measure", "name", "slo_id", "slo_query_type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AdditionalQueryFilters = all.AdditionalQueryFilters
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
