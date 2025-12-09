// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAmazonDataFirehoseSource The `amazon_data_firehose` source ingests logs from AWS Data Firehose.
type ObservabilityPipelineAmazonDataFirehoseSource struct {
	// AWS authentication credentials used for accessing AWS services such as S3.
	// If omitted, the system’s default credentials are used (for example, the IAM role and environment variables).
	Auth *ObservabilityPipelineAwsAuth `json:"auth,omitempty"`
	// The unique identifier for this component. Used to reference this component in other parts of the pipeline (e.g., as input to downstream components).
	Id string `json:"id"`
	// Configuration for enabling TLS encryption between the pipeline component and external services.
	Tls *ObservabilityPipelineTls `json:"tls,omitempty"`
	// The source type. The value should always be `amazon_data_firehose`.
	Type ObservabilityPipelineAmazonDataFirehoseSourceType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineAmazonDataFirehoseSource instantiates a new ObservabilityPipelineAmazonDataFirehoseSource object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineAmazonDataFirehoseSource(id string, typeVar ObservabilityPipelineAmazonDataFirehoseSourceType) *ObservabilityPipelineAmazonDataFirehoseSource {
	this := ObservabilityPipelineAmazonDataFirehoseSource{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineAmazonDataFirehoseSourceWithDefaults instantiates a new ObservabilityPipelineAmazonDataFirehoseSource object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineAmazonDataFirehoseSourceWithDefaults() *ObservabilityPipelineAmazonDataFirehoseSource {
	this := ObservabilityPipelineAmazonDataFirehoseSource{}
	var typeVar ObservabilityPipelineAmazonDataFirehoseSourceType = OBSERVABILITYPIPELINEAMAZONDATAFIREHOSESOURCETYPE_AMAZON_DATA_FIREHOSE
	this.Type = typeVar
	return &this
}

// GetAuth returns the Auth field value if set, zero value otherwise.
func (o *ObservabilityPipelineAmazonDataFirehoseSource) GetAuth() ObservabilityPipelineAwsAuth {
	if o == nil || o.Auth == nil {
		var ret ObservabilityPipelineAwsAuth
		return ret
	}
	return *o.Auth
}

// GetAuthOk returns a tuple with the Auth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonDataFirehoseSource) GetAuthOk() (*ObservabilityPipelineAwsAuth, bool) {
	if o == nil || o.Auth == nil {
		return nil, false
	}
	return o.Auth, true
}

// HasAuth returns a boolean if a field has been set.
func (o *ObservabilityPipelineAmazonDataFirehoseSource) HasAuth() bool {
	return o != nil && o.Auth != nil
}

// SetAuth gets a reference to the given ObservabilityPipelineAwsAuth and assigns it to the Auth field.
func (o *ObservabilityPipelineAmazonDataFirehoseSource) SetAuth(v ObservabilityPipelineAwsAuth) {
	o.Auth = &v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineAmazonDataFirehoseSource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonDataFirehoseSource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineAmazonDataFirehoseSource) SetId(v string) {
	o.Id = v
}

// GetTls returns the Tls field value if set, zero value otherwise.
func (o *ObservabilityPipelineAmazonDataFirehoseSource) GetTls() ObservabilityPipelineTls {
	if o == nil || o.Tls == nil {
		var ret ObservabilityPipelineTls
		return ret
	}
	return *o.Tls
}

// GetTlsOk returns a tuple with the Tls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonDataFirehoseSource) GetTlsOk() (*ObservabilityPipelineTls, bool) {
	if o == nil || o.Tls == nil {
		return nil, false
	}
	return o.Tls, true
}

// HasTls returns a boolean if a field has been set.
func (o *ObservabilityPipelineAmazonDataFirehoseSource) HasTls() bool {
	return o != nil && o.Tls != nil
}

// SetTls gets a reference to the given ObservabilityPipelineTls and assigns it to the Tls field.
func (o *ObservabilityPipelineAmazonDataFirehoseSource) SetTls(v ObservabilityPipelineTls) {
	o.Tls = &v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineAmazonDataFirehoseSource) GetType() ObservabilityPipelineAmazonDataFirehoseSourceType {
	if o == nil {
		var ret ObservabilityPipelineAmazonDataFirehoseSourceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonDataFirehoseSource) GetTypeOk() (*ObservabilityPipelineAmazonDataFirehoseSourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineAmazonDataFirehoseSource) SetType(v ObservabilityPipelineAmazonDataFirehoseSourceType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineAmazonDataFirehoseSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Auth != nil {
		toSerialize["auth"] = o.Auth
	}
	toSerialize["id"] = o.Id
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
func (o *ObservabilityPipelineAmazonDataFirehoseSource) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Auth *ObservabilityPipelineAwsAuth                      `json:"auth,omitempty"`
		Id   *string                                            `json:"id"`
		Tls  *ObservabilityPipelineTls                          `json:"tls,omitempty"`
		Type *ObservabilityPipelineAmazonDataFirehoseSourceType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auth", "id", "tls", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Auth != nil && all.Auth.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Auth = all.Auth
	o.Id = *all.Id
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
