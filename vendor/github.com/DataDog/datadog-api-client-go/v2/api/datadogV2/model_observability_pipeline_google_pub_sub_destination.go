// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineGooglePubSubDestination The `google_pubsub` destination publishes logs to a Google Cloud Pub/Sub topic.
type ObservabilityPipelineGooglePubSubDestination struct {
	// GCP credentials used to authenticate with Google Cloud Storage.
	Auth *ObservabilityPipelineGcpAuth `json:"auth,omitempty"`
	// Encoding format for log events.
	Encoding ObservabilityPipelineGooglePubSubDestinationEncoding `json:"encoding"`
	// The unique identifier for this component.
	Id string `json:"id"`
	// A list of component IDs whose output is used as the `input` for this component.
	Inputs []string `json:"inputs"`
	// The GCP project ID that owns the Pub/Sub topic.
	Project string `json:"project"`
	// Configuration for enabling TLS encryption between the pipeline component and external services.
	Tls *ObservabilityPipelineTls `json:"tls,omitempty"`
	// The Pub/Sub topic name to publish logs to.
	Topic string `json:"topic"`
	// The destination type. The value should always be `google_pubsub`.
	Type ObservabilityPipelineGooglePubSubDestinationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineGooglePubSubDestination instantiates a new ObservabilityPipelineGooglePubSubDestination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineGooglePubSubDestination(encoding ObservabilityPipelineGooglePubSubDestinationEncoding, id string, inputs []string, project string, topic string, typeVar ObservabilityPipelineGooglePubSubDestinationType) *ObservabilityPipelineGooglePubSubDestination {
	this := ObservabilityPipelineGooglePubSubDestination{}
	this.Encoding = encoding
	this.Id = id
	this.Inputs = inputs
	this.Project = project
	this.Topic = topic
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineGooglePubSubDestinationWithDefaults instantiates a new ObservabilityPipelineGooglePubSubDestination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineGooglePubSubDestinationWithDefaults() *ObservabilityPipelineGooglePubSubDestination {
	this := ObservabilityPipelineGooglePubSubDestination{}
	var typeVar ObservabilityPipelineGooglePubSubDestinationType = OBSERVABILITYPIPELINEGOOGLEPUBSUBDESTINATIONTYPE_GOOGLE_PUBSUB
	this.Type = typeVar
	return &this
}

// GetAuth returns the Auth field value if set, zero value otherwise.
func (o *ObservabilityPipelineGooglePubSubDestination) GetAuth() ObservabilityPipelineGcpAuth {
	if o == nil || o.Auth == nil {
		var ret ObservabilityPipelineGcpAuth
		return ret
	}
	return *o.Auth
}

// GetAuthOk returns a tuple with the Auth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGooglePubSubDestination) GetAuthOk() (*ObservabilityPipelineGcpAuth, bool) {
	if o == nil || o.Auth == nil {
		return nil, false
	}
	return o.Auth, true
}

// HasAuth returns a boolean if a field has been set.
func (o *ObservabilityPipelineGooglePubSubDestination) HasAuth() bool {
	return o != nil && o.Auth != nil
}

// SetAuth gets a reference to the given ObservabilityPipelineGcpAuth and assigns it to the Auth field.
func (o *ObservabilityPipelineGooglePubSubDestination) SetAuth(v ObservabilityPipelineGcpAuth) {
	o.Auth = &v
}

// GetEncoding returns the Encoding field value.
func (o *ObservabilityPipelineGooglePubSubDestination) GetEncoding() ObservabilityPipelineGooglePubSubDestinationEncoding {
	if o == nil {
		var ret ObservabilityPipelineGooglePubSubDestinationEncoding
		return ret
	}
	return o.Encoding
}

// GetEncodingOk returns a tuple with the Encoding field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGooglePubSubDestination) GetEncodingOk() (*ObservabilityPipelineGooglePubSubDestinationEncoding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Encoding, true
}

// SetEncoding sets field value.
func (o *ObservabilityPipelineGooglePubSubDestination) SetEncoding(v ObservabilityPipelineGooglePubSubDestinationEncoding) {
	o.Encoding = v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineGooglePubSubDestination) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGooglePubSubDestination) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineGooglePubSubDestination) SetId(v string) {
	o.Id = v
}

