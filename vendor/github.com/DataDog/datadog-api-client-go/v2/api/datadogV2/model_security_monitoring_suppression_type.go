// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringSuppressionType The type of the resource. The value should always be `suppressions`.
type SecurityMonitoringSuppressionType string

// List of SecurityMonitoringSuppressionType.
const (
	SECURITYMONITORINGSUPPRESSIONTYPE_SUPPRESSIONS SecurityMonitoringSuppressionType = "suppressions"
)

var allowedSecurityMonitoringSuppressionTypeEnumValues = []SecurityMonitoringSuppressionType{
	SECURITYMONITORINGSUPPRESSIONTYPE_SUPPRESSIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringSuppressionType) GetAllowedValues() []SecurityMonitoringSuppressionType {
	return allowedSecurityMonitoringSuppressionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringSuppressionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringSuppressionType(value)
	return nil
}

// NewSecurityMonitoringSuppressionTypeFromValue returns a pointer to a valid SecurityMonitoringSuppressionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringSuppressionTypeFromValue(v string) (*SecurityMonitoringSuppressionType, error) {
	ev := SecurityMonitoringSuppressionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringSuppressionType: valid values are %v", v, allowedSecurityMonitoringSuppressionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringSuppressionType) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringSuppressionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringSuppressionType value.
func (v SecurityMonitoringSuppressionType) Ptr() *SecurityMonitoringSuppressionType {
	return &v
}
