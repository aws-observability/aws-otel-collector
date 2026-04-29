// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecureEmbedGetResponseType Resource type for secure embed get responses.
type SecureEmbedGetResponseType string

// List of SecureEmbedGetResponseType.
const (
	SECUREEMBEDGETRESPONSETYPE_SECURE_EMBED_GET_RESPONSE SecureEmbedGetResponseType = "secure_embed_get_response"
)

var allowedSecureEmbedGetResponseTypeEnumValues = []SecureEmbedGetResponseType{
	SECUREEMBEDGETRESPONSETYPE_SECURE_EMBED_GET_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecureEmbedGetResponseType) GetAllowedValues() []SecureEmbedGetResponseType {
	return allowedSecureEmbedGetResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecureEmbedGetResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecureEmbedGetResponseType(value)
	return nil
}

// NewSecureEmbedGetResponseTypeFromValue returns a pointer to a valid SecureEmbedGetResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecureEmbedGetResponseTypeFromValue(v string) (*SecureEmbedGetResponseType, error) {
	ev := SecureEmbedGetResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecureEmbedGetResponseType: valid values are %v", v, allowedSecureEmbedGetResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecureEmbedGetResponseType) IsValid() bool {
	for _, existing := range allowedSecureEmbedGetResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecureEmbedGetResponseType value.
func (v SecureEmbedGetResponseType) Ptr() *SecureEmbedGetResponseType {
	return &v
}
