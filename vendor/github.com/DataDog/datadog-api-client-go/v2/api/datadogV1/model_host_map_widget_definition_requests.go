// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HostMapWidgetDefinitionRequests Query definition for the host map widget. Supports two mutually exclusive formats distinguished by the presence of `request_type`: the legacy metric-based format (`fill`/`size`) and the infrastructure-backed format (`request_type`, `node_type`, `enrichments`).
type HostMapWidgetDefinitionRequests struct {
	// Infrastructure-backed request for the host map widget. Supports entity-based
	// visualization with metric query enrichments, tag-based filtering, flexible grouping,
	// and hierarchical views.
	Child *HostMapWidgetInfrastructureRequest `json:"child,omitempty"`
	// List of conditional formatting rules applied to fill values.
	ConditionalFormats []WidgetConditionalFormat `json:"conditional_formats,omitempty"`
	// Metric or event queries joined to the entity set. Each formula specifies a visual dimension.
	Enrichments []HostMapWidgetScalarRequest `json:"enrichments,omitempty"`
	// Updated host map.
	Fill *HostMapRequest `json:"fill,omitempty"`
	// Filter string for the entity set in tag format (for example, `env:prod`).
	Filter *string `json:"filter,omitempty"`
	// Defines how entities are grouped into tiles. The ordering of entries implies
	// the grouping hierarchy.
	GroupBy []HostMapWidgetGroupBy `json:"group_by,omitempty"`
	// Whether to hide entities that have no group assignment.
	NoGroupHosts *bool `json:"no_group_hosts,omitempty"`
	// Whether to hide entities that have no enrichment data.
	NoMetricHosts *bool `json:"no_metric_hosts,omitempty"`
	// Which type of infrastructure entity to visualize in the host map.
	NodeType *HostMapWidgetNodeType `json:"node_type,omitempty"`
	// Identifies this as an infrastructure-backed host map request.
	RequestType *HostMapWidgetInfrastructureRequestRequestType `json:"request_type,omitempty"`
	// Updated host map.
	Size *HostMapRequest `json:"size,omitempty"`
	// Style configuration for the infrastructure host map.
	Style *HostMapWidgetInfrastructureStyle `json:"style,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewHostMapWidgetDefinitionRequests instantiates a new HostMapWidgetDefinitionRequests object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHostMapWidgetDefinitionRequests() *HostMapWidgetDefinitionRequests {
	this := HostMapWidgetDefinitionRequests{}
	return &this
}

// NewHostMapWidgetDefinitionRequestsWithDefaults instantiates a new HostMapWidgetDefinitionRequests object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHostMapWidgetDefinitionRequestsWithDefaults() *HostMapWidgetDefinitionRequests {
	this := HostMapWidgetDefinitionRequests{}
	return &this
}

// GetChild returns the Child field value if set, zero value otherwise.
func (o *HostMapWidgetDefinitionRequests) GetChild() HostMapWidgetInfrastructureRequest {
	if o == nil || o.Child == nil {
		var ret HostMapWidgetInfrastructureRequest
		return ret
	}
	return *o.Child
}

// GetChildOk returns a tuple with the Child field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapWidgetDefinitionRequests) GetChildOk() (*HostMapWidgetInfrastructureRequest, bool) {
	if o == nil || o.Child == nil {
		return nil, false
	}
	return o.Child, true
}

// HasChild returns a boolean if a field has been set.
func (o *HostMapWidgetDefinitionRequests) HasChild() bool {
	return o != nil && o.Child != nil
}

// SetChild gets a reference to the given HostMapWidgetInfrastructureRequest and assigns it to the Child field.
func (o *HostMapWidgetDefinitionRequests) SetChild(v HostMapWidgetInfrastructureRequest) {
	o.Child = &v
}

// GetConditionalFormats returns the ConditionalFormats field value if set, zero value otherwise.
func (o *HostMapWidgetDefinitionRequests) GetConditionalFormats() []WidgetConditionalFormat {
	if o == nil || o.ConditionalFormats == nil {
		var ret []WidgetConditionalFormat
		return ret
	}
	return o.ConditionalFormats
}

// GetConditionalFormatsOk returns a tuple with the ConditionalFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapWidgetDefinitionRequests) GetConditionalFormatsOk() (*[]WidgetConditionalFormat, bool) {
	if o == nil || o.ConditionalFormats == nil {
		return nil, false
	}
	return &o.ConditionalFormats, true
}

// HasConditionalFormats returns a boolean if a field has been set.
func (o *HostMapWidgetDefinitionRequests) HasConditionalFormats() bool {
	return o != nil && o.ConditionalFormats != nil
}

// SetConditionalFormats gets a reference to the given []WidgetConditionalFormat and assigns it to the ConditionalFormats field.
func (o *HostMapWidgetDefinitionRequests) SetConditionalFormats(v []WidgetConditionalFormat) {
	o.ConditionalFormats = v
}

// GetEnrichments returns the Enrichments field value if set, zero value otherwise.
func (o *HostMapWidgetDefinitionRequests) GetEnrichments() []HostMapWidgetScalarRequest {
	if o == nil || o.Enrichments == nil {
		var ret []HostMapWidgetScalarRequest
		return ret
	}
	return o.Enrichments
}

// GetEnrichmentsOk returns a tuple with the Enrichments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapWidgetDefinitionRequests) GetEnrichmentsOk() (*[]HostMapWidgetScalarRequest, bool) {
	if o == nil || o.Enrichments == nil {
		return nil, false
	}
	return &o.Enrichments, true
}

// HasEnrichments returns a boolean if a field has been set.
func (o *HostMapWidgetDefinitionRequests) HasEnrichments() bool {
	return o != nil && o.Enrichments != nil
}

// SetEnrichments gets a reference to the given []HostMapWidgetScalarRequest and assigns it to the Enrichments field.
func (o *HostMapWidgetDefinitionRequests) SetEnrichments(v []HostMapWidgetScalarRequest) {
	o.Enrichments = v
}

// GetFill returns the Fill field value if set, zero value otherwise.
func (o *HostMapWidgetDefinitionRequests) GetFill() HostMapRequest {
	if o == nil || o.Fill == nil {
		var ret HostMapRequest
		return ret
	}
	return *o.Fill
}

// GetFillOk returns a tuple with the Fill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapWidgetDefinitionRequests) GetFillOk() (*HostMapRequest, bool) {
	if o == nil || o.Fill == nil {
		return nil, false
	}
	return o.Fill, true
}

// HasFill returns a boolean if a field has been set.
func (o *HostMapWidgetDefinitionRequests) HasFill() bool {
	return o != nil && o.Fill != nil
}

// SetFill gets a reference to the given HostMapRequest and assigns it to the Fill field.
func (o *HostMapWidgetDefinitionRequests) SetFill(v HostMapRequest) {
	o.Fill = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *HostMapWidgetDefinitionRequests) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapWidgetDefinitionRequests) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *HostMapWidgetDefinitionRequests) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *HostMapWidgetDefinitionRequests) SetFilter(v string) {
	o.Filter = &v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *HostMapWidgetDefinitionRequests) GetGroupBy() []HostMapWidgetGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []HostMapWidgetGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapWidgetDefinitionRequests) GetGroupByOk() (*[]HostMapWidgetGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *HostMapWidgetDefinitionRequests) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []HostMapWidgetGroupBy and assigns it to the GroupBy field.
func (o *HostMapWidgetDefinitionRequests) SetGroupBy(v []HostMapWidgetGroupBy) {
	o.GroupBy = v
}

// GetNoGroupHosts returns the NoGroupHosts field value if set, zero value otherwise.
func (o *HostMapWidgetDefinitionRequests) GetNoGroupHosts() bool {
	if o == nil || o.NoGroupHosts == nil {
		var ret bool
		return ret
	}
	return *o.NoGroupHosts
}

// GetNoGroupHostsOk returns a tuple with the NoGroupHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapWidgetDefinitionRequests) GetNoGroupHostsOk() (*bool, bool) {
	if o == nil || o.NoGroupHosts == nil {
		return nil, false
	}
	return o.NoGroupHosts, true
}

// HasNoGroupHosts returns a boolean if a field has been set.
func (o *HostMapWidgetDefinitionRequests) HasNoGroupHosts() bool {
	return o != nil && o.NoGroupHosts != nil
}

// SetNoGroupHosts gets a reference to the given bool and assigns it to the NoGroupHosts field.
func (o *HostMapWidgetDefinitionRequests) SetNoGroupHosts(v bool) {
	o.NoGroupHosts = &v
}

// GetNoMetricHosts returns the NoMetricHosts field value if set, zero value otherwise.
func (o *HostMapWidgetDefinitionRequests) GetNoMetricHosts() bool {
	if o == nil || o.NoMetricHosts == nil {
		var ret bool
		return ret
	}
	return *o.NoMetricHosts
}

// GetNoMetricHostsOk returns a tuple with the NoMetricHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapWidgetDefinitionRequests) GetNoMetricHostsOk() (*bool, bool) {
	if o == nil || o.NoMetricHosts == nil {
		return nil, false
	}
	return o.NoMetricHosts, true
}

// HasNoMetricHosts returns a boolean if a field has been set.
func (o *HostMapWidgetDefinitionRequests) HasNoMetricHosts() bool {
	return o != nil && o.NoMetricHosts != nil
}

// SetNoMetricHosts gets a reference to the given bool and assigns it to the NoMetricHosts field.
func (o *HostMapWidgetDefinitionRequests) SetNoMetricHosts(v bool) {
	o.NoMetricHosts = &v
}

// GetNodeType returns the NodeType field value if set, zero value otherwise.
func (o *HostMapWidgetDefinitionRequests) GetNodeType() HostMapWidgetNodeType {
	if o == nil || o.NodeType == nil {
		var ret HostMapWidgetNodeType
		return ret
	}
	return *o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapWidgetDefinitionRequests) GetNodeTypeOk() (*HostMapWidgetNodeType, bool) {
	if o == nil || o.NodeType == nil {
		return nil, false
	}
	return o.NodeType, true
}

// HasNodeType returns a boolean if a field has been set.
func (o *HostMapWidgetDefinitionRequests) HasNodeType() bool {
	return o != nil && o.NodeType != nil
}

// SetNodeType gets a reference to the given HostMapWidgetNodeType and assigns it to the NodeType field.
func (o *HostMapWidgetDefinitionRequests) SetNodeType(v HostMapWidgetNodeType) {
	o.NodeType = &v
}

// GetRequestType returns the RequestType field value if set, zero value otherwise.
func (o *HostMapWidgetDefinitionRequests) GetRequestType() HostMapWidgetInfrastructureRequestRequestType {
	if o == nil || o.RequestType == nil {
		var ret HostMapWidgetInfrastructureRequestRequestType
		return ret
	}
	return *o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapWidgetDefinitionRequests) GetRequestTypeOk() (*HostMapWidgetInfrastructureRequestRequestType, bool) {
	if o == nil || o.RequestType == nil {
		return nil, false
	}
	return o.RequestType, true
}

// HasRequestType returns a boolean if a field has been set.
func (o *HostMapWidgetDefinitionRequests) HasRequestType() bool {
	return o != nil && o.RequestType != nil
}

// SetRequestType gets a reference to the given HostMapWidgetInfrastructureRequestRequestType and assigns it to the RequestType field.
func (o *HostMapWidgetDefinitionRequests) SetRequestType(v HostMapWidgetInfrastructureRequestRequestType) {
	o.RequestType = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *HostMapWidgetDefinitionRequests) GetSize() HostMapRequest {
	if o == nil || o.Size == nil {
		var ret HostMapRequest
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapWidgetDefinitionRequests) GetSizeOk() (*HostMapRequest, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *HostMapWidgetDefinitionRequests) HasSize() bool {
	return o != nil && o.Size != nil
}

// SetSize gets a reference to the given HostMapRequest and assigns it to the Size field.
func (o *HostMapWidgetDefinitionRequests) SetSize(v HostMapRequest) {
	o.Size = &v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *HostMapWidgetDefinitionRequests) GetStyle() HostMapWidgetInfrastructureStyle {
	if o == nil || o.Style == nil {
		var ret HostMapWidgetInfrastructureStyle
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapWidgetDefinitionRequests) GetStyleOk() (*HostMapWidgetInfrastructureStyle, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *HostMapWidgetDefinitionRequests) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given HostMapWidgetInfrastructureStyle and assigns it to the Style field.
func (o *HostMapWidgetDefinitionRequests) SetStyle(v HostMapWidgetInfrastructureStyle) {
	o.Style = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o HostMapWidgetDefinitionRequests) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Child != nil {
		toSerialize["child"] = o.Child
	}
	if o.ConditionalFormats != nil {
		toSerialize["conditional_formats"] = o.ConditionalFormats
	}
	if o.Enrichments != nil {
		toSerialize["enrichments"] = o.Enrichments
	}
	if o.Fill != nil {
		toSerialize["fill"] = o.Fill
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	if o.NoGroupHosts != nil {
		toSerialize["no_group_hosts"] = o.NoGroupHosts
	}
	if o.NoMetricHosts != nil {
		toSerialize["no_metric_hosts"] = o.NoMetricHosts
	}
	if o.NodeType != nil {
		toSerialize["node_type"] = o.NodeType
	}
	if o.RequestType != nil {
		toSerialize["request_type"] = o.RequestType
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
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
func (o *HostMapWidgetDefinitionRequests) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Child              *HostMapWidgetInfrastructureRequest            `json:"child,omitempty"`
		ConditionalFormats []WidgetConditionalFormat                      `json:"conditional_formats,omitempty"`
		Enrichments        []HostMapWidgetScalarRequest                   `json:"enrichments,omitempty"`
		Fill               *HostMapRequest                                `json:"fill,omitempty"`
		Filter             *string                                        `json:"filter,omitempty"`
		GroupBy            []HostMapWidgetGroupBy                         `json:"group_by,omitempty"`
		NoGroupHosts       *bool                                          `json:"no_group_hosts,omitempty"`
		NoMetricHosts      *bool                                          `json:"no_metric_hosts,omitempty"`
		NodeType           *HostMapWidgetNodeType                         `json:"node_type,omitempty"`
		RequestType        *HostMapWidgetInfrastructureRequestRequestType `json:"request_type,omitempty"`
		Size               *HostMapRequest                                `json:"size,omitempty"`
		Style              *HostMapWidgetInfrastructureStyle              `json:"style,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"child", "conditional_formats", "enrichments", "fill", "filter", "group_by", "no_group_hosts", "no_metric_hosts", "node_type", "request_type", "size", "style"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Child != nil && all.Child.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Child = all.Child
	o.ConditionalFormats = all.ConditionalFormats
	o.Enrichments = all.Enrichments
	if all.Fill != nil && all.Fill.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Fill = all.Fill
	o.Filter = all.Filter
	o.GroupBy = all.GroupBy
	o.NoGroupHosts = all.NoGroupHosts
	o.NoMetricHosts = all.NoMetricHosts
	if all.NodeType != nil && !all.NodeType.IsValid() {
		hasInvalidField = true
	} else {
		o.NodeType = all.NodeType
	}
	if all.RequestType != nil && !all.RequestType.IsValid() {
		hasInvalidField = true
	} else {
		o.RequestType = all.RequestType
	}
	if all.Size != nil && all.Size.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Size = all.Size
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
