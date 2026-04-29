// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineKafkaDestination The `kafka` destination sends logs to Apache Kafka topics.
//
// **Supported pipeline types:** logs
type ObservabilityPipelineKafkaDestination struct {
	// Name of the environment variable or secret that holds the Kafka bootstrap servers list.
	BootstrapServersKey *string `json:"bootstrap_servers_key,omitempty"`
	// Compression codec for Kafka messages.
	Compression *ObservabilityPipelineKafkaDestinationCompression `json:"compression,omitempty"`
	// Encoding format for log events.
	Encoding ObservabilityPipelineKafkaDestinationEncoding `json:"encoding"`
	// The field name to use for Kafka message headers.
	HeadersKey *string `json:"headers_key,omitempty"`
	// The unique identifier for this component.
	Id string `json:"id"`
	// A list of component IDs whose output is used as the `input` for this component.
	Inputs []string `json:"inputs"`
	// The field name to use as the Kafka message key.
	KeyField *string `json:"key_field,omitempty"`
	// Optional list of advanced Kafka producer configuration options, defined as key-value pairs.
	LibrdkafkaOptions []ObservabilityPipelineKafkaLibrdkafkaOption `json:"librdkafka_options,omitempty"`
	// Maximum time in milliseconds to wait for message delivery confirmation.
	MessageTimeoutMs *int64 `json:"message_timeout_ms,omitempty"`
	// Duration in seconds for the rate limit window.
	RateLimitDurationSecs *int64 `json:"rate_limit_duration_secs,omitempty"`
	// Maximum number of messages allowed per rate limit duration.
	RateLimitNum *int64 `json:"rate_limit_num,omitempty"`
	// Specifies the SASL mechanism for authenticating with a Kafka cluster.
	Sasl *ObservabilityPipelineKafkaSasl `json:"sasl,omitempty"`
	// Socket timeout in milliseconds for network requests.
	SocketTimeoutMs *int64 `json:"socket_timeout_ms,omitempty"`
	// Configuration for enabling TLS encryption between the pipeline component and external services.
	Tls *ObservabilityPipelineTls `json:"tls,omitempty"`
	// The Kafka topic name to publish logs to.
	Topic string `json:"topic"`
	// The destination type. The value should always be `kafka`.
	Type ObservabilityPipelineKafkaDestinationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineKafkaDestination instantiates a new ObservabilityPipelineKafkaDestination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineKafkaDestination(encoding ObservabilityPipelineKafkaDestinationEncoding, id string, inputs []string, topic string, typeVar ObservabilityPipelineKafkaDestinationType) *ObservabilityPipelineKafkaDestination {
	this := ObservabilityPipelineKafkaDestination{}
	this.Encoding = encoding
	this.Id = id
	this.Inputs = inputs
	this.Topic = topic
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineKafkaDestinationWithDefaults instantiates a new ObservabilityPipelineKafkaDestination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineKafkaDestinationWithDefaults() *ObservabilityPipelineKafkaDestination {
	this := ObservabilityPipelineKafkaDestination{}
	var typeVar ObservabilityPipelineKafkaDestinationType = OBSERVABILITYPIPELINEKAFKADESTINATIONTYPE_KAFKA
	this.Type = typeVar
	return &this
}

// GetBootstrapServersKey returns the BootstrapServersKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineKafkaDestination) GetBootstrapServersKey() string {
	if o == nil || o.BootstrapServersKey == nil {
		var ret string
		return ret
	}
	return *o.BootstrapServersKey
}

// GetBootstrapServersKeyOk returns a tuple with the BootstrapServersKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineKafkaDestination) GetBootstrapServersKeyOk() (*string, bool) {
	if o == nil || o.BootstrapServersKey == nil {
		return nil, false
	}
	return o.BootstrapServersKey, true
}

// HasBootstrapServersKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineKafkaDestination) HasBootstrapServersKey() bool {
	return o != nil && o.BootstrapServersKey != nil
}

// SetBootstrapServersKey gets a reference to the given string and assigns it to the BootstrapServersKey field.
func (o *ObservabilityPipelineKafkaDestination) SetBootstrapServersKey(v string) {
	o.BootstrapServersKey = &v
}

