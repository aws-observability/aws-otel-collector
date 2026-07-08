// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WebhooksOAuth2ClientCredentialsType OAuth2 client credentials resource type.
type WebhooksOAuth2ClientCredentialsType string

// List of WebhooksOAuth2ClientCredentialsType.
const (
	WEBHOOKSOAUTH2CLIENTCREDENTIALSTYPE_WEBHOOKS_AUTH_METHOD_OAUTH2_CLIENT_CREDENTIALS WebhooksOAuth2ClientCredentialsType = "webhooks-auth-method-oauth2-client-credentials"
)

var allowedWebhooksOAuth2ClientCredentialsTypeEnumValues = []WebhooksOAuth2ClientCredentialsType{
	WEBHOOKSOAUTH2CLIENTCREDENTIALSTYPE_WEBHOOKS_AUTH_METHOD_OAUTH2_CLIENT_CREDENTIALS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WebhooksOAuth2ClientCredentialsType) GetAllowedValues() []WebhooksOAuth2ClientCredentialsType {
	return allowedWebhooksOAuth2ClientCredentialsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WebhooksOAuth2ClientCredentialsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WebhooksOAuth2ClientCredentialsType(value)
	return nil
}

// NewWebhooksOAuth2ClientCredentialsTypeFromValue returns a pointer to a valid WebhooksOAuth2ClientCredentialsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWebhooksOAuth2ClientCredentialsTypeFromValue(v string) (*WebhooksOAuth2ClientCredentialsType, error) {
	ev := WebhooksOAuth2ClientCredentialsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WebhooksOAuth2ClientCredentialsType: valid values are %v", v, allowedWebhooksOAuth2ClientCredentialsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WebhooksOAuth2ClientCredentialsType) IsValid() bool {
	for _, existing := range allowedWebhooksOAuth2ClientCredentialsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WebhooksOAuth2ClientCredentialsType value.
func (v WebhooksOAuth2ClientCredentialsType) Ptr() *WebhooksOAuth2ClientCredentialsType {
	return &v
}
