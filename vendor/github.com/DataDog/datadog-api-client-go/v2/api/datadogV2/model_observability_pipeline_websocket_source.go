// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineWebsocketSource The `websocket` source ingests logs from a WebSocket server using the `ws://` or `wss://` protocol.
//
// **Supported pipeline types:** logs.
type ObservabilityPipelineWebsocketSource struct {
	// Authentication strategy for the WebSocket source connection.
	AuthStrategy ObservabilityPipelineWebsocketSourceAuthStrategy `json:"auth_strategy"`
	// Name of the environment variable or secret that holds the custom authorization header value. Used when `auth_strategy` is `custom`.
	CustomKey *string `json:"custom_key,omitempty"`
	// The decoding format used to interpret incoming logs.
	Decoding ObservabilityPipelineDecoding `json:"decoding"`
	// The unique identifier for this component.
	Id string `json:"id"`
	// Name of the environment variable or secret that holds the password. Used when `auth_strategy` is `basic`.
	PasswordKey *string `json:"password_key,omitempty"`
	// TLS configuration for the WebSocket source. Use `enabled` for standard `wss://` connections, or `with_client_cert` to present a client certificate for mutual TLS.
	Tls *ObservabilityPipelineWebsocketSourceTls `json:"tls,omitempty"`
	// Name of the environment variable or secret that holds the bearer token. Used when `auth_strategy` is `bearer`.
	TokenKey *string `json:"token_key,omitempty"`
	// The source type. The value should always be `websocket`.
	Type ObservabilityPipelineWebsocketSourceType `json:"type"`
	// Name of the environment variable or secret that holds the WebSocket server URI (`ws://` or `wss://`).
	UriKey *string `json:"uri_key,omitempty"`
	// Name of the environment variable or secret that holds the username. Used when `auth_strategy` is `basic`.
	UsernameKey *string `json:"username_key,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineWebsocketSource instantiates a new ObservabilityPipelineWebsocketSource object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineWebsocketSource(authStrategy ObservabilityPipelineWebsocketSourceAuthStrategy, decoding ObservabilityPipelineDecoding, id string, typeVar ObservabilityPipelineWebsocketSourceType) *ObservabilityPipelineWebsocketSource {
	this := ObservabilityPipelineWebsocketSource{}
	this.AuthStrategy = authStrategy
	this.Decoding = decoding
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineWebsocketSourceWithDefaults instantiates a new ObservabilityPipelineWebsocketSource object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineWebsocketSourceWithDefaults() *ObservabilityPipelineWebsocketSource {
	this := ObservabilityPipelineWebsocketSource{}
	var typeVar ObservabilityPipelineWebsocketSourceType = OBSERVABILITYPIPELINEWEBSOCKETSOURCETYPE_WEBSOCKET
	this.Type = typeVar
	return &this
}

// GetAuthStrategy returns the AuthStrategy field value.
func (o *ObservabilityPipelineWebsocketSource) GetAuthStrategy() ObservabilityPipelineWebsocketSourceAuthStrategy {
	if o == nil {
		var ret ObservabilityPipelineWebsocketSourceAuthStrategy
		return ret
	}
	return o.AuthStrategy
}

// GetAuthStrategyOk returns a tuple with the AuthStrategy field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineWebsocketSource) GetAuthStrategyOk() (*ObservabilityPipelineWebsocketSourceAuthStrategy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthStrategy, true
}

// SetAuthStrategy sets field value.
func (o *ObservabilityPipelineWebsocketSource) SetAuthStrategy(v ObservabilityPipelineWebsocketSourceAuthStrategy) {
	o.AuthStrategy = v
}

// GetCustomKey returns the CustomKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineWebsocketSource) GetCustomKey() string {
	if o == nil || o.CustomKey == nil {
		var ret string
		return ret
	}
	return *o.CustomKey
}

// GetCustomKeyOk returns a tuple with the CustomKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineWebsocketSource) GetCustomKeyOk() (*string, bool) {
	if o == nil || o.CustomKey == nil {
		return nil, false
	}
	return o.CustomKey, true
}

// HasCustomKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineWebsocketSource) HasCustomKey() bool {
	return o != nil && o.CustomKey != nil
}

// SetCustomKey gets a reference to the given string and assigns it to the CustomKey field.
func (o *ObservabilityPipelineWebsocketSource) SetCustomKey(v string) {
	o.CustomKey = &v
}

// GetDecoding returns the Decoding field value.
func (o *ObservabilityPipelineWebsocketSource) GetDecoding() ObservabilityPipelineDecoding {
	if o == nil {
		var ret ObservabilityPipelineDecoding
		return ret
	}
	return o.Decoding
}

// GetDecodingOk returns a tuple with the Decoding field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineWebsocketSource) GetDecodingOk() (*ObservabilityPipelineDecoding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Decoding, true
}

// SetDecoding sets field value.
func (o *ObservabilityPipelineWebsocketSource) SetDecoding(v ObservabilityPipelineDecoding) {
	o.Decoding = v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineWebsocketSource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineWebsocketSource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineWebsocketSource) SetId(v string) {
	o.Id = v
}

// GetPasswordKey returns the PasswordKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineWebsocketSource) GetPasswordKey() string {
	if o == nil || o.PasswordKey == nil {
		var ret string
		return ret
	}
	return *o.PasswordKey
}

// GetPasswordKeyOk returns a tuple with the PasswordKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineWebsocketSource) GetPasswordKeyOk() (*string, bool) {
	if o == nil || o.PasswordKey == nil {
		return nil, false
	}
	return o.PasswordKey, true
}

// HasPasswordKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineWebsocketSource) HasPasswordKey() bool {
	return o != nil && o.PasswordKey != nil
}

// SetPasswordKey gets a reference to the given string and assigns it to the PasswordKey field.
func (o *ObservabilityPipelineWebsocketSource) SetPasswordKey(v string) {
	o.PasswordKey = &v
}

// GetTls returns the Tls field value if set, zero value otherwise.
func (o *ObservabilityPipelineWebsocketSource) GetTls() ObservabilityPipelineWebsocketSourceTls {
	if o == nil || o.Tls == nil {
		var ret ObservabilityPipelineWebsocketSourceTls
		return ret
	}
	return *o.Tls
}

// GetTlsOk returns a tuple with the Tls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineWebsocketSource) GetTlsOk() (*ObservabilityPipelineWebsocketSourceTls, bool) {
	if o == nil || o.Tls == nil {
		return nil, false
	}
	return o.Tls, true
}

// HasTls returns a boolean if a field has been set.
func (o *ObservabilityPipelineWebsocketSource) HasTls() bool {
	return o != nil && o.Tls != nil
}

// SetTls gets a reference to the given ObservabilityPipelineWebsocketSourceTls and assigns it to the Tls field.
func (o *ObservabilityPipelineWebsocketSource) SetTls(v ObservabilityPipelineWebsocketSourceTls) {
	o.Tls = &v
}

// GetTokenKey returns the TokenKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineWebsocketSource) GetTokenKey() string {
	if o == nil || o.TokenKey == nil {
		var ret string
		return ret
	}
	return *o.TokenKey
}

// GetTokenKeyOk returns a tuple with the TokenKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineWebsocketSource) GetTokenKeyOk() (*string, bool) {
	if o == nil || o.TokenKey == nil {
		return nil, false
	}
	return o.TokenKey, true
}

// HasTokenKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineWebsocketSource) HasTokenKey() bool {
	return o != nil && o.TokenKey != nil
}

// SetTokenKey gets a reference to the given string and assigns it to the TokenKey field.
func (o *ObservabilityPipelineWebsocketSource) SetTokenKey(v string) {
	o.TokenKey = &v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineWebsocketSource) GetType() ObservabilityPipelineWebsocketSourceType {
	if o == nil {
		var ret ObservabilityPipelineWebsocketSourceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineWebsocketSource) GetTypeOk() (*ObservabilityPipelineWebsocketSourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineWebsocketSource) SetType(v ObservabilityPipelineWebsocketSourceType) {
	o.Type = v
}

// GetUriKey returns the UriKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineWebsocketSource) GetUriKey() string {
	if o == nil || o.UriKey == nil {
		var ret string
		return ret
	}
	return *o.UriKey
}

// GetUriKeyOk returns a tuple with the UriKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineWebsocketSource) GetUriKeyOk() (*string, bool) {
	if o == nil || o.UriKey == nil {
		return nil, false
	}
	return o.UriKey, true
}

// HasUriKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineWebsocketSource) HasUriKey() bool {
	return o != nil && o.UriKey != nil
}

// SetUriKey gets a reference to the given string and assigns it to the UriKey field.
func (o *ObservabilityPipelineWebsocketSource) SetUriKey(v string) {
	o.UriKey = &v
}

// GetUsernameKey returns the UsernameKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineWebsocketSource) GetUsernameKey() string {
	if o == nil || o.UsernameKey == nil {
		var ret string
		return ret
	}
	return *o.UsernameKey
}

// GetUsernameKeyOk returns a tuple with the UsernameKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineWebsocketSource) GetUsernameKeyOk() (*string, bool) {
	if o == nil || o.UsernameKey == nil {
		return nil, false
	}
	return o.UsernameKey, true
}

// HasUsernameKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineWebsocketSource) HasUsernameKey() bool {
	return o != nil && o.UsernameKey != nil
}

// SetUsernameKey gets a reference to the given string and assigns it to the UsernameKey field.
func (o *ObservabilityPipelineWebsocketSource) SetUsernameKey(v string) {
	o.UsernameKey = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineWebsocketSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["auth_strategy"] = o.AuthStrategy
	if o.CustomKey != nil {
		toSerialize["custom_key"] = o.CustomKey
	}
	toSerialize["decoding"] = o.Decoding
	toSerialize["id"] = o.Id
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
func (o *ObservabilityPipelineWebsocketSource) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthStrategy *ObservabilityPipelineWebsocketSourceAuthStrategy `json:"auth_strategy"`
		CustomKey    *string                                           `json:"custom_key,omitempty"`
		Decoding     *ObservabilityPipelineDecoding                    `json:"decoding"`
		Id           *string                                           `json:"id"`
		PasswordKey  *string                                           `json:"password_key,omitempty"`
		Tls          *ObservabilityPipelineWebsocketSourceTls          `json:"tls,omitempty"`
		TokenKey     *string                                           `json:"token_key,omitempty"`
		Type         *ObservabilityPipelineWebsocketSourceType         `json:"type"`
		UriKey       *string                                           `json:"uri_key,omitempty"`
		UsernameKey  *string                                           `json:"username_key,omitempty"`
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
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auth_strategy", "custom_key", "decoding", "id", "password_key", "tls", "token_key", "type", "uri_key", "username_key"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.AuthStrategy.IsValid() {
		hasInvalidField = true
	} else {
		o.AuthStrategy = *all.AuthStrategy
	}
	o.CustomKey = all.CustomKey
	if !all.Decoding.IsValid() {
		hasInvalidField = true
	} else {
		o.Decoding = *all.Decoding
	}
	o.Id = *all.Id
	o.PasswordKey = all.PasswordKey
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
