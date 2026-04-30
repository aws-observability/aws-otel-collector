// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsFunnelWidgetDefinition The user journey funnel visualization displays conversion funnels based on user journey data from Product Analytics.
type ProductAnalyticsFunnelWidgetDefinition struct {
	// The description of the widget.
	Description *string `json:"description,omitempty"`
	// Display mode for grouped funnel results.
	GroupedDisplay *FunnelGroupedDisplay `json:"grouped_display,omitempty"`
	// Request payload used to query items.
	Requests []ProductAnalyticsFunnelRequest `json:"requests"`
	// Time setting for the widget.
	Time *WidgetTime `json:"time,omitempty"`
	// The title of the widget.
	Title *string `json:"title,omitempty"`
	// How to align the text on the widget.
	TitleAlign *WidgetTextAlign `json:"title_align,omitempty"`
	// The size of the title.
	TitleSize *string `json:"title_size,omitempty"`
	// Type of funnel widget.
	Type FunnelWidgetDefinitionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewProductAnalyticsFunnelWidgetDefinition instantiates a new ProductAnalyticsFunnelWidgetDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsFunnelWidgetDefinition(requests []ProductAnalyticsFunnelRequest, typeVar FunnelWidgetDefinitionType) *ProductAnalyticsFunnelWidgetDefinition {
	this := ProductAnalyticsFunnelWidgetDefinition{}
	this.Requests = requests
	this.Type = typeVar
	return &this
}

// NewProductAnalyticsFunnelWidgetDefinitionWithDefaults instantiates a new ProductAnalyticsFunnelWidgetDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsFunnelWidgetDefinitionWithDefaults() *ProductAnalyticsFunnelWidgetDefinition {
	this := ProductAnalyticsFunnelWidgetDefinition{}
	var typeVar FunnelWidgetDefinitionType = FUNNELWIDGETDEFINITIONTYPE_FUNNEL
	this.Type = typeVar
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProductAnalyticsFunnelWidgetDefinition) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFunnelWidgetDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProductAnalyticsFunnelWidgetDefinition) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ProductAnalyticsFunnelWidgetDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetGroupedDisplay returns the GroupedDisplay field value if set, zero value otherwise.
func (o *ProductAnalyticsFunnelWidgetDefinition) GetGroupedDisplay() FunnelGroupedDisplay {
	if o == nil || o.GroupedDisplay == nil {
		var ret FunnelGroupedDisplay
		return ret
	}
	return *o.GroupedDisplay
}

// GetGroupedDisplayOk returns a tuple with the GroupedDisplay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFunnelWidgetDefinition) GetGroupedDisplayOk() (*FunnelGroupedDisplay, bool) {
	if o == nil || o.GroupedDisplay == nil {
		return nil, false
	}
	return o.GroupedDisplay, true
}

// HasGroupedDisplay returns a boolean if a field has been set.
func (o *ProductAnalyticsFunnelWidgetDefinition) HasGroupedDisplay() bool {
	return o != nil && o.GroupedDisplay != nil
}

// SetGroupedDisplay gets a reference to the given FunnelGroupedDisplay and assigns it to the GroupedDisplay field.
func (o *ProductAnalyticsFunnelWidgetDefinition) SetGroupedDisplay(v FunnelGroupedDisplay) {
	o.GroupedDisplay = &v
}

// GetRequests returns the Requests field value.
func (o *ProductAnalyticsFunnelWidgetDefinition) GetRequests() []ProductAnalyticsFunnelRequest {
	if o == nil {
		var ret []ProductAnalyticsFunnelRequest
		return ret
	}
	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFunnelWidgetDefinition) GetRequestsOk() (*[]ProductAnalyticsFunnelRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requests, true
}

