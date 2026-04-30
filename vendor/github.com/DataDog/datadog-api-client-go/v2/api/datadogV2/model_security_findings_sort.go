// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityFindingsSort The sort parameters when querying security findings.
type SecurityFindingsSort string

// List of SecurityFindingsSort.
const (
	SECURITYFINDINGSSORT_DETECTION_CHANGED_AT_ASC  SecurityFindingsSort = "@detection_changed_at"
	SECURITYFINDINGSSORT_DETECTION_CHANGED_AT_DESC SecurityFindingsSort = "-@detection_changed_at"
)

var allowedSecurityFindingsSortEnumValues = []SecurityFindingsSort{
	SECURITYFINDINGSSORT_DETECTION_CHANGED_AT_ASC,
	SECURITYFINDINGSSORT_DETECTION_CHANGED_AT_DESC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityFindingsSort) GetAllowedValues() []SecurityFindingsSort {
	return allowedSecurityFindingsSortEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityFindingsSort) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityFindingsSort(value)
	return nil
}

// NewSecurityFindingsSortFromValue returns a pointer to a valid SecurityFindingsSort
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityFindingsSortFromValue(v string) (*SecurityFindingsSort, error) {
	ev := SecurityFindingsSort(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityFindingsSort: valid values are %v", v, allowedSecurityFindingsSortEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityFindingsSort) IsValid() bool {
	for _, existing := range allowedSecurityFindingsSortEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityFindingsSort value.
func (v SecurityFindingsSort) Ptr() *SecurityFindingsSort {
	return &v
}
