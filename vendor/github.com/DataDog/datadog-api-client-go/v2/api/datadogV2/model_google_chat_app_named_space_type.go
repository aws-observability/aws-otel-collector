// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GoogleChatAppNamedSpaceType Google Chat space resource type.
type GoogleChatAppNamedSpaceType string

// List of GoogleChatAppNamedSpaceType.
const (
	GOOGLECHATAPPNAMEDSPACETYPE_GOOGLE_CHAT_APP_NAMED_SPACE_TYPE GoogleChatAppNamedSpaceType = "google-chat-app-named-space"
)

var allowedGoogleChatAppNamedSpaceTypeEnumValues = []GoogleChatAppNamedSpaceType{
	GOOGLECHATAPPNAMEDSPACETYPE_GOOGLE_CHAT_APP_NAMED_SPACE_TYPE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GoogleChatAppNamedSpaceType) GetAllowedValues() []GoogleChatAppNamedSpaceType {
	return allowedGoogleChatAppNamedSpaceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GoogleChatAppNamedSpaceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GoogleChatAppNamedSpaceType(value)
	return nil
}

// NewGoogleChatAppNamedSpaceTypeFromValue returns a pointer to a valid GoogleChatAppNamedSpaceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGoogleChatAppNamedSpaceTypeFromValue(v string) (*GoogleChatAppNamedSpaceType, error) {
	ev := GoogleChatAppNamedSpaceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GoogleChatAppNamedSpaceType: valid values are %v", v, allowedGoogleChatAppNamedSpaceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GoogleChatAppNamedSpaceType) IsValid() bool {
	for _, existing := range allowedGoogleChatAppNamedSpaceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GoogleChatAppNamedSpaceType value.
func (v GoogleChatAppNamedSpaceType) Ptr() *GoogleChatAppNamedSpaceType {
	return &v
}