// SetRequests sets field value.
func (o *ProductAnalyticsFunnelWidgetDefinition) SetRequests(v []ProductAnalyticsFunnelRequest) {
	o.Requests = v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *ProductAnalyticsFunnelWidgetDefinition) GetTime() WidgetTime {
	if o == nil || o.Time == nil {
		var ret WidgetTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFunnelWidgetDefinition) GetTimeOk() (*WidgetTime, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *ProductAnalyticsFunnelWidgetDefinition) HasTime() bool {
	return o != nil && o.Time != nil
}

// SetTime gets a reference to the given WidgetTime and assigns it to the Time field.
func (o *ProductAnalyticsFunnelWidgetDefinition) SetTime(v WidgetTime) {
	o.Time = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ProductAnalyticsFunnelWidgetDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFunnelWidgetDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ProductAnalyticsFunnelWidgetDefinition) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ProductAnalyticsFunnelWidgetDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetTitleAlign returns the TitleAlign field value if set, zero value otherwise.
func (o *ProductAnalyticsFunnelWidgetDefinition) GetTitleAlign() WidgetTextAlign {
	if o == nil || o.TitleAlign == nil {
		var ret WidgetTextAlign
		return ret
	}
	return *o.TitleAlign
}

// GetTitleAlignOk returns a tuple with the TitleAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFunnelWidgetDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool) {
	if o == nil || o.TitleAlign == nil {
		return nil, false
	}
	return o.TitleAlign, true
}

// HasTitleAlign returns a boolean if a field has been set.
func (o *ProductAnalyticsFunnelWidgetDefinition) HasTitleAlign() bool {
	return o != nil && o.TitleAlign != nil
}

// SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.
func (o *ProductAnalyticsFunnelWidgetDefinition) SetTitleAlign(v WidgetTextAlign) {
	o.TitleAlign = &v
}

// GetTitleSize returns the TitleSize field value if set, zero value otherwise.
func (o *ProductAnalyticsFunnelWidgetDefinition) GetTitleSize() string {
	if o == nil || o.TitleSize == nil {
		var ret string
		return ret
	}
	return *o.TitleSize
}

// GetTitleSizeOk returns a tuple with the TitleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFunnelWidgetDefinition) GetTitleSizeOk() (*string, bool) {
	if o == nil || o.TitleSize == nil {
		return nil, false
	}
	return o.TitleSize, true
}

// HasTitleSize returns a boolean if a field has been set.
func (o *ProductAnalyticsFunnelWidgetDefinition) HasTitleSize() bool {
	return o != nil && o.TitleSize != nil
}

// SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.
func (o *ProductAnalyticsFunnelWidgetDefinition) SetTitleSize(v string) {
	o.TitleSize = &v
}

// GetType returns the Type field value.
func (o *ProductAnalyticsFunnelWidgetDefinition) GetType() FunnelWidgetDefinitionType {
	if o == nil {
		var ret FunnelWidgetDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFunnelWidgetDefinition) GetTypeOk() (*FunnelWidgetDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ProductAnalyticsFunnelWidgetDefinition) SetType(v FunnelWidgetDefinitionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsFunnelWidgetDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.GroupedDisplay != nil {
		toSerialize["grouped_display"] = o.GroupedDisplay
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
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsFunnelWidgetDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description    *string                          `json:"description,omitempty"`
		GroupedDisplay *FunnelGroupedDisplay            `json:"grouped_display,omitempty"`
		Requests       *[]ProductAnalyticsFunnelRequest `json:"requests"`
		Time           *WidgetTime                      `json:"time,omitempty"`
		Title          *string                          `json:"title,omitempty"`
		TitleAlign     *WidgetTextAlign                 `json:"title_align,omitempty"`
		TitleSize      *string                          `json:"title_size,omitempty"`
		Type           *FunnelWidgetDefinitionType      `json:"type"`
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
	o.Description = all.Description
	if all.GroupedDisplay != nil && !all.GroupedDisplay.IsValid() {
		hasInvalidField = true
	} else {
		o.GroupedDisplay = all.GroupedDisplay
	}
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

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
