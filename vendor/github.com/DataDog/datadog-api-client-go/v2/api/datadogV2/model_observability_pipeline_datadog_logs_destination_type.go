// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineDatadogLogsDestinationType The destination type. The value should always be `datadog_logs`.
type ObservabilityPipelineDatadogLogsDestinationType string

// List of ObservabilityPipelineDatadogLogsDestinationType.
const (
	OBSERVABILITYPIPELINEDATADOGLOGSDESTINATIONTYPE_DATADOG_LOGS ObservabilityPipelineDatadogLogsDestinationType = "datadog_logs"
)

var allowedObservabilityPipelineDatadogLogsDestinationTypeEnumValues = []ObservabilityPipelineDatadogLogsDestinationType{
	OBSERVABILITYPIPELINEDATADOGLOGSDESTINATIONTYPE_DATADOG_LOGS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineDatadogLogsDestinationType) GetAllowedValues() []ObservabilityPipelineDatadogLogsDestinationType {
	return allowedObservabilityPipelineDatadogLogsDestinationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineDatadogLogsDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineDatadogLogsDestinationType(value)
	return nil
}

// NewObservabilityPipelineDatadogLogsDestinationTypeFromValue returns a pointer to a valid ObservabilityPipelineDatadogLogsDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineDatadogLogsDestinationTypeFromValue(v string) (*ObservabilityPipelineDatadogLogsDestinationType, error) {
	ev := ObservabilityPipelineDatadogLogsDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineDatadogLogsDestinationType: valid values are %v", v, allowedObservabilityPipelineDatadogLogsDestinationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineDatadogLogsDestinationType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineDatadogLogsDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineDatadogLogsDestinationType value.
func (v ObservabilityPipelineDatadogLogsDestinationType) Ptr() *ObservabilityPipelineDatadogLogsDestinationType {
	return &v
}
