// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAmazonS3GenericCompressionZstdType The compression type. Always `zstd`.
type ObservabilityPipelineAmazonS3GenericCompressionZstdType string

// List of ObservabilityPipelineAmazonS3GenericCompressionZstdType.
const (
	OBSERVABILITYPIPELINEAMAZONS3GENERICCOMPRESSIONZSTDTYPE_ZSTD ObservabilityPipelineAmazonS3GenericCompressionZstdType = "zstd"
)

var allowedObservabilityPipelineAmazonS3GenericCompressionZstdTypeEnumValues = []ObservabilityPipelineAmazonS3GenericCompressionZstdType{
	OBSERVABILITYPIPELINEAMAZONS3GENERICCOMPRESSIONZSTDTYPE_ZSTD,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineAmazonS3GenericCompressionZstdType) GetAllowedValues() []ObservabilityPipelineAmazonS3GenericCompressionZstdType {
	return allowedObservabilityPipelineAmazonS3GenericCompressionZstdTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineAmazonS3GenericCompressionZstdType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineAmazonS3GenericCompressionZstdType(value)
	return nil
}

// NewObservabilityPipelineAmazonS3GenericCompressionZstdTypeFromValue returns a pointer to a valid ObservabilityPipelineAmazonS3GenericCompressionZstdType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineAmazonS3GenericCompressionZstdTypeFromValue(v string) (*ObservabilityPipelineAmazonS3GenericCompressionZstdType, error) {
	ev := ObservabilityPipelineAmazonS3GenericCompressionZstdType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineAmazonS3GenericCompressionZstdType: valid values are %v", v, allowedObservabilityPipelineAmazonS3GenericCompressionZstdTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineAmazonS3GenericCompressionZstdType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineAmazonS3GenericCompressionZstdTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineAmazonS3GenericCompressionZstdType value.
func (v ObservabilityPipelineAmazonS3GenericCompressionZstdType) Ptr() *ObservabilityPipelineAmazonS3GenericCompressionZstdType {
	return &v
}
