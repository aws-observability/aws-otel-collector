// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnCallPageTargetType The kind of target, `team_id` | `team_handle` | `user_id`.
type OnCallPageTargetType string

// List of OnCallPageTargetType.
const (
	ONCALLPAGETARGETTYPE_TEAM_ID     OnCallPageTargetType = "team_id"
	ONCALLPAGETARGETTYPE_TEAM_HANDLE OnCallPageTargetType = "team_handle"
	ONCALLPAGETARGETTYPE_USER_ID     OnCallPageTargetType = "user_id"
)

var allowedOnCallPageTargetTypeEnumValues = []OnCallPageTargetType{
	ONCALLPAGETARGETTYPE_TEAM_ID,
	ONCALLPAGETARGETTYPE_TEAM_HANDLE,
	ONCALLPAGETARGETTYPE_USER_ID,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OnCallPageTargetType) GetAllowedValues() []OnCallPageTargetType {
	return allowedOnCallPageTargetTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OnCallPageTargetType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OnCallPageTargetType(value)
	return nil
}

// NewOnCallPageTargetTypeFromValue returns a pointer to a valid OnCallPageTargetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOnCallPageTargetTypeFromValue(v string) (*OnCallPageTargetType, error) {
	ev := OnCallPageTargetType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OnCallPageTargetType: valid values are %v", v, allowedOnCallPageTargetTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OnCallPageTargetType) IsValid() bool {
	for _, existing := range allowedOnCallPageTargetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OnCallPageTargetType value.
func (v OnCallPageTargetType) Ptr() *OnCallPageTargetType {
	return &v
}
