// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SharedDashboardSelectableTemplateVariable A template variable that viewers can modify on the shared dashboard.
type SharedDashboardSelectableTemplateVariable struct {
	// Whether viewers can see all tag values for the template variable and specify any value.
	AllowAnyValue bool `json:"allow_any_value"`
	// Default selected values for the variable.
	DefaultValues []string `json:"default_values"`
	// Name of the template variable.
	Name string `json:"name"`
	// Tag prefix for the variable.
	Prefix string `json:"prefix"`
	// Type of the template variable.
	Type string `json:"type"`
	// Restricts which tag values are visible to the viewer.
	VisibleTags []string `json:"visible_tags"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSharedDashboardSelectableTemplateVariable instantiates a new SharedDashboardSelectableTemplateVariable object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSharedDashboardSelectableTemplateVariable(allowAnyValue bool, defaultValues []string, name string, prefix string, typeVar string, visibleTags []string) *SharedDashboardSelectableTemplateVariable {
	this := SharedDashboardSelectableTemplateVariable{}
	this.AllowAnyValue = allowAnyValue
	this.DefaultValues = defaultValues
	this.Name = name
	this.Prefix = prefix
	this.Type = typeVar
	this.VisibleTags = visibleTags
	return &this
}

// NewSharedDashboardSelectableTemplateVariableWithDefaults instantiates a new SharedDashboardSelectableTemplateVariable object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSharedDashboardSelectableTemplateVariableWithDefaults() *SharedDashboardSelectableTemplateVariable {
	this := SharedDashboardSelectableTemplateVariable{}
	return &this
}

// GetAllowAnyValue returns the AllowAnyValue field value.
func (o *SharedDashboardSelectableTemplateVariable) GetAllowAnyValue() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.AllowAnyValue
}

// GetAllowAnyValueOk returns a tuple with the AllowAnyValue field value
// and a boolean to check if the value has been set.
func (o *SharedDashboardSelectableTemplateVariable) GetAllowAnyValueOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowAnyValue, true
}

// SetAllowAnyValue sets field value.
func (o *SharedDashboardSelectableTemplateVariable) SetAllowAnyValue(v bool) {
	o.AllowAnyValue = v
}

// GetDefaultValues returns the DefaultValues field value.
func (o *SharedDashboardSelectableTemplateVariable) GetDefaultValues() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DefaultValues
}

// GetDefaultValuesOk returns a tuple with the DefaultValues field value
// and a boolean to check if the value has been set.
func (o *SharedDashboardSelectableTemplateVariable) GetDefaultValuesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultValues, true
}

// SetDefaultValues sets field value.
func (o *SharedDashboardSelectableTemplateVariable) SetDefaultValues(v []string) {
	o.DefaultValues = v
}

// GetName returns the Name field value.
func (o *SharedDashboardSelectableTemplateVariable) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SharedDashboardSelectableTemplateVariable) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SharedDashboardSelectableTemplateVariable) SetName(v string) {
	o.Name = v
}

// GetPrefix returns the Prefix field value.
func (o *SharedDashboardSelectableTemplateVariable) GetPrefix() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value
// and a boolean to check if the value has been set.
func (o *SharedDashboardSelectableTemplateVariable) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prefix, true
}

// SetPrefix sets field value.
func (o *SharedDashboardSelectableTemplateVariable) SetPrefix(v string) {
	o.Prefix = v
}

// GetType returns the Type field value.
func (o *SharedDashboardSelectableTemplateVariable) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SharedDashboardSelectableTemplateVariable) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SharedDashboardSelectableTemplateVariable) SetType(v string) {
	o.Type = v
}

// GetVisibleTags returns the VisibleTags field value.
func (o *SharedDashboardSelectableTemplateVariable) GetVisibleTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.VisibleTags
}

// GetVisibleTagsOk returns a tuple with the VisibleTags field value
// and a boolean to check if the value has been set.
func (o *SharedDashboardSelectableTemplateVariable) GetVisibleTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VisibleTags, true
}

// SetVisibleTags sets field value.
func (o *SharedDashboardSelectableTemplateVariable) SetVisibleTags(v []string) {
	o.VisibleTags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SharedDashboardSelectableTemplateVariable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["allow_any_value"] = o.AllowAnyValue
	toSerialize["default_values"] = o.DefaultValues
	toSerialize["name"] = o.Name
	toSerialize["prefix"] = o.Prefix
	toSerialize["type"] = o.Type
	toSerialize["visible_tags"] = o.VisibleTags

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SharedDashboardSelectableTemplateVariable) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AllowAnyValue *bool     `json:"allow_any_value"`
		DefaultValues *[]string `json:"default_values"`
		Name          *string   `json:"name"`
		Prefix        *string   `json:"prefix"`
		Type          *string   `json:"type"`
		VisibleTags   *[]string `json:"visible_tags"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AllowAnyValue == nil {
		return fmt.Errorf("required field allow_any_value missing")
	}
	if all.DefaultValues == nil {
		return fmt.Errorf("required field default_values missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Prefix == nil {
		return fmt.Errorf("required field prefix missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.VisibleTags == nil {
		return fmt.Errorf("required field visible_tags missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"allow_any_value", "default_values", "name", "prefix", "type", "visible_tags"})
	} else {
		return err
	}
	o.AllowAnyValue = *all.AllowAnyValue
	o.DefaultValues = *all.DefaultValues
	o.Name = *all.Name
	o.Prefix = *all.Prefix
	o.Type = *all.Type
	o.VisibleTags = *all.VisibleTags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
