// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAmazonS3Destination The `amazon_s3` destination sends your logs in Datadog-rehydratable format to an Amazon S3 bucket for archiving.
//
// **Supported pipeline types:** logs
type ObservabilityPipelineAmazonS3Destination struct {
	// AWS authentication credentials used for accessing AWS services such as S3.
	// If omitted, the system’s default credentials are used (for example, the IAM role and environment variables).
	Auth *ObservabilityPipelineAwsAuth `json:"auth,omitempty"`
	// S3 bucket name.
	Bucket string `json:"bucket"`
	// Configuration for buffer settings on destination components.
	Buffer *ObservabilityPipelineBufferOptions `json:"buffer,omitempty"`
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
	// Configuration for enabling TLS encryption between the pipeline component and external services.
	Tls *ObservabilityPipelineTls `json:"tls,omitempty"`
	// The destination type. Always `amazon_s3`.
	Type ObservabilityPipelineAmazonS3DestinationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineAmazonS3Destination instantiates a new ObservabilityPipelineAmazonS3Destination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineAmazonS3Destination(bucket string, id string, inputs []string, region string, storageClass ObservabilityPipelineAmazonS3DestinationStorageClass, typeVar ObservabilityPipelineAmazonS3DestinationType) *ObservabilityPipelineAmazonS3Destination {
	this := ObservabilityPipelineAmazonS3Destination{}
	this.Bucket = bucket
	this.Id = id
	this.Inputs = inputs
	this.Region = region
	this.StorageClass = storageClass
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineAmazonS3DestinationWithDefaults instantiates a new ObservabilityPipelineAmazonS3Destination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineAmazonS3DestinationWithDefaults() *ObservabilityPipelineAmazonS3Destination {
	this := ObservabilityPipelineAmazonS3Destination{}
	var typeVar ObservabilityPipelineAmazonS3DestinationType = OBSERVABILITYPIPELINEAMAZONS3DESTINATIONTYPE_AMAZON_S3
	this.Type = typeVar
	return &this
}

// GetAuth returns the Auth field value if set, zero value otherwise.
func (o *ObservabilityPipelineAmazonS3Destination) GetAuth() ObservabilityPipelineAwsAuth {
	if o == nil || o.Auth == nil {
		var ret ObservabilityPipelineAwsAuth
		return ret
	}
	return *o.Auth
}

// GetAuthOk returns a tuple with the Auth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonS3Destination) GetAuthOk() (*ObservabilityPipelineAwsAuth, bool) {
	if o == nil || o.Auth == nil {
		return nil, false
	}
	return o.Auth, true
}

// HasAuth returns a boolean if a field has been set.
func (o *ObservabilityPipelineAmazonS3Destination) HasAuth() bool {
	return o != nil && o.Auth != nil
}

// SetAuth gets a reference to the given ObservabilityPipelineAwsAuth and assigns it to the Auth field.
func (o *ObservabilityPipelineAmazonS3Destination) SetAuth(v ObservabilityPipelineAwsAuth) {
	o.Auth = &v
}

// GetBucket returns the Bucket field value.
func (o *ObservabilityPipelineAmazonS3Destination) GetBucket() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonS3Destination) GetBucketOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bucket, true
}

// SetBucket sets field value.
func (o *ObservabilityPipelineAmazonS3Destination) SetBucket(v string) {
	o.Bucket = v
}

// GetBuffer returns the Buffer field value if set, zero value otherwise.
func (o *ObservabilityPipelineAmazonS3Destination) GetBuffer() ObservabilityPipelineBufferOptions {
	if o == nil || o.Buffer == nil {
		var ret ObservabilityPipelineBufferOptions
		return ret
	}
	return *o.Buffer
}

// GetBufferOk returns a tuple with the Buffer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonS3Destination) GetBufferOk() (*ObservabilityPipelineBufferOptions, bool) {
	if o == nil || o.Buffer == nil {
		return nil, false
	}
	return o.Buffer, true
}

// HasBuffer returns a boolean if a field has been set.
func (o *ObservabilityPipelineAmazonS3Destination) HasBuffer() bool {
	return o != nil && o.Buffer != nil
}

// SetBuffer gets a reference to the given ObservabilityPipelineBufferOptions and assigns it to the Buffer field.
func (o *ObservabilityPipelineAmazonS3Destination) SetBuffer(v ObservabilityPipelineBufferOptions) {
	o.Buffer = &v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineAmazonS3Destination) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonS3Destination) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineAmazonS3Destination) SetId(v string) {
	o.Id = v
}

// GetInputs returns the Inputs field value.
func (o *ObservabilityPipelineAmazonS3Destination) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonS3Destination) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *ObservabilityPipelineAmazonS3Destination) SetInputs(v []string) {
	o.Inputs = v
}

// GetKeyPrefix returns the KeyPrefix field value if set, zero value otherwise.
func (o *ObservabilityPipelineAmazonS3Destination) GetKeyPrefix() string {
	if o == nil || o.KeyPrefix == nil {
		var ret string
		return ret
	}
	return *o.KeyPrefix
}

