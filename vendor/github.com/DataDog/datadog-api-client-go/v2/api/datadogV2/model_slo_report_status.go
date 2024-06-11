// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SLOReportStatus The status of the SLO report job.
type SLOReportStatus string

// List of SLOReportStatus.
const (
	SLOREPORTSTATUS_IN_PROGRESS           SLOReportStatus = "in_progress"
	SLOREPORTSTATUS_COMPLETED             SLOReportStatus = "completed"
	SLOREPORTSTATUS_COMPLETED_WITH_ERRORS SLOReportStatus = "completed_with_errors"
	SLOREPORTSTATUS_FAILED                SLOReportStatus = "failed"
)

var allowedSLOReportStatusEnumValues = []SLOReportStatus{
	SLOREPORTSTATUS_IN_PROGRESS,
	SLOREPORTSTATUS_COMPLETED,
	SLOREPORTSTATUS_COMPLETED_WITH_ERRORS,
	SLOREPORTSTATUS_FAILED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SLOReportStatus) GetAllowedValues() []SLOReportStatus {
	return allowedSLOReportStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SLOReportStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SLOReportStatus(value)
	return nil
}

// NewSLOReportStatusFromValue returns a pointer to a valid SLOReportStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSLOReportStatusFromValue(v string) (*SLOReportStatus, error) {
	ev := SLOReportStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SLOReportStatus: valid values are %v", v, allowedSLOReportStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SLOReportStatus) IsValid() bool {
	for _, existing := range allowedSLOReportStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SLOReportStatus value.
func (v SLOReportStatus) Ptr() *SLOReportStatus {
	return &v
}
