// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardGlobalTimeLiveSpan Dashboard global time live_span selection
type DashboardGlobalTimeLiveSpan string

// List of DashboardGlobalTimeLiveSpan.
const (
	DASHBOARDGLOBALTIMELIVESPAN_PAST_FIFTEEN_MINUTES DashboardGlobalTimeLiveSpan = "15m"
	DASHBOARDGLOBALTIMELIVESPAN_PAST_ONE_HOUR        DashboardGlobalTimeLiveSpan = "1h"
	DASHBOARDGLOBALTIMELIVESPAN_PAST_FOUR_HOURS      DashboardGlobalTimeLiveSpan = "4h"
	DASHBOARDGLOBALTIMELIVESPAN_PAST_ONE_DAY         DashboardGlobalTimeLiveSpan = "1d"
	DASHBOARDGLOBALTIMELIVESPAN_PAST_TWO_DAYS        DashboardGlobalTimeLiveSpan = "2d"
	DASHBOARDGLOBALTIMELIVESPAN_PAST_ONE_WEEK        DashboardGlobalTimeLiveSpan = "1w"
	DASHBOARDGLOBALTIMELIVESPAN_PAST_ONE_MONTH       DashboardGlobalTimeLiveSpan = "1mo"
	DASHBOARDGLOBALTIMELIVESPAN_PAST_THREE_MONTHS    DashboardGlobalTimeLiveSpan = "3mo"
)

var allowedDashboardGlobalTimeLiveSpanEnumValues = []DashboardGlobalTimeLiveSpan{
	DASHBOARDGLOBALTIMELIVESPAN_PAST_FIFTEEN_MINUTES,
	DASHBOARDGLOBALTIMELIVESPAN_PAST_ONE_HOUR,
	DASHBOARDGLOBALTIMELIVESPAN_PAST_FOUR_HOURS,
	DASHBOARDGLOBALTIMELIVESPAN_PAST_ONE_DAY,
	DASHBOARDGLOBALTIMELIVESPAN_PAST_TWO_DAYS,
	DASHBOARDGLOBALTIMELIVESPAN_PAST_ONE_WEEK,
	DASHBOARDGLOBALTIMELIVESPAN_PAST_ONE_MONTH,
	DASHBOARDGLOBALTIMELIVESPAN_PAST_THREE_MONTHS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DashboardGlobalTimeLiveSpan) GetAllowedValues() []DashboardGlobalTimeLiveSpan {
	return allowedDashboardGlobalTimeLiveSpanEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DashboardGlobalTimeLiveSpan) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DashboardGlobalTimeLiveSpan(value)
	return nil
}

// NewDashboardGlobalTimeLiveSpanFromValue returns a pointer to a valid DashboardGlobalTimeLiveSpan
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDashboardGlobalTimeLiveSpanFromValue(v string) (*DashboardGlobalTimeLiveSpan, error) {
	ev := DashboardGlobalTimeLiveSpan(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DashboardGlobalTimeLiveSpan: valid values are %v", v, allowedDashboardGlobalTimeLiveSpanEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DashboardGlobalTimeLiveSpan) IsValid() bool {
	for _, existing := range allowedDashboardGlobalTimeLiveSpanEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardGlobalTimeLiveSpan value.
func (v DashboardGlobalTimeLiveSpan) Ptr() *DashboardGlobalTimeLiveSpan {
	return &v
}
