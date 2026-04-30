// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineHttpClientDestination The `http_client` destination sends data to an HTTP endpoint.
//
// **Supported pipeline types:** logs, metrics
type ObservabilityPipelineHttpClientDestination struct {
	// HTTP authentication strategy.
	AuthStrategy *ObservabilityPipelineHttpClientDestinationAuthStrategy `json:"auth_strategy,omitempty"`
	// Compression configuration for HTTP requests.
	Compression *ObservabilityPipelineHttpClientDestinationCompression `json:"compression,omitempty"`
	// Name of the environment variable or secret that holds a custom header value (used with custom auth strategies).
	CustomKey *string `json:"custom_key,omitempty"`
	// Encoding format for log events.
	Encoding ObservabilityPipelineHttpClientDestinationEncoding `json:"encoding"`
	// The unique identifier for this component.
	Id string `json:"id"`
	// A list of component IDs whose output is used as the input for this component.
	Inputs []string `json:"inputs"`
	// Name of the environment variable or secret that holds the password (used when `auth_strategy` is `basic`).
	PasswordKey *string `json:"password_key,omitempty"`
	// Configuration for enabling TLS encryption between the pipeline component and external services.
	Tls *ObservabilityPipelineTls `json:"tls,omitempty"`
	// Name of the environment variable or secret that holds the bearer token (used when `auth_strategy` is `bearer`).
	TokenKey *string `json:"token_key,omitempty"`
	// The destination type. The value should always be `http_client`.
	Type ObservabilityPipelineHttpClientDestinationType `json:"type"`
	// Name of the environment variable or secret that holds the HTTP endpoint URI.
	UriKey *string `json:"uri_key,omitempty"`
	// Name of the environment variable or secret that holds the username (used when `auth_strategy` is `basic`).
	UsernameKey *string `json:"username_key,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineHttpClientDestination instantiates a new ObservabilityPipelineHttpClientDestination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineHttpClientDestination(encoding ObservabilityPipelineHttpClientDestinationEncoding, id string, inputs []string, typeVar ObservabilityPipelineHttpClientDestinationType) *ObservabilityPipelineHttpClientDestination {
	this := ObservabilityPipelineHttpClientDestination{}
	this.Encoding = encoding
	this.Id = id
	this.Inputs = inputs
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineHttpClientDestinationWithDefaults instantiates a new ObservabilityPipelineHttpClientDestination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineHttpClientDestinationWithDefaults() *ObservabilityPipelineHttpClientDestination {
	this := ObservabilityPipelineHttpClientDestination{}
	var typeVar ObservabilityPipelineHttpClientDestinationType = OBSERVABILITYPIPELINEHTTPCLIENTDESTINATIONTYPE_HTTP_CLIENT
	this.Type = typeVar
	return &this
}

// GetAuthStrategy returns the AuthStrategy field value if set, zero value otherwise.
func (o *ObservabilityPipelineHttpClientDestination) GetAuthStrategy() ObservabilityPipelineHttpClientDestinationAuthStrategy {
	if o == nil || o.AuthStrategy == nil {
		var ret ObservabilityPipelineHttpClientDestinationAuthStrategy
		return ret
	}
	return *o.AuthStrategy
}

// GetAuthStrategyOk returns a tuple with the AuthStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineHttpClientDestination) GetAuthStrategyOk() (*ObservabilityPipelineHttpClientDestinationAuthStrategy, bool) {
	if o == nil || o.AuthStrategy == nil {
		return nil, false
	}
	return o.AuthStrategy, true
}

// HasAuthStrategy returns a boolean if a field has been set.
func (o *ObservabilityPipelineHttpClientDestination) HasAuthStrategy() bool {
	return o != nil && o.AuthStrategy != nil
}

// SetAuthStrategy gets a reference to the given ObservabilityPipelineHttpClientDestinationAuthStrategy and assigns it to the AuthStrategy field.
func (o *ObservabilityPipelineHttpClientDestination) SetAuthStrategy(v ObservabilityPipelineHttpClientDestinationAuthStrategy) {
	o.AuthStrategy = &v
}

// GetCompression returns the Compression field value if set, zero value otherwise.
func (o *ObservabilityPipelineHttpClientDestination) GetCompression() ObservabilityPipelineHttpClientDestinationCompression {
	if o == nil || o.Compression == nil {
		var ret ObservabilityPipelineHttpClientDestinationCompression
		return ret
	}
	return *o.Compression
}

// GetCompressionOk returns a tuple with the Compression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineHttpClientDestination) GetCompressionOk() (*ObservabilityPipelineHttpClientDestinationCompression, bool) {
	if o == nil || o.Compression == nil {
		return nil, false
	}
	return o.Compression, true
}

// HasCompression returns a boolean if a field has been set.
func (o *ObservabilityPipelineHttpClientDestination) HasCompression() bool {
	return o != nil && o.Compression != nil
}

// SetCompression gets a reference to the given ObservabilityPipelineHttpClientDestinationCompression and assigns it to the Compression field.
func (o *ObservabilityPipelineHttpClientDestination) SetCompression(v ObservabilityPipelineHttpClientDestinationCompression) {
	o.Compression = &v
}

// GetCustomKey returns the CustomKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineHttpClientDestination) GetCustomKey() string {
	if o == nil || o.CustomKey == nil {
		var ret string
		return ret
	}
	return *o.CustomKey
}

// GetCustomKeyOk returns a tuple with the CustomKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineHttpClientDestination) GetCustomKeyOk() (*string, bool) {
	if o == nil || o.CustomKey == nil {
		return nil, false
	}
	return o.CustomKey, true
}

// HasCustomKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineHttpClientDestination) HasCustomKey() bool {
	return o != nil && o.CustomKey != nil
}

// SetCustomKey gets a reference to the given string and assigns it to the CustomKey field.
func (o *ObservabilityPipelineHttpClientDestination) SetCustomKey(v string) {
	o.CustomKey = &v
}

// GetEncoding returns the Encoding field value.
func (o *ObservabilityPipelineHttpClientDestination) GetEncoding() ObservabilityPipelineHttpClientDestinationEncoding {
	if o == nil {
		var ret ObservabilityPipelineHttpClientDestinationEncoding
		return ret
	}
	return o.Encoding
}

// GetEncodingOk returns a tuple with the Encoding field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineHttpClientDestination) GetEncodingOk() (*ObservabilityPipelineHttpClientDestinationEncoding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Encoding, true
}

// SetEncoding sets field value.
func (o *ObservabilityPipelineHttpClientDestination) SetEncoding(v ObservabilityPipelineHttpClientDestinationEncoding) {
	o.Encoding = v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineHttpClientDestination) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineHttpClientDestination) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineHttpClientDestination) SetId(v string) {
	o.Id = v
}

// GetInputs returns the Inputs field value.
func (o *ObservabilityPipelineHttpClientDestination) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineHttpClientDestination) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *ObservabilityPipelineHttpClientDestination) SetInputs(v []string) {
	o.Inputs = v
}

// GetPasswordKey returns the PasswordKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineHttpClientDestination) GetPasswordKey() string {
	if o == nil || o.PasswordKey == nil {
		var ret string
		return ret
	}
	return *o.PasswordKey
}

// GetPasswordKeyOk returns a tuple with the PasswordKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineHttpClientDestination) GetPasswordKeyOk() (*string, bool) {
	if o == nil || o.PasswordKey == nil {
		return nil, false
	}
	return o.PasswordKey, true
}

// HasPasswordKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineHttpClientDestination) HasPasswordKey() bool {
	return o != nil && o.PasswordKey != nil
}

// SetPasswordKey gets a reference to the given string and assigns it to the PasswordKey field.
func (o *ObservabilityPipelineHttpClientDestination) SetPasswordKey(v string) {
	o.PasswordKey = &v
}

// GetTls returns the Tls field value if set, zero value otherwise.
func (o *ObservabilityPipelineHttpClientDestination) GetTls() ObservabilityPipelineTls {
	if o == nil || o.Tls == nil {
		var ret ObservabilityPipelineTls
		return ret
	}
	return *o.Tls
}

// GetTlsOk returns a tuple with the Tls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineHttpClientDestination) GetTlsOk() (*ObservabilityPipelineTls, bool) {
	if o == nil || o.Tls == nil {
		return nil, false
	}
	return o.Tls, true
}

// HasTls returns a boolean if a field has been set.
func (o *ObservabilityPipelineHttpClientDestination) HasTls() bool {
	return o != nil && o.Tls != nil
}

// SetTls gets a reference to the given ObservabilityPipelineTls and assigns it to the Tls field.
func (o *ObservabilityPipelineHttpClientDestination) SetTls(v ObservabilityPipelineTls) {
	o.Tls = &v
}

// GetTokenKey returns the TokenKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineHttpClientDestination) GetTokenKey() string {
	if o == nil || o.TokenKey == nil {
		var ret string
		return ret
	}
	return *o.TokenKey
}

// GetTokenKeyOk returns a tuple with the TokenKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineHttpClientDestination) GetTokenKeyOk() (*string, bool) {
	if o == nil || o.TokenKey == nil {
		return nil, false
	}
	return o.TokenKey, true
}

// HasTokenKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineHttpClientDestination) HasTokenKey() bool {
	return o != nil && o.TokenKey != nil
}

// SetTokenKey gets a reference to the given string and assigns it to the TokenKey field.
func (o *ObservabilityPipelineHttpClientDestination) SetTokenKey(v string) {
	o.TokenKey = &v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineHttpClientDestination) GetType() ObservabilityPipelineHttpClientDestinationType {
	if o == nil {
		var ret ObservabilityPipelineHttpClientDestinationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineHttpClientDestination) GetTypeOk() (*ObservabilityPipelineHttpClientDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineHttpClientDestination) SetType(v ObservabilityPipelineHttpClientDestinationType) {
	o.Type = v
}

// GetUriKey returns the UriKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineHttpClientDestination) GetUriKey() string {
	if o == nil || o.UriKey == nil {
		var ret string
		return ret
	}
	return *o.UriKey
}

// GetUriKeyOk returns a tuple with the UriKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineHttpClientDestination) GetUriKeyOk() (*string, bool) {
	if o == nil || o.UriKey == nil {
		return nil, false
	}
	return o.UriKey, true
}

// HasUriKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineHttpClientDestination) HasUriKey() bool {
	return o != nil && o.UriKey != nil
}

// SetUriKey gets a reference to the given string and assigns it to the UriKey field.
func (o *ObservabilityPipelineHttpClientDestination) SetUriKey(v string) {
	o.UriKey = &v
}

// GetUsernameKey returns the UsernameKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineHttpClientDestination) GetUsernameKey() string {
	if o == nil || o.UsernameKey == nil {
		var ret string
		return ret
	}
	return *o.UsernameKey
}

// GetUsernameKeyOk returns a tuple with the UsernameKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineHttpClientDestination) GetUsernameKeyOk() (*string, bool) {
	if o == nil || o.UsernameKey == nil {
		return nil, false
	}
	return o.UsernameKey, true
}

// HasUsernameKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineHttpClientDestination) HasUsernameKey() bool {
	return o != nil && o.UsernameKey != nil
}

// SetUsernameKey gets a reference to the given string and assigns it to the UsernameKey field.
func (o *ObservabilityPipelineHttpClientDestination) SetUsernameKey(v string) {
	o.UsernameKey = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineHttpClientDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AuthStrategy != nil {
		toSerialize["auth_strategy"] = o.AuthStrategy
	}
	if o.Compression != nil {
		toSerialize["compression"] = o.Compression
	}
	if o.CustomKey != nil {
		toSerialize["custom_key"] = o.CustomKey
	}
	toSerialize["encoding"] = o.Encoding
	toSerialize["id"] = o.Id
	toSerialize["inputs"] = o.Inputs
	if o.PasswordKey != nil {
		toSerialize["password_key"] = o.PasswordKey
	}
	if o.Tls != nil {
		toSerialize["tls"] = o.Tls
	}
	if o.TokenKey != nil {
		toSerialize["token_key"] = o.TokenKey
	}
	toSerialize["type"] = o.Type
	if o.UriKey != nil {
		toSerialize["uri_key"] = o.UriKey
	}
	if o.UsernameKey != nil {
		toSerialize["username_key"] = o.UsernameKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineHttpClientDestination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthStrategy *ObservabilityPipelineHttpClientDestinationAuthStrategy `json:"auth_strategy,omitempty"`
		Compression  *ObservabilityPipelineHttpClientDestinationCompression  `json:"compression,omitempty"`
		CustomKey    *string                                                 `json:"custom_key,omitempty"`
		Encoding     *ObservabilityPipelineHttpClientDestinationEncoding     `json:"encoding"`
		Id           *string                                                 `json:"id"`
		Inputs       *[]string                                               `json:"inputs"`
		PasswordKey  *string                                                 `json:"password_key,omitempty"`
		Tls          *ObservabilityPipelineTls                               `json:"tls,omitempty"`
		TokenKey     *string                                                 `json:"token_key,omitempty"`
		Type         *ObservabilityPipelineHttpClientDestinationType         `json:"type"`
		UriKey       *string                                                 `json:"uri_key,omitempty"`
		UsernameKey  *string                                                 `json:"username_key,omitempty"`
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
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auth_strategy", "compression", "custom_key", "encoding", "id", "inputs", "password_key", "tls", "token_key", "type", "uri_key", "username_key"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AuthStrategy != nil && !all.AuthStrategy.IsValid() {
		hasInvalidField = true
	} else {
		o.AuthStrategy = all.AuthStrategy
	}
	if all.Compression != nil && all.Compression.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Compression = all.Compression
	o.CustomKey = all.CustomKey
	if !all.Encoding.IsValid() {
		hasInvalidField = true
	} else {
		o.Encoding = *all.Encoding
	}
	o.Id = *all.Id
	o.Inputs = *all.Inputs
	o.PasswordKey = all.PasswordKey
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
	o.UriKey = all.UriKey
	o.UsernameKey = all.UsernameKey

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
