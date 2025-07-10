// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OutputSchemaParameters The definition of `OutputSchemaParameters` object.
type OutputSchemaParameters struct {
	// The `OutputSchemaParameters` `defaultValue`.
	DefaultValue interface{} `json:"defaultValue,omitempty"`
	// The `OutputSchemaParameters` `description`.
	Description *string `json:"description,omitempty"`
	// The `OutputSchemaParameters` `label`.
	Label *string `json:"label,omitempty"`
	// The `OutputSchemaParameters` `name`.
	Name string `json:"name"`
	// The definition of `OutputSchemaParametersType` object.
	Type OutputSchemaParametersType `json:"type"`
	// The `OutputSchemaParameters` `value`.
	Value interface{} `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOutputSchemaParameters instantiates a new OutputSchemaParameters object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOutputSchemaParameters(name string, typeVar OutputSchemaParametersType) *OutputSchemaParameters {
	this := OutputSchemaParameters{}
	this.Name = name
	this.Type = typeVar
	return &this
}

// NewOutputSchemaParametersWithDefaults instantiates a new OutputSchemaParameters object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOutputSchemaParametersWithDefaults() *OutputSchemaParameters {
	this := OutputSchemaParameters{}
	return &this
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *OutputSchemaParameters) GetDefaultValue() interface{} {
	if o == nil || o.DefaultValue == nil {
		var ret interface{}
		return ret
	}
	return o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputSchemaParameters) GetDefaultValueOk() (*interface{}, bool) {
	if o == nil || o.DefaultValue == nil {
		return nil, false
	}
	return &o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *OutputSchemaParameters) HasDefaultValue() bool {
	return o != nil && o.DefaultValue != nil
}

// SetDefaultValue gets a reference to the given interface{} and assigns it to the DefaultValue field.
func (o *OutputSchemaParameters) SetDefaultValue(v interface{}) {
	o.DefaultValue = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OutputSchemaParameters) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputSchemaParameters) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OutputSchemaParameters) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OutputSchemaParameters) SetDescription(v string) {
	o.Description = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *OutputSchemaParameters) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputSchemaParameters) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *OutputSchemaParameters) HasLabel() bool {
	return o != nil && o.Label != nil
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *OutputSchemaParameters) SetLabel(v string) {
	o.Label = &v
}

// GetName returns the Name field value.
func (o *OutputSchemaParameters) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OutputSchemaParameters) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *OutputSchemaParameters) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value.
func (o *OutputSchemaParameters) GetType() OutputSchemaParametersType {
	if o == nil {
		var ret OutputSchemaParametersType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OutputSchemaParameters) GetTypeOk() (*OutputSchemaParametersType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *OutputSchemaParameters) SetType(v OutputSchemaParametersType) {
	o.Type = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *OutputSchemaParameters) GetValue() interface{} {
	if o == nil || o.Value == nil {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputSchemaParameters) GetValueOk() (*interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *OutputSchemaParameters) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *OutputSchemaParameters) SetValue(v interface{}) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OutputSchemaParameters) MarshalJSON() ([]byte, error) {
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
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OutputSchemaParameters) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DefaultValue interface{}                 `json:"defaultValue,omitempty"`
		Description  *string                     `json:"description,omitempty"`
		Label        *string                     `json:"label,omitempty"`
		Name         *string                     `json:"name"`
		Type         *OutputSchemaParametersType `json:"type"`
		Value        interface{}                 `json:"value,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"defaultValue", "description", "label", "name", "type", "value"})
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
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
