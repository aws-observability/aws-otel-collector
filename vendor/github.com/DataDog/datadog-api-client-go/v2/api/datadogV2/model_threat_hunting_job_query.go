// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ThreatHuntingJobQuery Query for selecting logs analyzed by the threat hunting job.
type ThreatHuntingJobQuery struct {
	// The aggregation type.
	Aggregation *SecurityMonitoringRuleQueryAggregation `json:"aggregation,omitempty"`
	// Source of events, either logs, audit trail, or Datadog events.
	DataSource *SecurityMonitoringStandardDataSource `json:"dataSource,omitempty"`
	// Field for which the cardinality is measured. Sent as an array.
	DistinctFields []string `json:"distinctFields,omitempty"`
	// Fields to group by.
	GroupByFields []string `json:"groupByFields,omitempty"`
	// When false, events without a group-by value are ignored by the query. When true, events with missing group-by fields are processed with `N/A`, replacing the missing values.
	HasOptionalGroupByFields *bool `json:"hasOptionalGroupByFields,omitempty"`
	// Group of target fields to aggregate over when using the sum, max, geo data, or new value aggregations. The sum, max, and geo data aggregations only accept one value in this list, whereas the new value aggregation accepts up to five values.
	Metrics []string `json:"metrics,omitempty"`
	// Name of the query.
	Name *string `json:"name,omitempty"`
	// Query to run on logs.
	Query *string `json:"query,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewThreatHuntingJobQuery instantiates a new ThreatHuntingJobQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewThreatHuntingJobQuery() *ThreatHuntingJobQuery {
	this := ThreatHuntingJobQuery{}
	var dataSource SecurityMonitoringStandardDataSource = SECURITYMONITORINGSTANDARDDATASOURCE_LOGS
	this.DataSource = &dataSource
	var hasOptionalGroupByFields bool = false
	this.HasOptionalGroupByFields = &hasOptionalGroupByFields
	return &this
}

// NewThreatHuntingJobQueryWithDefaults instantiates a new ThreatHuntingJobQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewThreatHuntingJobQueryWithDefaults() *ThreatHuntingJobQuery {
	this := ThreatHuntingJobQuery{}
	var dataSource SecurityMonitoringStandardDataSource = SECURITYMONITORINGSTANDARDDATASOURCE_LOGS
	this.DataSource = &dataSource
	var hasOptionalGroupByFields bool = false
	this.HasOptionalGroupByFields = &hasOptionalGroupByFields
	return &this
}

// GetAggregation returns the Aggregation field value if set, zero value otherwise.
func (o *ThreatHuntingJobQuery) GetAggregation() SecurityMonitoringRuleQueryAggregation {
	if o == nil || o.Aggregation == nil {
		var ret SecurityMonitoringRuleQueryAggregation
		return ret
	}
	return *o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatHuntingJobQuery) GetAggregationOk() (*SecurityMonitoringRuleQueryAggregation, bool) {
	if o == nil || o.Aggregation == nil {
		return nil, false
	}
	return o.Aggregation, true
}

// HasAggregation returns a boolean if a field has been set.
func (o *ThreatHuntingJobQuery) HasAggregation() bool {
	return o != nil && o.Aggregation != nil
}

// SetAggregation gets a reference to the given SecurityMonitoringRuleQueryAggregation and assigns it to the Aggregation field.
func (o *ThreatHuntingJobQuery) SetAggregation(v SecurityMonitoringRuleQueryAggregation) {
	o.Aggregation = &v
}

// GetDataSource returns the DataSource field value if set, zero value otherwise.
func (o *ThreatHuntingJobQuery) GetDataSource() SecurityMonitoringStandardDataSource {
	if o == nil || o.DataSource == nil {
		var ret SecurityMonitoringStandardDataSource
		return ret
	}
	return *o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatHuntingJobQuery) GetDataSourceOk() (*SecurityMonitoringStandardDataSource, bool) {
	if o == nil || o.DataSource == nil {
		return nil, false
	}
	return o.DataSource, true
}

// HasDataSource returns a boolean if a field has been set.
func (o *ThreatHuntingJobQuery) HasDataSource() bool {
	return o != nil && o.DataSource != nil
}

// SetDataSource gets a reference to the given SecurityMonitoringStandardDataSource and assigns it to the DataSource field.
func (o *ThreatHuntingJobQuery) SetDataSource(v SecurityMonitoringStandardDataSource) {
	o.DataSource = &v
}

// GetDistinctFields returns the DistinctFields field value if set, zero value otherwise.
func (o *ThreatHuntingJobQuery) GetDistinctFields() []string {
	if o == nil || o.DistinctFields == nil {
		var ret []string
		return ret
	}
	return o.DistinctFields
}

// GetDistinctFieldsOk returns a tuple with the DistinctFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatHuntingJobQuery) GetDistinctFieldsOk() (*[]string, bool) {
	if o == nil || o.DistinctFields == nil {
		return nil, false
	}
	return &o.DistinctFields, true
}

// HasDistinctFields returns a boolean if a field has been set.
func (o *ThreatHuntingJobQuery) HasDistinctFields() bool {
	return o != nil && o.DistinctFields != nil
}

// SetDistinctFields gets a reference to the given []string and assigns it to the DistinctFields field.
func (o *ThreatHuntingJobQuery) SetDistinctFields(v []string) {
	o.DistinctFields = v
}

// GetGroupByFields returns the GroupByFields field value if set, zero value otherwise.
func (o *ThreatHuntingJobQuery) GetGroupByFields() []string {
	if o == nil || o.GroupByFields == nil {
		var ret []string
		return ret
	}
	return o.GroupByFields
}

// GetGroupByFieldsOk returns a tuple with the GroupByFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatHuntingJobQuery) GetGroupByFieldsOk() (*[]string, bool) {
	if o == nil || o.GroupByFields == nil {
		return nil, false
	}
	return &o.GroupByFields, true
}

// HasGroupByFields returns a boolean if a field has been set.
func (o *ThreatHuntingJobQuery) HasGroupByFields() bool {
	return o != nil && o.GroupByFields != nil
}

// SetGroupByFields gets a reference to the given []string and assigns it to the GroupByFields field.
func (o *ThreatHuntingJobQuery) SetGroupByFields(v []string) {
	o.GroupByFields = v
}

// GetHasOptionalGroupByFields returns the HasOptionalGroupByFields field value if set, zero value otherwise.
func (o *ThreatHuntingJobQuery) GetHasOptionalGroupByFields() bool {
	if o == nil || o.HasOptionalGroupByFields == nil {
		var ret bool
		return ret
	}
	return *o.HasOptionalGroupByFields
}

// GetHasOptionalGroupByFieldsOk returns a tuple with the HasOptionalGroupByFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatHuntingJobQuery) GetHasOptionalGroupByFieldsOk() (*bool, bool) {
	if o == nil || o.HasOptionalGroupByFields == nil {
		return nil, false
	}
	return o.HasOptionalGroupByFields, true
}

// HasHasOptionalGroupByFields returns a boolean if a field has been set.
func (o *ThreatHuntingJobQuery) HasHasOptionalGroupByFields() bool {
	return o != nil && o.HasOptionalGroupByFields != nil
}

// SetHasOptionalGroupByFields gets a reference to the given bool and assigns it to the HasOptionalGroupByFields field.
func (o *ThreatHuntingJobQuery) SetHasOptionalGroupByFields(v bool) {
	o.HasOptionalGroupByFields = &v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *ThreatHuntingJobQuery) GetMetrics() []string {
	if o == nil || o.Metrics == nil {
		var ret []string
		return ret
	}
	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatHuntingJobQuery) GetMetricsOk() (*[]string, bool) {
	if o == nil || o.Metrics == nil {
		return nil, false
	}
	return &o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *ThreatHuntingJobQuery) HasMetrics() bool {
	return o != nil && o.Metrics != nil
}

// SetMetrics gets a reference to the given []string and assigns it to the Metrics field.
func (o *ThreatHuntingJobQuery) SetMetrics(v []string) {
	o.Metrics = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ThreatHuntingJobQuery) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatHuntingJobQuery) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ThreatHuntingJobQuery) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ThreatHuntingJobQuery) SetName(v string) {
	o.Name = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *ThreatHuntingJobQuery) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatHuntingJobQuery) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *ThreatHuntingJobQuery) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *ThreatHuntingJobQuery) SetQuery(v string) {
	o.Query = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ThreatHuntingJobQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Aggregation != nil {
		toSerialize["aggregation"] = o.Aggregation
	}
	if o.DataSource != nil {
		toSerialize["dataSource"] = o.DataSource
	}
	if o.DistinctFields != nil {
		toSerialize["distinctFields"] = o.DistinctFields
	}
	if o.GroupByFields != nil {
		toSerialize["groupByFields"] = o.GroupByFields
	}
	if o.HasOptionalGroupByFields != nil {
		toSerialize["hasOptionalGroupByFields"] = o.HasOptionalGroupByFields
	}
	if o.Metrics != nil {
		toSerialize["metrics"] = o.Metrics
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ThreatHuntingJobQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregation              *SecurityMonitoringRuleQueryAggregation `json:"aggregation,omitempty"`
		DataSource               *SecurityMonitoringStandardDataSource   `json:"dataSource,omitempty"`
		DistinctFields           []string                                `json:"distinctFields,omitempty"`
		GroupByFields            []string                                `json:"groupByFields,omitempty"`
		HasOptionalGroupByFields *bool                                   `json:"hasOptionalGroupByFields,omitempty"`
		Metrics                  []string                                `json:"metrics,omitempty"`
		Name                     *string                                 `json:"name,omitempty"`
		Query                    *string                                 `json:"query,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aggregation", "dataSource", "distinctFields", "groupByFields", "hasOptionalGroupByFields", "metrics", "name", "query"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Aggregation != nil && !all.Aggregation.IsValid() {
		hasInvalidField = true
	} else {
		o.Aggregation = all.Aggregation
	}
	if all.DataSource != nil && !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = all.DataSource
	}
	o.DistinctFields = all.DistinctFields
	o.GroupByFields = all.GroupByFields
	o.HasOptionalGroupByFields = all.HasOptionalGroupByFields
	o.Metrics = all.Metrics
	o.Name = all.Name
	o.Query = all.Query

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
