// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAmazonS3DestinationType The destination type. Always `amazon_s3`.
type ObservabilityPipelineAmazonS3DestinationType string

// List of ObservabilityPipelineAmazonS3DestinationType.
const (
	OBSERVABILITYPIPELINEAMAZONS3DESTINATIONTYPE_AMAZON_S3 ObservabilityPipelineAmazonS3DestinationType = "amazon_s3"
)

var allowedObservabilityPipelineAmazonS3DestinationTypeEnumValues = []ObservabilityPipelineAmazonS3DestinationType{
	OBSERVABILITYPIPELINEAMAZONS3DESTINATIONTYPE_AMAZON_S3,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineAmazonS3DestinationType) GetAllowedValues() []ObservabilityPipelineAmazonS3DestinationType {
	return allowedObservabilityPipelineAmazonS3DestinationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineAmazonS3DestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineAmazonS3DestinationType(value)
	return nil
}

// NewObservabilityPipelineAmazonS3DestinationTypeFromValue returns a pointer to a valid ObservabilityPipelineAmazonS3DestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineAmazonS3DestinationTypeFromValue(v string) (*ObservabilityPipelineAmazonS3DestinationType, error) {
	ev := ObservabilityPipelineAmazonS3DestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineAmazonS3DestinationType: valid values are %v", v, allowedObservabilityPipelineAmazonS3DestinationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineAmazonS3DestinationType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineAmazonS3DestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineAmazonS3DestinationType value.
func (v ObservabilityPipelineAmazonS3DestinationType) Ptr() *ObservabilityPipelineAmazonS3DestinationType {
	return &v
}
