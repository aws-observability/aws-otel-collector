// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReportScheduleIncludedResourceType JSON:API resource type for an included report resource.
type ReportScheduleIncludedResourceType string

// List of ReportScheduleIncludedResourceType.
const (
	REPORTSCHEDULEINCLUDEDRESOURCETYPE_RESOURCE ReportScheduleIncludedResourceType = "resource"
)

var allowedReportScheduleIncludedResourceTypeEnumValues = []ReportScheduleIncludedResourceType{
	REPORTSCHEDULEINCLUDEDRESOURCETYPE_RESOURCE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ReportScheduleIncludedResourceType) GetAllowedValues() []ReportScheduleIncludedResourceType {
	return allowedReportScheduleIncludedResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ReportScheduleIncludedResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ReportScheduleIncludedResourceType(value)
	return nil
}

// NewReportScheduleIncludedResourceTypeFromValue returns a pointer to a valid ReportScheduleIncludedResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewReportScheduleIncludedResourceTypeFromValue(v string) (*ReportScheduleIncludedResourceType, error) {
	ev := ReportScheduleIncludedResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ReportScheduleIncludedResourceType: valid values are %v", v, allowedReportScheduleIncludedResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ReportScheduleIncludedResourceType) IsValid() bool {
	for _, existing := range allowedReportScheduleIncludedResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReportScheduleIncludedResourceType value.
func (v ReportScheduleIncludedResourceType) Ptr() *ReportScheduleIncludedResourceType {
	return &v
}
