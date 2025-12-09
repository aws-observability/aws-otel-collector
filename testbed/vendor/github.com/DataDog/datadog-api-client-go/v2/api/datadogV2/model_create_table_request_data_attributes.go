// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateTableRequestDataAttributes Attributes that define the reference table's configuration and properties.
type CreateTableRequestDataAttributes struct {
	// Optional text describing the purpose or contents of this reference table.
	Description *string `json:"description,omitempty"`
	// Metadata specifying where and how to access the reference table's data file.
	FileMetadata *CreateTableRequestDataAttributesFileMetadata `json:"file_metadata,omitempty"`
	// Schema defining the structure and columns of the reference table.
	Schema CreateTableRequestDataAttributesSchema `json:"schema"`
	// The source type for creating reference table data. Only these source types can be created through this API.
	Source ReferenceTableCreateSourceType `json:"source"`
	// Name to identify this reference table.
	TableName string `json:"table_name"`
	// Tags for organizing and filtering reference tables.
	Tags []string `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateTableRequestDataAttributes instantiates a new CreateTableRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateTableRequestDataAttributes(schema CreateTableRequestDataAttributesSchema, source ReferenceTableCreateSourceType, tableName string) *CreateTableRequestDataAttributes {
	this := CreateTableRequestDataAttributes{}
	this.Schema = schema
	this.Source = source
	this.TableName = tableName
	return &this
}

// NewCreateTableRequestDataAttributesWithDefaults instantiates a new CreateTableRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateTableRequestDataAttributesWithDefaults() *CreateTableRequestDataAttributes {
	this := CreateTableRequestDataAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateTableRequestDataAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTableRequestDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateTableRequestDataAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateTableRequestDataAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetFileMetadata returns the FileMetadata field value if set, zero value otherwise.
func (o *CreateTableRequestDataAttributes) GetFileMetadata() CreateTableRequestDataAttributesFileMetadata {
	if o == nil || o.FileMetadata == nil {
		var ret CreateTableRequestDataAttributesFileMetadata
		return ret
	}
	return *o.FileMetadata
}

// GetFileMetadataOk returns a tuple with the FileMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTableRequestDataAttributes) GetFileMetadataOk() (*CreateTableRequestDataAttributesFileMetadata, bool) {
	if o == nil || o.FileMetadata == nil {
		return nil, false
	}
	return o.FileMetadata, true
}

// HasFileMetadata returns a boolean if a field has been set.
func (o *CreateTableRequestDataAttributes) HasFileMetadata() bool {
	return o != nil && o.FileMetadata != nil
}

// SetFileMetadata gets a reference to the given CreateTableRequestDataAttributesFileMetadata and assigns it to the FileMetadata field.
func (o *CreateTableRequestDataAttributes) SetFileMetadata(v CreateTableRequestDataAttributesFileMetadata) {
	o.FileMetadata = &v
}

// GetSchema returns the Schema field value.
func (o *CreateTableRequestDataAttributes) GetSchema() CreateTableRequestDataAttributesSchema {
	if o == nil {
		var ret CreateTableRequestDataAttributesSchema
		return ret
	}
	return o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value
// and a boolean to check if the value has been set.
func (o *CreateTableRequestDataAttributes) GetSchemaOk() (*CreateTableRequestDataAttributesSchema, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schema, true
}

// SetSchema sets field value.
func (o *CreateTableRequestDataAttributes) SetSchema(v CreateTableRequestDataAttributesSchema) {
	o.Schema = v
}

// GetSource returns the Source field value.
func (o *CreateTableRequestDataAttributes) GetSource() ReferenceTableCreateSourceType {
	if o == nil {
		var ret ReferenceTableCreateSourceType
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *CreateTableRequestDataAttributes) GetSourceOk() (*ReferenceTableCreateSourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *CreateTableRequestDataAttributes) SetSource(v ReferenceTableCreateSourceType) {
	o.Source = v
}

// GetTableName returns the TableName field value.
func (o *CreateTableRequestDataAttributes) GetTableName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TableName
}

// GetTableNameOk returns a tuple with the TableName field value
// and a boolean to check if the value has been set.
func (o *CreateTableRequestDataAttributes) GetTableNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TableName, true
}

// SetTableName sets field value.
func (o *CreateTableRequestDataAttributes) SetTableName(v string) {
	o.TableName = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateTableRequestDataAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTableRequestDataAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateTableRequestDataAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CreateTableRequestDataAttributes) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateTableRequestDataAttributes) MarshalJSON() ([]byte, error) {
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
	toSerialize["schema"] = o.Schema
	toSerialize["source"] = o.Source
	toSerialize["table_name"] = o.TableName
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateTableRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description  *string                                       `json:"description,omitempty"`
		FileMetadata *CreateTableRequestDataAttributesFileMetadata `json:"file_metadata,omitempty"`
		Schema       *CreateTableRequestDataAttributesSchema       `json:"schema"`
		Source       *ReferenceTableCreateSourceType               `json:"source"`
		TableName    *string                                       `json:"table_name"`
		Tags         []string                                      `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Schema == nil {
		return fmt.Errorf("required field schema missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	if all.TableName == nil {
		return fmt.Errorf("required field table_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "file_metadata", "schema", "source", "table_name", "tags"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = all.Description
	o.FileMetadata = all.FileMetadata
	if all.Schema.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Schema = *all.Schema
	if !all.Source.IsValid() {
		hasInvalidField = true
	} else {
		o.Source = *all.Source
	}
	o.TableName = *all.TableName
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
