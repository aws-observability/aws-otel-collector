// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAmazonS3GenericCompressionGzipType The compression type. Always `gzip`.
type ObservabilityPipelineAmazonS3GenericCompressionGzipType string

// List of ObservabilityPipelineAmazonS3GenericCompressionGzipType.
const (
	OBSERVABILITYPIPELINEAMAZONS3GENERICCOMPRESSIONGZIPTYPE_GZIP ObservabilityPipelineAmazonS3GenericCompressionGzipType = "gzip"
)

var allowedObservabilityPipelineAmazonS3GenericCompressionGzipTypeEnumValues = []ObservabilityPipelineAmazonS3GenericCompressionGzipType{
	OBSERVABILITYPIPELINEAMAZONS3GENERICCOMPRESSIONGZIPTYPE_GZIP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineAmazonS3GenericCompressionGzipType) GetAllowedValues() []ObservabilityPipelineAmazonS3GenericCompressionGzipType {
	return allowedObservabilityPipelineAmazonS3GenericCompressionGzipTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineAmazonS3GenericCompressionGzipType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineAmazonS3GenericCompressionGzipType(value)
	return nil
}

// NewObservabilityPipelineAmazonS3GenericCompressionGzipTypeFromValue returns a pointer to a valid ObservabilityPipelineAmazonS3GenericCompressionGzipType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineAmazonS3GenericCompressionGzipTypeFromValue(v string) (*ObservabilityPipelineAmazonS3GenericCompressionGzipType, error) {
	ev := ObservabilityPipelineAmazonS3GenericCompressionGzipType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineAmazonS3GenericCompressionGzipType: valid values are %v", v, allowedObservabilityPipelineAmazonS3GenericCompressionGzipTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineAmazonS3GenericCompressionGzipType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineAmazonS3GenericCompressionGzipTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineAmazonS3GenericCompressionGzipType value.
func (v ObservabilityPipelineAmazonS3GenericCompressionGzipType) Ptr() *ObservabilityPipelineAmazonS3GenericCompressionGzipType {
	return &v
}
