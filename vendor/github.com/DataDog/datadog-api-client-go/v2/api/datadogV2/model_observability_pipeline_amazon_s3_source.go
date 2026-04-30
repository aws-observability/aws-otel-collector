// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAmazonS3Source The `amazon_s3` source ingests logs from an Amazon S3 bucket.
// It supports AWS authentication, TLS encryption, and configurable compression.
//
// **Supported pipeline types:** logs
type ObservabilityPipelineAmazonS3Source struct {
	// AWS authentication credentials used for accessing AWS services such as S3.
	// If omitted, the system’s default credentials are used (for example, the IAM role and environment variables).
	Auth *ObservabilityPipelineAwsAuth `json:"auth,omitempty"`
	// Compression format for objects retrieved from the S3 bucket. Use `auto` to detect compression from the object's Content-Encoding header or file extension.
	Compression *ObservabilityPipelineAmazonS3SourceCompression `json:"compression,omitempty"`
	// The unique identifier for this component. Used in other parts of the pipeline to reference this component (for example, as the `input` to downstream components).
	Id string `json:"id"`
	// AWS region where the S3 bucket resides.
	Region string `json:"region"`
	// Configuration for enabling TLS encryption between the pipeline component and external services.
	Tls *ObservabilityPipelineTls `json:"tls,omitempty"`
	// The source type. Always `amazon_s3`.
	Type ObservabilityPipelineAmazonS3SourceType `json:"type"`
	// Name of the environment variable or secret that holds the S3 bucket URL.
	UrlKey *string `json:"url_key,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineAmazonS3Source instantiates a new ObservabilityPipelineAmazonS3Source object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineAmazonS3Source(id string, region string, typeVar ObservabilityPipelineAmazonS3SourceType) *ObservabilityPipelineAmazonS3Source {
	this := ObservabilityPipelineAmazonS3Source{}
	this.Id = id
	this.Region = region
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineAmazonS3SourceWithDefaults instantiates a new ObservabilityPipelineAmazonS3Source object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineAmazonS3SourceWithDefaults() *ObservabilityPipelineAmazonS3Source {
	this := ObservabilityPipelineAmazonS3Source{}
	var typeVar ObservabilityPipelineAmazonS3SourceType = OBSERVABILITYPIPELINEAMAZONS3SOURCETYPE_AMAZON_S3
	this.Type = typeVar
	return &this
}

// GetAuth returns the Auth field value if set, zero value otherwise.
func (o *ObservabilityPipelineAmazonS3Source) GetAuth() ObservabilityPipelineAwsAuth {
	if o == nil || o.Auth == nil {
		var ret ObservabilityPipelineAwsAuth
		return ret
	}
	return *o.Auth
}

// GetAuthOk returns a tuple with the Auth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonS3Source) GetAuthOk() (*ObservabilityPipelineAwsAuth, bool) {
	if o == nil || o.Auth == nil {
		return nil, false
	}
	return o.Auth, true
}

// HasAuth returns a boolean if a field has been set.
func (o *ObservabilityPipelineAmazonS3Source) HasAuth() bool {
	return o != nil && o.Auth != nil
}

// SetAuth gets a reference to the given ObservabilityPipelineAwsAuth and assigns it to the Auth field.
func (o *ObservabilityPipelineAmazonS3Source) SetAuth(v ObservabilityPipelineAwsAuth) {
	o.Auth = &v
}

// GetCompression returns the Compression field value if set, zero value otherwise.
func (o *ObservabilityPipelineAmazonS3Source) GetCompression() ObservabilityPipelineAmazonS3SourceCompression {
	if o == nil || o.Compression == nil {
		var ret ObservabilityPipelineAmazonS3SourceCompression
		return ret
	}
	return *o.Compression
}

// GetCompressionOk returns a tuple with the Compression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonS3Source) GetCompressionOk() (*ObservabilityPipelineAmazonS3SourceCompression, bool) {
	if o == nil || o.Compression == nil {
		return nil, false
	}
	return o.Compression, true
}

// HasCompression returns a boolean if a field has been set.
func (o *ObservabilityPipelineAmazonS3Source) HasCompression() bool {
	return o != nil && o.Compression != nil
}

// SetCompression gets a reference to the given ObservabilityPipelineAmazonS3SourceCompression and assigns it to the Compression field.
func (o *ObservabilityPipelineAmazonS3Source) SetCompression(v ObservabilityPipelineAmazonS3SourceCompression) {
	o.Compression = &v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineAmazonS3Source) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonS3Source) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineAmazonS3Source) SetId(v string) {
	o.Id = v
}

// GetRegion returns the Region field value.
func (o *ObservabilityPipelineAmazonS3Source) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonS3Source) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value.
func (o *ObservabilityPipelineAmazonS3Source) SetRegion(v string) {
	o.Region = v
}

// GetTls returns the Tls field value if set, zero value otherwise.
func (o *ObservabilityPipelineAmazonS3Source) GetTls() ObservabilityPipelineTls {
	if o == nil || o.Tls == nil {
		var ret ObservabilityPipelineTls
		return ret
	}
	return *o.Tls
}

// GetTlsOk returns a tuple with the Tls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonS3Source) GetTlsOk() (*ObservabilityPipelineTls, bool) {
	if o == nil || o.Tls == nil {
		return nil, false
	}
	return o.Tls, true
}

// HasTls returns a boolean if a field has been set.
func (o *ObservabilityPipelineAmazonS3Source) HasTls() bool {
	return o != nil && o.Tls != nil
}

// SetTls gets a reference to the given ObservabilityPipelineTls and assigns it to the Tls field.
func (o *ObservabilityPipelineAmazonS3Source) SetTls(v ObservabilityPipelineTls) {
	o.Tls = &v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineAmazonS3Source) GetType() ObservabilityPipelineAmazonS3SourceType {
	if o == nil {
		var ret ObservabilityPipelineAmazonS3SourceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonS3Source) GetTypeOk() (*ObservabilityPipelineAmazonS3SourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineAmazonS3Source) SetType(v ObservabilityPipelineAmazonS3SourceType) {
	o.Type = v
}

// GetUrlKey returns the UrlKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineAmazonS3Source) GetUrlKey() string {
	if o == nil || o.UrlKey == nil {
		var ret string
		return ret
	}
	return *o.UrlKey
}

// GetUrlKeyOk returns a tuple with the UrlKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonS3Source) GetUrlKeyOk() (*string, bool) {
	if o == nil || o.UrlKey == nil {
		return nil, false
	}
	return o.UrlKey, true
}

// HasUrlKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineAmazonS3Source) HasUrlKey() bool {
	return o != nil && o.UrlKey != nil
}

// SetUrlKey gets a reference to the given string and assigns it to the UrlKey field.
func (o *ObservabilityPipelineAmazonS3Source) SetUrlKey(v string) {
	o.UrlKey = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineAmazonS3Source) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Auth != nil {
		toSerialize["auth"] = o.Auth
	}
	if o.Compression != nil {
		toSerialize["compression"] = o.Compression
	}
	toSerialize["id"] = o.Id
	toSerialize["region"] = o.Region
	if o.Tls != nil {
		toSerialize["tls"] = o.Tls
	}
	toSerialize["type"] = o.Type
	if o.UrlKey != nil {
		toSerialize["url_key"] = o.UrlKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineAmazonS3Source) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Auth        *ObservabilityPipelineAwsAuth                   `json:"auth,omitempty"`
		Compression *ObservabilityPipelineAmazonS3SourceCompression `json:"compression,omitempty"`
		Id          *string                                         `json:"id"`
		Region      *string                                         `json:"region"`
		Tls         *ObservabilityPipelineTls                       `json:"tls,omitempty"`
		Type        *ObservabilityPipelineAmazonS3SourceType        `json:"type"`
		UrlKey      *string                                         `json:"url_key,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Region == nil {
		return fmt.Errorf("required field region missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auth", "compression", "id", "region", "tls", "type", "url_key"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Auth != nil && all.Auth.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Auth = all.Auth
	if all.Compression != nil && !all.Compression.IsValid() {
		hasInvalidField = true
	} else {
		o.Compression = all.Compression
	}
	o.Id = *all.Id
	o.Region = *all.Region
	if all.Tls != nil && all.Tls.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Tls = all.Tls
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.UrlKey = all.UrlKey

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
