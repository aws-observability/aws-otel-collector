// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecureEmbedShareType The type of share. Always `secure_embed`.
type SecureEmbedShareType string

// List of SecureEmbedShareType.
const (
	SECUREEMBEDSHARETYPE_SECURE_EMBED SecureEmbedShareType = "secure_embed"
)

var allowedSecureEmbedShareTypeEnumValues = []SecureEmbedShareType{
	SECUREEMBEDSHARETYPE_SECURE_EMBED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecureEmbedShareType) GetAllowedValues() []SecureEmbedShareType {
	return allowedSecureEmbedShareTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecureEmbedShareType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecureEmbedShareType(value)
	return nil
}

// NewSecureEmbedShareTypeFromValue returns a pointer to a valid SecureEmbedShareType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecureEmbedShareTypeFromValue(v string) (*SecureEmbedShareType, error) {
	ev := SecureEmbedShareType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecureEmbedShareType: valid values are %v", v, allowedSecureEmbedShareTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecureEmbedShareType) IsValid() bool {
	for _, existing := range allowedSecureEmbedShareTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecureEmbedShareType value.
func (v SecureEmbedShareType) Ptr() *SecureEmbedShareType {
	return &v
}
