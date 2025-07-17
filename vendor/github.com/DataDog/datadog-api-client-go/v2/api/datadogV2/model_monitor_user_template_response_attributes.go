// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorUserTemplateResponseAttributes Attributes for a monitor user template.
type MonitorUserTemplateResponseAttributes struct {
	// The created timestamp of the template.
	Created *time.Time `json:"created,omitempty"`
	// A brief description of the monitor user template.
	Description datadog.NullableString `json:"description,omitempty"`
	// The last modified timestamp. When the template version was created.
	Modified *time.Time `json:"modified,omitempty"`
	// A valid monitor definition in the same format as the [V1 Monitor API](https://docs.datadoghq.com/api/latest/monitors/#create-a-monitor).
	MonitorDefinition map[string]interface{} `json:"monitor_definition,omitempty"`
	// The definition of `MonitorUserTemplateTags` object.
	Tags []string `json:"tags,omitempty"`
	// The definition of `MonitorUserTemplateTemplateVariables` object.
	TemplateVariables []MonitorUserTemplateTemplateVariablesItems `json:"template_variables,omitempty"`
	// The title of the monitor user template.
	Title *string `json:"title,omitempty"`
	// The version of the monitor user template.
	Version datadog.NullableInt64 `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMonitorUserTemplateResponseAttributes instantiates a new MonitorUserTemplateResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorUserTemplateResponseAttributes() *MonitorUserTemplateResponseAttributes {
	this := MonitorUserTemplateResponseAttributes{}
	return &this
}

// NewMonitorUserTemplateResponseAttributesWithDefaults instantiates a new MonitorUserTemplateResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorUserTemplateResponseAttributesWithDefaults() *MonitorUserTemplateResponseAttributes {
	this := MonitorUserTemplateResponseAttributes{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *MonitorUserTemplateResponseAttributes) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorUserTemplateResponseAttributes) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *MonitorUserTemplateResponseAttributes) HasCreated() bool {
	return o != nil && o.Created != nil
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *MonitorUserTemplateResponseAttributes) SetCreated(v time.Time) {
	o.Created = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonitorUserTemplateResponseAttributes) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonitorUserTemplateResponseAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *MonitorUserTemplateResponseAttributes) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given datadog.NullableString and assigns it to the Description field.
func (o *MonitorUserTemplateResponseAttributes) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *MonitorUserTemplateResponseAttributes) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *MonitorUserTemplateResponseAttributes) UnsetDescription() {
	o.Description.Unset()
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *MonitorUserTemplateResponseAttributes) GetModified() time.Time {
	if o == nil || o.Modified == nil {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorUserTemplateResponseAttributes) GetModifiedOk() (*time.Time, bool) {
	if o == nil || o.Modified == nil {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *MonitorUserTemplateResponseAttributes) HasModified() bool {
	return o != nil && o.Modified != nil
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *MonitorUserTemplateResponseAttributes) SetModified(v time.Time) {
	o.Modified = &v
}

// GetMonitorDefinition returns the MonitorDefinition field value if set, zero value otherwise.
func (o *MonitorUserTemplateResponseAttributes) GetMonitorDefinition() map[string]interface{} {
	if o == nil || o.MonitorDefinition == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.MonitorDefinition
}

// GetMonitorDefinitionOk returns a tuple with the MonitorDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorUserTemplateResponseAttributes) GetMonitorDefinitionOk() (*map[string]interface{}, bool) {
	if o == nil || o.MonitorDefinition == nil {
		return nil, false
	}
	return &o.MonitorDefinition, true
}

// HasMonitorDefinition returns a boolean if a field has been set.
func (o *MonitorUserTemplateResponseAttributes) HasMonitorDefinition() bool {
	return o != nil && o.MonitorDefinition != nil
}

// SetMonitorDefinition gets a reference to the given map[string]interface{} and assigns it to the MonitorDefinition field.
func (o *MonitorUserTemplateResponseAttributes) SetMonitorDefinition(v map[string]interface{}) {
	o.MonitorDefinition = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *MonitorUserTemplateResponseAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorUserTemplateResponseAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *MonitorUserTemplateResponseAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *MonitorUserTemplateResponseAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetTemplateVariables returns the TemplateVariables field value if set, zero value otherwise.
func (o *MonitorUserTemplateResponseAttributes) GetTemplateVariables() []MonitorUserTemplateTemplateVariablesItems {
	if o == nil || o.TemplateVariables == nil {
		var ret []MonitorUserTemplateTemplateVariablesItems
		return ret
	}
	return o.TemplateVariables
}

// GetTemplateVariablesOk returns a tuple with the TemplateVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorUserTemplateResponseAttributes) GetTemplateVariablesOk() (*[]MonitorUserTemplateTemplateVariablesItems, bool) {
	if o == nil || o.TemplateVariables == nil {
		return nil, false
	}
	return &o.TemplateVariables, true
}

// HasTemplateVariables returns a boolean if a field has been set.
func (o *MonitorUserTemplateResponseAttributes) HasTemplateVariables() bool {
	return o != nil && o.TemplateVariables != nil
}

// SetTemplateVariables gets a reference to the given []MonitorUserTemplateTemplateVariablesItems and assigns it to the TemplateVariables field.
func (o *MonitorUserTemplateResponseAttributes) SetTemplateVariables(v []MonitorUserTemplateTemplateVariablesItems) {
	o.TemplateVariables = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *MonitorUserTemplateResponseAttributes) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorUserTemplateResponseAttributes) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *MonitorUserTemplateResponseAttributes) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *MonitorUserTemplateResponseAttributes) SetTitle(v string) {
	o.Title = &v
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonitorUserTemplateResponseAttributes) GetVersion() int64 {
	if o == nil || o.Version.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonitorUserTemplateResponseAttributes) GetVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *MonitorUserTemplateResponseAttributes) HasVersion() bool {
	return o != nil && o.Version.IsSet()
}

// SetVersion gets a reference to the given datadog.NullableInt64 and assigns it to the Version field.
func (o *MonitorUserTemplateResponseAttributes) SetVersion(v int64) {
	o.Version.Set(&v)
}

// SetVersionNil sets the value for Version to be an explicit nil.
func (o *MonitorUserTemplateResponseAttributes) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil.
func (o *MonitorUserTemplateResponseAttributes) UnsetVersion() {
	o.Version.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorUserTemplateResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Created != nil {
		if o.Created.Nanosecond() == 0 {
			toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Modified != nil {
		if o.Modified.Nanosecond() == 0 {
			toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.MonitorDefinition != nil {
		toSerialize["monitor_definition"] = o.MonitorDefinition
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.TemplateVariables != nil {
		toSerialize["template_variables"] = o.TemplateVariables
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorUserTemplateResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Created           *time.Time                                  `json:"created,omitempty"`
		Description       datadog.NullableString                      `json:"description,omitempty"`
		Modified          *time.Time                                  `json:"modified,omitempty"`
		MonitorDefinition map[string]interface{}                      `json:"monitor_definition,omitempty"`
		Tags              []string                                    `json:"tags,omitempty"`
		TemplateVariables []MonitorUserTemplateTemplateVariablesItems `json:"template_variables,omitempty"`
		Title             *string                                     `json:"title,omitempty"`
		Version           datadog.NullableInt64                       `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created", "description", "modified", "monitor_definition", "tags", "template_variables", "title", "version"})
	} else {
		return err
	}
	o.Created = all.Created
	o.Description = all.Description
	o.Modified = all.Modified
	o.MonitorDefinition = all.MonitorDefinition
	o.Tags = all.Tags
	o.TemplateVariables = all.TemplateVariables
	o.Title = all.Title
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
