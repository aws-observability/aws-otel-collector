// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineHttpServerSourceAuthStrategy HTTP authentication method.
type ObservabilityPipelineHttpServerSourceAuthStrategy string

// List of ObservabilityPipelineHttpServerSourceAuthStrategy.
const (
	OBSERVABILITYPIPELINEHTTPSERVERSOURCEAUTHSTRATEGY_NONE  ObservabilityPipelineHttpServerSourceAuthStrategy = "none"
	OBSERVABILITYPIPELINEHTTPSERVERSOURCEAUTHSTRATEGY_PLAIN ObservabilityPipelineHttpServerSourceAuthStrategy = "plain"
)

var allowedObservabilityPipelineHttpServerSourceAuthStrategyEnumValues = []ObservabilityPipelineHttpServerSourceAuthStrategy{
	OBSERVABILITYPIPELINEHTTPSERVERSOURCEAUTHSTRATEGY_NONE,
	OBSERVABILITYPIPELINEHTTPSERVERSOURCEAUTHSTRATEGY_PLAIN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineHttpServerSourceAuthStrategy) GetAllowedValues() []ObservabilityPipelineHttpServerSourceAuthStrategy {
	return allowedObservabilityPipelineHttpServerSourceAuthStrategyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineHttpServerSourceAuthStrategy) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineHttpServerSourceAuthStrategy(value)
	return nil
}

// NewObservabilityPipelineHttpServerSourceAuthStrategyFromValue returns a pointer to a valid ObservabilityPipelineHttpServerSourceAuthStrategy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineHttpServerSourceAuthStrategyFromValue(v string) (*ObservabilityPipelineHttpServerSourceAuthStrategy, error) {
	ev := ObservabilityPipelineHttpServerSourceAuthStrategy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineHttpServerSourceAuthStrategy: valid values are %v", v, allowedObservabilityPipelineHttpServerSourceAuthStrategyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineHttpServerSourceAuthStrategy) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineHttpServerSourceAuthStrategyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineHttpServerSourceAuthStrategy value.
func (v ObservabilityPipelineHttpServerSourceAuthStrategy) Ptr() *ObservabilityPipelineHttpServerSourceAuthStrategy {
	return &v
}
