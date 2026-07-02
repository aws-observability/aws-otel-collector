// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WebhooksOAuth2ClientCredentialsUpdateData OAuth2 client credentials data for an update request.
type WebhooksOAuth2ClientCredentialsUpdateData struct {
	// OAuth2 client credentials attributes for an update request.
	Attributes WebhooksOAuth2ClientCredentialsUpdateAttributes `json:"attributes"`
	// OAuth2 client credentials resource type.
	Type WebhooksOAuth2ClientCredentialsType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWebhooksOAuth2ClientCredentialsUpdateData instantiates a new WebhooksOAuth2ClientCredentialsUpdateData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWebhooksOAuth2ClientCredentialsUpdateData(attributes WebhooksOAuth2ClientCredentialsUpdateAttributes, typeVar WebhooksOAuth2ClientCredentialsType) *WebhooksOAuth2ClientCredentialsUpdateData {
	this := WebhooksOAuth2ClientCredentialsUpdateData{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewWebhooksOAuth2ClientCredentialsUpdateDataWithDefaults instantiates a new WebhooksOAuth2ClientCredentialsUpdateData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWebhooksOAuth2ClientCredentialsUpdateDataWithDefaults() *WebhooksOAuth2ClientCredentialsUpdateData {
	this := WebhooksOAuth2ClientCredentialsUpdateData{}
	var typeVar WebhooksOAuth2ClientCredentialsType = WEBHOOKSOAUTH2CLIENTCREDENTIALSTYPE_WEBHOOKS_AUTH_METHOD_OAUTH2_CLIENT_CREDENTIALS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *WebhooksOAuth2ClientCredentialsUpdateData) GetAttributes() WebhooksOAuth2ClientCredentialsUpdateAttributes {
	if o == nil {
		var ret WebhooksOAuth2ClientCredentialsUpdateAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *WebhooksOAuth2ClientCredentialsUpdateData) GetAttributesOk() (*WebhooksOAuth2ClientCredentialsUpdateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *WebhooksOAuth2ClientCredentialsUpdateData) SetAttributes(v WebhooksOAuth2ClientCredentialsUpdateAttributes) {
	o.Attributes = v
}

// GetType returns the Type field value.
func (o *WebhooksOAuth2ClientCredentialsUpdateData) GetType() WebhooksOAuth2ClientCredentialsType {
	if o == nil {
		var ret WebhooksOAuth2ClientCredentialsType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WebhooksOAuth2ClientCredentialsUpdateData) GetTypeOk() (*WebhooksOAuth2ClientCredentialsType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *WebhooksOAuth2ClientCredentialsUpdateData) SetType(v WebhooksOAuth2ClientCredentialsType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o WebhooksOAuth2ClientCredentialsUpdateData) MarshalJSON() ([]byte, error) {
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
func (o *WebhooksOAuth2ClientCredentialsUpdateData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *WebhooksOAuth2ClientCredentialsUpdateAttributes `json:"attributes"`
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
