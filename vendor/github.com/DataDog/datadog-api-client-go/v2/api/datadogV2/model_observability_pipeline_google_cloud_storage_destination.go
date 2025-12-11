// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineGoogleCloudStorageDestination The `google_cloud_storage` destination stores logs in a Google Cloud Storage (GCS) bucket.
// It requires a bucket name, GCP authentication, and metadata fields.
type ObservabilityPipelineGoogleCloudStorageDestination struct {
	// Access control list setting for objects written to the bucket.
	Acl ObservabilityPipelineGoogleCloudStorageDestinationAcl `json:"acl"`
	// GCP credentials used to authenticate with Google Cloud Storage.
	Auth ObservabilityPipelineGcpAuth `json:"auth"`
	// Name of the GCS bucket.
	Bucket string `json:"bucket"`
	// Unique identifier for the destination component.
	Id string `json:"id"`
	// A list of component IDs whose output is used as the `input` for this component.
	Inputs []string `json:"inputs"`
	// Optional prefix for object keys within the GCS bucket.
	KeyPrefix *string `json:"key_prefix,omitempty"`
	// Custom metadata to attach to each object uploaded to the GCS bucket.
	Metadata []ObservabilityPipelineMetadataEntry `json:"metadata,omitempty"`
	// Storage class used for objects stored in GCS.
	StorageClass ObservabilityPipelineGoogleCloudStorageDestinationStorageClass `json:"storage_class"`
	// The destination type. Always `google_cloud_storage`.
	Type ObservabilityPipelineGoogleCloudStorageDestinationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineGoogleCloudStorageDestination instantiates a new ObservabilityPipelineGoogleCloudStorageDestination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineGoogleCloudStorageDestination(acl ObservabilityPipelineGoogleCloudStorageDestinationAcl, auth ObservabilityPipelineGcpAuth, bucket string, id string, inputs []string, storageClass ObservabilityPipelineGoogleCloudStorageDestinationStorageClass, typeVar ObservabilityPipelineGoogleCloudStorageDestinationType) *ObservabilityPipelineGoogleCloudStorageDestination {
	this := ObservabilityPipelineGoogleCloudStorageDestination{}
	this.Acl = acl
	this.Auth = auth
	this.Bucket = bucket
	this.Id = id
	this.Inputs = inputs
	this.StorageClass = storageClass
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineGoogleCloudStorageDestinationWithDefaults instantiates a new ObservabilityPipelineGoogleCloudStorageDestination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineGoogleCloudStorageDestinationWithDefaults() *ObservabilityPipelineGoogleCloudStorageDestination {
	this := ObservabilityPipelineGoogleCloudStorageDestination{}
	var typeVar ObservabilityPipelineGoogleCloudStorageDestinationType = OBSERVABILITYPIPELINEGOOGLECLOUDSTORAGEDESTINATIONTYPE_GOOGLE_CLOUD_STORAGE
	this.Type = typeVar
	return &this
}

// GetAcl returns the Acl field value.
func (o *ObservabilityPipelineGoogleCloudStorageDestination) GetAcl() ObservabilityPipelineGoogleCloudStorageDestinationAcl {
	if o == nil {
		var ret ObservabilityPipelineGoogleCloudStorageDestinationAcl
		return ret
	}
	return o.Acl
}

// GetAclOk returns a tuple with the Acl field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGoogleCloudStorageDestination) GetAclOk() (*ObservabilityPipelineGoogleCloudStorageDestinationAcl, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Acl, true
}

// SetAcl sets field value.
func (o *ObservabilityPipelineGoogleCloudStorageDestination) SetAcl(v ObservabilityPipelineGoogleCloudStorageDestinationAcl) {
	o.Acl = v
}

// GetAuth returns the Auth field value.
func (o *ObservabilityPipelineGoogleCloudStorageDestination) GetAuth() ObservabilityPipelineGcpAuth {
	if o == nil {
		var ret ObservabilityPipelineGcpAuth
		return ret
	}
	return o.Auth
}

// GetAuthOk returns a tuple with the Auth field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGoogleCloudStorageDestination) GetAuthOk() (*ObservabilityPipelineGcpAuth, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Auth, true
}

// SetAuth sets field value.
func (o *ObservabilityPipelineGoogleCloudStorageDestination) SetAuth(v ObservabilityPipelineGcpAuth) {
	o.Auth = v
}

// GetBucket returns the Bucket field value.
func (o *ObservabilityPipelineGoogleCloudStorageDestination) GetBucket() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGoogleCloudStorageDestination) GetBucketOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bucket, true
}

// SetBucket sets field value.
func (o *ObservabilityPipelineGoogleCloudStorageDestination) SetBucket(v string) {
	o.Bucket = v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineGoogleCloudStorageDestination) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGoogleCloudStorageDestination) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineGoogleCloudStorageDestination) SetId(v string) {
	o.Id = v
}

// GetInputs returns the Inputs field value.
func (o *ObservabilityPipelineGoogleCloudStorageDestination) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGoogleCloudStorageDestination) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *ObservabilityPipelineGoogleCloudStorageDestination) SetInputs(v []string) {
	o.Inputs = v
}

