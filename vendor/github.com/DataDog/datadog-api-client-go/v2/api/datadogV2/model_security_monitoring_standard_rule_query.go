// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringStandardRuleQuery Query for matching rule.
type SecurityMonitoringStandardRuleQuery struct {
	// The aggregation type.
	Aggregation *SecurityMonitoringRuleQueryAggregation `json:"aggregation,omitempty"`
	// Query extension to append to the logs query.
	CustomQueryExtension *string `json:"customQueryExtension,omitempty"`
	// Source of events, either logs, audit trail, or Datadog events.
	DataSource *SecurityMonitoringStandardDataSource `json:"dataSource,omitempty"`
	// Field for which the cardinality is measured. Sent as an array.
	DistinctFields []string `json:"distinctFields,omitempty"`
	// Fields to group by.
	GroupByFields []string `json:"groupByFields,omitempty"`
	// When false, events without a group-by value are ignored by the rule. When true, events with missing group-by fields are processed with `N/A`, replacing the missing values.
	HasOptionalGroupByFields *bool `json:"hasOptionalGroupByFields,omitempty"`
	// **This field is currently unstable and might be removed in a minor version upgrade.**
	// The index to run the query on, if the `dataSource` is `logs`. Only used for scheduled rules - in other words, when the `schedulingOptions` field is present in the rule payload.
	Index *string `json:"index,omitempty"`
	// List of indexes to query when the `dataSource` is `logs`. Only used for scheduled rules, such as when the `schedulingOptions` field is present in the rule payload.
	Indexes []string `json:"indexes,omitempty"`
	// (Deprecated) The target field to aggregate over when using the sum or max
	// aggregations. `metrics` field should be used instead.
	// Deprecated
	Metric *string `json:"metric,omitempty"`
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

// NewSecurityMonitoringStandardRuleQuery instantiates a new SecurityMonitoringStandardRuleQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringStandardRuleQuery() *SecurityMonitoringStandardRuleQuery {
	this := SecurityMonitoringStandardRuleQuery{}
	var dataSource SecurityMonitoringStandardDataSource = SECURITYMONITORINGSTANDARDDATASOURCE_LOGS
	this.DataSource = &dataSource
	var hasOptionalGroupByFields bool = false
	this.HasOptionalGroupByFields = &hasOptionalGroupByFields
	return &this
}

// NewSecurityMonitoringStandardRuleQueryWithDefaults instantiates a new SecurityMonitoringStandardRuleQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringStandardRuleQueryWithDefaults() *SecurityMonitoringStandardRuleQuery {
	this := SecurityMonitoringStandardRuleQuery{}
	var dataSource SecurityMonitoringStandardDataSource = SECURITYMONITORINGSTANDARDDATASOURCE_LOGS
	this.DataSource = &dataSource
	var hasOptionalGroupByFields bool = false
	this.HasOptionalGroupByFields = &hasOptionalGroupByFields
	return &this
}

// GetAggregation returns the Aggregation field value if set, zero value otherwise.
func (o *SecurityMonitoringStandardRuleQuery) GetAggregation() SecurityMonitoringRuleQueryAggregation {
	if o == nil || o.Aggregation == nil {
		var ret SecurityMonitoringRuleQueryAggregation
		return ret
	}
	return *o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringStandardRuleQuery) GetAggregationOk() (*SecurityMonitoringRuleQueryAggregation, bool) {
	if o == nil || o.Aggregation == nil {
		return nil, false
	}
	return o.Aggregation, true
}

// HasAggregation returns a boolean if a field has been set.
func (o *SecurityMonitoringStandardRuleQuery) HasAggregation() bool {
	return o != nil && o.Aggregation != nil
}

// SetAggregation gets a reference to the given SecurityMonitoringRuleQueryAggregation and assigns it to the Aggregation field.
func (o *SecurityMonitoringStandardRuleQuery) SetAggregation(v SecurityMonitoringRuleQueryAggregation) {
	o.Aggregation = &v
}

// GetCustomQueryExtension returns the CustomQueryExtension field value if set, zero value otherwise.
func (o *SecurityMonitoringStandardRuleQuery) GetCustomQueryExtension() string {
	if o == nil || o.CustomQueryExtension == nil {
		var ret string
		return ret
	}
	return *o.CustomQueryExtension
}

// GetCustomQueryExtensionOk returns a tuple with the CustomQueryExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringStandardRuleQuery) GetCustomQueryExtensionOk() (*string, bool) {
	if o == nil || o.CustomQueryExtension == nil {
		return nil, false
	}
	return o.CustomQueryExtension, true
}

// HasCustomQueryExtension returns a boolean if a field has been set.
func (o *SecurityMonitoringStandardRuleQuery) HasCustomQueryExtension() bool {
	return o != nil && o.CustomQueryExtension != nil
}

