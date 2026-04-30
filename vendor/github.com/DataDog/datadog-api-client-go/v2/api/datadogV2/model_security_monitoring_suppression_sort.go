// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringSuppressionSort The sort parameters used for querying suppression rules.
type SecurityMonitoringSuppressionSort string

// List of SecurityMonitoringSuppressionSort.
const (
	SECURITYMONITORINGSUPPRESSIONSORT_NAME                       SecurityMonitoringSuppressionSort = "name"
	SECURITYMONITORINGSUPPRESSIONSORT_START_DATE                 SecurityMonitoringSuppressionSort = "start_date"
	SECURITYMONITORINGSUPPRESSIONSORT_EXPIRATION_DATE            SecurityMonitoringSuppressionSort = "expiration_date"
	SECURITYMONITORINGSUPPRESSIONSORT_UPDATE_DATE                SecurityMonitoringSuppressionSort = "update_date"
	SECURITYMONITORINGSUPPRESSIONSORT_ENABLED                    SecurityMonitoringSuppressionSort = "enabled"
	SECURITYMONITORINGSUPPRESSIONSORT_NAME_DESCENDING            SecurityMonitoringSuppressionSort = "-name"
	SECURITYMONITORINGSUPPRESSIONSORT_START_DATE_DESCENDING      SecurityMonitoringSuppressionSort = "-start_date"
	SECURITYMONITORINGSUPPRESSIONSORT_EXPIRATION_DATE_DESCENDING SecurityMonitoringSuppressionSort = "-expiration_date"
	SECURITYMONITORINGSUPPRESSIONSORT_UPDATE_DATE_DESCENDING     SecurityMonitoringSuppressionSort = "-update_date"
	SECURITYMONITORINGSUPPRESSIONSORT_CREATION_DATE_DESCENDING   SecurityMonitoringSuppressionSort = "-creation_date"
	SECURITYMONITORINGSUPPRESSIONSORT_ENABLED_DESCENDING         SecurityMonitoringSuppressionSort = "-enabled"
)

var allowedSecurityMonitoringSuppressionSortEnumValues = []SecurityMonitoringSuppressionSort{
	SECURITYMONITORINGSUPPRESSIONSORT_NAME,
	SECURITYMONITORINGSUPPRESSIONSORT_START_DATE,
	SECURITYMONITORINGSUPPRESSIONSORT_EXPIRATION_DATE,
	SECURITYMONITORINGSUPPRESSIONSORT_UPDATE_DATE,
	SECURITYMONITORINGSUPPRESSIONSORT_ENABLED,
	SECURITYMONITORINGSUPPRESSIONSORT_NAME_DESCENDING,
	SECURITYMONITORINGSUPPRESSIONSORT_START_DATE_DESCENDING,
	SECURITYMONITORINGSUPPRESSIONSORT_EXPIRATION_DATE_DESCENDING,
	SECURITYMONITORINGSUPPRESSIONSORT_UPDATE_DATE_DESCENDING,
	SECURITYMONITORINGSUPPRESSIONSORT_CREATION_DATE_DESCENDING,
	SECURITYMONITORINGSUPPRESSIONSORT_ENABLED_DESCENDING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringSuppressionSort) GetAllowedValues() []SecurityMonitoringSuppressionSort {
	return allowedSecurityMonitoringSuppressionSortEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringSuppressionSort) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringSuppressionSort(value)
	return nil
}

// NewSecurityMonitoringSuppressionSortFromValue returns a pointer to a valid SecurityMonitoringSuppressionSort
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringSuppressionSortFromValue(v string) (*SecurityMonitoringSuppressionSort, error) {
	ev := SecurityMonitoringSuppressionSort(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringSuppressionSort: valid values are %v", v, allowedSecurityMonitoringSuppressionSortEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringSuppressionSort) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringSuppressionSortEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringSuppressionSort value.
func (v SecurityMonitoringSuppressionSort) Ptr() *SecurityMonitoringSuppressionSort {
	return &v
}
