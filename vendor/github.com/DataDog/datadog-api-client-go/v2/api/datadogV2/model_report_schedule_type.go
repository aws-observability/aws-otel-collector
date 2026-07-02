// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReportScheduleType JSON:API resource type for report schedules.
type ReportScheduleType string

// List of ReportScheduleType.
const (
	REPORTSCHEDULETYPE_SCHEDULE ReportScheduleType = "schedule"
)

var allowedReportScheduleTypeEnumValues = []ReportScheduleType{
	REPORTSCHEDULETYPE_SCHEDULE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ReportScheduleType) GetAllowedValues() []ReportScheduleType {
	return allowedReportScheduleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ReportScheduleType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ReportScheduleType(value)
	return nil
}

// NewReportScheduleTypeFromValue returns a pointer to a valid ReportScheduleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewReportScheduleTypeFromValue(v string) (*ReportScheduleType, error) {
	ev := ReportScheduleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ReportScheduleType: valid values are %v", v, allowedReportScheduleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ReportScheduleType) IsValid() bool {
	for _, existing := range allowedReportScheduleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReportScheduleType value.
func (v ReportScheduleType) Ptr() *ReportScheduleType {
	return &v
}
