// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSplunkHecMetricsDestination The `splunk_hec_metrics` destination forwards metrics to Splunk using the HTTP Event Collector (HEC).
//
// **Supported pipeline types:** metrics
type ObservabilityPipelineSplunkHecMetricsDestination struct {
	// Configuration for buffer settings on destination components.
	Buffer *ObservabilityPipelineBufferOptions `json:"buffer,omitempty"`
	// Compression algorithm applied when sending metrics to Splunk HEC.
	Compression *ObservabilityPipelineSplunkHecMetricsDestinationCompression `json:"compression,omitempty"`
	// Optional default namespace for metrics sent to Splunk HEC.
	DefaultNamespace *string `json:"default_namespace,omitempty"`
	// Name of the environment variable or secret that holds the Splunk HEC endpoint URL.
	EndpointUrlKey *string `json:"endpoint_url_key,omitempty"`
	// The unique identifier for this component. Used in other parts of the pipeline to reference this component (for example, as the `input` to downstream components).
	Id string `json:"id"`
	// Optional name of the Splunk index where metrics are written.
	Index *string `json:"index,omitempty"`
	// A list of component IDs whose output is used as the `input` for this component.
	Inputs []string `json:"inputs"`
	// The Splunk source field value for metric events.
	Source *string `json:"source,omitempty"`
	// The Splunk sourcetype to assign to metric events.
	Sourcetype *string `json:"sourcetype,omitempty"`
	// Configuration for enabling TLS encryption between the pipeline component and external services.
	Tls *ObservabilityPipelineTls `json:"tls,omitempty"`
	// Name of the environment variable or secret that holds the Splunk HEC token.
	TokenKey *string `json:"token_key,omitempty"`
	// The destination type. Always `splunk_hec_metrics`.
	Type ObservabilityPipelineSplunkHecMetricsDestinationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineSplunkHecMetricsDestination instantiates a new ObservabilityPipelineSplunkHecMetricsDestination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineSplunkHecMetricsDestination(id string, inputs []string, typeVar ObservabilityPipelineSplunkHecMetricsDestinationType) *ObservabilityPipelineSplunkHecMetricsDestination {
	this := ObservabilityPipelineSplunkHecMetricsDestination{}
	var compression ObservabilityPipelineSplunkHecMetricsDestinationCompression = OBSERVABILITYPIPELINESPLUNKHECMETRICSDESTINATIONCOMPRESSION_NONE
	this.Compression = &compression
	this.Id = id
	this.Inputs = inputs
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineSplunkHecMetricsDestinationWithDefaults instantiates a new ObservabilityPipelineSplunkHecMetricsDestination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineSplunkHecMetricsDestinationWithDefaults() *ObservabilityPipelineSplunkHecMetricsDestination {
	this := ObservabilityPipelineSplunkHecMetricsDestination{}
	var compression ObservabilityPipelineSplunkHecMetricsDestinationCompression = OBSERVABILITYPIPELINESPLUNKHECMETRICSDESTINATIONCOMPRESSION_NONE
	this.Compression = &compression
	var typeVar ObservabilityPipelineSplunkHecMetricsDestinationType = OBSERVABILITYPIPELINESPLUNKHECMETRICSDESTINATIONTYPE_SPLUNK_HEC_METRICS
	this.Type = typeVar
	return &this
}

// GetBuffer returns the Buffer field value if set, zero value otherwise.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) GetBuffer() ObservabilityPipelineBufferOptions {
	if o == nil || o.Buffer == nil {
		var ret ObservabilityPipelineBufferOptions
		return ret
	}
	return *o.Buffer
}

// GetBufferOk returns a tuple with the Buffer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) GetBufferOk() (*ObservabilityPipelineBufferOptions, bool) {
	if o == nil || o.Buffer == nil {
		return nil, false
	}
	return o.Buffer, true
}

// HasBuffer returns a boolean if a field has been set.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) HasBuffer() bool {
	return o != nil && o.Buffer != nil
}

// SetBuffer gets a reference to the given ObservabilityPipelineBufferOptions and assigns it to the Buffer field.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) SetBuffer(v ObservabilityPipelineBufferOptions) {
	o.Buffer = &v
}

