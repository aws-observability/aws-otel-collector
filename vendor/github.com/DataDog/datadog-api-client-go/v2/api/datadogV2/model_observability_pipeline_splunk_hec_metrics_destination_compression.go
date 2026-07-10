// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSplunkHecMetricsDestinationCompression Compression algorithm applied when sending metrics to Splunk HEC.
type ObservabilityPipelineSplunkHecMetricsDestinationCompression string

// List of ObservabilityPipelineSplunkHecMetricsDestinationCompression.
const (
	OBSERVABILITYPIPELINESPLUNKHECMETRICSDESTINATIONCOMPRESSION_NONE ObservabilityPipelineSplunkHecMetricsDestinationCompression = "none"
	OBSERVABILITYPIPELINESPLUNKHECMETRICSDESTINATIONCOMPRESSION_GZIP ObservabilityPipelineSplunkHecMetricsDestinationCompression = "gzip"
)

var allowedObservabilityPipelineSplunkHecMetricsDestinationCompressionEnumValues = []ObservabilityPipelineSplunkHecMetricsDestinationCompression{
	OBSERVABILITYPIPELINESPLUNKHECMETRICSDESTINATIONCOMPRESSION_NONE,
	OBSERVABILITYPIPELINESPLUNKHECMETRICSDESTINATIONCOMPRESSION_GZIP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineSplunkHecMetricsDestinationCompression) GetAllowedValues() []ObservabilityPipelineSplunkHecMetricsDestinationCompression {
	return allowedObservabilityPipelineSplunkHecMetricsDestinationCompressionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineSplunkHecMetricsDestinationCompression) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineSplunkHecMetricsDestinationCompression(value)
	return nil
}

// NewObservabilityPipelineSplunkHecMetricsDestinationCompressionFromValue returns a pointer to a valid ObservabilityPipelineSplunkHecMetricsDestinationCompression
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineSplunkHecMetricsDestinationCompressionFromValue(v string) (*ObservabilityPipelineSplunkHecMetricsDestinationCompression, error) {
	ev := ObservabilityPipelineSplunkHecMetricsDestinationCompression(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineSplunkHecMetricsDestinationCompression: valid values are %v", v, allowedObservabilityPipelineSplunkHecMetricsDestinationCompressionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineSplunkHecMetricsDestinationCompression) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineSplunkHecMetricsDestinationCompressionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineSplunkHecMetricsDestinationCompression value.
func (v ObservabilityPipelineSplunkHecMetricsDestinationCompression) Ptr() *ObservabilityPipelineSplunkHecMetricsDestinationCompression {
	return &v
}
