// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentUserDefinedFieldAttributesUpdateRequest Attributes for updating an incident user-defined field. All fields are optional.
type IncidentUserDefinedFieldAttributesUpdateRequest struct {
	// The section in which the field appears: "what_happened" or "why_it_happened". When null, the field appears in the Attributes section.
	Category NullableIncidentUserDefinedFieldCategory `json:"category,omitempty"`
	// The lifecycle stage at which the app prompts users to fill out this field. Cannot be set on required fields.
	Collected NullableIncidentUserDefinedFieldCollected `json:"collected,omitempty"`
	// The default value for the field. Must be one of the valid values when valid_values is set.
	DefaultValue datadog.NullableString `json:"default_value,omitempty"`
	// The human-readable name shown in the UI.
	DisplayName *string `json:"display_name,omitempty"`
	// A decimal string representing the field's display order in the UI.
	Ordinal datadog.NullableString `json:"ordinal,omitempty"`
	// When true, users must fill out this field on incidents.
	Required datadog.NullableBool `json:"required,omitempty"`
	// The list of allowed values for dropdown and multiselect fields. Limited to 1000 values.
	ValidValues []IncidentUserDefinedFieldValidValue `json:"valid_values,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentUserDefinedFieldAttributesUpdateRequest instantiates a new IncidentUserDefinedFieldAttributesUpdateRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentUserDefinedFieldAttributesUpdateRequest() *IncidentUserDefinedFieldAttributesUpdateRequest {
	this := IncidentUserDefinedFieldAttributesUpdateRequest{}
	return &this
}

// NewIncidentUserDefinedFieldAttributesUpdateRequestWithDefaults instantiates a new IncidentUserDefinedFieldAttributesUpdateRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentUserDefinedFieldAttributesUpdateRequestWithDefaults() *IncidentUserDefinedFieldAttributesUpdateRequest {
	this := IncidentUserDefinedFieldAttributesUpdateRequest{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentUserDefinedFieldAttributesUpdateRequest) GetCategory() IncidentUserDefinedFieldCategory {
	if o == nil || o.Category.Get() == nil {
		var ret IncidentUserDefinedFieldCategory
		return ret
	}
	return *o.Category.Get()
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentUserDefinedFieldAttributesUpdateRequest) GetCategoryOk() (*IncidentUserDefinedFieldCategory, bool) {
	if o == nil {
		return nil, false
	}
	return o.Category.Get(), o.Category.IsSet()
}

// HasCategory returns a boolean if a field has been set.
func (o *IncidentUserDefinedFieldAttributesUpdateRequest) HasCategory() bool {
	return o != nil && o.Category.IsSet()
}

// SetCategory gets a reference to the given NullableIncidentUserDefinedFieldCategory and assigns it to the Category field.
func (o *IncidentUserDefinedFieldAttributesUpdateRequest) SetCategory(v IncidentUserDefinedFieldCategory) {
	o.Category.Set(&v)
}

// SetCategoryNil sets the value for Category to be an explicit nil.
func (o *IncidentUserDefinedFieldAttributesUpdateRequest) SetCategoryNil() {
	o.Category.Set(nil)
}

// UnsetCategory ensures that no value is present for Category, not even an explicit nil.
func (o *IncidentUserDefinedFieldAttributesUpdateRequest) UnsetCategory() {
	o.Category.Unset()
}

// GetCollected returns the Collected field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentUserDefinedFieldAttributesUpdateRequest) GetCollected() IncidentUserDefinedFieldCollected {
	if o == nil || o.Collected.Get() == nil {
		var ret IncidentUserDefinedFieldCollected
		return ret
	}
	return *o.Collected.Get()
}

// GetCollectedOk returns a tuple with the Collected field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentUserDefinedFieldAttributesUpdateRequest) GetCollectedOk() (*IncidentUserDefinedFieldCollected, bool) {
	if o == nil {
		return nil, false
	}
	return o.Collected.Get(), o.Collected.IsSet()
}

// HasCollected returns a boolean if a field has been set.
func (o *IncidentUserDefinedFieldAttributesUpdateRequest) HasCollected() bool {
	return o != nil && o.Collected.IsSet()
}

// SetCollected gets a reference to the given NullableIncidentUserDefinedFieldCollected and assigns it to the Collected field.
func (o *IncidentUserDefinedFieldAttributesUpdateRequest) SetCollected(v IncidentUserDefinedFieldCollected) {
	o.Collected.Set(&v)
}

// SetCollectedNil sets the value for Collected to be an explicit nil.
func (o *IncidentUserDefinedFieldAttributesUpdateRequest) SetCollectedNil() {
	o.Collected.Set(nil)
}

// UnsetCollected ensures that no value is present for Collected, not even an explicit nil.
func (o *IncidentUserDefinedFieldAttributesUpdateRequest) UnsetCollected() {
	o.Collected.Unset()
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentUserDefinedFieldAttributesUpdateRequest) GetDefaultValue() string {
	if o == nil || o.DefaultValue.Get() == nil {
		var ret string
		return ret
	}
	return *o.DefaultValue.Get()
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentUserDefinedFieldAttributesUpdateRequest) GetDefaultValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultValue.Get(), o.DefaultValue.IsSet()
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *IncidentUserDefinedFieldAttributesUpdateRequest) HasDefaultValue() bool {
	return o != nil && o.DefaultValue.IsSet()
}

// SetDefaultValue gets a reference to the given datadog.NullableString and assigns it to the DefaultValue field.
func (o *IncidentUserDefinedFieldAttributesUpdateRequest) SetDefaultValue(v string) {
	o.DefaultValue.Set(&v)
}

// SetDefaultValueNil sets the value for DefaultValue to be an explicit nil.
func (o *IncidentUserDefinedFieldAttributesUpdateRequest) SetDefaultValueNil() {
	o.DefaultValue.Set(nil)
}

// UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil.
func (o *IncidentUserDefinedFieldAttributesUpdateRequest) UnsetDefaultValue() {
	o.DefaultValue.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *IncidentUserDefinedFieldAttributesUpdateRequest) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedFieldAttributesUpdateRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *IncidentUserDefinedFieldAttributesUpdateRequest) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *IncidentUserDefinedFieldAttributesUpdateRequest) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetOrdinal returns the Ordinal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentUserDefinedFieldAttributesUpdateRequest) GetOrdinal() string {
	if o == nil || o.Ordinal.Get() == nil {
		var ret string
		return ret
	}
	return *o.Ordinal.Get()
}

// GetOrdinalOk returns a tuple with the Ordinal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentUserDefinedFieldAttributesUpdateRequest) GetOrdinalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ordinal.Get(), o.Ordinal.IsSet()
}

// HasOrdinal returns a boolean if a field has been set.
func (o *IncidentUserDefinedFieldAttributesUpdateRequest) HasOrdinal() bool {
	return o != nil && o.Ordinal.IsSet()
}

// SetOrdinal gets a reference to the given datadog.NullableString and assigns it to the Ordinal field.
func (o *IncidentUserDefinedFieldAttributesUpdateRequest) SetOrdinal(v string) {
	o.Ordinal.Set(&v)
}

// SetOrdinalNil sets the value for Ordinal to be an explicit nil.
func (o *IncidentUserDefinedFieldAttributesUpdateRequest) SetOrdinalNil() {
	o.Ordinal.Set(nil)
}

// UnsetOrdinal ensures that no value is present for Ordinal, not even an explicit nil.
func (o *IncidentUserDefinedFieldAttributesUpdateRequest) UnsetOrdinal() {
	o.Ordinal.Unset()
}

// GetRequired returns the Required field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentUserDefinedFieldAttributesUpdateRequest) GetRequired() bool {
	if o == nil || o.Required.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Required.Get()
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentUserDefinedFieldAttributesUpdateRequest) GetRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Required.Get(), o.Required.IsSet()
}

// HasRequired returns a boolean if a field has been set.
func (o *IncidentUserDefinedFieldAttributesUpdateRequest) HasRequired() bool {
	return o != nil && o.Required.IsSet()
}

// SetRequired gets a reference to the given datadog.NullableBool and assigns it to the Required field.
func (o *IncidentUserDefinedFieldAttributesUpdateRequest) SetRequired(v bool) {
	o.Required.Set(&v)
}

// SetRequiredNil sets the value for Required to be an explicit nil.
func (o *IncidentUserDefinedFieldAttributesUpdateRequest) SetRequiredNil() {
	o.Required.Set(nil)
}

// UnsetRequired ensures that no value is present for Required, not even an explicit nil.
func (o *IncidentUserDefinedFieldAttributesUpdateRequest) UnsetRequired() {
	o.Required.Unset()
}

// GetValidValues returns the ValidValues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentUserDefinedFieldAttributesUpdateRequest) GetValidValues() []IncidentUserDefinedFieldValidValue {
	if o == nil {
		var ret []IncidentUserDefinedFieldValidValue
		return ret
	}
	return o.ValidValues
}

// GetValidValuesOk returns a tuple with the ValidValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentUserDefinedFieldAttributesUpdateRequest) GetValidValuesOk() (*[]IncidentUserDefinedFieldValidValue, bool) {
	if o == nil || o.ValidValues == nil {
		return nil, false
	}
	return &o.ValidValues, true
}

// HasValidValues returns a boolean if a field has been set.
func (o *IncidentUserDefinedFieldAttributesUpdateRequest) HasValidValues() bool {
	return o != nil && o.ValidValues != nil
}

// SetValidValues gets a reference to the given []IncidentUserDefinedFieldValidValue and assigns it to the ValidValues field.
func (o *IncidentUserDefinedFieldAttributesUpdateRequest) SetValidValues(v []IncidentUserDefinedFieldValidValue) {
	o.ValidValues = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentUserDefinedFieldAttributesUpdateRequest) MarshalJSON() ([]byte, error) {
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
	if o.Ordinal.IsSet() {
		toSerialize["ordinal"] = o.Ordinal.Get()
	}
	if o.Required.IsSet() {
		toSerialize["required"] = o.Required.Get()
	}
	if o.ValidValues != nil {
		toSerialize["valid_values"] = o.ValidValues
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentUserDefinedFieldAttributesUpdateRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Category     NullableIncidentUserDefinedFieldCategory  `json:"category,omitempty"`
		Collected    NullableIncidentUserDefinedFieldCollected `json:"collected,omitempty"`
		DefaultValue datadog.NullableString                    `json:"default_value,omitempty"`
		DisplayName  *string                                   `json:"display_name,omitempty"`
		Ordinal      datadog.NullableString                    `json:"ordinal,omitempty"`
		Required     datadog.NullableBool                      `json:"required,omitempty"`
		ValidValues  []IncidentUserDefinedFieldValidValue      `json:"valid_values,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"category", "collected", "default_value", "display_name", "ordinal", "required", "valid_values"})
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
	o.Ordinal = all.Ordinal
	o.Required = all.Required
	o.ValidValues = all.ValidValues

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
