// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DataObservabilityMonitorRunType The JSON:API resource type for a data observability monitor run.
type DataObservabilityMonitorRunType string

// List of DataObservabilityMonitorRunType.
const (
	DATAOBSERVABILITYMONITORRUNTYPE_MONITOR_RUN DataObservabilityMonitorRunType = "monitor_run"
)

var allowedDataObservabilityMonitorRunTypeEnumValues = []DataObservabilityMonitorRunType{
	DATAOBSERVABILITYMONITORRUNTYPE_MONITOR_RUN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DataObservabilityMonitorRunType) GetAllowedValues() []DataObservabilityMonitorRunType {
	return allowedDataObservabilityMonitorRunTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DataObservabilityMonitorRunType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DataObservabilityMonitorRunType(value)
	return nil
}

// NewDataObservabilityMonitorRunTypeFromValue returns a pointer to a valid DataObservabilityMonitorRunType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDataObservabilityMonitorRunTypeFromValue(v string) (*DataObservabilityMonitorRunType, error) {
	ev := DataObservabilityMonitorRunType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DataObservabilityMonitorRunType: valid values are %v", v, allowedDataObservabilityMonitorRunTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DataObservabilityMonitorRunType) IsValid() bool {
	for _, existing := range allowedDataObservabilityMonitorRunTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataObservabilityMonitorRunType value.
func (v DataObservabilityMonitorRunType) Ptr() *DataObservabilityMonitorRunType {
	return &v
}
