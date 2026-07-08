// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GoogleChatTargetAudienceType Google Chat target audience resource type.
type GoogleChatTargetAudienceType string

// List of GoogleChatTargetAudienceType.
const (
	GOOGLECHATTARGETAUDIENCETYPE_GOOGLE_CHAT_TARGET_AUDIENCE_TYPE GoogleChatTargetAudienceType = "google-chat-target-audience"
)

var allowedGoogleChatTargetAudienceTypeEnumValues = []GoogleChatTargetAudienceType{
	GOOGLECHATTARGETAUDIENCETYPE_GOOGLE_CHAT_TARGET_AUDIENCE_TYPE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GoogleChatTargetAudienceType) GetAllowedValues() []GoogleChatTargetAudienceType {
	return allowedGoogleChatTargetAudienceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GoogleChatTargetAudienceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GoogleChatTargetAudienceType(value)
	return nil
}

// NewGoogleChatTargetAudienceTypeFromValue returns a pointer to a valid GoogleChatTargetAudienceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGoogleChatTargetAudienceTypeFromValue(v string) (*GoogleChatTargetAudienceType, error) {
	ev := GoogleChatTargetAudienceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GoogleChatTargetAudienceType: valid values are %v", v, allowedGoogleChatTargetAudienceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GoogleChatTargetAudienceType) IsValid() bool {
	for _, existing := range allowedGoogleChatTargetAudienceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GoogleChatTargetAudienceType value.
func (v GoogleChatTargetAudienceType) Ptr() *GoogleChatTargetAudienceType {
	return &v
}
