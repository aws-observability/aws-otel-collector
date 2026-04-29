// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SankeyWidgetDefinition The Sankey diagram visualizes the flow of data between categories, stages or sets of values.
type SankeyWidgetDefinition struct {
	// List of Sankey widget requests.
	Requests []SankeyWidgetRequest `json:"requests"`
	// Whether to show links for "other" category.
	ShowOtherLinks *bool `json:"show_other_links,omitempty"`
	// Whether to sort nodes in the Sankey diagram.
	SortNodes *bool `json:"sort_nodes,omitempty"`
	// Time setting for the widget.
	Time *WidgetTime `json:"time,omitempty"`
	// Title of your widget.
	Title *string `json:"title,omitempty"`
	// How to align the text on the widget.
	TitleAlign *WidgetTextAlign `json:"title_align,omitempty"`
	// Size of the title.
	TitleSize *string `json:"title_size,omitempty"`
	// Type of the Sankey widget.
	Type SankeyWidgetDefinitionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewSankeyWidgetDefinition instantiates a new SankeyWidgetDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSankeyWidgetDefinition(requests []SankeyWidgetRequest, typeVar SankeyWidgetDefinitionType) *SankeyWidgetDefinition {
	this := SankeyWidgetDefinition{}
	this.Requests = requests
	this.Type = typeVar
	return &this
}

// NewSankeyWidgetDefinitionWithDefaults instantiates a new SankeyWidgetDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSankeyWidgetDefinitionWithDefaults() *SankeyWidgetDefinition {
	this := SankeyWidgetDefinition{}
	var typeVar SankeyWidgetDefinitionType = SANKEYWIDGETDEFINITIONTYPE_SANKEY
	this.Type = typeVar
	return &this
}

// GetRequests returns the Requests field value.
func (o *SankeyWidgetDefinition) GetRequests() []SankeyWidgetRequest {
	if o == nil {
		var ret []SankeyWidgetRequest
		return ret
	}
	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *SankeyWidgetDefinition) GetRequestsOk() (*[]SankeyWidgetRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requests, true
}

// SetRequests sets field value.
func (o *SankeyWidgetDefinition) SetRequests(v []SankeyWidgetRequest) {
	o.Requests = v
}

// GetShowOtherLinks returns the ShowOtherLinks field value if set, zero value otherwise.
func (o *SankeyWidgetDefinition) GetShowOtherLinks() bool {
	if o == nil || o.ShowOtherLinks == nil {
		var ret bool
		return ret
	}
	return *o.ShowOtherLinks
}

// GetShowOtherLinksOk returns a tuple with the ShowOtherLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyWidgetDefinition) GetShowOtherLinksOk() (*bool, bool) {
	if o == nil || o.ShowOtherLinks == nil {
		return nil, false
	}
	return o.ShowOtherLinks, true
}

// HasShowOtherLinks returns a boolean if a field has been set.
func (o *SankeyWidgetDefinition) HasShowOtherLinks() bool {
	return o != nil && o.ShowOtherLinks != nil
}

// SetShowOtherLinks gets a reference to the given bool and assigns it to the ShowOtherLinks field.
func (o *SankeyWidgetDefinition) SetShowOtherLinks(v bool) {
	o.ShowOtherLinks = &v
}

// GetSortNodes returns the SortNodes field value if set, zero value otherwise.
func (o *SankeyWidgetDefinition) GetSortNodes() bool {
	if o == nil || o.SortNodes == nil {
		var ret bool
		return ret
	}
	return *o.SortNodes
}

// GetSortNodesOk returns a tuple with the SortNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyWidgetDefinition) GetSortNodesOk() (*bool, bool) {
	if o == nil || o.SortNodes == nil {
		return nil, false
	}
	return o.SortNodes, true
}

// HasSortNodes returns a boolean if a field has been set.
func (o *SankeyWidgetDefinition) HasSortNodes() bool {
	return o != nil && o.SortNodes != nil
}

