// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BlueprintMetadataAttributes The attributes of a blueprint metadata resource.
type BlueprintMetadataAttributes struct {
	// The timestamp when the blueprint was created.
	CreatedAt time.Time `json:"created_at"`
	// A description of what the blueprint does.
	Description string `json:"description"`
	// The human-readable name of the blueprint.
	Name string `json:"name"`
	// The unique slug identifier of the blueprint.
	Slug string `json:"slug"`
	// Tags associated with the blueprint.
	Tags []string `json:"tags,omitempty"`
	// The background style of the blueprint tile.
	TileBackground *string `json:"tile_background,omitempty"`
	// The fully qualified name of the action used as the tile icon.
	TileIconActionFqn *string `json:"tile_icon_action_fqn,omitempty"`
	// The timestamp when the blueprint was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBlueprintMetadataAttributes instantiates a new BlueprintMetadataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBlueprintMetadataAttributes(createdAt time.Time, description string, name string, slug string, updatedAt time.Time) *BlueprintMetadataAttributes {
	this := BlueprintMetadataAttributes{}
	this.CreatedAt = createdAt
	this.Description = description
	this.Name = name
	this.Slug = slug
	this.UpdatedAt = updatedAt
	return &this
}

// NewBlueprintMetadataAttributesWithDefaults instantiates a new BlueprintMetadataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBlueprintMetadataAttributesWithDefaults() *BlueprintMetadataAttributes {
	this := BlueprintMetadataAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *BlueprintMetadataAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *BlueprintMetadataAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *BlueprintMetadataAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDescription returns the Description field value.
func (o *BlueprintMetadataAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *BlueprintMetadataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *BlueprintMetadataAttributes) SetDescription(v string) {
	o.Description = v
}

// GetName returns the Name field value.
func (o *BlueprintMetadataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BlueprintMetadataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *BlueprintMetadataAttributes) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value.
func (o *BlueprintMetadataAttributes) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *BlueprintMetadataAttributes) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value.
func (o *BlueprintMetadataAttributes) SetSlug(v string) {
	o.Slug = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *BlueprintMetadataAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintMetadataAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *BlueprintMetadataAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *BlueprintMetadataAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetTileBackground returns the TileBackground field value if set, zero value otherwise.
func (o *BlueprintMetadataAttributes) GetTileBackground() string {
	if o == nil || o.TileBackground == nil {
		var ret string
		return ret
	}
	return *o.TileBackground
}

// GetTileBackgroundOk returns a tuple with the TileBackground field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintMetadataAttributes) GetTileBackgroundOk() (*string, bool) {
	if o == nil || o.TileBackground == nil {
		return nil, false
	}
	return o.TileBackground, true
}

// HasTileBackground returns a boolean if a field has been set.
func (o *BlueprintMetadataAttributes) HasTileBackground() bool {
	return o != nil && o.TileBackground != nil
}

// SetTileBackground gets a reference to the given string and assigns it to the TileBackground field.
func (o *BlueprintMetadataAttributes) SetTileBackground(v string) {
	o.TileBackground = &v
}

// GetTileIconActionFqn returns the TileIconActionFqn field value if set, zero value otherwise.
func (o *BlueprintMetadataAttributes) GetTileIconActionFqn() string {
	if o == nil || o.TileIconActionFqn == nil {
		var ret string
		return ret
	}
	return *o.TileIconActionFqn
}

// GetTileIconActionFqnOk returns a tuple with the TileIconActionFqn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintMetadataAttributes) GetTileIconActionFqnOk() (*string, bool) {
	if o == nil || o.TileIconActionFqn == nil {
		return nil, false
	}
	return o.TileIconActionFqn, true
}

// HasTileIconActionFqn returns a boolean if a field has been set.
func (o *BlueprintMetadataAttributes) HasTileIconActionFqn() bool {
	return o != nil && o.TileIconActionFqn != nil
}

// SetTileIconActionFqn gets a reference to the given string and assigns it to the TileIconActionFqn field.
func (o *BlueprintMetadataAttributes) SetTileIconActionFqn(v string) {
	o.TileIconActionFqn = &v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *BlueprintMetadataAttributes) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *BlueprintMetadataAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *BlueprintMetadataAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o BlueprintMetadataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["description"] = o.Description
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.TileBackground != nil {
		toSerialize["tile_background"] = o.TileBackground
	}
	if o.TileIconActionFqn != nil {
		toSerialize["tile_icon_action_fqn"] = o.TileIconActionFqn
	}
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
func (o *BlueprintMetadataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt         *time.Time `json:"created_at"`
		Description       *string    `json:"description"`
		Name              *string    `json:"name"`
		Slug              *string    `json:"slug"`
		Tags              []string   `json:"tags,omitempty"`
		TileBackground    *string    `json:"tile_background,omitempty"`
		TileIconActionFqn *string    `json:"tile_icon_action_fqn,omitempty"`
		UpdatedAt         *time.Time `json:"updated_at"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Slug == nil {
		return fmt.Errorf("required field slug missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "description", "name", "slug", "tags", "tile_background", "tile_icon_action_fqn", "updated_at"})
	} else {
		return err
	}
	o.CreatedAt = *all.CreatedAt
	o.Description = *all.Description
	o.Name = *all.Name
	o.Slug = *all.Slug
	o.Tags = all.Tags
	o.TileBackground = all.TileBackground
	o.TileIconActionFqn = all.TileIconActionFqn
	o.UpdatedAt = *all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
