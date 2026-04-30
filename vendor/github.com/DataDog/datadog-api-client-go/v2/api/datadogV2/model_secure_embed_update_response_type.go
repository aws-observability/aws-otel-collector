// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecureEmbedUpdateResponseType Resource type for secure embed update responses.
type SecureEmbedUpdateResponseType string

// List of SecureEmbedUpdateResponseType.
const (
	SECUREEMBEDUPDATERESPONSETYPE_SECURE_EMBED_UPDATE_RESPONSE SecureEmbedUpdateResponseType = "secure_embed_update_response"
)

var allowedSecureEmbedUpdateResponseTypeEnumValues = []SecureEmbedUpdateResponseType{
	SECUREEMBEDUPDATERESPONSETYPE_SECURE_EMBED_UPDATE_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecureEmbedUpdateResponseType) GetAllowedValues() []SecureEmbedUpdateResponseType {
	return allowedSecureEmbedUpdateResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecureEmbedUpdateResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecureEmbedUpdateResponseType(value)
	return nil
}

// NewSecureEmbedUpdateResponseTypeFromValue returns a pointer to a valid SecureEmbedUpdateResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecureEmbedUpdateResponseTypeFromValue(v string) (*SecureEmbedUpdateResponseType, error) {
	ev := SecureEmbedUpdateResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecureEmbedUpdateResponseType: valid values are %v", v, allowedSecureEmbedUpdateResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecureEmbedUpdateResponseType) IsValid() bool {
	for _, existing := range allowedSecureEmbedUpdateResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecureEmbedUpdateResponseType value.
func (v SecureEmbedUpdateResponseType) Ptr() *SecureEmbedUpdateResponseType {
	return &v
}
