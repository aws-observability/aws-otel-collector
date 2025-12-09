// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsArchiveDestinationS3 The S3 archive destination.
type LogsArchiveDestinationS3 struct {
	// The bucket where the archive will be stored.
	Bucket string `json:"bucket"`
	// The S3 encryption settings.
	Encryption *LogsArchiveEncryptionS3 `json:"encryption,omitempty"`
	// The S3 Archive's integration destination.
	Integration LogsArchiveIntegrationS3 `json:"integration"`
	// The archive path.
	Path *string `json:"path,omitempty"`
	// The storage class where the archive will be stored.
	StorageClass *LogsArchiveStorageClassS3Type `json:"storage_class,omitempty"`
	// Type of the S3 archive destination.
	Type LogsArchiveDestinationS3Type `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLogsArchiveDestinationS3 instantiates a new LogsArchiveDestinationS3 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsArchiveDestinationS3(bucket string, integration LogsArchiveIntegrationS3, typeVar LogsArchiveDestinationS3Type) *LogsArchiveDestinationS3 {
	this := LogsArchiveDestinationS3{}
	this.Bucket = bucket
	this.Integration = integration
	var storageClass LogsArchiveStorageClassS3Type = LOGSARCHIVESTORAGECLASSS3TYPE_STANDARD
	this.StorageClass = &storageClass
	this.Type = typeVar
	return &this
}

// NewLogsArchiveDestinationS3WithDefaults instantiates a new LogsArchiveDestinationS3 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsArchiveDestinationS3WithDefaults() *LogsArchiveDestinationS3 {
	this := LogsArchiveDestinationS3{}
	var storageClass LogsArchiveStorageClassS3Type = LOGSARCHIVESTORAGECLASSS3TYPE_STANDARD
	this.StorageClass = &storageClass
	var typeVar LogsArchiveDestinationS3Type = LOGSARCHIVEDESTINATIONS3TYPE_S3
	this.Type = typeVar
	return &this
}

// GetBucket returns the Bucket field value.
func (o *LogsArchiveDestinationS3) GetBucket() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value
// and a boolean to check if the value has been set.
func (o *LogsArchiveDestinationS3) GetBucketOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bucket, true
}

// SetBucket sets field value.
func (o *LogsArchiveDestinationS3) SetBucket(v string) {
	o.Bucket = v
}

// GetEncryption returns the Encryption field value if set, zero value otherwise.
func (o *LogsArchiveDestinationS3) GetEncryption() LogsArchiveEncryptionS3 {
	if o == nil || o.Encryption == nil {
		var ret LogsArchiveEncryptionS3
		return ret
	}
	return *o.Encryption
}

// GetEncryptionOk returns a tuple with the Encryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsArchiveDestinationS3) GetEncryptionOk() (*LogsArchiveEncryptionS3, bool) {
	if o == nil || o.Encryption == nil {
		return nil, false
	}
	return o.Encryption, true
}

// HasEncryption returns a boolean if a field has been set.
func (o *LogsArchiveDestinationS3) HasEncryption() bool {
	return o != nil && o.Encryption != nil
}

// SetEncryption gets a reference to the given LogsArchiveEncryptionS3 and assigns it to the Encryption field.
func (o *LogsArchiveDestinationS3) SetEncryption(v LogsArchiveEncryptionS3) {
	o.Encryption = &v
}

// GetIntegration returns the Integration field value.
func (o *LogsArchiveDestinationS3) GetIntegration() LogsArchiveIntegrationS3 {
	if o == nil {
		var ret LogsArchiveIntegrationS3
		return ret
	}
	return o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value
// and a boolean to check if the value has been set.
func (o *LogsArchiveDestinationS3) GetIntegrationOk() (*LogsArchiveIntegrationS3, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Integration, true
}

// SetIntegration sets field value.
func (o *LogsArchiveDestinationS3) SetIntegration(v LogsArchiveIntegrationS3) {
	o.Integration = v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *LogsArchiveDestinationS3) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsArchiveDestinationS3) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *LogsArchiveDestinationS3) HasPath() bool {
	return o != nil && o.Path != nil
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *LogsArchiveDestinationS3) SetPath(v string) {
	o.Path = &v
}

// GetStorageClass returns the StorageClass field value if set, zero value otherwise.
func (o *LogsArchiveDestinationS3) GetStorageClass() LogsArchiveStorageClassS3Type {
	if o == nil || o.StorageClass == nil {
		var ret LogsArchiveStorageClassS3Type
		return ret
	}
	return *o.StorageClass
}

// GetStorageClassOk returns a tuple with the StorageClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsArchiveDestinationS3) GetStorageClassOk() (*LogsArchiveStorageClassS3Type, bool) {
	if o == nil || o.StorageClass == nil {
		return nil, false
	}
	return o.StorageClass, true
}

// HasStorageClass returns a boolean if a field has been set.
func (o *LogsArchiveDestinationS3) HasStorageClass() bool {
	return o != nil && o.StorageClass != nil
}

// SetStorageClass gets a reference to the given LogsArchiveStorageClassS3Type and assigns it to the StorageClass field.
func (o *LogsArchiveDestinationS3) SetStorageClass(v LogsArchiveStorageClassS3Type) {
	o.StorageClass = &v
}

// GetType returns the Type field value.
func (o *LogsArchiveDestinationS3) GetType() LogsArchiveDestinationS3Type {
	if o == nil {
		var ret LogsArchiveDestinationS3Type
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LogsArchiveDestinationS3) GetTypeOk() (*LogsArchiveDestinationS3Type, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *LogsArchiveDestinationS3) SetType(v LogsArchiveDestinationS3Type) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsArchiveDestinationS3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["bucket"] = o.Bucket
	if o.Encryption != nil {
		toSerialize["encryption"] = o.Encryption
	}
	toSerialize["integration"] = o.Integration
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.StorageClass != nil {
		toSerialize["storage_class"] = o.StorageClass
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogsArchiveDestinationS3) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Bucket       *string                        `json:"bucket"`
		Encryption   *LogsArchiveEncryptionS3       `json:"encryption,omitempty"`
		Integration  *LogsArchiveIntegrationS3      `json:"integration"`
		Path         *string                        `json:"path,omitempty"`
		StorageClass *LogsArchiveStorageClassS3Type `json:"storage_class,omitempty"`
		Type         *LogsArchiveDestinationS3Type  `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Bucket == nil {
		return fmt.Errorf("required field bucket missing")
	}
	if all.Integration == nil {
		return fmt.Errorf("required field integration missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"bucket", "encryption", "integration", "path", "storage_class", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Bucket = *all.Bucket
	if all.Encryption != nil && all.Encryption.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Encryption = all.Encryption
	if all.Integration.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Integration = *all.Integration
	o.Path = all.Path
	if all.StorageClass != nil && !all.StorageClass.IsValid() {
		hasInvalidField = true
	} else {
		o.StorageClass = all.StorageClass
	}
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
