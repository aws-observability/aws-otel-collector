// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MicrosoftTeamsApiHandleInfoType Handle resource type.
type MicrosoftTeamsApiHandleInfoType string

// List of MicrosoftTeamsApiHandleInfoType.
const (
	MICROSOFTTEAMSAPIHANDLEINFOTYPE_MS_TEAMS_HANDLE_INFO MicrosoftTeamsApiHandleInfoType = "ms-teams-handle-info"
)

var allowedMicrosoftTeamsApiHandleInfoTypeEnumValues = []MicrosoftTeamsApiHandleInfoType{
	MICROSOFTTEAMSAPIHANDLEINFOTYPE_MS_TEAMS_HANDLE_INFO,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MicrosoftTeamsApiHandleInfoType) GetAllowedValues() []MicrosoftTeamsApiHandleInfoType {
	return allowedMicrosoftTeamsApiHandleInfoTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MicrosoftTeamsApiHandleInfoType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MicrosoftTeamsApiHandleInfoType(value)
	return nil
}

// NewMicrosoftTeamsApiHandleInfoTypeFromValue returns a pointer to a valid MicrosoftTeamsApiHandleInfoType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMicrosoftTeamsApiHandleInfoTypeFromValue(v string) (*MicrosoftTeamsApiHandleInfoType, error) {
	ev := MicrosoftTeamsApiHandleInfoType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MicrosoftTeamsApiHandleInfoType: valid values are %v", v, allowedMicrosoftTeamsApiHandleInfoTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MicrosoftTeamsApiHandleInfoType) IsValid() bool {
	for _, existing := range allowedMicrosoftTeamsApiHandleInfoTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MicrosoftTeamsApiHandleInfoType value.
func (v MicrosoftTeamsApiHandleInfoType) Ptr() *MicrosoftTeamsApiHandleInfoType {
	return &v
}
