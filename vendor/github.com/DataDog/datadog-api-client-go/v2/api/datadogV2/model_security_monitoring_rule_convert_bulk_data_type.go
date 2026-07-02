// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleConvertBulkDataType The type of the resource.
type SecurityMonitoringRuleConvertBulkDataType string

// List of SecurityMonitoringRuleConvertBulkDataType.
const (
	SECURITYMONITORINGRULECONVERTBULKDATATYPE_SECURITY_MONITORING_RULES_CONVERT_BULK SecurityMonitoringRuleConvertBulkDataType = "security_monitoring_rules_convert_bulk"
)

var allowedSecurityMonitoringRuleConvertBulkDataTypeEnumValues = []SecurityMonitoringRuleConvertBulkDataType{
	SECURITYMONITORINGRULECONVERTBULKDATATYPE_SECURITY_MONITORING_RULES_CONVERT_BULK,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringRuleConvertBulkDataType) GetAllowedValues() []SecurityMonitoringRuleConvertBulkDataType {
	return allowedSecurityMonitoringRuleConvertBulkDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringRuleConvertBulkDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringRuleConvertBulkDataType(value)
	return nil
}

// NewSecurityMonitoringRuleConvertBulkDataTypeFromValue returns a pointer to a valid SecurityMonitoringRuleConvertBulkDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringRuleConvertBulkDataTypeFromValue(v string) (*SecurityMonitoringRuleConvertBulkDataType, error) {
	ev := SecurityMonitoringRuleConvertBulkDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringRuleConvertBulkDataType: valid values are %v", v, allowedSecurityMonitoringRuleConvertBulkDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringRuleConvertBulkDataType) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringRuleConvertBulkDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringRuleConvertBulkDataType value.
func (v SecurityMonitoringRuleConvertBulkDataType) Ptr() *SecurityMonitoringRuleConvertBulkDataType {
	return &v
}
