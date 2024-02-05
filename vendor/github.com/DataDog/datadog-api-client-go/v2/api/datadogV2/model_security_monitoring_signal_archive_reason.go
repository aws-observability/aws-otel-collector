// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringSignalArchiveReason Reason a signal is archived.
type SecurityMonitoringSignalArchiveReason string

// List of SecurityMonitoringSignalArchiveReason.
const (
	SECURITYMONITORINGSIGNALARCHIVEREASON_NONE                     SecurityMonitoringSignalArchiveReason = "none"
	SECURITYMONITORINGSIGNALARCHIVEREASON_FALSE_POSITIVE           SecurityMonitoringSignalArchiveReason = "false_positive"
	SECURITYMONITORINGSIGNALARCHIVEREASON_TESTING_OR_MAINTENANCE   SecurityMonitoringSignalArchiveReason = "testing_or_maintenance"
	SECURITYMONITORINGSIGNALARCHIVEREASON_INVESTIGATED_CASE_OPENED SecurityMonitoringSignalArchiveReason = "investigated_case_opened"
	SECURITYMONITORINGSIGNALARCHIVEREASON_OTHER                    SecurityMonitoringSignalArchiveReason = "other"
)

var allowedSecurityMonitoringSignalArchiveReasonEnumValues = []SecurityMonitoringSignalArchiveReason{
	SECURITYMONITORINGSIGNALARCHIVEREASON_NONE,
	SECURITYMONITORINGSIGNALARCHIVEREASON_FALSE_POSITIVE,
	SECURITYMONITORINGSIGNALARCHIVEREASON_TESTING_OR_MAINTENANCE,
	SECURITYMONITORINGSIGNALARCHIVEREASON_INVESTIGATED_CASE_OPENED,
	SECURITYMONITORINGSIGNALARCHIVEREASON_OTHER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringSignalArchiveReason) GetAllowedValues() []SecurityMonitoringSignalArchiveReason {
	return allowedSecurityMonitoringSignalArchiveReasonEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringSignalArchiveReason) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringSignalArchiveReason(value)
	return nil
}

// NewSecurityMonitoringSignalArchiveReasonFromValue returns a pointer to a valid SecurityMonitoringSignalArchiveReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringSignalArchiveReasonFromValue(v string) (*SecurityMonitoringSignalArchiveReason, error) {
	ev := SecurityMonitoringSignalArchiveReason(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringSignalArchiveReason: valid values are %v", v, allowedSecurityMonitoringSignalArchiveReasonEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringSignalArchiveReason) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringSignalArchiveReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringSignalArchiveReason value.
func (v SecurityMonitoringSignalArchiveReason) Ptr() *SecurityMonitoringSignalArchiveReason {
	return &v
}
