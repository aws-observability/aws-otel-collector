// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PowerpackWidgetDefinition The powerpack widget allows you to keep similar graphs together on your timeboard. Each group has a custom header, can hold one to many graphs, and is collapsible.
type PowerpackWidgetDefinition struct {
	// Background color of the powerpack title.
	BackgroundColor *string `json:"background_color,omitempty"`
	// URL of image to display as a banner for the powerpack.
	BannerImg *string `json:"banner_img,omitempty"`
	// UUID of the associated powerpack.
	PowerpackId string `json:"powerpack_id"`
	// Whether to show the title or not.
	ShowTitle *bool `json:"show_title,omitempty"`
	// Powerpack template variables.
	TemplateVariables *PowerpackTemplateVariables `json:"template_variables,omitempty"`
	// Title of the widget.
	Title *string `json:"title,omitempty"`
	// Type of the powerpack widget.
	Type PowerpackWidgetDefinitionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPowerpackWidgetDefinition instantiates a new PowerpackWidgetDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPowerpackWidgetDefinition(powerpackId string, typeVar PowerpackWidgetDefinitionType) *PowerpackWidgetDefinition {
	this := PowerpackWidgetDefinition{}
	this.PowerpackId = powerpackId
	var showTitle bool = true
	this.ShowTitle = &showTitle
	this.Type = typeVar
	return &this
}

// NewPowerpackWidgetDefinitionWithDefaults instantiates a new PowerpackWidgetDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPowerpackWidgetDefinitionWithDefaults() *PowerpackWidgetDefinition {
	this := PowerpackWidgetDefinition{}
	var showTitle bool = true
	this.ShowTitle = &showTitle
	var typeVar PowerpackWidgetDefinitionType = POWERPACKWIDGETDEFINITIONTYPE_POWERPACK
	this.Type = typeVar
	return &this
}

// GetBackgroundColor returns the BackgroundColor field value if set, zero value otherwise.
func (o *PowerpackWidgetDefinition) GetBackgroundColor() string {
	if o == nil || o.BackgroundColor == nil {
		var ret string
		return ret
	}
	return *o.BackgroundColor
}

// GetBackgroundColorOk returns a tuple with the BackgroundColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerpackWidgetDefinition) GetBackgroundColorOk() (*string, bool) {
	if o == nil || o.BackgroundColor == nil {
		return nil, false
	}
	return o.BackgroundColor, true
}

// HasBackgroundColor returns a boolean if a field has been set.
func (o *PowerpackWidgetDefinition) HasBackgroundColor() bool {
	return o != nil && o.BackgroundColor != nil
}

// SetBackgroundColor gets a reference to the given string and assigns it to the BackgroundColor field.
func (o *PowerpackWidgetDefinition) SetBackgroundColor(v string) {
	o.BackgroundColor = &v
}

// GetBannerImg returns the BannerImg field value if set, zero value otherwise.
func (o *PowerpackWidgetDefinition) GetBannerImg() string {
	if o == nil || o.BannerImg == nil {
		var ret string
		return ret
	}
	return *o.BannerImg
}

// GetBannerImgOk returns a tuple with the BannerImg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerpackWidgetDefinition) GetBannerImgOk() (*string, bool) {
	if o == nil || o.BannerImg == nil {
		return nil, false
	}
	return o.BannerImg, true
}

// HasBannerImg returns a boolean if a field has been set.
func (o *PowerpackWidgetDefinition) HasBannerImg() bool {
	return o != nil && o.BannerImg != nil
}

// SetBannerImg gets a reference to the given string and assigns it to the BannerImg field.
func (o *PowerpackWidgetDefinition) SetBannerImg(v string) {
	o.BannerImg = &v
}

// GetPowerpackId returns the PowerpackId field value.
func (o *PowerpackWidgetDefinition) GetPowerpackId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PowerpackId
}

// GetPowerpackIdOk returns a tuple with the PowerpackId field value
// and a boolean to check if the value has been set.
func (o *PowerpackWidgetDefinition) GetPowerpackIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PowerpackId, true
}

