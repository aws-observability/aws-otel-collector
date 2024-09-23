// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeomapWidgetRequest An updated geomap widget.
type GeomapWidgetRequest struct {
	// Widget columns.
	Columns []ListStreamColumn `json:"columns,omitempty"`
	// List of formulas that operate on queries.
	Formulas []WidgetFormula `json:"formulas,omitempty"`
	// The log query.
	LogQuery *LogQueryDefinition `json:"log_query,omitempty"`
	// The widget metrics query.
	Q *string `json:"q,omitempty"`
	// List of queries that can be returned directly or used in formulas.
	Queries []FormulaAndFunctionQueryDefinition `json:"queries,omitempty"`
	// Updated list stream widget.
	Query *ListStreamQuery `json:"query,omitempty"`
	// Timeseries, scalar, or event list response. Event list response formats are supported by Geomap widgets.
	ResponseFormat *FormulaAndFunctionResponseFormat `json:"response_format,omitempty"`
	// The log query.
	RumQuery *LogQueryDefinition `json:"rum_query,omitempty"`
	// The log query.
	SecurityQuery *LogQueryDefinition `json:"security_query,omitempty"`
	// The controls for sorting the widget.
	Sort *WidgetSortBy `json:"sort,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGeomapWidgetRequest instantiates a new GeomapWidgetRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGeomapWidgetRequest() *GeomapWidgetRequest {
	this := GeomapWidgetRequest{}
	return &this
}

// NewGeomapWidgetRequestWithDefaults instantiates a new GeomapWidgetRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGeomapWidgetRequestWithDefaults() *GeomapWidgetRequest {
	this := GeomapWidgetRequest{}
	return &this
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *GeomapWidgetRequest) GetColumns() []ListStreamColumn {
	if o == nil || o.Columns == nil {
		var ret []ListStreamColumn
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeomapWidgetRequest) GetColumnsOk() (*[]ListStreamColumn, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return &o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *GeomapWidgetRequest) HasColumns() bool {
	return o != nil && o.Columns != nil
}

// SetColumns gets a reference to the given []ListStreamColumn and assigns it to the Columns field.
func (o *GeomapWidgetRequest) SetColumns(v []ListStreamColumn) {
	o.Columns = v
}

// GetFormulas returns the Formulas field value if set, zero value otherwise.
func (o *GeomapWidgetRequest) GetFormulas() []WidgetFormula {
	if o == nil || o.Formulas == nil {
		var ret []WidgetFormula
		return ret
	}
	return o.Formulas
}

// GetFormulasOk returns a tuple with the Formulas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeomapWidgetRequest) GetFormulasOk() (*[]WidgetFormula, bool) {
	if o == nil || o.Formulas == nil {
		return nil, false
	}
	return &o.Formulas, true
}

// HasFormulas returns a boolean if a field has been set.
func (o *GeomapWidgetRequest) HasFormulas() bool {
	return o != nil && o.Formulas != nil
}

// SetFormulas gets a reference to the given []WidgetFormula and assigns it to the Formulas field.
func (o *GeomapWidgetRequest) SetFormulas(v []WidgetFormula) {
	o.Formulas = v
}

// GetLogQuery returns the LogQuery field value if set, zero value otherwise.
func (o *GeomapWidgetRequest) GetLogQuery() LogQueryDefinition {
	if o == nil || o.LogQuery == nil {
		var ret LogQueryDefinition
		return ret
	}
	return *o.LogQuery
}

// GetLogQueryOk returns a tuple with the LogQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeomapWidgetRequest) GetLogQueryOk() (*LogQueryDefinition, bool) {
	if o == nil || o.LogQuery == nil {
		return nil, false
	}
	return o.LogQuery, true
}

// HasLogQuery returns a boolean if a field has been set.
func (o *GeomapWidgetRequest) HasLogQuery() bool {
	return o != nil && o.LogQuery != nil
}

// SetLogQuery gets a reference to the given LogQueryDefinition and assigns it to the LogQuery field.
func (o *GeomapWidgetRequest) SetLogQuery(v LogQueryDefinition) {
	o.LogQuery = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *GeomapWidgetRequest) GetQ() string {
	if o == nil || o.Q == nil {
		var ret string
		return ret
	}
	return *o.Q
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeomapWidgetRequest) GetQOk() (*string, bool) {
	if o == nil || o.Q == nil {
		return nil, false
	}
	return o.Q, true
}

// HasQ returns a boolean if a field has been set.
func (o *GeomapWidgetRequest) HasQ() bool {
	return o != nil && o.Q != nil
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *GeomapWidgetRequest) SetQ(v string) {
	o.Q = &v
}

// GetQueries returns the Queries field value if set, zero value otherwise.
func (o *GeomapWidgetRequest) GetQueries() []FormulaAndFunctionQueryDefinition {
	if o == nil || o.Queries == nil {
		var ret []FormulaAndFunctionQueryDefinition
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeomapWidgetRequest) GetQueriesOk() (*[]FormulaAndFunctionQueryDefinition, bool) {
	if o == nil || o.Queries == nil {
		return nil, false
	}
	return &o.Queries, true
}

// HasQueries returns a boolean if a field has been set.
func (o *GeomapWidgetRequest) HasQueries() bool {
	return o != nil && o.Queries != nil
}

// SetQueries gets a reference to the given []FormulaAndFunctionQueryDefinition and assigns it to the Queries field.
func (o *GeomapWidgetRequest) SetQueries(v []FormulaAndFunctionQueryDefinition) {
	o.Queries = v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *GeomapWidgetRequest) GetQuery() ListStreamQuery {
	if o == nil || o.Query == nil {
		var ret ListStreamQuery
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeomapWidgetRequest) GetQueryOk() (*ListStreamQuery, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *GeomapWidgetRequest) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given ListStreamQuery and assigns it to the Query field.
func (o *GeomapWidgetRequest) SetQuery(v ListStreamQuery) {
	o.Query = &v
}

// GetResponseFormat returns the ResponseFormat field value if set, zero value otherwise.
func (o *GeomapWidgetRequest) GetResponseFormat() FormulaAndFunctionResponseFormat {
	if o == nil || o.ResponseFormat == nil {
		var ret FormulaAndFunctionResponseFormat
		return ret
	}
	return *o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeomapWidgetRequest) GetResponseFormatOk() (*FormulaAndFunctionResponseFormat, bool) {
	if o == nil || o.ResponseFormat == nil {
		return nil, false
	}
	return o.ResponseFormat, true
}

// HasResponseFormat returns a boolean if a field has been set.
func (o *GeomapWidgetRequest) HasResponseFormat() bool {
	return o != nil && o.ResponseFormat != nil
}

// SetResponseFormat gets a reference to the given FormulaAndFunctionResponseFormat and assigns it to the ResponseFormat field.
func (o *GeomapWidgetRequest) SetResponseFormat(v FormulaAndFunctionResponseFormat) {
	o.ResponseFormat = &v
}

// GetRumQuery returns the RumQuery field value if set, zero value otherwise.
func (o *GeomapWidgetRequest) GetRumQuery() LogQueryDefinition {
	if o == nil || o.RumQuery == nil {
		var ret LogQueryDefinition
		return ret
	}
	return *o.RumQuery
}

// GetRumQueryOk returns a tuple with the RumQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeomapWidgetRequest) GetRumQueryOk() (*LogQueryDefinition, bool) {
	if o == nil || o.RumQuery == nil {
		return nil, false
	}
	return o.RumQuery, true
}

// HasRumQuery returns a boolean if a field has been set.
func (o *GeomapWidgetRequest) HasRumQuery() bool {
	return o != nil && o.RumQuery != nil
}

// SetRumQuery gets a reference to the given LogQueryDefinition and assigns it to the RumQuery field.
func (o *GeomapWidgetRequest) SetRumQuery(v LogQueryDefinition) {
	o.RumQuery = &v
}

// GetSecurityQuery returns the SecurityQuery field value if set, zero value otherwise.
func (o *GeomapWidgetRequest) GetSecurityQuery() LogQueryDefinition {
	if o == nil || o.SecurityQuery == nil {
		var ret LogQueryDefinition
		return ret
	}
	return *o.SecurityQuery
}

// GetSecurityQueryOk returns a tuple with the SecurityQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeomapWidgetRequest) GetSecurityQueryOk() (*LogQueryDefinition, bool) {
	if o == nil || o.SecurityQuery == nil {
		return nil, false
	}
	return o.SecurityQuery, true
}

// HasSecurityQuery returns a boolean if a field has been set.
func (o *GeomapWidgetRequest) HasSecurityQuery() bool {
	return o != nil && o.SecurityQuery != nil
}

// SetSecurityQuery gets a reference to the given LogQueryDefinition and assigns it to the SecurityQuery field.
func (o *GeomapWidgetRequest) SetSecurityQuery(v LogQueryDefinition) {
	o.SecurityQuery = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *GeomapWidgetRequest) GetSort() WidgetSortBy {
	if o == nil || o.Sort == nil {
		var ret WidgetSortBy
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeomapWidgetRequest) GetSortOk() (*WidgetSortBy, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *GeomapWidgetRequest) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given WidgetSortBy and assigns it to the Sort field.
func (o *GeomapWidgetRequest) SetSort(v WidgetSortBy) {
	o.Sort = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GeomapWidgetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	if o.Formulas != nil {
		toSerialize["formulas"] = o.Formulas
	}
	if o.LogQuery != nil {
		toSerialize["log_query"] = o.LogQuery
	}
	if o.Q != nil {
		toSerialize["q"] = o.Q
	}
	if o.Queries != nil {
		toSerialize["queries"] = o.Queries
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GeomapWidgetRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Columns        []ListStreamColumn                  `json:"columns,omitempty"`
		Formulas       []WidgetFormula                     `json:"formulas,omitempty"`
		LogQuery       *LogQueryDefinition                 `json:"log_query,omitempty"`
		Q              *string                             `json:"q,omitempty"`
		Queries        []FormulaAndFunctionQueryDefinition `json:"queries,omitempty"`
		Query          *ListStreamQuery                    `json:"query,omitempty"`
		ResponseFormat *FormulaAndFunctionResponseFormat   `json:"response_format,omitempty"`
		RumQuery       *LogQueryDefinition                 `json:"rum_query,omitempty"`
		SecurityQuery  *LogQueryDefinition                 `json:"security_query,omitempty"`
		Sort           *WidgetSortBy                       `json:"sort,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"columns", "formulas", "log_query", "q", "queries", "query", "response_format", "rum_query", "security_query", "sort"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Columns = all.Columns
	o.Formulas = all.Formulas
	if all.LogQuery != nil && all.LogQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LogQuery = all.LogQuery
	o.Q = all.Q
	o.Queries = all.Queries
	if all.Query != nil && all.Query.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Query = all.Query
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

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
