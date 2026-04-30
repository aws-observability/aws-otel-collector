// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GoogleChatOrganizationHandleType Organization handle resource type.
type GoogleChatOrganizationHandleType string

// List of GoogleChatOrganizationHandleType.
const (
	GOOGLECHATORGANIZATIONHANDLETYPE_GOOGLE_CHAT_ORGANIZATION_HANDLE_TYPE GoogleChatOrganizationHandleType = "google-chat-organization-handle"
)

var allowedGoogleChatOrganizationHandleTypeEnumValues = []GoogleChatOrganizationHandleType{
	GOOGLECHATORGANIZATIONHANDLETYPE_GOOGLE_CHAT_ORGANIZATION_HANDLE_TYPE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GoogleChatOrganizationHandleType) GetAllowedValues() []GoogleChatOrganizationHandleType {
	return allowedGoogleChatOrganizationHandleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GoogleChatOrganizationHandleType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GoogleChatOrganizationHandleType(value)
	return nil
}

// NewGoogleChatOrganizationHandleTypeFromValue returns a pointer to a valid GoogleChatOrganizationHandleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGoogleChatOrganizationHandleTypeFromValue(v string) (*GoogleChatOrganizationHandleType, error) {
	ev := GoogleChatOrganizationHandleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GoogleChatOrganizationHandleType: valid values are %v", v, allowedGoogleChatOrganizationHandleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GoogleChatOrganizationHandleType) IsValid() bool {
	for _, existing := range allowedGoogleChatOrganizationHandleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GoogleChatOrganizationHandleType value.
func (v GoogleChatOrganizationHandleType) Ptr() *GoogleChatOrganizationHandleType {
	return &v
}
