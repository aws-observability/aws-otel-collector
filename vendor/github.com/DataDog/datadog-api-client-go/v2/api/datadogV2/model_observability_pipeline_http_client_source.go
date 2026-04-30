// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineHttpClientSource The `http_client` source scrapes logs from HTTP endpoints at regular intervals.
//
// **Supported pipeline types:** logs
type ObservabilityPipelineHttpClientSource struct {
	// Optional authentication strategy for HTTP requests.
	AuthStrategy *ObservabilityPipelineHttpClientSourceAuthStrategy `json:"auth_strategy,omitempty"`
	// Name of the environment variable or secret that holds a custom header value (used with custom auth strategies).
	CustomKey *string `json:"custom_key,omitempty"`
	// The decoding format used to interpret incoming logs.
	Decoding ObservabilityPipelineDecoding `json:"decoding"`
	// Name of the environment variable or secret that holds the HTTP endpoint URL to scrape.
	EndpointUrlKey *string `json:"endpoint_url_key,omitempty"`
	// The unique identifier for this component. Used in other parts of the pipeline to reference this component (for example, as the `input` to downstream components).
	Id string `json:"id"`
	// Name of the environment variable or secret that holds the password (used when `auth_strategy` is `basic`).
	PasswordKey *string `json:"password_key,omitempty"`
	// The interval (in seconds) between HTTP scrape requests.
	ScrapeIntervalSecs *int64 `json:"scrape_interval_secs,omitempty"`
	// The timeout (in seconds) for each scrape request.
	ScrapeTimeoutSecs *int64 `json:"scrape_timeout_secs,omitempty"`
	// Configuration for enabling TLS encryption between the pipeline component and external services.
	Tls *ObservabilityPipelineTls `json:"tls,omitempty"`
	// Name of the environment variable or secret that holds the bearer token (used when `auth_strategy` is `bearer`).
	TokenKey *string `json:"token_key,omitempty"`
	// The source type. The value should always be `http_client`.
	Type ObservabilityPipelineHttpClientSourceType `json:"type"`
	// Name of the environment variable or secret that holds the username (used when `auth_strategy` is `basic`).
	UsernameKey *string `json:"username_key,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineHttpClientSource instantiates a new ObservabilityPipelineHttpClientSource object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineHttpClientSource(decoding ObservabilityPipelineDecoding, id string, typeVar ObservabilityPipelineHttpClientSourceType) *ObservabilityPipelineHttpClientSource {
	this := ObservabilityPipelineHttpClientSource{}
	this.Decoding = decoding
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineHttpClientSourceWithDefaults instantiates a new ObservabilityPipelineHttpClientSource object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineHttpClientSourceWithDefaults() *ObservabilityPipelineHttpClientSource {
	this := ObservabilityPipelineHttpClientSource{}
	var typeVar ObservabilityPipelineHttpClientSourceType = OBSERVABILITYPIPELINEHTTPCLIENTSOURCETYPE_HTTP_CLIENT
	this.Type = typeVar
	return &this
}

// GetAuthStrategy returns the AuthStrategy field value if set, zero value otherwise.
func (o *ObservabilityPipelineHttpClientSource) GetAuthStrategy() ObservabilityPipelineHttpClientSourceAuthStrategy {
	if o == nil || o.AuthStrategy == nil {
		var ret ObservabilityPipelineHttpClientSourceAuthStrategy
		return ret
	}
	return *o.AuthStrategy
}

// GetAuthStrategyOk returns a tuple with the AuthStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineHttpClientSource) GetAuthStrategyOk() (*ObservabilityPipelineHttpClientSourceAuthStrategy, bool) {
	if o == nil || o.AuthStrategy == nil {
		return nil, false
	}
	return o.AuthStrategy, true
}

// HasAuthStrategy returns a boolean if a field has been set.
func (o *ObservabilityPipelineHttpClientSource) HasAuthStrategy() bool {
	return o != nil && o.AuthStrategy != nil
}

// SetAuthStrategy gets a reference to the given ObservabilityPipelineHttpClientSourceAuthStrategy and assigns it to the AuthStrategy field.
func (o *ObservabilityPipelineHttpClientSource) SetAuthStrategy(v ObservabilityPipelineHttpClientSourceAuthStrategy) {
	o.AuthStrategy = &v
}

// GetCustomKey returns the CustomKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineHttpClientSource) GetCustomKey() string {
	if o == nil || o.CustomKey == nil {
		var ret string
		return ret
	}
	return *o.CustomKey
}

// GetCustomKeyOk returns a tuple with the CustomKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineHttpClientSource) GetCustomKeyOk() (*string, bool) {
	if o == nil || o.CustomKey == nil {
		return nil, false
	}
	return o.CustomKey, true
}

// HasCustomKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineHttpClientSource) HasCustomKey() bool {
	return o != nil && o.CustomKey != nil
}

// SetCustomKey gets a reference to the given string and assigns it to the CustomKey field.
func (o *ObservabilityPipelineHttpClientSource) SetCustomKey(v string) {
	o.CustomKey = &v
}

// GetDecoding returns the Decoding field value.
func (o *ObservabilityPipelineHttpClientSource) GetDecoding() ObservabilityPipelineDecoding {
	if o == nil {
		var ret ObservabilityPipelineDecoding
		return ret
	}
	return o.Decoding
}

// GetDecodingOk returns a tuple with the Decoding field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineHttpClientSource) GetDecodingOk() (*ObservabilityPipelineDecoding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Decoding, true
}

// SetDecoding sets field value.
func (o *ObservabilityPipelineHttpClientSource) SetDecoding(v ObservabilityPipelineDecoding) {
	o.Decoding = v
}

// GetEndpointUrlKey returns the EndpointUrlKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineHttpClientSource) GetEndpointUrlKey() string {
	if o == nil || o.EndpointUrlKey == nil {
		var ret string
		return ret
	}
	return *o.EndpointUrlKey
}

// GetEndpointUrlKeyOk returns a tuple with the EndpointUrlKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineHttpClientSource) GetEndpointUrlKeyOk() (*string, bool) {
	if o == nil || o.EndpointUrlKey == nil {
		return nil, false
	}
	return o.EndpointUrlKey, true
}

// HasEndpointUrlKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineHttpClientSource) HasEndpointUrlKey() bool {
	return o != nil && o.EndpointUrlKey != nil
}

// SetEndpointUrlKey gets a reference to the given string and assigns it to the EndpointUrlKey field.
func (o *ObservabilityPipelineHttpClientSource) SetEndpointUrlKey(v string) {
	o.EndpointUrlKey = &v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineHttpClientSource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineHttpClientSource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineHttpClientSource) SetId(v string) {
	o.Id = v
}

// GetPasswordKey returns the PasswordKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineHttpClientSource) GetPasswordKey() string {
	if o == nil || o.PasswordKey == nil {
		var ret string
		return ret
	}
	return *o.PasswordKey
}

// GetPasswordKeyOk returns a tuple with the PasswordKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineHttpClientSource) GetPasswordKeyOk() (*string, bool) {
	if o == nil || o.PasswordKey == nil {
		return nil, false
	}
	return o.PasswordKey, true
}

// HasPasswordKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineHttpClientSource) HasPasswordKey() bool {
	return o != nil && o.PasswordKey != nil
}

// SetPasswordKey gets a reference to the given string and assigns it to the PasswordKey field.
func (o *ObservabilityPipelineHttpClientSource) SetPasswordKey(v string) {
	o.PasswordKey = &v
}

// GetScrapeIntervalSecs returns the ScrapeIntervalSecs field value if set, zero value otherwise.
func (o *ObservabilityPipelineHttpClientSource) GetScrapeIntervalSecs() int64 {
	if o == nil || o.ScrapeIntervalSecs == nil {
		var ret int64
		return ret
	}
	return *o.ScrapeIntervalSecs
}

// GetScrapeIntervalSecsOk returns a tuple with the ScrapeIntervalSecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineHttpClientSource) GetScrapeIntervalSecsOk() (*int64, bool) {
	if o == nil || o.ScrapeIntervalSecs == nil {
		return nil, false
	}
	return o.ScrapeIntervalSecs, true
}

// HasScrapeIntervalSecs returns a boolean if a field has been set.
func (o *ObservabilityPipelineHttpClientSource) HasScrapeIntervalSecs() bool {
	return o != nil && o.ScrapeIntervalSecs != nil
}

// SetScrapeIntervalSecs gets a reference to the given int64 and assigns it to the ScrapeIntervalSecs field.
func (o *ObservabilityPipelineHttpClientSource) SetScrapeIntervalSecs(v int64) {
	o.ScrapeIntervalSecs = &v
}

// GetScrapeTimeoutSecs returns the ScrapeTimeoutSecs field value if set, zero value otherwise.
func (o *ObservabilityPipelineHttpClientSource) GetScrapeTimeoutSecs() int64 {
	if o == nil || o.ScrapeTimeoutSecs == nil {
		var ret int64
		return ret
	}
	return *o.ScrapeTimeoutSecs
}

// GetScrapeTimeoutSecsOk returns a tuple with the ScrapeTimeoutSecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineHttpClientSource) GetScrapeTimeoutSecsOk() (*int64, bool) {
	if o == nil || o.ScrapeTimeoutSecs == nil {
		return nil, false
	}
	return o.ScrapeTimeoutSecs, true
}

// HasScrapeTimeoutSecs returns a boolean if a field has been set.
func (o *ObservabilityPipelineHttpClientSource) HasScrapeTimeoutSecs() bool {
	return o != nil && o.ScrapeTimeoutSecs != nil
}

// SetScrapeTimeoutSecs gets a reference to the given int64 and assigns it to the ScrapeTimeoutSecs field.
func (o *ObservabilityPipelineHttpClientSource) SetScrapeTimeoutSecs(v int64) {
	o.ScrapeTimeoutSecs = &v
}

// GetTls returns the Tls field value if set, zero value otherwise.
func (o *ObservabilityPipelineHttpClientSource) GetTls() ObservabilityPipelineTls {
	if o == nil || o.Tls == nil {
		var ret ObservabilityPipelineTls
		return ret
	}
	return *o.Tls
}

// GetTlsOk returns a tuple with the Tls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineHttpClientSource) GetTlsOk() (*ObservabilityPipelineTls, bool) {
	if o == nil || o.Tls == nil {
		return nil, false
	}
	return o.Tls, true
}

// HasTls returns a boolean if a field has been set.
func (o *ObservabilityPipelineHttpClientSource) HasTls() bool {
	return o != nil && o.Tls != nil
}

// SetTls gets a reference to the given ObservabilityPipelineTls and assigns it to the Tls field.
func (o *ObservabilityPipelineHttpClientSource) SetTls(v ObservabilityPipelineTls) {
	o.Tls = &v
}

// GetTokenKey returns the TokenKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineHttpClientSource) GetTokenKey() string {
	if o == nil || o.TokenKey == nil {
		var ret string
		return ret
	}
	return *o.TokenKey
}

// GetTokenKeyOk returns a tuple with the TokenKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineHttpClientSource) GetTokenKeyOk() (*string, bool) {
	if o == nil || o.TokenKey == nil {
		return nil, false
	}
	return o.TokenKey, true
}

// HasTokenKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineHttpClientSource) HasTokenKey() bool {
	return o != nil && o.TokenKey != nil
}

// SetTokenKey gets a reference to the given string and assigns it to the TokenKey field.
func (o *ObservabilityPipelineHttpClientSource) SetTokenKey(v string) {
	o.TokenKey = &v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineHttpClientSource) GetType() ObservabilityPipelineHttpClientSourceType {
	if o == nil {
		var ret ObservabilityPipelineHttpClientSourceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineHttpClientSource) GetTypeOk() (*ObservabilityPipelineHttpClientSourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineHttpClientSource) SetType(v ObservabilityPipelineHttpClientSourceType) {
	o.Type = v
}

// GetUsernameKey returns the UsernameKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineHttpClientSource) GetUsernameKey() string {
	if o == nil || o.UsernameKey == nil {
		var ret string
		return ret
	}
	return *o.UsernameKey
}

// GetUsernameKeyOk returns a tuple with the UsernameKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineHttpClientSource) GetUsernameKeyOk() (*string, bool) {
	if o == nil || o.UsernameKey == nil {
		return nil, false
	}
	return o.UsernameKey, true
}

// HasUsernameKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineHttpClientSource) HasUsernameKey() bool {
	return o != nil && o.UsernameKey != nil
}

// SetUsernameKey gets a reference to the given string and assigns it to the UsernameKey field.
func (o *ObservabilityPipelineHttpClientSource) SetUsernameKey(v string) {
	o.UsernameKey = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineHttpClientSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AuthStrategy != nil {
		toSerialize["auth_strategy"] = o.AuthStrategy
	}
	if o.CustomKey != nil {
		toSerialize["custom_key"] = o.CustomKey
	}
	toSerialize["decoding"] = o.Decoding
	if o.EndpointUrlKey != nil {
		toSerialize["endpoint_url_key"] = o.EndpointUrlKey
	}
	toSerialize["id"] = o.Id
	if o.PasswordKey != nil {
		toSerialize["password_key"] = o.PasswordKey
	}
	if o.ScrapeIntervalSecs != nil {
		toSerialize["scrape_interval_secs"] = o.ScrapeIntervalSecs
	}
	if o.ScrapeTimeoutSecs != nil {
		toSerialize["scrape_timeout_secs"] = o.ScrapeTimeoutSecs
	}
	if o.Tls != nil {
		toSerialize["tls"] = o.Tls
	}
	if o.TokenKey != nil {
		toSerialize["token_key"] = o.TokenKey
	}
	toSerialize["type"] = o.Type
	if o.UsernameKey != nil {
		toSerialize["username_key"] = o.UsernameKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineHttpClientSource) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthStrategy       *ObservabilityPipelineHttpClientSourceAuthStrategy `json:"auth_strategy,omitempty"`
		CustomKey          *string                                            `json:"custom_key,omitempty"`
		Decoding           *ObservabilityPipelineDecoding                     `json:"decoding"`
		EndpointUrlKey     *string                                            `json:"endpoint_url_key,omitempty"`
		Id                 *string                                            `json:"id"`
		PasswordKey        *string                                            `json:"password_key,omitempty"`
		ScrapeIntervalSecs *int64                                             `json:"scrape_interval_secs,omitempty"`
		ScrapeTimeoutSecs  *int64                                             `json:"scrape_timeout_secs,omitempty"`
		Tls                *ObservabilityPipelineTls                          `json:"tls,omitempty"`
		TokenKey           *string                                            `json:"token_key,omitempty"`
		Type               *ObservabilityPipelineHttpClientSourceType         `json:"type"`
		UsernameKey        *string                                            `json:"username_key,omitempty"`
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
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auth_strategy", "custom_key", "decoding", "endpoint_url_key", "id", "password_key", "scrape_interval_secs", "scrape_timeout_secs", "tls", "token_key", "type", "username_key"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AuthStrategy != nil && !all.AuthStrategy.IsValid() {
		hasInvalidField = true
	} else {
		o.AuthStrategy = all.AuthStrategy
	}
	o.CustomKey = all.CustomKey
	if !all.Decoding.IsValid() {
		hasInvalidField = true
	} else {
		o.Decoding = *all.Decoding
	}
	o.EndpointUrlKey = all.EndpointUrlKey
	o.Id = *all.Id
	o.PasswordKey = all.PasswordKey
	o.ScrapeIntervalSecs = all.ScrapeIntervalSecs
	o.ScrapeTimeoutSecs = all.ScrapeTimeoutSecs
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
	o.UsernameKey = all.UsernameKey

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
