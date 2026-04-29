// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAmazonS3SourceCompression Compression format for objects retrieved from the S3 bucket. Use `auto` to detect compression from the object's Content-Encoding header or file extension.
type ObservabilityPipelineAmazonS3SourceCompression string

// List of ObservabilityPipelineAmazonS3SourceCompression.
const (
	OBSERVABILITYPIPELINEAMAZONS3SOURCECOMPRESSION_AUTO ObservabilityPipelineAmazonS3SourceCompression = "auto"
	OBSERVABILITYPIPELINEAMAZONS3SOURCECOMPRESSION_NONE ObservabilityPipelineAmazonS3SourceCompression = "none"
	OBSERVABILITYPIPELINEAMAZONS3SOURCECOMPRESSION_GZIP ObservabilityPipelineAmazonS3SourceCompression = "gzip"
	OBSERVABILITYPIPELINEAMAZONS3SOURCECOMPRESSION_ZSTD ObservabilityPipelineAmazonS3SourceCompression = "zstd"
)

var allowedObservabilityPipelineAmazonS3SourceCompressionEnumValues = []ObservabilityPipelineAmazonS3SourceCompression{
	OBSERVABILITYPIPELINEAMAZONS3SOURCECOMPRESSION_AUTO,
	OBSERVABILITYPIPELINEAMAZONS3SOURCECOMPRESSION_NONE,
	OBSERVABILITYPIPELINEAMAZONS3SOURCECOMPRESSION_GZIP,
	OBSERVABILITYPIPELINEAMAZONS3SOURCECOMPRESSION_ZSTD,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineAmazonS3SourceCompression) GetAllowedValues() []ObservabilityPipelineAmazonS3SourceCompression {
	return allowedObservabilityPipelineAmazonS3SourceCompressionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineAmazonS3SourceCompression) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineAmazonS3SourceCompression(value)
	return nil
}

// NewObservabilityPipelineAmazonS3SourceCompressionFromValue returns a pointer to a valid ObservabilityPipelineAmazonS3SourceCompression
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineAmazonS3SourceCompressionFromValue(v string) (*ObservabilityPipelineAmazonS3SourceCompression, error) {
	ev := ObservabilityPipelineAmazonS3SourceCompression(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineAmazonS3SourceCompression: valid values are %v", v, allowedObservabilityPipelineAmazonS3SourceCompressionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineAmazonS3SourceCompression) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineAmazonS3SourceCompressionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineAmazonS3SourceCompression value.
func (v ObservabilityPipelineAmazonS3SourceCompression) Ptr() *ObservabilityPipelineAmazonS3SourceCompression {
	return &v
}
