// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityFindingsStatus The status of the response.
type SecurityFindingsStatus string

// List of SecurityFindingsStatus.
const (
	SECURITYFINDINGSSTATUS_DONE    SecurityFindingsStatus = "done"
	SECURITYFINDINGSSTATUS_TIMEOUT SecurityFindingsStatus = "timeout"
)

var allowedSecurityFindingsStatusEnumValues = []SecurityFindingsStatus{
	SECURITYFINDINGSSTATUS_DONE,
	SECURITYFINDINGSSTATUS_TIMEOUT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityFindingsStatus) GetAllowedValues() []SecurityFindingsStatus {
	return allowedSecurityFindingsStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityFindingsStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityFindingsStatus(value)
	return nil
}

// NewSecurityFindingsStatusFromValue returns a pointer to a valid SecurityFindingsStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityFindingsStatusFromValue(v string) (*SecurityFindingsStatus, error) {
	ev := SecurityFindingsStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityFindingsStatus: valid values are %v", v, allowedSecurityFindingsStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityFindingsStatus) IsValid() bool {
	for _, existing := range allowedSecurityFindingsStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityFindingsStatus value.
func (v SecurityFindingsStatus) Ptr() *SecurityFindingsStatus {
	return &v
}
