// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineOpenSearchDestination The `opensearch` destination writes logs to an OpenSearch cluster.
//
// **Supported pipeline types:** logs
type ObservabilityPipelineOpenSearchDestination struct {
	// Authentication settings for the Elasticsearch destination.
	// When `strategy` is `basic`, use `username_key` and `password_key` to reference credentials stored in environment variables or secrets.
	Auth *ObservabilityPipelineElasticsearchDestinationAuth `json:"auth,omitempty"`
	// Configuration for buffer settings on destination components.
	Buffer *ObservabilityPipelineBufferOptions `json:"buffer,omitempty"`
	// The index to write logs to.
	BulkIndex *string `json:"bulk_index,omitempty"`
	// Configuration options for writing to OpenSearch Data Streams instead of a fixed index.
	DataStream *ObservabilityPipelineOpenSearchDestinationDataStream `json:"data_stream,omitempty"`
	// Name of the environment variable or secret that holds the OpenSearch endpoint URL.
	EndpointUrlKey *string `json:"endpoint_url_key,omitempty"`
	// The unique identifier for this component.
	Id string `json:"id"`
	// A list of component IDs whose output is used as the `input` for this component.
	Inputs []string `json:"inputs"`
	// The destination type. The value should always be `opensearch`.
	Type ObservabilityPipelineOpenSearchDestinationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineOpenSearchDestination instantiates a new ObservabilityPipelineOpenSearchDestination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineOpenSearchDestination(id string, inputs []string, typeVar ObservabilityPipelineOpenSearchDestinationType) *ObservabilityPipelineOpenSearchDestination {
	this := ObservabilityPipelineOpenSearchDestination{}
	this.Id = id
	this.Inputs = inputs
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineOpenSearchDestinationWithDefaults instantiates a new ObservabilityPipelineOpenSearchDestination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineOpenSearchDestinationWithDefaults() *ObservabilityPipelineOpenSearchDestination {
	this := ObservabilityPipelineOpenSearchDestination{}
	var typeVar ObservabilityPipelineOpenSearchDestinationType = OBSERVABILITYPIPELINEOPENSEARCHDESTINATIONTYPE_OPENSEARCH
	this.Type = typeVar
	return &this
}

// GetAuth returns the Auth field value if set, zero value otherwise.
func (o *ObservabilityPipelineOpenSearchDestination) GetAuth() ObservabilityPipelineElasticsearchDestinationAuth {
	if o == nil || o.Auth == nil {
		var ret ObservabilityPipelineElasticsearchDestinationAuth
		return ret
	}
	return *o.Auth
}

// GetAuthOk returns a tuple with the Auth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOpenSearchDestination) GetAuthOk() (*ObservabilityPipelineElasticsearchDestinationAuth, bool) {
	if o == nil || o.Auth == nil {
		return nil, false
	}
	return o.Auth, true
}

// HasAuth returns a boolean if a field has been set.
func (o *ObservabilityPipelineOpenSearchDestination) HasAuth() bool {
	return o != nil && o.Auth != nil
}

// SetAuth gets a reference to the given ObservabilityPipelineElasticsearchDestinationAuth and assigns it to the Auth field.
func (o *ObservabilityPipelineOpenSearchDestination) SetAuth(v ObservabilityPipelineElasticsearchDestinationAuth) {
	o.Auth = &v
}

// GetBuffer returns the Buffer field value if set, zero value otherwise.
func (o *ObservabilityPipelineOpenSearchDestination) GetBuffer() ObservabilityPipelineBufferOptions {
	if o == nil || o.Buffer == nil {
		var ret ObservabilityPipelineBufferOptions
		return ret
	}
	return *o.Buffer
}

// GetBufferOk returns a tuple with the Buffer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOpenSearchDestination) GetBufferOk() (*ObservabilityPipelineBufferOptions, bool) {
	if o == nil || o.Buffer == nil {
		return nil, false
	}
	return o.Buffer, true
}

// HasBuffer returns a boolean if a field has been set.
func (o *ObservabilityPipelineOpenSearchDestination) HasBuffer() bool {
	return o != nil && o.Buffer != nil
}

// SetBuffer gets a reference to the given ObservabilityPipelineBufferOptions and assigns it to the Buffer field.
func (o *ObservabilityPipelineOpenSearchDestination) SetBuffer(v ObservabilityPipelineBufferOptions) {
	o.Buffer = &v
}

// GetBulkIndex returns the BulkIndex field value if set, zero value otherwise.
func (o *ObservabilityPipelineOpenSearchDestination) GetBulkIndex() string {
	if o == nil || o.BulkIndex == nil {
		var ret string
		return ret
	}
	return *o.BulkIndex
}

// GetBulkIndexOk returns a tuple with the BulkIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOpenSearchDestination) GetBulkIndexOk() (*string, bool) {
	if o == nil || o.BulkIndex == nil {
		return nil, false
	}
	return o.BulkIndex, true
}

