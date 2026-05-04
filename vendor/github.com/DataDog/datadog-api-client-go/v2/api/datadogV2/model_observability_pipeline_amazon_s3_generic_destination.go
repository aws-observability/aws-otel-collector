// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAmazonS3GenericDestination The `amazon_s3_generic` destination sends your logs to an Amazon S3 bucket.
//
// **Supported pipeline types:** logs
type ObservabilityPipelineAmazonS3GenericDestination struct {
	// AWS authentication credentials used for accessing AWS services such as S3.
	// If omitted, the system’s default credentials are used (for example, the IAM role and environment variables).
	Auth *ObservabilityPipelineAwsAuth `json:"auth,omitempty"`
	// Event batching settings
	BatchSettings *ObservabilityPipelineAmazonS3GenericBatchSettings `json:"batch_settings,omitempty"`
	// S3 bucket name.
	Bucket string `json:"bucket"`
	// Compression algorithm applied to encoded logs.
	Compression ObservabilityPipelineAmazonS3GenericCompression `json:"compression"`
	// Encoding format for the destination.
	Encoding ObservabilityPipelineAmazonS3GenericEncoding `json:"encoding"`
	// Unique identifier for the destination component.
	Id string `json:"id"`
	// A list of component IDs whose output is used as the `input` for this component.
	Inputs []string `json:"inputs"`
	// Optional prefix for object keys.
	KeyPrefix *string `json:"key_prefix,omitempty"`
	// AWS region of the S3 bucket.
	Region string `json:"region"`
	// S3 storage class.
	StorageClass ObservabilityPipelineAmazonS3DestinationStorageClass `json:"storage_class"`
	// The destination type. Always `amazon_s3_generic`.
	Type ObservabilityPipelineAmazonS3GenericDestinationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineAmazonS3GenericDestination instantiates a new ObservabilityPipelineAmazonS3GenericDestination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineAmazonS3GenericDestination(bucket string, compression ObservabilityPipelineAmazonS3GenericCompression, encoding ObservabilityPipelineAmazonS3GenericEncoding, id string, inputs []string, region string, storageClass ObservabilityPipelineAmazonS3DestinationStorageClass, typeVar ObservabilityPipelineAmazonS3GenericDestinationType) *ObservabilityPipelineAmazonS3GenericDestination {
	this := ObservabilityPipelineAmazonS3GenericDestination{}
	this.Bucket = bucket
	this.Compression = compression
	this.Encoding = encoding
	this.Id = id
	this.Inputs = inputs
	this.Region = region
	this.StorageClass = storageClass
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineAmazonS3GenericDestinationWithDefaults instantiates a new ObservabilityPipelineAmazonS3GenericDestination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineAmazonS3GenericDestinationWithDefaults() *ObservabilityPipelineAmazonS3GenericDestination {
	this := ObservabilityPipelineAmazonS3GenericDestination{}
	var typeVar ObservabilityPipelineAmazonS3GenericDestinationType = OBSERVABILITYPIPELINEAMAZONS3GENERICDESTINATIONTYPE_GENERIC_ARCHIVES_S3
	this.Type = typeVar
	return &this
}

// GetAuth returns the Auth field value if set, zero value otherwise.
func (o *ObservabilityPipelineAmazonS3GenericDestination) GetAuth() ObservabilityPipelineAwsAuth {
	if o == nil || o.Auth == nil {
		var ret ObservabilityPipelineAwsAuth
		return ret
	}
	return *o.Auth
}

// GetAuthOk returns a tuple with the Auth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonS3GenericDestination) GetAuthOk() (*ObservabilityPipelineAwsAuth, bool) {
	if o == nil || o.Auth == nil {
		return nil, false
	}
	return o.Auth, true
}

// HasAuth returns a boolean if a field has been set.
func (o *ObservabilityPipelineAmazonS3GenericDestination) HasAuth() bool {
	return o != nil && o.Auth != nil
}

// SetAuth gets a reference to the given ObservabilityPipelineAwsAuth and assigns it to the Auth field.
func (o *ObservabilityPipelineAmazonS3GenericDestination) SetAuth(v ObservabilityPipelineAwsAuth) {
	o.Auth = &v
}

// GetBatchSettings returns the BatchSettings field value if set, zero value otherwise.
func (o *ObservabilityPipelineAmazonS3GenericDestination) GetBatchSettings() ObservabilityPipelineAmazonS3GenericBatchSettings {
	if o == nil || o.BatchSettings == nil {
		var ret ObservabilityPipelineAmazonS3GenericBatchSettings
		return ret
	}
	return *o.BatchSettings
}

// GetBatchSettingsOk returns a tuple with the BatchSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonS3GenericDestination) GetBatchSettingsOk() (*ObservabilityPipelineAmazonS3GenericBatchSettings, bool) {
	if o == nil || o.BatchSettings == nil {
		return nil, false
	}
	return o.BatchSettings, true
}

