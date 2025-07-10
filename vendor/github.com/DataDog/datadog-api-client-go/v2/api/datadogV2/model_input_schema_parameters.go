// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// InputSchemaParameters The definition of `InputSchemaParameters` object.
type InputSchemaParameters struct {
	// The `InputSchemaParameters` `defaultValue`.
	DefaultValue interface{} `json:"defaultValue,omitempty"`
	// The `InputSchemaParameters` `description`.
	Description *string `json:"description,omitempty"`
	// The `InputSchemaParameters` `label`.
	Label *string `json:"label,omitempty"`
	// The `InputSchemaParameters` `name`.
	Name string `json:"name"`
	// The definition of `InputSchemaParametersType` object.
	Type InputSchemaParametersType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInputSchemaParameters instantiates a new InputSchemaParameters object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInputSchemaParameters(name string, typeVar InputSchemaParametersType) *InputSchemaParameters {
	this := InputSchemaParameters{}
	this.Name = name
	this.Type = typeVar
	return &this
}

// NewInputSchemaParametersWithDefaults instantiates a new InputSchemaParameters object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInputSchemaParametersWithDefaults() *InputSchemaParameters {
	this := InputSchemaParameters{}
	return &this
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *InputSchemaParameters) GetDefaultValue() interface{} {
	if o == nil || o.DefaultValue == nil {
		var ret interface{}
		return ret
	}
	return o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputSchemaParameters) GetDefaultValueOk() (*interface{}, bool) {
	if o == nil || o.DefaultValue == nil {
		return nil, false
	}
	return &o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *InputSchemaParameters) HasDefaultValue() bool {
	return o != nil && o.DefaultValue != nil
}

// SetDefaultValue gets a reference to the given interface{} and assigns it to the DefaultValue field.
func (o *InputSchemaParameters) SetDefaultValue(v interface{}) {
	o.DefaultValue = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InputSchemaParameters) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputSchemaParameters) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InputSchemaParameters) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InputSchemaParameters) SetDescription(v string) {
	o.Description = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *InputSchemaParameters) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputSchemaParameters) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *InputSchemaParameters) HasLabel() bool {
	return o != nil && o.Label != nil
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *InputSchemaParameters) SetLabel(v string) {
	o.Label = &v
}

// GetName returns the Name field value.
func (o *InputSchemaParameters) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InputSchemaParameters) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *InputSchemaParameters) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value.
func (o *InputSchemaParameters) GetType() InputSchemaParametersType {
	if o == nil {
		var ret InputSchemaParametersType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InputSchemaParameters) GetTypeOk() (*InputSchemaParametersType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *InputSchemaParameters) SetType(v InputSchemaParametersType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o InputSchemaParameters) MarshalJSON() ([]byte, error) {
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
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InputSchemaParameters) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DefaultValue interface{}                `json:"defaultValue,omitempty"`
		Description  *string                    `json:"description,omitempty"`
		Label        *string                    `json:"label,omitempty"`
		Name         *string                    `json:"name"`
		Type         *InputSchemaParametersType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"defaultValue", "description", "label", "name", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DefaultValue = all.DefaultValue
	o.Description = all.Description
	o.Label = all.Label
	o.Name = *all.Name
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
