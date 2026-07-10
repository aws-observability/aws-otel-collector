// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WebhooksOAuth2ClientCredentialsRelationshipData Relationship data referencing an OAuth2 client credentials resource.
type WebhooksOAuth2ClientCredentialsRelationshipData struct {
	// The ID of the OAuth2 client credentials resource.
	Id *string `json:"id,omitempty"`
	// OAuth2 client credentials resource type.
	Type *WebhooksOAuth2ClientCredentialsType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWebhooksOAuth2ClientCredentialsRelationshipData instantiates a new WebhooksOAuth2ClientCredentialsRelationshipData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWebhooksOAuth2ClientCredentialsRelationshipData() *WebhooksOAuth2ClientCredentialsRelationshipData {
	this := WebhooksOAuth2ClientCredentialsRelationshipData{}
	var typeVar WebhooksOAuth2ClientCredentialsType = WEBHOOKSOAUTH2CLIENTCREDENTIALSTYPE_WEBHOOKS_AUTH_METHOD_OAUTH2_CLIENT_CREDENTIALS
	this.Type = &typeVar
	return &this
}

// NewWebhooksOAuth2ClientCredentialsRelationshipDataWithDefaults instantiates a new WebhooksOAuth2ClientCredentialsRelationshipData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWebhooksOAuth2ClientCredentialsRelationshipDataWithDefaults() *WebhooksOAuth2ClientCredentialsRelationshipData {
	this := WebhooksOAuth2ClientCredentialsRelationshipData{}
	var typeVar WebhooksOAuth2ClientCredentialsType = WEBHOOKSOAUTH2CLIENTCREDENTIALSTYPE_WEBHOOKS_AUTH_METHOD_OAUTH2_CLIENT_CREDENTIALS
	this.Type = &typeVar
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WebhooksOAuth2ClientCredentialsRelationshipData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksOAuth2ClientCredentialsRelationshipData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WebhooksOAuth2ClientCredentialsRelationshipData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WebhooksOAuth2ClientCredentialsRelationshipData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WebhooksOAuth2ClientCredentialsRelationshipData) GetType() WebhooksOAuth2ClientCredentialsType {
	if o == nil || o.Type == nil {
		var ret WebhooksOAuth2ClientCredentialsType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksOAuth2ClientCredentialsRelationshipData) GetTypeOk() (*WebhooksOAuth2ClientCredentialsType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WebhooksOAuth2ClientCredentialsRelationshipData) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given WebhooksOAuth2ClientCredentialsType and assigns it to the Type field.
func (o *WebhooksOAuth2ClientCredentialsRelationshipData) SetType(v WebhooksOAuth2ClientCredentialsType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o WebhooksOAuth2ClientCredentialsRelationshipData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WebhooksOAuth2ClientCredentialsRelationshipData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id   *string                              `json:"id,omitempty"`
		Type *WebhooksOAuth2ClientCredentialsType `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
