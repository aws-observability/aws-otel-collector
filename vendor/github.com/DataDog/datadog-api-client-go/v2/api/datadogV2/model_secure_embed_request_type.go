// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecureEmbedRequestType Resource type for secure embed create requests.
type SecureEmbedRequestType string

// List of SecureEmbedRequestType.
const (
	SECUREEMBEDREQUESTTYPE_SECURE_EMBED_REQUEST SecureEmbedRequestType = "secure_embed_request"
)

var allowedSecureEmbedRequestTypeEnumValues = []SecureEmbedRequestType{
	SECUREEMBEDREQUESTTYPE_SECURE_EMBED_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecureEmbedRequestType) GetAllowedValues() []SecureEmbedRequestType {
	return allowedSecureEmbedRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecureEmbedRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecureEmbedRequestType(value)
	return nil
}

// NewSecureEmbedRequestTypeFromValue returns a pointer to a valid SecureEmbedRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecureEmbedRequestTypeFromValue(v string) (*SecureEmbedRequestType, error) {
	ev := SecureEmbedRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecureEmbedRequestType: valid values are %v", v, allowedSecureEmbedRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecureEmbedRequestType) IsValid() bool {
	for _, existing := range allowedSecureEmbedRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecureEmbedRequestType value.
func (v SecureEmbedRequestType) Ptr() *SecureEmbedRequestType {
	return &v
}
