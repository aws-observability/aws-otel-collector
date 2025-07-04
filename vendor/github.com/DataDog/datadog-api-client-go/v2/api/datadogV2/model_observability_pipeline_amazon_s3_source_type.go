// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAmazonS3SourceType The source type. Always `amazon_s3`.
type ObservabilityPipelineAmazonS3SourceType string

// List of ObservabilityPipelineAmazonS3SourceType.
const (
	OBSERVABILITYPIPELINEAMAZONS3SOURCETYPE_AMAZON_S3 ObservabilityPipelineAmazonS3SourceType = "amazon_s3"
)

var allowedObservabilityPipelineAmazonS3SourceTypeEnumValues = []ObservabilityPipelineAmazonS3SourceType{
	OBSERVABILITYPIPELINEAMAZONS3SOURCETYPE_AMAZON_S3,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineAmazonS3SourceType) GetAllowedValues() []ObservabilityPipelineAmazonS3SourceType {
	return allowedObservabilityPipelineAmazonS3SourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineAmazonS3SourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineAmazonS3SourceType(value)
	return nil
}

// NewObservabilityPipelineAmazonS3SourceTypeFromValue returns a pointer to a valid ObservabilityPipelineAmazonS3SourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineAmazonS3SourceTypeFromValue(v string) (*ObservabilityPipelineAmazonS3SourceType, error) {
	ev := ObservabilityPipelineAmazonS3SourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineAmazonS3SourceType: valid values are %v", v, allowedObservabilityPipelineAmazonS3SourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineAmazonS3SourceType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineAmazonS3SourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineAmazonS3SourceType value.
func (v ObservabilityPipelineAmazonS3SourceType) Ptr() *ObservabilityPipelineAmazonS3SourceType {
	return &v
}
