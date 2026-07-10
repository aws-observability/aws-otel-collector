// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ModelLabProjectAttributes Attributes of a Model Lab project.
type ModelLabProjectAttributes struct {
	// The storage location for project artifacts.
	ArtifactStorageLocation string `json:"artifact_storage_location"`
	// The date and time the project was created.
	CreatedAt time.Time `json:"created_at"`
	// The date and time the project was soft-deleted.
	DeletedAt datadog.NullableTime `json:"deleted_at,omitempty"`
	// A description of the project.
	Description string `json:"description"`
	// An optional external URL associated with the project.
	ExternalUrl datadog.NullableString `json:"external_url,omitempty"`
	// Whether the project is starred by the current user.
	IsStarred bool `json:"is_starred"`
	// The name of the project.
	Name string `json:"name"`
	// The UUID of the project owner.
	OwnerId datadog.NullableString `json:"owner_id,omitempty"`
	// The list of tags associated with the project.
	Tags []ModelLabTag `json:"tags"`
	// The date and time the project was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewModelLabProjectAttributes instantiates a new ModelLabProjectAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewModelLabProjectAttributes(artifactStorageLocation string, createdAt time.Time, description string, isStarred bool, name string, tags []ModelLabTag, updatedAt time.Time) *ModelLabProjectAttributes {
	this := ModelLabProjectAttributes{}
	this.ArtifactStorageLocation = artifactStorageLocation
	this.CreatedAt = createdAt
	this.Description = description
	this.IsStarred = isStarred
	this.Name = name
	this.Tags = tags
	this.UpdatedAt = updatedAt
	return &this
}

// NewModelLabProjectAttributesWithDefaults instantiates a new ModelLabProjectAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewModelLabProjectAttributesWithDefaults() *ModelLabProjectAttributes {
	this := ModelLabProjectAttributes{}
	return &this
}

// GetArtifactStorageLocation returns the ArtifactStorageLocation field value.
func (o *ModelLabProjectAttributes) GetArtifactStorageLocation() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ArtifactStorageLocation
}

// GetArtifactStorageLocationOk returns a tuple with the ArtifactStorageLocation field value
// and a boolean to check if the value has been set.
func (o *ModelLabProjectAttributes) GetArtifactStorageLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ArtifactStorageLocation, true
}

// SetArtifactStorageLocation sets field value.
func (o *ModelLabProjectAttributes) SetArtifactStorageLocation(v string) {
	o.ArtifactStorageLocation = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *ModelLabProjectAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ModelLabProjectAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *ModelLabProjectAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelLabProjectAttributes) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ModelLabProjectAttributes) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *ModelLabProjectAttributes) HasDeletedAt() bool {
	return o != nil && o.DeletedAt.IsSet()
}

// SetDeletedAt gets a reference to the given datadog.NullableTime and assigns it to the DeletedAt field.
func (o *ModelLabProjectAttributes) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}

// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil.
func (o *ModelLabProjectAttributes) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil.
func (o *ModelLabProjectAttributes) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetDescription returns the Description field value.
func (o *ModelLabProjectAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ModelLabProjectAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *ModelLabProjectAttributes) SetDescription(v string) {
	o.Description = v
}

// GetExternalUrl returns the ExternalUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelLabProjectAttributes) GetExternalUrl() string {
	if o == nil || o.ExternalUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExternalUrl.Get()
}

// GetExternalUrlOk returns a tuple with the ExternalUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ModelLabProjectAttributes) GetExternalUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalUrl.Get(), o.ExternalUrl.IsSet()
}

// HasExternalUrl returns a boolean if a field has been set.
func (o *ModelLabProjectAttributes) HasExternalUrl() bool {
	return o != nil && o.ExternalUrl.IsSet()
}

// SetExternalUrl gets a reference to the given datadog.NullableString and assigns it to the ExternalUrl field.
func (o *ModelLabProjectAttributes) SetExternalUrl(v string) {
	o.ExternalUrl.Set(&v)
}

// SetExternalUrlNil sets the value for ExternalUrl to be an explicit nil.
func (o *ModelLabProjectAttributes) SetExternalUrlNil() {
	o.ExternalUrl.Set(nil)
}

// UnsetExternalUrl ensures that no value is present for ExternalUrl, not even an explicit nil.
func (o *ModelLabProjectAttributes) UnsetExternalUrl() {
	o.ExternalUrl.Unset()
}

