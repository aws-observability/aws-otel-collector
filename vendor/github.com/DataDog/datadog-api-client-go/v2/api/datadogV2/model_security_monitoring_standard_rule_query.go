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
	// Field for which the cardinality is measured. Sent as an array.
	DistinctFields []string `json:"distinctFields,omitempty"`
	// Fields to group by.
	GroupByFields []string `json:"groupByFields,omitempty"`
	// When false, events without a group-by value are ignored by the rule. When true, events with missing group-by fields are processed with `N/A`, replacing the missing values.
	HasOptionalGroupByFields *bool `json:"hasOptionalGroupByFields,omitempty"`
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
	return &this
}

// NewSecurityMonitoringStandardRuleQueryWithDefaults instantiates a new SecurityMonitoringStandardRuleQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringStandardRuleQueryWithDefaults() *SecurityMonitoringStandardRuleQuery {
	this := SecurityMonitoringStandardRuleQuery{}
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
	if o.DistinctFields != nil {
		toSerialize["distinctFields"] = o.DistinctFields
	}
	if o.GroupByFields != nil {
		toSerialize["groupByFields"] = o.GroupByFields
	}
	if o.HasOptionalGroupByFields != nil {
		toSerialize["hasOptionalGroupByFields"] = o.HasOptionalGroupByFields
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
		DistinctFields           []string                                `json:"distinctFields,omitempty"`
		GroupByFields            []string                                `json:"groupByFields,omitempty"`
		HasOptionalGroupByFields *bool                                   `json:"hasOptionalGroupByFields,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"aggregation", "distinctFields", "groupByFields", "hasOptionalGroupByFields", "metric", "metrics", "name", "query"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Aggregation != nil && !all.Aggregation.IsValid() {
		hasInvalidField = true
	} else {
		o.Aggregation = all.Aggregation
	}
	o.DistinctFields = all.DistinctFields
	o.GroupByFields = all.GroupByFields
	o.HasOptionalGroupByFields = all.HasOptionalGroupByFields
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
