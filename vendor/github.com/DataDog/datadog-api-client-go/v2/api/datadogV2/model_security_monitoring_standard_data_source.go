// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringStandardDataSource Source of events, either logs, audit trail, or Datadog events.
type SecurityMonitoringStandardDataSource string

// List of SecurityMonitoringStandardDataSource.
const (
	SECURITYMONITORINGSTANDARDDATASOURCE_LOGS             SecurityMonitoringStandardDataSource = "logs"
	SECURITYMONITORINGSTANDARDDATASOURCE_AUDIT            SecurityMonitoringStandardDataSource = "audit"
	SECURITYMONITORINGSTANDARDDATASOURCE_APP_SEC_SPANS    SecurityMonitoringStandardDataSource = "app_sec_spans"
	SECURITYMONITORINGSTANDARDDATASOURCE_SPANS            SecurityMonitoringStandardDataSource = "spans"
	SECURITYMONITORINGSTANDARDDATASOURCE_SECURITY_RUNTIME SecurityMonitoringStandardDataSource = "security_runtime"
	SECURITYMONITORINGSTANDARDDATASOURCE_NETWORK          SecurityMonitoringStandardDataSource = "network"
	SECURITYMONITORINGSTANDARDDATASOURCE_EVENTS           SecurityMonitoringStandardDataSource = "events"
)

var allowedSecurityMonitoringStandardDataSourceEnumValues = []SecurityMonitoringStandardDataSource{
	SECURITYMONITORINGSTANDARDDATASOURCE_LOGS,
	SECURITYMONITORINGSTANDARDDATASOURCE_AUDIT,
	SECURITYMONITORINGSTANDARDDATASOURCE_APP_SEC_SPANS,
	SECURITYMONITORINGSTANDARDDATASOURCE_SPANS,
	SECURITYMONITORINGSTANDARDDATASOURCE_SECURITY_RUNTIME,
	SECURITYMONITORINGSTANDARDDATASOURCE_NETWORK,
	SECURITYMONITORINGSTANDARDDATASOURCE_EVENTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringStandardDataSource) GetAllowedValues() []SecurityMonitoringStandardDataSource {
	return allowedSecurityMonitoringStandardDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringStandardDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringStandardDataSource(value)
	return nil
}

// NewSecurityMonitoringStandardDataSourceFromValue returns a pointer to a valid SecurityMonitoringStandardDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringStandardDataSourceFromValue(v string) (*SecurityMonitoringStandardDataSource, error) {
	ev := SecurityMonitoringStandardDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringStandardDataSource: valid values are %v", v, allowedSecurityMonitoringStandardDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringStandardDataSource) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringStandardDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringStandardDataSource value.
func (v SecurityMonitoringStandardDataSource) Ptr() *SecurityMonitoringStandardDataSource {
	return &v
}
