// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReportScheduleListResponsePaginationType The pagination type.
type ReportScheduleListResponsePaginationType string

// List of ReportScheduleListResponsePaginationType.
const (
	REPORTSCHEDULELISTRESPONSEPAGINATIONTYPE_OFFSET_LIMIT ReportScheduleListResponsePaginationType = "offset_limit"
)

var allowedReportScheduleListResponsePaginationTypeEnumValues = []ReportScheduleListResponsePaginationType{
	REPORTSCHEDULELISTRESPONSEPAGINATIONTYPE_OFFSET_LIMIT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ReportScheduleListResponsePaginationType) GetAllowedValues() []ReportScheduleListResponsePaginationType {
	return allowedReportScheduleListResponsePaginationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ReportScheduleListResponsePaginationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ReportScheduleListResponsePaginationType(value)
	return nil
}

// NewReportScheduleListResponsePaginationTypeFromValue returns a pointer to a valid ReportScheduleListResponsePaginationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewReportScheduleListResponsePaginationTypeFromValue(v string) (*ReportScheduleListResponsePaginationType, error) {
	ev := ReportScheduleListResponsePaginationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ReportScheduleListResponsePaginationType: valid values are %v", v, allowedReportScheduleListResponsePaginationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ReportScheduleListResponsePaginationType) IsValid() bool {
	for _, existing := range allowedReportScheduleListResponsePaginationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReportScheduleListResponsePaginationType value.
func (v ReportScheduleListResponsePaginationType) Ptr() *ReportScheduleListResponsePaginationType {
	return &v
}
