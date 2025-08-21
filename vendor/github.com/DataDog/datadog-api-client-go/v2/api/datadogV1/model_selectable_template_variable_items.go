// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SelectableTemplateVariableItems Object containing the template variable's name, associated tag/attribute, default value and selectable values.
type SelectableTemplateVariableItems struct {
	// The default value of the template variable.
	DefaultValue *string `json:"default_value,omitempty"`
	// Name of the template variable.
	Name *string `json:"name,omitempty"`
	// The tag/attribute key associated with the template variable.
	Prefix *string `json:"prefix,omitempty"`
	// The type of variable. This is to differentiate between filter variables (interpolated in query) and group by variables (interpolated into group by).
	Type datadog.NullableString `json:"type,omitempty"`
	// List of visible tag values on the shared dashboard.
	VisibleTags datadog.NullableList[string] `json:"visible_tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSelectableTemplateVariableItems instantiates a new SelectableTemplateVariableItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSelectableTemplateVariableItems() *SelectableTemplateVariableItems {
	this := SelectableTemplateVariableItems{}
	return &this
}

// NewSelectableTemplateVariableItemsWithDefaults instantiates a new SelectableTemplateVariableItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSelectableTemplateVariableItemsWithDefaults() *SelectableTemplateVariableItems {
	this := SelectableTemplateVariableItems{}
	return &this
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *SelectableTemplateVariableItems) GetDefaultValue() string {
	if o == nil || o.DefaultValue == nil {
		var ret string
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectableTemplateVariableItems) GetDefaultValueOk() (*string, bool) {
	if o == nil || o.DefaultValue == nil {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *SelectableTemplateVariableItems) HasDefaultValue() bool {
	return o != nil && o.DefaultValue != nil
}

// SetDefaultValue gets a reference to the given string and assigns it to the DefaultValue field.
func (o *SelectableTemplateVariableItems) SetDefaultValue(v string) {
	o.DefaultValue = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SelectableTemplateVariableItems) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectableTemplateVariableItems) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SelectableTemplateVariableItems) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SelectableTemplateVariableItems) SetName(v string) {
	o.Name = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *SelectableTemplateVariableItems) GetPrefix() string {
	if o == nil || o.Prefix == nil {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectableTemplateVariableItems) GetPrefixOk() (*string, bool) {
	if o == nil || o.Prefix == nil {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *SelectableTemplateVariableItems) HasPrefix() bool {
	return o != nil && o.Prefix != nil
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *SelectableTemplateVariableItems) SetPrefix(v string) {
	o.Prefix = &v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SelectableTemplateVariableItems) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SelectableTemplateVariableItems) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *SelectableTemplateVariableItems) HasType() bool {
	return o != nil && o.Type.IsSet()
}

// SetType gets a reference to the given datadog.NullableString and assigns it to the Type field.
func (o *SelectableTemplateVariableItems) SetType(v string) {
	o.Type.Set(&v)
}

// SetTypeNil sets the value for Type to be an explicit nil.
func (o *SelectableTemplateVariableItems) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil.
func (o *SelectableTemplateVariableItems) UnsetType() {
	o.Type.Unset()
}

// GetVisibleTags returns the VisibleTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SelectableTemplateVariableItems) GetVisibleTags() []string {
	if o == nil || o.VisibleTags.Get() == nil {
		var ret []string
		return ret
	}
	return *o.VisibleTags.Get()
}

// GetVisibleTagsOk returns a tuple with the VisibleTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SelectableTemplateVariableItems) GetVisibleTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VisibleTags.Get(), o.VisibleTags.IsSet()
}

// HasVisibleTags returns a boolean if a field has been set.
func (o *SelectableTemplateVariableItems) HasVisibleTags() bool {
	return o != nil && o.VisibleTags.IsSet()
}

// SetVisibleTags gets a reference to the given datadog.NullableList[string] and assigns it to the VisibleTags field.
func (o *SelectableTemplateVariableItems) SetVisibleTags(v []string) {
	o.VisibleTags.Set(&v)
}

// SetVisibleTagsNil sets the value for VisibleTags to be an explicit nil.
func (o *SelectableTemplateVariableItems) SetVisibleTagsNil() {
	o.VisibleTags.Set(nil)
}

// UnsetVisibleTags ensures that no value is present for VisibleTags, not even an explicit nil.
func (o *SelectableTemplateVariableItems) UnsetVisibleTags() {
	o.VisibleTags.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o SelectableTemplateVariableItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DefaultValue != nil {
		toSerialize["default_value"] = o.DefaultValue
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Prefix != nil {
		toSerialize["prefix"] = o.Prefix
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.VisibleTags.IsSet() {
		toSerialize["visible_tags"] = o.VisibleTags.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SelectableTemplateVariableItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DefaultValue *string                      `json:"default_value,omitempty"`
		Name         *string                      `json:"name,omitempty"`
		Prefix       *string                      `json:"prefix,omitempty"`
		Type         datadog.NullableString       `json:"type,omitempty"`
		VisibleTags  datadog.NullableList[string] `json:"visible_tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"default_value", "name", "prefix", "type", "visible_tags"})
	} else {
		return err
	}
	o.DefaultValue = all.DefaultValue
	o.Name = all.Name
	o.Prefix = all.Prefix
	o.Type = all.Type
	o.VisibleTags = all.VisibleTags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
