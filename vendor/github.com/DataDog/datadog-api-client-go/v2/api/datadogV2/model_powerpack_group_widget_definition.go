// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PowerpackGroupWidgetDefinition Powerpack group widget object.
type PowerpackGroupWidgetDefinition struct {
	// Layout type of widgets.
	LayoutType string `json:"layout_type"`
	// Boolean indicating whether powerpack group title should be visible or not.
	ShowTitle *bool `json:"show_title,omitempty"`
	// Name for the group widget.
	Title *string `json:"title,omitempty"`
	// Type of widget, must be group.
	Type string `json:"type"`
	// Widgets inside the powerpack.
	Widgets []PowerpackInnerWidgets `json:"widgets"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPowerpackGroupWidgetDefinition instantiates a new PowerpackGroupWidgetDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPowerpackGroupWidgetDefinition(layoutType string, typeVar string, widgets []PowerpackInnerWidgets) *PowerpackGroupWidgetDefinition {
	this := PowerpackGroupWidgetDefinition{}
	this.LayoutType = layoutType
	this.Type = typeVar
	this.Widgets = widgets
	return &this
}

// NewPowerpackGroupWidgetDefinitionWithDefaults instantiates a new PowerpackGroupWidgetDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPowerpackGroupWidgetDefinitionWithDefaults() *PowerpackGroupWidgetDefinition {
	this := PowerpackGroupWidgetDefinition{}
	return &this
}

// GetLayoutType returns the LayoutType field value.
func (o *PowerpackGroupWidgetDefinition) GetLayoutType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.LayoutType
}

// GetLayoutTypeOk returns a tuple with the LayoutType field value
// and a boolean to check if the value has been set.
func (o *PowerpackGroupWidgetDefinition) GetLayoutTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LayoutType, true
}

// SetLayoutType sets field value.
func (o *PowerpackGroupWidgetDefinition) SetLayoutType(v string) {
	o.LayoutType = v
}

// GetShowTitle returns the ShowTitle field value if set, zero value otherwise.
func (o *PowerpackGroupWidgetDefinition) GetShowTitle() bool {
	if o == nil || o.ShowTitle == nil {
		var ret bool
		return ret
	}
	return *o.ShowTitle
}

// GetShowTitleOk returns a tuple with the ShowTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerpackGroupWidgetDefinition) GetShowTitleOk() (*bool, bool) {
	if o == nil || o.ShowTitle == nil {
		return nil, false
	}
	return o.ShowTitle, true
}

// HasShowTitle returns a boolean if a field has been set.
func (o *PowerpackGroupWidgetDefinition) HasShowTitle() bool {
	return o != nil && o.ShowTitle != nil
}

// SetShowTitle gets a reference to the given bool and assigns it to the ShowTitle field.
func (o *PowerpackGroupWidgetDefinition) SetShowTitle(v bool) {
	o.ShowTitle = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *PowerpackGroupWidgetDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerpackGroupWidgetDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *PowerpackGroupWidgetDefinition) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *PowerpackGroupWidgetDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value.
func (o *PowerpackGroupWidgetDefinition) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PowerpackGroupWidgetDefinition) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *PowerpackGroupWidgetDefinition) SetType(v string) {
	o.Type = v
}

// GetWidgets returns the Widgets field value.
func (o *PowerpackGroupWidgetDefinition) GetWidgets() []PowerpackInnerWidgets {
	if o == nil {
		var ret []PowerpackInnerWidgets
		return ret
	}
	return o.Widgets
}

// GetWidgetsOk returns a tuple with the Widgets field value
// and a boolean to check if the value has been set.
func (o *PowerpackGroupWidgetDefinition) GetWidgetsOk() (*[]PowerpackInnerWidgets, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Widgets, true
}

// SetWidgets sets field value.
func (o *PowerpackGroupWidgetDefinition) SetWidgets(v []PowerpackInnerWidgets) {
	o.Widgets = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PowerpackGroupWidgetDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["layout_type"] = o.LayoutType
	if o.ShowTitle != nil {
		toSerialize["show_title"] = o.ShowTitle
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	toSerialize["type"] = o.Type
	toSerialize["widgets"] = o.Widgets

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PowerpackGroupWidgetDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		LayoutType *string                  `json:"layout_type"`
		ShowTitle  *bool                    `json:"show_title,omitempty"`
		Title      *string                  `json:"title,omitempty"`
		Type       *string                  `json:"type"`
		Widgets    *[]PowerpackInnerWidgets `json:"widgets"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.LayoutType == nil {
		return fmt.Errorf("required field layout_type missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Widgets == nil {
		return fmt.Errorf("required field widgets missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"layout_type", "show_title", "title", "type", "widgets"})
	} else {
		return err
	}
	o.LayoutType = *all.LayoutType
	o.ShowTitle = all.ShowTitle
	o.Title = all.Title
	o.Type = *all.Type
	o.Widgets = *all.Widgets

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