// GetIsStarred returns the IsStarred field value.
func (o *ModelLabProjectAttributes) GetIsStarred() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsStarred
}

// GetIsStarredOk returns a tuple with the IsStarred field value
// and a boolean to check if the value has been set.
func (o *ModelLabProjectAttributes) GetIsStarredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsStarred, true
}

// SetIsStarred sets field value.
func (o *ModelLabProjectAttributes) SetIsStarred(v bool) {
	o.IsStarred = v
}

// GetName returns the Name field value.
func (o *ModelLabProjectAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ModelLabProjectAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ModelLabProjectAttributes) SetName(v string) {
	o.Name = v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelLabProjectAttributes) GetOwnerId() string {
	if o == nil || o.OwnerId.Get() == nil {
		var ret string
		return ret
	}
	return *o.OwnerId.Get()
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ModelLabProjectAttributes) GetOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OwnerId.Get(), o.OwnerId.IsSet()
}

// HasOwnerId returns a boolean if a field has been set.
func (o *ModelLabProjectAttributes) HasOwnerId() bool {
	return o != nil && o.OwnerId.IsSet()
}

// SetOwnerId gets a reference to the given datadog.NullableString and assigns it to the OwnerId field.
func (o *ModelLabProjectAttributes) SetOwnerId(v string) {
	o.OwnerId.Set(&v)
}

// SetOwnerIdNil sets the value for OwnerId to be an explicit nil.
func (o *ModelLabProjectAttributes) SetOwnerIdNil() {
	o.OwnerId.Set(nil)
}

// UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil.
func (o *ModelLabProjectAttributes) UnsetOwnerId() {
	o.OwnerId.Unset()
}

// GetTags returns the Tags field value.
func (o *ModelLabProjectAttributes) GetTags() []ModelLabTag {
	if o == nil {
		var ret []ModelLabTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *ModelLabProjectAttributes) GetTagsOk() (*[]ModelLabTag, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value.
func (o *ModelLabProjectAttributes) SetTags(v []ModelLabTag) {
	o.Tags = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *ModelLabProjectAttributes) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ModelLabProjectAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *ModelLabProjectAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ModelLabProjectAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["artifact_storage_location"] = o.ArtifactStorageLocation
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.DeletedAt.IsSet() {
		toSerialize["deleted_at"] = o.DeletedAt.Get()
	}
	toSerialize["description"] = o.Description
	if o.ExternalUrl.IsSet() {
		toSerialize["external_url"] = o.ExternalUrl.Get()
	}
	toSerialize["is_starred"] = o.IsStarred
	toSerialize["name"] = o.Name
	if o.OwnerId.IsSet() {
		toSerialize["owner_id"] = o.OwnerId.Get()
	}
	toSerialize["tags"] = o.Tags
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ModelLabProjectAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ArtifactStorageLocation *string                `json:"artifact_storage_location"`
		CreatedAt               *time.Time             `json:"created_at"`
		DeletedAt               datadog.NullableTime   `json:"deleted_at,omitempty"`
		Description             *string                `json:"description"`
		ExternalUrl             datadog.NullableString `json:"external_url,omitempty"`
		IsStarred               *bool                  `json:"is_starred"`
		Name                    *string                `json:"name"`
		OwnerId                 datadog.NullableString `json:"owner_id,omitempty"`
		Tags                    *[]ModelLabTag         `json:"tags"`
		UpdatedAt               *time.Time             `json:"updated_at"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ArtifactStorageLocation == nil {
		return fmt.Errorf("required field artifact_storage_location missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.IsStarred == nil {
		return fmt.Errorf("required field is_starred missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Tags == nil {
		return fmt.Errorf("required field tags missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"artifact_storage_location", "created_at", "deleted_at", "description", "external_url", "is_starred", "name", "owner_id", "tags", "updated_at"})
	} else {
		return err
	}
	o.ArtifactStorageLocation = *all.ArtifactStorageLocation
	o.CreatedAt = *all.CreatedAt
	o.DeletedAt = all.DeletedAt
	o.Description = *all.Description
	o.ExternalUrl = all.ExternalUrl
	o.IsStarred = *all.IsStarred
	o.Name = *all.Name
	o.OwnerId = all.OwnerId
	o.Tags = *all.Tags
	o.UpdatedAt = *all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