// HasBatchSettings returns a boolean if a field has been set.
func (o *ObservabilityPipelineAmazonS3GenericDestination) HasBatchSettings() bool {
	return o != nil && o.BatchSettings != nil
}

// SetBatchSettings gets a reference to the given ObservabilityPipelineAmazonS3GenericBatchSettings and assigns it to the BatchSettings field.
func (o *ObservabilityPipelineAmazonS3GenericDestination) SetBatchSettings(v ObservabilityPipelineAmazonS3GenericBatchSettings) {
	o.BatchSettings = &v
}

// GetBucket returns the Bucket field value.
func (o *ObservabilityPipelineAmazonS3GenericDestination) GetBucket() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonS3GenericDestination) GetBucketOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bucket, true
}

// SetBucket sets field value.
func (o *ObservabilityPipelineAmazonS3GenericDestination) SetBucket(v string) {
	o.Bucket = v
}

// GetCompression returns the Compression field value.
func (o *ObservabilityPipelineAmazonS3GenericDestination) GetCompression() ObservabilityPipelineAmazonS3GenericCompression {
	if o == nil {
		var ret ObservabilityPipelineAmazonS3GenericCompression
		return ret
	}
	return o.Compression
}

// GetCompressionOk returns a tuple with the Compression field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonS3GenericDestination) GetCompressionOk() (*ObservabilityPipelineAmazonS3GenericCompression, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Compression, true
}

// SetCompression sets field value.
func (o *ObservabilityPipelineAmazonS3GenericDestination) SetCompression(v ObservabilityPipelineAmazonS3GenericCompression) {
	o.Compression = v
}

// GetEncoding returns the Encoding field value.
func (o *ObservabilityPipelineAmazonS3GenericDestination) GetEncoding() ObservabilityPipelineAmazonS3GenericEncoding {
	if o == nil {
		var ret ObservabilityPipelineAmazonS3GenericEncoding
		return ret
	}
	return o.Encoding
}

// GetEncodingOk returns a tuple with the Encoding field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonS3GenericDestination) GetEncodingOk() (*ObservabilityPipelineAmazonS3GenericEncoding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Encoding, true
}

// SetEncoding sets field value.
func (o *ObservabilityPipelineAmazonS3GenericDestination) SetEncoding(v ObservabilityPipelineAmazonS3GenericEncoding) {
	o.Encoding = v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineAmazonS3GenericDestination) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonS3GenericDestination) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineAmazonS3GenericDestination) SetId(v string) {
	o.Id = v
}

// GetInputs returns the Inputs field value.
func (o *ObservabilityPipelineAmazonS3GenericDestination) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonS3GenericDestination) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *ObservabilityPipelineAmazonS3GenericDestination) SetInputs(v []string) {
	o.Inputs = v
}

// GetKeyPrefix returns the KeyPrefix field value if set, zero value otherwise.
func (o *ObservabilityPipelineAmazonS3GenericDestination) GetKeyPrefix() string {
	if o == nil || o.KeyPrefix == nil {
		var ret string
		return ret
	}
	return *o.KeyPrefix
}

// GetKeyPrefixOk returns a tuple with the KeyPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonS3GenericDestination) GetKeyPrefixOk() (*string, bool) {
	if o == nil || o.KeyPrefix == nil {
		return nil, false
	}
	return o.KeyPrefix, true
}

// HasKeyPrefix returns a boolean if a field has been set.
func (o *ObservabilityPipelineAmazonS3GenericDestination) HasKeyPrefix() bool {
	return o != nil && o.KeyPrefix != nil
}

