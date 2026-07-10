// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PointPlotWidgetDefinition The point plot displays individual data points over time.
type PointPlotWidgetDefinition struct {
	// List of custom links.
	CustomLinks []WidgetCustomLink `json:"custom_links,omitempty"`
	// The description of the widget.
	Description *string `json:"description,omitempty"`
	// Legend configuration for the point plot widget.
	Legend *PointPlotWidgetLegend `json:"legend,omitempty"`
	// List of markers for the widget.
	Markers []WidgetMarker `json:"markers,omitempty"`
	// List of request configurations for the widget.
	Requests []PointPlotWidgetRequest `json:"requests"`
	// Time setting for the widget.
	Time *WidgetTime `json:"time,omitempty"`
	// Title of the widget.
	Title *string `json:"title,omitempty"`
	// How to align the text on the widget.
	TitleAlign *WidgetTextAlign `json:"title_align,omitempty"`
	// Size of the title.
	TitleSize *string `json:"title_size,omitempty"`
	// Type of the point plot widget.
	Type PointPlotWidgetDefinitionType `json:"type"`
	// Axis controls for the widget.
	Yaxis *WidgetAxis `json:"yaxis,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPointPlotWidgetDefinition instantiates a new PointPlotWidgetDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPointPlotWidgetDefinition(requests []PointPlotWidgetRequest, typeVar PointPlotWidgetDefinitionType) *PointPlotWidgetDefinition {
	this := PointPlotWidgetDefinition{}
	this.Requests = requests
	this.Type = typeVar
	return &this
}

// NewPointPlotWidgetDefinitionWithDefaults instantiates a new PointPlotWidgetDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPointPlotWidgetDefinitionWithDefaults() *PointPlotWidgetDefinition {
	this := PointPlotWidgetDefinition{}
	var typeVar PointPlotWidgetDefinitionType = POINTPLOTWIDGETDEFINITIONTYPE_POINT_PLOT
	this.Type = typeVar
	return &this
}

// GetCustomLinks returns the CustomLinks field value if set, zero value otherwise.
func (o *PointPlotWidgetDefinition) GetCustomLinks() []WidgetCustomLink {
	if o == nil || o.CustomLinks == nil {
		var ret []WidgetCustomLink
		return ret
	}
	return o.CustomLinks
}

// GetCustomLinksOk returns a tuple with the CustomLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PointPlotWidgetDefinition) GetCustomLinksOk() (*[]WidgetCustomLink, bool) {
	if o == nil || o.CustomLinks == nil {
		return nil, false
	}
	return &o.CustomLinks, true
}

// HasCustomLinks returns a boolean if a field has been set.
func (o *PointPlotWidgetDefinition) HasCustomLinks() bool {
	return o != nil && o.CustomLinks != nil
}

// SetCustomLinks gets a reference to the given []WidgetCustomLink and assigns it to the CustomLinks field.
func (o *PointPlotWidgetDefinition) SetCustomLinks(v []WidgetCustomLink) {
	o.CustomLinks = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PointPlotWidgetDefinition) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PointPlotWidgetDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PointPlotWidgetDefinition) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PointPlotWidgetDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetLegend returns the Legend field value if set, zero value otherwise.
func (o *PointPlotWidgetDefinition) GetLegend() PointPlotWidgetLegend {
	if o == nil || o.Legend == nil {
		var ret PointPlotWidgetLegend
		return ret
	}
	return *o.Legend
}

// GetLegendOk returns a tuple with the Legend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PointPlotWidgetDefinition) GetLegendOk() (*PointPlotWidgetLegend, bool) {
	if o == nil || o.Legend == nil {
		return nil, false
	}
	return o.Legend, true
}

// HasLegend returns a boolean if a field has been set.
func (o *PointPlotWidgetDefinition) HasLegend() bool {
	return o != nil && o.Legend != nil
}

// SetLegend gets a reference to the given PointPlotWidgetLegend and assigns it to the Legend field.
func (o *PointPlotWidgetDefinition) SetLegend(v PointPlotWidgetLegend) {
	o.Legend = &v
}

// GetMarkers returns the Markers field value if set, zero value otherwise.
func (o *PointPlotWidgetDefinition) GetMarkers() []WidgetMarker {
	if o == nil || o.Markers == nil {
		var ret []WidgetMarker
		return ret
	}
	return o.Markers
}

// GetMarkersOk returns a tuple with the Markers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PointPlotWidgetDefinition) GetMarkersOk() (*[]WidgetMarker, bool) {
	if o == nil || o.Markers == nil {
		return nil, false
	}
	return &o.Markers, true
}

// HasMarkers returns a boolean if a field has been set.
func (o *PointPlotWidgetDefinition) HasMarkers() bool {
	return o != nil && o.Markers != nil
}

// SetMarkers gets a reference to the given []WidgetMarker and assigns it to the Markers field.
func (o *PointPlotWidgetDefinition) SetMarkers(v []WidgetMarker) {
	o.Markers = v
}

// GetRequests returns the Requests field value.
func (o *PointPlotWidgetDefinition) GetRequests() []PointPlotWidgetRequest {
	if o == nil {
		var ret []PointPlotWidgetRequest
		return ret
	}
	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *PointPlotWidgetDefinition) GetRequestsOk() (*[]PointPlotWidgetRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requests, true
}

// SetRequests sets field value.
func (o *PointPlotWidgetDefinition) SetRequests(v []PointPlotWidgetRequest) {
	o.Requests = v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *PointPlotWidgetDefinition) GetTime() WidgetTime {
	if o == nil || o.Time == nil {
		var ret WidgetTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PointPlotWidgetDefinition) GetTimeOk() (*WidgetTime, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *PointPlotWidgetDefinition) HasTime() bool {
	return o != nil && o.Time != nil
}

// SetTime gets a reference to the given WidgetTime and assigns it to the Time field.
func (o *PointPlotWidgetDefinition) SetTime(v WidgetTime) {
	o.Time = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *PointPlotWidgetDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PointPlotWidgetDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *PointPlotWidgetDefinition) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *PointPlotWidgetDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetTitleAlign returns the TitleAlign field value if set, zero value otherwise.
func (o *PointPlotWidgetDefinition) GetTitleAlign() WidgetTextAlign {
	if o == nil || o.TitleAlign == nil {
		var ret WidgetTextAlign
		return ret
	}
	return *o.TitleAlign
}

// GetTitleAlignOk returns a tuple with the TitleAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PointPlotWidgetDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool) {
	if o == nil || o.TitleAlign == nil {
		return nil, false
	}
	return o.TitleAlign, true
}

// HasTitleAlign returns a boolean if a field has been set.
func (o *PointPlotWidgetDefinition) HasTitleAlign() bool {
	return o != nil && o.TitleAlign != nil
}

// SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.
func (o *PointPlotWidgetDefinition) SetTitleAlign(v WidgetTextAlign) {
	o.TitleAlign = &v
}

// GetTitleSize returns the TitleSize field value if set, zero value otherwise.
func (o *PointPlotWidgetDefinition) GetTitleSize() string {
	if o == nil || o.TitleSize == nil {
		var ret string
		return ret
	}
	return *o.TitleSize
}

// GetTitleSizeOk returns a tuple with the TitleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PointPlotWidgetDefinition) GetTitleSizeOk() (*string, bool) {
	if o == nil || o.TitleSize == nil {
		return nil, false
	}
	return o.TitleSize, true
}

// HasTitleSize returns a boolean if a field has been set.
func (o *PointPlotWidgetDefinition) HasTitleSize() bool {
	return o != nil && o.TitleSize != nil
}

// SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.
func (o *PointPlotWidgetDefinition) SetTitleSize(v string) {
	o.TitleSize = &v
}

// GetType returns the Type field value.
func (o *PointPlotWidgetDefinition) GetType() PointPlotWidgetDefinitionType {
	if o == nil {
		var ret PointPlotWidgetDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PointPlotWidgetDefinition) GetTypeOk() (*PointPlotWidgetDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *PointPlotWidgetDefinition) SetType(v PointPlotWidgetDefinitionType) {
	o.Type = v
}

// GetYaxis returns the Yaxis field value if set, zero value otherwise.
func (o *PointPlotWidgetDefinition) GetYaxis() WidgetAxis {
	if o == nil || o.Yaxis == nil {
		var ret WidgetAxis
		return ret
	}
	return *o.Yaxis
}

// GetYaxisOk returns a tuple with the Yaxis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PointPlotWidgetDefinition) GetYaxisOk() (*WidgetAxis, bool) {
	if o == nil || o.Yaxis == nil {
		return nil, false
	}
	return o.Yaxis, true
}

// HasYaxis returns a boolean if a field has been set.
func (o *PointPlotWidgetDefinition) HasYaxis() bool {
	return o != nil && o.Yaxis != nil
}

// SetYaxis gets a reference to the given WidgetAxis and assigns it to the Yaxis field.
func (o *PointPlotWidgetDefinition) SetYaxis(v WidgetAxis) {
	o.Yaxis = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PointPlotWidgetDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CustomLinks != nil {
		toSerialize["custom_links"] = o.CustomLinks
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Legend != nil {
		toSerialize["legend"] = o.Legend
	}
	if o.Markers != nil {
		toSerialize["markers"] = o.Markers
	}
	toSerialize["requests"] = o.Requests
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
	if o.Yaxis != nil {
		toSerialize["yaxis"] = o.Yaxis
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PointPlotWidgetDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CustomLinks []WidgetCustomLink             `json:"custom_links,omitempty"`
		Description *string                        `json:"description,omitempty"`
		Legend      *PointPlotWidgetLegend         `json:"legend,omitempty"`
		Markers     []WidgetMarker                 `json:"markers,omitempty"`
		Requests    *[]PointPlotWidgetRequest      `json:"requests"`
		Time        *WidgetTime                    `json:"time,omitempty"`
		Title       *string                        `json:"title,omitempty"`
		TitleAlign  *WidgetTextAlign               `json:"title_align,omitempty"`
		TitleSize   *string                        `json:"title_size,omitempty"`
		Type        *PointPlotWidgetDefinitionType `json:"type"`
		Yaxis       *WidgetAxis                    `json:"yaxis,omitempty"`
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
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"custom_links", "description", "legend", "markers", "requests", "time", "title", "title_align", "title_size", "type", "yaxis"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CustomLinks = all.CustomLinks
	o.Description = all.Description
	if all.Legend != nil && all.Legend.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Legend = all.Legend
	o.Markers = all.Markers
	o.Requests = *all.Requests
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
	if all.Yaxis != nil && all.Yaxis.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Yaxis = all.Yaxis

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
