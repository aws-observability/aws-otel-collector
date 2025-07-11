// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAmazonOpenSearchDestinationAuthStrategy The authentication strategy to use.
type ObservabilityPipelineAmazonOpenSearchDestinationAuthStrategy string

// List of ObservabilityPipelineAmazonOpenSearchDestinationAuthStrategy.
const (
	OBSERVABILITYPIPELINEAMAZONOPENSEARCHDESTINATIONAUTHSTRATEGY_BASIC ObservabilityPipelineAmazonOpenSearchDestinationAuthStrategy = "basic"
	OBSERVABILITYPIPELINEAMAZONOPENSEARCHDESTINATIONAUTHSTRATEGY_AWS   ObservabilityPipelineAmazonOpenSearchDestinationAuthStrategy = "aws"
)

var allowedObservabilityPipelineAmazonOpenSearchDestinationAuthStrategyEnumValues = []ObservabilityPipelineAmazonOpenSearchDestinationAuthStrategy{
	OBSERVABILITYPIPELINEAMAZONOPENSEARCHDESTINATIONAUTHSTRATEGY_BASIC,
	OBSERVABILITYPIPELINEAMAZONOPENSEARCHDESTINATIONAUTHSTRATEGY_AWS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineAmazonOpenSearchDestinationAuthStrategy) GetAllowedValues() []ObservabilityPipelineAmazonOpenSearchDestinationAuthStrategy {
	return allowedObservabilityPipelineAmazonOpenSearchDestinationAuthStrategyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineAmazonOpenSearchDestinationAuthStrategy) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineAmazonOpenSearchDestinationAuthStrategy(value)
	return nil
}

// NewObservabilityPipelineAmazonOpenSearchDestinationAuthStrategyFromValue returns a pointer to a valid ObservabilityPipelineAmazonOpenSearchDestinationAuthStrategy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineAmazonOpenSearchDestinationAuthStrategyFromValue(v string) (*ObservabilityPipelineAmazonOpenSearchDestinationAuthStrategy, error) {
	ev := ObservabilityPipelineAmazonOpenSearchDestinationAuthStrategy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineAmazonOpenSearchDestinationAuthStrategy: valid values are %v", v, allowedObservabilityPipelineAmazonOpenSearchDestinationAuthStrategyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineAmazonOpenSearchDestinationAuthStrategy) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineAmazonOpenSearchDestinationAuthStrategyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineAmazonOpenSearchDestinationAuthStrategy value.
func (v ObservabilityPipelineAmazonOpenSearchDestinationAuthStrategy) Ptr() *ObservabilityPipelineAmazonOpenSearchDestinationAuthStrategy {
	return &v
}
