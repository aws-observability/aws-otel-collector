// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentUserDefinedFieldAttributesCreateRequest Attributes for creating an incident user-defined field.
type IncidentUserDefinedFieldAttributesCreateRequest struct {
	// The section in which the field appears: "what_happened" or "why_it_happened". When null, the field appears in the Attributes section.
	Category NullableIncidentUserDefinedFieldCategory `json:"category,omitempty"`
	// The lifecycle stage at which the app prompts users to fill out this field. Cannot be set on required fields.
	Collected NullableIncidentUserDefinedFieldCollected `json:"collected,omitempty"`
	// The default value for the field. Must be one of the valid values when valid_values is set.
	DefaultValue datadog.NullableString `json:"default_value,omitempty"`
	// The human-readable name shown in the UI. Defaults to a formatted version of the name if not provided.
	DisplayName *string `json:"display_name,omitempty"`
	// The unique identifier of the field. Must start with a letter or digit and contain only letters, digits, underscores, or periods.
	Name string `json:"name"`
	// A decimal string representing the field's display order in the UI.
	Ordinal datadog.NullableString `json:"ordinal,omitempty"`
	// When true, users must fill out this field on incidents.
	Required *bool `json:"required,omitempty"`
	// For metric tag-type fields only, the metric tag key that powers the autocomplete options.
	TagKey datadog.NullableString `json:"tag_key,omitempty"`
	// The data type of the field. 1=dropdown, 2=multiselect, 3=textbox, 4=textarray, 5=metrictag, 6=autocomplete, 7=number, 8=datetime.
	Type IncidentUserDefinedFieldFieldType `json:"type"`
	// The list of allowed values for dropdown and multiselect fields. Limited to 1000 values.
	ValidValues []IncidentUserDefinedFieldValidValue `json:"valid_values,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentUserDefinedFieldAttributesCreateRequest instantiates a new IncidentUserDefinedFieldAttributesCreateRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentUserDefinedFieldAttributesCreateRequest(name string, typeVar IncidentUserDefinedFieldFieldType) *IncidentUserDefinedFieldAttributesCreateRequest {
	this := IncidentUserDefinedFieldAttributesCreateRequest{}
	this.Name = name
	this.Type = typeVar
	return &this
}

// NewIncidentUserDefinedFieldAttributesCreateRequestWithDefaults instantiates a new IncidentUserDefinedFieldAttributesCreateRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentUserDefinedFieldAttributesCreateRequestWithDefaults() *IncidentUserDefinedFieldAttributesCreateRequest {
	this := IncidentUserDefinedFieldAttributesCreateRequest{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentUserDefinedFieldAttributesCreateRequest) GetCategory() IncidentUserDefinedFieldCategory {
	if o == nil || o.Category.Get() == nil {
		var ret IncidentUserDefinedFieldCategory
		return ret
	}
	return *o.Category.Get()
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) GetCategoryOk() (*IncidentUserDefinedFieldCategory, bool) {
	if o == nil {
		return nil, false
	}
	return o.Category.Get(), o.Category.IsSet()
}

// HasCategory returns a boolean if a field has been set.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) HasCategory() bool {
	return o != nil && o.Category.IsSet()
}

// SetCategory gets a reference to the given NullableIncidentUserDefinedFieldCategory and assigns it to the Category field.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) SetCategory(v IncidentUserDefinedFieldCategory) {
	o.Category.Set(&v)
}

// SetCategoryNil sets the value for Category to be an explicit nil.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) SetCategoryNil() {
	o.Category.Set(nil)
}

// UnsetCategory ensures that no value is present for Category, not even an explicit nil.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) UnsetCategory() {
	o.Category.Unset()
}

// GetCollected returns the Collected field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentUserDefinedFieldAttributesCreateRequest) GetCollected() IncidentUserDefinedFieldCollected {
	if o == nil || o.Collected.Get() == nil {
		var ret IncidentUserDefinedFieldCollected
		return ret
	}
	return *o.Collected.Get()
}

// GetCollectedOk returns a tuple with the Collected field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) GetCollectedOk() (*IncidentUserDefinedFieldCollected, bool) {
	if o == nil {
		return nil, false
	}
	return o.Collected.Get(), o.Collected.IsSet()
}

// HasCollected returns a boolean if a field has been set.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) HasCollected() bool {
	return o != nil && o.Collected.IsSet()
}

// SetCollected gets a reference to the given NullableIncidentUserDefinedFieldCollected and assigns it to the Collected field.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) SetCollected(v IncidentUserDefinedFieldCollected) {
	o.Collected.Set(&v)
}

// SetCollectedNil sets the value for Collected to be an explicit nil.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) SetCollectedNil() {
	o.Collected.Set(nil)
}

// UnsetCollected ensures that no value is present for Collected, not even an explicit nil.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) UnsetCollected() {
	o.Collected.Unset()
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentUserDefinedFieldAttributesCreateRequest) GetDefaultValue() string {
	if o == nil || o.DefaultValue.Get() == nil {
		var ret string
		return ret
	}
	return *o.DefaultValue.Get()
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) GetDefaultValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultValue.Get(), o.DefaultValue.IsSet()
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) HasDefaultValue() bool {
	return o != nil && o.DefaultValue.IsSet()
}

// SetDefaultValue gets a reference to the given datadog.NullableString and assigns it to the DefaultValue field.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) SetDefaultValue(v string) {
	o.DefaultValue.Set(&v)
}

// SetDefaultValueNil sets the value for DefaultValue to be an explicit nil.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) SetDefaultValueNil() {
	o.DefaultValue.Set(nil)
}

// UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) UnsetDefaultValue() {
	o.DefaultValue.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetName returns the Name field value.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) SetName(v string) {
	o.Name = v
}

// GetOrdinal returns the Ordinal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentUserDefinedFieldAttributesCreateRequest) GetOrdinal() string {
	if o == nil || o.Ordinal.Get() == nil {
		var ret string
		return ret
	}
	return *o.Ordinal.Get()
}

// GetOrdinalOk returns a tuple with the Ordinal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) GetOrdinalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ordinal.Get(), o.Ordinal.IsSet()
}

// HasOrdinal returns a boolean if a field has been set.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) HasOrdinal() bool {
	return o != nil && o.Ordinal.IsSet()
}

// SetOrdinal gets a reference to the given datadog.NullableString and assigns it to the Ordinal field.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) SetOrdinal(v string) {
	o.Ordinal.Set(&v)
}

// SetOrdinalNil sets the value for Ordinal to be an explicit nil.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) SetOrdinalNil() {
	o.Ordinal.Set(nil)
}

// UnsetOrdinal ensures that no value is present for Ordinal, not even an explicit nil.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) UnsetOrdinal() {
	o.Ordinal.Unset()
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) HasRequired() bool {
	return o != nil && o.Required != nil
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) SetRequired(v bool) {
	o.Required = &v
}

// GetTagKey returns the TagKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentUserDefinedFieldAttributesCreateRequest) GetTagKey() string {
	if o == nil || o.TagKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.TagKey.Get()
}

// GetTagKeyOk returns a tuple with the TagKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) GetTagKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TagKey.Get(), o.TagKey.IsSet()
}

// HasTagKey returns a boolean if a field has been set.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) HasTagKey() bool {
	return o != nil && o.TagKey.IsSet()
}

// SetTagKey gets a reference to the given datadog.NullableString and assigns it to the TagKey field.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) SetTagKey(v string) {
	o.TagKey.Set(&v)
}

// SetTagKeyNil sets the value for TagKey to be an explicit nil.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) SetTagKeyNil() {
	o.TagKey.Set(nil)
}

// UnsetTagKey ensures that no value is present for TagKey, not even an explicit nil.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) UnsetTagKey() {
	o.TagKey.Unset()
}

// GetType returns the Type field value.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) GetType() IncidentUserDefinedFieldFieldType {
	if o == nil {
		var ret IncidentUserDefinedFieldFieldType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) GetTypeOk() (*IncidentUserDefinedFieldFieldType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) SetType(v IncidentUserDefinedFieldFieldType) {
	o.Type = v
}

// GetValidValues returns the ValidValues field value if set, zero value otherwise.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) GetValidValues() []IncidentUserDefinedFieldValidValue {
	if o == nil || o.ValidValues == nil {
		var ret []IncidentUserDefinedFieldValidValue
		return ret
	}
	return o.ValidValues
}

// GetValidValuesOk returns a tuple with the ValidValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) GetValidValuesOk() (*[]IncidentUserDefinedFieldValidValue, bool) {
	if o == nil || o.ValidValues == nil {
		return nil, false
	}
	return &o.ValidValues, true
}

// HasValidValues returns a boolean if a field has been set.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) HasValidValues() bool {
	return o != nil && o.ValidValues != nil
}

// SetValidValues gets a reference to the given []IncidentUserDefinedFieldValidValue and assigns it to the ValidValues field.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) SetValidValues(v []IncidentUserDefinedFieldValidValue) {
	o.ValidValues = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentUserDefinedFieldAttributesCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Category.IsSet() {
		toSerialize["category"] = o.Category.Get()
	}
	if o.Collected.IsSet() {
		toSerialize["collected"] = o.Collected.Get()
	}
	if o.DefaultValue.IsSet() {
		toSerialize["default_value"] = o.DefaultValue.Get()
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	toSerialize["name"] = o.Name
	if o.Ordinal.IsSet() {
		toSerialize["ordinal"] = o.Ordinal.Get()
	}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}
	if o.TagKey.IsSet() {
		toSerialize["tag_key"] = o.TagKey.Get()
	}
	toSerialize["type"] = o.Type
	if o.ValidValues != nil {
		toSerialize["valid_values"] = o.ValidValues
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentUserDefinedFieldAttributesCreateRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Category     NullableIncidentUserDefinedFieldCategory  `json:"category,omitempty"`
		Collected    NullableIncidentUserDefinedFieldCollected `json:"collected,omitempty"`
		DefaultValue datadog.NullableString                    `json:"default_value,omitempty"`
		DisplayName  *string                                   `json:"display_name,omitempty"`
		Name         *string                                   `json:"name"`
		Ordinal      datadog.NullableString                    `json:"ordinal,omitempty"`
		Required     *bool                                     `json:"required,omitempty"`
		TagKey       datadog.NullableString                    `json:"tag_key,omitempty"`
		Type         *IncidentUserDefinedFieldFieldType        `json:"type"`
		ValidValues  []IncidentUserDefinedFieldValidValue      `json:"valid_values,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"category", "collected", "default_value", "display_name", "name", "ordinal", "required", "tag_key", "type", "valid_values"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Category.Get() != nil && !all.Category.Get().IsValid() {
		hasInvalidField = true
	} else {
		o.Category = all.Category
	}
	if all.Collected.Get() != nil && !all.Collected.Get().IsValid() {
		hasInvalidField = true
	} else {
		o.Collected = all.Collected
	}
	o.DefaultValue = all.DefaultValue
	o.DisplayName = all.DisplayName
	o.Name = *all.Name
	o.Ordinal = all.Ordinal
	o.Required = all.Required
	o.TagKey = all.TagKey
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.ValidValues = all.ValidValues

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
