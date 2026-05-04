// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetAttributes Attributes of a widget resource.
type WidgetAttributes struct {
	// ISO 8601 timestamp of when the widget was created.
	CreatedAt string `json:"created_at"`
	// The definition of a widget, including its type and configuration.
	Definition WidgetDefinition `json:"definition"`
	// Will be implemented soon. Currently always returns false.
	IsFavorited bool `json:"is_favorited"`
	// ISO 8601 timestamp of when the widget was last modified.
	ModifiedAt string `json:"modified_at"`
	// User-defined tags for organizing widgets.
	Tags datadog.NullableList[string] `json:"tags"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWidgetAttributes instantiates a new WidgetAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWidgetAttributes(createdAt string, definition WidgetDefinition, isFavorited bool, modifiedAt string, tags datadog.NullableList[string]) *WidgetAttributes {
	this := WidgetAttributes{}
	this.CreatedAt = createdAt
	this.Definition = definition
	this.IsFavorited = isFavorited
	this.ModifiedAt = modifiedAt
	this.Tags = tags
	return &this
}

// NewWidgetAttributesWithDefaults instantiates a new WidgetAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWidgetAttributesWithDefaults() *WidgetAttributes {
	this := WidgetAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *WidgetAttributes) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *WidgetAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *WidgetAttributes) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetDefinition returns the Definition field value.
func (o *WidgetAttributes) GetDefinition() WidgetDefinition {
	if o == nil {
		var ret WidgetDefinition
		return ret
	}
	return o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value
// and a boolean to check if the value has been set.
func (o *WidgetAttributes) GetDefinitionOk() (*WidgetDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Definition, true
}

// SetDefinition sets field value.
func (o *WidgetAttributes) SetDefinition(v WidgetDefinition) {
	o.Definition = v
}

// GetIsFavorited returns the IsFavorited field value.
func (o *WidgetAttributes) GetIsFavorited() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsFavorited
}

// GetIsFavoritedOk returns a tuple with the IsFavorited field value
// and a boolean to check if the value has been set.
func (o *WidgetAttributes) GetIsFavoritedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsFavorited, true
}

// SetIsFavorited sets field value.
func (o *WidgetAttributes) SetIsFavorited(v bool) {
	o.IsFavorited = v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *WidgetAttributes) GetModifiedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *WidgetAttributes) GetModifiedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *WidgetAttributes) SetModifiedAt(v string) {
	o.ModifiedAt = v
}

// GetTags returns the Tags field value.
// If the value is explicit nil, the zero value for []string will be returned.
func (o *WidgetAttributes) GetTags() []string {
	if o == nil || o.Tags.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Tags.Get()
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *WidgetAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// SetTags sets field value.
func (o *WidgetAttributes) SetTags(v []string) {
	o.Tags.Set(&v)
}

// MarshalJSON serializes the struct using spec logic.
func (o WidgetAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["definition"] = o.Definition
	toSerialize["is_favorited"] = o.IsFavorited
	toSerialize["modified_at"] = o.ModifiedAt
	toSerialize["tags"] = o.Tags.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WidgetAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt   *string                      `json:"created_at"`
		Definition  *WidgetDefinition            `json:"definition"`
		IsFavorited *bool                        `json:"is_favorited"`
		ModifiedAt  *string                      `json:"modified_at"`
		Tags        datadog.NullableList[string] `json:"tags"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.Definition == nil {
		return fmt.Errorf("required field definition missing")
	}
	if all.IsFavorited == nil {
		return fmt.Errorf("required field is_favorited missing")
	}
	if all.ModifiedAt == nil {
		return fmt.Errorf("required field modified_at missing")
	}
	if !all.Tags.IsSet() {
		return fmt.Errorf("required field tags missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "definition", "is_favorited", "modified_at", "tags"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = *all.CreatedAt
	if all.Definition.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Definition = *all.Definition
	o.IsFavorited = *all.IsFavorited
	o.ModifiedAt = *all.ModifiedAt
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
