// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecretRuleDataAttributesMatchValidation
type SecretRuleDataAttributesMatchValidation struct {
	//
	Endpoint *string `json:"endpoint,omitempty"`
	//
	Hosts []string `json:"hosts,omitempty"`
	//
	HttpMethod *string `json:"http_method,omitempty"`
	//
	InvalidHttpStatusCode []SecretRuleDataAttributesMatchValidationInvalidHttpStatusCodeItems `json:"invalid_http_status_code,omitempty"`
	//
	RequestHeaders map[string]string `json:"request_headers,omitempty"`
	//
	TimeoutSeconds *int64 `json:"timeout_seconds,omitempty"`
	//
	Type *string `json:"type,omitempty"`
	//
	ValidHttpStatusCode []SecretRuleDataAttributesMatchValidationValidHttpStatusCodeItems `json:"valid_http_status_code,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecretRuleDataAttributesMatchValidation instantiates a new SecretRuleDataAttributesMatchValidation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecretRuleDataAttributesMatchValidation() *SecretRuleDataAttributesMatchValidation {
	this := SecretRuleDataAttributesMatchValidation{}
	return &this
}

// NewSecretRuleDataAttributesMatchValidationWithDefaults instantiates a new SecretRuleDataAttributesMatchValidation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecretRuleDataAttributesMatchValidationWithDefaults() *SecretRuleDataAttributesMatchValidation {
	this := SecretRuleDataAttributesMatchValidation{}
	return &this
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *SecretRuleDataAttributesMatchValidation) GetEndpoint() string {
	if o == nil || o.Endpoint == nil {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretRuleDataAttributesMatchValidation) GetEndpointOk() (*string, bool) {
	if o == nil || o.Endpoint == nil {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *SecretRuleDataAttributesMatchValidation) HasEndpoint() bool {
	return o != nil && o.Endpoint != nil
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *SecretRuleDataAttributesMatchValidation) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetHosts returns the Hosts field value if set, zero value otherwise.
func (o *SecretRuleDataAttributesMatchValidation) GetHosts() []string {
	if o == nil || o.Hosts == nil {
		var ret []string
		return ret
	}
	return o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretRuleDataAttributesMatchValidation) GetHostsOk() (*[]string, bool) {
	if o == nil || o.Hosts == nil {
		return nil, false
	}
	return &o.Hosts, true
}

// HasHosts returns a boolean if a field has been set.
func (o *SecretRuleDataAttributesMatchValidation) HasHosts() bool {
	return o != nil && o.Hosts != nil
}

// SetHosts gets a reference to the given []string and assigns it to the Hosts field.
func (o *SecretRuleDataAttributesMatchValidation) SetHosts(v []string) {
	o.Hosts = v
}

// GetHttpMethod returns the HttpMethod field value if set, zero value otherwise.
func (o *SecretRuleDataAttributesMatchValidation) GetHttpMethod() string {
	if o == nil || o.HttpMethod == nil {
		var ret string
		return ret
	}
	return *o.HttpMethod
}

// GetHttpMethodOk returns a tuple with the HttpMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretRuleDataAttributesMatchValidation) GetHttpMethodOk() (*string, bool) {
	if o == nil || o.HttpMethod == nil {
		return nil, false
	}
	return o.HttpMethod, true
}

// HasHttpMethod returns a boolean if a field has been set.
func (o *SecretRuleDataAttributesMatchValidation) HasHttpMethod() bool {
	return o != nil && o.HttpMethod != nil
}

// SetHttpMethod gets a reference to the given string and assigns it to the HttpMethod field.
func (o *SecretRuleDataAttributesMatchValidation) SetHttpMethod(v string) {
	o.HttpMethod = &v
}

// GetInvalidHttpStatusCode returns the InvalidHttpStatusCode field value if set, zero value otherwise.
func (o *SecretRuleDataAttributesMatchValidation) GetInvalidHttpStatusCode() []SecretRuleDataAttributesMatchValidationInvalidHttpStatusCodeItems {
	if o == nil || o.InvalidHttpStatusCode == nil {
		var ret []SecretRuleDataAttributesMatchValidationInvalidHttpStatusCodeItems
		return ret
	}
	return o.InvalidHttpStatusCode
}

// GetInvalidHttpStatusCodeOk returns a tuple with the InvalidHttpStatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretRuleDataAttributesMatchValidation) GetInvalidHttpStatusCodeOk() (*[]SecretRuleDataAttributesMatchValidationInvalidHttpStatusCodeItems, bool) {
	if o == nil || o.InvalidHttpStatusCode == nil {
		return nil, false
	}
	return &o.InvalidHttpStatusCode, true
}

// HasInvalidHttpStatusCode returns a boolean if a field has been set.
func (o *SecretRuleDataAttributesMatchValidation) HasInvalidHttpStatusCode() bool {
	return o != nil && o.InvalidHttpStatusCode != nil
}

// SetInvalidHttpStatusCode gets a reference to the given []SecretRuleDataAttributesMatchValidationInvalidHttpStatusCodeItems and assigns it to the InvalidHttpStatusCode field.
func (o *SecretRuleDataAttributesMatchValidation) SetInvalidHttpStatusCode(v []SecretRuleDataAttributesMatchValidationInvalidHttpStatusCodeItems) {
	o.InvalidHttpStatusCode = v
}

// GetRequestHeaders returns the RequestHeaders field value if set, zero value otherwise.
func (o *SecretRuleDataAttributesMatchValidation) GetRequestHeaders() map[string]string {
	if o == nil || o.RequestHeaders == nil {
		var ret map[string]string
		return ret
	}
	return o.RequestHeaders
}

// GetRequestHeadersOk returns a tuple with the RequestHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretRuleDataAttributesMatchValidation) GetRequestHeadersOk() (*map[string]string, bool) {
	if o == nil || o.RequestHeaders == nil {
		return nil, false
	}
	return &o.RequestHeaders, true
}

// HasRequestHeaders returns a boolean if a field has been set.
func (o *SecretRuleDataAttributesMatchValidation) HasRequestHeaders() bool {
	return o != nil && o.RequestHeaders != nil
}

// SetRequestHeaders gets a reference to the given map[string]string and assigns it to the RequestHeaders field.
func (o *SecretRuleDataAttributesMatchValidation) SetRequestHeaders(v map[string]string) {
	o.RequestHeaders = v
}

// GetTimeoutSeconds returns the TimeoutSeconds field value if set, zero value otherwise.
func (o *SecretRuleDataAttributesMatchValidation) GetTimeoutSeconds() int64 {
	if o == nil || o.TimeoutSeconds == nil {
		var ret int64
		return ret
	}
	return *o.TimeoutSeconds
}

// GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretRuleDataAttributesMatchValidation) GetTimeoutSecondsOk() (*int64, bool) {
	if o == nil || o.TimeoutSeconds == nil {
		return nil, false
	}
	return o.TimeoutSeconds, true
}

// HasTimeoutSeconds returns a boolean if a field has been set.
func (o *SecretRuleDataAttributesMatchValidation) HasTimeoutSeconds() bool {
	return o != nil && o.TimeoutSeconds != nil
}

// SetTimeoutSeconds gets a reference to the given int64 and assigns it to the TimeoutSeconds field.
func (o *SecretRuleDataAttributesMatchValidation) SetTimeoutSeconds(v int64) {
	o.TimeoutSeconds = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SecretRuleDataAttributesMatchValidation) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretRuleDataAttributesMatchValidation) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SecretRuleDataAttributesMatchValidation) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SecretRuleDataAttributesMatchValidation) SetType(v string) {
	o.Type = &v
}

// GetValidHttpStatusCode returns the ValidHttpStatusCode field value if set, zero value otherwise.
func (o *SecretRuleDataAttributesMatchValidation) GetValidHttpStatusCode() []SecretRuleDataAttributesMatchValidationValidHttpStatusCodeItems {
	if o == nil || o.ValidHttpStatusCode == nil {
		var ret []SecretRuleDataAttributesMatchValidationValidHttpStatusCodeItems
		return ret
	}
	return o.ValidHttpStatusCode
}

// GetValidHttpStatusCodeOk returns a tuple with the ValidHttpStatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretRuleDataAttributesMatchValidation) GetValidHttpStatusCodeOk() (*[]SecretRuleDataAttributesMatchValidationValidHttpStatusCodeItems, bool) {
	if o == nil || o.ValidHttpStatusCode == nil {
		return nil, false
	}
	return &o.ValidHttpStatusCode, true
}

// HasValidHttpStatusCode returns a boolean if a field has been set.
func (o *SecretRuleDataAttributesMatchValidation) HasValidHttpStatusCode() bool {
	return o != nil && o.ValidHttpStatusCode != nil
}

// SetValidHttpStatusCode gets a reference to the given []SecretRuleDataAttributesMatchValidationValidHttpStatusCodeItems and assigns it to the ValidHttpStatusCode field.
func (o *SecretRuleDataAttributesMatchValidation) SetValidHttpStatusCode(v []SecretRuleDataAttributesMatchValidationValidHttpStatusCodeItems) {
	o.ValidHttpStatusCode = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecretRuleDataAttributesMatchValidation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Endpoint != nil {
		toSerialize["endpoint"] = o.Endpoint
	}
	if o.Hosts != nil {
		toSerialize["hosts"] = o.Hosts
	}
	if o.HttpMethod != nil {
		toSerialize["http_method"] = o.HttpMethod
	}
	if o.InvalidHttpStatusCode != nil {
		toSerialize["invalid_http_status_code"] = o.InvalidHttpStatusCode
	}
	if o.RequestHeaders != nil {
		toSerialize["request_headers"] = o.RequestHeaders
	}
	if o.TimeoutSeconds != nil {
		toSerialize["timeout_seconds"] = o.TimeoutSeconds
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.ValidHttpStatusCode != nil {
		toSerialize["valid_http_status_code"] = o.ValidHttpStatusCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecretRuleDataAttributesMatchValidation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Endpoint              *string                                                             `json:"endpoint,omitempty"`
		Hosts                 []string                                                            `json:"hosts,omitempty"`
		HttpMethod            *string                                                             `json:"http_method,omitempty"`
		InvalidHttpStatusCode []SecretRuleDataAttributesMatchValidationInvalidHttpStatusCodeItems `json:"invalid_http_status_code,omitempty"`
		RequestHeaders        map[string]string                                                   `json:"request_headers,omitempty"`
		TimeoutSeconds        *int64                                                              `json:"timeout_seconds,omitempty"`
		Type                  *string                                                             `json:"type,omitempty"`
		ValidHttpStatusCode   []SecretRuleDataAttributesMatchValidationValidHttpStatusCodeItems   `json:"valid_http_status_code,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"endpoint", "hosts", "http_method", "invalid_http_status_code", "request_headers", "timeout_seconds", "type", "valid_http_status_code"})
	} else {
		return err
	}
	o.Endpoint = all.Endpoint
	o.Hosts = all.Hosts
	o.HttpMethod = all.HttpMethod
	o.InvalidHttpStatusCode = all.InvalidHttpStatusCode
	o.RequestHeaders = all.RequestHeaders
	o.TimeoutSeconds = all.TimeoutSeconds
	o.Type = all.Type
	o.ValidHttpStatusCode = all.ValidHttpStatusCode

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
