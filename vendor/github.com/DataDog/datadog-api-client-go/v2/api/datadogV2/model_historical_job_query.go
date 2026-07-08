// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HistoricalJobQuery Query for selecting logs analyzed by the historical job.
type HistoricalJobQuery struct {
	// Additional filters appended to the query at evaluation time.
	AdditionalFilters *string `json:"additionalFilters,omitempty"`
	// The aggregation type.
	Aggregation *SecurityMonitoringRuleQueryAggregation `json:"aggregation,omitempty"`
	// Fields used to correlate results across queries in sequence detection rules.
	CorrelatedByFields []string `json:"correlatedByFields,omitempty"`
	// Zero-based index of the query to correlate with in sequence detection rules. Up to 10 queries are supported, so valid values are 0 to 9.
	CorrelatedQueryIndex *int64 `json:"correlatedQueryIndex,omitempty"`
	// Custom query extension used to refine the base query.
	CustomQueryExtension *string `json:"customQueryExtension,omitempty"`
	// Source of events, either logs, audit trail, security signals, or Datadog events. `app_sec_spans` is deprecated in favor of `spans`.
	DataSource *SecurityMonitoringStandardDataSource `json:"dataSource,omitempty"`
	// IDs of reference datasets used by this query.
	DatasetIds []string `json:"datasetIds,omitempty"`
	// Field for which the cardinality is measured. Sent as an array.
	DistinctFields []string `json:"distinctFields,omitempty"`
	// Fields to group by.
	GroupByFields []string `json:"groupByFields,omitempty"`
	// When false, events without a group-by value are ignored by the query. When true, events with missing group-by fields are processed with `N/A`, replacing the missing values.
	HasOptionalGroupByFields *bool `json:"hasOptionalGroupByFields,omitempty"`
	// Index used to load the data for this query.
	Index *string `json:"index,omitempty"`
	// Indexes used to load the data for this query. Mutually exclusive with `index`.
	Indexes []string `json:"indexes,omitempty"`
	// Group of target fields to aggregate over when using the sum, max, geo data, or new value aggregations. The sum, max, and geo data aggregations only accept one value in this list, whereas the new value aggregation accepts up to five values.
	Metrics []string `json:"metrics,omitempty"`
	// Name of the query.
	Name *string `json:"name,omitempty"`
	// Query to run on logs.
	Query *string `json:"query,omitempty"`
	// Language used to parse the query string.
	QueryLanguage *string `json:"queryLanguage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewHistoricalJobQuery instantiates a new HistoricalJobQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHistoricalJobQuery() *HistoricalJobQuery {
	this := HistoricalJobQuery{}
	var dataSource SecurityMonitoringStandardDataSource = SECURITYMONITORINGSTANDARDDATASOURCE_LOGS
	this.DataSource = &dataSource
	var hasOptionalGroupByFields bool = false
	this.HasOptionalGroupByFields = &hasOptionalGroupByFields
	return &this
}

// NewHistoricalJobQueryWithDefaults instantiates a new HistoricalJobQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHistoricalJobQueryWithDefaults() *HistoricalJobQuery {
	this := HistoricalJobQuery{}
	var dataSource SecurityMonitoringStandardDataSource = SECURITYMONITORINGSTANDARDDATASOURCE_LOGS
	this.DataSource = &dataSource
	var hasOptionalGroupByFields bool = false
	this.HasOptionalGroupByFields = &hasOptionalGroupByFields
	return &this
}

// GetAdditionalFilters returns the AdditionalFilters field value if set, zero value otherwise.
func (o *HistoricalJobQuery) GetAdditionalFilters() string {
	if o == nil || o.AdditionalFilters == nil {
		var ret string
		return ret
	}
	return *o.AdditionalFilters
}

// GetAdditionalFiltersOk returns a tuple with the AdditionalFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalJobQuery) GetAdditionalFiltersOk() (*string, bool) {
	if o == nil || o.AdditionalFilters == nil {
		return nil, false
	}
	return o.AdditionalFilters, true
}

// HasAdditionalFilters returns a boolean if a field has been set.
func (o *HistoricalJobQuery) HasAdditionalFilters() bool {
	return o != nil && o.AdditionalFilters != nil
}

// SetAdditionalFilters gets a reference to the given string and assigns it to the AdditionalFilters field.
func (o *HistoricalJobQuery) SetAdditionalFilters(v string) {
	o.AdditionalFilters = &v
}

// GetAggregation returns the Aggregation field value if set, zero value otherwise.
func (o *HistoricalJobQuery) GetAggregation() SecurityMonitoringRuleQueryAggregation {
	if o == nil || o.Aggregation == nil {
		var ret SecurityMonitoringRuleQueryAggregation
		return ret
	}
	return *o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalJobQuery) GetAggregationOk() (*SecurityMonitoringRuleQueryAggregation, bool) {
	if o == nil || o.Aggregation == nil {
		return nil, false
	}
	return o.Aggregation, true
}

// HasAggregation returns a boolean if a field has been set.
func (o *HistoricalJobQuery) HasAggregation() bool {
	return o != nil && o.Aggregation != nil
}

// SetAggregation gets a reference to the given SecurityMonitoringRuleQueryAggregation and assigns it to the Aggregation field.
func (o *HistoricalJobQuery) SetAggregation(v SecurityMonitoringRuleQueryAggregation) {
	o.Aggregation = &v
}

// GetCorrelatedByFields returns the CorrelatedByFields field value if set, zero value otherwise.
func (o *HistoricalJobQuery) GetCorrelatedByFields() []string {
	if o == nil || o.CorrelatedByFields == nil {
		var ret []string
		return ret
	}
	return o.CorrelatedByFields
}

// GetCorrelatedByFieldsOk returns a tuple with the CorrelatedByFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalJobQuery) GetCorrelatedByFieldsOk() (*[]string, bool) {
	if o == nil || o.CorrelatedByFields == nil {
		return nil, false
	}
	return &o.CorrelatedByFields, true
}

// HasCorrelatedByFields returns a boolean if a field has been set.
func (o *HistoricalJobQuery) HasCorrelatedByFields() bool {
	return o != nil && o.CorrelatedByFields != nil
}

// SetCorrelatedByFields gets a reference to the given []string and assigns it to the CorrelatedByFields field.
func (o *HistoricalJobQuery) SetCorrelatedByFields(v []string) {
	o.CorrelatedByFields = v
}

// GetCorrelatedQueryIndex returns the CorrelatedQueryIndex field value if set, zero value otherwise.
func (o *HistoricalJobQuery) GetCorrelatedQueryIndex() int64 {
	if o == nil || o.CorrelatedQueryIndex == nil {
		var ret int64
		return ret
	}
	return *o.CorrelatedQueryIndex
}

// GetCorrelatedQueryIndexOk returns a tuple with the CorrelatedQueryIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalJobQuery) GetCorrelatedQueryIndexOk() (*int64, bool) {
	if o == nil || o.CorrelatedQueryIndex == nil {
		return nil, false
	}
	return o.CorrelatedQueryIndex, true
}

// HasCorrelatedQueryIndex returns a boolean if a field has been set.
func (o *HistoricalJobQuery) HasCorrelatedQueryIndex() bool {
	return o != nil && o.CorrelatedQueryIndex != nil
}

// SetCorrelatedQueryIndex gets a reference to the given int64 and assigns it to the CorrelatedQueryIndex field.
func (o *HistoricalJobQuery) SetCorrelatedQueryIndex(v int64) {
	o.CorrelatedQueryIndex = &v
}

// GetCustomQueryExtension returns the CustomQueryExtension field value if set, zero value otherwise.
func (o *HistoricalJobQuery) GetCustomQueryExtension() string {
	if o == nil || o.CustomQueryExtension == nil {
		var ret string
		return ret
	}
	return *o.CustomQueryExtension
}

// GetCustomQueryExtensionOk returns a tuple with the CustomQueryExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalJobQuery) GetCustomQueryExtensionOk() (*string, bool) {
	if o == nil || o.CustomQueryExtension == nil {
		return nil, false
	}
	return o.CustomQueryExtension, true
}

// HasCustomQueryExtension returns a boolean if a field has been set.
func (o *HistoricalJobQuery) HasCustomQueryExtension() bool {
	return o != nil && o.CustomQueryExtension != nil
}

// SetCustomQueryExtension gets a reference to the given string and assigns it to the CustomQueryExtension field.
func (o *HistoricalJobQuery) SetCustomQueryExtension(v string) {
	o.CustomQueryExtension = &v
}

// GetDataSource returns the DataSource field value if set, zero value otherwise.
func (o *HistoricalJobQuery) GetDataSource() SecurityMonitoringStandardDataSource {
	if o == nil || o.DataSource == nil {
		var ret SecurityMonitoringStandardDataSource
		return ret
	}
	return *o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalJobQuery) GetDataSourceOk() (*SecurityMonitoringStandardDataSource, bool) {
	if o == nil || o.DataSource == nil {
		return nil, false
	}
	return o.DataSource, true
}

// HasDataSource returns a boolean if a field has been set.
func (o *HistoricalJobQuery) HasDataSource() bool {
	return o != nil && o.DataSource != nil
}

// SetDataSource gets a reference to the given SecurityMonitoringStandardDataSource and assigns it to the DataSource field.
func (o *HistoricalJobQuery) SetDataSource(v SecurityMonitoringStandardDataSource) {
	o.DataSource = &v
}

// GetDatasetIds returns the DatasetIds field value if set, zero value otherwise.
func (o *HistoricalJobQuery) GetDatasetIds() []string {
	if o == nil || o.DatasetIds == nil {
		var ret []string
		return ret
	}
	return o.DatasetIds
}

// GetDatasetIdsOk returns a tuple with the DatasetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalJobQuery) GetDatasetIdsOk() (*[]string, bool) {
	if o == nil || o.DatasetIds == nil {
		return nil, false
	}
	return &o.DatasetIds, true
}

// HasDatasetIds returns a boolean if a field has been set.
func (o *HistoricalJobQuery) HasDatasetIds() bool {
	return o != nil && o.DatasetIds != nil
}

// SetDatasetIds gets a reference to the given []string and assigns it to the DatasetIds field.
func (o *HistoricalJobQuery) SetDatasetIds(v []string) {
	o.DatasetIds = v
}

// GetDistinctFields returns the DistinctFields field value if set, zero value otherwise.
func (o *HistoricalJobQuery) GetDistinctFields() []string {
	if o == nil || o.DistinctFields == nil {
		var ret []string
		return ret
	}
	return o.DistinctFields
}

// GetDistinctFieldsOk returns a tuple with the DistinctFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalJobQuery) GetDistinctFieldsOk() (*[]string, bool) {
	if o == nil || o.DistinctFields == nil {
		return nil, false
	}
	return &o.DistinctFields, true
}

// HasDistinctFields returns a boolean if a field has been set.
func (o *HistoricalJobQuery) HasDistinctFields() bool {
	return o != nil && o.DistinctFields != nil
}

// SetDistinctFields gets a reference to the given []string and assigns it to the DistinctFields field.
func (o *HistoricalJobQuery) SetDistinctFields(v []string) {
	o.DistinctFields = v
}

// GetGroupByFields returns the GroupByFields field value if set, zero value otherwise.
func (o *HistoricalJobQuery) GetGroupByFields() []string {
	if o == nil || o.GroupByFields == nil {
		var ret []string
		return ret
	}
	return o.GroupByFields
}

// GetGroupByFieldsOk returns a tuple with the GroupByFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalJobQuery) GetGroupByFieldsOk() (*[]string, bool) {
	if o == nil || o.GroupByFields == nil {
		return nil, false
	}
	return &o.GroupByFields, true
}

// HasGroupByFields returns a boolean if a field has been set.
func (o *HistoricalJobQuery) HasGroupByFields() bool {
	return o != nil && o.GroupByFields != nil
}

// SetGroupByFields gets a reference to the given []string and assigns it to the GroupByFields field.
func (o *HistoricalJobQuery) SetGroupByFields(v []string) {
	o.GroupByFields = v
}

// GetHasOptionalGroupByFields returns the HasOptionalGroupByFields field value if set, zero value otherwise.
func (o *HistoricalJobQuery) GetHasOptionalGroupByFields() bool {
	if o == nil || o.HasOptionalGroupByFields == nil {
		var ret bool
		return ret
	}
	return *o.HasOptionalGroupByFields
}

// GetHasOptionalGroupByFieldsOk returns a tuple with the HasOptionalGroupByFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalJobQuery) GetHasOptionalGroupByFieldsOk() (*bool, bool) {
	if o == nil || o.HasOptionalGroupByFields == nil {
		return nil, false
	}
	return o.HasOptionalGroupByFields, true
}

// HasHasOptionalGroupByFields returns a boolean if a field has been set.
func (o *HistoricalJobQuery) HasHasOptionalGroupByFields() bool {
	return o != nil && o.HasOptionalGroupByFields != nil
}

// SetHasOptionalGroupByFields gets a reference to the given bool and assigns it to the HasOptionalGroupByFields field.
func (o *HistoricalJobQuery) SetHasOptionalGroupByFields(v bool) {
	o.HasOptionalGroupByFields = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *HistoricalJobQuery) GetIndex() string {
	if o == nil || o.Index == nil {
		var ret string
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalJobQuery) GetIndexOk() (*string, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *HistoricalJobQuery) HasIndex() bool {
	return o != nil && o.Index != nil
}

// SetIndex gets a reference to the given string and assigns it to the Index field.
func (o *HistoricalJobQuery) SetIndex(v string) {
	o.Index = &v
}

// GetIndexes returns the Indexes field value if set, zero value otherwise.
func (o *HistoricalJobQuery) GetIndexes() []string {
	if o == nil || o.Indexes == nil {
		var ret []string
		return ret
	}
	return o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalJobQuery) GetIndexesOk() (*[]string, bool) {
	if o == nil || o.Indexes == nil {
		return nil, false
	}
	return &o.Indexes, true
}

// HasIndexes returns a boolean if a field has been set.
func (o *HistoricalJobQuery) HasIndexes() bool {
	return o != nil && o.Indexes != nil
}

// SetIndexes gets a reference to the given []string and assigns it to the Indexes field.
func (o *HistoricalJobQuery) SetIndexes(v []string) {
	o.Indexes = v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *HistoricalJobQuery) GetMetrics() []string {
	if o == nil || o.Metrics == nil {
		var ret []string
		return ret
	}
	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalJobQuery) GetMetricsOk() (*[]string, bool) {
	if o == nil || o.Metrics == nil {
		return nil, false
	}
	return &o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *HistoricalJobQuery) HasMetrics() bool {
	return o != nil && o.Metrics != nil
}

// SetMetrics gets a reference to the given []string and assigns it to the Metrics field.
func (o *HistoricalJobQuery) SetMetrics(v []string) {
	o.Metrics = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HistoricalJobQuery) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalJobQuery) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HistoricalJobQuery) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HistoricalJobQuery) SetName(v string) {
	o.Name = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *HistoricalJobQuery) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalJobQuery) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *HistoricalJobQuery) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *HistoricalJobQuery) SetQuery(v string) {
	o.Query = &v
}

// GetQueryLanguage returns the QueryLanguage field value if set, zero value otherwise.
func (o *HistoricalJobQuery) GetQueryLanguage() string {
	if o == nil || o.QueryLanguage == nil {
		var ret string
		return ret
	}
	return *o.QueryLanguage
}

// GetQueryLanguageOk returns a tuple with the QueryLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalJobQuery) GetQueryLanguageOk() (*string, bool) {
	if o == nil || o.QueryLanguage == nil {
		return nil, false
	}
	return o.QueryLanguage, true
}

// HasQueryLanguage returns a boolean if a field has been set.
func (o *HistoricalJobQuery) HasQueryLanguage() bool {
	return o != nil && o.QueryLanguage != nil
}

// SetQueryLanguage gets a reference to the given string and assigns it to the QueryLanguage field.
func (o *HistoricalJobQuery) SetQueryLanguage(v string) {
	o.QueryLanguage = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o HistoricalJobQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AdditionalFilters != nil {
		toSerialize["additionalFilters"] = o.AdditionalFilters
	}
	if o.Aggregation != nil {
		toSerialize["aggregation"] = o.Aggregation
	}
	if o.CorrelatedByFields != nil {
		toSerialize["correlatedByFields"] = o.CorrelatedByFields
	}
	if o.CorrelatedQueryIndex != nil {
		toSerialize["correlatedQueryIndex"] = o.CorrelatedQueryIndex
	}
	if o.CustomQueryExtension != nil {
		toSerialize["customQueryExtension"] = o.CustomQueryExtension
	}
	if o.DataSource != nil {
		toSerialize["dataSource"] = o.DataSource
	}
	if o.DatasetIds != nil {
		toSerialize["datasetIds"] = o.DatasetIds
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
	if o.Metrics != nil {
		toSerialize["metrics"] = o.Metrics
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.QueryLanguage != nil {
		toSerialize["queryLanguage"] = o.QueryLanguage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HistoricalJobQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AdditionalFilters        *string                                 `json:"additionalFilters,omitempty"`
		Aggregation              *SecurityMonitoringRuleQueryAggregation `json:"aggregation,omitempty"`
		CorrelatedByFields       []string                                `json:"correlatedByFields,omitempty"`
		CorrelatedQueryIndex     *int64                                  `json:"correlatedQueryIndex,omitempty"`
		CustomQueryExtension     *string                                 `json:"customQueryExtension,omitempty"`
		DataSource               *SecurityMonitoringStandardDataSource   `json:"dataSource,omitempty"`
		DatasetIds               []string                                `json:"datasetIds,omitempty"`
		DistinctFields           []string                                `json:"distinctFields,omitempty"`
		GroupByFields            []string                                `json:"groupByFields,omitempty"`
		HasOptionalGroupByFields *bool                                   `json:"hasOptionalGroupByFields,omitempty"`
		Index                    *string                                 `json:"index,omitempty"`
		Indexes                  []string                                `json:"indexes,omitempty"`
		Metrics                  []string                                `json:"metrics,omitempty"`
		Name                     *string                                 `json:"name,omitempty"`
		Query                    *string                                 `json:"query,omitempty"`
		QueryLanguage            *string                                 `json:"queryLanguage,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"additionalFilters", "aggregation", "correlatedByFields", "correlatedQueryIndex", "customQueryExtension", "dataSource", "datasetIds", "distinctFields", "groupByFields", "hasOptionalGroupByFields", "index", "indexes", "metrics", "name", "query", "queryLanguage"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AdditionalFilters = all.AdditionalFilters
	if all.Aggregation != nil && !all.Aggregation.IsValid() {
		hasInvalidField = true
	} else {
		o.Aggregation = all.Aggregation
	}
	o.CorrelatedByFields = all.CorrelatedByFields
	o.CorrelatedQueryIndex = all.CorrelatedQueryIndex
	o.CustomQueryExtension = all.CustomQueryExtension
	if all.DataSource != nil && !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = all.DataSource
	}
	o.DatasetIds = all.DatasetIds
	o.DistinctFields = all.DistinctFields
	o.GroupByFields = all.GroupByFields
	o.HasOptionalGroupByFields = all.HasOptionalGroupByFields
	o.Index = all.Index
	o.Indexes = all.Indexes
	o.Metrics = all.Metrics
	o.Name = all.Name
	o.Query = all.Query
	o.QueryLanguage = all.QueryLanguage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
