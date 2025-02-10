// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MicrosoftTeamsTenantBasedHandleType Specifies the tenant-based handle resource type.
type MicrosoftTeamsTenantBasedHandleType string

// List of MicrosoftTeamsTenantBasedHandleType.
const (
	MICROSOFTTEAMSTENANTBASEDHANDLETYPE_TENANT_BASED_HANDLE MicrosoftTeamsTenantBasedHandleType = "tenant-based-handle"
)

var allowedMicrosoftTeamsTenantBasedHandleTypeEnumValues = []MicrosoftTeamsTenantBasedHandleType{
	MICROSOFTTEAMSTENANTBASEDHANDLETYPE_TENANT_BASED_HANDLE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MicrosoftTeamsTenantBasedHandleType) GetAllowedValues() []MicrosoftTeamsTenantBasedHandleType {
	return allowedMicrosoftTeamsTenantBasedHandleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MicrosoftTeamsTenantBasedHandleType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MicrosoftTeamsTenantBasedHandleType(value)
	return nil
}

// NewMicrosoftTeamsTenantBasedHandleTypeFromValue returns a pointer to a valid MicrosoftTeamsTenantBasedHandleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMicrosoftTeamsTenantBasedHandleTypeFromValue(v string) (*MicrosoftTeamsTenantBasedHandleType, error) {
	ev := MicrosoftTeamsTenantBasedHandleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MicrosoftTeamsTenantBasedHandleType: valid values are %v", v, allowedMicrosoftTeamsTenantBasedHandleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MicrosoftTeamsTenantBasedHandleType) IsValid() bool {
	for _, existing := range allowedMicrosoftTeamsTenantBasedHandleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MicrosoftTeamsTenantBasedHandleType value.
func (v MicrosoftTeamsTenantBasedHandleType) Ptr() *MicrosoftTeamsTenantBasedHandleType {
	return &v
}
