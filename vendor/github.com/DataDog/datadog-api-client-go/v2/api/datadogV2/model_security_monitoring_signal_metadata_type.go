// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringSignalMetadataType The type of event.
type SecurityMonitoringSignalMetadataType string

// List of SecurityMonitoringSignalMetadataType.
const (
	SECURITYMONITORINGSIGNALMETADATATYPE_SIGNAL_METADATA SecurityMonitoringSignalMetadataType = "signal_metadata"
)

var allowedSecurityMonitoringSignalMetadataTypeEnumValues = []SecurityMonitoringSignalMetadataType{
	SECURITYMONITORINGSIGNALMETADATATYPE_SIGNAL_METADATA,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringSignalMetadataType) GetAllowedValues() []SecurityMonitoringSignalMetadataType {
	return allowedSecurityMonitoringSignalMetadataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringSignalMetadataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringSignalMetadataType(value)
	return nil
}

// NewSecurityMonitoringSignalMetadataTypeFromValue returns a pointer to a valid SecurityMonitoringSignalMetadataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringSignalMetadataTypeFromValue(v string) (*SecurityMonitoringSignalMetadataType, error) {
	ev := SecurityMonitoringSignalMetadataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringSignalMetadataType: valid values are %v", v, allowedSecurityMonitoringSignalMetadataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringSignalMetadataType) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringSignalMetadataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringSignalMetadataType value.
func (v SecurityMonitoringSignalMetadataType) Ptr() *SecurityMonitoringSignalMetadataType {
	return &v
}
