// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineCrowdStrikeNextGenSiemDestinationCompressionAlgorithm Compression algorithm for log events.
type ObservabilityPipelineCrowdStrikeNextGenSiemDestinationCompressionAlgorithm string

// List of ObservabilityPipelineCrowdStrikeNextGenSiemDestinationCompressionAlgorithm.
const (
	OBSERVABILITYPIPELINECROWDSTRIKENEXTGENSIEMDESTINATIONCOMPRESSIONALGORITHM_GZIP ObservabilityPipelineCrowdStrikeNextGenSiemDestinationCompressionAlgorithm = "gzip"
	OBSERVABILITYPIPELINECROWDSTRIKENEXTGENSIEMDESTINATIONCOMPRESSIONALGORITHM_ZLIB ObservabilityPipelineCrowdStrikeNextGenSiemDestinationCompressionAlgorithm = "zlib"
)

var allowedObservabilityPipelineCrowdStrikeNextGenSiemDestinationCompressionAlgorithmEnumValues = []ObservabilityPipelineCrowdStrikeNextGenSiemDestinationCompressionAlgorithm{
	OBSERVABILITYPIPELINECROWDSTRIKENEXTGENSIEMDESTINATIONCOMPRESSIONALGORITHM_GZIP,
	OBSERVABILITYPIPELINECROWDSTRIKENEXTGENSIEMDESTINATIONCOMPRESSIONALGORITHM_ZLIB,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineCrowdStrikeNextGenSiemDestinationCompressionAlgorithm) GetAllowedValues() []ObservabilityPipelineCrowdStrikeNextGenSiemDestinationCompressionAlgorithm {
	return allowedObservabilityPipelineCrowdStrikeNextGenSiemDestinationCompressionAlgorithmEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineCrowdStrikeNextGenSiemDestinationCompressionAlgorithm) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineCrowdStrikeNextGenSiemDestinationCompressionAlgorithm(value)
	return nil
}

// NewObservabilityPipelineCrowdStrikeNextGenSiemDestinationCompressionAlgorithmFromValue returns a pointer to a valid ObservabilityPipelineCrowdStrikeNextGenSiemDestinationCompressionAlgorithm
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineCrowdStrikeNextGenSiemDestinationCompressionAlgorithmFromValue(v string) (*ObservabilityPipelineCrowdStrikeNextGenSiemDestinationCompressionAlgorithm, error) {
	ev := ObservabilityPipelineCrowdStrikeNextGenSiemDestinationCompressionAlgorithm(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineCrowdStrikeNextGenSiemDestinationCompressionAlgorithm: valid values are %v", v, allowedObservabilityPipelineCrowdStrikeNextGenSiemDestinationCompressionAlgorithmEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineCrowdStrikeNextGenSiemDestinationCompressionAlgorithm) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineCrowdStrikeNextGenSiemDestinationCompressionAlgorithmEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineCrowdStrikeNextGenSiemDestinationCompressionAlgorithm value.
func (v ObservabilityPipelineCrowdStrikeNextGenSiemDestinationCompressionAlgorithm) Ptr() *ObservabilityPipelineCrowdStrikeNextGenSiemDestinationCompressionAlgorithm {
	return &v
}