// HasBulkIndex returns a boolean if a field has been set.
func (o *ObservabilityPipelineOpenSearchDestination) HasBulkIndex() bool {
	return o != nil && o.BulkIndex != nil
}

// SetBulkIndex gets a reference to the given string and assigns it to the BulkIndex field.
func (o *ObservabilityPipelineOpenSearchDestination) SetBulkIndex(v string) {
	o.BulkIndex = &v
}

// GetDataStream returns the DataStream field value if set, zero value otherwise.
func (o *ObservabilityPipelineOpenSearchDestination) GetDataStream() ObservabilityPipelineOpenSearchDestinationDataStream {
	if o == nil || o.DataStream == nil {
		var ret ObservabilityPipelineOpenSearchDestinationDataStream
		return ret
	}
	return *o.DataStream
}

// GetDataStreamOk returns a tuple with the DataStream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOpenSearchDestination) GetDataStreamOk() (*ObservabilityPipelineOpenSearchDestinationDataStream, bool) {
	if o == nil || o.DataStream == nil {
		return nil, false
	}
	return o.DataStream, true
}

// HasDataStream returns a boolean if a field has been set.
func (o *ObservabilityPipelineOpenSearchDestination) HasDataStream() bool {
	return o != nil && o.DataStream != nil
}

// SetDataStream gets a reference to the given ObservabilityPipelineOpenSearchDestinationDataStream and assigns it to the DataStream field.
func (o *ObservabilityPipelineOpenSearchDestination) SetDataStream(v ObservabilityPipelineOpenSearchDestinationDataStream) {
	o.DataStream = &v
}

// GetEndpointUrlKey returns the EndpointUrlKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineOpenSearchDestination) GetEndpointUrlKey() string {
	if o == nil || o.EndpointUrlKey == nil {
		var ret string
		return ret
	}
	return *o.EndpointUrlKey
}

// GetEndpointUrlKeyOk returns a tuple with the EndpointUrlKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOpenSearchDestination) GetEndpointUrlKeyOk() (*string, bool) {
	if o == nil || o.EndpointUrlKey == nil {
		return nil, false
	}
	return o.EndpointUrlKey, true
}

// HasEndpointUrlKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineOpenSearchDestination) HasEndpointUrlKey() bool {
	return o != nil && o.EndpointUrlKey != nil
}

// SetEndpointUrlKey gets a reference to the given string and assigns it to the EndpointUrlKey field.
func (o *ObservabilityPipelineOpenSearchDestination) SetEndpointUrlKey(v string) {
	o.EndpointUrlKey = &v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineOpenSearchDestination) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOpenSearchDestination) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineOpenSearchDestination) SetId(v string) {
	o.Id = v
}

// GetInputs returns the Inputs field value.
func (o *ObservabilityPipelineOpenSearchDestination) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOpenSearchDestination) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *ObservabilityPipelineOpenSearchDestination) SetInputs(v []string) {
	o.Inputs = v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineOpenSearchDestination) GetType() ObservabilityPipelineOpenSearchDestinationType {
	if o == nil {
		var ret ObservabilityPipelineOpenSearchDestinationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOpenSearchDestination) GetTypeOk() (*ObservabilityPipelineOpenSearchDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineOpenSearchDestination) SetType(v ObservabilityPipelineOpenSearchDestinationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineOpenSearchDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Auth != nil {
		toSerialize["auth"] = o.Auth
	}
	if o.Buffer != nil {
		toSerialize["buffer"] = o.Buffer
	}
	if o.BulkIndex != nil {
		toSerialize["bulk_index"] = o.BulkIndex
	}
	if o.DataStream != nil {
		toSerialize["data_stream"] = o.DataStream
	}
	if o.EndpointUrlKey != nil {
		toSerialize["endpoint_url_key"] = o.EndpointUrlKey
	}
	toSerialize["id"] = o.Id
	toSerialize["inputs"] = o.Inputs
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineOpenSearchDestination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Auth           *ObservabilityPipelineElasticsearchDestinationAuth    `json:"auth,omitempty"`
		Buffer         *ObservabilityPipelineBufferOptions                   `json:"buffer,omitempty"`
		BulkIndex      *string                                               `json:"bulk_index,omitempty"`
		DataStream     *ObservabilityPipelineOpenSearchDestinationDataStream `json:"data_stream,omitempty"`
		EndpointUrlKey *string                                               `json:"endpoint_url_key,omitempty"`
		Id             *string                                               `json:"id"`
		Inputs         *[]string                                             `json:"inputs"`
		Type           *ObservabilityPipelineOpenSearchDestinationType       `json:"type"`
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
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auth", "buffer", "bulk_index", "data_stream", "endpoint_url_key", "id", "inputs", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Auth != nil && all.Auth.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Auth = all.Auth
	o.Buffer = all.Buffer
	o.BulkIndex = all.BulkIndex
	if all.DataStream != nil && all.DataStream.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DataStream = all.DataStream
	o.EndpointUrlKey = all.EndpointUrlKey
	o.Id = *all.Id
	o.Inputs = *all.Inputs
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
