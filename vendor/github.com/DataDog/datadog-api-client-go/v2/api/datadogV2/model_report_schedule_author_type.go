// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReportScheduleAuthorType JSON:API resource type for the included report author.
type ReportScheduleAuthorType string

// List of ReportScheduleAuthorType.
const (
	REPORTSCHEDULEAUTHORTYPE_USERS ReportScheduleAuthorType = "users"
)

var allowedReportScheduleAuthorTypeEnumValues = []ReportScheduleAuthorType{
	REPORTSCHEDULEAUTHORTYPE_USERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ReportScheduleAuthorType) GetAllowedValues() []ReportScheduleAuthorType {
	return allowedReportScheduleAuthorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ReportScheduleAuthorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ReportScheduleAuthorType(value)
	return nil
}

// NewReportScheduleAuthorTypeFromValue returns a pointer to a valid ReportScheduleAuthorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewReportScheduleAuthorTypeFromValue(v string) (*ReportScheduleAuthorType, error) {
	ev := ReportScheduleAuthorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ReportScheduleAuthorType: valid values are %v", v, allowedReportScheduleAuthorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ReportScheduleAuthorType) IsValid() bool {
	for _, existing := range allowedReportScheduleAuthorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReportScheduleAuthorType value.
func (v ReportScheduleAuthorType) Ptr() *ReportScheduleAuthorType {
	return &v
}
