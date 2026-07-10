// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HostMapWidgetInfrastructureRequestLeaf Infrastructure-backed host map child request (leaf node, no further nesting supported).
type HostMapWidgetInfrastructureRequestLeaf struct {
	// List of conditional formatting rules applied to fill values.
	ConditionalFormats []WidgetConditionalFormat `json:"conditional_formats,omitempty"`
	// Metric or event queries joined to the entity set. Each formula specifies a visual dimension.
	Enrichments []HostMapWidgetScalarRequest `json:"enrichments"`
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
	NodeType HostMapWidgetNodeType `json:"node_type"`
	// Identifies this as an infrastructure-backed host map request.
	RequestType HostMapWidgetInfrastructureRequestRequestType `json:"request_type"`
	// Style configuration for the infrastructure host map.
	Style *HostMapWidgetInfrastructureStyle `json:"style,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewHostMapWidgetInfrastructureRequestLeaf instantiates a new HostMapWidgetInfrastructureRequestLeaf object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHostMapWidgetInfrastructureRequestLeaf(enrichments []HostMapWidgetScalarRequest, nodeType HostMapWidgetNodeType, requestType HostMapWidgetInfrastructureRequestRequestType) *HostMapWidgetInfrastructureRequestLeaf {
	this := HostMapWidgetInfrastructureRequestLeaf{}
	this.Enrichments = enrichments
	this.NodeType = nodeType
	this.RequestType = requestType
	return &this
}

// NewHostMapWidgetInfrastructureRequestLeafWithDefaults instantiates a new HostMapWidgetInfrastructureRequestLeaf object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHostMapWidgetInfrastructureRequestLeafWithDefaults() *HostMapWidgetInfrastructureRequestLeaf {
	this := HostMapWidgetInfrastructureRequestLeaf{}
	return &this
}

// GetConditionalFormats returns the ConditionalFormats field value if set, zero value otherwise.
func (o *HostMapWidgetInfrastructureRequestLeaf) GetConditionalFormats() []WidgetConditionalFormat {
	if o == nil || o.ConditionalFormats == nil {
		var ret []WidgetConditionalFormat
		return ret
	}
	return o.ConditionalFormats
}

// GetConditionalFormatsOk returns a tuple with the ConditionalFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapWidgetInfrastructureRequestLeaf) GetConditionalFormatsOk() (*[]WidgetConditionalFormat, bool) {
	if o == nil || o.ConditionalFormats == nil {
		return nil, false
	}
	return &o.ConditionalFormats, true
}

// HasConditionalFormats returns a boolean if a field has been set.
func (o *HostMapWidgetInfrastructureRequestLeaf) HasConditionalFormats() bool {
	return o != nil && o.ConditionalFormats != nil
}

// SetConditionalFormats gets a reference to the given []WidgetConditionalFormat and assigns it to the ConditionalFormats field.
func (o *HostMapWidgetInfrastructureRequestLeaf) SetConditionalFormats(v []WidgetConditionalFormat) {
	o.ConditionalFormats = v
}

// GetEnrichments returns the Enrichments field value.
func (o *HostMapWidgetInfrastructureRequestLeaf) GetEnrichments() []HostMapWidgetScalarRequest {
	if o == nil {
		var ret []HostMapWidgetScalarRequest
		return ret
	}
	return o.Enrichments
}

// GetEnrichmentsOk returns a tuple with the Enrichments field value
// and a boolean to check if the value has been set.
func (o *HostMapWidgetInfrastructureRequestLeaf) GetEnrichmentsOk() (*[]HostMapWidgetScalarRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enrichments, true
}

// SetEnrichments sets field value.
func (o *HostMapWidgetInfrastructureRequestLeaf) SetEnrichments(v []HostMapWidgetScalarRequest) {
	o.Enrichments = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *HostMapWidgetInfrastructureRequestLeaf) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapWidgetInfrastructureRequestLeaf) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *HostMapWidgetInfrastructureRequestLeaf) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *HostMapWidgetInfrastructureRequestLeaf) SetFilter(v string) {
	o.Filter = &v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *HostMapWidgetInfrastructureRequestLeaf) GetGroupBy() []HostMapWidgetGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []HostMapWidgetGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapWidgetInfrastructureRequestLeaf) GetGroupByOk() (*[]HostMapWidgetGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *HostMapWidgetInfrastructureRequestLeaf) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []HostMapWidgetGroupBy and assigns it to the GroupBy field.
func (o *HostMapWidgetInfrastructureRequestLeaf) SetGroupBy(v []HostMapWidgetGroupBy) {
	o.GroupBy = v
}

// GetNoGroupHosts returns the NoGroupHosts field value if set, zero value otherwise.
func (o *HostMapWidgetInfrastructureRequestLeaf) GetNoGroupHosts() bool {
	if o == nil || o.NoGroupHosts == nil {
		var ret bool
		return ret
	}
	return *o.NoGroupHosts
}

// GetNoGroupHostsOk returns a tuple with the NoGroupHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapWidgetInfrastructureRequestLeaf) GetNoGroupHostsOk() (*bool, bool) {
	if o == nil || o.NoGroupHosts == nil {
		return nil, false
	}
	return o.NoGroupHosts, true
}

// HasNoGroupHosts returns a boolean if a field has been set.
func (o *HostMapWidgetInfrastructureRequestLeaf) HasNoGroupHosts() bool {
	return o != nil && o.NoGroupHosts != nil
}

// SetNoGroupHosts gets a reference to the given bool and assigns it to the NoGroupHosts field.
func (o *HostMapWidgetInfrastructureRequestLeaf) SetNoGroupHosts(v bool) {
	o.NoGroupHosts = &v
}

// GetNoMetricHosts returns the NoMetricHosts field value if set, zero value otherwise.
func (o *HostMapWidgetInfrastructureRequestLeaf) GetNoMetricHosts() bool {
	if o == nil || o.NoMetricHosts == nil {
		var ret bool
		return ret
	}
	return *o.NoMetricHosts
}

// GetNoMetricHostsOk returns a tuple with the NoMetricHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapWidgetInfrastructureRequestLeaf) GetNoMetricHostsOk() (*bool, bool) {
	if o == nil || o.NoMetricHosts == nil {
		return nil, false
	}
	return o.NoMetricHosts, true
}

// HasNoMetricHosts returns a boolean if a field has been set.
func (o *HostMapWidgetInfrastructureRequestLeaf) HasNoMetricHosts() bool {
	return o != nil && o.NoMetricHosts != nil
}

// SetNoMetricHosts gets a reference to the given bool and assigns it to the NoMetricHosts field.
func (o *HostMapWidgetInfrastructureRequestLeaf) SetNoMetricHosts(v bool) {
	o.NoMetricHosts = &v
}

// GetNodeType returns the NodeType field value.
func (o *HostMapWidgetInfrastructureRequestLeaf) GetNodeType() HostMapWidgetNodeType {
	if o == nil {
		var ret HostMapWidgetNodeType
		return ret
	}
	return o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value
// and a boolean to check if the value has been set.
func (o *HostMapWidgetInfrastructureRequestLeaf) GetNodeTypeOk() (*HostMapWidgetNodeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeType, true
}

// SetNodeType sets field value.
func (o *HostMapWidgetInfrastructureRequestLeaf) SetNodeType(v HostMapWidgetNodeType) {
	o.NodeType = v
}

// GetRequestType returns the RequestType field value.
func (o *HostMapWidgetInfrastructureRequestLeaf) GetRequestType() HostMapWidgetInfrastructureRequestRequestType {
	if o == nil {
		var ret HostMapWidgetInfrastructureRequestRequestType
		return ret
	}
	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *HostMapWidgetInfrastructureRequestLeaf) GetRequestTypeOk() (*HostMapWidgetInfrastructureRequestRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value.
func (o *HostMapWidgetInfrastructureRequestLeaf) SetRequestType(v HostMapWidgetInfrastructureRequestRequestType) {
	o.RequestType = v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *HostMapWidgetInfrastructureRequestLeaf) GetStyle() HostMapWidgetInfrastructureStyle {
	if o == nil || o.Style == nil {
		var ret HostMapWidgetInfrastructureStyle
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostMapWidgetInfrastructureRequestLeaf) GetStyleOk() (*HostMapWidgetInfrastructureStyle, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *HostMapWidgetInfrastructureRequestLeaf) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given HostMapWidgetInfrastructureStyle and assigns it to the Style field.
func (o *HostMapWidgetInfrastructureRequestLeaf) SetStyle(v HostMapWidgetInfrastructureStyle) {
	o.Style = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o HostMapWidgetInfrastructureRequestLeaf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ConditionalFormats != nil {
		toSerialize["conditional_formats"] = o.ConditionalFormats
	}
	toSerialize["enrichments"] = o.Enrichments
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
	toSerialize["node_type"] = o.NodeType
	toSerialize["request_type"] = o.RequestType
	if o.Style != nil {
		toSerialize["style"] = o.Style
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HostMapWidgetInfrastructureRequestLeaf) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConditionalFormats []WidgetConditionalFormat                      `json:"conditional_formats,omitempty"`
		Enrichments        *[]HostMapWidgetScalarRequest                  `json:"enrichments"`
		Filter             *string                                        `json:"filter,omitempty"`
		GroupBy            []HostMapWidgetGroupBy                         `json:"group_by,omitempty"`
		NoGroupHosts       *bool                                          `json:"no_group_hosts,omitempty"`
		NoMetricHosts      *bool                                          `json:"no_metric_hosts,omitempty"`
		NodeType           *HostMapWidgetNodeType                         `json:"node_type"`
		RequestType        *HostMapWidgetInfrastructureRequestRequestType `json:"request_type"`
		Style              *HostMapWidgetInfrastructureStyle              `json:"style,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Enrichments == nil {
		return fmt.Errorf("required field enrichments missing")
	}
	if all.NodeType == nil {
		return fmt.Errorf("required field node_type missing")
	}
	if all.RequestType == nil {
		return fmt.Errorf("required field request_type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"conditional_formats", "enrichments", "filter", "group_by", "no_group_hosts", "no_metric_hosts", "node_type", "request_type", "style"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ConditionalFormats = all.ConditionalFormats
	o.Enrichments = *all.Enrichments
	o.Filter = all.Filter
	o.GroupBy = all.GroupBy
	o.NoGroupHosts = all.NoGroupHosts
	o.NoMetricHosts = all.NoMetricHosts
	if !all.NodeType.IsValid() {
		hasInvalidField = true
	} else {
		o.NodeType = *all.NodeType
	}
	if !all.RequestType.IsValid() {
		hasInvalidField = true
	} else {
		o.RequestType = *all.RequestType
	}
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
