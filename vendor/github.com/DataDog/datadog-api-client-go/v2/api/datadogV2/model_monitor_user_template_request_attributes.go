// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorUserTemplateRequestAttributes Attributes for a monitor user template.
type MonitorUserTemplateRequestAttributes struct {
	// A brief description of the monitor user template.
	Description datadog.NullableString `json:"description,omitempty"`
	// A valid monitor definition in the same format as the [V1 Monitor API](https://docs.datadoghq.com/api/latest/monitors/#create-a-monitor).
	MonitorDefinition map[string]interface{} `json:"monitor_definition"`
	// The definition of `MonitorUserTemplateTags` object.
	Tags []string `json:"tags"`
	// The definition of `MonitorUserTemplateTemplateVariables` object.
	TemplateVariables []MonitorUserTemplateTemplateVariablesItems `json:"template_variables,omitempty"`
	// The title of the monitor user template.
	Title string `json:"title"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewMonitorUserTemplateRequestAttributes instantiates a new MonitorUserTemplateRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorUserTemplateRequestAttributes(monitorDefinition map[string]interface{}, tags []string, title string) *MonitorUserTemplateRequestAttributes {
	this := MonitorUserTemplateRequestAttributes{}
	this.MonitorDefinition = monitorDefinition
	this.Tags = tags
	this.Title = title
	return &this
}

// NewMonitorUserTemplateRequestAttributesWithDefaults instantiates a new MonitorUserTemplateRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorUserTemplateRequestAttributesWithDefaults() *MonitorUserTemplateRequestAttributes {
	this := MonitorUserTemplateRequestAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonitorUserTemplateRequestAttributes) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonitorUserTemplateRequestAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *MonitorUserTemplateRequestAttributes) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given datadog.NullableString and assigns it to the Description field.
func (o *MonitorUserTemplateRequestAttributes) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *MonitorUserTemplateRequestAttributes) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *MonitorUserTemplateRequestAttributes) UnsetDescription() {
	o.Description.Unset()
}

// GetMonitorDefinition returns the MonitorDefinition field value.
func (o *MonitorUserTemplateRequestAttributes) GetMonitorDefinition() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.MonitorDefinition
}

// GetMonitorDefinitionOk returns a tuple with the MonitorDefinition field value
// and a boolean to check if the value has been set.
func (o *MonitorUserTemplateRequestAttributes) GetMonitorDefinitionOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonitorDefinition, true
}

// SetMonitorDefinition sets field value.
func (o *MonitorUserTemplateRequestAttributes) SetMonitorDefinition(v map[string]interface{}) {
	o.MonitorDefinition = v
}

// GetTags returns the Tags field value.
func (o *MonitorUserTemplateRequestAttributes) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *MonitorUserTemplateRequestAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value.
func (o *MonitorUserTemplateRequestAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetTemplateVariables returns the TemplateVariables field value if set, zero value otherwise.
func (o *MonitorUserTemplateRequestAttributes) GetTemplateVariables() []MonitorUserTemplateTemplateVariablesItems {
	if o == nil || o.TemplateVariables == nil {
		var ret []MonitorUserTemplateTemplateVariablesItems
		return ret
	}
	return o.TemplateVariables
}

// GetTemplateVariablesOk returns a tuple with the TemplateVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorUserTemplateRequestAttributes) GetTemplateVariablesOk() (*[]MonitorUserTemplateTemplateVariablesItems, bool) {
	if o == nil || o.TemplateVariables == nil {
		return nil, false
	}
	return &o.TemplateVariables, true
}

// HasTemplateVariables returns a boolean if a field has been set.
func (o *MonitorUserTemplateRequestAttributes) HasTemplateVariables() bool {
	return o != nil && o.TemplateVariables != nil
}

// SetTemplateVariables gets a reference to the given []MonitorUserTemplateTemplateVariablesItems and assigns it to the TemplateVariables field.
func (o *MonitorUserTemplateRequestAttributes) SetTemplateVariables(v []MonitorUserTemplateTemplateVariablesItems) {
	o.TemplateVariables = v
}

// GetTitle returns the Title field value.
func (o *MonitorUserTemplateRequestAttributes) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *MonitorUserTemplateRequestAttributes) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *MonitorUserTemplateRequestAttributes) SetTitle(v string) {
	o.Title = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorUserTemplateRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["monitor_definition"] = o.MonitorDefinition
	toSerialize["tags"] = o.Tags
	if o.TemplateVariables != nil {
		toSerialize["template_variables"] = o.TemplateVariables
	}
	toSerialize["title"] = o.Title
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorUserTemplateRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description       datadog.NullableString                      `json:"description,omitempty"`
		MonitorDefinition *map[string]interface{}                     `json:"monitor_definition"`
		Tags              *[]string                                   `json:"tags"`
		TemplateVariables []MonitorUserTemplateTemplateVariablesItems `json:"template_variables,omitempty"`
		Title             *string                                     `json:"title"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.MonitorDefinition == nil {
		return fmt.Errorf("required field monitor_definition missing")
	}
	if all.Tags == nil {
		return fmt.Errorf("required field tags missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	o.Description = all.Description
	o.MonitorDefinition = *all.MonitorDefinition
	o.Tags = *all.Tags
	o.TemplateVariables = all.TemplateVariables
	o.Title = *all.Title

	return nil
}
