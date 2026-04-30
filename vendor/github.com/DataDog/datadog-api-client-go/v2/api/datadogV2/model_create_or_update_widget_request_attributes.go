// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateOrUpdateWidgetRequestAttributes Attributes for creating or updating a widget.
type CreateOrUpdateWidgetRequestAttributes struct {
	// The definition of a widget, including its type and configuration.
	Definition WidgetDefinition `json:"definition"`
	// User-defined tags for organizing the widget.
	Tags datadog.NullableList[string] `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateOrUpdateWidgetRequestAttributes instantiates a new CreateOrUpdateWidgetRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateOrUpdateWidgetRequestAttributes(definition WidgetDefinition) *CreateOrUpdateWidgetRequestAttributes {
	this := CreateOrUpdateWidgetRequestAttributes{}
	this.Definition = definition
	return &this
}

// NewCreateOrUpdateWidgetRequestAttributesWithDefaults instantiates a new CreateOrUpdateWidgetRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateOrUpdateWidgetRequestAttributesWithDefaults() *CreateOrUpdateWidgetRequestAttributes {
	this := CreateOrUpdateWidgetRequestAttributes{}
	return &this
}

// GetDefinition returns the Definition field value.
func (o *CreateOrUpdateWidgetRequestAttributes) GetDefinition() WidgetDefinition {
	if o == nil {
		var ret WidgetDefinition
		return ret
	}
	return o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateWidgetRequestAttributes) GetDefinitionOk() (*WidgetDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Definition, true
}

// SetDefinition sets field value.
func (o *CreateOrUpdateWidgetRequestAttributes) SetDefinition(v WidgetDefinition) {
	o.Definition = v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOrUpdateWidgetRequestAttributes) GetTags() []string {
	if o == nil || o.Tags.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Tags.Get()
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CreateOrUpdateWidgetRequestAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// HasTags returns a boolean if a field has been set.
func (o *CreateOrUpdateWidgetRequestAttributes) HasTags() bool {
	return o != nil && o.Tags.IsSet()
}

// SetTags gets a reference to the given datadog.NullableList[string] and assigns it to the Tags field.
func (o *CreateOrUpdateWidgetRequestAttributes) SetTags(v []string) {
	o.Tags.Set(&v)
}

// SetTagsNil sets the value for Tags to be an explicit nil.
func (o *CreateOrUpdateWidgetRequestAttributes) SetTagsNil() {
	o.Tags.Set(nil)
}

// UnsetTags ensures that no value is present for Tags, not even an explicit nil.
func (o *CreateOrUpdateWidgetRequestAttributes) UnsetTags() {
	o.Tags.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateOrUpdateWidgetRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["definition"] = o.Definition
	if o.Tags.IsSet() {
		toSerialize["tags"] = o.Tags.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateOrUpdateWidgetRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Definition *WidgetDefinition            `json:"definition"`
		Tags       datadog.NullableList[string] `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Definition == nil {
		return fmt.Errorf("required field definition missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"definition", "tags"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Definition.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Definition = *all.Definition
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
