// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetType Widget types that are allowed to be stored as individual records.
// This is not a complete list of dashboard and notebook widget types.
type WidgetType string

// List of WidgetType.
const (
	WIDGETTYPE_BAR_CHART          WidgetType = "bar_chart"
	WIDGETTYPE_CHANGE             WidgetType = "change"
	WIDGETTYPE_CLOUD_COST_SUMMARY WidgetType = "cloud_cost_summary"
	WIDGETTYPE_COHORT             WidgetType = "cohort"
	WIDGETTYPE_FUNNEL             WidgetType = "funnel"
	WIDGETTYPE_GEOMAP             WidgetType = "geomap"
	WIDGETTYPE_LIST_STREAM        WidgetType = "list_stream"
	WIDGETTYPE_QUERY_TABLE        WidgetType = "query_table"
	WIDGETTYPE_QUERY_VALUE        WidgetType = "query_value"
	WIDGETTYPE_RETENTION_CURVE    WidgetType = "retention_curve"
	WIDGETTYPE_SANKEY             WidgetType = "sankey"
	WIDGETTYPE_SUNBURST           WidgetType = "sunburst"
	WIDGETTYPE_TIMESERIES         WidgetType = "timeseries"
	WIDGETTYPE_TOPLIST            WidgetType = "toplist"
	WIDGETTYPE_TREEMAP            WidgetType = "treemap"
)

var allowedWidgetTypeEnumValues = []WidgetType{
	WIDGETTYPE_BAR_CHART,
	WIDGETTYPE_CHANGE,
	WIDGETTYPE_CLOUD_COST_SUMMARY,
	WIDGETTYPE_COHORT,
	WIDGETTYPE_FUNNEL,
	WIDGETTYPE_GEOMAP,
	WIDGETTYPE_LIST_STREAM,
	WIDGETTYPE_QUERY_TABLE,
	WIDGETTYPE_QUERY_VALUE,
	WIDGETTYPE_RETENTION_CURVE,
	WIDGETTYPE_SANKEY,
	WIDGETTYPE_SUNBURST,
	WIDGETTYPE_TIMESERIES,
	WIDGETTYPE_TOPLIST,
	WIDGETTYPE_TREEMAP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *WidgetType) GetAllowedValues() []WidgetType {
	return allowedWidgetTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WidgetType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WidgetType(value)
	return nil
}

// NewWidgetTypeFromValue returns a pointer to a valid WidgetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWidgetTypeFromValue(v string) (*WidgetType, error) {
	ev := WidgetType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WidgetType: valid values are %v", v, allowedWidgetTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WidgetType) IsValid() bool {
	for _, existing := range allowedWidgetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WidgetType value.
func (v WidgetType) Ptr() *WidgetType {
	return &v
}
