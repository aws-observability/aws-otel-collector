// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PowerpackAttributes Powerpack attribute object.
type PowerpackAttributes struct {
	// Description of this powerpack.
	Description *string `json:"description,omitempty"`
	// Powerpack group widget definition object.
	GroupWidget PowerpackGroupWidget `json:"group_widget"`
	// Name of the powerpack.
	Name string `json:"name"`
	// List of tags to identify this powerpack.
	Tags []string `json:"tags,omitempty"`
	// List of template variables for this powerpack.
	TemplateVariables []PowerpackTemplateVariable `json:"template_variables,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPowerpackAttributes instantiates a new PowerpackAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPowerpackAttributes(groupWidget PowerpackGroupWidget, name string) *PowerpackAttributes {
	this := PowerpackAttributes{}
	this.GroupWidget = groupWidget
	this.Name = name
	return &this
}

// NewPowerpackAttributesWithDefaults instantiates a new PowerpackAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPowerpackAttributesWithDefaults() *PowerpackAttributes {
	this := PowerpackAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PowerpackAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerpackAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PowerpackAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PowerpackAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetGroupWidget returns the GroupWidget field value.
func (o *PowerpackAttributes) GetGroupWidget() PowerpackGroupWidget {
	if o == nil {
		var ret PowerpackGroupWidget
		return ret
	}
	return o.GroupWidget
}

// GetGroupWidgetOk returns a tuple with the GroupWidget field value
// and a boolean to check if the value has been set.
func (o *PowerpackAttributes) GetGroupWidgetOk() (*PowerpackGroupWidget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupWidget, true
}

// SetGroupWidget sets field value.
func (o *PowerpackAttributes) SetGroupWidget(v PowerpackGroupWidget) {
	o.GroupWidget = v
}

// GetName returns the Name field value.
func (o *PowerpackAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PowerpackAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *PowerpackAttributes) SetName(v string) {
	o.Name = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PowerpackAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerpackAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PowerpackAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *PowerpackAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetTemplateVariables returns the TemplateVariables field value if set, zero value otherwise.
func (o *PowerpackAttributes) GetTemplateVariables() []PowerpackTemplateVariable {
	if o == nil || o.TemplateVariables == nil {
		var ret []PowerpackTemplateVariable
		return ret
	}
	return o.TemplateVariables
}

// GetTemplateVariablesOk returns a tuple with the TemplateVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerpackAttributes) GetTemplateVariablesOk() (*[]PowerpackTemplateVariable, bool) {
	if o == nil || o.TemplateVariables == nil {
		return nil, false
	}
	return &o.TemplateVariables, true
}

// HasTemplateVariables returns a boolean if a field has been set.
func (o *PowerpackAttributes) HasTemplateVariables() bool {
	return o != nil && o.TemplateVariables != nil
}

// SetTemplateVariables gets a reference to the given []PowerpackTemplateVariable and assigns it to the TemplateVariables field.
func (o *PowerpackAttributes) SetTemplateVariables(v []PowerpackTemplateVariable) {
	o.TemplateVariables = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PowerpackAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	toSerialize["group_widget"] = o.GroupWidget
	toSerialize["name"] = o.Name
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.TemplateVariables != nil {
		toSerialize["template_variables"] = o.TemplateVariables
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PowerpackAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description       *string                     `json:"description,omitempty"`
		GroupWidget       *PowerpackGroupWidget       `json:"group_widget"`
		Name              *string                     `json:"name"`
		Tags              []string                    `json:"tags,omitempty"`
		TemplateVariables []PowerpackTemplateVariable `json:"template_variables,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.GroupWidget == nil {
		return fmt.Errorf("required field group_widget missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "group_widget", "name", "tags", "template_variables"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = all.Description
	if all.GroupWidget.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.GroupWidget = *all.GroupWidget
	o.Name = *all.Name
	o.Tags = all.Tags
	o.TemplateVariables = all.TemplateVariables

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