// SetCustomQueryExtension gets a reference to the given string and assigns it to the CustomQueryExtension field.
func (o *SecurityMonitoringStandardRuleQuery) SetCustomQueryExtension(v string) {
	o.CustomQueryExtension = &v
}

// GetDataSource returns the DataSource field value if set, zero value otherwise.
func (o *SecurityMonitoringStandardRuleQuery) GetDataSource() SecurityMonitoringStandardDataSource {
	if o == nil || o.DataSource == nil {
		var ret SecurityMonitoringStandardDataSource
		return ret
	}
	return *o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringStandardRuleQuery) GetDataSourceOk() (*SecurityMonitoringStandardDataSource, bool) {
	if o == nil || o.DataSource == nil {
		return nil, false
	}
	return o.DataSource, true
}

// HasDataSource returns a boolean if a field has been set.
func (o *SecurityMonitoringStandardRuleQuery) HasDataSource() bool {
	return o != nil && o.DataSource != nil
}

// SetDataSource gets a reference to the given SecurityMonitoringStandardDataSource and assigns it to the DataSource field.
func (o *SecurityMonitoringStandardRuleQuery) SetDataSource(v SecurityMonitoringStandardDataSource) {
	o.DataSource = &v
}

// GetDistinctFields returns the DistinctFields field value if set, zero value otherwise.
func (o *SecurityMonitoringStandardRuleQuery) GetDistinctFields() []string {
	if o == nil || o.DistinctFields == nil {
		var ret []string
		return ret
	}
	return o.DistinctFields
}

// GetDistinctFieldsOk returns a tuple with the DistinctFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringStandardRuleQuery) GetDistinctFieldsOk() (*[]string, bool) {
	if o == nil || o.DistinctFields == nil {
		return nil, false
	}
	return &o.DistinctFields, true
}

// HasDistinctFields returns a boolean if a field has been set.
func (o *SecurityMonitoringStandardRuleQuery) HasDistinctFields() bool {
	return o != nil && o.DistinctFields != nil
}

// SetDistinctFields gets a reference to the given []string and assigns it to the DistinctFields field.
func (o *SecurityMonitoringStandardRuleQuery) SetDistinctFields(v []string) {
	o.DistinctFields = v
}

// GetGroupByFields returns the GroupByFields field value if set, zero value otherwise.
func (o *SecurityMonitoringStandardRuleQuery) GetGroupByFields() []string {
	if o == nil || o.GroupByFields == nil {
		var ret []string
		return ret
	}
	return o.GroupByFields
}

// GetGroupByFieldsOk returns a tuple with the GroupByFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringStandardRuleQuery) GetGroupByFieldsOk() (*[]string, bool) {
	if o == nil || o.GroupByFields == nil {
		return nil, false
	}
	return &o.GroupByFields, true
}

// HasGroupByFields returns a boolean if a field has been set.
func (o *SecurityMonitoringStandardRuleQuery) HasGroupByFields() bool {
	return o != nil && o.GroupByFields != nil
}

// SetGroupByFields gets a reference to the given []string and assigns it to the GroupByFields field.
func (o *SecurityMonitoringStandardRuleQuery) SetGroupByFields(v []string) {
	o.GroupByFields = v
}

// GetHasOptionalGroupByFields returns the HasOptionalGroupByFields field value if set, zero value otherwise.
func (o *SecurityMonitoringStandardRuleQuery) GetHasOptionalGroupByFields() bool {
	if o == nil || o.HasOptionalGroupByFields == nil {
		var ret bool
		return ret
	}
	return *o.HasOptionalGroupByFields
}

// GetHasOptionalGroupByFieldsOk returns a tuple with the HasOptionalGroupByFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringStandardRuleQuery) GetHasOptionalGroupByFieldsOk() (*bool, bool) {
	if o == nil || o.HasOptionalGroupByFields == nil {
		return nil, false
	}
	return o.HasOptionalGroupByFields, true
}

// HasHasOptionalGroupByFields returns a boolean if a field has been set.
func (o *SecurityMonitoringStandardRuleQuery) HasHasOptionalGroupByFields() bool {
	return o != nil && o.HasOptionalGroupByFields != nil
}

