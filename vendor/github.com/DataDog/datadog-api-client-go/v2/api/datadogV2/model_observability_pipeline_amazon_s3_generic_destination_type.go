// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAmazonS3GenericDestinationType The destination type. Always `amazon_s3_generic`.
type ObservabilityPipelineAmazonS3GenericDestinationType string

// List of ObservabilityPipelineAmazonS3GenericDestinationType.
const (
	OBSERVABILITYPIPELINEAMAZONS3GENERICDESTINATIONTYPE_GENERIC_ARCHIVES_S3 ObservabilityPipelineAmazonS3GenericDestinationType = "amazon_s3_generic"
)

var allowedObservabilityPipelineAmazonS3GenericDestinationTypeEnumValues = []ObservabilityPipelineAmazonS3GenericDestinationType{
	OBSERVABILITYPIPELINEAMAZONS3GENERICDESTINATIONTYPE_GENERIC_ARCHIVES_S3,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineAmazonS3GenericDestinationType) GetAllowedValues() []ObservabilityPipelineAmazonS3GenericDestinationType {
	return allowedObservabilityPipelineAmazonS3GenericDestinationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineAmazonS3GenericDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineAmazonS3GenericDestinationType(value)
	return nil
}

// NewObservabilityPipelineAmazonS3GenericDestinationTypeFromValue returns a pointer to a valid ObservabilityPipelineAmazonS3GenericDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineAmazonS3GenericDestinationTypeFromValue(v string) (*ObservabilityPipelineAmazonS3GenericDestinationType, error) {
	ev := ObservabilityPipelineAmazonS3GenericDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineAmazonS3GenericDestinationType: valid values are %v", v, allowedObservabilityPipelineAmazonS3GenericDestinationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineAmazonS3GenericDestinationType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineAmazonS3GenericDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineAmazonS3GenericDestinationType value.
func (v ObservabilityPipelineAmazonS3GenericDestinationType) Ptr() *ObservabilityPipelineAmazonS3GenericDestinationType {
	return &v
}
