// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgGroupPolicyConfigAttributes Attributes of an org group policy config.
type OrgGroupPolicyConfigAttributes struct {
	// The allowed values for this config.
	AllowedValues []string `json:"allowed_values"`
	// The default value for this config.
	DefaultValue interface{} `json:"default_value"`
	// The description of the policy config.
	Description string `json:"description"`
	// The name of the policy config.
	Name string `json:"name"`
	// The type of the value for this config.
	ValueType string `json:"value_type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgGroupPolicyConfigAttributes instantiates a new OrgGroupPolicyConfigAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgGroupPolicyConfigAttributes(allowedValues []string, defaultValue interface{}, description string, name string, valueType string) *OrgGroupPolicyConfigAttributes {
	this := OrgGroupPolicyConfigAttributes{}
	this.AllowedValues = allowedValues
	this.DefaultValue = defaultValue
	this.Description = description
	this.Name = name
	this.ValueType = valueType
	return &this
}

// NewOrgGroupPolicyConfigAttributesWithDefaults instantiates a new OrgGroupPolicyConfigAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgGroupPolicyConfigAttributesWithDefaults() *OrgGroupPolicyConfigAttributes {
	this := OrgGroupPolicyConfigAttributes{}
	return &this
}

// GetAllowedValues returns the AllowedValues field value.
func (o *OrgGroupPolicyConfigAttributes) GetAllowedValues() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AllowedValues
}

// GetAllowedValuesOk returns a tuple with the AllowedValues field value
// and a boolean to check if the value has been set.
func (o *OrgGroupPolicyConfigAttributes) GetAllowedValuesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowedValues, true
}

// SetAllowedValues sets field value.
func (o *OrgGroupPolicyConfigAttributes) SetAllowedValues(v []string) {
	o.AllowedValues = v
}

// GetDefaultValue returns the DefaultValue field value.
func (o *OrgGroupPolicyConfigAttributes) GetDefaultValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value
// and a boolean to check if the value has been set.
func (o *OrgGroupPolicyConfigAttributes) GetDefaultValueOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultValue, true
}

// SetDefaultValue sets field value.
func (o *OrgGroupPolicyConfigAttributes) SetDefaultValue(v interface{}) {
	o.DefaultValue = v
}

// GetDescription returns the Description field value.
func (o *OrgGroupPolicyConfigAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *OrgGroupPolicyConfigAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *OrgGroupPolicyConfigAttributes) SetDescription(v string) {
	o.Description = v
}

// GetName returns the Name field value.
func (o *OrgGroupPolicyConfigAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OrgGroupPolicyConfigAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *OrgGroupPolicyConfigAttributes) SetName(v string) {
	o.Name = v
}

// GetValueType returns the ValueType field value.
func (o *OrgGroupPolicyConfigAttributes) GetValueType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value
// and a boolean to check if the value has been set.
func (o *OrgGroupPolicyConfigAttributes) GetValueTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValueType, true
}

// SetValueType sets field value.
func (o *OrgGroupPolicyConfigAttributes) SetValueType(v string) {
	o.ValueType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgGroupPolicyConfigAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["allowed_values"] = o.AllowedValues
	toSerialize["default_value"] = o.DefaultValue
	toSerialize["description"] = o.Description
	toSerialize["name"] = o.Name
	toSerialize["value_type"] = o.ValueType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgGroupPolicyConfigAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AllowedValues *[]string    `json:"allowed_values"`
		DefaultValue  *interface{} `json:"default_value"`
		Description   *string      `json:"description"`
		Name          *string      `json:"name"`
		ValueType     *string      `json:"value_type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AllowedValues == nil {
		return fmt.Errorf("required field allowed_values missing")
	}
	if all.DefaultValue == nil {
		return fmt.Errorf("required field default_value missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.ValueType == nil {
		return fmt.Errorf("required field value_type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"allowed_values", "default_value", "description", "name", "value_type"})
	} else {
		return err
	}
	o.AllowedValues = *all.AllowedValues
	o.DefaultValue = *all.DefaultValue
	o.Description = *all.Description
	o.Name = *all.Name
	o.ValueType = *all.ValueType

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
