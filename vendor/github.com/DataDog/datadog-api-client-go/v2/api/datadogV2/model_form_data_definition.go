// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormDataDefinition A JSON Schema definition that describes the form's data fields.
type FormDataDefinition struct {
	// A description shown to form respondents.
	Description *string `json:"description,omitempty"`
	// A map of field names to their JSON Schema definitions.
	Properties map[string]interface{} `json:"properties,omitempty"`
	// List of field names that must be answered.
	Required []string `json:"required,omitempty"`
	// The title of the form schema.
	Title *string `json:"title,omitempty"`
	// The root schema type.
	Type *FormDataDefinitionType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFormDataDefinition instantiates a new FormDataDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormDataDefinition() *FormDataDefinition {
	this := FormDataDefinition{}
	var typeVar FormDataDefinitionType = FORMDATADEFINITIONTYPE_OBJECT
	this.Type = &typeVar
	return &this
}

// NewFormDataDefinitionWithDefaults instantiates a new FormDataDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormDataDefinitionWithDefaults() *FormDataDefinition {
	this := FormDataDefinition{}
	var typeVar FormDataDefinitionType = FORMDATADEFINITIONTYPE_OBJECT
	this.Type = &typeVar
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FormDataDefinition) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FormDataDefinition) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FormDataDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *FormDataDefinition) GetProperties() map[string]interface{} {
	if o == nil || o.Properties == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataDefinition) GetPropertiesOk() (*map[string]interface{}, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return &o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *FormDataDefinition) HasProperties() bool {
	return o != nil && o.Properties != nil
}

// SetProperties gets a reference to the given map[string]interface{} and assigns it to the Properties field.
func (o *FormDataDefinition) SetProperties(v map[string]interface{}) {
	o.Properties = v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *FormDataDefinition) GetRequired() []string {
	if o == nil || o.Required == nil {
		var ret []string
		return ret
	}
	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataDefinition) GetRequiredOk() (*[]string, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return &o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *FormDataDefinition) HasRequired() bool {
	return o != nil && o.Required != nil
}

// SetRequired gets a reference to the given []string and assigns it to the Required field.
func (o *FormDataDefinition) SetRequired(v []string) {
	o.Required = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *FormDataDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *FormDataDefinition) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *FormDataDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FormDataDefinition) GetType() FormDataDefinitionType {
	if o == nil || o.Type == nil {
		var ret FormDataDefinitionType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataDefinition) GetTypeOk() (*FormDataDefinitionType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FormDataDefinition) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given FormDataDefinitionType and assigns it to the Type field.
func (o *FormDataDefinition) SetType(v FormDataDefinitionType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormDataDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormDataDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string                 `json:"description,omitempty"`
		Properties  map[string]interface{}  `json:"properties,omitempty"`
		Required    []string                `json:"required,omitempty"`
		Title       *string                 `json:"title,omitempty"`
		Type        *FormDataDefinitionType `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "properties", "required", "title", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = all.Description
	o.Properties = all.Properties
	o.Required = all.Required
	o.Title = all.Title
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