// GetCompression returns the Compression field value if set, zero value otherwise.
func (o *ObservabilityPipelineKafkaDestination) GetCompression() ObservabilityPipelineKafkaDestinationCompression {
	if o == nil || o.Compression == nil {
		var ret ObservabilityPipelineKafkaDestinationCompression
		return ret
	}
	return *o.Compression
}

// GetCompressionOk returns a tuple with the Compression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineKafkaDestination) GetCompressionOk() (*ObservabilityPipelineKafkaDestinationCompression, bool) {
	if o == nil || o.Compression == nil {
		return nil, false
	}
	return o.Compression, true
}

// HasCompression returns a boolean if a field has been set.
func (o *ObservabilityPipelineKafkaDestination) HasCompression() bool {
	return o != nil && o.Compression != nil
}

// SetCompression gets a reference to the given ObservabilityPipelineKafkaDestinationCompression and assigns it to the Compression field.
func (o *ObservabilityPipelineKafkaDestination) SetCompression(v ObservabilityPipelineKafkaDestinationCompression) {
	o.Compression = &v
}

// GetEncoding returns the Encoding field value.
func (o *ObservabilityPipelineKafkaDestination) GetEncoding() ObservabilityPipelineKafkaDestinationEncoding {
	if o == nil {
		var ret ObservabilityPipelineKafkaDestinationEncoding
		return ret
	}
	return o.Encoding
}

// GetEncodingOk returns a tuple with the Encoding field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineKafkaDestination) GetEncodingOk() (*ObservabilityPipelineKafkaDestinationEncoding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Encoding, true
}

// SetEncoding sets field value.
func (o *ObservabilityPipelineKafkaDestination) SetEncoding(v ObservabilityPipelineKafkaDestinationEncoding) {
	o.Encoding = v
}

// GetHeadersKey returns the HeadersKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineKafkaDestination) GetHeadersKey() string {
	if o == nil || o.HeadersKey == nil {
		var ret string
		return ret
	}
	return *o.HeadersKey
}

// GetHeadersKeyOk returns a tuple with the HeadersKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineKafkaDestination) GetHeadersKeyOk() (*string, bool) {
	if o == nil || o.HeadersKey == nil {
		return nil, false
	}
	return o.HeadersKey, true
}

// HasHeadersKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineKafkaDestination) HasHeadersKey() bool {
	return o != nil && o.HeadersKey != nil
}

// SetHeadersKey gets a reference to the given string and assigns it to the HeadersKey field.
func (o *ObservabilityPipelineKafkaDestination) SetHeadersKey(v string) {
	o.HeadersKey = &v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineKafkaDestination) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineKafkaDestination) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineKafkaDestination) SetId(v string) {
	o.Id = v
}

// GetInputs returns the Inputs field value.
func (o *ObservabilityPipelineKafkaDestination) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineKafkaDestination) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *ObservabilityPipelineKafkaDestination) SetInputs(v []string) {
	o.Inputs = v
}

// GetKeyField returns the KeyField field value if set, zero value otherwise.
func (o *ObservabilityPipelineKafkaDestination) GetKeyField() string {
	if o == nil || o.KeyField == nil {
		var ret string
		return ret
	}
	return *o.KeyField
}

// GetKeyFieldOk returns a tuple with the KeyField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineKafkaDestination) GetKeyFieldOk() (*string, bool) {
	if o == nil || o.KeyField == nil {
		return nil, false
	}
	return o.KeyField, true
}

// HasKeyField returns a boolean if a field has been set.
func (o *ObservabilityPipelineKafkaDestination) HasKeyField() bool {
	return o != nil && o.KeyField != nil
}

// SetKeyField gets a reference to the given string and assigns it to the KeyField field.
func (o *ObservabilityPipelineKafkaDestination) SetKeyField(v string) {
	o.KeyField = &v
}

// GetLibrdkafkaOptions returns the LibrdkafkaOptions field value if set, zero value otherwise.
func (o *ObservabilityPipelineKafkaDestination) GetLibrdkafkaOptions() []ObservabilityPipelineKafkaLibrdkafkaOption {
	if o == nil || o.LibrdkafkaOptions == nil {
		var ret []ObservabilityPipelineKafkaLibrdkafkaOption
		return ret
	}
	return o.LibrdkafkaOptions
}