// GetCompression returns the Compression field value if set, zero value otherwise.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) GetCompression() ObservabilityPipelineSplunkHecMetricsDestinationCompression {
	if o == nil || o.Compression == nil {
		var ret ObservabilityPipelineSplunkHecMetricsDestinationCompression
		return ret
	}
	return *o.Compression
}

// GetCompressionOk returns a tuple with the Compression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) GetCompressionOk() (*ObservabilityPipelineSplunkHecMetricsDestinationCompression, bool) {
	if o == nil || o.Compression == nil {
		return nil, false
	}
	return o.Compression, true
}

// HasCompression returns a boolean if a field has been set.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) HasCompression() bool {
	return o != nil && o.Compression != nil
}

// SetCompression gets a reference to the given ObservabilityPipelineSplunkHecMetricsDestinationCompression and assigns it to the Compression field.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) SetCompression(v ObservabilityPipelineSplunkHecMetricsDestinationCompression) {
	o.Compression = &v
}

// GetDefaultNamespace returns the DefaultNamespace field value if set, zero value otherwise.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) GetDefaultNamespace() string {
	if o == nil || o.DefaultNamespace == nil {
		var ret string
		return ret
	}
	return *o.DefaultNamespace
}

// GetDefaultNamespaceOk returns a tuple with the DefaultNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) GetDefaultNamespaceOk() (*string, bool) {
	if o == nil || o.DefaultNamespace == nil {
		return nil, false
	}
	return o.DefaultNamespace, true
}

// HasDefaultNamespace returns a boolean if a field has been set.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) HasDefaultNamespace() bool {
	return o != nil && o.DefaultNamespace != nil
}

// SetDefaultNamespace gets a reference to the given string and assigns it to the DefaultNamespace field.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) SetDefaultNamespace(v string) {
	o.DefaultNamespace = &v
}

// GetEndpointUrlKey returns the EndpointUrlKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) GetEndpointUrlKey() string {
	if o == nil || o.EndpointUrlKey == nil {
		var ret string
		return ret
	}
	return *o.EndpointUrlKey
}

// GetEndpointUrlKeyOk returns a tuple with the EndpointUrlKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) GetEndpointUrlKeyOk() (*string, bool) {
	if o == nil || o.EndpointUrlKey == nil {
		return nil, false
	}
	return o.EndpointUrlKey, true
}

// HasEndpointUrlKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) HasEndpointUrlKey() bool {
	return o != nil && o.EndpointUrlKey != nil
}

// SetEndpointUrlKey gets a reference to the given string and assigns it to the EndpointUrlKey field.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) SetEndpointUrlKey(v string) {
	o.EndpointUrlKey = &v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) SetId(v string) {
	o.Id = v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) GetIndex() string {
	if o == nil || o.Index == nil {
		var ret string
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) GetIndexOk() (*string, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) HasIndex() bool {
	return o != nil && o.Index != nil
}

// SetIndex gets a reference to the given string and assigns it to the Index field.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) SetIndex(v string) {
	o.Index = &v
}

// GetInputs returns the Inputs field value.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) SetInputs(v []string) {
	o.Inputs = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) SetSource(v string) {
	o.Source = &v
}

// GetSourcetype returns the Sourcetype field value if set, zero value otherwise.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) GetSourcetype() string {
	if o == nil || o.Sourcetype == nil {
		var ret string
		return ret
	}
	return *o.Sourcetype
}

// GetSourcetypeOk returns a tuple with the Sourcetype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) GetSourcetypeOk() (*string, bool) {
	if o == nil || o.Sourcetype == nil {
		return nil, false
	}
	return o.Sourcetype, true
}

// HasSourcetype returns a boolean if a field has been set.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) HasSourcetype() bool {
	return o != nil && o.Sourcetype != nil
}

// SetSourcetype gets a reference to the given string and assigns it to the Sourcetype field.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) SetSourcetype(v string) {
	o.Sourcetype = &v
}

// GetTls returns the Tls field value if set, zero value otherwise.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) GetTls() ObservabilityPipelineTls {
	if o == nil || o.Tls == nil {
		var ret ObservabilityPipelineTls
		return ret
	}
	return *o.Tls
}

