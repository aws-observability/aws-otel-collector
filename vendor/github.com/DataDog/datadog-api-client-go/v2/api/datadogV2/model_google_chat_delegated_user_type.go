// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GoogleChatDelegatedUserType Google Chat delegated user resource type.
type GoogleChatDelegatedUserType string

// List of GoogleChatDelegatedUserType.
const (
	GOOGLECHATDELEGATEDUSERTYPE_GOOGLE_CHAT_DELEGATED_USER_TYPE GoogleChatDelegatedUserType = "google-chat-delegated-user"
)

var allowedGoogleChatDelegatedUserTypeEnumValues = []GoogleChatDelegatedUserType{
	GOOGLECHATDELEGATEDUSERTYPE_GOOGLE_CHAT_DELEGATED_USER_TYPE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GoogleChatDelegatedUserType) GetAllowedValues() []GoogleChatDelegatedUserType {
	return allowedGoogleChatDelegatedUserTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GoogleChatDelegatedUserType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GoogleChatDelegatedUserType(value)
	return nil
}

// NewGoogleChatDelegatedUserTypeFromValue returns a pointer to a valid GoogleChatDelegatedUserType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGoogleChatDelegatedUserTypeFromValue(v string) (*GoogleChatDelegatedUserType, error) {
	ev := GoogleChatDelegatedUserType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GoogleChatDelegatedUserType: valid values are %v", v, allowedGoogleChatDelegatedUserTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GoogleChatDelegatedUserType) IsValid() bool {
	for _, existing := range allowedGoogleChatDelegatedUserTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GoogleChatDelegatedUserType value.
func (v GoogleChatDelegatedUserType) Ptr() *GoogleChatDelegatedUserType {
	return &v
}