// SetHasOptionalGroupByFields gets a reference to the given bool and assigns it to the HasOptionalGroupByFields field.
func (o *SecurityMonitoringStandardRuleQuery) SetHasOptionalGroupByFields(v bool) {
	o.HasOptionalGroupByFields = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *SecurityMonitoringStandardRuleQuery) GetIndex() string {
	if o == nil || o.Index == nil {
		var ret string
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringStandardRuleQuery) GetIndexOk() (*string, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *SecurityMonitoringStandardRuleQuery) HasIndex() bool {
	return o != nil && o.Index != nil
}

// SetIndex gets a reference to the given string and assigns it to the Index field.
func (o *SecurityMonitoringStandardRuleQuery) SetIndex(v string) {
	o.Index = &v
}

// GetIndexes returns the Indexes field value if set, zero value otherwise.
func (o *SecurityMonitoringStandardRuleQuery) GetIndexes() []string {
	if o == nil || o.Indexes == nil {
		var ret []string
		return ret
	}
	return o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringStandardRuleQuery) GetIndexesOk() (*[]string, bool) {
	if o == nil || o.Indexes == nil {
		return nil, false
	}
	return &o.Indexes, true
}

// HasIndexes returns a boolean if a field has been set.
func (o *SecurityMonitoringStandardRuleQuery) HasIndexes() bool {
	return o != nil && o.Indexes != nil
}

// SetIndexes gets a reference to the given []string and assigns it to the Indexes field.
func (o *SecurityMonitoringStandardRuleQuery) SetIndexes(v []string) {
	o.Indexes = v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
// Deprecated
func (o *SecurityMonitoringStandardRuleQuery) GetMetric() string {
	if o == nil || o.Metric == nil {
		var ret string
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *SecurityMonitoringStandardRuleQuery) GetMetricOk() (*string, bool) {
	if o == nil || o.Metric == nil {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *SecurityMonitoringStandardRuleQuery) HasMetric() bool {
	return o != nil && o.Metric != nil
}

// SetMetric gets a reference to the given string and assigns it to the Metric field.
// Deprecated
func (o *SecurityMonitoringStandardRuleQuery) SetMetric(v string) {
	o.Metric = &v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *SecurityMonitoringStandardRuleQuery) GetMetrics() []string {
	if o == nil || o.Metrics == nil {
		var ret []string
		return ret
	}
	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringStandardRuleQuery) GetMetricsOk() (*[]string, bool) {
	if o == nil || o.Metrics == nil {
		return nil, false
	}
	return &o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *SecurityMonitoringStandardRuleQuery) HasMetrics() bool {
	return o != nil && o.Metrics != nil
}

// SetMetrics gets a reference to the given []string and assigns it to the Metrics field.
func (o *SecurityMonitoringStandardRuleQuery) SetMetrics(v []string) {
	o.Metrics = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SecurityMonitoringStandardRuleQuery) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringStandardRuleQuery) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SecurityMonitoringStandardRuleQuery) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SecurityMonitoringStandardRuleQuery) SetName(v string) {
	o.Name = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *SecurityMonitoringStandardRuleQuery) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringStandardRuleQuery) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *SecurityMonitoringStandardRuleQuery) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *SecurityMonitoringStandardRuleQuery) SetQuery(v string) {
	o.Query = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringStandardRuleQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Aggregation != nil {
		toSerialize["aggregation"] = o.Aggregation
	}
	if o.CustomQueryExtension != nil {
		toSerialize["customQueryExtension"] = o.CustomQueryExtension
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
	if o.Index != nil {
		toSerialize["index"] = o.Index
	}
	if o.Indexes != nil {
		toSerialize["indexes"] = o.Indexes
	}
	if o.Metric != nil {
		toSerialize["metric"] = o.Metric
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
func (o *SecurityMonitoringStandardRuleQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregation              *SecurityMonitoringRuleQueryAggregation `json:"aggregation,omitempty"`
		CustomQueryExtension     *string                                 `json:"customQueryExtension,omitempty"`
		DataSource               *SecurityMonitoringStandardDataSource   `json:"dataSource,omitempty"`
		DistinctFields           []string                                `json:"distinctFields,omitempty"`
		GroupByFields            []string                                `json:"groupByFields,omitempty"`
		HasOptionalGroupByFields *bool                                   `json:"hasOptionalGroupByFields,omitempty"`
		Index                    *string                                 `json:"index,omitempty"`
		Indexes                  []string                                `json:"indexes,omitempty"`
		Metric                   *string                                 `json:"metric,omitempty"`
		Metrics                  []string                                `json:"metrics,omitempty"`
		Name                     *string                                 `json:"name,omitempty"`
		Query                    *string                                 `json:"query,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aggregation", "customQueryExtension", "dataSource", "distinctFields", "groupByFields", "hasOptionalGroupByFields", "index", "indexes", "metric", "metrics", "name", "query"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Aggregation != nil && !all.Aggregation.IsValid() {
		hasInvalidField = true
	} else {
		o.Aggregation = all.Aggregation
	}
	o.CustomQueryExtension = all.CustomQueryExtension
	if all.DataSource != nil && !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = all.DataSource
	}
	o.DistinctFields = all.DistinctFields
	o.GroupByFields = all.GroupByFields
	o.HasOptionalGroupByFields = all.HasOptionalGroupByFields
	o.Index = all.Index
	o.Indexes = all.Indexes
	o.Metric = all.Metric
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
