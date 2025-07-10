// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAmazonOpenSearchDestinationType The destination type. The value should always be `amazon_opensearch`.
type ObservabilityPipelineAmazonOpenSearchDestinationType string

// List of ObservabilityPipelineAmazonOpenSearchDestinationType.
const (
	OBSERVABILITYPIPELINEAMAZONOPENSEARCHDESTINATIONTYPE_AMAZON_OPENSEARCH ObservabilityPipelineAmazonOpenSearchDestinationType = "amazon_opensearch"
)

var allowedObservabilityPipelineAmazonOpenSearchDestinationTypeEnumValues = []ObservabilityPipelineAmazonOpenSearchDestinationType{
	OBSERVABILITYPIPELINEAMAZONOPENSEARCHDESTINATIONTYPE_AMAZON_OPENSEARCH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineAmazonOpenSearchDestinationType) GetAllowedValues() []ObservabilityPipelineAmazonOpenSearchDestinationType {
	return allowedObservabilityPipelineAmazonOpenSearchDestinationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineAmazonOpenSearchDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineAmazonOpenSearchDestinationType(value)
	return nil
}

// NewObservabilityPipelineAmazonOpenSearchDestinationTypeFromValue returns a pointer to a valid ObservabilityPipelineAmazonOpenSearchDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineAmazonOpenSearchDestinationTypeFromValue(v string) (*ObservabilityPipelineAmazonOpenSearchDestinationType, error) {
	ev := ObservabilityPipelineAmazonOpenSearchDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineAmazonOpenSearchDestinationType: valid values are %v", v, allowedObservabilityPipelineAmazonOpenSearchDestinationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineAmazonOpenSearchDestinationType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineAmazonOpenSearchDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineAmazonOpenSearchDestinationType value.
func (v ObservabilityPipelineAmazonOpenSearchDestinationType) Ptr() *ObservabilityPipelineAmazonOpenSearchDestinationType {
	return &v
}
