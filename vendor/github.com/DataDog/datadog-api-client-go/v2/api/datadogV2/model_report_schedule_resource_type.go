// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReportScheduleResourceType The type of dashboard resource the report schedule targets.
type ReportScheduleResourceType string

// List of ReportScheduleResourceType.
const (
	REPORTSCHEDULERESOURCETYPE_DASHBOARD             ReportScheduleResourceType = "dashboard"
	REPORTSCHEDULERESOURCETYPE_INTEGRATION_DASHBOARD ReportScheduleResourceType = "integration_dashboard"
)

var allowedReportScheduleResourceTypeEnumValues = []ReportScheduleResourceType{
	REPORTSCHEDULERESOURCETYPE_DASHBOARD,
	REPORTSCHEDULERESOURCETYPE_INTEGRATION_DASHBOARD,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ReportScheduleResourceType) GetAllowedValues() []ReportScheduleResourceType {
	return allowedReportScheduleResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ReportScheduleResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ReportScheduleResourceType(value)
	return nil
}

// NewReportScheduleResourceTypeFromValue returns a pointer to a valid ReportScheduleResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewReportScheduleResourceTypeFromValue(v string) (*ReportScheduleResourceType, error) {
	ev := ReportScheduleResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ReportScheduleResourceType: valid values are %v", v, allowedReportScheduleResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ReportScheduleResourceType) IsValid() bool {
	for _, existing := range allowedReportScheduleResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReportScheduleResourceType value.
func (v ReportScheduleResourceType) Ptr() *ReportScheduleResourceType {
	return &v
}
