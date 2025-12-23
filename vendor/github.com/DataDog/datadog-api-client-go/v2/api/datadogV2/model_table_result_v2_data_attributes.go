// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableResultV2DataAttributes Attributes that define the reference table's configuration and properties.
type TableResultV2DataAttributes struct {
	// UUID of the user who created the reference table.
	CreatedBy *string `json:"created_by,omitempty"`
	// Optional text describing the purpose or contents of this reference table.
	Description *string `json:"description,omitempty"`
	// Metadata specifying where and how to access the reference table's data file.
	//
	// For cloud storage tables (S3/GCS/Azure):
	//   - sync_enabled and access_details will always be present
	//   - error fields (error_message, error_row_count, error_type) are present only when errors occur
	//
	// For local file tables:
	//   - error fields (error_message, error_row_count) are present only when errors occur
	//   - sync_enabled, access_details are never present
	FileMetadata *TableResultV2DataAttributesFileMetadata `json:"file_metadata,omitempty"`
	// UUID of the user who last updated the reference table.
	LastUpdatedBy *string `json:"last_updated_by,omitempty"`
	// The number of successfully processed rows in the reference table.
	RowCount *int64 `json:"row_count,omitempty"`
	// Schema defining the structure and columns of the reference table.
	Schema *TableResultV2DataAttributesSchema `json:"schema,omitempty"`
	// The source type for reference table data. Includes all possible source types that can appear in responses.
	Source *ReferenceTableSourceType `json:"source,omitempty"`
	// The processing status of the table.
	Status *string `json:"status,omitempty"`
	// Unique name to identify this reference table. Used in enrichment processors and API calls.
	TableName *string `json:"table_name,omitempty"`
	// Tags for organizing and filtering reference tables.
	Tags []string `json:"tags,omitempty"`
	// When the reference table was last updated, in ISO 8601 format.
	UpdatedAt *string `json:"updated_at,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTableResultV2DataAttributes instantiates a new TableResultV2DataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTableResultV2DataAttributes() *TableResultV2DataAttributes {
	this := TableResultV2DataAttributes{}
	return &this
}

// NewTableResultV2DataAttributesWithDefaults instantiates a new TableResultV2DataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTableResultV2DataAttributesWithDefaults() *TableResultV2DataAttributes {
	this := TableResultV2DataAttributes{}
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *TableResultV2DataAttributes) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableResultV2DataAttributes) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *TableResultV2DataAttributes) HasCreatedBy() bool {
	return o != nil && o.CreatedBy != nil
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *TableResultV2DataAttributes) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TableResultV2DataAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableResultV2DataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TableResultV2DataAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TableResultV2DataAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetFileMetadata returns the FileMetadata field value if set, zero value otherwise.
func (o *TableResultV2DataAttributes) GetFileMetadata() TableResultV2DataAttributesFileMetadata {
	if o == nil || o.FileMetadata == nil {
		var ret TableResultV2DataAttributesFileMetadata
		return ret
	}
	return *o.FileMetadata
}

// GetFileMetadataOk returns a tuple with the FileMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableResultV2DataAttributes) GetFileMetadataOk() (*TableResultV2DataAttributesFileMetadata, bool) {
	if o == nil || o.FileMetadata == nil {
		return nil, false
	}
	return o.FileMetadata, true
}

// HasFileMetadata returns a boolean if a field has been set.
func (o *TableResultV2DataAttributes) HasFileMetadata() bool {
	return o != nil && o.FileMetadata != nil
}

// SetFileMetadata gets a reference to the given TableResultV2DataAttributesFileMetadata and assigns it to the FileMetadata field.
func (o *TableResultV2DataAttributes) SetFileMetadata(v TableResultV2DataAttributesFileMetadata) {
	o.FileMetadata = &v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value if set, zero value otherwise.
func (o *TableResultV2DataAttributes) GetLastUpdatedBy() string {
	if o == nil || o.LastUpdatedBy == nil {
		var ret string
		return ret
	}
	return *o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableResultV2DataAttributes) GetLastUpdatedByOk() (*string, bool) {
	if o == nil || o.LastUpdatedBy == nil {
		return nil, false
	}
	return o.LastUpdatedBy, true
}

// HasLastUpdatedBy returns a boolean if a field has been set.
func (o *TableResultV2DataAttributes) HasLastUpdatedBy() bool {
	return o != nil && o.LastUpdatedBy != nil
}

// SetLastUpdatedBy gets a reference to the given string and assigns it to the LastUpdatedBy field.
func (o *TableResultV2DataAttributes) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = &v
}

// GetRowCount returns the RowCount field value if set, zero value otherwise.
func (o *TableResultV2DataAttributes) GetRowCount() int64 {
	if o == nil || o.RowCount == nil {
		var ret int64
		return ret
	}
	return *o.RowCount
}

// GetRowCountOk returns a tuple with the RowCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableResultV2DataAttributes) GetRowCountOk() (*int64, bool) {
	if o == nil || o.RowCount == nil {
		return nil, false
	}
	return o.RowCount, true
}

// HasRowCount returns a boolean if a field has been set.
func (o *TableResultV2DataAttributes) HasRowCount() bool {
	return o != nil && o.RowCount != nil
}

// SetRowCount gets a reference to the given int64 and assigns it to the RowCount field.
func (o *TableResultV2DataAttributes) SetRowCount(v int64) {
	o.RowCount = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *TableResultV2DataAttributes) GetSchema() TableResultV2DataAttributesSchema {
	if o == nil || o.Schema == nil {
		var ret TableResultV2DataAttributesSchema
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableResultV2DataAttributes) GetSchemaOk() (*TableResultV2DataAttributesSchema, bool) {
	if o == nil || o.Schema == nil {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *TableResultV2DataAttributes) HasSchema() bool {
	return o != nil && o.Schema != nil
}

// SetSchema gets a reference to the given TableResultV2DataAttributesSchema and assigns it to the Schema field.
func (o *TableResultV2DataAttributes) SetSchema(v TableResultV2DataAttributesSchema) {
	o.Schema = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *TableResultV2DataAttributes) GetSource() ReferenceTableSourceType {
	if o == nil || o.Source == nil {
		var ret ReferenceTableSourceType
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableResultV2DataAttributes) GetSourceOk() (*ReferenceTableSourceType, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *TableResultV2DataAttributes) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given ReferenceTableSourceType and assigns it to the Source field.
func (o *TableResultV2DataAttributes) SetSource(v ReferenceTableSourceType) {
	o.Source = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TableResultV2DataAttributes) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableResultV2DataAttributes) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TableResultV2DataAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TableResultV2DataAttributes) SetStatus(v string) {
	o.Status = &v
}

// GetTableName returns the TableName field value if set, zero value otherwise.
func (o *TableResultV2DataAttributes) GetTableName() string {
	if o == nil || o.TableName == nil {
		var ret string
		return ret
	}
	return *o.TableName
}

// GetTableNameOk returns a tuple with the TableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableResultV2DataAttributes) GetTableNameOk() (*string, bool) {
	if o == nil || o.TableName == nil {
		return nil, false
	}
	return o.TableName, true
}

// HasTableName returns a boolean if a field has been set.
func (o *TableResultV2DataAttributes) HasTableName() bool {
	return o != nil && o.TableName != nil
}

// SetTableName gets a reference to the given string and assigns it to the TableName field.
func (o *TableResultV2DataAttributes) SetTableName(v string) {
	o.TableName = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *TableResultV2DataAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableResultV2DataAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *TableResultV2DataAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *TableResultV2DataAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *TableResultV2DataAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableResultV2DataAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *TableResultV2DataAttributes) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *TableResultV2DataAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TableResultV2DataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.FileMetadata != nil {
		toSerialize["file_metadata"] = o.FileMetadata
	}
	if o.LastUpdatedBy != nil {
		toSerialize["last_updated_by"] = o.LastUpdatedBy
	}
	if o.RowCount != nil {
		toSerialize["row_count"] = o.RowCount
	}
	if o.Schema != nil {
		toSerialize["schema"] = o.Schema
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.TableName != nil {
		toSerialize["table_name"] = o.TableName
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TableResultV2DataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedBy     *string                                  `json:"created_by,omitempty"`
		Description   *string                                  `json:"description,omitempty"`
		FileMetadata  *TableResultV2DataAttributesFileMetadata `json:"file_metadata,omitempty"`
		LastUpdatedBy *string                                  `json:"last_updated_by,omitempty"`
		RowCount      *int64                                   `json:"row_count,omitempty"`
		Schema        *TableResultV2DataAttributesSchema       `json:"schema,omitempty"`
		Source        *ReferenceTableSourceType                `json:"source,omitempty"`
		Status        *string                                  `json:"status,omitempty"`
		TableName     *string                                  `json:"table_name,omitempty"`
		Tags          []string                                 `json:"tags,omitempty"`
		UpdatedAt     *string                                  `json:"updated_at,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_by", "description", "file_metadata", "last_updated_by", "row_count", "schema", "source", "status", "table_name", "tags", "updated_at"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedBy = all.CreatedBy
	o.Description = all.Description
	if all.FileMetadata != nil && all.FileMetadata.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.FileMetadata = all.FileMetadata
	o.LastUpdatedBy = all.LastUpdatedBy
	o.RowCount = all.RowCount
	if all.Schema != nil && all.Schema.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Schema = all.Schema
	if all.Source != nil && !all.Source.IsValid() {
		hasInvalidField = true
	} else {
		o.Source = all.Source
	}
	o.Status = all.Status
	o.TableName = all.TableName
	o.Tags = all.Tags
	o.UpdatedAt = all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
