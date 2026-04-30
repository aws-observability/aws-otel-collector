// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecureEmbedSelectableTemplateVariable A template variable that viewers can modify on the secure embed shared dashboard.
type SecureEmbedSelectableTemplateVariable struct {
	// Default selected values for the variable.
	DefaultValues []string `json:"default_values,omitempty"`
	// Name of the template variable. Usually matches the prefix unless you want a different display name.
	Name *string `json:"name,omitempty"`
	// Tag prefix for the variable (e.g., `environment`, `service`).
	Prefix *string `json:"prefix,omitempty"`
	// Restrict which tag values are visible to the viewer.
	VisibleTags []string `json:"visible_tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecureEmbedSelectableTemplateVariable instantiates a new SecureEmbedSelectableTemplateVariable object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecureEmbedSelectableTemplateVariable() *SecureEmbedSelectableTemplateVariable {
	this := SecureEmbedSelectableTemplateVariable{}
	return &this
}

// NewSecureEmbedSelectableTemplateVariableWithDefaults instantiates a new SecureEmbedSelectableTemplateVariable object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecureEmbedSelectableTemplateVariableWithDefaults() *SecureEmbedSelectableTemplateVariable {
	this := SecureEmbedSelectableTemplateVariable{}
	return &this
}

// GetDefaultValues returns the DefaultValues field value if set, zero value otherwise.
func (o *SecureEmbedSelectableTemplateVariable) GetDefaultValues() []string {
	if o == nil || o.DefaultValues == nil {
		var ret []string
		return ret
	}
	return o.DefaultValues
}

// GetDefaultValuesOk returns a tuple with the DefaultValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureEmbedSelectableTemplateVariable) GetDefaultValuesOk() (*[]string, bool) {
	if o == nil || o.DefaultValues == nil {
		return nil, false
	}
	return &o.DefaultValues, true
}

// HasDefaultValues returns a boolean if a field has been set.
func (o *SecureEmbedSelectableTemplateVariable) HasDefaultValues() bool {
	return o != nil && o.DefaultValues != nil
}

// SetDefaultValues gets a reference to the given []string and assigns it to the DefaultValues field.
func (o *SecureEmbedSelectableTemplateVariable) SetDefaultValues(v []string) {
	o.DefaultValues = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SecureEmbedSelectableTemplateVariable) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureEmbedSelectableTemplateVariable) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SecureEmbedSelectableTemplateVariable) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SecureEmbedSelectableTemplateVariable) SetName(v string) {
	o.Name = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *SecureEmbedSelectableTemplateVariable) GetPrefix() string {
	if o == nil || o.Prefix == nil {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureEmbedSelectableTemplateVariable) GetPrefixOk() (*string, bool) {
	if o == nil || o.Prefix == nil {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *SecureEmbedSelectableTemplateVariable) HasPrefix() bool {
	return o != nil && o.Prefix != nil
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *SecureEmbedSelectableTemplateVariable) SetPrefix(v string) {
	o.Prefix = &v
}

// GetVisibleTags returns the VisibleTags field value if set, zero value otherwise.
func (o *SecureEmbedSelectableTemplateVariable) GetVisibleTags() []string {
	if o == nil || o.VisibleTags == nil {
		var ret []string
		return ret
	}
	return o.VisibleTags
}

// GetVisibleTagsOk returns a tuple with the VisibleTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecureEmbedSelectableTemplateVariable) GetVisibleTagsOk() (*[]string, bool) {
	if o == nil || o.VisibleTags == nil {
		return nil, false
	}
	return &o.VisibleTags, true
}

// HasVisibleTags returns a boolean if a field has been set.
func (o *SecureEmbedSelectableTemplateVariable) HasVisibleTags() bool {
	return o != nil && o.VisibleTags != nil
}

// SetVisibleTags gets a reference to the given []string and assigns it to the VisibleTags field.
func (o *SecureEmbedSelectableTemplateVariable) SetVisibleTags(v []string) {
	o.VisibleTags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecureEmbedSelectableTemplateVariable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DefaultValues != nil {
		toSerialize["default_values"] = o.DefaultValues
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Prefix != nil {
		toSerialize["prefix"] = o.Prefix
	}
	if o.VisibleTags != nil {
		toSerialize["visible_tags"] = o.VisibleTags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecureEmbedSelectableTemplateVariable) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DefaultValues []string `json:"default_values,omitempty"`
		Name          *string  `json:"name,omitempty"`
		Prefix        *string  `json:"prefix,omitempty"`
		VisibleTags   []string `json:"visible_tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"default_values", "name", "prefix", "visible_tags"})
	} else {
		return err
	}
	o.DefaultValues = all.DefaultValues
	o.Name = all.Name
	o.Prefix = all.Prefix
	o.VisibleTags = all.VisibleTags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