// GetLibrdkafkaOptionsOk returns a tuple with the LibrdkafkaOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineKafkaDestination) GetLibrdkafkaOptionsOk() (*[]ObservabilityPipelineKafkaLibrdkafkaOption, bool) {
	if o == nil || o.LibrdkafkaOptions == nil {
		return nil, false
	}
	return &o.LibrdkafkaOptions, true
}

// HasLibrdkafkaOptions returns a boolean if a field has been set.
func (o *ObservabilityPipelineKafkaDestination) HasLibrdkafkaOptions() bool {
	return o != nil && o.LibrdkafkaOptions != nil
}

// SetLibrdkafkaOptions gets a reference to the given []ObservabilityPipelineKafkaLibrdkafkaOption and assigns it to the LibrdkafkaOptions field.
func (o *ObservabilityPipelineKafkaDestination) SetLibrdkafkaOptions(v []ObservabilityPipelineKafkaLibrdkafkaOption) {
	o.LibrdkafkaOptions = v
}

// GetMessageTimeoutMs returns the MessageTimeoutMs field value if set, zero value otherwise.
func (o *ObservabilityPipelineKafkaDestination) GetMessageTimeoutMs() int64 {
	if o == nil || o.MessageTimeoutMs == nil {
		var ret int64
		return ret
	}
	return *o.MessageTimeoutMs
}

// GetMessageTimeoutMsOk returns a tuple with the MessageTimeoutMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineKafkaDestination) GetMessageTimeoutMsOk() (*int64, bool) {
	if o == nil || o.MessageTimeoutMs == nil {
		return nil, false
	}
	return o.MessageTimeoutMs, true
}

// HasMessageTimeoutMs returns a boolean if a field has been set.
func (o *ObservabilityPipelineKafkaDestination) HasMessageTimeoutMs() bool {
	return o != nil && o.MessageTimeoutMs != nil
}

// SetMessageTimeoutMs gets a reference to the given int64 and assigns it to the MessageTimeoutMs field.
func (o *ObservabilityPipelineKafkaDestination) SetMessageTimeoutMs(v int64) {
	o.MessageTimeoutMs = &v
}

// GetRateLimitDurationSecs returns the RateLimitDurationSecs field value if set, zero value otherwise.
func (o *ObservabilityPipelineKafkaDestination) GetRateLimitDurationSecs() int64 {
	if o == nil || o.RateLimitDurationSecs == nil {
		var ret int64
		return ret
	}
	return *o.RateLimitDurationSecs
}

// GetRateLimitDurationSecsOk returns a tuple with the RateLimitDurationSecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineKafkaDestination) GetRateLimitDurationSecsOk() (*int64, bool) {
	if o == nil || o.RateLimitDurationSecs == nil {
		return nil, false
	}
	return o.RateLimitDurationSecs, true
}

// HasRateLimitDurationSecs returns a boolean if a field has been set.
func (o *ObservabilityPipelineKafkaDestination) HasRateLimitDurationSecs() bool {
	return o != nil && o.RateLimitDurationSecs != nil
}

// SetRateLimitDurationSecs gets a reference to the given int64 and assigns it to the RateLimitDurationSecs field.
func (o *ObservabilityPipelineKafkaDestination) SetRateLimitDurationSecs(v int64) {
	o.RateLimitDurationSecs = &v
}

// GetRateLimitNum returns the RateLimitNum field value if set, zero value otherwise.
func (o *ObservabilityPipelineKafkaDestination) GetRateLimitNum() int64 {
	if o == nil || o.RateLimitNum == nil {
		var ret int64
		return ret
	}
	return *o.RateLimitNum
}

// GetRateLimitNumOk returns a tuple with the RateLimitNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineKafkaDestination) GetRateLimitNumOk() (*int64, bool) {
	if o == nil || o.RateLimitNum == nil {
		return nil, false
	}
	return o.RateLimitNum, true
}

