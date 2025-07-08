// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineHttpServerSource The `http_server` source collects logs over HTTP POST from external services.
type ObservabilityPipelineHttpServerSource struct {
	// HTTP authentication method.
	AuthStrategy ObservabilityPipelineHttpServerSourceAuthStrategy `json:"auth_strategy"`
	// The decoding format used to interpret incoming logs.
	Decoding ObservabilityPipelineDecoding `json:"decoding"`
	// Unique ID for the HTTP server source.
	Id string `json:"id"`
	// Configuration for enabling TLS encryption between the pipeline component and external services.
	Tls *ObservabilityPipelineTls `json:"tls,omitempty"`
	// The source type. The value should always be `http_server`.
	Type ObservabilityPipelineHttpServerSourceType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineHttpServerSource instantiates a new ObservabilityPipelineHttpServerSource object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineHttpServerSource(authStrategy ObservabilityPipelineHttpServerSourceAuthStrategy, decoding ObservabilityPipelineDecoding, id string, typeVar ObservabilityPipelineHttpServerSourceType) *ObservabilityPipelineHttpServerSource {
	this := ObservabilityPipelineHttpServerSource{}
	this.AuthStrategy = authStrategy
	this.Decoding = decoding
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineHttpServerSourceWithDefaults instantiates a new ObservabilityPipelineHttpServerSource object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineHttpServerSourceWithDefaults() *ObservabilityPipelineHttpServerSource {
	this := ObservabilityPipelineHttpServerSource{}
	var typeVar ObservabilityPipelineHttpServerSourceType = OBSERVABILITYPIPELINEHTTPSERVERSOURCETYPE_HTTP_SERVER
	this.Type = typeVar
	return &this
}

// GetAuthStrategy returns the AuthStrategy field value.
func (o *ObservabilityPipelineHttpServerSource) GetAuthStrategy() ObservabilityPipelineHttpServerSourceAuthStrategy {
	if o == nil {
		var ret ObservabilityPipelineHttpServerSourceAuthStrategy
		return ret
	}
	return o.AuthStrategy
}

// GetAuthStrategyOk returns a tuple with the AuthStrategy field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineHttpServerSource) GetAuthStrategyOk() (*ObservabilityPipelineHttpServerSourceAuthStrategy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthStrategy, true
}

// SetAuthStrategy sets field value.
func (o *ObservabilityPipelineHttpServerSource) SetAuthStrategy(v ObservabilityPipelineHttpServerSourceAuthStrategy) {
	o.AuthStrategy = v
}

// GetDecoding returns the Decoding field value.
func (o *ObservabilityPipelineHttpServerSource) GetDecoding() ObservabilityPipelineDecoding {
	if o == nil {
		var ret ObservabilityPipelineDecoding
		return ret
	}
	return o.Decoding
}

// GetDecodingOk returns a tuple with the Decoding field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineHttpServerSource) GetDecodingOk() (*ObservabilityPipelineDecoding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Decoding, true
}

// SetDecoding sets field value.
func (o *ObservabilityPipelineHttpServerSource) SetDecoding(v ObservabilityPipelineDecoding) {
	o.Decoding = v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineHttpServerSource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineHttpServerSource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineHttpServerSource) SetId(v string) {
	o.Id = v
}

// GetTls returns the Tls field value if set, zero value otherwise.
func (o *ObservabilityPipelineHttpServerSource) GetTls() ObservabilityPipelineTls {
	if o == nil || o.Tls == nil {
		var ret ObservabilityPipelineTls
		return ret
	}
	return *o.Tls
}

// GetTlsOk returns a tuple with the Tls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineHttpServerSource) GetTlsOk() (*ObservabilityPipelineTls, bool) {
	if o == nil || o.Tls == nil {
		return nil, false
	}
	return o.Tls, true
}

// HasTls returns a boolean if a field has been set.
func (o *ObservabilityPipelineHttpServerSource) HasTls() bool {
	return o != nil && o.Tls != nil
}

// SetTls gets a reference to the given ObservabilityPipelineTls and assigns it to the Tls field.
func (o *ObservabilityPipelineHttpServerSource) SetTls(v ObservabilityPipelineTls) {
	o.Tls = &v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineHttpServerSource) GetType() ObservabilityPipelineHttpServerSourceType {
	if o == nil {
		var ret ObservabilityPipelineHttpServerSourceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineHttpServerSource) GetTypeOk() (*ObservabilityPipelineHttpServerSourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineHttpServerSource) SetType(v ObservabilityPipelineHttpServerSourceType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineHttpServerSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["auth_strategy"] = o.AuthStrategy
	toSerialize["decoding"] = o.Decoding
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
func (o *ObservabilityPipelineHttpServerSource) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthStrategy *ObservabilityPipelineHttpServerSourceAuthStrategy `json:"auth_strategy"`
		Decoding     *ObservabilityPipelineDecoding                     `json:"decoding"`
		Id           *string                                            `json:"id"`
		Tls          *ObservabilityPipelineTls                          `json:"tls,omitempty"`
		Type         *ObservabilityPipelineHttpServerSourceType         `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AuthStrategy == nil {
		return fmt.Errorf("required field auth_strategy missing")
	}
	if all.Decoding == nil {
		return fmt.Errorf("required field decoding missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auth_strategy", "decoding", "id", "tls", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.AuthStrategy.IsValid() {
		hasInvalidField = true
	} else {
		o.AuthStrategy = *all.AuthStrategy
	}
	if !all.Decoding.IsValid() {
		hasInvalidField = true
	} else {
		o.Decoding = *all.Decoding
	}
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
