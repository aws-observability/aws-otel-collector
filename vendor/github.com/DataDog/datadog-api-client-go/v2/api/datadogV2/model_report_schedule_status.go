// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReportScheduleStatus Whether the schedule is currently delivering reports (`active`) or paused (`inactive`).
type ReportScheduleStatus string

// List of ReportScheduleStatus.
const (
	REPORTSCHEDULESTATUS_ACTIVE   ReportScheduleStatus = "active"
	REPORTSCHEDULESTATUS_INACTIVE ReportScheduleStatus = "inactive"
)

var allowedReportScheduleStatusEnumValues = []ReportScheduleStatus{
	REPORTSCHEDULESTATUS_ACTIVE,
	REPORTSCHEDULESTATUS_INACTIVE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ReportScheduleStatus) GetAllowedValues() []ReportScheduleStatus {
	return allowedReportScheduleStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ReportScheduleStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ReportScheduleStatus(value)
	return nil
}

// NewReportScheduleStatusFromValue returns a pointer to a valid ReportScheduleStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewReportScheduleStatusFromValue(v string) (*ReportScheduleStatus, error) {
	ev := ReportScheduleStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ReportScheduleStatus: valid values are %v", v, allowedReportScheduleStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ReportScheduleStatus) IsValid() bool {
	for _, existing := range allowedReportScheduleStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReportScheduleStatus value.
func (v ReportScheduleStatus) Ptr() *ReportScheduleStatus {
	return &v
}
