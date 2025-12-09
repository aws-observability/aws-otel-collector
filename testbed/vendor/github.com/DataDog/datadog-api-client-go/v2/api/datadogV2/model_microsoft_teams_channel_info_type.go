// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MicrosoftTeamsChannelInfoType Channel info resource type.
type MicrosoftTeamsChannelInfoType string

// List of MicrosoftTeamsChannelInfoType.
const (
	MICROSOFTTEAMSCHANNELINFOTYPE_MS_TEAMS_CHANNEL_INFO MicrosoftTeamsChannelInfoType = "ms-teams-channel-info"
)

var allowedMicrosoftTeamsChannelInfoTypeEnumValues = []MicrosoftTeamsChannelInfoType{
	MICROSOFTTEAMSCHANNELINFOTYPE_MS_TEAMS_CHANNEL_INFO,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MicrosoftTeamsChannelInfoType) GetAllowedValues() []MicrosoftTeamsChannelInfoType {
	return allowedMicrosoftTeamsChannelInfoTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MicrosoftTeamsChannelInfoType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MicrosoftTeamsChannelInfoType(value)
	return nil
}

// NewMicrosoftTeamsChannelInfoTypeFromValue returns a pointer to a valid MicrosoftTeamsChannelInfoType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMicrosoftTeamsChannelInfoTypeFromValue(v string) (*MicrosoftTeamsChannelInfoType, error) {
	ev := MicrosoftTeamsChannelInfoType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MicrosoftTeamsChannelInfoType: valid values are %v", v, allowedMicrosoftTeamsChannelInfoTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MicrosoftTeamsChannelInfoType) IsValid() bool {
	for _, existing := range allowedMicrosoftTeamsChannelInfoTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MicrosoftTeamsChannelInfoType value.
func (v MicrosoftTeamsChannelInfoType) Ptr() *MicrosoftTeamsChannelInfoType {
	return &v
}
