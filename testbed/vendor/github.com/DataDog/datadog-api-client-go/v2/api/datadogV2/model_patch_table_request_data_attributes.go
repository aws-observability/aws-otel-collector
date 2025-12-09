// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PatchTableRequestDataAttributes Attributes that define the updates to the reference table's configuration and properties.
type PatchTableRequestDataAttributes struct {
	// Optional text describing the purpose or contents of this reference table.
	Description *string `json:"description,omitempty"`
	// Metadata specifying where and how to access the reference table's data file.
	FileMetadata *PatchTableRequestDataAttributesFileMetadata `json:"file_metadata,omitempty"`
	// Schema defining the updates to the structure and columns of the reference table. Schema fields cannot be deleted or renamed.
	Schema *PatchTableRequestDataAttributesSchema `json:"schema,omitempty"`
	// Whether this table is synced automatically.
	SyncEnabled *bool `json:"sync_enabled,omitempty"`
	// Tags for organizing and filtering reference tables.
	Tags []string `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPatchTableRequestDataAttributes instantiates a new PatchTableRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPatchTableRequestDataAttributes() *PatchTableRequestDataAttributes {
	this := PatchTableRequestDataAttributes{}
	return &this
}

// NewPatchTableRequestDataAttributesWithDefaults instantiates a new PatchTableRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPatchTableRequestDataAttributesWithDefaults() *PatchTableRequestDataAttributes {
	this := PatchTableRequestDataAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchTableRequestDataAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchTableRequestDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchTableRequestDataAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchTableRequestDataAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetFileMetadata returns the FileMetadata field value if set, zero value otherwise.
func (o *PatchTableRequestDataAttributes) GetFileMetadata() PatchTableRequestDataAttributesFileMetadata {
	if o == nil || o.FileMetadata == nil {
		var ret PatchTableRequestDataAttributesFileMetadata
		return ret
	}
	return *o.FileMetadata
}

// GetFileMetadataOk returns a tuple with the FileMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchTableRequestDataAttributes) GetFileMetadataOk() (*PatchTableRequestDataAttributesFileMetadata, bool) {
	if o == nil || o.FileMetadata == nil {
		return nil, false
	}
	return o.FileMetadata, true
}

// HasFileMetadata returns a boolean if a field has been set.
func (o *PatchTableRequestDataAttributes) HasFileMetadata() bool {
	return o != nil && o.FileMetadata != nil
}

// SetFileMetadata gets a reference to the given PatchTableRequestDataAttributesFileMetadata and assigns it to the FileMetadata field.
func (o *PatchTableRequestDataAttributes) SetFileMetadata(v PatchTableRequestDataAttributesFileMetadata) {
	o.FileMetadata = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *PatchTableRequestDataAttributes) GetSchema() PatchTableRequestDataAttributesSchema {
	if o == nil || o.Schema == nil {
		var ret PatchTableRequestDataAttributesSchema
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchTableRequestDataAttributes) GetSchemaOk() (*PatchTableRequestDataAttributesSchema, bool) {
	if o == nil || o.Schema == nil {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *PatchTableRequestDataAttributes) HasSchema() bool {
	return o != nil && o.Schema != nil
}

// SetSchema gets a reference to the given PatchTableRequestDataAttributesSchema and assigns it to the Schema field.
func (o *PatchTableRequestDataAttributes) SetSchema(v PatchTableRequestDataAttributesSchema) {
	o.Schema = &v
}

// GetSyncEnabled returns the SyncEnabled field value if set, zero value otherwise.
func (o *PatchTableRequestDataAttributes) GetSyncEnabled() bool {
	if o == nil || o.SyncEnabled == nil {
		var ret bool
		return ret
	}
	return *o.SyncEnabled
}

// GetSyncEnabledOk returns a tuple with the SyncEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchTableRequestDataAttributes) GetSyncEnabledOk() (*bool, bool) {
	if o == nil || o.SyncEnabled == nil {
		return nil, false
	}
	return o.SyncEnabled, true
}

// HasSyncEnabled returns a boolean if a field has been set.
func (o *PatchTableRequestDataAttributes) HasSyncEnabled() bool {
	return o != nil && o.SyncEnabled != nil
}

// SetSyncEnabled gets a reference to the given bool and assigns it to the SyncEnabled field.
func (o *PatchTableRequestDataAttributes) SetSyncEnabled(v bool) {
	o.SyncEnabled = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchTableRequestDataAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchTableRequestDataAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchTableRequestDataAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *PatchTableRequestDataAttributes) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PatchTableRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.FileMetadata != nil {
		toSerialize["file_metadata"] = o.FileMetadata
	}
	if o.Schema != nil {
		toSerialize["schema"] = o.Schema
	}
	if o.SyncEnabled != nil {
		toSerialize["sync_enabled"] = o.SyncEnabled
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PatchTableRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description  *string                                      `json:"description,omitempty"`
		FileMetadata *PatchTableRequestDataAttributesFileMetadata `json:"file_metadata,omitempty"`
		Schema       *PatchTableRequestDataAttributesSchema       `json:"schema,omitempty"`
		SyncEnabled  *bool                                        `json:"sync_enabled,omitempty"`
		Tags         []string                                     `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "file_metadata", "schema", "sync_enabled", "tags"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = all.Description
	o.FileMetadata = all.FileMetadata
	if all.Schema != nil && all.Schema.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Schema = all.Schema
	o.SyncEnabled = all.SyncEnabled
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
