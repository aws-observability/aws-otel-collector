// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringSignalType The type of event.
type SecurityMonitoringSignalType string

// List of SecurityMonitoringSignalType.
const (
	SECURITYMONITORINGSIGNALTYPE_SIGNAL SecurityMonitoringSignalType = "signal"
)

var allowedSecurityMonitoringSignalTypeEnumValues = []SecurityMonitoringSignalType{
	SECURITYMONITORINGSIGNALTYPE_SIGNAL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringSignalType) GetAllowedValues() []SecurityMonitoringSignalType {
	return allowedSecurityMonitoringSignalTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringSignalType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringSignalType(value)
	return nil
}

// NewSecurityMonitoringSignalTypeFromValue returns a pointer to a valid SecurityMonitoringSignalType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringSignalTypeFromValue(v string) (*SecurityMonitoringSignalType, error) {
	ev := SecurityMonitoringSignalType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringSignalType: valid values are %v", v, allowedSecurityMonitoringSignalTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringSignalType) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringSignalTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringSignalType value.
func (v SecurityMonitoringSignalType) Ptr() *SecurityMonitoringSignalType {
	return &v
}