// SetKeyPrefix gets a reference to the given string and assigns it to the KeyPrefix field.
func (o *ObservabilityPipelineAmazonS3GenericDestination) SetKeyPrefix(v string) {
	o.KeyPrefix = &v
}

// GetRegion returns the Region field value.
func (o *ObservabilityPipelineAmazonS3GenericDestination) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonS3GenericDestination) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value.
func (o *ObservabilityPipelineAmazonS3GenericDestination) SetRegion(v string) {
	o.Region = v
}

// GetStorageClass returns the StorageClass field value.
func (o *ObservabilityPipelineAmazonS3GenericDestination) GetStorageClass() ObservabilityPipelineAmazonS3DestinationStorageClass {
	if o == nil {
		var ret ObservabilityPipelineAmazonS3DestinationStorageClass
		return ret
	}
	return o.StorageClass
}

// GetStorageClassOk returns a tuple with the StorageClass field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonS3GenericDestination) GetStorageClassOk() (*ObservabilityPipelineAmazonS3DestinationStorageClass, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageClass, true
}

// SetStorageClass sets field value.
func (o *ObservabilityPipelineAmazonS3GenericDestination) SetStorageClass(v ObservabilityPipelineAmazonS3DestinationStorageClass) {
	o.StorageClass = v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineAmazonS3GenericDestination) GetType() ObservabilityPipelineAmazonS3GenericDestinationType {
	if o == nil {
		var ret ObservabilityPipelineAmazonS3GenericDestinationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonS3GenericDestination) GetTypeOk() (*ObservabilityPipelineAmazonS3GenericDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineAmazonS3GenericDestination) SetType(v ObservabilityPipelineAmazonS3GenericDestinationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineAmazonS3GenericDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Auth != nil {
		toSerialize["auth"] = o.Auth
	}
	if o.BatchSettings != nil {
		toSerialize["batch_settings"] = o.BatchSettings
	}
	toSerialize["bucket"] = o.Bucket
	toSerialize["compression"] = o.Compression
	toSerialize["encoding"] = o.Encoding
	toSerialize["id"] = o.Id
	toSerialize["inputs"] = o.Inputs
	if o.KeyPrefix != nil {
		toSerialize["key_prefix"] = o.KeyPrefix
	}
	toSerialize["region"] = o.Region
	toSerialize["storage_class"] = o.StorageClass
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineAmazonS3GenericDestination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Auth          *ObservabilityPipelineAwsAuth                         `json:"auth,omitempty"`
		BatchSettings *ObservabilityPipelineAmazonS3GenericBatchSettings    `json:"batch_settings,omitempty"`
		Bucket        *string                                               `json:"bucket"`
		Compression   *ObservabilityPipelineAmazonS3GenericCompression      `json:"compression"`
		Encoding      *ObservabilityPipelineAmazonS3GenericEncoding         `json:"encoding"`
		Id            *string                                               `json:"id"`
		Inputs        *[]string                                             `json:"inputs"`
		KeyPrefix     *string                                               `json:"key_prefix,omitempty"`
		Region        *string                                               `json:"region"`
		StorageClass  *ObservabilityPipelineAmazonS3DestinationStorageClass `json:"storage_class"`
		Type          *ObservabilityPipelineAmazonS3GenericDestinationType  `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Bucket == nil {
		return fmt.Errorf("required field bucket missing")
	}
	if all.Compression == nil {
		return fmt.Errorf("required field compression missing")
	}
	if all.Encoding == nil {
		return fmt.Errorf("required field encoding missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Inputs == nil {
		return fmt.Errorf("required field inputs missing")
	}
	if all.Region == nil {
		return fmt.Errorf("required field region missing")
	}
	if all.StorageClass == nil {
		return fmt.Errorf("required field storage_class missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auth", "batch_settings", "bucket", "compression", "encoding", "id", "inputs", "key_prefix", "region", "storage_class", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Auth != nil && all.Auth.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Auth = all.Auth
	if all.BatchSettings != nil && all.BatchSettings.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.BatchSettings = all.BatchSettings
	o.Bucket = *all.Bucket
	o.Compression = *all.Compression
	o.Encoding = *all.Encoding
	o.Id = *all.Id
	o.Inputs = *all.Inputs
	o.KeyPrefix = all.KeyPrefix
	o.Region = *all.Region
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
