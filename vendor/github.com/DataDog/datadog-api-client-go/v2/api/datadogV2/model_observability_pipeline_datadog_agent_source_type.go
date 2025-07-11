// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineDatadogAgentSourceType The source type. The value should always be `datadog_agent`.
type ObservabilityPipelineDatadogAgentSourceType string

// List of ObservabilityPipelineDatadogAgentSourceType.
const (
	OBSERVABILITYPIPELINEDATADOGAGENTSOURCETYPE_DATADOG_AGENT ObservabilityPipelineDatadogAgentSourceType = "datadog_agent"
)

var allowedObservabilityPipelineDatadogAgentSourceTypeEnumValues = []ObservabilityPipelineDatadogAgentSourceType{
	OBSERVABILITYPIPELINEDATADOGAGENTSOURCETYPE_DATADOG_AGENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ObservabilityPipelineDatadogAgentSourceType) GetAllowedValues() []ObservabilityPipelineDatadogAgentSourceType {
	return allowedObservabilityPipelineDatadogAgentSourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ObservabilityPipelineDatadogAgentSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ObservabilityPipelineDatadogAgentSourceType(value)
	return nil
}

// NewObservabilityPipelineDatadogAgentSourceTypeFromValue returns a pointer to a valid ObservabilityPipelineDatadogAgentSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewObservabilityPipelineDatadogAgentSourceTypeFromValue(v string) (*ObservabilityPipelineDatadogAgentSourceType, error) {
	ev := ObservabilityPipelineDatadogAgentSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ObservabilityPipelineDatadogAgentSourceType: valid values are %v", v, allowedObservabilityPipelineDatadogAgentSourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ObservabilityPipelineDatadogAgentSourceType) IsValid() bool {
	for _, existing := range allowedObservabilityPipelineDatadogAgentSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObservabilityPipelineDatadogAgentSourceType value.
func (v ObservabilityPipelineDatadogAgentSourceType) Ptr() *ObservabilityPipelineDatadogAgentSourceType {
	return &v
}
