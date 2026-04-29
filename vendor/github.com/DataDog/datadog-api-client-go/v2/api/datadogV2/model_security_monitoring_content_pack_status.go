// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringContentPackStatus The current operational status of a content pack.
type SecurityMonitoringContentPackStatus string

// List of SecurityMonitoringContentPackStatus.
const (
	SECURITYMONITORINGCONTENTPACKSTATUS_INSTALL      SecurityMonitoringContentPackStatus = "install"
	SECURITYMONITORINGCONTENTPACKSTATUS_ACTIVATE     SecurityMonitoringContentPackStatus = "activate"
	SECURITYMONITORINGCONTENTPACKSTATUS_INITIALIZING SecurityMonitoringContentPackStatus = "initializing"
	SECURITYMONITORINGCONTENTPACKSTATUS_ACTIVE       SecurityMonitoringContentPackStatus = "active"
	SECURITYMONITORINGCONTENTPACKSTATUS_WARNING      SecurityMonitoringContentPackStatus = "warning"
	SECURITYMONITORINGCONTENTPACKSTATUS_BROKEN       SecurityMonitoringContentPackStatus = "broken"
)

var allowedSecurityMonitoringContentPackStatusEnumValues = []SecurityMonitoringContentPackStatus{
	SECURITYMONITORINGCONTENTPACKSTATUS_INSTALL,
	SECURITYMONITORINGCONTENTPACKSTATUS_ACTIVATE,
	SECURITYMONITORINGCONTENTPACKSTATUS_INITIALIZING,
	SECURITYMONITORINGCONTENTPACKSTATUS_ACTIVE,
	SECURITYMONITORINGCONTENTPACKSTATUS_WARNING,
	SECURITYMONITORINGCONTENTPACKSTATUS_BROKEN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringContentPackStatus) GetAllowedValues() []SecurityMonitoringContentPackStatus {
	return allowedSecurityMonitoringContentPackStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringContentPackStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringContentPackStatus(value)
	return nil
}

// NewSecurityMonitoringContentPackStatusFromValue returns a pointer to a valid SecurityMonitoringContentPackStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringContentPackStatusFromValue(v string) (*SecurityMonitoringContentPackStatus, error) {
	ev := SecurityMonitoringContentPackStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringContentPackStatus: valid values are %v", v, allowedSecurityMonitoringContentPackStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringContentPackStatus) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringContentPackStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringContentPackStatus value.
func (v SecurityMonitoringContentPackStatus) Ptr() *SecurityMonitoringContentPackStatus {
	return &v
}
