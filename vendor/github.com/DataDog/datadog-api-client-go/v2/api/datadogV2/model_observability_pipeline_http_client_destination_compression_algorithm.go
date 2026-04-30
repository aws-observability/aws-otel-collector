// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineHttpClientDestinationCompressionAlgorithm Compression algorithm.
type ObservabilityPipelineHttpClientDestinationCompressionAlgorithm string

// List of ObservabilityPipelineHttpClientDestinationCompressionAlgorithm.
const (
	OBSERVABILITYPIPELINEHTTPCLIENTDESTINATIONCOMPRESSIONALGORITHM_GZIP ObservabilityPipelineHttpClientDestinationCompressionAlgorithm = "gzip"
)

var allowedObservabilityPipelineHttpClientDestinationCompressionAlgorithmEnumValues = []ObservabilityPipelineHttpClientDestinationCompressionAlgorithm{
	OBSERVABILITYPIPELINEHTTPCLIENTDESTINATIONCOMPRESSIONALGORITHM_GZIP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineHttpClientDestinationCompressionAlgorithm) GetAllowedValues() []ObservabilityPipelineHttpClientDestinationCompressionAlgorithm {
	return allowedObservabilityPipelineHttpClientDestinationCompressionAlgorithmEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineHttpClientDestinationCompressionAlgorithm) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineHttpClientDestinationCompressionAlgorithm(value)
	return nil
}

// NewObservabilityPipelineHttpClientDestinationCompressionAlgorithmFromValue returns a pointer to a valid ObservabilityPipelineHttpClientDestinationCompressionAlgorithm
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineHttpClientDestinationCompressionAlgorithmFromValue(v string) (*ObservabilityPipelineHttpClientDestinationCompressionAlgorithm, error) {
	ev := ObservabilityPipelineHttpClientDestinationCompressionAlgorithm(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineHttpClientDestinationCompressionAlgorithm: valid values are %v", v, allowedObservabilityPipelineHttpClientDestinationCompressionAlgorithmEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineHttpClientDestinationCompressionAlgorithm) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineHttpClientDestinationCompressionAlgorithmEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineHttpClientDestinationCompressionAlgorithm value.
func (v ObservabilityPipelineHttpClientDestinationCompressionAlgorithm) Ptr() *ObservabilityPipelineHttpClientDestinationCompressionAlgorithm {
	return &v
}
