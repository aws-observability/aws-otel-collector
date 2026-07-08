// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WebhooksOAuth2ClientCredentialsCreateData OAuth2 client credentials data for a create request.
type WebhooksOAuth2ClientCredentialsCreateData struct {
	// OAuth2 client credentials attributes for a create request.
	Attributes WebhooksOAuth2ClientCredentialsCreateAttributes `json:"attributes"`
	// OAuth2 client credentials resource type.
	Type WebhooksOAuth2ClientCredentialsType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWebhooksOAuth2ClientCredentialsCreateData instantiates a new WebhooksOAuth2ClientCredentialsCreateData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWebhooksOAuth2ClientCredentialsCreateData(attributes WebhooksOAuth2ClientCredentialsCreateAttributes, typeVar WebhooksOAuth2ClientCredentialsType) *WebhooksOAuth2ClientCredentialsCreateData {
	this := WebhooksOAuth2ClientCredentialsCreateData{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewWebhooksOAuth2ClientCredentialsCreateDataWithDefaults instantiates a new WebhooksOAuth2ClientCredentialsCreateData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWebhooksOAuth2ClientCredentialsCreateDataWithDefaults() *WebhooksOAuth2ClientCredentialsCreateData {
	this := WebhooksOAuth2ClientCredentialsCreateData{}
	var typeVar WebhooksOAuth2ClientCredentialsType = WEBHOOKSOAUTH2CLIENTCREDENTIALSTYPE_WEBHOOKS_AUTH_METHOD_OAUTH2_CLIENT_CREDENTIALS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *WebhooksOAuth2ClientCredentialsCreateData) GetAttributes() WebhooksOAuth2ClientCredentialsCreateAttributes {
	if o == nil {
		var ret WebhooksOAuth2ClientCredentialsCreateAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *WebhooksOAuth2ClientCredentialsCreateData) GetAttributesOk() (*WebhooksOAuth2ClientCredentialsCreateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *WebhooksOAuth2ClientCredentialsCreateData) SetAttributes(v WebhooksOAuth2ClientCredentialsCreateAttributes) {
	o.Attributes = v
}

// GetType returns the Type field value.
func (o *WebhooksOAuth2ClientCredentialsCreateData) GetType() WebhooksOAuth2ClientCredentialsType {
	if o == nil {
		var ret WebhooksOAuth2ClientCredentialsType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WebhooksOAuth2ClientCredentialsCreateData) GetTypeOk() (*WebhooksOAuth2ClientCredentialsType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *WebhooksOAuth2ClientCredentialsCreateData) SetType(v WebhooksOAuth2ClientCredentialsType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o WebhooksOAuth2ClientCredentialsCreateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WebhooksOAuth2ClientCredentialsCreateData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *WebhooksOAuth2ClientCredentialsCreateAttributes `json:"attributes"`
		Type       *WebhooksOAuth2ClientCredentialsType             `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
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
