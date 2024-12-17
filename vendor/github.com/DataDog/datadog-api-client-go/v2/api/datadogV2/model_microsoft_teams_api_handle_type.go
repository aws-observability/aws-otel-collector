// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MicrosoftTeamsApiHandleType Specifies the handle resource type.
type MicrosoftTeamsApiHandleType string

// List of MicrosoftTeamsApiHandleType.
const (
	MICROSOFTTEAMSAPIHANDLETYPE_HANDLE MicrosoftTeamsApiHandleType = "handle"
)

var allowedMicrosoftTeamsApiHandleTypeEnumValues = []MicrosoftTeamsApiHandleType{
	MICROSOFTTEAMSAPIHANDLETYPE_HANDLE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MicrosoftTeamsApiHandleType) GetAllowedValues() []MicrosoftTeamsApiHandleType {
	return allowedMicrosoftTeamsApiHandleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MicrosoftTeamsApiHandleType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MicrosoftTeamsApiHandleType(value)
	return nil
}

// NewMicrosoftTeamsApiHandleTypeFromValue returns a pointer to a valid MicrosoftTeamsApiHandleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMicrosoftTeamsApiHandleTypeFromValue(v string) (*MicrosoftTeamsApiHandleType, error) {
	ev := MicrosoftTeamsApiHandleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MicrosoftTeamsApiHandleType: valid values are %v", v, allowedMicrosoftTeamsApiHandleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MicrosoftTeamsApiHandleType) IsValid() bool {
	for _, existing := range allowedMicrosoftTeamsApiHandleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MicrosoftTeamsApiHandleType value.
func (v MicrosoftTeamsApiHandleType) Ptr() *MicrosoftTeamsApiHandleType {
	return &v
}