// GetInputs returns the Inputs field value.
func (o *ObservabilityPipelineGooglePubSubDestination) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGooglePubSubDestination) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *ObservabilityPipelineGooglePubSubDestination) SetInputs(v []string) {
	o.Inputs = v
}

// GetProject returns the Project field value.
func (o *ObservabilityPipelineGooglePubSubDestination) GetProject() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGooglePubSubDestination) GetProjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value.
func (o *ObservabilityPipelineGooglePubSubDestination) SetProject(v string) {
	o.Project = v
}

// GetTls returns the Tls field value if set, zero value otherwise.
func (o *ObservabilityPipelineGooglePubSubDestination) GetTls() ObservabilityPipelineTls {
	if o == nil || o.Tls == nil {
		var ret ObservabilityPipelineTls
		return ret
	}
	return *o.Tls
}

// GetTlsOk returns a tuple with the Tls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGooglePubSubDestination) GetTlsOk() (*ObservabilityPipelineTls, bool) {
	if o == nil || o.Tls == nil {
		return nil, false
	}
	return o.Tls, true
}

// HasTls returns a boolean if a field has been set.
func (o *ObservabilityPipelineGooglePubSubDestination) HasTls() bool {
	return o != nil && o.Tls != nil
}

// SetTls gets a reference to the given ObservabilityPipelineTls and assigns it to the Tls field.
func (o *ObservabilityPipelineGooglePubSubDestination) SetTls(v ObservabilityPipelineTls) {
	o.Tls = &v
}

// GetTopic returns the Topic field value.
func (o *ObservabilityPipelineGooglePubSubDestination) GetTopic() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Topic
}

// GetTopicOk returns a tuple with the Topic field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGooglePubSubDestination) GetTopicOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Topic, true
}

// SetTopic sets field value.
func (o *ObservabilityPipelineGooglePubSubDestination) SetTopic(v string) {
	o.Topic = v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineGooglePubSubDestination) GetType() ObservabilityPipelineGooglePubSubDestinationType {
	if o == nil {
		var ret ObservabilityPipelineGooglePubSubDestinationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGooglePubSubDestination) GetTypeOk() (*ObservabilityPipelineGooglePubSubDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineGooglePubSubDestination) SetType(v ObservabilityPipelineGooglePubSubDestinationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineGooglePubSubDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Auth != nil {
		toSerialize["auth"] = o.Auth
	}
	toSerialize["encoding"] = o.Encoding
	toSerialize["id"] = o.Id
	toSerialize["inputs"] = o.Inputs
	toSerialize["project"] = o.Project
	if o.Tls != nil {
		toSerialize["tls"] = o.Tls
	}
	toSerialize["topic"] = o.Topic
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineGooglePubSubDestination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Auth     *ObservabilityPipelineGcpAuth                         `json:"auth,omitempty"`
		Encoding *ObservabilityPipelineGooglePubSubDestinationEncoding `json:"encoding"`
		Id       *string                                               `json:"id"`
		Inputs   *[]string                                             `json:"inputs"`
		Project  *string                                               `json:"project"`
		Tls      *ObservabilityPipelineTls                             `json:"tls,omitempty"`
		Topic    *string                                               `json:"topic"`
		Type     *ObservabilityPipelineGooglePubSubDestinationType     `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
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
	if all.Project == nil {
		return fmt.Errorf("required field project missing")
	}
	if all.Topic == nil {
		return fmt.Errorf("required field topic missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auth", "encoding", "id", "inputs", "project", "tls", "topic", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Auth != nil && all.Auth.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Auth = all.Auth
	if !all.Encoding.IsValid() {
		hasInvalidField = true
	} else {
		o.Encoding = *all.Encoding
	}
	o.Id = *all.Id
	o.Inputs = *all.Inputs
	o.Project = *all.Project
	if all.Tls != nil && all.Tls.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Tls = all.Tls
	o.Topic = *all.Topic
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
