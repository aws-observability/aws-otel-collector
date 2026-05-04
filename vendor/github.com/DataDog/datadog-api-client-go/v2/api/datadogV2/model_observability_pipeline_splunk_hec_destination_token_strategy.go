// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSplunkHecDestinationTokenStrategy Controls how the Splunk HEC token is supplied. Use `custom` to provide a token with `token_key`, or `from_source` to forward the token received from an upstream Splunk HEC source.
type ObservabilityPipelineSplunkHecDestinationTokenStrategy string

// List of ObservabilityPipelineSplunkHecDestinationTokenStrategy.
const (
	OBSERVABILITYPIPELINESPLUNKHECDESTINATIONTOKENSTRATEGY_CUSTOM      ObservabilityPipelineSplunkHecDestinationTokenStrategy = "custom"
	OBSERVABILITYPIPELINESPLUNKHECDESTINATIONTOKENSTRATEGY_FROM_SOURCE ObservabilityPipelineSplunkHecDestinationTokenStrategy = "from_source"
)

var allowedObservabilityPipelineSplunkHecDestinationTokenStrategyEnumValues = []ObservabilityPipelineSplunkHecDestinationTokenStrategy{
	OBSERVABILITYPIPELINESPLUNKHECDESTINATIONTOKENSTRATEGY_CUSTOM,
	OBSERVABILITYPIPELINESPLUNKHECDESTINATIONTOKENSTRATEGY_FROM_SOURCE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineSplunkHecDestinationTokenStrategy) GetAllowedValues() []ObservabilityPipelineSplunkHecDestinationTokenStrategy {
	return allowedObservabilityPipelineSplunkHecDestinationTokenStrategyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineSplunkHecDestinationTokenStrategy) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineSplunkHecDestinationTokenStrategy(value)
	return nil
}

// NewObservabilityPipelineSplunkHecDestinationTokenStrategyFromValue returns a pointer to a valid ObservabilityPipelineSplunkHecDestinationTokenStrategy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineSplunkHecDestinationTokenStrategyFromValue(v string) (*ObservabilityPipelineSplunkHecDestinationTokenStrategy, error) {
	ev := ObservabilityPipelineSplunkHecDestinationTokenStrategy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineSplunkHecDestinationTokenStrategy: valid values are %v", v, allowedObservabilityPipelineSplunkHecDestinationTokenStrategyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineSplunkHecDestinationTokenStrategy) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineSplunkHecDestinationTokenStrategyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineSplunkHecDestinationTokenStrategy value.
func (v ObservabilityPipelineSplunkHecDestinationTokenStrategy) Ptr() *ObservabilityPipelineSplunkHecDestinationTokenStrategy {
	return &v
}
