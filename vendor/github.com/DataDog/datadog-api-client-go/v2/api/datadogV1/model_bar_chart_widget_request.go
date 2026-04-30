// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BarChartWidgetRequest Updated bar chart widget.
type BarChartWidgetRequest struct {
	// The log query.
	ApmQuery *LogQueryDefinition `json:"apm_query,omitempty"`
	// The log query.
	AuditQuery *LogQueryDefinition `json:"audit_query,omitempty"`
	// List of conditional formats.
	ConditionalFormats []WidgetConditionalFormat `json:"conditional_formats,omitempty"`
	// The log query.
	EventQuery *LogQueryDefinition `json:"event_query,omitempty"`
	// List of formulas that operate on queries.
	Formulas []WidgetFormula `json:"formulas,omitempty"`
	// The log query.
	LogQuery *LogQueryDefinition `json:"log_query,omitempty"`
	// The log query.
	NetworkQuery *LogQueryDefinition `json:"network_query,omitempty"`
	// The process query to use in the widget.
	ProcessQuery *ProcessQueryDefinition `json:"process_query,omitempty"`
	// The log query.
	ProfileMetricsQuery *LogQueryDefinition `json:"profile_metrics_query,omitempty"`
	// Widget query. Deprecated - Use `queries` and `formulas` instead.
	// Deprecated
	Q *string `json:"q,omitempty"`
	// List of queries that can be returned directly or used in formulas.
	Queries []FormulaAndFunctionQueryDefinition `json:"queries,omitempty"`
	// Timeseries, scalar, or event list response. Event list response formats are supported by Geomap widgets.
	ResponseFormat *FormulaAndFunctionResponseFormat `json:"response_format,omitempty"`
	// The log query.
	RumQuery *LogQueryDefinition `json:"rum_query,omitempty"`
	// The log query.
	SecurityQuery *LogQueryDefinition `json:"security_query,omitempty"`
	// The controls for sorting the widget.
	Sort *WidgetSortBy `json:"sort,omitempty"`
	// Define request widget style.
	Style *WidgetRequestStyle `json:"style,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBarChartWidgetRequest instantiates a new BarChartWidgetRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBarChartWidgetRequest() *BarChartWidgetRequest {
	this := BarChartWidgetRequest{}
	return &this
}

// NewBarChartWidgetRequestWithDefaults instantiates a new BarChartWidgetRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBarChartWidgetRequestWithDefaults() *BarChartWidgetRequest {
	this := BarChartWidgetRequest{}
	return &this
}

// GetApmQuery returns the ApmQuery field value if set, zero value otherwise.
func (o *BarChartWidgetRequest) GetApmQuery() LogQueryDefinition {
	if o == nil || o.ApmQuery == nil {
		var ret LogQueryDefinition
		return ret
	}
	return *o.ApmQuery
}

// GetApmQueryOk returns a tuple with the ApmQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BarChartWidgetRequest) GetApmQueryOk() (*LogQueryDefinition, bool) {
	if o == nil || o.ApmQuery == nil {
		return nil, false
	}
	return o.ApmQuery, true
}

// HasApmQuery returns a boolean if a field has been set.
func (o *BarChartWidgetRequest) HasApmQuery() bool {
	return o != nil && o.ApmQuery != nil
}

// SetApmQuery gets a reference to the given LogQueryDefinition and assigns it to the ApmQuery field.
func (o *BarChartWidgetRequest) SetApmQuery(v LogQueryDefinition) {
	o.ApmQuery = &v
}

// GetAuditQuery returns the AuditQuery field value if set, zero value otherwise.
func (o *BarChartWidgetRequest) GetAuditQuery() LogQueryDefinition {
	if o == nil || o.AuditQuery == nil {
		var ret LogQueryDefinition
		return ret
	}
	return *o.AuditQuery
}

// GetAuditQueryOk returns a tuple with the AuditQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BarChartWidgetRequest) GetAuditQueryOk() (*LogQueryDefinition, bool) {
	if o == nil || o.AuditQuery == nil {
		return nil, false
	}
	return o.AuditQuery, true
}

// HasAuditQuery returns a boolean if a field has been set.
func (o *BarChartWidgetRequest) HasAuditQuery() bool {
	return o != nil && o.AuditQuery != nil
}

// SetAuditQuery gets a reference to the given LogQueryDefinition and assigns it to the AuditQuery field.
func (o *BarChartWidgetRequest) SetAuditQuery(v LogQueryDefinition) {
	o.AuditQuery = &v
}

// GetConditionalFormats returns the ConditionalFormats field value if set, zero value otherwise.
func (o *BarChartWidgetRequest) GetConditionalFormats() []WidgetConditionalFormat {
	if o == nil || o.ConditionalFormats == nil {
		var ret []WidgetConditionalFormat
		return ret
	}
	return o.ConditionalFormats
}

// GetConditionalFormatsOk returns a tuple with the ConditionalFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BarChartWidgetRequest) GetConditionalFormatsOk() (*[]WidgetConditionalFormat, bool) {
	if o == nil || o.ConditionalFormats == nil {
		return nil, false
	}
	return &o.ConditionalFormats, true
}

// HasConditionalFormats returns a boolean if a field has been set.
func (o *BarChartWidgetRequest) HasConditionalFormats() bool {
	return o != nil && o.ConditionalFormats != nil
}

// SetConditionalFormats gets a reference to the given []WidgetConditionalFormat and assigns it to the ConditionalFormats field.
func (o *BarChartWidgetRequest) SetConditionalFormats(v []WidgetConditionalFormat) {
	o.ConditionalFormats = v
}

// GetEventQuery returns the EventQuery field value if set, zero value otherwise.
func (o *BarChartWidgetRequest) GetEventQuery() LogQueryDefinition {
	if o == nil || o.EventQuery == nil {
		var ret LogQueryDefinition
		return ret
	}
	return *o.EventQuery
}

// GetEventQueryOk returns a tuple with the EventQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BarChartWidgetRequest) GetEventQueryOk() (*LogQueryDefinition, bool) {
	if o == nil || o.EventQuery == nil {
		return nil, false
	}
	return o.EventQuery, true
}

// HasEventQuery returns a boolean if a field has been set.
func (o *BarChartWidgetRequest) HasEventQuery() bool {
	return o != nil && o.EventQuery != nil
}

// SetEventQuery gets a reference to the given LogQueryDefinition and assigns it to the EventQuery field.
func (o *BarChartWidgetRequest) SetEventQuery(v LogQueryDefinition) {
	o.EventQuery = &v
}

// GetFormulas returns the Formulas field value if set, zero value otherwise.
func (o *BarChartWidgetRequest) GetFormulas() []WidgetFormula {
	if o == nil || o.Formulas == nil {
		var ret []WidgetFormula
		return ret
	}
	return o.Formulas
}

// GetFormulasOk returns a tuple with the Formulas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BarChartWidgetRequest) GetFormulasOk() (*[]WidgetFormula, bool) {
	if o == nil || o.Formulas == nil {
		return nil, false
	}
	return &o.Formulas, true
}

// HasFormulas returns a boolean if a field has been set.
func (o *BarChartWidgetRequest) HasFormulas() bool {
	return o != nil && o.Formulas != nil
}

// SetFormulas gets a reference to the given []WidgetFormula and assigns it to the Formulas field.
func (o *BarChartWidgetRequest) SetFormulas(v []WidgetFormula) {
	o.Formulas = v
}

// GetLogQuery returns the LogQuery field value if set, zero value otherwise.
func (o *BarChartWidgetRequest) GetLogQuery() LogQueryDefinition {
	if o == nil || o.LogQuery == nil {
		var ret LogQueryDefinition
		return ret
	}
	return *o.LogQuery
}

// GetLogQueryOk returns a tuple with the LogQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BarChartWidgetRequest) GetLogQueryOk() (*LogQueryDefinition, bool) {
	if o == nil || o.LogQuery == nil {
		return nil, false
	}
	return o.LogQuery, true
}

// HasLogQuery returns a boolean if a field has been set.
func (o *BarChartWidgetRequest) HasLogQuery() bool {
	return o != nil && o.LogQuery != nil
}

// SetLogQuery gets a reference to the given LogQueryDefinition and assigns it to the LogQuery field.
func (o *BarChartWidgetRequest) SetLogQuery(v LogQueryDefinition) {
	o.LogQuery = &v
}

// GetNetworkQuery returns the NetworkQuery field value if set, zero value otherwise.
func (o *BarChartWidgetRequest) GetNetworkQuery() LogQueryDefinition {
	if o == nil || o.NetworkQuery == nil {
		var ret LogQueryDefinition
		return ret
	}
	return *o.NetworkQuery
}

// GetNetworkQueryOk returns a tuple with the NetworkQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BarChartWidgetRequest) GetNetworkQueryOk() (*LogQueryDefinition, bool) {
	if o == nil || o.NetworkQuery == nil {
		return nil, false
	}
	return o.NetworkQuery, true
}

// HasNetworkQuery returns a boolean if a field has been set.
func (o *BarChartWidgetRequest) HasNetworkQuery() bool {
	return o != nil && o.NetworkQuery != nil
}

// SetNetworkQuery gets a reference to the given LogQueryDefinition and assigns it to the NetworkQuery field.
func (o *BarChartWidgetRequest) SetNetworkQuery(v LogQueryDefinition) {
	o.NetworkQuery = &v
}

// GetProcessQuery returns the ProcessQuery field value if set, zero value otherwise.
func (o *BarChartWidgetRequest) GetProcessQuery() ProcessQueryDefinition {
	if o == nil || o.ProcessQuery == nil {
		var ret ProcessQueryDefinition
		return ret
	}
	return *o.ProcessQuery
}

// GetProcessQueryOk returns a tuple with the ProcessQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BarChartWidgetRequest) GetProcessQueryOk() (*ProcessQueryDefinition, bool) {
	if o == nil || o.ProcessQuery == nil {
		return nil, false
	}
	return o.ProcessQuery, true
}

// HasProcessQuery returns a boolean if a field has been set.
func (o *BarChartWidgetRequest) HasProcessQuery() bool {
	return o != nil && o.ProcessQuery != nil
}

// SetProcessQuery gets a reference to the given ProcessQueryDefinition and assigns it to the ProcessQuery field.
func (o *BarChartWidgetRequest) SetProcessQuery(v ProcessQueryDefinition) {
	o.ProcessQuery = &v
}

// GetProfileMetricsQuery returns the ProfileMetricsQuery field value if set, zero value otherwise.
func (o *BarChartWidgetRequest) GetProfileMetricsQuery() LogQueryDefinition {
	if o == nil || o.ProfileMetricsQuery == nil {
		var ret LogQueryDefinition
		return ret
	}
	return *o.ProfileMetricsQuery
}

// GetProfileMetricsQueryOk returns a tuple with the ProfileMetricsQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BarChartWidgetRequest) GetProfileMetricsQueryOk() (*LogQueryDefinition, bool) {
	if o == nil || o.ProfileMetricsQuery == nil {
		return nil, false
	}
	return o.ProfileMetricsQuery, true
}

// HasProfileMetricsQuery returns a boolean if a field has been set.
func (o *BarChartWidgetRequest) HasProfileMetricsQuery() bool {
	return o != nil && o.ProfileMetricsQuery != nil
}

// SetProfileMetricsQuery gets a reference to the given LogQueryDefinition and assigns it to the ProfileMetricsQuery field.
func (o *BarChartWidgetRequest) SetProfileMetricsQuery(v LogQueryDefinition) {
	o.ProfileMetricsQuery = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
// Deprecated
func (o *BarChartWidgetRequest) GetQ() string {
	if o == nil || o.Q == nil {
		var ret string
		return ret
	}
	return *o.Q
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *BarChartWidgetRequest) GetQOk() (*string, bool) {
	if o == nil || o.Q == nil {
		return nil, false
	}
	return o.Q, true
}

// HasQ returns a boolean if a field has been set.
func (o *BarChartWidgetRequest) HasQ() bool {
	return o != nil && o.Q != nil
}

// SetQ gets a reference to the given string and assigns it to the Q field.
// Deprecated
func (o *BarChartWidgetRequest) SetQ(v string) {
	o.Q = &v
}

// GetQueries returns the Queries field value if set, zero value otherwise.
func (o *BarChartWidgetRequest) GetQueries() []FormulaAndFunctionQueryDefinition {
	if o == nil || o.Queries == nil {
		var ret []FormulaAndFunctionQueryDefinition
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BarChartWidgetRequest) GetQueriesOk() (*[]FormulaAndFunctionQueryDefinition, bool) {
	if o == nil || o.Queries == nil {
		return nil, false
	}
	return &o.Queries, true
}

// HasQueries returns a boolean if a field has been set.
func (o *BarChartWidgetRequest) HasQueries() bool {
	return o != nil && o.Queries != nil
}

// SetQueries gets a reference to the given []FormulaAndFunctionQueryDefinition and assigns it to the Queries field.
func (o *BarChartWidgetRequest) SetQueries(v []FormulaAndFunctionQueryDefinition) {
	o.Queries = v
}

// GetResponseFormat returns the ResponseFormat field value if set, zero value otherwise.
func (o *BarChartWidgetRequest) GetResponseFormat() FormulaAndFunctionResponseFormat {
	if o == nil || o.ResponseFormat == nil {
		var ret FormulaAndFunctionResponseFormat
		return ret
	}
	return *o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BarChartWidgetRequest) GetResponseFormatOk() (*FormulaAndFunctionResponseFormat, bool) {
	if o == nil || o.ResponseFormat == nil {
		return nil, false
	}
	return o.ResponseFormat, true
}

// HasResponseFormat returns a boolean if a field has been set.
func (o *BarChartWidgetRequest) HasResponseFormat() bool {
	return o != nil && o.ResponseFormat != nil
}

// SetResponseFormat gets a reference to the given FormulaAndFunctionResponseFormat and assigns it to the ResponseFormat field.
func (o *BarChartWidgetRequest) SetResponseFormat(v FormulaAndFunctionResponseFormat) {
	o.ResponseFormat = &v
}

// GetRumQuery returns the RumQuery field value if set, zero value otherwise.
func (o *BarChartWidgetRequest) GetRumQuery() LogQueryDefinition {
	if o == nil || o.RumQuery == nil {
		var ret LogQueryDefinition
		return ret
	}
	return *o.RumQuery
}

// GetRumQueryOk returns a tuple with the RumQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BarChartWidgetRequest) GetRumQueryOk() (*LogQueryDefinition, bool) {
	if o == nil || o.RumQuery == nil {
		return nil, false
	}
	return o.RumQuery, true
}

// HasRumQuery returns a boolean if a field has been set.
func (o *BarChartWidgetRequest) HasRumQuery() bool {
	return o != nil && o.RumQuery != nil
}

// SetRumQuery gets a reference to the given LogQueryDefinition and assigns it to the RumQuery field.
func (o *BarChartWidgetRequest) SetRumQuery(v LogQueryDefinition) {
	o.RumQuery = &v
}

// GetSecurityQuery returns the SecurityQuery field value if set, zero value otherwise.
func (o *BarChartWidgetRequest) GetSecurityQuery() LogQueryDefinition {
	if o == nil || o.SecurityQuery == nil {
		var ret LogQueryDefinition
		return ret
	}
	return *o.SecurityQuery
}

// GetSecurityQueryOk returns a tuple with the SecurityQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BarChartWidgetRequest) GetSecurityQueryOk() (*LogQueryDefinition, bool) {
	if o == nil || o.SecurityQuery == nil {
		return nil, false
	}
	return o.SecurityQuery, true
}

// HasSecurityQuery returns a boolean if a field has been set.
func (o *BarChartWidgetRequest) HasSecurityQuery() bool {
	return o != nil && o.SecurityQuery != nil
}

// SetSecurityQuery gets a reference to the given LogQueryDefinition and assigns it to the SecurityQuery field.
func (o *BarChartWidgetRequest) SetSecurityQuery(v LogQueryDefinition) {
	o.SecurityQuery = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *BarChartWidgetRequest) GetSort() WidgetSortBy {
	if o == nil || o.Sort == nil {
		var ret WidgetSortBy
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BarChartWidgetRequest) GetSortOk() (*WidgetSortBy, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *BarChartWidgetRequest) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given WidgetSortBy and assigns it to the Sort field.
func (o *BarChartWidgetRequest) SetSort(v WidgetSortBy) {
	o.Sort = &v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *BarChartWidgetRequest) GetStyle() WidgetRequestStyle {
	if o == nil || o.Style == nil {
		var ret WidgetRequestStyle
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BarChartWidgetRequest) GetStyleOk() (*WidgetRequestStyle, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *BarChartWidgetRequest) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given WidgetRequestStyle and assigns it to the Style field.
func (o *BarChartWidgetRequest) SetStyle(v WidgetRequestStyle) {
	o.Style = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o BarChartWidgetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ApmQuery != nil {
		toSerialize["apm_query"] = o.ApmQuery
	}
	if o.AuditQuery != nil {
		toSerialize["audit_query"] = o.AuditQuery
	}
	if o.ConditionalFormats != nil {
		toSerialize["conditional_formats"] = o.ConditionalFormats
	}
	if o.EventQuery != nil {
		toSerialize["event_query"] = o.EventQuery
	}
	if o.Formulas != nil {
		toSerialize["formulas"] = o.Formulas
	}
	if o.LogQuery != nil {
		toSerialize["log_query"] = o.LogQuery
	}
	if o.NetworkQuery != nil {
		toSerialize["network_query"] = o.NetworkQuery
	}
	if o.ProcessQuery != nil {
		toSerialize["process_query"] = o.ProcessQuery
	}
	if o.ProfileMetricsQuery != nil {
		toSerialize["profile_metrics_query"] = o.ProfileMetricsQuery
	}
	if o.Q != nil {
		toSerialize["q"] = o.Q
	}
	if o.Queries != nil {
		toSerialize["queries"] = o.Queries
	}
	if o.ResponseFormat != nil {
		toSerialize["response_format"] = o.ResponseFormat
	}
	if o.RumQuery != nil {
		toSerialize["rum_query"] = o.RumQuery
	}
	if o.SecurityQuery != nil {
		toSerialize["security_query"] = o.SecurityQuery
	}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	if o.Style != nil {
		toSerialize["style"] = o.Style
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BarChartWidgetRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApmQuery            *LogQueryDefinition                 `json:"apm_query,omitempty"`
		AuditQuery          *LogQueryDefinition                 `json:"audit_query,omitempty"`
		ConditionalFormats  []WidgetConditionalFormat           `json:"conditional_formats,omitempty"`
		EventQuery          *LogQueryDefinition                 `json:"event_query,omitempty"`
		Formulas            []WidgetFormula                     `json:"formulas,omitempty"`
		LogQuery            *LogQueryDefinition                 `json:"log_query,omitempty"`
		NetworkQuery        *LogQueryDefinition                 `json:"network_query,omitempty"`
		ProcessQuery        *ProcessQueryDefinition             `json:"process_query,omitempty"`
		ProfileMetricsQuery *LogQueryDefinition                 `json:"profile_metrics_query,omitempty"`
		Q                   *string                             `json:"q,omitempty"`
		Queries             []FormulaAndFunctionQueryDefinition `json:"queries,omitempty"`
		ResponseFormat      *FormulaAndFunctionResponseFormat   `json:"response_format,omitempty"`
		RumQuery            *LogQueryDefinition                 `json:"rum_query,omitempty"`
		SecurityQuery       *LogQueryDefinition                 `json:"security_query,omitempty"`
		Sort                *WidgetSortBy                       `json:"sort,omitempty"`
		Style               *WidgetRequestStyle                 `json:"style,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"apm_query", "audit_query", "conditional_formats", "event_query", "formulas", "log_query", "network_query", "process_query", "profile_metrics_query", "q", "queries", "response_format", "rum_query", "security_query", "sort", "style"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ApmQuery != nil && all.ApmQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ApmQuery = all.ApmQuery
	if all.AuditQuery != nil && all.AuditQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AuditQuery = all.AuditQuery
	o.ConditionalFormats = all.ConditionalFormats
	if all.EventQuery != nil && all.EventQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.EventQuery = all.EventQuery
	o.Formulas = all.Formulas
	if all.LogQuery != nil && all.LogQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LogQuery = all.LogQuery
	if all.NetworkQuery != nil && all.NetworkQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.NetworkQuery = all.NetworkQuery
	if all.ProcessQuery != nil && all.ProcessQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ProcessQuery = all.ProcessQuery
	if all.ProfileMetricsQuery != nil && all.ProfileMetricsQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ProfileMetricsQuery = all.ProfileMetricsQuery
	o.Q = all.Q
	o.Queries = all.Queries
	if all.ResponseFormat != nil && !all.ResponseFormat.IsValid() {
		hasInvalidField = true
	} else {
		o.ResponseFormat = all.ResponseFormat
	}
	if all.RumQuery != nil && all.RumQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.RumQuery = all.RumQuery
	if all.SecurityQuery != nil && all.SecurityQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SecurityQuery = all.SecurityQuery
	if all.Sort != nil && all.Sort.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Sort = all.Sort
	if all.Style != nil && all.Style.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Style = all.Style

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
