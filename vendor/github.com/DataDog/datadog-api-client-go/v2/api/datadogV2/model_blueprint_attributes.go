// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BlueprintAttributes The attributes of a blueprint resource.
type BlueprintAttributes struct {
	// The timestamp when the blueprint was created.
	CreatedAt time.Time `json:"created_at"`
	// The app definition type.
	Definition AppDefinitionType `json:"definition"`
	// A description of what the blueprint does.
	Description string `json:"description"`
	// Embedded datastore blueprints.
	EmbeddedDatastoreBlueprints map[string]interface{} `json:"embedded_datastore_blueprints,omitempty"`
	// Embedded native actions.
	EmbeddedNativeActions []map[string]interface{} `json:"embedded_native_actions,omitempty"`
	// Embedded workflow blueprints.
	EmbeddedWorkflowBlueprints map[string]interface{} `json:"embedded_workflow_blueprints,omitempty"`
	// The integration ID associated with the blueprint.
	IntegrationId *string `json:"integration_id,omitempty"`
	// Mocked outputs for testing the blueprint.
	MockedOutputs map[string]interface{} `json:"mocked_outputs,omitempty"`
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

// NewBlueprintAttributes instantiates a new BlueprintAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBlueprintAttributes(createdAt time.Time, definition AppDefinitionType, description string, name string, slug string, updatedAt time.Time) *BlueprintAttributes {
	this := BlueprintAttributes{}
	this.CreatedAt = createdAt
	this.Definition = definition
	this.Description = description
	this.Name = name
	this.Slug = slug
	this.UpdatedAt = updatedAt
	return &this
}

// NewBlueprintAttributesWithDefaults instantiates a new BlueprintAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBlueprintAttributesWithDefaults() *BlueprintAttributes {
	this := BlueprintAttributes{}
	var definition AppDefinitionType = APPDEFINITIONTYPE_APPDEFINITIONS
	this.Definition = definition
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *BlueprintAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *BlueprintAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *BlueprintAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDefinition returns the Definition field value.
func (o *BlueprintAttributes) GetDefinition() AppDefinitionType {
	if o == nil {
		var ret AppDefinitionType
		return ret
	}
	return o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value
// and a boolean to check if the value has been set.
func (o *BlueprintAttributes) GetDefinitionOk() (*AppDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Definition, true
}

// SetDefinition sets field value.
func (o *BlueprintAttributes) SetDefinition(v AppDefinitionType) {
	o.Definition = v
}

// GetDescription returns the Description field value.
func (o *BlueprintAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *BlueprintAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *BlueprintAttributes) SetDescription(v string) {
	o.Description = v
}

// GetEmbeddedDatastoreBlueprints returns the EmbeddedDatastoreBlueprints field value if set, zero value otherwise.
func (o *BlueprintAttributes) GetEmbeddedDatastoreBlueprints() map[string]interface{} {
	if o == nil || o.EmbeddedDatastoreBlueprints == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.EmbeddedDatastoreBlueprints
}

// GetEmbeddedDatastoreBlueprintsOk returns a tuple with the EmbeddedDatastoreBlueprints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintAttributes) GetEmbeddedDatastoreBlueprintsOk() (*map[string]interface{}, bool) {
	if o == nil || o.EmbeddedDatastoreBlueprints == nil {
		return nil, false
	}
	return &o.EmbeddedDatastoreBlueprints, true
}

// HasEmbeddedDatastoreBlueprints returns a boolean if a field has been set.
func (o *BlueprintAttributes) HasEmbeddedDatastoreBlueprints() bool {
	return o != nil && o.EmbeddedDatastoreBlueprints != nil
}

// SetEmbeddedDatastoreBlueprints gets a reference to the given map[string]interface{} and assigns it to the EmbeddedDatastoreBlueprints field.
func (o *BlueprintAttributes) SetEmbeddedDatastoreBlueprints(v map[string]interface{}) {
	o.EmbeddedDatastoreBlueprints = v
}

// GetEmbeddedNativeActions returns the EmbeddedNativeActions field value if set, zero value otherwise.
func (o *BlueprintAttributes) GetEmbeddedNativeActions() []map[string]interface{} {
	if o == nil || o.EmbeddedNativeActions == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.EmbeddedNativeActions
}

// GetEmbeddedNativeActionsOk returns a tuple with the EmbeddedNativeActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintAttributes) GetEmbeddedNativeActionsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.EmbeddedNativeActions == nil {
		return nil, false
	}
	return &o.EmbeddedNativeActions, true
}

