// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineKafkaSource The `kafka` source ingests data from Apache Kafka topics.
//
// **Supported pipeline types:** logs
type ObservabilityPipelineKafkaSource struct {
	// Name of the environment variable or secret that holds the Kafka bootstrap servers list.
	BootstrapServersKey *string `json:"bootstrap_servers_key,omitempty"`
	// Consumer group ID used by the Kafka client.
	GroupId string `json:"group_id"`
	// The unique identifier for this component. Used in other parts of the pipeline to reference this component (for example, as the `input` to downstream components).
	Id string `json:"id"`
	// Optional list of advanced Kafka client configuration options, defined as key-value pairs.
	LibrdkafkaOptions []ObservabilityPipelineKafkaLibrdkafkaOption `json:"librdkafka_options,omitempty"`
	// Specifies the SASL mechanism for authenticating with a Kafka cluster.
	Sasl *ObservabilityPipelineKafkaSasl `json:"sasl,omitempty"`
	// Configuration for enabling TLS encryption between the pipeline component and external services.
	Tls *ObservabilityPipelineTls `json:"tls,omitempty"`
	// A list of Kafka topic names to subscribe to. The source ingests messages from each topic specified.
	Topics []string `json:"topics"`
	// The source type. The value should always be `kafka`.
	Type ObservabilityPipelineKafkaSourceType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineKafkaSource instantiates a new ObservabilityPipelineKafkaSource object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineKafkaSource(groupId string, id string, topics []string, typeVar ObservabilityPipelineKafkaSourceType) *ObservabilityPipelineKafkaSource {
	this := ObservabilityPipelineKafkaSource{}
	this.GroupId = groupId
	this.Id = id
	this.Topics = topics
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineKafkaSourceWithDefaults instantiates a new ObservabilityPipelineKafkaSource object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineKafkaSourceWithDefaults() *ObservabilityPipelineKafkaSource {
	this := ObservabilityPipelineKafkaSource{}
	var typeVar ObservabilityPipelineKafkaSourceType = OBSERVABILITYPIPELINEKAFKASOURCETYPE_KAFKA
	this.Type = typeVar
	return &this
}

// GetBootstrapServersKey returns the BootstrapServersKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineKafkaSource) GetBootstrapServersKey() string {
	if o == nil || o.BootstrapServersKey == nil {
		var ret string
		return ret
	}
	return *o.BootstrapServersKey
}

// GetBootstrapServersKeyOk returns a tuple with the BootstrapServersKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineKafkaSource) GetBootstrapServersKeyOk() (*string, bool) {
	if o == nil || o.BootstrapServersKey == nil {
		return nil, false
	}
	return o.BootstrapServersKey, true
}

// HasBootstrapServersKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineKafkaSource) HasBootstrapServersKey() bool {
	return o != nil && o.BootstrapServersKey != nil
}

// SetBootstrapServersKey gets a reference to the given string and assigns it to the BootstrapServersKey field.
func (o *ObservabilityPipelineKafkaSource) SetBootstrapServersKey(v string) {
	o.BootstrapServersKey = &v
}

// GetGroupId returns the GroupId field value.
func (o *ObservabilityPipelineKafkaSource) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineKafkaSource) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value.
func (o *ObservabilityPipelineKafkaSource) SetGroupId(v string) {
	o.GroupId = v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineKafkaSource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineKafkaSource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineKafkaSource) SetId(v string) {
	o.Id = v
}

// GetLibrdkafkaOptions returns the LibrdkafkaOptions field value if set, zero value otherwise.
func (o *ObservabilityPipelineKafkaSource) GetLibrdkafkaOptions() []ObservabilityPipelineKafkaLibrdkafkaOption {
	if o == nil || o.LibrdkafkaOptions == nil {
		var ret []ObservabilityPipelineKafkaLibrdkafkaOption
		return ret
	}
	return o.LibrdkafkaOptions
}

// GetLibrdkafkaOptionsOk returns a tuple with the LibrdkafkaOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineKafkaSource) GetLibrdkafkaOptionsOk() (*[]ObservabilityPipelineKafkaLibrdkafkaOption, bool) {
	if o == nil || o.LibrdkafkaOptions == nil {
		return nil, false
	}
	return &o.LibrdkafkaOptions, true
}

// HasLibrdkafkaOptions returns a boolean if a field has been set.
func (o *ObservabilityPipelineKafkaSource) HasLibrdkafkaOptions() bool {
	return o != nil && o.LibrdkafkaOptions != nil
}

// SetLibrdkafkaOptions gets a reference to the given []ObservabilityPipelineKafkaLibrdkafkaOption and assigns it to the LibrdkafkaOptions field.
func (o *ObservabilityPipelineKafkaSource) SetLibrdkafkaOptions(v []ObservabilityPipelineKafkaLibrdkafkaOption) {
	o.LibrdkafkaOptions = v
}

