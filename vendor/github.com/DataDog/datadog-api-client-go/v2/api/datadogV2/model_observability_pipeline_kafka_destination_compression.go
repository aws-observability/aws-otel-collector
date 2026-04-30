// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineKafkaDestinationCompression Compression codec for Kafka messages.
type ObservabilityPipelineKafkaDestinationCompression string

// List of ObservabilityPipelineKafkaDestinationCompression.
const (
	OBSERVABILITYPIPELINEKAFKADESTINATIONCOMPRESSION_NONE   ObservabilityPipelineKafkaDestinationCompression = "none"
	OBSERVABILITYPIPELINEKAFKADESTINATIONCOMPRESSION_GZIP   ObservabilityPipelineKafkaDestinationCompression = "gzip"
	OBSERVABILITYPIPELINEKAFKADESTINATIONCOMPRESSION_SNAPPY ObservabilityPipelineKafkaDestinationCompression = "snappy"
	OBSERVABILITYPIPELINEKAFKADESTINATIONCOMPRESSION_LZ4    ObservabilityPipelineKafkaDestinationCompression = "lz4"
	OBSERVABILITYPIPELINEKAFKADESTINATIONCOMPRESSION_ZSTD   ObservabilityPipelineKafkaDestinationCompression = "zstd"
)

var allowedObservabilityPipelineKafkaDestinationCompressionEnumValues = []ObservabilityPipelineKafkaDestinationCompression{
	OBSERVABILITYPIPELINEKAFKADESTINATIONCOMPRESSION_NONE,
	OBSERVABILITYPIPELINEKAFKADESTINATIONCOMPRESSION_GZIP,
	OBSERVABILITYPIPELINEKAFKADESTINATIONCOMPRESSION_SNAPPY,
	OBSERVABILITYPIPELINEKAFKADESTINATIONCOMPRESSION_LZ4,
	OBSERVABILITYPIPELINEKAFKADESTINATIONCOMPRESSION_ZSTD,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineKafkaDestinationCompression) GetAllowedValues() []ObservabilityPipelineKafkaDestinationCompression {
	return allowedObservabilityPipelineKafkaDestinationCompressionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineKafkaDestinationCompression) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineKafkaDestinationCompression(value)
	return nil
}

// NewObservabilityPipelineKafkaDestinationCompressionFromValue returns a pointer to a valid ObservabilityPipelineKafkaDestinationCompression
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineKafkaDestinationCompressionFromValue(v string) (*ObservabilityPipelineKafkaDestinationCompression, error) {
	ev := ObservabilityPipelineKafkaDestinationCompression(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineKafkaDestinationCompression: valid values are %v", v, allowedObservabilityPipelineKafkaDestinationCompressionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineKafkaDestinationCompression) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineKafkaDestinationCompressionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineKafkaDestinationCompression value.
func (v ObservabilityPipelineKafkaDestinationCompression) Ptr() *ObservabilityPipelineKafkaDestinationCompression {
	return &v
}
