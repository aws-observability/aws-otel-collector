// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DataObservabilityMonitorRunStatus The status of a data observability monitor run.
type DataObservabilityMonitorRunStatus string

// List of DataObservabilityMonitorRunStatus.
const (
	DATAOBSERVABILITYMONITORRUNSTATUS_PENDING DataObservabilityMonitorRunStatus = "pending"
	DATAOBSERVABILITYMONITORRUNSTATUS_OK      DataObservabilityMonitorRunStatus = "ok"
	DATAOBSERVABILITYMONITORRUNSTATUS_WARN    DataObservabilityMonitorRunStatus = "warn"
	DATAOBSERVABILITYMONITORRUNSTATUS_ALERT   DataObservabilityMonitorRunStatus = "alert"
	DATAOBSERVABILITYMONITORRUNSTATUS_ERROR   DataObservabilityMonitorRunStatus = "error"
)

var allowedDataObservabilityMonitorRunStatusEnumValues = []DataObservabilityMonitorRunStatus{
	DATAOBSERVABILITYMONITORRUNSTATUS_PENDING,
	DATAOBSERVABILITYMONITORRUNSTATUS_OK,
	DATAOBSERVABILITYMONITORRUNSTATUS_WARN,
	DATAOBSERVABILITYMONITORRUNSTATUS_ALERT,
	DATAOBSERVABILITYMONITORRUNSTATUS_ERROR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DataObservabilityMonitorRunStatus) GetAllowedValues() []DataObservabilityMonitorRunStatus {
	return allowedDataObservabilityMonitorRunStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DataObservabilityMonitorRunStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DataObservabilityMonitorRunStatus(value)
	return nil
}

// NewDataObservabilityMonitorRunStatusFromValue returns a pointer to a valid DataObservabilityMonitorRunStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDataObservabilityMonitorRunStatusFromValue(v string) (*DataObservabilityMonitorRunStatus, error) {
	ev := DataObservabilityMonitorRunStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DataObservabilityMonitorRunStatus: valid values are %v", v, allowedDataObservabilityMonitorRunStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DataObservabilityMonitorRunStatus) IsValid() bool {
	for _, existing := range allowedDataObservabilityMonitorRunStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataObservabilityMonitorRunStatus value.
func (v DataObservabilityMonitorRunStatus) Ptr() *DataObservabilityMonitorRunStatus {
	return &v
}