// GetTlsOk returns a tuple with the Tls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) GetTlsOk() (*ObservabilityPipelineTls, bool) {
	if o == nil || o.Tls == nil {
		return nil, false
	}
	return o.Tls, true
}

// HasTls returns a boolean if a field has been set.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) HasTls() bool {
	return o != nil && o.Tls != nil
}

// SetTls gets a reference to the given ObservabilityPipelineTls and assigns it to the Tls field.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) SetTls(v ObservabilityPipelineTls) {
	o.Tls = &v
}

// GetTokenKey returns the TokenKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) GetTokenKey() string {
	if o == nil || o.TokenKey == nil {
		var ret string
		return ret
	}
	return *o.TokenKey
}

// GetTokenKeyOk returns a tuple with the TokenKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) GetTokenKeyOk() (*string, bool) {
	if o == nil || o.TokenKey == nil {
		return nil, false
	}
	return o.TokenKey, true
}

// HasTokenKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) HasTokenKey() bool {
	return o != nil && o.TokenKey != nil
}

// SetTokenKey gets a reference to the given string and assigns it to the TokenKey field.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) SetTokenKey(v string) {
	o.TokenKey = &v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) GetType() ObservabilityPipelineSplunkHecMetricsDestinationType {
	if o == nil {
		var ret ObservabilityPipelineSplunkHecMetricsDestinationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) GetTypeOk() (*ObservabilityPipelineSplunkHecMetricsDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) SetType(v ObservabilityPipelineSplunkHecMetricsDestinationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineSplunkHecMetricsDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Buffer != nil {
		toSerialize["buffer"] = o.Buffer
	}
	if o.Compression != nil {
		toSerialize["compression"] = o.Compression
	}
	if o.DefaultNamespace != nil {
		toSerialize["default_namespace"] = o.DefaultNamespace
	}
	if o.EndpointUrlKey != nil {
		toSerialize["endpoint_url_key"] = o.EndpointUrlKey
	}
	toSerialize["id"] = o.Id
	if o.Index != nil {
		toSerialize["index"] = o.Index
	}
	toSerialize["inputs"] = o.Inputs
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Sourcetype != nil {
		toSerialize["sourcetype"] = o.Sourcetype
	}
	if o.Tls != nil {
		toSerialize["tls"] = o.Tls
	}
	if o.TokenKey != nil {
		toSerialize["token_key"] = o.TokenKey
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineSplunkHecMetricsDestination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Buffer           *ObservabilityPipelineBufferOptions                          `json:"buffer,omitempty"`
		Compression      *ObservabilityPipelineSplunkHecMetricsDestinationCompression `json:"compression,omitempty"`
		DefaultNamespace *string                                                      `json:"default_namespace,omitempty"`
		EndpointUrlKey   *string                                                      `json:"endpoint_url_key,omitempty"`
		Id               *string                                                      `json:"id"`
		Index            *string                                                      `json:"index,omitempty"`
		Inputs           *[]string                                                    `json:"inputs"`
		Source           *string                                                      `json:"source,omitempty"`
		Sourcetype       *string                                                      `json:"sourcetype,omitempty"`
		Tls              *ObservabilityPipelineTls                                    `json:"tls,omitempty"`
		TokenKey         *string                                                      `json:"token_key,omitempty"`
		Type             *ObservabilityPipelineSplunkHecMetricsDestinationType        `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Inputs == nil {
		return fmt.Errorf("required field inputs missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"buffer", "compression", "default_namespace", "endpoint_url_key", "id", "index", "inputs", "source", "sourcetype", "tls", "token_key", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Buffer = all.Buffer
	if all.Compression != nil && !all.Compression.IsValid() {
		hasInvalidField = true
	} else {
		o.Compression = all.Compression
	}
	o.DefaultNamespace = all.DefaultNamespace
	o.EndpointUrlKey = all.EndpointUrlKey
	o.Id = *all.Id
	o.Index = all.Index
	o.Inputs = *all.Inputs
	o.Source = all.Source
	o.Sourcetype = all.Sourcetype
	if all.Tls != nil && all.Tls.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Tls = all.Tls
	o.TokenKey = all.TokenKey
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
