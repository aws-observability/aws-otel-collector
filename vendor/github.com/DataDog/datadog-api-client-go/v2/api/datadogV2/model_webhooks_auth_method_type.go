// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WebhooksAuthMethodType Webhooks auth method resource type.
type WebhooksAuthMethodType string

// List of WebhooksAuthMethodType.
const (
	WEBHOOKSAUTHMETHODTYPE_WEBHOOKS_AUTH_METHOD WebhooksAuthMethodType = "webhooks-auth-method"
)

var allowedWebhooksAuthMethodTypeEnumValues = []WebhooksAuthMethodType{
	WEBHOOKSAUTHMETHODTYPE_WEBHOOKS_AUTH_METHOD,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WebhooksAuthMethodType) GetAllowedValues() []WebhooksAuthMethodType {
	return allowedWebhooksAuthMethodTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WebhooksAuthMethodType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WebhooksAuthMethodType(value)
	return nil
}

// NewWebhooksAuthMethodTypeFromValue returns a pointer to a valid WebhooksAuthMethodType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWebhooksAuthMethodTypeFromValue(v string) (*WebhooksAuthMethodType, error) {
	ev := WebhooksAuthMethodType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WebhooksAuthMethodType: valid values are %v", v, allowedWebhooksAuthMethodTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WebhooksAuthMethodType) IsValid() bool {
	for _, existing := range allowedWebhooksAuthMethodTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WebhooksAuthMethodType value.
func (v WebhooksAuthMethodType) Ptr() *WebhooksAuthMethodType {
	return &v
}
