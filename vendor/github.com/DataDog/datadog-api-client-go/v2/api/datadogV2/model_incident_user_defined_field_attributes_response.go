// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentUserDefinedFieldAttributesResponse Attributes of an incident user-defined field.
type IncidentUserDefinedFieldAttributesResponse struct {
	// The resource type this field is attached to. Always "incidents".
	AttachedTo string `json:"attached_to"`
	// The section in which the field appears: "what_happened" or "why_it_happened". When null, the field appears in the Attributes section.
	Category NullableIncidentUserDefinedFieldCategory `json:"category"`
	// The lifecycle stage at which the app prompts users to fill out this field. Cannot be set on required fields.
	Collected NullableIncidentUserDefinedFieldCollected `json:"collected"`
	// Timestamp when the field was created.
	Created time.Time `json:"created"`
	// The default value for the field.
	DefaultValue datadog.NullableString `json:"default_value"`
	// Timestamp when the field was soft-deleted, or null if not deleted.
	Deleted datadog.NullableTime `json:"deleted"`
	// The human-readable name shown in the UI.
	DisplayName string `json:"display_name"`
	// Metadata for autocomplete-type user-defined fields, describing how to populate autocomplete options.
	Metadata NullableIncidentUserDefinedFieldMetadata `json:"metadata"`
	// Timestamp when the field was last modified.
	Modified datadog.NullableTime `json:"modified"`
	// The unique identifier of the field.
	Name string `json:"name"`
	// A decimal string representing the field's display order in the UI.
	Ordinal datadog.NullableString `json:"ordinal"`
	// Reserved for future use. Always null.
	Prerequisite datadog.NullableString `json:"prerequisite"`
	// When true, users must fill out this field on incidents.
	Required bool `json:"required"`
	// When true, this field is reserved for system use and cannot be deleted.
	Reserved bool `json:"reserved"`
	// Reserved for internal use. Always 0.
	TableId int64 `json:"table_id"`
	// For metric tag-type fields only, the metric tag key that powers the autocomplete options.
	TagKey datadog.NullableString `json:"tag_key"`
	// The data type of the field. 1=dropdown, 2=multiselect, 3=textbox, 4=textarray, 5=metrictag, 6=autocomplete, 7=number, 8=datetime.
	Type datadog.NullableInt32 `json:"type"`
	// The list of allowed values for dropdown, multiselect, and autocomplete fields.
	ValidValues datadog.NullableList[IncidentUserDefinedFieldValidValue] `json:"valid_values"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentUserDefinedFieldAttributesResponse instantiates a new IncidentUserDefinedFieldAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentUserDefinedFieldAttributesResponse(attachedTo string, category NullableIncidentUserDefinedFieldCategory, collected NullableIncidentUserDefinedFieldCollected, created time.Time, defaultValue datadog.NullableString, deleted datadog.NullableTime, displayName string, metadata NullableIncidentUserDefinedFieldMetadata, modified datadog.NullableTime, name string, ordinal datadog.NullableString, prerequisite datadog.NullableString, required bool, reserved bool, tableId int64, tagKey datadog.NullableString, typeVar datadog.NullableInt32, validValues datadog.NullableList[IncidentUserDefinedFieldValidValue]) *IncidentUserDefinedFieldAttributesResponse {
	this := IncidentUserDefinedFieldAttributesResponse{}
	this.AttachedTo = attachedTo
	this.Category = category
	this.Collected = collected
	this.Created = created
	this.DefaultValue = defaultValue
	this.Deleted = deleted
	this.DisplayName = displayName
	this.Metadata = metadata
	this.Modified = modified
	this.Name = name
	this.Ordinal = ordinal
	this.Prerequisite = prerequisite
	this.Required = required
	this.Reserved = reserved
	this.TableId = tableId
	this.TagKey = tagKey
	this.Type = typeVar
	this.ValidValues = validValues
	return &this
}

// NewIncidentUserDefinedFieldAttributesResponseWithDefaults instantiates a new IncidentUserDefinedFieldAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentUserDefinedFieldAttributesResponseWithDefaults() *IncidentUserDefinedFieldAttributesResponse {
	this := IncidentUserDefinedFieldAttributesResponse{}
	return &this
}

// GetAttachedTo returns the AttachedTo field value.
func (o *IncidentUserDefinedFieldAttributesResponse) GetAttachedTo() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AttachedTo
}

// GetAttachedToOk returns a tuple with the AttachedTo field value
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedFieldAttributesResponse) GetAttachedToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttachedTo, true
}

// SetAttachedTo sets field value.
func (o *IncidentUserDefinedFieldAttributesResponse) SetAttachedTo(v string) {
	o.AttachedTo = v
}

// GetCategory returns the Category field value.
// If the value is explicit nil, the zero value for IncidentUserDefinedFieldCategory will be returned.
func (o *IncidentUserDefinedFieldAttributesResponse) GetCategory() IncidentUserDefinedFieldCategory {
	if o == nil || o.Category.Get() == nil {
		var ret IncidentUserDefinedFieldCategory
		return ret
	}
	return *o.Category.Get()
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentUserDefinedFieldAttributesResponse) GetCategoryOk() (*IncidentUserDefinedFieldCategory, bool) {
	if o == nil {
		return nil, false
	}
	return o.Category.Get(), o.Category.IsSet()
}

// SetCategory sets field value.
func (o *IncidentUserDefinedFieldAttributesResponse) SetCategory(v IncidentUserDefinedFieldCategory) {
	o.Category.Set(&v)
}

// GetCollected returns the Collected field value.
// If the value is explicit nil, the zero value for IncidentUserDefinedFieldCollected will be returned.
func (o *IncidentUserDefinedFieldAttributesResponse) GetCollected() IncidentUserDefinedFieldCollected {
	if o == nil || o.Collected.Get() == nil {
		var ret IncidentUserDefinedFieldCollected
		return ret
	}
	return *o.Collected.Get()
}

// GetCollectedOk returns a tuple with the Collected field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentUserDefinedFieldAttributesResponse) GetCollectedOk() (*IncidentUserDefinedFieldCollected, bool) {
	if o == nil {
		return nil, false
	}
	return o.Collected.Get(), o.Collected.IsSet()
}

// SetCollected sets field value.
func (o *IncidentUserDefinedFieldAttributesResponse) SetCollected(v IncidentUserDefinedFieldCollected) {
	o.Collected.Set(&v)
}

// GetCreated returns the Created field value.
func (o *IncidentUserDefinedFieldAttributesResponse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedFieldAttributesResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value.
func (o *IncidentUserDefinedFieldAttributesResponse) SetCreated(v time.Time) {
	o.Created = v
}

// GetDefaultValue returns the DefaultValue field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *IncidentUserDefinedFieldAttributesResponse) GetDefaultValue() string {
	if o == nil || o.DefaultValue.Get() == nil {
		var ret string
		return ret
	}
	return *o.DefaultValue.Get()
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentUserDefinedFieldAttributesResponse) GetDefaultValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultValue.Get(), o.DefaultValue.IsSet()
}

// SetDefaultValue sets field value.
func (o *IncidentUserDefinedFieldAttributesResponse) SetDefaultValue(v string) {
	o.DefaultValue.Set(&v)
}

// GetDeleted returns the Deleted field value.
// If the value is explicit nil, the zero value for time.Time will be returned.
func (o *IncidentUserDefinedFieldAttributesResponse) GetDeleted() time.Time {
	if o == nil || o.Deleted.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Deleted.Get()
}

// GetDeletedOk returns a tuple with the Deleted field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentUserDefinedFieldAttributesResponse) GetDeletedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Deleted.Get(), o.Deleted.IsSet()
}

// SetDeleted sets field value.
func (o *IncidentUserDefinedFieldAttributesResponse) SetDeleted(v time.Time) {
	o.Deleted.Set(&v)
}

// GetDisplayName returns the DisplayName field value.
func (o *IncidentUserDefinedFieldAttributesResponse) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedFieldAttributesResponse) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value.
func (o *IncidentUserDefinedFieldAttributesResponse) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetMetadata returns the Metadata field value.
// If the value is explicit nil, the zero value for IncidentUserDefinedFieldMetadata will be returned.
func (o *IncidentUserDefinedFieldAttributesResponse) GetMetadata() IncidentUserDefinedFieldMetadata {
	if o == nil || o.Metadata.Get() == nil {
		var ret IncidentUserDefinedFieldMetadata
		return ret
	}
	return *o.Metadata.Get()
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentUserDefinedFieldAttributesResponse) GetMetadataOk() (*IncidentUserDefinedFieldMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metadata.Get(), o.Metadata.IsSet()
}

// SetMetadata sets field value.
func (o *IncidentUserDefinedFieldAttributesResponse) SetMetadata(v IncidentUserDefinedFieldMetadata) {
	o.Metadata.Set(&v)
}

// GetModified returns the Modified field value.
// If the value is explicit nil, the zero value for time.Time will be returned.
func (o *IncidentUserDefinedFieldAttributesResponse) GetModified() time.Time {
	if o == nil || o.Modified.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Modified.Get()
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentUserDefinedFieldAttributesResponse) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Modified.Get(), o.Modified.IsSet()
}

// SetModified sets field value.
func (o *IncidentUserDefinedFieldAttributesResponse) SetModified(v time.Time) {
	o.Modified.Set(&v)
}

// GetName returns the Name field value.
func (o *IncidentUserDefinedFieldAttributesResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedFieldAttributesResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *IncidentUserDefinedFieldAttributesResponse) SetName(v string) {
	o.Name = v
}

// GetOrdinal returns the Ordinal field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *IncidentUserDefinedFieldAttributesResponse) GetOrdinal() string {
	if o == nil || o.Ordinal.Get() == nil {
		var ret string
		return ret
	}
	return *o.Ordinal.Get()
}

// GetOrdinalOk returns a tuple with the Ordinal field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentUserDefinedFieldAttributesResponse) GetOrdinalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ordinal.Get(), o.Ordinal.IsSet()
}

// SetOrdinal sets field value.
func (o *IncidentUserDefinedFieldAttributesResponse) SetOrdinal(v string) {
	o.Ordinal.Set(&v)
}

// GetPrerequisite returns the Prerequisite field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *IncidentUserDefinedFieldAttributesResponse) GetPrerequisite() string {
	if o == nil || o.Prerequisite.Get() == nil {
		var ret string
		return ret
	}
	return *o.Prerequisite.Get()
}

// GetPrerequisiteOk returns a tuple with the Prerequisite field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentUserDefinedFieldAttributesResponse) GetPrerequisiteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prerequisite.Get(), o.Prerequisite.IsSet()
}

// SetPrerequisite sets field value.
func (o *IncidentUserDefinedFieldAttributesResponse) SetPrerequisite(v string) {
	o.Prerequisite.Set(&v)
}

// GetRequired returns the Required field value.
func (o *IncidentUserDefinedFieldAttributesResponse) GetRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedFieldAttributesResponse) GetRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Required, true
}

// SetRequired sets field value.
func (o *IncidentUserDefinedFieldAttributesResponse) SetRequired(v bool) {
	o.Required = v
}

// GetReserved returns the Reserved field value.
func (o *IncidentUserDefinedFieldAttributesResponse) GetReserved() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Reserved
}

// GetReservedOk returns a tuple with the Reserved field value
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedFieldAttributesResponse) GetReservedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reserved, true
}

// SetReserved sets field value.
func (o *IncidentUserDefinedFieldAttributesResponse) SetReserved(v bool) {
	o.Reserved = v
}

// GetTableId returns the TableId field value.
func (o *IncidentUserDefinedFieldAttributesResponse) GetTableId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TableId
}

// GetTableIdOk returns a tuple with the TableId field value
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedFieldAttributesResponse) GetTableIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TableId, true
}

// SetTableId sets field value.
func (o *IncidentUserDefinedFieldAttributesResponse) SetTableId(v int64) {
	o.TableId = v
}

// GetTagKey returns the TagKey field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *IncidentUserDefinedFieldAttributesResponse) GetTagKey() string {
	if o == nil || o.TagKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.TagKey.Get()
}

// GetTagKeyOk returns a tuple with the TagKey field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentUserDefinedFieldAttributesResponse) GetTagKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TagKey.Get(), o.TagKey.IsSet()
}

// SetTagKey sets field value.
func (o *IncidentUserDefinedFieldAttributesResponse) SetTagKey(v string) {
	o.TagKey.Set(&v)
}

// GetType returns the Type field value.
// If the value is explicit nil, the zero value for int32 will be returned.
func (o *IncidentUserDefinedFieldAttributesResponse) GetType() int32 {
	if o == nil || o.Type.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentUserDefinedFieldAttributesResponse) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// SetType sets field value.
func (o *IncidentUserDefinedFieldAttributesResponse) SetType(v int32) {
	o.Type.Set(&v)
}

// GetValidValues returns the ValidValues field value.
// If the value is explicit nil, the zero value for []IncidentUserDefinedFieldValidValue will be returned.
func (o *IncidentUserDefinedFieldAttributesResponse) GetValidValues() []IncidentUserDefinedFieldValidValue {
	if o == nil {
		var ret []IncidentUserDefinedFieldValidValue
		return ret
	}
	return *o.ValidValues.Get()
}

// GetValidValuesOk returns a tuple with the ValidValues field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentUserDefinedFieldAttributesResponse) GetValidValuesOk() (*[]IncidentUserDefinedFieldValidValue, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValidValues.Get(), o.ValidValues.IsSet()
}

// SetValidValues sets field value.
func (o *IncidentUserDefinedFieldAttributesResponse) SetValidValues(v []IncidentUserDefinedFieldValidValue) {
	o.ValidValues.Set(&v)
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentUserDefinedFieldAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attached_to"] = o.AttachedTo
	toSerialize["category"] = o.Category.Get()
	toSerialize["collected"] = o.Collected.Get()
	if o.Created.Nanosecond() == 0 {
		toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["default_value"] = o.DefaultValue.Get()
	toSerialize["deleted"] = o.Deleted.Get()
	toSerialize["display_name"] = o.DisplayName
	toSerialize["metadata"] = o.Metadata.Get()
	toSerialize["modified"] = o.Modified.Get()
	toSerialize["name"] = o.Name
	toSerialize["ordinal"] = o.Ordinal.Get()
	toSerialize["prerequisite"] = o.Prerequisite.Get()
	toSerialize["required"] = o.Required
	toSerialize["reserved"] = o.Reserved
	toSerialize["table_id"] = o.TableId
	toSerialize["tag_key"] = o.TagKey.Get()
	toSerialize["type"] = o.Type.Get()
	toSerialize["valid_values"] = o.ValidValues.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentUserDefinedFieldAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AttachedTo   *string                                                  `json:"attached_to"`
		Category     NullableIncidentUserDefinedFieldCategory                 `json:"category"`
		Collected    NullableIncidentUserDefinedFieldCollected                `json:"collected"`
		Created      *time.Time                                               `json:"created"`
		DefaultValue datadog.NullableString                                   `json:"default_value"`
		Deleted      datadog.NullableTime                                     `json:"deleted"`
		DisplayName  *string                                                  `json:"display_name"`
		Metadata     NullableIncidentUserDefinedFieldMetadata                 `json:"metadata"`
		Modified     datadog.NullableTime                                     `json:"modified"`
		Name         *string                                                  `json:"name"`
		Ordinal      datadog.NullableString                                   `json:"ordinal"`
		Prerequisite datadog.NullableString                                   `json:"prerequisite"`
		Required     *bool                                                    `json:"required"`
		Reserved     *bool                                                    `json:"reserved"`
		TableId      *int64                                                   `json:"table_id"`
		TagKey       datadog.NullableString                                   `json:"tag_key"`
		Type         datadog.NullableInt32                                    `json:"type"`
		ValidValues  datadog.NullableList[IncidentUserDefinedFieldValidValue] `json:"valid_values"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AttachedTo == nil {
		return fmt.Errorf("required field attached_to missing")
	}
	if !all.Category.IsSet() {
		return fmt.Errorf("required field category missing")
	}
	if !all.Collected.IsSet() {
		return fmt.Errorf("required field collected missing")
	}
	if all.Created == nil {
		return fmt.Errorf("required field created missing")
	}
	if !all.DefaultValue.IsSet() {
		return fmt.Errorf("required field default_value missing")
	}
	if !all.Deleted.IsSet() {
		return fmt.Errorf("required field deleted missing")
	}
	if all.DisplayName == nil {
		return fmt.Errorf("required field display_name missing")
	}
	if !all.Metadata.IsSet() {
		return fmt.Errorf("required field metadata missing")
	}
	if !all.Modified.IsSet() {
		return fmt.Errorf("required field modified missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if !all.Ordinal.IsSet() {
		return fmt.Errorf("required field ordinal missing")
	}
	if !all.Prerequisite.IsSet() {
		return fmt.Errorf("required field prerequisite missing")
	}
	if all.Required == nil {
		return fmt.Errorf("required field required missing")
	}
	if all.Reserved == nil {
		return fmt.Errorf("required field reserved missing")
	}
	if all.TableId == nil {
		return fmt.Errorf("required field table_id missing")
	}
	if !all.TagKey.IsSet() {
		return fmt.Errorf("required field tag_key missing")
	}
	if !all.Type.IsSet() {
		return fmt.Errorf("required field type missing")
	}
	if !all.ValidValues.IsSet() {
		return fmt.Errorf("required field valid_values missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attached_to", "category", "collected", "created", "default_value", "deleted", "display_name", "metadata", "modified", "name", "ordinal", "prerequisite", "required", "reserved", "table_id", "tag_key", "type", "valid_values"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AttachedTo = *all.AttachedTo
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
	o.Created = *all.Created
	o.DefaultValue = all.DefaultValue
	o.Deleted = all.Deleted
	o.DisplayName = *all.DisplayName
	o.Metadata = all.Metadata
	o.Modified = all.Modified
	o.Name = *all.Name
	o.Ordinal = all.Ordinal
	o.Prerequisite = all.Prerequisite
	o.Required = *all.Required
	o.Reserved = *all.Reserved
	o.TableId = *all.TableId
	o.TagKey = all.TagKey
	o.Type = all.Type
	o.ValidValues = all.ValidValues

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
