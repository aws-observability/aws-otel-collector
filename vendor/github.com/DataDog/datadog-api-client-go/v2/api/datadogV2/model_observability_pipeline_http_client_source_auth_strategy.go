// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineHttpClientSourceAuthStrategy Optional authentication strategy for HTTP requests.
type ObservabilityPipelineHttpClientSourceAuthStrategy string

// List of ObservabilityPipelineHttpClientSourceAuthStrategy.
const (
	OBSERVABILITYPIPELINEHTTPCLIENTSOURCEAUTHSTRATEGY_BASIC  ObservabilityPipelineHttpClientSourceAuthStrategy = "basic"
	OBSERVABILITYPIPELINEHTTPCLIENTSOURCEAUTHSTRATEGY_BEARER ObservabilityPipelineHttpClientSourceAuthStrategy = "bearer"
)

var allowedObservabilityPipelineHttpClientSourceAuthStrategyEnumValues = []ObservabilityPipelineHttpClientSourceAuthStrategy{
	OBSERVABILITYPIPELINEHTTPCLIENTSOURCEAUTHSTRATEGY_BASIC,
	OBSERVABILITYPIPELINEHTTPCLIENTSOURCEAUTHSTRATEGY_BEARER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineHttpClientSourceAuthStrategy) GetAllowedValues() []ObservabilityPipelineHttpClientSourceAuthStrategy {
	return allowedObservabilityPipelineHttpClientSourceAuthStrategyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineHttpClientSourceAuthStrategy) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineHttpClientSourceAuthStrategy(value)
	return nil
}

// NewObservabilityPipelineHttpClientSourceAuthStrategyFromValue returns a pointer to a valid ObservabilityPipelineHttpClientSourceAuthStrategy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineHttpClientSourceAuthStrategyFromValue(v string) (*ObservabilityPipelineHttpClientSourceAuthStrategy, error) {
	ev := ObservabilityPipelineHttpClientSourceAuthStrategy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineHttpClientSourceAuthStrategy: valid values are %v", v, allowedObservabilityPipelineHttpClientSourceAuthStrategyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineHttpClientSourceAuthStrategy) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineHttpClientSourceAuthStrategyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineHttpClientSourceAuthStrategy value.
func (v ObservabilityPipelineHttpClientSourceAuthStrategy) Ptr() *ObservabilityPipelineHttpClientSourceAuthStrategy {
	return &v
}