// GetKeyPrefixOk returns a tuple with the KeyPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonS3Destination) GetKeyPrefixOk() (*string, bool) {
	if o == nil || o.KeyPrefix == nil {
		return nil, false
	}
	return o.KeyPrefix, true
}

// HasKeyPrefix returns a boolean if a field has been set.
func (o *ObservabilityPipelineAmazonS3Destination) HasKeyPrefix() bool {
	return o != nil && o.KeyPrefix != nil
}

// SetKeyPrefix gets a reference to the given string and assigns it to the KeyPrefix field.
func (o *ObservabilityPipelineAmazonS3Destination) SetKeyPrefix(v string) {
	o.KeyPrefix = &v
}

// GetRegion returns the Region field value.
func (o *ObservabilityPipelineAmazonS3Destination) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonS3Destination) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value.
func (o *ObservabilityPipelineAmazonS3Destination) SetRegion(v string) {
	o.Region = v
}

// GetStorageClass returns the StorageClass field value.
func (o *ObservabilityPipelineAmazonS3Destination) GetStorageClass() ObservabilityPipelineAmazonS3DestinationStorageClass {
	if o == nil {
		var ret ObservabilityPipelineAmazonS3DestinationStorageClass
		return ret
	}
	return o.StorageClass
}

// GetStorageClassOk returns a tuple with the StorageClass field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonS3Destination) GetStorageClassOk() (*ObservabilityPipelineAmazonS3DestinationStorageClass, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageClass, true
}

// SetStorageClass sets field value.
func (o *ObservabilityPipelineAmazonS3Destination) SetStorageClass(v ObservabilityPipelineAmazonS3DestinationStorageClass) {
	o.StorageClass = v
}

// GetTls returns the Tls field value if set, zero value otherwise.
func (o *ObservabilityPipelineAmazonS3Destination) GetTls() ObservabilityPipelineTls {
	if o == nil || o.Tls == nil {
		var ret ObservabilityPipelineTls
		return ret
	}
	return *o.Tls
}

// GetTlsOk returns a tuple with the Tls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonS3Destination) GetTlsOk() (*ObservabilityPipelineTls, bool) {
	if o == nil || o.Tls == nil {
		return nil, false
	}
	return o.Tls, true
}

// HasTls returns a boolean if a field has been set.
func (o *ObservabilityPipelineAmazonS3Destination) HasTls() bool {
	return o != nil && o.Tls != nil
}

// SetTls gets a reference to the given ObservabilityPipelineTls and assigns it to the Tls field.
func (o *ObservabilityPipelineAmazonS3Destination) SetTls(v ObservabilityPipelineTls) {
	o.Tls = &v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineAmazonS3Destination) GetType() ObservabilityPipelineAmazonS3DestinationType {
	if o == nil {
		var ret ObservabilityPipelineAmazonS3DestinationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonS3Destination) GetTypeOk() (*ObservabilityPipelineAmazonS3DestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineAmazonS3Destination) SetType(v ObservabilityPipelineAmazonS3DestinationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineAmazonS3Destination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Auth != nil {
		toSerialize["auth"] = o.Auth
	}
	toSerialize["bucket"] = o.Bucket
	if o.Buffer != nil {
		toSerialize["buffer"] = o.Buffer
	}
	toSerialize["id"] = o.Id
	toSerialize["inputs"] = o.Inputs
	if o.KeyPrefix != nil {
		toSerialize["key_prefix"] = o.KeyPrefix
	}
	toSerialize["region"] = o.Region
	toSerialize["storage_class"] = o.StorageClass
	if o.Tls != nil {
		toSerialize["tls"] = o.Tls
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineAmazonS3Destination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Auth         *ObservabilityPipelineAwsAuth                         `json:"auth,omitempty"`
		Bucket       *string                                               `json:"bucket"`
		Buffer       *ObservabilityPipelineBufferOptions                   `json:"buffer,omitempty"`
		Id           *string                                               `json:"id"`
		Inputs       *[]string                                             `json:"inputs"`
		KeyPrefix    *string                                               `json:"key_prefix,omitempty"`
		Region       *string                                               `json:"region"`
		StorageClass *ObservabilityPipelineAmazonS3DestinationStorageClass `json:"storage_class"`
		Tls          *ObservabilityPipelineTls                             `json:"tls,omitempty"`
		Type         *ObservabilityPipelineAmazonS3DestinationType         `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
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
		datadog.DeleteKeys(additionalProperties, &[]string{"auth", "bucket", "buffer", "id", "inputs", "key_prefix", "region", "storage_class", "tls", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Auth != nil && all.Auth.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Auth = all.Auth
	o.Bucket = *all.Bucket
	o.Buffer = all.Buffer
	o.Id = *all.Id
	o.Inputs = *all.Inputs
	o.KeyPrefix = all.KeyPrefix
	o.Region = *all.Region
	if !all.StorageClass.IsValid() {
		hasInvalidField = true
	} else {
		o.StorageClass = *all.StorageClass
	}
	if all.Tls != nil && all.Tls.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Tls = all.Tls
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
