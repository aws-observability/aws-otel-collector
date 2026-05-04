// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineBufferOptionsWhenFull Behavior when the buffer is full (block and stop accepting new events, or drop new events)
type ObservabilityPipelineBufferOptionsWhenFull string

// List of ObservabilityPipelineBufferOptionsWhenFull.
const (
	OBSERVABILITYPIPELINEBUFFEROPTIONSWHENFULL_BLOCK       ObservabilityPipelineBufferOptionsWhenFull = "block"
	OBSERVABILITYPIPELINEBUFFEROPTIONSWHENFULL_DROP_NEWEST ObservabilityPipelineBufferOptionsWhenFull = "drop_newest"
)

var allowedObservabilityPipelineBufferOptionsWhenFullEnumValues = []ObservabilityPipelineBufferOptionsWhenFull{
	OBSERVABILITYPIPELINEBUFFEROPTIONSWHENFULL_BLOCK,
	OBSERVABILITYPIPELINEBUFFEROPTIONSWHENFULL_DROP_NEWEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineBufferOptionsWhenFull) GetAllowedValues() []ObservabilityPipelineBufferOptionsWhenFull {
	return allowedObservabilityPipelineBufferOptionsWhenFullEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineBufferOptionsWhenFull) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineBufferOptionsWhenFull(value)
	return nil
}

// NewObservabilityPipelineBufferOptionsWhenFullFromValue returns a pointer to a valid ObservabilityPipelineBufferOptionsWhenFull
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineBufferOptionsWhenFullFromValue(v string) (*ObservabilityPipelineBufferOptionsWhenFull, error) {
	ev := ObservabilityPipelineBufferOptionsWhenFull(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineBufferOptionsWhenFull: valid values are %v", v, allowedObservabilityPipelineBufferOptionsWhenFullEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineBufferOptionsWhenFull) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineBufferOptionsWhenFullEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineBufferOptionsWhenFull value.
func (v ObservabilityPipelineBufferOptionsWhenFull) Ptr() *ObservabilityPipelineBufferOptionsWhenFull {
	return &v
}