// HasEmbeddedNativeActions returns a boolean if a field has been set.
func (o *BlueprintAttributes) HasEmbeddedNativeActions() bool {
	return o != nil && o.EmbeddedNativeActions != nil
}

// SetEmbeddedNativeActions gets a reference to the given []map[string]interface{} and assigns it to the EmbeddedNativeActions field.
func (o *BlueprintAttributes) SetEmbeddedNativeActions(v []map[string]interface{}) {
	o.EmbeddedNativeActions = v
}

// GetEmbeddedWorkflowBlueprints returns the EmbeddedWorkflowBlueprints field value if set, zero value otherwise.
func (o *BlueprintAttributes) GetEmbeddedWorkflowBlueprints() map[string]interface{} {
	if o == nil || o.EmbeddedWorkflowBlueprints == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.EmbeddedWorkflowBlueprints
}

// GetEmbeddedWorkflowBlueprintsOk returns a tuple with the EmbeddedWorkflowBlueprints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintAttributes) GetEmbeddedWorkflowBlueprintsOk() (*map[string]interface{}, bool) {
	if o == nil || o.EmbeddedWorkflowBlueprints == nil {
		return nil, false
	}
	return &o.EmbeddedWorkflowBlueprints, true
}

// HasEmbeddedWorkflowBlueprints returns a boolean if a field has been set.
func (o *BlueprintAttributes) HasEmbeddedWorkflowBlueprints() bool {
	return o != nil && o.EmbeddedWorkflowBlueprints != nil
}

// SetEmbeddedWorkflowBlueprints gets a reference to the given map[string]interface{} and assigns it to the EmbeddedWorkflowBlueprints field.
func (o *BlueprintAttributes) SetEmbeddedWorkflowBlueprints(v map[string]interface{}) {
	o.EmbeddedWorkflowBlueprints = v
}

// GetIntegrationId returns the IntegrationId field value if set, zero value otherwise.
func (o *BlueprintAttributes) GetIntegrationId() string {
	if o == nil || o.IntegrationId == nil {
		var ret string
		return ret
	}
	return *o.IntegrationId
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintAttributes) GetIntegrationIdOk() (*string, bool) {
	if o == nil || o.IntegrationId == nil {
		return nil, false
	}
	return o.IntegrationId, true
}

// HasIntegrationId returns a boolean if a field has been set.
func (o *BlueprintAttributes) HasIntegrationId() bool {
	return o != nil && o.IntegrationId != nil
}

// SetIntegrationId gets a reference to the given string and assigns it to the IntegrationId field.
func (o *BlueprintAttributes) SetIntegrationId(v string) {
	o.IntegrationId = &v
}

// GetMockedOutputs returns the MockedOutputs field value if set, zero value otherwise.
func (o *BlueprintAttributes) GetMockedOutputs() map[string]interface{} {
	if o == nil || o.MockedOutputs == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.MockedOutputs
}

// GetMockedOutputsOk returns a tuple with the MockedOutputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintAttributes) GetMockedOutputsOk() (*map[string]interface{}, bool) {
	if o == nil || o.MockedOutputs == nil {
		return nil, false
	}
	return &o.MockedOutputs, true
}

// HasMockedOutputs returns a boolean if a field has been set.
func (o *BlueprintAttributes) HasMockedOutputs() bool {
	return o != nil && o.MockedOutputs != nil
}

// SetMockedOutputs gets a reference to the given map[string]interface{} and assigns it to the MockedOutputs field.
func (o *BlueprintAttributes) SetMockedOutputs(v map[string]interface{}) {
	o.MockedOutputs = v
}

// GetName returns the Name field value.
func (o *BlueprintAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BlueprintAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *BlueprintAttributes) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value.
func (o *BlueprintAttributes) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *BlueprintAttributes) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value.
func (o *BlueprintAttributes) SetSlug(v string) {
	o.Slug = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *BlueprintAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *BlueprintAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *BlueprintAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetTileBackground returns the TileBackground field value if set, zero value otherwise.
func (o *BlueprintAttributes) GetTileBackground() string {
	if o == nil || o.TileBackground == nil {
		var ret string
		return ret
	}
	return *o.TileBackground
}

// GetTileBackgroundOk returns a tuple with the TileBackground field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintAttributes) GetTileBackgroundOk() (*string, bool) {
	if o == nil || o.TileBackground == nil {
		return nil, false
	}
	return o.TileBackground, true
}

