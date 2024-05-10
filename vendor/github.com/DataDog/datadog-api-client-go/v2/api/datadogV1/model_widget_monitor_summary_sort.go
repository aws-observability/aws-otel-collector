// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetMonitorSummarySort Widget sorting methods.
type WidgetMonitorSummarySort string

// List of WidgetMonitorSummarySort.
const (
	WIDGETMONITORSUMMARYSORT_NAME                 WidgetMonitorSummarySort = "name"
	WIDGETMONITORSUMMARYSORT_GROUP                WidgetMonitorSummarySort = "group"
	WIDGETMONITORSUMMARYSORT_STATUS               WidgetMonitorSummarySort = "status"
	WIDGETMONITORSUMMARYSORT_TAGS                 WidgetMonitorSummarySort = "tags"
	WIDGETMONITORSUMMARYSORT_TRIGGERED            WidgetMonitorSummarySort = "triggered"
	WIDGETMONITORSUMMARYSORT_GROUP_ASCENDING      WidgetMonitorSummarySort = "group,asc"
	WIDGETMONITORSUMMARYSORT_GROUP_DESCENDING     WidgetMonitorSummarySort = "group,desc"
	WIDGETMONITORSUMMARYSORT_NAME_ASCENDING       WidgetMonitorSummarySort = "name,asc"
	WIDGETMONITORSUMMARYSORT_NAME_DESCENDING      WidgetMonitorSummarySort = "name,desc"
	WIDGETMONITORSUMMARYSORT_STATUS_ASCENDING     WidgetMonitorSummarySort = "status,asc"
	WIDGETMONITORSUMMARYSORT_STATUS_DESCENDING    WidgetMonitorSummarySort = "status,desc"
	WIDGETMONITORSUMMARYSORT_TAGS_ASCENDING       WidgetMonitorSummarySort = "tags,asc"
	WIDGETMONITORSUMMARYSORT_TAGS_DESCENDING      WidgetMonitorSummarySort = "tags,desc"
	WIDGETMONITORSUMMARYSORT_TRIGGERED_ASCENDING  WidgetMonitorSummarySort = "triggered,asc"
	WIDGETMONITORSUMMARYSORT_TRIGGERED_DESCENDING WidgetMonitorSummarySort = "triggered,desc"
	WIDGETMONITORSUMMARYSORT_PRIORITY_ASCENDING   WidgetMonitorSummarySort = "priority,asc"
	WIDGETMONITORSUMMARYSORT_PRIORITY_DESCENDING  WidgetMonitorSummarySort = "priority,desc"
)

var allowedWidgetMonitorSummarySortEnumValues = []WidgetMonitorSummarySort{
	WIDGETMONITORSUMMARYSORT_NAME,
	WIDGETMONITORSUMMARYSORT_GROUP,
	WIDGETMONITORSUMMARYSORT_STATUS,
	WIDGETMONITORSUMMARYSORT_TAGS,
	WIDGETMONITORSUMMARYSORT_TRIGGERED,
	WIDGETMONITORSUMMARYSORT_GROUP_ASCENDING,
	WIDGETMONITORSUMMARYSORT_GROUP_DESCENDING,
	WIDGETMONITORSUMMARYSORT_NAME_ASCENDING,
	WIDGETMONITORSUMMARYSORT_NAME_DESCENDING,
	WIDGETMONITORSUMMARYSORT_STATUS_ASCENDING,
	WIDGETMONITORSUMMARYSORT_STATUS_DESCENDING,
	WIDGETMONITORSUMMARYSORT_TAGS_ASCENDING,
	WIDGETMONITORSUMMARYSORT_TAGS_DESCENDING,
	WIDGETMONITORSUMMARYSORT_TRIGGERED_ASCENDING,
	WIDGETMONITORSUMMARYSORT_TRIGGERED_DESCENDING,
	WIDGETMONITORSUMMARYSORT_PRIORITY_ASCENDING,
	WIDGETMONITORSUMMARYSORT_PRIORITY_DESCENDING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WidgetMonitorSummarySort) GetAllowedValues() []WidgetMonitorSummarySort {
	return allowedWidgetMonitorSummarySortEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WidgetMonitorSummarySort) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WidgetMonitorSummarySort(value)
	return nil
}

// NewWidgetMonitorSummarySortFromValue returns a pointer to a valid WidgetMonitorSummarySort
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWidgetMonitorSummarySortFromValue(v string) (*WidgetMonitorSummarySort, error) {
	ev := WidgetMonitorSummarySort(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WidgetMonitorSummarySort: valid values are %v", v, allowedWidgetMonitorSummarySortEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WidgetMonitorSummarySort) IsValid() bool {
	for _, existing := range allowedWidgetMonitorSummarySortEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WidgetMonitorSummarySort value.
func (v WidgetMonitorSummarySort) Ptr() *WidgetMonitorSummarySort {
	return &v
}
