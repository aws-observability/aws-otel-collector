// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MicrosoftTeamsTenantBasedHandleInfoType Tenant-based handle resource type.
type MicrosoftTeamsTenantBasedHandleInfoType string

// List of MicrosoftTeamsTenantBasedHandleInfoType.
const (
	MICROSOFTTEAMSTENANTBASEDHANDLEINFOTYPE_MS_TEAMS_TENANT_BASED_HANDLE_INFO MicrosoftTeamsTenantBasedHandleInfoType = "ms-teams-tenant-based-handle-info"
)

var allowedMicrosoftTeamsTenantBasedHandleInfoTypeEnumValues = []MicrosoftTeamsTenantBasedHandleInfoType{
	MICROSOFTTEAMSTENANTBASEDHANDLEINFOTYPE_MS_TEAMS_TENANT_BASED_HANDLE_INFO,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MicrosoftTeamsTenantBasedHandleInfoType) GetAllowedValues() []MicrosoftTeamsTenantBasedHandleInfoType {
	return allowedMicrosoftTeamsTenantBasedHandleInfoTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MicrosoftTeamsTenantBasedHandleInfoType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MicrosoftTeamsTenantBasedHandleInfoType(value)
	return nil
}

// NewMicrosoftTeamsTenantBasedHandleInfoTypeFromValue returns a pointer to a valid MicrosoftTeamsTenantBasedHandleInfoType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMicrosoftTeamsTenantBasedHandleInfoTypeFromValue(v string) (*MicrosoftTeamsTenantBasedHandleInfoType, error) {
	ev := MicrosoftTeamsTenantBasedHandleInfoType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MicrosoftTeamsTenantBasedHandleInfoType: valid values are %v", v, allowedMicrosoftTeamsTenantBasedHandleInfoTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MicrosoftTeamsTenantBasedHandleInfoType) IsValid() bool {
	for _, existing := range allowedMicrosoftTeamsTenantBasedHandleInfoTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MicrosoftTeamsTenantBasedHandleInfoType value.
func (v MicrosoftTeamsTenantBasedHandleInfoType) Ptr() *MicrosoftTeamsTenantBasedHandleInfoType {
	return &v
}
