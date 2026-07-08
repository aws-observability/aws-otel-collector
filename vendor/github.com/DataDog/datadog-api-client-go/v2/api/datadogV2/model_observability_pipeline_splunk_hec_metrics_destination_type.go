// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSplunkHecMetricsDestinationType The destination type. Always `splunk_hec_metrics`.
type ObservabilityPipelineSplunkHecMetricsDestinationType string

// List of ObservabilityPipelineSplunkHecMetricsDestinationType.
const (
	OBSERVABILITYPIPELINESPLUNKHECMETRICSDESTINATIONTYPE_SPLUNK_HEC_METRICS ObservabilityPipelineSplunkHecMetricsDestinationType = "splunk_hec_metrics"
)

var allowedObservabilityPipelineSplunkHecMetricsDestinationTypeEnumValues = []ObservabilityPipelineSplunkHecMetricsDestinationType{
	OBSERVABILITYPIPELINESPLUNKHECMETRICSDESTINATIONTYPE_SPLUNK_HEC_METRICS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineSplunkHecMetricsDestinationType) GetAllowedValues() []ObservabilityPipelineSplunkHecMetricsDestinationType {
	return allowedObservabilityPipelineSplunkHecMetricsDestinationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineSplunkHecMetricsDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineSplunkHecMetricsDestinationType(value)
	return nil
}

// NewObservabilityPipelineSplunkHecMetricsDestinationTypeFromValue returns a pointer to a valid ObservabilityPipelineSplunkHecMetricsDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineSplunkHecMetricsDestinationTypeFromValue(v string) (*ObservabilityPipelineSplunkHecMetricsDestinationType, error) {
	ev := ObservabilityPipelineSplunkHecMetricsDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineSplunkHecMetricsDestinationType: valid values are %v", v, allowedObservabilityPipelineSplunkHecMetricsDestinationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineSplunkHecMetricsDestinationType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineSplunkHecMetricsDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineSplunkHecMetricsDestinationType value.
func (v ObservabilityPipelineSplunkHecMetricsDestinationType) Ptr() *ObservabilityPipelineSplunkHecMetricsDestinationType {
	return &v
}
