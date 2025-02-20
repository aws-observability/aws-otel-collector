// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// InputSchemaDataAttributesParametersItemsDataAttributes The definition of `InputSchemaDataAttributesParametersItemsDataAttributes` object.
type InputSchemaDataAttributesParametersItemsDataAttributes struct {
	// The `attributes` `defaultValue`.
	DefaultValue interface{} `json:"defaultValue,omitempty"`
	// The `attributes` `description`.
	Description *string `json:"description,omitempty"`
	// The `attributes` `enum`.
	Enum []string `json:"enum,omitempty"`
	// The `attributes` `label`.
	Label *string `json:"label,omitempty"`
	// The `attributes` `name`.
	Name *string `json:"name,omitempty"`
	// The `attributes` `type`.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInputSchemaDataAttributesParametersItemsDataAttributes instantiates a new InputSchemaDataAttributesParametersItemsDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInputSchemaDataAttributesParametersItemsDataAttributes() *InputSchemaDataAttributesParametersItemsDataAttributes {
	this := InputSchemaDataAttributesParametersItemsDataAttributes{}
	return &this
}

// NewInputSchemaDataAttributesParametersItemsDataAttributesWithDefaults instantiates a new InputSchemaDataAttributesParametersItemsDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInputSchemaDataAttributesParametersItemsDataAttributesWithDefaults() *InputSchemaDataAttributesParametersItemsDataAttributes {
	this := InputSchemaDataAttributesParametersItemsDataAttributes{}
	return &this
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *InputSchemaDataAttributesParametersItemsDataAttributes) GetDefaultValue() interface{} {
	if o == nil || o.DefaultValue == nil {
		var ret interface{}
		return ret
	}
	return o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputSchemaDataAttributesParametersItemsDataAttributes) GetDefaultValueOk() (*interface{}, bool) {
	if o == nil || o.DefaultValue == nil {
		return nil, false
	}
	return &o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *InputSchemaDataAttributesParametersItemsDataAttributes) HasDefaultValue() bool {
	return o != nil && o.DefaultValue != nil
}

// SetDefaultValue gets a reference to the given interface{} and assigns it to the DefaultValue field.
func (o *InputSchemaDataAttributesParametersItemsDataAttributes) SetDefaultValue(v interface{}) {
	o.DefaultValue = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InputSchemaDataAttributesParametersItemsDataAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputSchemaDataAttributesParametersItemsDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InputSchemaDataAttributesParametersItemsDataAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InputSchemaDataAttributesParametersItemsDataAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetEnum returns the Enum field value if set, zero value otherwise.
func (o *InputSchemaDataAttributesParametersItemsDataAttributes) GetEnum() []string {
	if o == nil || o.Enum == nil {
		var ret []string
		return ret
	}
	return o.Enum
}

// GetEnumOk returns a tuple with the Enum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputSchemaDataAttributesParametersItemsDataAttributes) GetEnumOk() (*[]string, bool) {
	if o == nil || o.Enum == nil {
		return nil, false
	}
	return &o.Enum, true
}

// HasEnum returns a boolean if a field has been set.
func (o *InputSchemaDataAttributesParametersItemsDataAttributes) HasEnum() bool {
	return o != nil && o.Enum != nil
}

// SetEnum gets a reference to the given []string and assigns it to the Enum field.
func (o *InputSchemaDataAttributesParametersItemsDataAttributes) SetEnum(v []string) {
	o.Enum = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *InputSchemaDataAttributesParametersItemsDataAttributes) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputSchemaDataAttributesParametersItemsDataAttributes) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *InputSchemaDataAttributesParametersItemsDataAttributes) HasLabel() bool {
	return o != nil && o.Label != nil
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *InputSchemaDataAttributesParametersItemsDataAttributes) SetLabel(v string) {
	o.Label = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InputSchemaDataAttributesParametersItemsDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputSchemaDataAttributesParametersItemsDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InputSchemaDataAttributesParametersItemsDataAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InputSchemaDataAttributesParametersItemsDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InputSchemaDataAttributesParametersItemsDataAttributes) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputSchemaDataAttributesParametersItemsDataAttributes) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InputSchemaDataAttributesParametersItemsDataAttributes) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InputSchemaDataAttributesParametersItemsDataAttributes) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o InputSchemaDataAttributesParametersItemsDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DefaultValue != nil {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Enum != nil {
		toSerialize["enum"] = o.Enum
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
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
func (o *InputSchemaDataAttributesParametersItemsDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DefaultValue interface{} `json:"defaultValue,omitempty"`
		Description  *string     `json:"description,omitempty"`
		Enum         []string    `json:"enum,omitempty"`
		Label        *string     `json:"label,omitempty"`
		Name         *string     `json:"name,omitempty"`
		Type         *string     `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"defaultValue", "description", "enum", "label", "name", "type"})
	} else {
		return err
	}
	o.DefaultValue = all.DefaultValue
	o.Description = all.Description
	o.Enum = all.Enum
	o.Label = all.Label
	o.Name = all.Name
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