// GetKeyPrefix returns the KeyPrefix field value if set, zero value otherwise.
func (o *ObservabilityPipelineGoogleCloudStorageDestination) GetKeyPrefix() string {
	if o == nil || o.KeyPrefix == nil {
		var ret string
		return ret
	}
	return *o.KeyPrefix
}

// GetKeyPrefixOk returns a tuple with the KeyPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGoogleCloudStorageDestination) GetKeyPrefixOk() (*string, bool) {
	if o == nil || o.KeyPrefix == nil {
		return nil, false
	}
	return o.KeyPrefix, true
}

// HasKeyPrefix returns a boolean if a field has been set.
func (o *ObservabilityPipelineGoogleCloudStorageDestination) HasKeyPrefix() bool {
	return o != nil && o.KeyPrefix != nil
}

// SetKeyPrefix gets a reference to the given string and assigns it to the KeyPrefix field.
func (o *ObservabilityPipelineGoogleCloudStorageDestination) SetKeyPrefix(v string) {
	o.KeyPrefix = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ObservabilityPipelineGoogleCloudStorageDestination) GetMetadata() []ObservabilityPipelineMetadataEntry {
	if o == nil || o.Metadata == nil {
		var ret []ObservabilityPipelineMetadataEntry
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGoogleCloudStorageDestination) GetMetadataOk() (*[]ObservabilityPipelineMetadataEntry, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ObservabilityPipelineGoogleCloudStorageDestination) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given []ObservabilityPipelineMetadataEntry and assigns it to the Metadata field.
func (o *ObservabilityPipelineGoogleCloudStorageDestination) SetMetadata(v []ObservabilityPipelineMetadataEntry) {
	o.Metadata = v
}

// GetStorageClass returns the StorageClass field value.
func (o *ObservabilityPipelineGoogleCloudStorageDestination) GetStorageClass() ObservabilityPipelineGoogleCloudStorageDestinationStorageClass {
	if o == nil {
		var ret ObservabilityPipelineGoogleCloudStorageDestinationStorageClass
		return ret
	}
	return o.StorageClass
}

// GetStorageClassOk returns a tuple with the StorageClass field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGoogleCloudStorageDestination) GetStorageClassOk() (*ObservabilityPipelineGoogleCloudStorageDestinationStorageClass, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageClass, true
}

// SetStorageClass sets field value.
func (o *ObservabilityPipelineGoogleCloudStorageDestination) SetStorageClass(v ObservabilityPipelineGoogleCloudStorageDestinationStorageClass) {
	o.StorageClass = v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineGoogleCloudStorageDestination) GetType() ObservabilityPipelineGoogleCloudStorageDestinationType {
	if o == nil {
		var ret ObservabilityPipelineGoogleCloudStorageDestinationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGoogleCloudStorageDestination) GetTypeOk() (*ObservabilityPipelineGoogleCloudStorageDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineGoogleCloudStorageDestination) SetType(v ObservabilityPipelineGoogleCloudStorageDestinationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineGoogleCloudStorageDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["acl"] = o.Acl
	toSerialize["auth"] = o.Auth
	toSerialize["bucket"] = o.Bucket
	toSerialize["id"] = o.Id
	toSerialize["inputs"] = o.Inputs
	if o.KeyPrefix != nil {
		toSerialize["key_prefix"] = o.KeyPrefix
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["storage_class"] = o.StorageClass
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineGoogleCloudStorageDestination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Acl          *ObservabilityPipelineGoogleCloudStorageDestinationAcl          `json:"acl"`
		Auth         *ObservabilityPipelineGcpAuth                                   `json:"auth"`
		Bucket       *string                                                         `json:"bucket"`
		Id           *string                                                         `json:"id"`
		Inputs       *[]string                                                       `json:"inputs"`
		KeyPrefix    *string                                                         `json:"key_prefix,omitempty"`
		Metadata     []ObservabilityPipelineMetadataEntry                            `json:"metadata,omitempty"`
		StorageClass *ObservabilityPipelineGoogleCloudStorageDestinationStorageClass `json:"storage_class"`
		Type         *ObservabilityPipelineGoogleCloudStorageDestinationType         `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Acl == nil {
		return fmt.Errorf("required field acl missing")
	}
	if all.Auth == nil {
		return fmt.Errorf("required field auth missing")
	}
	if all.Bucket == nil {
		return fmt.Errorf("required field bucket missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Inputs == nil {
		return fmt.Errorf("required field inputs missing")
	}
	if all.StorageClass == nil {
		return fmt.Errorf("required field storage_class missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"acl", "auth", "bucket", "id", "inputs", "key_prefix", "metadata", "storage_class", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Acl.IsValid() {
		hasInvalidField = true
	} else {
		o.Acl = *all.Acl
	}
	if all.Auth.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Auth = *all.Auth
	o.Bucket = *all.Bucket
	o.Id = *all.Id
	o.Inputs = *all.Inputs
	o.KeyPrefix = all.KeyPrefix
	o.Metadata = all.Metadata
	if !all.StorageClass.IsValid() {
		hasInvalidField = true
	} else {
		o.StorageClass = *all.StorageClass
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