// HasTileBackground returns a boolean if a field has been set.
func (o *BlueprintAttributes) HasTileBackground() bool {
	return o != nil && o.TileBackground != nil
}

// SetTileBackground gets a reference to the given string and assigns it to the TileBackground field.
func (o *BlueprintAttributes) SetTileBackground(v string) {
	o.TileBackground = &v
}

// GetTileIconActionFqn returns the TileIconActionFqn field value if set, zero value otherwise.
func (o *BlueprintAttributes) GetTileIconActionFqn() string {
	if o == nil || o.TileIconActionFqn == nil {
		var ret string
		return ret
	}
	return *o.TileIconActionFqn
}

// GetTileIconActionFqnOk returns a tuple with the TileIconActionFqn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintAttributes) GetTileIconActionFqnOk() (*string, bool) {
	if o == nil || o.TileIconActionFqn == nil {
		return nil, false
	}
	return o.TileIconActionFqn, true
}

// HasTileIconActionFqn returns a boolean if a field has been set.
func (o *BlueprintAttributes) HasTileIconActionFqn() bool {
	return o != nil && o.TileIconActionFqn != nil
}

// SetTileIconActionFqn gets a reference to the given string and assigns it to the TileIconActionFqn field.
func (o *BlueprintAttributes) SetTileIconActionFqn(v string) {
	o.TileIconActionFqn = &v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *BlueprintAttributes) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *BlueprintAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *BlueprintAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o BlueprintAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["definition"] = o.Definition
	toSerialize["description"] = o.Description
	if o.EmbeddedDatastoreBlueprints != nil {
		toSerialize["embedded_datastore_blueprints"] = o.EmbeddedDatastoreBlueprints
	}
	if o.EmbeddedNativeActions != nil {
		toSerialize["embedded_native_actions"] = o.EmbeddedNativeActions
	}
	if o.EmbeddedWorkflowBlueprints != nil {
		toSerialize["embedded_workflow_blueprints"] = o.EmbeddedWorkflowBlueprints
	}
	if o.IntegrationId != nil {
		toSerialize["integration_id"] = o.IntegrationId
	}
	if o.MockedOutputs != nil {
		toSerialize["mocked_outputs"] = o.MockedOutputs
	}
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
func (o *BlueprintAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt                   *time.Time               `json:"created_at"`
		Definition                  *AppDefinitionType       `json:"definition"`
		Description                 *string                  `json:"description"`
		EmbeddedDatastoreBlueprints map[string]interface{}   `json:"embedded_datastore_blueprints,omitempty"`
		EmbeddedNativeActions       []map[string]interface{} `json:"embedded_native_actions,omitempty"`
		EmbeddedWorkflowBlueprints  map[string]interface{}   `json:"embedded_workflow_blueprints,omitempty"`
		IntegrationId               *string                  `json:"integration_id,omitempty"`
		MockedOutputs               map[string]interface{}   `json:"mocked_outputs,omitempty"`
		Name                        *string                  `json:"name"`
		Slug                        *string                  `json:"slug"`
		Tags                        []string                 `json:"tags,omitempty"`
		TileBackground              *string                  `json:"tile_background,omitempty"`
		TileIconActionFqn           *string                  `json:"tile_icon_action_fqn,omitempty"`
		UpdatedAt                   *time.Time               `json:"updated_at"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "definition", "description", "embedded_datastore_blueprints", "embedded_native_actions", "embedded_workflow_blueprints", "integration_id", "mocked_outputs", "name", "slug", "tags", "tile_background", "tile_icon_action_fqn", "updated_at"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = *all.CreatedAt
	if !all.Definition.IsValid() {
		hasInvalidField = true
	} else {
		o.Definition = *all.Definition
	}
	o.Description = *all.Description
	o.EmbeddedDatastoreBlueprints = all.EmbeddedDatastoreBlueprints
	o.EmbeddedNativeActions = all.EmbeddedNativeActions
	o.EmbeddedWorkflowBlueprints = all.EmbeddedWorkflowBlueprints
	o.IntegrationId = all.IntegrationId
	o.MockedOutputs = all.MockedOutputs
	o.Name = *all.Name
	o.Slug = *all.Slug
	o.Tags = all.Tags
	o.TileBackground = all.TileBackground
	o.TileIconActionFqn = all.TileIconActionFqn
	o.UpdatedAt = *all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
