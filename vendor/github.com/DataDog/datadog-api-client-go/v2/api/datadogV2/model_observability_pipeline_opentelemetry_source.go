// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineOpentelemetrySource The `opentelemetry` source receives telemetry data using the OpenTelemetry Protocol (OTLP) over gRPC and HTTP.
//
// **Supported pipeline types:** logs, metrics
type ObservabilityPipelineOpentelemetrySource struct {
	// Environment variable name containing the gRPC server address for receiving OTLP data. Must be a valid environment variable name (alphanumeric characters and underscores only).
	GrpcAddressKey *string `json:"grpc_address_key,omitempty"`
	// Environment variable name containing the HTTP server address for receiving OTLP data. Must be a valid environment variable name (alphanumeric characters and underscores only).
	HttpAddressKey *string `json:"http_address_key,omitempty"`
	// The unique identifier for this component. Used in other parts of the pipeline to reference this component (for example, as the `input` to downstream components).
	Id string `json:"id"`
	// Configuration for enabling TLS encryption between the pipeline component and external services.
	Tls *ObservabilityPipelineTls `json:"tls,omitempty"`
	// The source type. The value should always be `opentelemetry`.
	Type ObservabilityPipelineOpentelemetrySourceType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineOpentelemetrySource instantiates a new ObservabilityPipelineOpentelemetrySource object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineOpentelemetrySource(id string, typeVar ObservabilityPipelineOpentelemetrySourceType) *ObservabilityPipelineOpentelemetrySource {
	this := ObservabilityPipelineOpentelemetrySource{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineOpentelemetrySourceWithDefaults instantiates a new ObservabilityPipelineOpentelemetrySource object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineOpentelemetrySourceWithDefaults() *ObservabilityPipelineOpentelemetrySource {
	this := ObservabilityPipelineOpentelemetrySource{}
	var typeVar ObservabilityPipelineOpentelemetrySourceType = OBSERVABILITYPIPELINEOPENTELEMETRYSOURCETYPE_OPENTELEMETRY
	this.Type = typeVar
	return &this
}

// GetGrpcAddressKey returns the GrpcAddressKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineOpentelemetrySource) GetGrpcAddressKey() string {
	if o == nil || o.GrpcAddressKey == nil {
		var ret string
		return ret
	}
	return *o.GrpcAddressKey
}

// GetGrpcAddressKeyOk returns a tuple with the GrpcAddressKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOpentelemetrySource) GetGrpcAddressKeyOk() (*string, bool) {
	if o == nil || o.GrpcAddressKey == nil {
		return nil, false
	}
	return o.GrpcAddressKey, true
}

// HasGrpcAddressKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineOpentelemetrySource) HasGrpcAddressKey() bool {
	return o != nil && o.GrpcAddressKey != nil
}

// SetGrpcAddressKey gets a reference to the given string and assigns it to the GrpcAddressKey field.
func (o *ObservabilityPipelineOpentelemetrySource) SetGrpcAddressKey(v string) {
	o.GrpcAddressKey = &v
}

// GetHttpAddressKey returns the HttpAddressKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineOpentelemetrySource) GetHttpAddressKey() string {
	if o == nil || o.HttpAddressKey == nil {
		var ret string
		return ret
	}
	return *o.HttpAddressKey
}

// GetHttpAddressKeyOk returns a tuple with the HttpAddressKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOpentelemetrySource) GetHttpAddressKeyOk() (*string, bool) {
	if o == nil || o.HttpAddressKey == nil {
		return nil, false
	}
	return o.HttpAddressKey, true
}

// HasHttpAddressKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineOpentelemetrySource) HasHttpAddressKey() bool {
	return o != nil && o.HttpAddressKey != nil
}

// SetHttpAddressKey gets a reference to the given string and assigns it to the HttpAddressKey field.
func (o *ObservabilityPipelineOpentelemetrySource) SetHttpAddressKey(v string) {
	o.HttpAddressKey = &v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineOpentelemetrySource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOpentelemetrySource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineOpentelemetrySource) SetId(v string) {
	o.Id = v
}

// GetTls returns the Tls field value if set, zero value otherwise.
func (o *ObservabilityPipelineOpentelemetrySource) GetTls() ObservabilityPipelineTls {
	if o == nil || o.Tls == nil {
		var ret ObservabilityPipelineTls
		return ret
	}
	return *o.Tls
}

// GetTlsOk returns a tuple with the Tls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOpentelemetrySource) GetTlsOk() (*ObservabilityPipelineTls, bool) {
	if o == nil || o.Tls == nil {
		return nil, false
	}
	return o.Tls, true
}

// HasTls returns a boolean if a field has been set.
func (o *ObservabilityPipelineOpentelemetrySource) HasTls() bool {
	return o != nil && o.Tls != nil
}

// SetTls gets a reference to the given ObservabilityPipelineTls and assigns it to the Tls field.
func (o *ObservabilityPipelineOpentelemetrySource) SetTls(v ObservabilityPipelineTls) {
	o.Tls = &v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineOpentelemetrySource) GetType() ObservabilityPipelineOpentelemetrySourceType {
	if o == nil {
		var ret ObservabilityPipelineOpentelemetrySourceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOpentelemetrySource) GetTypeOk() (*ObservabilityPipelineOpentelemetrySourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineOpentelemetrySource) SetType(v ObservabilityPipelineOpentelemetrySourceType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineOpentelemetrySource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.GrpcAddressKey != nil {
		toSerialize["grpc_address_key"] = o.GrpcAddressKey
	}
	if o.HttpAddressKey != nil {
		toSerialize["http_address_key"] = o.HttpAddressKey
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
func (o *ObservabilityPipelineOpentelemetrySource) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		GrpcAddressKey *string                                       `json:"grpc_address_key,omitempty"`
		HttpAddressKey *string                                       `json:"http_address_key,omitempty"`
		Id             *string                                       `json:"id"`
		Tls            *ObservabilityPipelineTls                     `json:"tls,omitempty"`
		Type           *ObservabilityPipelineOpentelemetrySourceType `json:"type"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"grpc_address_key", "http_address_key", "id", "tls", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.GrpcAddressKey = all.GrpcAddressKey
	o.HttpAddressKey = all.HttpAddressKey
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
