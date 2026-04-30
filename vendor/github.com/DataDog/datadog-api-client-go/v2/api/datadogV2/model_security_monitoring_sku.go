// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringSKU The Cloud SIEM pricing model (SKU) for the organization.
type SecurityMonitoringSKU string

// List of SecurityMonitoringSKU.
const (
	SECURITYMONITORINGSKU_PER_GB_ANALYZED              SecurityMonitoringSKU = "per_gb_analyzed"
	SECURITYMONITORINGSKU_PER_EVENT_IN_SIEM_INDEX_2023 SecurityMonitoringSKU = "per_event_in_siem_index_2023"
	SECURITYMONITORINGSKU_ADD_ON_2024                  SecurityMonitoringSKU = "add_on_2024"
)

var allowedSecurityMonitoringSKUEnumValues = []SecurityMonitoringSKU{
	SECURITYMONITORINGSKU_PER_GB_ANALYZED,
	SECURITYMONITORINGSKU_PER_EVENT_IN_SIEM_INDEX_2023,
	SECURITYMONITORINGSKU_ADD_ON_2024,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringSKU) GetAllowedValues() []SecurityMonitoringSKU {
	return allowedSecurityMonitoringSKUEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringSKU) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringSKU(value)
	return nil
}

// NewSecurityMonitoringSKUFromValue returns a pointer to a valid SecurityMonitoringSKU
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringSKUFromValue(v string) (*SecurityMonitoringSKU, error) {
	ev := SecurityMonitoringSKU(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringSKU: valid values are %v", v, allowedSecurityMonitoringSKUEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringSKU) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringSKUEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringSKU value.
func (v SecurityMonitoringSKU) Ptr() *SecurityMonitoringSKU {
	return &v
}
