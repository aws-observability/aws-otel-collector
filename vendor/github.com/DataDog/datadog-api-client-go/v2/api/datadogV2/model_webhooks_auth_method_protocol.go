// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WebhooksAuthMethodProtocol Authentication protocol used by the auth method.
type WebhooksAuthMethodProtocol string

// List of WebhooksAuthMethodProtocol.
const (
	WEBHOOKSAUTHMETHODPROTOCOL_OAUTH2_CLIENT_CREDENTIALS WebhooksAuthMethodProtocol = "oauth2-client-credentials"
)

var allowedWebhooksAuthMethodProtocolEnumValues = []WebhooksAuthMethodProtocol{
	WEBHOOKSAUTHMETHODPROTOCOL_OAUTH2_CLIENT_CREDENTIALS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WebhooksAuthMethodProtocol) GetAllowedValues() []WebhooksAuthMethodProtocol {
	return allowedWebhooksAuthMethodProtocolEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WebhooksAuthMethodProtocol) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WebhooksAuthMethodProtocol(value)
	return nil
}

// NewWebhooksAuthMethodProtocolFromValue returns a pointer to a valid WebhooksAuthMethodProtocol
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWebhooksAuthMethodProtocolFromValue(v string) (*WebhooksAuthMethodProtocol, error) {
	ev := WebhooksAuthMethodProtocol(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WebhooksAuthMethodProtocol: valid values are %v", v, allowedWebhooksAuthMethodProtocolEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WebhooksAuthMethodProtocol) IsValid() bool {
	for _, existing := range allowedWebhooksAuthMethodProtocolEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WebhooksAuthMethodProtocol value.
func (v WebhooksAuthMethodProtocol) Ptr() *WebhooksAuthMethodProtocol {
	return &v
}