// GetSasl returns the Sasl field value if set, zero value otherwise.
func (o *ObservabilityPipelineKafkaSource) GetSasl() ObservabilityPipelineKafkaSasl {
	if o == nil || o.Sasl == nil {
		var ret ObservabilityPipelineKafkaSasl
		return ret
	}
	return *o.Sasl
}

// GetSaslOk returns a tuple with the Sasl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineKafkaSource) GetSaslOk() (*ObservabilityPipelineKafkaSasl, bool) {
	if o == nil || o.Sasl == nil {
		return nil, false
	}
	return o.Sasl, true
}

// HasSasl returns a boolean if a field has been set.
func (o *ObservabilityPipelineKafkaSource) HasSasl() bool {
	return o != nil && o.Sasl != nil
}

// SetSasl gets a reference to the given ObservabilityPipelineKafkaSasl and assigns it to the Sasl field.
func (o *ObservabilityPipelineKafkaSource) SetSasl(v ObservabilityPipelineKafkaSasl) {
	o.Sasl = &v
}

// GetTls returns the Tls field value if set, zero value otherwise.
func (o *ObservabilityPipelineKafkaSource) GetTls() ObservabilityPipelineTls {
	if o == nil || o.Tls == nil {
		var ret ObservabilityPipelineTls
		return ret
	}
	return *o.Tls
}

// GetTlsOk returns a tuple with the Tls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineKafkaSource) GetTlsOk() (*ObservabilityPipelineTls, bool) {
	if o == nil || o.Tls == nil {
		return nil, false
	}
	return o.Tls, true
}

// HasTls returns a boolean if a field has been set.
func (o *ObservabilityPipelineKafkaSource) HasTls() bool {
	return o != nil && o.Tls != nil
}

// SetTls gets a reference to the given ObservabilityPipelineTls and assigns it to the Tls field.
func (o *ObservabilityPipelineKafkaSource) SetTls(v ObservabilityPipelineTls) {
	o.Tls = &v
}

// GetTopics returns the Topics field value.
func (o *ObservabilityPipelineKafkaSource) GetTopics() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Topics
}

// GetTopicsOk returns a tuple with the Topics field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineKafkaSource) GetTopicsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Topics, true
}

// SetTopics sets field value.
func (o *ObservabilityPipelineKafkaSource) SetTopics(v []string) {
	o.Topics = v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineKafkaSource) GetType() ObservabilityPipelineKafkaSourceType {
	if o == nil {
		var ret ObservabilityPipelineKafkaSourceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineKafkaSource) GetTypeOk() (*ObservabilityPipelineKafkaSourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineKafkaSource) SetType(v ObservabilityPipelineKafkaSourceType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineKafkaSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.BootstrapServersKey != nil {
		toSerialize["bootstrap_servers_key"] = o.BootstrapServersKey
	}
	toSerialize["group_id"] = o.GroupId
	toSerialize["id"] = o.Id
	if o.LibrdkafkaOptions != nil {
		toSerialize["librdkafka_options"] = o.LibrdkafkaOptions
	}
	if o.Sasl != nil {
		toSerialize["sasl"] = o.Sasl
	}
	if o.Tls != nil {
		toSerialize["tls"] = o.Tls
	}
	toSerialize["topics"] = o.Topics
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineKafkaSource) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BootstrapServersKey *string                                      `json:"bootstrap_servers_key,omitempty"`
		GroupId             *string                                      `json:"group_id"`
		Id                  *string                                      `json:"id"`
		LibrdkafkaOptions   []ObservabilityPipelineKafkaLibrdkafkaOption `json:"librdkafka_options,omitempty"`
		Sasl                *ObservabilityPipelineKafkaSasl              `json:"sasl,omitempty"`
		Tls                 *ObservabilityPipelineTls                    `json:"tls,omitempty"`
		Topics              *[]string                                    `json:"topics"`
		Type                *ObservabilityPipelineKafkaSourceType        `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.GroupId == nil {
		return fmt.Errorf("required field group_id missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Topics == nil {
		return fmt.Errorf("required field topics missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"bootstrap_servers_key", "group_id", "id", "librdkafka_options", "sasl", "tls", "topics", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.BootstrapServersKey = all.BootstrapServersKey
	o.GroupId = *all.GroupId
	o.Id = *all.Id
	o.LibrdkafkaOptions = all.LibrdkafkaOptions
	if all.Sasl != nil && all.Sasl.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Sasl = all.Sasl
	if all.Tls != nil && all.Tls.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Tls = all.Tls
	o.Topics = *all.Topics
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
