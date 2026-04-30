// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateFeatureFlagAttributes Attributes for updating a feature flag.
type UpdateFeatureFlagAttributes struct {
	// The description of the feature flag.
	Description *string `json:"description,omitempty"`
	// JSON schema for validation when value_type is JSON.
	JsonSchema datadog.NullableString `json:"json_schema,omitempty"`
	// The name of the feature flag.
	Name *string `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpdateFeatureFlagAttributes instantiates a new UpdateFeatureFlagAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpdateFeatureFlagAttributes() *UpdateFeatureFlagAttributes {
	this := UpdateFeatureFlagAttributes{}
	return &this
}

// NewUpdateFeatureFlagAttributesWithDefaults instantiates a new UpdateFeatureFlagAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpdateFeatureFlagAttributesWithDefaults() *UpdateFeatureFlagAttributes {
	this := UpdateFeatureFlagAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateFeatureFlagAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFeatureFlagAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateFeatureFlagAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateFeatureFlagAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetJsonSchema returns the JsonSchema field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateFeatureFlagAttributes) GetJsonSchema() string {
	if o == nil || o.JsonSchema.Get() == nil {
		var ret string
		return ret
	}
	return *o.JsonSchema.Get()
}

// GetJsonSchemaOk returns a tuple with the JsonSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UpdateFeatureFlagAttributes) GetJsonSchemaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.JsonSchema.Get(), o.JsonSchema.IsSet()
}

// HasJsonSchema returns a boolean if a field has been set.
func (o *UpdateFeatureFlagAttributes) HasJsonSchema() bool {
	return o != nil && o.JsonSchema.IsSet()
}

// SetJsonSchema gets a reference to the given datadog.NullableString and assigns it to the JsonSchema field.
func (o *UpdateFeatureFlagAttributes) SetJsonSchema(v string) {
	o.JsonSchema.Set(&v)
}

// SetJsonSchemaNil sets the value for JsonSchema to be an explicit nil.
func (o *UpdateFeatureFlagAttributes) SetJsonSchemaNil() {
	o.JsonSchema.Set(nil)
}

// UnsetJsonSchema ensures that no value is present for JsonSchema, not even an explicit nil.
func (o *UpdateFeatureFlagAttributes) UnsetJsonSchema() {
	o.JsonSchema.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateFeatureFlagAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFeatureFlagAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateFeatureFlagAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateFeatureFlagAttributes) SetName(v string) {
	o.Name = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpdateFeatureFlagAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.JsonSchema.IsSet() {
		toSerialize["json_schema"] = o.JsonSchema.Get()
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpdateFeatureFlagAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string                `json:"description,omitempty"`
		JsonSchema  datadog.NullableString `json:"json_schema,omitempty"`
		Name        *string                `json:"name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "json_schema", "name"})
	} else {
		return err
	}
	o.Description = all.Description
	o.JsonSchema = all.JsonSchema
	o.Name = all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
