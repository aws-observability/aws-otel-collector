// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineGooglePubSubSource The `google_pubsub` source ingests logs from a Google Cloud Pub/Sub subscription.
//
// **Supported pipeline types:** logs
type ObservabilityPipelineGooglePubSubSource struct {
	// Google Cloud credentials used to authenticate with Google Cloud Storage.
	Auth *ObservabilityPipelineGcpAuth `json:"auth,omitempty"`
	// The decoding format used to interpret incoming logs.
	Decoding ObservabilityPipelineDecoding `json:"decoding"`
	// The unique identifier for this component. Used in other parts of the pipeline to reference this component (for example, as the `input` to downstream components).
	Id string `json:"id"`
	// The Google Cloud project ID that owns the Pub/Sub subscription.
	Project string `json:"project"`
	// The Pub/Sub subscription name from which messages are consumed.
	Subscription string `json:"subscription"`
	// Configuration for enabling TLS encryption between the pipeline component and external services.
	Tls *ObservabilityPipelineTls `json:"tls,omitempty"`
	// The source type. The value should always be `google_pubsub`.
	Type ObservabilityPipelineGooglePubSubSourceType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineGooglePubSubSource instantiates a new ObservabilityPipelineGooglePubSubSource object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineGooglePubSubSource(decoding ObservabilityPipelineDecoding, id string, project string, subscription string, typeVar ObservabilityPipelineGooglePubSubSourceType) *ObservabilityPipelineGooglePubSubSource {
	this := ObservabilityPipelineGooglePubSubSource{}
	this.Decoding = decoding
	this.Id = id
	this.Project = project
	this.Subscription = subscription
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineGooglePubSubSourceWithDefaults instantiates a new ObservabilityPipelineGooglePubSubSource object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineGooglePubSubSourceWithDefaults() *ObservabilityPipelineGooglePubSubSource {
	this := ObservabilityPipelineGooglePubSubSource{}
	var typeVar ObservabilityPipelineGooglePubSubSourceType = OBSERVABILITYPIPELINEGOOGLEPUBSUBSOURCETYPE_GOOGLE_PUBSUB
	this.Type = typeVar
	return &this
}

// GetAuth returns the Auth field value if set, zero value otherwise.
func (o *ObservabilityPipelineGooglePubSubSource) GetAuth() ObservabilityPipelineGcpAuth {
	if o == nil || o.Auth == nil {
		var ret ObservabilityPipelineGcpAuth
		return ret
	}
	return *o.Auth
}

// GetAuthOk returns a tuple with the Auth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGooglePubSubSource) GetAuthOk() (*ObservabilityPipelineGcpAuth, bool) {
	if o == nil || o.Auth == nil {
		return nil, false
	}
	return o.Auth, true
}

// HasAuth returns a boolean if a field has been set.
func (o *ObservabilityPipelineGooglePubSubSource) HasAuth() bool {
	return o != nil && o.Auth != nil
}

// SetAuth gets a reference to the given ObservabilityPipelineGcpAuth and assigns it to the Auth field.
func (o *ObservabilityPipelineGooglePubSubSource) SetAuth(v ObservabilityPipelineGcpAuth) {
	o.Auth = &v
}

// GetDecoding returns the Decoding field value.
func (o *ObservabilityPipelineGooglePubSubSource) GetDecoding() ObservabilityPipelineDecoding {
	if o == nil {
		var ret ObservabilityPipelineDecoding
		return ret
	}
	return o.Decoding
}

// GetDecodingOk returns a tuple with the Decoding field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGooglePubSubSource) GetDecodingOk() (*ObservabilityPipelineDecoding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Decoding, true
}

// SetDecoding sets field value.
func (o *ObservabilityPipelineGooglePubSubSource) SetDecoding(v ObservabilityPipelineDecoding) {
	o.Decoding = v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineGooglePubSubSource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGooglePubSubSource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineGooglePubSubSource) SetId(v string) {
	o.Id = v
}

// GetProject returns the Project field value.
func (o *ObservabilityPipelineGooglePubSubSource) GetProject() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGooglePubSubSource) GetProjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value.
func (o *ObservabilityPipelineGooglePubSubSource) SetProject(v string) {
	o.Project = v
}

// GetSubscription returns the Subscription field value.
func (o *ObservabilityPipelineGooglePubSubSource) GetSubscription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGooglePubSubSource) GetSubscriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subscription, true
}

// SetSubscription sets field value.
func (o *ObservabilityPipelineGooglePubSubSource) SetSubscription(v string) {
	o.Subscription = v
}

// GetTls returns the Tls field value if set, zero value otherwise.
func (o *ObservabilityPipelineGooglePubSubSource) GetTls() ObservabilityPipelineTls {
	if o == nil || o.Tls == nil {
		var ret ObservabilityPipelineTls
		return ret
	}
	return *o.Tls
}

// GetTlsOk returns a tuple with the Tls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGooglePubSubSource) GetTlsOk() (*ObservabilityPipelineTls, bool) {
	if o == nil || o.Tls == nil {
		return nil, false
	}
	return o.Tls, true
}

// HasTls returns a boolean if a field has been set.
func (o *ObservabilityPipelineGooglePubSubSource) HasTls() bool {
	return o != nil && o.Tls != nil
}

// SetTls gets a reference to the given ObservabilityPipelineTls and assigns it to the Tls field.
func (o *ObservabilityPipelineGooglePubSubSource) SetTls(v ObservabilityPipelineTls) {
	o.Tls = &v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineGooglePubSubSource) GetType() ObservabilityPipelineGooglePubSubSourceType {
	if o == nil {
		var ret ObservabilityPipelineGooglePubSubSourceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGooglePubSubSource) GetTypeOk() (*ObservabilityPipelineGooglePubSubSourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineGooglePubSubSource) SetType(v ObservabilityPipelineGooglePubSubSourceType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineGooglePubSubSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Auth != nil {
		toSerialize["auth"] = o.Auth
	}
	toSerialize["decoding"] = o.Decoding
	toSerialize["id"] = o.Id
	toSerialize["project"] = o.Project
	toSerialize["subscription"] = o.Subscription
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
func (o *ObservabilityPipelineGooglePubSubSource) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Auth         *ObservabilityPipelineGcpAuth                `json:"auth,omitempty"`
		Decoding     *ObservabilityPipelineDecoding               `json:"decoding"`
		Id           *string                                      `json:"id"`
		Project      *string                                      `json:"project"`
		Subscription *string                                      `json:"subscription"`
		Tls          *ObservabilityPipelineTls                    `json:"tls,omitempty"`
		Type         *ObservabilityPipelineGooglePubSubSourceType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Decoding == nil {
		return fmt.Errorf("required field decoding missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Project == nil {
		return fmt.Errorf("required field project missing")
	}
	if all.Subscription == nil {
		return fmt.Errorf("required field subscription missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auth", "decoding", "id", "project", "subscription", "tls", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Auth != nil && all.Auth.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Auth = all.Auth
	if !all.Decoding.IsValid() {
		hasInvalidField = true
	} else {
		o.Decoding = *all.Decoding
	}
	o.Id = *all.Id
	o.Project = *all.Project
	o.Subscription = *all.Subscription
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
