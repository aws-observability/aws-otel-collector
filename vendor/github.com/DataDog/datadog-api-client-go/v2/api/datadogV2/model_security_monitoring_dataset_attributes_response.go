// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringDatasetAttributesResponse The attributes of a Cloud SIEM dataset.
type SecurityMonitoringDatasetAttributesResponse struct {
	// The creation timestamp of the dataset, in ISO 8601 format.
	CreatedAt string `json:"createdAt"`
	// The Datadog handle of the user who created the dataset.
	CreatedByHandle string `json:"createdByHandle"`
	// The display name of the user who created the dataset.
	CreatedByName string `json:"createdByName"`
	// The definition of the dataset. The shape depends on the value of `data_source`.
	// Use `reference_table` or `managed_resource` for a referential dataset, or one of the
	// event platform sources (for example `logs`, `audit`, `events`, `spans`, `rum`) for
	// an event platform dataset.
	Definition SecurityMonitoringDatasetDefinition `json:"definition"`
	// The description of the dataset.
	Description string `json:"description"`
	// The UUID of the dataset.
	Id string `json:"id"`
	// Whether the dataset is an out-of-the-box dataset provided by Datadog.
	IsDefault bool `json:"isDefault"`
	// Whether the dataset is marked as deprecated.
	IsDeprecated bool `json:"isDeprecated"`
	// The timestamp of the last modification of the dataset, in ISO 8601 format.
	ModifiedAt string `json:"modifiedAt"`
	// The unique name of the dataset.
	Name string `json:"name"`
	// The Datadog handle of the user who last updated the dataset.
	UpdatedByHandle datadog.NullableString `json:"updatedByHandle"`
	// The display name of the user who last updated the dataset.
	UpdatedByName datadog.NullableString `json:"updatedByName"`
	// The current version of the dataset.
	Version int64 `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringDatasetAttributesResponse instantiates a new SecurityMonitoringDatasetAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringDatasetAttributesResponse(createdAt string, createdByHandle string, createdByName string, definition SecurityMonitoringDatasetDefinition, description string, id string, isDefault bool, isDeprecated bool, modifiedAt string, name string, updatedByHandle datadog.NullableString, updatedByName datadog.NullableString, version int64) *SecurityMonitoringDatasetAttributesResponse {
	this := SecurityMonitoringDatasetAttributesResponse{}
	this.CreatedAt = createdAt
	this.CreatedByHandle = createdByHandle
	this.CreatedByName = createdByName
	this.Definition = definition
	this.Description = description
	this.Id = id
	this.IsDefault = isDefault
	this.IsDeprecated = isDeprecated
	this.ModifiedAt = modifiedAt
	this.Name = name
	this.UpdatedByHandle = updatedByHandle
	this.UpdatedByName = updatedByName
	this.Version = version
	return &this
}

// NewSecurityMonitoringDatasetAttributesResponseWithDefaults instantiates a new SecurityMonitoringDatasetAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringDatasetAttributesResponseWithDefaults() *SecurityMonitoringDatasetAttributesResponse {
	this := SecurityMonitoringDatasetAttributesResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *SecurityMonitoringDatasetAttributesResponse) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetAttributesResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *SecurityMonitoringDatasetAttributesResponse) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetCreatedByHandle returns the CreatedByHandle field value.
func (o *SecurityMonitoringDatasetAttributesResponse) GetCreatedByHandle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedByHandle
}

// GetCreatedByHandleOk returns a tuple with the CreatedByHandle field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetAttributesResponse) GetCreatedByHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedByHandle, true
}

// SetCreatedByHandle sets field value.
func (o *SecurityMonitoringDatasetAttributesResponse) SetCreatedByHandle(v string) {
	o.CreatedByHandle = v
}

// GetCreatedByName returns the CreatedByName field value.
func (o *SecurityMonitoringDatasetAttributesResponse) GetCreatedByName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedByName
}

// GetCreatedByNameOk returns a tuple with the CreatedByName field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetAttributesResponse) GetCreatedByNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedByName, true
}

// SetCreatedByName sets field value.
func (o *SecurityMonitoringDatasetAttributesResponse) SetCreatedByName(v string) {
	o.CreatedByName = v
}

// GetDefinition returns the Definition field value.
func (o *SecurityMonitoringDatasetAttributesResponse) GetDefinition() SecurityMonitoringDatasetDefinition {
	if o == nil {
		var ret SecurityMonitoringDatasetDefinition
		return ret
	}
	return o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetAttributesResponse) GetDefinitionOk() (*SecurityMonitoringDatasetDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Definition, true
}

// SetDefinition sets field value.
func (o *SecurityMonitoringDatasetAttributesResponse) SetDefinition(v SecurityMonitoringDatasetDefinition) {
	o.Definition = v
}

// GetDescription returns the Description field value.
func (o *SecurityMonitoringDatasetAttributesResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetAttributesResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *SecurityMonitoringDatasetAttributesResponse) SetDescription(v string) {
	o.Description = v
}

// GetId returns the Id field value.
func (o *SecurityMonitoringDatasetAttributesResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetAttributesResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *SecurityMonitoringDatasetAttributesResponse) SetId(v string) {
	o.Id = v
}

// GetIsDefault returns the IsDefault field value.
func (o *SecurityMonitoringDatasetAttributesResponse) GetIsDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetAttributesResponse) GetIsDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDefault, true
}

// SetIsDefault sets field value.
func (o *SecurityMonitoringDatasetAttributesResponse) SetIsDefault(v bool) {
	o.IsDefault = v
}

// GetIsDeprecated returns the IsDeprecated field value.
func (o *SecurityMonitoringDatasetAttributesResponse) GetIsDeprecated() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsDeprecated
}

// GetIsDeprecatedOk returns a tuple with the IsDeprecated field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetAttributesResponse) GetIsDeprecatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDeprecated, true
}

// SetIsDeprecated sets field value.
func (o *SecurityMonitoringDatasetAttributesResponse) SetIsDeprecated(v bool) {
	o.IsDeprecated = v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *SecurityMonitoringDatasetAttributesResponse) GetModifiedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetAttributesResponse) GetModifiedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *SecurityMonitoringDatasetAttributesResponse) SetModifiedAt(v string) {
	o.ModifiedAt = v
}

// GetName returns the Name field value.
func (o *SecurityMonitoringDatasetAttributesResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetAttributesResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SecurityMonitoringDatasetAttributesResponse) SetName(v string) {
	o.Name = v
}

// GetUpdatedByHandle returns the UpdatedByHandle field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *SecurityMonitoringDatasetAttributesResponse) GetUpdatedByHandle() string {
	if o == nil || o.UpdatedByHandle.Get() == nil {
		var ret string
		return ret
	}
	return *o.UpdatedByHandle.Get()
}

// GetUpdatedByHandleOk returns a tuple with the UpdatedByHandle field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SecurityMonitoringDatasetAttributesResponse) GetUpdatedByHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedByHandle.Get(), o.UpdatedByHandle.IsSet()
}

// SetUpdatedByHandle sets field value.
func (o *SecurityMonitoringDatasetAttributesResponse) SetUpdatedByHandle(v string) {
	o.UpdatedByHandle.Set(&v)
}

// GetUpdatedByName returns the UpdatedByName field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *SecurityMonitoringDatasetAttributesResponse) GetUpdatedByName() string {
	if o == nil || o.UpdatedByName.Get() == nil {
		var ret string
		return ret
	}
	return *o.UpdatedByName.Get()
}

// GetUpdatedByNameOk returns a tuple with the UpdatedByName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SecurityMonitoringDatasetAttributesResponse) GetUpdatedByNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedByName.Get(), o.UpdatedByName.IsSet()
}

// SetUpdatedByName sets field value.
func (o *SecurityMonitoringDatasetAttributesResponse) SetUpdatedByName(v string) {
	o.UpdatedByName.Set(&v)
}

// GetVersion returns the Version field value.
func (o *SecurityMonitoringDatasetAttributesResponse) GetVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetAttributesResponse) GetVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *SecurityMonitoringDatasetAttributesResponse) SetVersion(v int64) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringDatasetAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["createdByHandle"] = o.CreatedByHandle
	toSerialize["createdByName"] = o.CreatedByName
	toSerialize["definition"] = o.Definition
	toSerialize["description"] = o.Description
	toSerialize["id"] = o.Id
	toSerialize["isDefault"] = o.IsDefault
	toSerialize["isDeprecated"] = o.IsDeprecated
	toSerialize["modifiedAt"] = o.ModifiedAt
	toSerialize["name"] = o.Name
	toSerialize["updatedByHandle"] = o.UpdatedByHandle.Get()
	toSerialize["updatedByName"] = o.UpdatedByName.Get()
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringDatasetAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt       *string                              `json:"createdAt"`
		CreatedByHandle *string                              `json:"createdByHandle"`
		CreatedByName   *string                              `json:"createdByName"`
		Definition      *SecurityMonitoringDatasetDefinition `json:"definition"`
		Description     *string                              `json:"description"`
		Id              *string                              `json:"id"`
		IsDefault       *bool                                `json:"isDefault"`
		IsDeprecated    *bool                                `json:"isDeprecated"`
		ModifiedAt      *string                              `json:"modifiedAt"`
		Name            *string                              `json:"name"`
		UpdatedByHandle datadog.NullableString               `json:"updatedByHandle"`
		UpdatedByName   datadog.NullableString               `json:"updatedByName"`
		Version         *int64                               `json:"version"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	if all.CreatedByHandle == nil {
		return fmt.Errorf("required field createdByHandle missing")
	}
	if all.CreatedByName == nil {
		return fmt.Errorf("required field createdByName missing")
	}
	if all.Definition == nil {
		return fmt.Errorf("required field definition missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.IsDefault == nil {
		return fmt.Errorf("required field isDefault missing")
	}
	if all.IsDeprecated == nil {
		return fmt.Errorf("required field isDeprecated missing")
	}
	if all.ModifiedAt == nil {
		return fmt.Errorf("required field modifiedAt missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if !all.UpdatedByHandle.IsSet() {
		return fmt.Errorf("required field updatedByHandle missing")
	}
	if !all.UpdatedByName.IsSet() {
		return fmt.Errorf("required field updatedByName missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"createdAt", "createdByHandle", "createdByName", "definition", "description", "id", "isDefault", "isDeprecated", "modifiedAt", "name", "updatedByHandle", "updatedByName", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = *all.CreatedAt
	o.CreatedByHandle = *all.CreatedByHandle
	o.CreatedByName = *all.CreatedByName
	if all.Definition.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Definition = *all.Definition
	o.Description = *all.Description
	o.Id = *all.Id
	o.IsDefault = *all.IsDefault
	o.IsDeprecated = *all.IsDeprecated
	o.ModifiedAt = *all.ModifiedAt
	o.Name = *all.Name
	o.UpdatedByHandle = all.UpdatedByHandle
	o.UpdatedByName = all.UpdatedByName
	o.Version = *all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
