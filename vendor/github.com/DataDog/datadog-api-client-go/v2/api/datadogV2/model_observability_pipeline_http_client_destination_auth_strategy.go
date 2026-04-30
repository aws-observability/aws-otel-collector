// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineHttpClientDestinationAuthStrategy HTTP authentication strategy.
type ObservabilityPipelineHttpClientDestinationAuthStrategy string

// List of ObservabilityPipelineHttpClientDestinationAuthStrategy.
const (
	OBSERVABILITYPIPELINEHTTPCLIENTDESTINATIONAUTHSTRATEGY_NONE   ObservabilityPipelineHttpClientDestinationAuthStrategy = "none"
	OBSERVABILITYPIPELINEHTTPCLIENTDESTINATIONAUTHSTRATEGY_BASIC  ObservabilityPipelineHttpClientDestinationAuthStrategy = "basic"
	OBSERVABILITYPIPELINEHTTPCLIENTDESTINATIONAUTHSTRATEGY_BEARER ObservabilityPipelineHttpClientDestinationAuthStrategy = "bearer"
)

var allowedObservabilityPipelineHttpClientDestinationAuthStrategyEnumValues = []ObservabilityPipelineHttpClientDestinationAuthStrategy{
	OBSERVABILITYPIPELINEHTTPCLIENTDESTINATIONAUTHSTRATEGY_NONE,
	OBSERVABILITYPIPELINEHTTPCLIENTDESTINATIONAUTHSTRATEGY_BASIC,
	OBSERVABILITYPIPELINEHTTPCLIENTDESTINATIONAUTHSTRATEGY_BEARER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineHttpClientDestinationAuthStrategy) GetAllowedValues() []ObservabilityPipelineHttpClientDestinationAuthStrategy {
	return allowedObservabilityPipelineHttpClientDestinationAuthStrategyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineHttpClientDestinationAuthStrategy) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineHttpClientDestinationAuthStrategy(value)
	return nil
}

// NewObservabilityPipelineHttpClientDestinationAuthStrategyFromValue returns a pointer to a valid ObservabilityPipelineHttpClientDestinationAuthStrategy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineHttpClientDestinationAuthStrategyFromValue(v string) (*ObservabilityPipelineHttpClientDestinationAuthStrategy, error) {
	ev := ObservabilityPipelineHttpClientDestinationAuthStrategy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineHttpClientDestinationAuthStrategy: valid values are %v", v, allowedObservabilityPipelineHttpClientDestinationAuthStrategyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineHttpClientDestinationAuthStrategy) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineHttpClientDestinationAuthStrategyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineHttpClientDestinationAuthStrategy value.
func (v ObservabilityPipelineHttpClientDestinationAuthStrategy) Ptr() *ObservabilityPipelineHttpClientDestinationAuthStrategy {
	return &v
}
