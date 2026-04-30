// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineElasticsearchDestinationCompressionAlgorithm The compression algorithm applied when sending data to Elasticsearch.
type ObservabilityPipelineElasticsearchDestinationCompressionAlgorithm string

// List of ObservabilityPipelineElasticsearchDestinationCompressionAlgorithm.
const (
	OBSERVABILITYPIPELINEELASTICSEARCHDESTINATIONCOMPRESSIONALGORITHM_NONE   ObservabilityPipelineElasticsearchDestinationCompressionAlgorithm = "none"
	OBSERVABILITYPIPELINEELASTICSEARCHDESTINATIONCOMPRESSIONALGORITHM_GZIP   ObservabilityPipelineElasticsearchDestinationCompressionAlgorithm = "gzip"
	OBSERVABILITYPIPELINEELASTICSEARCHDESTINATIONCOMPRESSIONALGORITHM_ZLIB   ObservabilityPipelineElasticsearchDestinationCompressionAlgorithm = "zlib"
	OBSERVABILITYPIPELINEELASTICSEARCHDESTINATIONCOMPRESSIONALGORITHM_ZSTD   ObservabilityPipelineElasticsearchDestinationCompressionAlgorithm = "zstd"
	OBSERVABILITYPIPELINEELASTICSEARCHDESTINATIONCOMPRESSIONALGORITHM_SNAPPY ObservabilityPipelineElasticsearchDestinationCompressionAlgorithm = "snappy"
)

var allowedObservabilityPipelineElasticsearchDestinationCompressionAlgorithmEnumValues = []ObservabilityPipelineElasticsearchDestinationCompressionAlgorithm{
	OBSERVABILITYPIPELINEELASTICSEARCHDESTINATIONCOMPRESSIONALGORITHM_NONE,
	OBSERVABILITYPIPELINEELASTICSEARCHDESTINATIONCOMPRESSIONALGORITHM_GZIP,
	OBSERVABILITYPIPELINEELASTICSEARCHDESTINATIONCOMPRESSIONALGORITHM_ZLIB,
	OBSERVABILITYPIPELINEELASTICSEARCHDESTINATIONCOMPRESSIONALGORITHM_ZSTD,
	OBSERVABILITYPIPELINEELASTICSEARCHDESTINATIONCOMPRESSIONALGORITHM_SNAPPY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineElasticsearchDestinationCompressionAlgorithm) GetAllowedValues() []ObservabilityPipelineElasticsearchDestinationCompressionAlgorithm {
	return allowedObservabilityPipelineElasticsearchDestinationCompressionAlgorithmEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineElasticsearchDestinationCompressionAlgorithm) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineElasticsearchDestinationCompressionAlgorithm(value)
	return nil
}

// NewObservabilityPipelineElasticsearchDestinationCompressionAlgorithmFromValue returns a pointer to a valid ObservabilityPipelineElasticsearchDestinationCompressionAlgorithm
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineElasticsearchDestinationCompressionAlgorithmFromValue(v string) (*ObservabilityPipelineElasticsearchDestinationCompressionAlgorithm, error) {
	ev := ObservabilityPipelineElasticsearchDestinationCompressionAlgorithm(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineElasticsearchDestinationCompressionAlgorithm: valid values are %v", v, allowedObservabilityPipelineElasticsearchDestinationCompressionAlgorithmEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineElasticsearchDestinationCompressionAlgorithm) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineElasticsearchDestinationCompressionAlgorithmEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineElasticsearchDestinationCompressionAlgorithm value.
func (v ObservabilityPipelineElasticsearchDestinationCompressionAlgorithm) Ptr() *ObservabilityPipelineElasticsearchDestinationCompressionAlgorithm {
	return &v
}
