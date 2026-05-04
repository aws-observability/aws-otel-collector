// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleSort The sort parameters used for querying security monitoring rules.
type SecurityMonitoringRuleSort string

// List of SecurityMonitoringRuleSort.
const (
	SECURITYMONITORINGRULESORT_NAME                        SecurityMonitoringRuleSort = "name"
	SECURITYMONITORINGRULESORT_CREATION_DATE               SecurityMonitoringRuleSort = "creation_date"
	SECURITYMONITORINGRULESORT_UPDATE_DATE                 SecurityMonitoringRuleSort = "update_date"
	SECURITYMONITORINGRULESORT_ENABLED                     SecurityMonitoringRuleSort = "enabled"
	SECURITYMONITORINGRULESORT_TYPE                        SecurityMonitoringRuleSort = "type"
	SECURITYMONITORINGRULESORT_HIGHEST_SEVERITY            SecurityMonitoringRuleSort = "highest_severity"
	SECURITYMONITORINGRULESORT_SOURCE                      SecurityMonitoringRuleSort = "source"
	SECURITYMONITORINGRULESORT_NAME_DESCENDING             SecurityMonitoringRuleSort = "-name"
	SECURITYMONITORINGRULESORT_CREATION_DATE_DESCENDING    SecurityMonitoringRuleSort = "-creation_date"
	SECURITYMONITORINGRULESORT_UPDATE_DATE_DESCENDING      SecurityMonitoringRuleSort = "-update_date"
	SECURITYMONITORINGRULESORT_ENABLED_DESCENDING          SecurityMonitoringRuleSort = "-enabled"
	SECURITYMONITORINGRULESORT_TYPE_DESCENDING             SecurityMonitoringRuleSort = "-type"
	SECURITYMONITORINGRULESORT_HIGHEST_SEVERITY_DESCENDING SecurityMonitoringRuleSort = "-highest_severity"
	SECURITYMONITORINGRULESORT_SOURCE_DESCENDING           SecurityMonitoringRuleSort = "-source"
)

var allowedSecurityMonitoringRuleSortEnumValues = []SecurityMonitoringRuleSort{
	SECURITYMONITORINGRULESORT_NAME,
	SECURITYMONITORINGRULESORT_CREATION_DATE,
	SECURITYMONITORINGRULESORT_UPDATE_DATE,
	SECURITYMONITORINGRULESORT_ENABLED,
	SECURITYMONITORINGRULESORT_TYPE,
	SECURITYMONITORINGRULESORT_HIGHEST_SEVERITY,
	SECURITYMONITORINGRULESORT_SOURCE,
	SECURITYMONITORINGRULESORT_NAME_DESCENDING,
	SECURITYMONITORINGRULESORT_CREATION_DATE_DESCENDING,
	SECURITYMONITORINGRULESORT_UPDATE_DATE_DESCENDING,
	SECURITYMONITORINGRULESORT_ENABLED_DESCENDING,
	SECURITYMONITORINGRULESORT_TYPE_DESCENDING,
	SECURITYMONITORINGRULESORT_HIGHEST_SEVERITY_DESCENDING,
	SECURITYMONITORINGRULESORT_SOURCE_DESCENDING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringRuleSort) GetAllowedValues() []SecurityMonitoringRuleSort {
	return allowedSecurityMonitoringRuleSortEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringRuleSort) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringRuleSort(value)
	return nil
}

// NewSecurityMonitoringRuleSortFromValue returns a pointer to a valid SecurityMonitoringRuleSort
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringRuleSortFromValue(v string) (*SecurityMonitoringRuleSort, error) {
	ev := SecurityMonitoringRuleSort(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringRuleSort: valid values are %v", v, allowedSecurityMonitoringRuleSortEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringRuleSort) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringRuleSortEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringRuleSort value.
func (v SecurityMonitoringRuleSort) Ptr() *SecurityMonitoringRuleSort {
	return &v
}