// HasRateLimitNum returns a boolean if a field has been set.
func (o *ObservabilityPipelineKafkaDestination) HasRateLimitNum() bool {
	return o != nil && o.RateLimitNum != nil
}

// SetRateLimitNum gets a reference to the given int64 and assigns it to the RateLimitNum field.
func (o *ObservabilityPipelineKafkaDestination) SetRateLimitNum(v int64) {
	o.RateLimitNum = &v
}

// GetSasl returns the Sasl field value if set, zero value otherwise.
func (o *ObservabilityPipelineKafkaDestination) GetSasl() ObservabilityPipelineKafkaSasl {
	if o == nil || o.Sasl == nil {
		var ret ObservabilityPipelineKafkaSasl
		return ret
	}
	return *o.Sasl
}

// GetSaslOk returns a tuple with the Sasl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineKafkaDestination) GetSaslOk() (*ObservabilityPipelineKafkaSasl, bool) {
	if o == nil || o.Sasl == nil {
		return nil, false
	}
	return o.Sasl, true
}

// HasSasl returns a boolean if a field has been set.
func (o *ObservabilityPipelineKafkaDestination) HasSasl() bool {
	return o != nil && o.Sasl != nil
}

// SetSasl gets a reference to the given ObservabilityPipelineKafkaSasl and assigns it to the Sasl field.
func (o *ObservabilityPipelineKafkaDestination) SetSasl(v ObservabilityPipelineKafkaSasl) {
	o.Sasl = &v
}

// GetSocketTimeoutMs returns the SocketTimeoutMs field value if set, zero value otherwise.
func (o *ObservabilityPipelineKafkaDestination) GetSocketTimeoutMs() int64 {
	if o == nil || o.SocketTimeoutMs == nil {
		var ret int64
		return ret
	}
	return *o.SocketTimeoutMs
}

// GetSocketTimeoutMsOk returns a tuple with the SocketTimeoutMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineKafkaDestination) GetSocketTimeoutMsOk() (*int64, bool) {
	if o == nil || o.SocketTimeoutMs == nil {
		return nil, false
	}
	return o.SocketTimeoutMs, true
}

// HasSocketTimeoutMs returns a boolean if a field has been set.
func (o *ObservabilityPipelineKafkaDestination) HasSocketTimeoutMs() bool {
	return o != nil && o.SocketTimeoutMs != nil
}

// SetSocketTimeoutMs gets a reference to the given int64 and assigns it to the SocketTimeoutMs field.
func (o *ObservabilityPipelineKafkaDestination) SetSocketTimeoutMs(v int64) {
	o.SocketTimeoutMs = &v
}

// GetTls returns the Tls field value if set, zero value otherwise.
func (o *ObservabilityPipelineKafkaDestination) GetTls() ObservabilityPipelineTls {
	if o == nil || o.Tls == nil {
		var ret ObservabilityPipelineTls
		return ret
	}
	return *o.Tls
}

// GetTlsOk returns a tuple with the Tls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineKafkaDestination) GetTlsOk() (*ObservabilityPipelineTls, bool) {
	if o == nil || o.Tls == nil {
		return nil, false
	}
	return o.Tls, true
}

// HasTls returns a boolean if a field has been set.
func (o *ObservabilityPipelineKafkaDestination) HasTls() bool {
	return o != nil && o.Tls != nil
}

// SetTls gets a reference to the given ObservabilityPipelineTls and assigns it to the Tls field.
func (o *ObservabilityPipelineKafkaDestination) SetTls(v ObservabilityPipelineTls) {
	o.Tls = &v
}

// GetTopic returns the Topic field value.
func (o *ObservabilityPipelineKafkaDestination) GetTopic() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Topic
}

// GetTopicOk returns a tuple with the Topic field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineKafkaDestination) GetTopicOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Topic, true
}

