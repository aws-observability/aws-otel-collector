// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AlertGraphWidgetDefinition Alert graphs are timeseries graphs showing the current status of any monitor defined on your system.
type AlertGraphWidgetDefinition struct {
	// ID of the alert to use in the widget.
	AlertId string `json:"alert_id"`
	// Time setting for the widget.
	Time *WidgetTime `json:"time,omitempty"`
	// The title of the widget.
	Title *string `json:"title,omitempty"`
	// How to align the text on the widget.
	TitleAlign *WidgetTextAlign `json:"title_align,omitempty"`
	// Size of the title.
	TitleSize *string `json:"title_size,omitempty"`
	// Type of the alert graph widget.
	Type AlertGraphWidgetDefinitionType `json:"type"`
	// Whether to display the Alert Graph as a timeseries or a top list.
	VizType WidgetVizType `json:"viz_type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAlertGraphWidgetDefinition instantiates a new AlertGraphWidgetDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAlertGraphWidgetDefinition(alertId string, typeVar AlertGraphWidgetDefinitionType, vizType WidgetVizType) *AlertGraphWidgetDefinition {
	this := AlertGraphWidgetDefinition{}
	this.AlertId = alertId
	this.Type = typeVar
	this.VizType = vizType
	return &this
}

// NewAlertGraphWidgetDefinitionWithDefaults instantiates a new AlertGraphWidgetDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAlertGraphWidgetDefinitionWithDefaults() *AlertGraphWidgetDefinition {
	this := AlertGraphWidgetDefinition{}
	var typeVar AlertGraphWidgetDefinitionType = ALERTGRAPHWIDGETDEFINITIONTYPE_ALERT_GRAPH
	this.Type = typeVar
	return &this
}

// GetAlertId returns the AlertId field value.
func (o *AlertGraphWidgetDefinition) GetAlertId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AlertId
}

// GetAlertIdOk returns a tuple with the AlertId field value
// and a boolean to check if the value has been set.
func (o *AlertGraphWidgetDefinition) GetAlertIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlertId, true
}

// SetAlertId sets field value.
func (o *AlertGraphWidgetDefinition) SetAlertId(v string) {
	o.AlertId = v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *AlertGraphWidgetDefinition) GetTime() WidgetTime {
	if o == nil || o.Time == nil {
		var ret WidgetTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertGraphWidgetDefinition) GetTimeOk() (*WidgetTime, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *AlertGraphWidgetDefinition) HasTime() bool {
	return o != nil && o.Time != nil
}

// SetTime gets a reference to the given WidgetTime and assigns it to the Time field.
func (o *AlertGraphWidgetDefinition) SetTime(v WidgetTime) {
	o.Time = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AlertGraphWidgetDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertGraphWidgetDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AlertGraphWidgetDefinition) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AlertGraphWidgetDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetTitleAlign returns the TitleAlign field value if set, zero value otherwise.
func (o *AlertGraphWidgetDefinition) GetTitleAlign() WidgetTextAlign {
	if o == nil || o.TitleAlign == nil {
		var ret WidgetTextAlign
		return ret
	}
	return *o.TitleAlign
}

// GetTitleAlignOk returns a tuple with the TitleAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertGraphWidgetDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool) {
	if o == nil || o.TitleAlign == nil {
		return nil, false
	}
	return o.TitleAlign, true
}

// HasTitleAlign returns a boolean if a field has been set.
func (o *AlertGraphWidgetDefinition) HasTitleAlign() bool {
	return o != nil && o.TitleAlign != nil
}

// SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.
func (o *AlertGraphWidgetDefinition) SetTitleAlign(v WidgetTextAlign) {
	o.TitleAlign = &v
}

// GetTitleSize returns the TitleSize field value if set, zero value otherwise.
func (o *AlertGraphWidgetDefinition) GetTitleSize() string {
	if o == nil || o.TitleSize == nil {
		var ret string
		return ret
	}
	return *o.TitleSize
}

// GetTitleSizeOk returns a tuple with the TitleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertGraphWidgetDefinition) GetTitleSizeOk() (*string, bool) {
	if o == nil || o.TitleSize == nil {
		return nil, false
	}
	return o.TitleSize, true
}

// HasTitleSize returns a boolean if a field has been set.
func (o *AlertGraphWidgetDefinition) HasTitleSize() bool {
	return o != nil && o.TitleSize != nil
}

// SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.
func (o *AlertGraphWidgetDefinition) SetTitleSize(v string) {
	o.TitleSize = &v
}

// GetType returns the Type field value.
func (o *AlertGraphWidgetDefinition) GetType() AlertGraphWidgetDefinitionType {
	if o == nil {
		var ret AlertGraphWidgetDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AlertGraphWidgetDefinition) GetTypeOk() (*AlertGraphWidgetDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AlertGraphWidgetDefinition) SetType(v AlertGraphWidgetDefinitionType) {
	o.Type = v
}

// GetVizType returns the VizType field value.
func (o *AlertGraphWidgetDefinition) GetVizType() WidgetVizType {
	if o == nil {
		var ret WidgetVizType
		return ret
	}
	return o.VizType
}

// GetVizTypeOk returns a tuple with the VizType field value
// and a boolean to check if the value has been set.
func (o *AlertGraphWidgetDefinition) GetVizTypeOk() (*WidgetVizType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VizType, true
}

// SetVizType sets field value.
func (o *AlertGraphWidgetDefinition) SetVizType(v WidgetVizType) {
	o.VizType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AlertGraphWidgetDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["alert_id"] = o.AlertId
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
	toSerialize["viz_type"] = o.VizType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AlertGraphWidgetDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AlertId    *string                         `json:"alert_id"`
		Time       *WidgetTime                     `json:"time,omitempty"`
		Title      *string                         `json:"title,omitempty"`
		TitleAlign *WidgetTextAlign                `json:"title_align,omitempty"`
		TitleSize  *string                         `json:"title_size,omitempty"`
		Type       *AlertGraphWidgetDefinitionType `json:"type"`
		VizType    *WidgetVizType                  `json:"viz_type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AlertId == nil {
		return fmt.Errorf("required field alert_id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.VizType == nil {
		return fmt.Errorf("required field viz_type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"alert_id", "time", "title", "title_align", "title_size", "type", "viz_type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AlertId = *all.AlertId
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
	if !all.VizType.IsValid() {
		hasInvalidField = true
	} else {
		o.VizType = *all.VizType
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