// SetSortNodes gets a reference to the given bool and assigns it to the SortNodes field.
func (o *SankeyWidgetDefinition) SetSortNodes(v bool) {
	o.SortNodes = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *SankeyWidgetDefinition) GetTime() WidgetTime {
	if o == nil || o.Time == nil {
		var ret WidgetTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyWidgetDefinition) GetTimeOk() (*WidgetTime, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *SankeyWidgetDefinition) HasTime() bool {
	return o != nil && o.Time != nil
}

// SetTime gets a reference to the given WidgetTime and assigns it to the Time field.
func (o *SankeyWidgetDefinition) SetTime(v WidgetTime) {
	o.Time = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *SankeyWidgetDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyWidgetDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *SankeyWidgetDefinition) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *SankeyWidgetDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetTitleAlign returns the TitleAlign field value if set, zero value otherwise.
func (o *SankeyWidgetDefinition) GetTitleAlign() WidgetTextAlign {
	if o == nil || o.TitleAlign == nil {
		var ret WidgetTextAlign
		return ret
	}
	return *o.TitleAlign
}

// GetTitleAlignOk returns a tuple with the TitleAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyWidgetDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool) {
	if o == nil || o.TitleAlign == nil {
		return nil, false
	}
	return o.TitleAlign, true
}

// HasTitleAlign returns a boolean if a field has been set.
func (o *SankeyWidgetDefinition) HasTitleAlign() bool {
	return o != nil && o.TitleAlign != nil
}

// SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.
func (o *SankeyWidgetDefinition) SetTitleAlign(v WidgetTextAlign) {
	o.TitleAlign = &v
}

// GetTitleSize returns the TitleSize field value if set, zero value otherwise.
func (o *SankeyWidgetDefinition) GetTitleSize() string {
	if o == nil || o.TitleSize == nil {
		var ret string
		return ret
	}
	return *o.TitleSize
}

// GetTitleSizeOk returns a tuple with the TitleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyWidgetDefinition) GetTitleSizeOk() (*string, bool) {
	if o == nil || o.TitleSize == nil {
		return nil, false
	}
	return o.TitleSize, true
}

// HasTitleSize returns a boolean if a field has been set.
func (o *SankeyWidgetDefinition) HasTitleSize() bool {
	return o != nil && o.TitleSize != nil
}

// SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.
func (o *SankeyWidgetDefinition) SetTitleSize(v string) {
	o.TitleSize = &v
}

// GetType returns the Type field value.
func (o *SankeyWidgetDefinition) GetType() SankeyWidgetDefinitionType {
	if o == nil {
		var ret SankeyWidgetDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SankeyWidgetDefinition) GetTypeOk() (*SankeyWidgetDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SankeyWidgetDefinition) SetType(v SankeyWidgetDefinitionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SankeyWidgetDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["requests"] = o.Requests
	if o.ShowOtherLinks != nil {
		toSerialize["show_other_links"] = o.ShowOtherLinks
	}
	if o.SortNodes != nil {
		toSerialize["sort_nodes"] = o.SortNodes
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.TitleAlign != nil {
		toSerialize["title_align"] = o.TitleAlign
	}
	if o.TitleSize != nil {
		toSerialize["title_size"] = o.TitleSize
	}
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SankeyWidgetDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Requests       *[]SankeyWidgetRequest      `json:"requests"`
		ShowOtherLinks *bool                       `json:"show_other_links,omitempty"`
		SortNodes      *bool                       `json:"sort_nodes,omitempty"`
		Time           *WidgetTime                 `json:"time,omitempty"`
		Title          *string                     `json:"title,omitempty"`
		TitleAlign     *WidgetTextAlign            `json:"title_align,omitempty"`
		TitleSize      *string                     `json:"title_size,omitempty"`
		Type           *SankeyWidgetDefinitionType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Requests == nil {
		return fmt.Errorf("required field requests missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.Requests = *all.Requests
	o.ShowOtherLinks = all.ShowOtherLinks
	o.SortNodes = all.SortNodes
	o.Time = all.Time
	o.Title = all.Title
	if all.TitleAlign != nil && !all.TitleAlign.IsValid() {
		hasInvalidField = true
	} else {
		o.TitleAlign = all.TitleAlign
	}
	o.TitleSize = all.TitleSize
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
