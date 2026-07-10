// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GoogleChatOrganizationType Google Chat organization resource type.
type GoogleChatOrganizationType string

// List of GoogleChatOrganizationType.
const (
	GOOGLECHATORGANIZATIONTYPE_GOOGLE_CHAT_ORGANIZATION_TYPE GoogleChatOrganizationType = "google-chat-organization"
)

var allowedGoogleChatOrganizationTypeEnumValues = []GoogleChatOrganizationType{
	GOOGLECHATORGANIZATIONTYPE_GOOGLE_CHAT_ORGANIZATION_TYPE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GoogleChatOrganizationType) GetAllowedValues() []GoogleChatOrganizationType {
	return allowedGoogleChatOrganizationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GoogleChatOrganizationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GoogleChatOrganizationType(value)
	return nil
}

// NewGoogleChatOrganizationTypeFromValue returns a pointer to a valid GoogleChatOrganizationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGoogleChatOrganizationTypeFromValue(v string) (*GoogleChatOrganizationType, error) {
	ev := GoogleChatOrganizationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GoogleChatOrganizationType: valid values are %v", v, allowedGoogleChatOrganizationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GoogleChatOrganizationType) IsValid() bool {
	for _, existing := range allowedGoogleChatOrganizationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GoogleChatOrganizationType value.
func (v GoogleChatOrganizationType) Ptr() *GoogleChatOrganizationType {
	return &v
}
