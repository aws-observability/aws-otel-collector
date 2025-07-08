// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HTTPTokenAuth The definition of `HTTPTokenAuth` object.
type HTTPTokenAuth struct {
	// The definition of `HTTPBody` object.
	Body *HTTPBody `json:"body,omitempty"`
	// The `HTTPTokenAuth` `headers`.
	Headers []HTTPHeader `json:"headers,omitempty"`
	// The `HTTPTokenAuth` `tokens`.
	Tokens []HTTPToken `json:"tokens,omitempty"`
	// The definition of `HTTPTokenAuthType` object.
	Type HTTPTokenAuthType `json:"type"`
	// The `HTTPTokenAuth` `url_parameters`.
	UrlParameters []UrlParam `json:"url_parameters,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewHTTPTokenAuth instantiates a new HTTPTokenAuth object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHTTPTokenAuth(typeVar HTTPTokenAuthType) *HTTPTokenAuth {
	this := HTTPTokenAuth{}
	this.Type = typeVar
	return &this
}

// NewHTTPTokenAuthWithDefaults instantiates a new HTTPTokenAuth object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHTTPTokenAuthWithDefaults() *HTTPTokenAuth {
	this := HTTPTokenAuth{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *HTTPTokenAuth) GetBody() HTTPBody {
	if o == nil || o.Body == nil {
		var ret HTTPBody
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTPTokenAuth) GetBodyOk() (*HTTPBody, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *HTTPTokenAuth) HasBody() bool {
	return o != nil && o.Body != nil
}

// SetBody gets a reference to the given HTTPBody and assigns it to the Body field.
func (o *HTTPTokenAuth) SetBody(v HTTPBody) {
	o.Body = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *HTTPTokenAuth) GetHeaders() []HTTPHeader {
	if o == nil || o.Headers == nil {
		var ret []HTTPHeader
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTPTokenAuth) GetHeadersOk() (*[]HTTPHeader, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return &o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *HTTPTokenAuth) HasHeaders() bool {
	return o != nil && o.Headers != nil
}

// SetHeaders gets a reference to the given []HTTPHeader and assigns it to the Headers field.
func (o *HTTPTokenAuth) SetHeaders(v []HTTPHeader) {
	o.Headers = v
}

// GetTokens returns the Tokens field value if set, zero value otherwise.
func (o *HTTPTokenAuth) GetTokens() []HTTPToken {
	if o == nil || o.Tokens == nil {
		var ret []HTTPToken
		return ret
	}
	return o.Tokens
}

// GetTokensOk returns a tuple with the Tokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTPTokenAuth) GetTokensOk() (*[]HTTPToken, bool) {
	if o == nil || o.Tokens == nil {
		return nil, false
	}
	return &o.Tokens, true
}

// HasTokens returns a boolean if a field has been set.
func (o *HTTPTokenAuth) HasTokens() bool {
	return o != nil && o.Tokens != nil
}

// SetTokens gets a reference to the given []HTTPToken and assigns it to the Tokens field.
func (o *HTTPTokenAuth) SetTokens(v []HTTPToken) {
	o.Tokens = v
}

// GetType returns the Type field value.
func (o *HTTPTokenAuth) GetType() HTTPTokenAuthType {
	if o == nil {
		var ret HTTPTokenAuthType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *HTTPTokenAuth) GetTypeOk() (*HTTPTokenAuthType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *HTTPTokenAuth) SetType(v HTTPTokenAuthType) {
	o.Type = v
}

// GetUrlParameters returns the UrlParameters field value if set, zero value otherwise.
func (o *HTTPTokenAuth) GetUrlParameters() []UrlParam {
	if o == nil || o.UrlParameters == nil {
		var ret []UrlParam
		return ret
	}
	return o.UrlParameters
}

// GetUrlParametersOk returns a tuple with the UrlParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTPTokenAuth) GetUrlParametersOk() (*[]UrlParam, bool) {
	if o == nil || o.UrlParameters == nil {
		return nil, false
	}
	return &o.UrlParameters, true
}

// HasUrlParameters returns a boolean if a field has been set.
func (o *HTTPTokenAuth) HasUrlParameters() bool {
	return o != nil && o.UrlParameters != nil
}

// SetUrlParameters gets a reference to the given []UrlParam and assigns it to the UrlParameters field.
func (o *HTTPTokenAuth) SetUrlParameters(v []UrlParam) {
	o.UrlParameters = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HTTPTokenAuth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	if o.Tokens != nil {
		toSerialize["tokens"] = o.Tokens
	}
	toSerialize["type"] = o.Type
	if o.UrlParameters != nil {
		toSerialize["url_parameters"] = o.UrlParameters
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HTTPTokenAuth) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Body          *HTTPBody          `json:"body,omitempty"`
		Headers       []HTTPHeader       `json:"headers,omitempty"`
		Tokens        []HTTPToken        `json:"tokens,omitempty"`
		Type          *HTTPTokenAuthType `json:"type"`
		UrlParameters []UrlParam         `json:"url_parameters,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"body", "headers", "tokens", "type", "url_parameters"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Body != nil && all.Body.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Body = all.Body
	o.Headers = all.Headers
	o.Tokens = all.Tokens
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.UrlParameters = all.UrlParameters

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
