// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineDatadogMetricsDestinationType The destination type. The value should always be `datadog_metrics`.
type ObservabilityPipelineDatadogMetricsDestinationType string

// List of ObservabilityPipelineDatadogMetricsDestinationType.
const (
	OBSERVABILITYPIPELINEDATADOGMETRICSDESTINATIONTYPE_DATADOG_METRICS ObservabilityPipelineDatadogMetricsDestinationType = "datadog_metrics"
)

var allowedObservabilityPipelineDatadogMetricsDestinationTypeEnumValues = []ObservabilityPipelineDatadogMetricsDestinationType{
	OBSERVABILITYPIPELINEDATADOGMETRICSDESTINATIONTYPE_DATADOG_METRICS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineDatadogMetricsDestinationType) GetAllowedValues() []ObservabilityPipelineDatadogMetricsDestinationType {
	return allowedObservabilityPipelineDatadogMetricsDestinationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineDatadogMetricsDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineDatadogMetricsDestinationType(value)
	return nil
}

// NewObservabilityPipelineDatadogMetricsDestinationTypeFromValue returns a pointer to a valid ObservabilityPipelineDatadogMetricsDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineDatadogMetricsDestinationTypeFromValue(v string) (*ObservabilityPipelineDatadogMetricsDestinationType, error) {
	ev := ObservabilityPipelineDatadogMetricsDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineDatadogMetricsDestinationType: valid values are %v", v, allowedObservabilityPipelineDatadogMetricsDestinationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineDatadogMetricsDestinationType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineDatadogMetricsDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineDatadogMetricsDestinationType value.
func (v ObservabilityPipelineDatadogMetricsDestinationType) Ptr() *ObservabilityPipelineDatadogMetricsDestinationType {
	return &v
}
