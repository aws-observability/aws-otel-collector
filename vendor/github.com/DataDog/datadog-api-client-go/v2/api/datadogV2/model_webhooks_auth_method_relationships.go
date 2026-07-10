// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WebhooksAuthMethodRelationships Relationships of a webhooks auth method to its protocol-specific resource.
type WebhooksAuthMethodRelationships struct {
	// Relationship pointing to the OAuth2 client credentials resource for this auth method.
	Oauth2ClientCredentials *WebhooksOAuth2ClientCredentialsRelationship `json:"oauth2-client-credentials,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWebhooksAuthMethodRelationships instantiates a new WebhooksAuthMethodRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWebhooksAuthMethodRelationships() *WebhooksAuthMethodRelationships {
	this := WebhooksAuthMethodRelationships{}
	return &this
}

// NewWebhooksAuthMethodRelationshipsWithDefaults instantiates a new WebhooksAuthMethodRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWebhooksAuthMethodRelationshipsWithDefaults() *WebhooksAuthMethodRelationships {
	this := WebhooksAuthMethodRelationships{}
	return &this
}

// GetOauth2ClientCredentials returns the Oauth2ClientCredentials field value if set, zero value otherwise.
func (o *WebhooksAuthMethodRelationships) GetOauth2ClientCredentials() WebhooksOAuth2ClientCredentialsRelationship {
	if o == nil || o.Oauth2ClientCredentials == nil {
		var ret WebhooksOAuth2ClientCredentialsRelationship
		return ret
	}
	return *o.Oauth2ClientCredentials
}

// GetOauth2ClientCredentialsOk returns a tuple with the Oauth2ClientCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksAuthMethodRelationships) GetOauth2ClientCredentialsOk() (*WebhooksOAuth2ClientCredentialsRelationship, bool) {
	if o == nil || o.Oauth2ClientCredentials == nil {
		return nil, false
	}
	return o.Oauth2ClientCredentials, true
}

// HasOauth2ClientCredentials returns a boolean if a field has been set.
func (o *WebhooksAuthMethodRelationships) HasOauth2ClientCredentials() bool {
	return o != nil && o.Oauth2ClientCredentials != nil
}

// SetOauth2ClientCredentials gets a reference to the given WebhooksOAuth2ClientCredentialsRelationship and assigns it to the Oauth2ClientCredentials field.
func (o *WebhooksAuthMethodRelationships) SetOauth2ClientCredentials(v WebhooksOAuth2ClientCredentialsRelationship) {
	o.Oauth2ClientCredentials = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o WebhooksAuthMethodRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Oauth2ClientCredentials != nil {
		toSerialize["oauth2-client-credentials"] = o.Oauth2ClientCredentials
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WebhooksAuthMethodRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Oauth2ClientCredentials *WebhooksOAuth2ClientCredentialsRelationship `json:"oauth2-client-credentials,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"oauth2-client-credentials"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Oauth2ClientCredentials != nil && all.Oauth2ClientCredentials.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Oauth2ClientCredentials = all.Oauth2ClientCredentials

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
