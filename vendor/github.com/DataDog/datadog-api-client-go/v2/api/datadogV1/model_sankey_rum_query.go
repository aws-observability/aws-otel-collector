// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SankeyRumQuery Sankey widget with RUM data source query.
type SankeyRumQuery struct {
	// Product Analytics/RUM audience filters.
	AudienceFilters *ProductAnalyticsAudienceFilters `json:"audience_filters,omitempty"`
	// Sankey widget with RUM data source.
	DataSource SankeyRumDataSource `json:"data_source"`
	// Entries per step.
	EntriesPerStep *int64 `json:"entries_per_step,omitempty"`
	// Join keys.
	JoinKeys *SankeyJoinKeys `json:"join_keys,omitempty"`
	// Sankey mode for RUM queries.
	Mode SankeyRumQueryMode `json:"mode"`
	// Number of steps.
	NumberOfSteps *int64 `json:"number_of_steps,omitempty"`
	// Filter applied to occurrence counts when building a Product Analytics audience.
	Occurrences *ProductAnalyticsAudienceOccurrenceFilter `json:"occurrences,omitempty"`
	// RUM event search query used to filter views or actions.
	QueryString string `json:"query_string"`
	// Source.
	Source *string `json:"source,omitempty"`
	// Subquery ID.
	SubqueryId *string `json:"subquery_id,omitempty"`
	// Target.
	Target *string `json:"target,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewSankeyRumQuery instantiates a new SankeyRumQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSankeyRumQuery(dataSource SankeyRumDataSource, mode SankeyRumQueryMode, queryString string) *SankeyRumQuery {
	this := SankeyRumQuery{}
	this.DataSource = dataSource
	this.Mode = mode
	this.QueryString = queryString
	return &this
}

// NewSankeyRumQueryWithDefaults instantiates a new SankeyRumQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSankeyRumQueryWithDefaults() *SankeyRumQuery {
	this := SankeyRumQuery{}
	var dataSource SankeyRumDataSource = SANKEYRUMDATASOURCE_RUM
	this.DataSource = dataSource
	var mode SankeyRumQueryMode = SANKEYRUMQUERYMODE_SOURCE
	this.Mode = mode
	return &this
}

// GetAudienceFilters returns the AudienceFilters field value if set, zero value otherwise.
func (o *SankeyRumQuery) GetAudienceFilters() ProductAnalyticsAudienceFilters {
	if o == nil || o.AudienceFilters == nil {
		var ret ProductAnalyticsAudienceFilters
		return ret
	}
	return *o.AudienceFilters
}

// GetAudienceFiltersOk returns a tuple with the AudienceFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyRumQuery) GetAudienceFiltersOk() (*ProductAnalyticsAudienceFilters, bool) {
	if o == nil || o.AudienceFilters == nil {
		return nil, false
	}
	return o.AudienceFilters, true
}

// HasAudienceFilters returns a boolean if a field has been set.
func (o *SankeyRumQuery) HasAudienceFilters() bool {
	return o != nil && o.AudienceFilters != nil
}

// SetAudienceFilters gets a reference to the given ProductAnalyticsAudienceFilters and assigns it to the AudienceFilters field.
func (o *SankeyRumQuery) SetAudienceFilters(v ProductAnalyticsAudienceFilters) {
	o.AudienceFilters = &v
}

// GetDataSource returns the DataSource field value.
func (o *SankeyRumQuery) GetDataSource() SankeyRumDataSource {
	if o == nil {
		var ret SankeyRumDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *SankeyRumQuery) GetDataSourceOk() (*SankeyRumDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *SankeyRumQuery) SetDataSource(v SankeyRumDataSource) {
	o.DataSource = v
}

// GetEntriesPerStep returns the EntriesPerStep field value if set, zero value otherwise.
func (o *SankeyRumQuery) GetEntriesPerStep() int64 {
	if o == nil || o.EntriesPerStep == nil {
		var ret int64
		return ret
	}
	return *o.EntriesPerStep
}

// GetEntriesPerStepOk returns a tuple with the EntriesPerStep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyRumQuery) GetEntriesPerStepOk() (*int64, bool) {
	if o == nil || o.EntriesPerStep == nil {
		return nil, false
	}
	return o.EntriesPerStep, true
}

// HasEntriesPerStep returns a boolean if a field has been set.
func (o *SankeyRumQuery) HasEntriesPerStep() bool {
	return o != nil && o.EntriesPerStep != nil
}

// SetEntriesPerStep gets a reference to the given int64 and assigns it to the EntriesPerStep field.
func (o *SankeyRumQuery) SetEntriesPerStep(v int64) {
	o.EntriesPerStep = &v
}

// GetJoinKeys returns the JoinKeys field value if set, zero value otherwise.
func (o *SankeyRumQuery) GetJoinKeys() SankeyJoinKeys {
	if o == nil || o.JoinKeys == nil {
		var ret SankeyJoinKeys
		return ret
	}
	return *o.JoinKeys
}

// GetJoinKeysOk returns a tuple with the JoinKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyRumQuery) GetJoinKeysOk() (*SankeyJoinKeys, bool) {
	if o == nil || o.JoinKeys == nil {
		return nil, false
	}
	return o.JoinKeys, true
}

// HasJoinKeys returns a boolean if a field has been set.
func (o *SankeyRumQuery) HasJoinKeys() bool {
	return o != nil && o.JoinKeys != nil
}

// SetJoinKeys gets a reference to the given SankeyJoinKeys and assigns it to the JoinKeys field.
func (o *SankeyRumQuery) SetJoinKeys(v SankeyJoinKeys) {
	o.JoinKeys = &v
}

// GetMode returns the Mode field value.
func (o *SankeyRumQuery) GetMode() SankeyRumQueryMode {
	if o == nil {
		var ret SankeyRumQueryMode
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *SankeyRumQuery) GetModeOk() (*SankeyRumQueryMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *SankeyRumQuery) SetMode(v SankeyRumQueryMode) {
	o.Mode = v
}

// GetNumberOfSteps returns the NumberOfSteps field value if set, zero value otherwise.
func (o *SankeyRumQuery) GetNumberOfSteps() int64 {
	if o == nil || o.NumberOfSteps == nil {
		var ret int64
		return ret
	}
	return *o.NumberOfSteps
}

// GetNumberOfStepsOk returns a tuple with the NumberOfSteps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyRumQuery) GetNumberOfStepsOk() (*int64, bool) {
	if o == nil || o.NumberOfSteps == nil {
		return nil, false
	}
	return o.NumberOfSteps, true
}

// HasNumberOfSteps returns a boolean if a field has been set.
func (o *SankeyRumQuery) HasNumberOfSteps() bool {
	return o != nil && o.NumberOfSteps != nil
}

// SetNumberOfSteps gets a reference to the given int64 and assigns it to the NumberOfSteps field.
func (o *SankeyRumQuery) SetNumberOfSteps(v int64) {
	o.NumberOfSteps = &v
}

// GetOccurrences returns the Occurrences field value if set, zero value otherwise.
func (o *SankeyRumQuery) GetOccurrences() ProductAnalyticsAudienceOccurrenceFilter {
	if o == nil || o.Occurrences == nil {
		var ret ProductAnalyticsAudienceOccurrenceFilter
		return ret
	}
	return *o.Occurrences
}

// GetOccurrencesOk returns a tuple with the Occurrences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyRumQuery) GetOccurrencesOk() (*ProductAnalyticsAudienceOccurrenceFilter, bool) {
	if o == nil || o.Occurrences == nil {
		return nil, false
	}
	return o.Occurrences, true
}

// HasOccurrences returns a boolean if a field has been set.
func (o *SankeyRumQuery) HasOccurrences() bool {
	return o != nil && o.Occurrences != nil
}

// SetOccurrences gets a reference to the given ProductAnalyticsAudienceOccurrenceFilter and assigns it to the Occurrences field.
func (o *SankeyRumQuery) SetOccurrences(v ProductAnalyticsAudienceOccurrenceFilter) {
	o.Occurrences = &v
}

// GetQueryString returns the QueryString field value.
func (o *SankeyRumQuery) GetQueryString() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value
// and a boolean to check if the value has been set.
func (o *SankeyRumQuery) GetQueryStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryString, true
}

// SetQueryString sets field value.
func (o *SankeyRumQuery) SetQueryString(v string) {
	o.QueryString = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *SankeyRumQuery) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyRumQuery) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *SankeyRumQuery) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *SankeyRumQuery) SetSource(v string) {
	o.Source = &v
}

// GetSubqueryId returns the SubqueryId field value if set, zero value otherwise.
func (o *SankeyRumQuery) GetSubqueryId() string {
	if o == nil || o.SubqueryId == nil {
		var ret string
		return ret
	}
	return *o.SubqueryId
}

// GetSubqueryIdOk returns a tuple with the SubqueryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyRumQuery) GetSubqueryIdOk() (*string, bool) {
	if o == nil || o.SubqueryId == nil {
		return nil, false
	}
	return o.SubqueryId, true
}

// HasSubqueryId returns a boolean if a field has been set.
func (o *SankeyRumQuery) HasSubqueryId() bool {
	return o != nil && o.SubqueryId != nil
}

// SetSubqueryId gets a reference to the given string and assigns it to the SubqueryId field.
func (o *SankeyRumQuery) SetSubqueryId(v string) {
	o.SubqueryId = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *SankeyRumQuery) GetTarget() string {
	if o == nil || o.Target == nil {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyRumQuery) GetTargetOk() (*string, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *SankeyRumQuery) HasTarget() bool {
	return o != nil && o.Target != nil
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *SankeyRumQuery) SetTarget(v string) {
	o.Target = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SankeyRumQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AudienceFilters != nil {
		toSerialize["audience_filters"] = o.AudienceFilters
	}
	toSerialize["data_source"] = o.DataSource
	if o.EntriesPerStep != nil {
		toSerialize["entries_per_step"] = o.EntriesPerStep
	}
	if o.JoinKeys != nil {
		toSerialize["join_keys"] = o.JoinKeys
	}
	toSerialize["mode"] = o.Mode
	if o.NumberOfSteps != nil {
		toSerialize["number_of_steps"] = o.NumberOfSteps
	}
	if o.Occurrences != nil {
		toSerialize["occurrences"] = o.Occurrences
	}
	toSerialize["query_string"] = o.QueryString
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.SubqueryId != nil {
		toSerialize["subquery_id"] = o.SubqueryId
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SankeyRumQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AudienceFilters *ProductAnalyticsAudienceFilters          `json:"audience_filters,omitempty"`
		DataSource      *SankeyRumDataSource                      `json:"data_source"`
		EntriesPerStep  *int64                                    `json:"entries_per_step,omitempty"`
		JoinKeys        *SankeyJoinKeys                           `json:"join_keys,omitempty"`
		Mode            *SankeyRumQueryMode                       `json:"mode"`
		NumberOfSteps   *int64                                    `json:"number_of_steps,omitempty"`
		Occurrences     *ProductAnalyticsAudienceOccurrenceFilter `json:"occurrences,omitempty"`
		QueryString     *string                                   `json:"query_string"`
		Source          *string                                   `json:"source,omitempty"`
		SubqueryId      *string                                   `json:"subquery_id,omitempty"`
		Target          *string                                   `json:"target,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.Mode == nil {
		return fmt.Errorf("required field mode missing")
	}
	if all.QueryString == nil {
		return fmt.Errorf("required field query_string missing")
	}

	hasInvalidField := false
	if all.AudienceFilters != nil && all.AudienceFilters.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AudienceFilters = all.AudienceFilters
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.EntriesPerStep = all.EntriesPerStep
	if all.JoinKeys != nil && all.JoinKeys.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.JoinKeys = all.JoinKeys
	if !all.Mode.IsValid() {
		hasInvalidField = true
	} else {
		o.Mode = *all.Mode
	}
	o.NumberOfSteps = all.NumberOfSteps
	if all.Occurrences != nil && all.Occurrences.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Occurrences = all.Occurrences
	o.QueryString = *all.QueryString
	o.Source = all.Source
	o.SubqueryId = all.SubqueryId
	o.Target = all.Target

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
