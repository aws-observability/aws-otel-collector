// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HTTPTokenAuthUpdate The definition of `HTTPTokenAuthUpdate` object.
type HTTPTokenAuthUpdate struct {
	// The definition of `HTTPBody` object.
	Body *HTTPBody `json:"body,omitempty"`
	// The `HTTPTokenAuthUpdate` `headers`.
	Headers []HTTPHeaderUpdate `json:"headers,omitempty"`
	// The `HTTPTokenAuthUpdate` `tokens`.
	Tokens []HTTPTokenUpdate `json:"tokens,omitempty"`
	// The definition of `HTTPTokenAuthType` object.
	Type HTTPTokenAuthType `json:"type"`
	// The `HTTPTokenAuthUpdate` `url_parameters`.
	UrlParameters []UrlParamUpdate `json:"url_parameters,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewHTTPTokenAuthUpdate instantiates a new HTTPTokenAuthUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHTTPTokenAuthUpdate(typeVar HTTPTokenAuthType) *HTTPTokenAuthUpdate {
	this := HTTPTokenAuthUpdate{}
	this.Type = typeVar
	return &this
}

// NewHTTPTokenAuthUpdateWithDefaults instantiates a new HTTPTokenAuthUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHTTPTokenAuthUpdateWithDefaults() *HTTPTokenAuthUpdate {
	this := HTTPTokenAuthUpdate{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *HTTPTokenAuthUpdate) GetBody() HTTPBody {
	if o == nil || o.Body == nil {
		var ret HTTPBody
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTPTokenAuthUpdate) GetBodyOk() (*HTTPBody, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *HTTPTokenAuthUpdate) HasBody() bool {
	return o != nil && o.Body != nil
}

// SetBody gets a reference to the given HTTPBody and assigns it to the Body field.
func (o *HTTPTokenAuthUpdate) SetBody(v HTTPBody) {
	o.Body = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *HTTPTokenAuthUpdate) GetHeaders() []HTTPHeaderUpdate {
	if o == nil || o.Headers == nil {
		var ret []HTTPHeaderUpdate
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTPTokenAuthUpdate) GetHeadersOk() (*[]HTTPHeaderUpdate, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return &o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *HTTPTokenAuthUpdate) HasHeaders() bool {
	return o != nil && o.Headers != nil
}

// SetHeaders gets a reference to the given []HTTPHeaderUpdate and assigns it to the Headers field.
func (o *HTTPTokenAuthUpdate) SetHeaders(v []HTTPHeaderUpdate) {
	o.Headers = v
}

// GetTokens returns the Tokens field value if set, zero value otherwise.
func (o *HTTPTokenAuthUpdate) GetTokens() []HTTPTokenUpdate {
	if o == nil || o.Tokens == nil {
		var ret []HTTPTokenUpdate
		return ret
	}
	return o.Tokens
}

// GetTokensOk returns a tuple with the Tokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTPTokenAuthUpdate) GetTokensOk() (*[]HTTPTokenUpdate, bool) {
	if o == nil || o.Tokens == nil {
		return nil, false
	}
	return &o.Tokens, true
}

// HasTokens returns a boolean if a field has been set.
func (o *HTTPTokenAuthUpdate) HasTokens() bool {
	return o != nil && o.Tokens != nil
}

// SetTokens gets a reference to the given []HTTPTokenUpdate and assigns it to the Tokens field.
func (o *HTTPTokenAuthUpdate) SetTokens(v []HTTPTokenUpdate) {
	o.Tokens = v
}

// GetType returns the Type field value.
func (o *HTTPTokenAuthUpdate) GetType() HTTPTokenAuthType {
	if o == nil {
		var ret HTTPTokenAuthType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *HTTPTokenAuthUpdate) GetTypeOk() (*HTTPTokenAuthType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *HTTPTokenAuthUpdate) SetType(v HTTPTokenAuthType) {
	o.Type = v
}

// GetUrlParameters returns the UrlParameters field value if set, zero value otherwise.
func (o *HTTPTokenAuthUpdate) GetUrlParameters() []UrlParamUpdate {
	if o == nil || o.UrlParameters == nil {
		var ret []UrlParamUpdate
		return ret
	}
	return o.UrlParameters
}

// GetUrlParametersOk returns a tuple with the UrlParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTPTokenAuthUpdate) GetUrlParametersOk() (*[]UrlParamUpdate, bool) {
	if o == nil || o.UrlParameters == nil {
		return nil, false
	}
	return &o.UrlParameters, true
}

// HasUrlParameters returns a boolean if a field has been set.
func (o *HTTPTokenAuthUpdate) HasUrlParameters() bool {
	return o != nil && o.UrlParameters != nil
}

// SetUrlParameters gets a reference to the given []UrlParamUpdate and assigns it to the UrlParameters field.
func (o *HTTPTokenAuthUpdate) SetUrlParameters(v []UrlParamUpdate) {
	o.UrlParameters = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HTTPTokenAuthUpdate) MarshalJSON() ([]byte, error) {
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
func (o *HTTPTokenAuthUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Body          *HTTPBody          `json:"body,omitempty"`
		Headers       []HTTPHeaderUpdate `json:"headers,omitempty"`
		Tokens        []HTTPTokenUpdate  `json:"tokens,omitempty"`
		Type          *HTTPTokenAuthType `json:"type"`
		UrlParameters []UrlParamUpdate   `json:"url_parameters,omitempty"`
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
