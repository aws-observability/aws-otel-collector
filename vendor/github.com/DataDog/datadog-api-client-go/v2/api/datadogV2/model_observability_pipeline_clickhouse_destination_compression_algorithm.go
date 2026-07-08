// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineClickhouseDestinationCompressionAlgorithm The compression algorithm applied to outbound HTTP requests.
type ObservabilityPipelineClickhouseDestinationCompressionAlgorithm string

// List of ObservabilityPipelineClickhouseDestinationCompressionAlgorithm.
const (
	OBSERVABILITYPIPELINECLICKHOUSEDESTINATIONCOMPRESSIONALGORITHM_GZIP ObservabilityPipelineClickhouseDestinationCompressionAlgorithm = "gzip"
	OBSERVABILITYPIPELINECLICKHOUSEDESTINATIONCOMPRESSIONALGORITHM_NONE ObservabilityPipelineClickhouseDestinationCompressionAlgorithm = "none"
)

var allowedObservabilityPipelineClickhouseDestinationCompressionAlgorithmEnumValues = []ObservabilityPipelineClickhouseDestinationCompressionAlgorithm{
	OBSERVABILITYPIPELINECLICKHOUSEDESTINATIONCOMPRESSIONALGORITHM_GZIP,
	OBSERVABILITYPIPELINECLICKHOUSEDESTINATIONCOMPRESSIONALGORITHM_NONE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineClickhouseDestinationCompressionAlgorithm) GetAllowedValues() []ObservabilityPipelineClickhouseDestinationCompressionAlgorithm {
	return allowedObservabilityPipelineClickhouseDestinationCompressionAlgorithmEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineClickhouseDestinationCompressionAlgorithm) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineClickhouseDestinationCompressionAlgorithm(value)
	return nil
}

// NewObservabilityPipelineClickhouseDestinationCompressionAlgorithmFromValue returns a pointer to a valid ObservabilityPipelineClickhouseDestinationCompressionAlgorithm
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineClickhouseDestinationCompressionAlgorithmFromValue(v string) (*ObservabilityPipelineClickhouseDestinationCompressionAlgorithm, error) {
	ev := ObservabilityPipelineClickhouseDestinationCompressionAlgorithm(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineClickhouseDestinationCompressionAlgorithm: valid values are %v", v, allowedObservabilityPipelineClickhouseDestinationCompressionAlgorithmEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineClickhouseDestinationCompressionAlgorithm) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineClickhouseDestinationCompressionAlgorithmEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineClickhouseDestinationCompressionAlgorithm value.
func (v ObservabilityPipelineClickhouseDestinationCompressionAlgorithm) Ptr() *ObservabilityPipelineClickhouseDestinationCompressionAlgorithm {
	return &v
}