// SetTopic sets field value.
func (o *ObservabilityPipelineKafkaDestination) SetTopic(v string) {
	o.Topic = v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineKafkaDestination) GetType() ObservabilityPipelineKafkaDestinationType {
	if o == nil {
		var ret ObservabilityPipelineKafkaDestinationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineKafkaDestination) GetTypeOk() (*ObservabilityPipelineKafkaDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineKafkaDestination) SetType(v ObservabilityPipelineKafkaDestinationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineKafkaDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.BootstrapServersKey != nil {
		toSerialize["bootstrap_servers_key"] = o.BootstrapServersKey
	}
	if o.Compression != nil {
		toSerialize["compression"] = o.Compression
	}
	toSerialize["encoding"] = o.Encoding
	if o.HeadersKey != nil {
		toSerialize["headers_key"] = o.HeadersKey
	}
	toSerialize["id"] = o.Id
	toSerialize["inputs"] = o.Inputs
	if o.KeyField != nil {
		toSerialize["key_field"] = o.KeyField
	}
	if o.LibrdkafkaOptions != nil {
		toSerialize["librdkafka_options"] = o.LibrdkafkaOptions
	}
	if o.MessageTimeoutMs != nil {
		toSerialize["message_timeout_ms"] = o.MessageTimeoutMs
	}
	if o.RateLimitDurationSecs != nil {
		toSerialize["rate_limit_duration_secs"] = o.RateLimitDurationSecs
	}
	if o.RateLimitNum != nil {
		toSerialize["rate_limit_num"] = o.RateLimitNum
	}
	if o.Sasl != nil {
		toSerialize["sasl"] = o.Sasl
	}
	if o.SocketTimeoutMs != nil {
		toSerialize["socket_timeout_ms"] = o.SocketTimeoutMs
	}
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
func (o *ObservabilityPipelineKafkaDestination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BootstrapServersKey   *string                                           `json:"bootstrap_servers_key,omitempty"`
		Compression           *ObservabilityPipelineKafkaDestinationCompression `json:"compression,omitempty"`
		Encoding              *ObservabilityPipelineKafkaDestinationEncoding    `json:"encoding"`
		HeadersKey            *string                                           `json:"headers_key,omitempty"`
		Id                    *string                                           `json:"id"`
		Inputs                *[]string                                         `json:"inputs"`
		KeyField              *string                                           `json:"key_field,omitempty"`
		LibrdkafkaOptions     []ObservabilityPipelineKafkaLibrdkafkaOption      `json:"librdkafka_options,omitempty"`
		MessageTimeoutMs      *int64                                            `json:"message_timeout_ms,omitempty"`
		RateLimitDurationSecs *int64                                            `json:"rate_limit_duration_secs,omitempty"`
		RateLimitNum          *int64                                            `json:"rate_limit_num,omitempty"`
		Sasl                  *ObservabilityPipelineKafkaSasl                   `json:"sasl,omitempty"`
		SocketTimeoutMs       *int64                                            `json:"socket_timeout_ms,omitempty"`
		Tls                   *ObservabilityPipelineTls                         `json:"tls,omitempty"`
		Topic                 *string                                           `json:"topic"`
		Type                  *ObservabilityPipelineKafkaDestinationType        `json:"type"`
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
	if all.Topic == nil {
		return fmt.Errorf("required field topic missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"bootstrap_servers_key", "compression", "encoding", "headers_key", "id", "inputs", "key_field", "librdkafka_options", "message_timeout_ms", "rate_limit_duration_secs", "rate_limit_num", "sasl", "socket_timeout_ms", "tls", "topic", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.BootstrapServersKey = all.BootstrapServersKey
	if all.Compression != nil && !all.Compression.IsValid() {
		hasInvalidField = true
	} else {
		o.Compression = all.Compression
	}
	if !all.Encoding.IsValid() {
		hasInvalidField = true
	} else {
		o.Encoding = *all.Encoding
	}
	o.HeadersKey = all.HeadersKey
	o.Id = *all.Id
	o.Inputs = *all.Inputs
	o.KeyField = all.KeyField
	o.LibrdkafkaOptions = all.LibrdkafkaOptions
	o.MessageTimeoutMs = all.MessageTimeoutMs
	o.RateLimitDurationSecs = all.RateLimitDurationSecs
	o.RateLimitNum = all.RateLimitNum
	if all.Sasl != nil && all.Sasl.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Sasl = all.Sasl
	o.SocketTimeoutMs = all.SocketTimeoutMs
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