// SetPowerpackId sets field value.
func (o *PowerpackWidgetDefinition) SetPowerpackId(v string) {
	o.PowerpackId = v
}

// GetShowTitle returns the ShowTitle field value if set, zero value otherwise.
func (o *PowerpackWidgetDefinition) GetShowTitle() bool {
	if o == nil || o.ShowTitle == nil {
		var ret bool
		return ret
	}
	return *o.ShowTitle
}

// GetShowTitleOk returns a tuple with the ShowTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerpackWidgetDefinition) GetShowTitleOk() (*bool, bool) {
	if o == nil || o.ShowTitle == nil {
		return nil, false
	}
	return o.ShowTitle, true
}

// HasShowTitle returns a boolean if a field has been set.
func (o *PowerpackWidgetDefinition) HasShowTitle() bool {
	return o != nil && o.ShowTitle != nil
}

// SetShowTitle gets a reference to the given bool and assigns it to the ShowTitle field.
func (o *PowerpackWidgetDefinition) SetShowTitle(v bool) {
	o.ShowTitle = &v
}

// GetTemplateVariables returns the TemplateVariables field value if set, zero value otherwise.
func (o *PowerpackWidgetDefinition) GetTemplateVariables() PowerpackTemplateVariables {
	if o == nil || o.TemplateVariables == nil {
		var ret PowerpackTemplateVariables
		return ret
	}
	return *o.TemplateVariables
}

// GetTemplateVariablesOk returns a tuple with the TemplateVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerpackWidgetDefinition) GetTemplateVariablesOk() (*PowerpackTemplateVariables, bool) {
	if o == nil || o.TemplateVariables == nil {
		return nil, false
	}
	return o.TemplateVariables, true
}

// HasTemplateVariables returns a boolean if a field has been set.
func (o *PowerpackWidgetDefinition) HasTemplateVariables() bool {
	return o != nil && o.TemplateVariables != nil
}

// SetTemplateVariables gets a reference to the given PowerpackTemplateVariables and assigns it to the TemplateVariables field.
func (o *PowerpackWidgetDefinition) SetTemplateVariables(v PowerpackTemplateVariables) {
	o.TemplateVariables = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *PowerpackWidgetDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerpackWidgetDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *PowerpackWidgetDefinition) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *PowerpackWidgetDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value.
func (o *PowerpackWidgetDefinition) GetType() PowerpackWidgetDefinitionType {
	if o == nil {
		var ret PowerpackWidgetDefinitionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PowerpackWidgetDefinition) GetTypeOk() (*PowerpackWidgetDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *PowerpackWidgetDefinition) SetType(v PowerpackWidgetDefinitionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PowerpackWidgetDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.BackgroundColor != nil {
		toSerialize["background_color"] = o.BackgroundColor
	}
	if o.BannerImg != nil {
		toSerialize["banner_img"] = o.BannerImg
	}
	toSerialize["powerpack_id"] = o.PowerpackId
	if o.ShowTitle != nil {
		toSerialize["show_title"] = o.ShowTitle
	}
	if o.TemplateVariables != nil {
		toSerialize["template_variables"] = o.TemplateVariables
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PowerpackWidgetDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BackgroundColor   *string                        `json:"background_color,omitempty"`
		BannerImg         *string                        `json:"banner_img,omitempty"`
		PowerpackId       *string                        `json:"powerpack_id"`
		ShowTitle         *bool                          `json:"show_title,omitempty"`
		TemplateVariables *PowerpackTemplateVariables    `json:"template_variables,omitempty"`
		Title             *string                        `json:"title,omitempty"`
		Type              *PowerpackWidgetDefinitionType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.PowerpackId == nil {
		return fmt.Errorf("required field powerpack_id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"background_color", "banner_img", "powerpack_id", "show_title", "template_variables", "title", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.BackgroundColor = all.BackgroundColor
	o.BannerImg = all.BannerImg
	o.PowerpackId = *all.PowerpackId
	o.ShowTitle = all.ShowTitle
	if all.TemplateVariables != nil && all.TemplateVariables.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TemplateVariables = all.TemplateVariables
	o.Title = all.Title
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
