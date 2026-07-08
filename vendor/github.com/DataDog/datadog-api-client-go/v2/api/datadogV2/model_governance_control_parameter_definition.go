// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GovernanceControlParameterDefinition The definition of a configurable parameter on a control or mitigation.
type GovernanceControlParameterDefinition struct {
	// The default value of the parameter. The JSON type depends on the parameter's `type`.
	DefaultValue interface{} `json:"default_value"`
	// A human-readable description of the parameter.
	Description string `json:"description"`
	// The human-readable name of the parameter.
	DisplayName string `json:"display_name"`
	// Whether the parameter is hidden from the UI.
	Hidden bool `json:"hidden"`
	// The machine-readable name of the parameter.
	Name string `json:"name"`
	// Whether the parameter must be provided.
	Required bool `json:"required"`
	// The supported values for an enumerated parameter.
	SupportedValues []GovernanceControlSupportedValue `json:"supported_values"`
	// The type of the parameter, such as `integer`, `string`, `boolean`, `enum`, or `pattern_list`.
	Type string `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGovernanceControlParameterDefinition instantiates a new GovernanceControlParameterDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGovernanceControlParameterDefinition(defaultValue interface{}, description string, displayName string, hidden bool, name string, required bool, supportedValues []GovernanceControlSupportedValue, typeVar string) *GovernanceControlParameterDefinition {
	this := GovernanceControlParameterDefinition{}
	this.DefaultValue = defaultValue
	this.Description = description
	this.DisplayName = displayName
	this.Hidden = hidden
	this.Name = name
	this.Required = required
	this.SupportedValues = supportedValues
	this.Type = typeVar
	return &this
}

// NewGovernanceControlParameterDefinitionWithDefaults instantiates a new GovernanceControlParameterDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGovernanceControlParameterDefinitionWithDefaults() *GovernanceControlParameterDefinition {
	this := GovernanceControlParameterDefinition{}
	return &this
}

// GetDefaultValue returns the DefaultValue field value.
func (o *GovernanceControlParameterDefinition) GetDefaultValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlParameterDefinition) GetDefaultValueOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultValue, true
}

// SetDefaultValue sets field value.
func (o *GovernanceControlParameterDefinition) SetDefaultValue(v interface{}) {
	o.DefaultValue = v
}

// GetDescription returns the Description field value.
func (o *GovernanceControlParameterDefinition) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlParameterDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *GovernanceControlParameterDefinition) SetDescription(v string) {
	o.Description = v
}

// GetDisplayName returns the DisplayName field value.
func (o *GovernanceControlParameterDefinition) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlParameterDefinition) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value.
func (o *GovernanceControlParameterDefinition) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetHidden returns the Hidden field value.
func (o *GovernanceControlParameterDefinition) GetHidden() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlParameterDefinition) GetHiddenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hidden, true
}

// SetHidden sets field value.
func (o *GovernanceControlParameterDefinition) SetHidden(v bool) {
	o.Hidden = v
}

// GetName returns the Name field value.
func (o *GovernanceControlParameterDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlParameterDefinition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *GovernanceControlParameterDefinition) SetName(v string) {
	o.Name = v
}

// GetRequired returns the Required field value.
func (o *GovernanceControlParameterDefinition) GetRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlParameterDefinition) GetRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Required, true
}

// SetRequired sets field value.
func (o *GovernanceControlParameterDefinition) SetRequired(v bool) {
	o.Required = v
}

// GetSupportedValues returns the SupportedValues field value.
func (o *GovernanceControlParameterDefinition) GetSupportedValues() []GovernanceControlSupportedValue {
	if o == nil {
		var ret []GovernanceControlSupportedValue
		return ret
	}
	return o.SupportedValues
}

// GetSupportedValuesOk returns a tuple with the SupportedValues field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlParameterDefinition) GetSupportedValuesOk() (*[]GovernanceControlSupportedValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportedValues, true
}

// SetSupportedValues sets field value.
func (o *GovernanceControlParameterDefinition) SetSupportedValues(v []GovernanceControlSupportedValue) {
	o.SupportedValues = v
}

// GetType returns the Type field value.
func (o *GovernanceControlParameterDefinition) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlParameterDefinition) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *GovernanceControlParameterDefinition) SetType(v string) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GovernanceControlParameterDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["default_value"] = o.DefaultValue
	toSerialize["description"] = o.Description
	toSerialize["display_name"] = o.DisplayName
	toSerialize["hidden"] = o.Hidden
	toSerialize["name"] = o.Name
	toSerialize["required"] = o.Required
	toSerialize["supported_values"] = o.SupportedValues
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GovernanceControlParameterDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DefaultValue    *interface{}                       `json:"default_value"`
		Description     *string                            `json:"description"`
		DisplayName     *string                            `json:"display_name"`
		Hidden          *bool                              `json:"hidden"`
		Name            *string                            `json:"name"`
		Required        *bool                              `json:"required"`
		SupportedValues *[]GovernanceControlSupportedValue `json:"supported_values"`
		Type            *string                            `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DefaultValue == nil {
		return fmt.Errorf("required field default_value missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.DisplayName == nil {
		return fmt.Errorf("required field display_name missing")
	}
	if all.Hidden == nil {
		return fmt.Errorf("required field hidden missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Required == nil {
		return fmt.Errorf("required field required missing")
	}
	if all.SupportedValues == nil {
		return fmt.Errorf("required field supported_values missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"default_value", "description", "display_name", "hidden", "name", "required", "supported_values", "type"})
	} else {
		return err
	}
	o.DefaultValue = *all.DefaultValue
	o.Description = *all.Description
	o.DisplayName = *all.DisplayName
	o.Hidden = *all.Hidden
	o.Name = *all.Name
	o.Required = *all.Required
	o.SupportedValues = *all.SupportedValues
	o.Type = *all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
