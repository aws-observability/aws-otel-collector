// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WebhooksOAuth2ClientCredentialsResponseData OAuth2 client credentials data from a response.
type WebhooksOAuth2ClientCredentialsResponseData struct {
	// OAuth2 client credentials attributes returned by the API. The `client_secret` is never echoed.
	Attributes WebhooksOAuth2ClientCredentialsResponseAttributes `json:"attributes"`
	// The ID of the OAuth2 client credentials auth method.
	Id string `json:"id"`
	// OAuth2 client credentials resource type.
	Type WebhooksOAuth2ClientCredentialsType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWebhooksOAuth2ClientCredentialsResponseData instantiates a new WebhooksOAuth2ClientCredentialsResponseData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWebhooksOAuth2ClientCredentialsResponseData(attributes WebhooksOAuth2ClientCredentialsResponseAttributes, id string, typeVar WebhooksOAuth2ClientCredentialsType) *WebhooksOAuth2ClientCredentialsResponseData {
	this := WebhooksOAuth2ClientCredentialsResponseData{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewWebhooksOAuth2ClientCredentialsResponseDataWithDefaults instantiates a new WebhooksOAuth2ClientCredentialsResponseData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWebhooksOAuth2ClientCredentialsResponseDataWithDefaults() *WebhooksOAuth2ClientCredentialsResponseData {
	this := WebhooksOAuth2ClientCredentialsResponseData{}
	var typeVar WebhooksOAuth2ClientCredentialsType = WEBHOOKSOAUTH2CLIENTCREDENTIALSTYPE_WEBHOOKS_AUTH_METHOD_OAUTH2_CLIENT_CREDENTIALS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *WebhooksOAuth2ClientCredentialsResponseData) GetAttributes() WebhooksOAuth2ClientCredentialsResponseAttributes {
	if o == nil {
		var ret WebhooksOAuth2ClientCredentialsResponseAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *WebhooksOAuth2ClientCredentialsResponseData) GetAttributesOk() (*WebhooksOAuth2ClientCredentialsResponseAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *WebhooksOAuth2ClientCredentialsResponseData) SetAttributes(v WebhooksOAuth2ClientCredentialsResponseAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *WebhooksOAuth2ClientCredentialsResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WebhooksOAuth2ClientCredentialsResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *WebhooksOAuth2ClientCredentialsResponseData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *WebhooksOAuth2ClientCredentialsResponseData) GetType() WebhooksOAuth2ClientCredentialsType {
	if o == nil {
		var ret WebhooksOAuth2ClientCredentialsType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WebhooksOAuth2ClientCredentialsResponseData) GetTypeOk() (*WebhooksOAuth2ClientCredentialsType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *WebhooksOAuth2ClientCredentialsResponseData) SetType(v WebhooksOAuth2ClientCredentialsType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o WebhooksOAuth2ClientCredentialsResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WebhooksOAuth2ClientCredentialsResponseData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *WebhooksOAuth2ClientCredentialsResponseAttributes `json:"attributes"`
		Id         *string                                            `json:"id"`
		Type       *WebhooksOAuth2ClientCredentialsType               `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	o.Id = *all.Id
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
