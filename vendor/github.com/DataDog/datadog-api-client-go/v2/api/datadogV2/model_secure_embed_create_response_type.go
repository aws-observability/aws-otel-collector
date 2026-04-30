// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecureEmbedCreateResponseType Resource type for secure embed create responses.
type SecureEmbedCreateResponseType string

// List of SecureEmbedCreateResponseType.
const (
	SECUREEMBEDCREATERESPONSETYPE_SECURE_EMBED_CREATE_RESPONSE SecureEmbedCreateResponseType = "secure_embed_create_response"
)

var allowedSecureEmbedCreateResponseTypeEnumValues = []SecureEmbedCreateResponseType{
	SECUREEMBEDCREATERESPONSETYPE_SECURE_EMBED_CREATE_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecureEmbedCreateResponseType) GetAllowedValues() []SecureEmbedCreateResponseType {
	return allowedSecureEmbedCreateResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecureEmbedCreateResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecureEmbedCreateResponseType(value)
	return nil
}

// NewSecureEmbedCreateResponseTypeFromValue returns a pointer to a valid SecureEmbedCreateResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecureEmbedCreateResponseTypeFromValue(v string) (*SecureEmbedCreateResponseType, error) {
	ev := SecureEmbedCreateResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecureEmbedCreateResponseType: valid values are %v", v, allowedSecureEmbedCreateResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecureEmbedCreateResponseType) IsValid() bool {
	for _, existing := range allowedSecureEmbedCreateResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecureEmbedCreateResponseType value.
func (v SecureEmbedCreateResponseType) Ptr() *SecureEmbedCreateResponseType {
	return &v
}
