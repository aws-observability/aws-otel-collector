// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorType The type of the monitor. For more information about `type`, see the [monitor options](https://docs.datadoghq.com/monitors/guide/monitor_api_options/) docs.
type MonitorType string

// List of MonitorType.
const (
	MONITORTYPE_COMPOSITE                 MonitorType = "composite"
	MONITORTYPE_EVENT_ALERT               MonitorType = "event alert"
	MONITORTYPE_LOG_ALERT                 MonitorType = "log alert"
	MONITORTYPE_METRIC_ALERT              MonitorType = "metric alert"
	MONITORTYPE_PROCESS_ALERT             MonitorType = "process alert"
	MONITORTYPE_QUERY_ALERT               MonitorType = "query alert"
	MONITORTYPE_RUM_ALERT                 MonitorType = "rum alert"
	MONITORTYPE_SERVICE_CHECK             MonitorType = "service check"
	MONITORTYPE_SYNTHETICS_ALERT          MonitorType = "synthetics alert"
	MONITORTYPE_TRACE_ANALYTICS_ALERT     MonitorType = "trace-analytics alert"
	MONITORTYPE_SLO_ALERT                 MonitorType = "slo alert"
	MONITORTYPE_EVENT_V2_ALERT            MonitorType = "event-v2 alert"
	MONITORTYPE_AUDIT_ALERT               MonitorType = "audit alert"
	MONITORTYPE_CI_PIPELINES_ALERT        MonitorType = "ci-pipelines alert"
	MONITORTYPE_CI_TESTS_ALERT            MonitorType = "ci-tests alert"
	MONITORTYPE_ERROR_TRACKING_ALERT      MonitorType = "error-tracking alert"
	MONITORTYPE_DATABASE_MONITORING_ALERT MonitorType = "database-monitoring alert"
	MONITORTYPE_NETWORK_PERFORMANCE_ALERT MonitorType = "network-performance alert"
	MONITORTYPE_COST_ALERT                MonitorType = "cost alert"
)

var allowedMonitorTypeEnumValues = []MonitorType{
	MONITORTYPE_COMPOSITE,
	MONITORTYPE_EVENT_ALERT,
	MONITORTYPE_LOG_ALERT,
	MONITORTYPE_METRIC_ALERT,
	MONITORTYPE_PROCESS_ALERT,
	MONITORTYPE_QUERY_ALERT,
	MONITORTYPE_RUM_ALERT,
	MONITORTYPE_SERVICE_CHECK,
	MONITORTYPE_SYNTHETICS_ALERT,
	MONITORTYPE_TRACE_ANALYTICS_ALERT,
	MONITORTYPE_SLO_ALERT,
	MONITORTYPE_EVENT_V2_ALERT,
	MONITORTYPE_AUDIT_ALERT,
	MONITORTYPE_CI_PIPELINES_ALERT,
	MONITORTYPE_CI_TESTS_ALERT,
	MONITORTYPE_ERROR_TRACKING_ALERT,
	MONITORTYPE_DATABASE_MONITORING_ALERT,
	MONITORTYPE_NETWORK_PERFORMANCE_ALERT,
	MONITORTYPE_COST_ALERT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MonitorType) GetAllowedValues() []MonitorType {
	return allowedMonitorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MonitorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MonitorType(value)
	return nil
}

// NewMonitorTypeFromValue returns a pointer to a valid MonitorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMonitorTypeFromValue(v string) (*MonitorType, error) {
	ev := MonitorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MonitorType: valid values are %v", v, allowedMonitorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MonitorType) IsValid() bool {
	for _, existing := range allowedMonitorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MonitorType value.
func (v MonitorType) Ptr() *MonitorType {
	return &v
}
