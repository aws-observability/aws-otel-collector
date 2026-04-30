// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAmazonS3GenericCompressionSnappyType The compression type. Always `snappy`.
type ObservabilityPipelineAmazonS3GenericCompressionSnappyType string

// List of ObservabilityPipelineAmazonS3GenericCompressionSnappyType.
const (
	OBSERVABILITYPIPELINEAMAZONS3GENERICCOMPRESSIONSNAPPYTYPE_SNAPPY ObservabilityPipelineAmazonS3GenericCompressionSnappyType = "snappy"
)

var allowedObservabilityPipelineAmazonS3GenericCompressionSnappyTypeEnumValues = []ObservabilityPipelineAmazonS3GenericCompressionSnappyType{
	OBSERVABILITYPIPELINEAMAZONS3GENERICCOMPRESSIONSNAPPYTYPE_SNAPPY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineAmazonS3GenericCompressionSnappyType) GetAllowedValues() []ObservabilityPipelineAmazonS3GenericCompressionSnappyType {
	return allowedObservabilityPipelineAmazonS3GenericCompressionSnappyTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineAmazonS3GenericCompressionSnappyType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineAmazonS3GenericCompressionSnappyType(value)
	return nil
}

// NewObservabilityPipelineAmazonS3GenericCompressionSnappyTypeFromValue returns a pointer to a valid ObservabilityPipelineAmazonS3GenericCompressionSnappyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineAmazonS3GenericCompressionSnappyTypeFromValue(v string) (*ObservabilityPipelineAmazonS3GenericCompressionSnappyType, error) {
	ev := ObservabilityPipelineAmazonS3GenericCompressionSnappyType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineAmazonS3GenericCompressionSnappyType: valid values are %v", v, allowedObservabilityPipelineAmazonS3GenericCompressionSnappyTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineAmazonS3GenericCompressionSnappyType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineAmazonS3GenericCompressionSnappyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineAmazonS3GenericCompressionSnappyType value.
func (v ObservabilityPipelineAmazonS3GenericCompressionSnappyType) Ptr() *ObservabilityPipelineAmazonS3GenericCompressionSnappyType {
	return &v
}
