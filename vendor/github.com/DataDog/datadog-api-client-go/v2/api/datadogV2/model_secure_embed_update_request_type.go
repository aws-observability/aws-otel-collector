// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecureEmbedUpdateRequestType Resource type for secure embed update requests.
type SecureEmbedUpdateRequestType string

// List of SecureEmbedUpdateRequestType.
const (
	SECUREEMBEDUPDATEREQUESTTYPE_SECURE_EMBED_UPDATE_REQUEST SecureEmbedUpdateRequestType = "secure_embed_update_request"
)

var allowedSecureEmbedUpdateRequestTypeEnumValues = []SecureEmbedUpdateRequestType{
	SECUREEMBEDUPDATEREQUESTTYPE_SECURE_EMBED_UPDATE_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecureEmbedUpdateRequestType) GetAllowedValues() []SecureEmbedUpdateRequestType {
	return allowedSecureEmbedUpdateRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecureEmbedUpdateRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecureEmbedUpdateRequestType(value)
	return nil
}

// NewSecureEmbedUpdateRequestTypeFromValue returns a pointer to a valid SecureEmbedUpdateRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecureEmbedUpdateRequestTypeFromValue(v string) (*SecureEmbedUpdateRequestType, error) {
	ev := SecureEmbedUpdateRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecureEmbedUpdateRequestType: valid values are %v", v, allowedSecureEmbedUpdateRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecureEmbedUpdateRequestType) IsValid() bool {
	for _, existing := range allowedSecureEmbedUpdateRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecureEmbedUpdateRequestType value.
func (v SecureEmbedUpdateRequestType) Ptr() *SecureEmbedUpdateRequestType {
	return &v
}
