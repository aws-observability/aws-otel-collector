// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaAndFunctionApmMetricsQueryDefinition A formula and functions APM metrics query.
type FormulaAndFunctionApmMetricsQueryDefinition struct {
	// Data source for APM metrics queries.
	DataSource FormulaAndFunctionApmMetricsDataSource `json:"data_source"`
	// Optional fields to group the query results by.
	GroupBy []string `json:"group_by,omitempty"`
	// Name of this query to use in formulas.
	Name string `json:"name"`
	// Optional operation mode to aggregate across operation names.
	OperationMode *string `json:"operation_mode,omitempty"`
	// Name of operation on service. If not provided, the primary operation name is used.
	OperationName *string `json:"operation_name,omitempty"`
	// Tags to query for a specific downstream entity (peer.service, peer.db_instance, peer.s3, peer.s3.bucket, etc.).
	PeerTags []string `json:"peer_tags,omitempty"`
	// Additional filters for the query using metrics query syntax (e.g., env, primary_tag).
	QueryFilter *string `json:"query_filter,omitempty"`
	// The hash of a specific resource to filter by.
	ResourceHash *string `json:"resource_hash,omitempty"`
	// The full name of a specific resource to filter by.
	ResourceName *string `json:"resource_name,omitempty"`
	// APM service name.
	Service *string `json:"service,omitempty"`
	// Describes the relationship between the span, its parents, and its children in a trace.
	SpanKind *FormulaAndFunctionApmMetricsSpanKind `json:"span_kind,omitempty"`
	// APM metric stat name.
	Stat FormulaAndFunctionApmMetricStatName `json:"stat"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFormulaAndFunctionApmMetricsQueryDefinition instantiates a new FormulaAndFunctionApmMetricsQueryDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormulaAndFunctionApmMetricsQueryDefinition(dataSource FormulaAndFunctionApmMetricsDataSource, name string, stat FormulaAndFunctionApmMetricStatName) *FormulaAndFunctionApmMetricsQueryDefinition {
	this := FormulaAndFunctionApmMetricsQueryDefinition{}
	this.DataSource = dataSource
	this.Name = name
	this.Stat = stat
	return &this
}

// NewFormulaAndFunctionApmMetricsQueryDefinitionWithDefaults instantiates a new FormulaAndFunctionApmMetricsQueryDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormulaAndFunctionApmMetricsQueryDefinitionWithDefaults() *FormulaAndFunctionApmMetricsQueryDefinition {
	this := FormulaAndFunctionApmMetricsQueryDefinition{}
	return &this
}

// GetDataSource returns the DataSource field value.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) GetDataSource() FormulaAndFunctionApmMetricsDataSource {
	if o == nil {
		var ret FormulaAndFunctionApmMetricsDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) GetDataSourceOk() (*FormulaAndFunctionApmMetricsDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) SetDataSource(v FormulaAndFunctionApmMetricsDataSource) {
	o.DataSource = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) GetGroupBy() []string {
	if o == nil || o.GroupBy == nil {
		var ret []string
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) GetGroupByOk() (*[]string, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []string and assigns it to the GroupBy field.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) SetGroupBy(v []string) {
	o.GroupBy = v
}

// GetName returns the Name field value.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) SetName(v string) {
	o.Name = v
}

// GetOperationMode returns the OperationMode field value if set, zero value otherwise.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) GetOperationMode() string {
	if o == nil || o.OperationMode == nil {
		var ret string
		return ret
	}
	return *o.OperationMode
}

// GetOperationModeOk returns a tuple with the OperationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) GetOperationModeOk() (*string, bool) {
	if o == nil || o.OperationMode == nil {
		return nil, false
	}
	return o.OperationMode, true
}

// HasOperationMode returns a boolean if a field has been set.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) HasOperationMode() bool {
	return o != nil && o.OperationMode != nil
}

// SetOperationMode gets a reference to the given string and assigns it to the OperationMode field.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) SetOperationMode(v string) {
	o.OperationMode = &v
}

// GetOperationName returns the OperationName field value if set, zero value otherwise.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) GetOperationName() string {
	if o == nil || o.OperationName == nil {
		var ret string
		return ret
	}
	return *o.OperationName
}

// GetOperationNameOk returns a tuple with the OperationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) GetOperationNameOk() (*string, bool) {
	if o == nil || o.OperationName == nil {
		return nil, false
	}
	return o.OperationName, true
}

// HasOperationName returns a boolean if a field has been set.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) HasOperationName() bool {
	return o != nil && o.OperationName != nil
}

// SetOperationName gets a reference to the given string and assigns it to the OperationName field.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) SetOperationName(v string) {
	o.OperationName = &v
}

// GetPeerTags returns the PeerTags field value if set, zero value otherwise.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) GetPeerTags() []string {
	if o == nil || o.PeerTags == nil {
		var ret []string
		return ret
	}
	return o.PeerTags
}

// GetPeerTagsOk returns a tuple with the PeerTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) GetPeerTagsOk() (*[]string, bool) {
	if o == nil || o.PeerTags == nil {
		return nil, false
	}
	return &o.PeerTags, true
}

// HasPeerTags returns a boolean if a field has been set.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) HasPeerTags() bool {
	return o != nil && o.PeerTags != nil
}

// SetPeerTags gets a reference to the given []string and assigns it to the PeerTags field.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) SetPeerTags(v []string) {
	o.PeerTags = v
}

// GetQueryFilter returns the QueryFilter field value if set, zero value otherwise.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) GetQueryFilter() string {
	if o == nil || o.QueryFilter == nil {
		var ret string
		return ret
	}
	return *o.QueryFilter
}

// GetQueryFilterOk returns a tuple with the QueryFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) GetQueryFilterOk() (*string, bool) {
	if o == nil || o.QueryFilter == nil {
		return nil, false
	}
	return o.QueryFilter, true
}

// HasQueryFilter returns a boolean if a field has been set.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) HasQueryFilter() bool {
	return o != nil && o.QueryFilter != nil
}

// SetQueryFilter gets a reference to the given string and assigns it to the QueryFilter field.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) SetQueryFilter(v string) {
	o.QueryFilter = &v
}

// GetResourceHash returns the ResourceHash field value if set, zero value otherwise.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) GetResourceHash() string {
	if o == nil || o.ResourceHash == nil {
		var ret string
		return ret
	}
	return *o.ResourceHash
}

// GetResourceHashOk returns a tuple with the ResourceHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) GetResourceHashOk() (*string, bool) {
	if o == nil || o.ResourceHash == nil {
		return nil, false
	}
	return o.ResourceHash, true
}

// HasResourceHash returns a boolean if a field has been set.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) HasResourceHash() bool {
	return o != nil && o.ResourceHash != nil
}

// SetResourceHash gets a reference to the given string and assigns it to the ResourceHash field.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) SetResourceHash(v string) {
	o.ResourceHash = &v
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) GetResourceName() string {
	if o == nil || o.ResourceName == nil {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) GetResourceNameOk() (*string, bool) {
	if o == nil || o.ResourceName == nil {
		return nil, false
	}
	return o.ResourceName, true
}

// HasResourceName returns a boolean if a field has been set.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) HasResourceName() bool {
	return o != nil && o.ResourceName != nil
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) SetResourceName(v string) {
	o.ResourceName = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) HasService() bool {
	return o != nil && o.Service != nil
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) SetService(v string) {
	o.Service = &v
}

// GetSpanKind returns the SpanKind field value if set, zero value otherwise.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) GetSpanKind() FormulaAndFunctionApmMetricsSpanKind {
	if o == nil || o.SpanKind == nil {
		var ret FormulaAndFunctionApmMetricsSpanKind
		return ret
	}
	return *o.SpanKind
}

// GetSpanKindOk returns a tuple with the SpanKind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) GetSpanKindOk() (*FormulaAndFunctionApmMetricsSpanKind, bool) {
	if o == nil || o.SpanKind == nil {
		return nil, false
	}
	return o.SpanKind, true
}

// HasSpanKind returns a boolean if a field has been set.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) HasSpanKind() bool {
	return o != nil && o.SpanKind != nil
}

// SetSpanKind gets a reference to the given FormulaAndFunctionApmMetricsSpanKind and assigns it to the SpanKind field.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) SetSpanKind(v FormulaAndFunctionApmMetricsSpanKind) {
	o.SpanKind = &v
}

// GetStat returns the Stat field value.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) GetStat() FormulaAndFunctionApmMetricStatName {
	if o == nil {
		var ret FormulaAndFunctionApmMetricStatName
		return ret
	}
	return o.Stat
}

// GetStatOk returns a tuple with the Stat field value
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) GetStatOk() (*FormulaAndFunctionApmMetricStatName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stat, true
}

// SetStat sets field value.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) SetStat(v FormulaAndFunctionApmMetricStatName) {
	o.Stat = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormulaAndFunctionApmMetricsQueryDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data_source"] = o.DataSource
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	toSerialize["name"] = o.Name
	if o.OperationMode != nil {
		toSerialize["operation_mode"] = o.OperationMode
	}
	if o.OperationName != nil {
		toSerialize["operation_name"] = o.OperationName
	}
	if o.PeerTags != nil {
		toSerialize["peer_tags"] = o.PeerTags
	}
	if o.QueryFilter != nil {
		toSerialize["query_filter"] = o.QueryFilter
	}
	if o.ResourceHash != nil {
		toSerialize["resource_hash"] = o.ResourceHash
	}
	if o.ResourceName != nil {
		toSerialize["resource_name"] = o.ResourceName
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if o.SpanKind != nil {
		toSerialize["span_kind"] = o.SpanKind
	}
	toSerialize["stat"] = o.Stat

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormulaAndFunctionApmMetricsQueryDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataSource    *FormulaAndFunctionApmMetricsDataSource `json:"data_source"`
		GroupBy       []string                                `json:"group_by,omitempty"`
		Name          *string                                 `json:"name"`
		OperationMode *string                                 `json:"operation_mode,omitempty"`
		OperationName *string                                 `json:"operation_name,omitempty"`
		PeerTags      []string                                `json:"peer_tags,omitempty"`
		QueryFilter   *string                                 `json:"query_filter,omitempty"`
		ResourceHash  *string                                 `json:"resource_hash,omitempty"`
		ResourceName  *string                                 `json:"resource_name,omitempty"`
		Service       *string                                 `json:"service,omitempty"`
		SpanKind      *FormulaAndFunctionApmMetricsSpanKind   `json:"span_kind,omitempty"`
		Stat          *FormulaAndFunctionApmMetricStatName    `json:"stat"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Stat == nil {
		return fmt.Errorf("required field stat missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data_source", "group_by", "name", "operation_mode", "operation_name", "peer_tags", "query_filter", "resource_hash", "resource_name", "service", "span_kind", "stat"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.GroupBy = all.GroupBy
	o.Name = *all.Name
	o.OperationMode = all.OperationMode
	o.OperationName = all.OperationName
	o.PeerTags = all.PeerTags
	o.QueryFilter = all.QueryFilter
	o.ResourceHash = all.ResourceHash
	o.ResourceName = all.ResourceName
	o.Service = all.Service
	if all.SpanKind != nil && !all.SpanKind.IsValid() {
		hasInvalidField = true
	} else {
		o.SpanKind = all.SpanKind
	}
	if !all.Stat.IsValid() {
		hasInvalidField = true
	} else {
		o.Stat = *all.Stat
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
