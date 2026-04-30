// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleBulkExportDataType The type of the resource.
type SecurityMonitoringRuleBulkExportDataType string

// List of SecurityMonitoringRuleBulkExportDataType.
const (
	SECURITYMONITORINGRULEBULKEXPORTDATATYPE_SECURITY_MONITORING_RULES_BULK_EXPORT SecurityMonitoringRuleBulkExportDataType = "security_monitoring_rules_bulk_export"
)

var allowedSecurityMonitoringRuleBulkExportDataTypeEnumValues = []SecurityMonitoringRuleBulkExportDataType{
	SECURITYMONITORINGRULEBULKEXPORTDATATYPE_SECURITY_MONITORING_RULES_BULK_EXPORT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringRuleBulkExportDataType) GetAllowedValues() []SecurityMonitoringRuleBulkExportDataType {
	return allowedSecurityMonitoringRuleBulkExportDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringRuleBulkExportDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringRuleBulkExportDataType(value)
	return nil
}

// NewSecurityMonitoringRuleBulkExportDataTypeFromValue returns a pointer to a valid SecurityMonitoringRuleBulkExportDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringRuleBulkExportDataTypeFromValue(v string) (*SecurityMonitoringRuleBulkExportDataType, error) {
	ev := SecurityMonitoringRuleBulkExportDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringRuleBulkExportDataType: valid values are %v", v, allowedSecurityMonitoringRuleBulkExportDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringRuleBulkExportDataType) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringRuleBulkExportDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringRuleBulkExportDataType value.
func (v SecurityMonitoringRuleBulkExportDataType) Ptr() *SecurityMonitoringRuleBulkExportDataType {
	return &v
}
